package module

import (
    "github.com/pkg/errors"
    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Locals groups together values that are used in several different places
// while creating the DynamoDB related resources.  Having a single struct
// makes it easier to pass this information around without having to thread
// dozens of individual parameters through every helper function.
type Locals struct {
    // Target is the full CR-like API object that was supplied by the caller
    // and describes the desired state of the DynamoDB table.
    Target *awsdynamodbv1.AwsDynamodb

    // TableName is the final name that should be used when creating the
    // dynamodb.Table resource.  In the simplest case it is exactly the name
    // that the user has put into spec.table_name but callers are free to add
    // their own suffixes or prefixes prior to provisioning.
    TableName string

    // Tags is a plain map representation of the key/value tags that should be
    // attached to every AWS resource (where supported).
    Tags map[string]string
}

// initializeLocals converts the protobuf-based stack input into a strongly
// typed Go struct that contains the pre-processed values we need during
// provisioning.  Every piece of generic, re-usable information that more
// than one sub-resource needs should live here so we avoid duplicating the
// conversion logic in multiple places.
func initializeLocals(_ interface{}, stackInput *awsdynamodbv1.AwsDynamodbStackInput) (*Locals, error) {
    if stackInput == nil {
        return nil, errors.New("stackInput must not be nil")
    }

    target := stackInput.GetTarget()
    if target == nil {
        return nil, errors.New("stackInput.target must be provided")
    }

    spec := target.GetSpec()
    if spec == nil {
        return nil, errors.New("target.spec must be provided")
    }

    // The base table name is taken exactly as supplied by the user.
    tableName := spec.GetTableName()
    if tableName == "" {
        return nil, errors.New("spec.table_name must not be empty")
    }

    // Convert the user provided map[string]string into a Go map so the value
    // can be passed straight into the AWS SDK resource constructors.
    tags := map[string]string{}
    for k, v := range spec.GetTags() {
        tags[k] = v
    }

    // Inject a handful of standard tags that help operators understand where
    // the resource came from.  We only add them when the caller did *not*
    // already define a tag with the same key so that the user's preference
    // always wins.
    if _, ok := tags["managed-by"]; !ok {
        tags["managed-by"] = "pulumi"
    }
    if _, ok := tags["provisioner"]; !ok {
        tags["provisioner"] = "project-planton"
    }

    return &Locals{
        Target:    target,
        TableName: tableName,
        Tags:      tags,
    }, nil
}
