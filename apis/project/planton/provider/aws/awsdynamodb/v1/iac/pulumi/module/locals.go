package module

import (
    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Locals aggregates common values required while creating resources in this
// module. It contains the incoming API-resource as well as the merged set of
// tags that will be applied to every AWS resource created by the module.
type Locals struct {
    AwsDynamodb *awsdynamodbv1.AwsDynamodb

    // Final set of tags to be attached to AWS resources.
    Tags map[string]string
}

// initializeLocals populates a Locals instance, merging any user-supplied tags
// with the standard Planton tags that must be present on every AWS resource.
func initializeLocals(stackInput *awsdynamodbv1.AwsDynamodbStackInput) (*Locals, error) {
    target := stackInput.GetTarget()
    if target == nil {
        return nil, errors.New("awsdynamodb target is required in stackInput")
    }

    spec := target.GetSpec()
    if spec == nil {
        return nil, errors.New("awsdynamodb spec cannot be nil")
    }

    // Start with a copy of user-defined tags.
    tags := map[string]string{}
    for k, v := range spec.GetTags() {
        tags[k] = v
    }

    // Standard Planton tags – add or override as required.
    tags["ResourceKind"] = "AwsDynamodb"
    tags["ResourceId"] = spec.GetTableName()

    return &Locals{
        AwsDynamodb: target,
        Tags:        tags,
    }, nil
}