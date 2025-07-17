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
// Storing those values avoids passing around long parameter lists and guarantees
// that every nested Pulumi resource sees a consistent view of the shared data.
//
// NOTE: add new fields here when additional cross-cutting information becomes
//       necessary.
//
// The struct intentionally keeps exported fields so that sub-resource helpers
// living in other files of the same package can freely access them.
type Locals struct {
    // Target is the strongly-typed representation of the desired DynamoDB table
    // coming from the stack input.
    Target *awsdynamodbv1.AwsDynamodb

    // Tags is the final set of key/value tags that will be attached to every
    // AWS resource created by this stack (the table itself, KMS key, …).
    Tags map[string]string
}

// initializeLocals prepares a Locals instance from the given stack input.
//
// System-level tags are injected first and then user-supplied tags from the
// spec take precedence (so users can intentionally override defaults if they
// wish to).
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
                // Ignore empty keys just in case – validation should already
                // enforce this but we prefer to be defensive.
                if k == "" {
                    continue
                }
                l.Tags[k] = v
            }
        }
    }

    return l, nil
}
