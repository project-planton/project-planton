package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Locals aggregates values that are shared across multiple sub-resources
// (e.g. pre-calculated tags or the fully-typed target API resource). Keeping
// them in one place simplifies the signature of the helper functions that
// create the individual Pulumi resources.
type Locals struct {
    // Target holds the user-supplied API resource that is being provisioned.
    Target *awsdynamodbv1.AwsDynamodb

    // Labels is a generic key/value collection typically used for naming or
    // tagging where a plain Go map[string]string is more convenient than
    // Pulumi’s StringMap.
    Labels map[string]string

    // Tags is a Pulumi-level representation of AWS resource tags that can be
    // passed directly into the AWS provider SDK. Values are already converted
    // to pulumi.String inputs, so no further transformation is required when
    // assigning them to the Tag property of a resource.
    Tags pulumi.StringMap
}

// initializeLocals derives helper values that are used across multiple
// sub-resources. It must be called exactly once by main.go before any actual
// resource creation happens.
func initializeLocals(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) (*Locals, error) {
    if ctx == nil {
        return nil, errors.New("pulumi context is nil")
    }
    if stackInput == nil {
        return nil, errors.New("AwsDynamodbStackInput is nil")
    }
    if stackInput.Target == nil {
        return nil, errors.New("target resource is nil in stack input")
    }

    locals := &Locals{
        Target: stackInput.Target,
        Labels: map[string]string{},
        Tags:   pulumi.StringMap{},
    }

    // Attempt to enrich labels and tags from the user-provided spec. We guard
    // every access behind nil checks so that the code keeps compiling even
    // when the proto definitions evolve.
    if spec := stackInput.Target.GetSpec(); spec != nil {
        // Table name label – useful for naming ancillary resources such as
        // IAM roles or monitoring dashboards.
        if tableName := spec.GetTableName(); tableName != "" {
            locals.Labels["table_name"] = tableName
        }

        // Copy user-defined tags verbatim, converting them to Pulumi inputs.
        for k, v := range spec.GetTags() {
            if v != "" {
                locals.Tags[k] = pulumi.String(v)
            }
        }
    }

    return locals, nil
}
