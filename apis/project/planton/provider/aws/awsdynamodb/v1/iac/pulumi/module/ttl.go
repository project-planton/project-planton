package module

import (
    "fmt"

    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// timeToLive configures the TTL (Time-To-Live) settings for a DynamoDB table.
// If TTL is not requested in the incoming specification, the function is a
// no-op and returns (nil, nil).
func timeToLive(
    ctx *pulumi.Context,
    locals *Locals,
    awsProvider *aws.Provider,
    tableName pulumi.StringInput,
) (*dynamodb.TableItemTtl, error) {
    // Safely pull the TTL specification from the locals structure.
    var ttlSpec *awsdynamodbpb.TimeToLiveSpecification
    if res := locals.Resource; res != nil {
        if spec := res.GetSpec(); spec != nil {
            ttlSpec = spec.GetTtlSpecification()
        }
    }

    // Nothing to do when the spec is nil or explicitly disabled.
    if ttlSpec == nil || !ttlSpec.GetTtlEnabled() {
        return nil, nil
    }

    attrName := ttlSpec.GetAttributeName()
    if attrName == "" {
        return nil, errors.New("ttl_specification.attribute_name must be provided when ttl_enabled is true")
    }

    resourceName := fmt.Sprintf("%s-ttl", locals.Resource.GetSpec().GetTableName())

    ttlResource, err := dynamodb.NewTableItemTtl(ctx, resourceName, &dynamodb.TableItemTtlArgs{
        TableName:     tableName,
        Enabled:       pulumi.Bool(true),
        AttributeName: pulumi.String(attrName),
    }, pulumi.Provider(awsProvider))
    if err != nil {
        return nil, errors.Wrap(err, "creating DynamoDB TTL configuration")
    }

    return ttlResource, nil
}
