package awsdynamodb

import (
    "fmt"

    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// BuildGlobalSecondaryIndexArgs converts a list of protobuf-defined GlobalSecondaryIndex
// messages into the shape expected by Pulumi's aws.dynamodb.Table resource. The resulting
// slice can be assigned directly to dynamodb.TableArgs.GlobalSecondaryIndexes.
func BuildGlobalSecondaryIndexArgs(
    gsis []*awsdynamodbpb.GlobalSecondaryIndex,
    billingMode awsdynamodbpb.BillingMode,
) ([]dynamodb.TableGlobalSecondaryIndexArgs, error) {
    if len(gsis) == 0 {
        return nil, nil
    }

    out := make([]dynamodb.TableGlobalSecondaryIndexArgs, 0, len(gsis))
    for _, gsi := range gsis {
        if gsi == nil {
            continue // Nothing to do.
        }

        // Extract HASH / RANGE keys.
        hashKey, rangeKey, err := extractKeys(gsi.KeySchema)
        if err != nil {
            return nil, fmt.Errorf("gsi %q: %w", gsi.IndexName, err)
        }

        projType, nonKeyAttrs := convertProjection(gsi.Projection)

        args := dynamodb.TableGlobalSecondaryIndexArgs{
            Name:            pulumi.StringPtr(gsi.IndexName),
            HashKey:         pulumi.StringPtr(hashKey),
            NonKeyAttributes: nonKeyAttrs,
            ProjectionType:  pulumi.StringPtr(projType),
        }

        if rangeKey != "" {
            args.RangeKey = pulumi.StringPtr(rangeKey)
        }

        // Include capacity only when the table operates in PROVISIONED billing mode.
        if billingMode == awsdynamodbpb.BillingMode_PROVISIONED {
            args.ReadCapacity = pulumi.IntPtr(int(gsi.GetProvisionedThroughput().GetReadCapacityUnits()))
            args.WriteCapacity = pulumi.IntPtr(int(gsi.GetProvisionedThroughput().GetWriteCapacityUnits()))
        }

        out = append(out, args)
    }

    return out, nil
}

// BuildLocalSecondaryIndexArgs converts protobuf-defined LocalSecondaryIndex messages into
// Pulumi aws.dynamodb.TableLocalSecondaryIndexArgs instances. The resulting slice can be
// assigned to dynamodb.TableArgs.LocalSecondaryIndexes.
func BuildLocalSecondaryIndexArgs(
    lsis []*awsdynamodbpb.LocalSecondaryIndex,
) ([]dynamodb.TableLocalSecondaryIndexArgs, error) {
    if len(lsis) == 0 {
        return nil, nil
    }

    out := make([]dynamodb.TableLocalSecondaryIndexArgs, 0, len(lsis))
    for _, lsi := range lsis {
        if lsi == nil {
            continue
        }

        _, rangeKey, err := extractKeys(lsi.KeySchema)
        if err != nil {
            return nil, fmt.Errorf("lsi %q: %w", lsi.IndexName, err)
        }
        if rangeKey == "" {
            return nil, fmt.Errorf("lsi %q: RANGE key not found", lsi.IndexName)
        }

        projType, nonKeyAttrs := convertProjection(lsi.Projection)

        args := dynamodb.TableLocalSecondaryIndexArgs{
            Name:            pulumi.StringPtr(lsi.IndexName),
            RangeKey:        pulumi.StringPtr(rangeKey),
            NonKeyAttributes: nonKeyAttrs,
            ProjectionType:  pulumi.StringPtr(projType),
        }

        out = append(out, args)
    }

    return out, nil
}

// extractKeys scans a protobuf key schema and returns the HASH and RANGE attribute names.
// HASH is required. RANGE may be empty when the schema only specifies a partition key.
func extractKeys(schema []*awsdynamodbpb.KeySchemaElement) (hashKey, rangeKey string, err error) {
    for _, elem := range schema {
        if elem == nil {
            continue
        }
        switch elem.KeyType {
        case awsdynamodbpb.KeyType_HASH:
            hashKey = elem.AttributeName
        case awsdynamodbpb.KeyType_RANGE:
            rangeKey = elem.AttributeName
        }
    }
    if hashKey == "" {
        err = fmt.Errorf("HASH key not defined in key schema")
    }
    return
}

// convertProjection converts a protobuf Projection message into the ProjectionType and
// NonKeyAttributes required by Pulumi.
func convertProjection(proj *awsdynamodbpb.Projection) (projType string, nonKeyAttrs pulumi.StringArray) {
    if proj == nil {
        // Default to projecting all attributes when the message is missing (should not happen).
        return "ALL", pulumi.StringArray{}
    }

    switch proj.ProjectionType {
    case awsdynamodbpb.ProjectionType_ALL:
        projType = "ALL"
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        projType = "KEYS_ONLY"
    case awsdynamodbpb.ProjectionType_INCLUDE:
        projType = "INCLUDE"
    default:
        projType = "ALL" // Fallback for unexpected enum values.
    }

    for _, a := range proj.NonKeyAttributes {
        nonKeyAttrs = append(nonKeyAttrs, pulumi.String(a))
    }

    return
}
