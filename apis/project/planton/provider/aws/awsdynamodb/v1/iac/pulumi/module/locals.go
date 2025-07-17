package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Locals gathers frequently-used, derived values that are shared between
// sub-resources during the table provisioning process.
//
//   • Target – a convenient reference to the API-level AwsDynamodb object that
//     is being materialised.
//   • Tags   – a merged map of tags coming from the user-supplied spec plus
//     system tags injected by Project Planton.
//
// Additional helper fields (KmsKeyArn, {G,L}siNames, Labels …) are exposed so
// that sibling helpers can easily exchange data without complex parameter
// lists.
//
// NOTE: keep fields exported so that helpers in other files can freely access
//       them.
//
//nolint:revive // exported fields are intentional here.
type Locals struct {
    // Target is the strongly-typed representation of the desired DynamoDB table
    // coming from the stack input.
    Target *awsdynamodbv1.AwsDynamodb

    // Tags is the final set of key/value tags that will be attached to every
    // AWS resource created by this stack (the table itself, KMS key, …).
    Tags map[string]string

    // Labels is an alias for Tags. Some helpers historically refer to the
    // merged tag set as “labels”.  Keeping both prevents widespread refactors.
    Labels map[string]string

    // KmsKeyArn stores the ARN of the CMK created/used for SSE-KMS (when
    // applicable).
    KmsKeyArn string

    // GsiNames & LsiNames are populated after table creation so they can be
    // exported by the stack.
    GsiNames []string
    LsiNames []string
}

// initializeLocals prepares a Locals instance from the given stack input.
//
// System-level tags are injected first and then user-supplied tags from the
// spec take precedence (so users can intentionally override defaults).
func initializeLocals(ctx *pulumi.Context, in *awsdynamodbv1.AwsDynamodbStackInput) (*Locals, error) {
    if in == nil {
        return nil, errors.New("initializeLocals: stack input must not be nil")
    }

    l := &Locals{
        Target: in.GetTarget(),
        Tags:   map[string]string{},
    }

    // 1. System / default tags – always present.
    l.Tags["managed-by"] = "project-planton"
    l.Tags["iac-provider"] = "pulumi"

    // 2. User-defined tags – override defaults when keys collide.
    if tgt := in.GetTarget(); tgt != nil {
        if spec := tgt.GetSpec(); spec != nil {
            for k, v := range spec.GetTags() {
                if k == "" {
                    continue
                }
                l.Tags[k] = v
            }
        }
    }

    // Labels currently equal the merged Tags. Exposed via a distinct field so
    // that call-sites that historically used “labels” keep compiling.
    l.Labels = l.Tags

    return l, nil
}
