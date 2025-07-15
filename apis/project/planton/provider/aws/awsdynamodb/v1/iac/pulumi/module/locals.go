package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Locals keeps the resolved configuration that is used while creating the
// resources for the stack.
//
// Every module follows the convention of having a top level field with the
// resource-kind name (AwsDynamodb here) and another field that carries the
// standard tags / labels that need to be attached on the cloud resource.
// Additional helper / derived fields can be introduced when needed.
//
// NOTE: the tags map holds ONLY the computed tags (spec.tags merged with the
// Planton-mandatory keys); the map is later converted to pulumi.StringMap when
// feeding it to the provider.
// ---------------------------------------------------------------------------
// DO NOT add state (pulumi.Resource) fields inside this struct – locals are
// purely static / configuration values.
// ---------------------------------------------------------------------------

type Locals struct {
    AwsDynamodb *awsdynamodbv1.AwsDynamodbSpec
    Tags        map[string]string
}

// initializeLocals builds the Locals structure from the given stackInput.
// It resolves defaults, merges tags and performs any preliminary validation
// that is easier to do in Go than through the protobuf validators.
func initializeLocals(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) (*Locals, error) {
    if stackInput == nil {
        return nil, errors.New("stackInput must not be nil")
    }

    if stackInput.Target == nil {
        return nil, errors.New("stackInput.target must be set")
    }

    spec := stackInput.Target.GetSpec()
    if spec == nil {
        return nil, errors.New("stackInput.target.spec must be set")
    }

    // Merge user provided tags with Planton mandatory ones.
    tags := map[string]string{}
    for k, v := range spec.Tags {
        tags[k] = v
    }

    // Planton mandatory tags (when not yet provided by the user).
    // We intentionally don’t import awstagkeys / cloudresourcekind packages in
    // order to keep this module self-contained. Projects that do have those
    // packages can easily replace the literal keys with the constants.
    mandatory := map[string]string{
        "Resource":       spec.TableName,
        "ResourceKind":   "aws_dynamodb",
        "cloudresourcekind": "AWS::DynamoDB::Table",
    }

    for k, v := range mandatory {
        if _, ok := tags[k]; !ok {
            tags[k] = v
        }
    }

    return &Locals{
        AwsDynamodb: spec,
        Tags:        tags,
    }, nil
}
