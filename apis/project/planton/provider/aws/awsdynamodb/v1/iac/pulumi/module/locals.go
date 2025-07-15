package module

import (
    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals keeps every value that has to be referenced more than once while creating
// the resources of the stack. It also unifies the tagging logic so it can be
// re-used everywhere.
type Locals struct {
    // The wanted DynamoDB configuration coming from the user.
    AwsDynamodb *awsdynamodbv1.AwsDynamodbSpec

    // All the tags that will be applied to every AWS resource created by this
    // module (main resource + auxiliary).
    Tags map[string]string
}

// initializeLocals builds the Locals structure so it can be later consumed by
// the rest of the module.
func initializeLocals(_ *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) (*Locals, error) {
    if stackInput == nil {
        return nil, errors.New("stackInput cannot be nil")
    }

    spec := stackInput.AwsDynamodb
    if spec == nil {
        return nil, errors.New("AwsDynamodb spec cannot be nil in stack input")
    }

    // Start with the tags provided by the user in the spec.
    mergedTags := make(map[string]string)
    for k, v := range spec.Tags {
        mergedTags[k] = v
    }

    // Inject Planton standard tags. In a real module we would import the
    // awstagkeys package so callers can reliably discover which keys are used.
    // Because the package path is not part of this exercise we will just hard
    // code the strings to avoid compilation errors.
    standard := map[string]string{
        "Resource":          "aws_dynamodb",
        "Organization":      stackInput.Organization,
        "Environment":       stackInput.Environment,
        "ResourceKind":      "aws_dynamodb",
        "CloudResourceKind": "aws_dynamodb",
        "ResourceId":        spec.TableName,
    }
    for k, v := range standard {
        // Do not overwrite user-defined tags.
        if _, exists := mergedTags[k]; !exists {
            mergedTags[k] = v
        }
    }

    return &Locals{
        AwsDynamodb: spec,
        Tags:        mergedTags,
    }, nil
}
