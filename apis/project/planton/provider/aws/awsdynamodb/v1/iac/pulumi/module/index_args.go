package awsdynamodb

// index_args.go contains helper utilities that translate the high-level
// AwsDynamodbSpec.{global_secondary_indexes,local_secondary_indexes}
// sections into the corresponding Pulumi input types expected by
// aws.dynamodb.Table – namely TableGlobalSecondaryIndexArgs and
// TableLocalSecondaryIndexArgs. Keeping this logic isolated makes the
// stack-construction code easier to read and unit-test.

import (
    "fmt"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// BuildGlobalSecondaryIndexArgs converts a slice of spec-level global
// secondary index definitions into Pulumi inputs. It validates that the
// HASH key is present and maps projection / capacity settings.
func BuildGlobalSecondaryIndexArgs(gsis []*awsdynamodbpb.GlobalSecondaryIndex) (dynamodb.TableGlobalSecondaryIndexArray, error) {
    var result dynamodb.TableGlobalSecondaryIndexArray

    for _, gsi := range gsis {
        if gsi == nil {
            // Skip nil entries to be defensive – the protobuf validation layer
            // should guarantee this never happens, but it costs us nothing to
            // double-check.
            continue
        }

        // ------------------------------------------------------------------
        // Extract HASH / RANGE keys from the key schema definition
        // ------------------------------------------------------------------
        var (
            hashKey  string
            rangeKey *string
        )
        for _, ks := range gsi.GetKeySchema() {
            switch ks.GetKeyType() {
            case awsdynamodbpb.KeyType_HASH:
                hashKey = ks.GetAttributeName()
            case awsdynamodbpb.KeyType_RANGE:
                r := ks.GetAttributeName()
                rangeKey = &r
            }
        }
        if hashKey == "" {
            return nil, fmt.Errorf("GSI %q is missing a HASH key entry", gsi.GetIndexName())
        }

        // ------------------------------------------------------------------
        // Projection settings
        // ------------------------------------------------------------------
        projType, err := projectionTypeToString(gsi.GetProjection().GetProjectionType())
        if err != nil {
            return nil, fmt.Errorf("GSI %q: %w", gsi.GetIndexName(), err)
        }
        var nonKeyAttrs pulumi.StringArray
        for _, attr := range gsi.GetProjection().GetNonKeyAttributes() {
            nonKeyAttrs = append(nonKeyAttrs, pulumi.String(attr))
        }

        // ------------------------------------------------------------------
        // Capacity – only set when present (i.e. PROVISIONED billing mode)
        // ------------------------------------------------------------------
        var readCap, writeCap pulumi.IntPtrInput
        if pt := gsi.GetProvisionedThroughput(); pt != nil {
            readCap = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
            writeCap = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
        }

        // ------------------------------------------------------------------
        // Assemble the Pulumi args struct
        // ------------------------------------------------------------------
        args := dynamodb.TableGlobalSecondaryIndexArgs{
            Name:            pulumi.String(gsi.GetIndexName()),
            HashKey:         pulumi.String(hashKey),
            ProjectionType:  pulumi.String(projType),
            NonKeyAttributes: nonKeyAttrs,
        }
        if rangeKey != nil {
            args.RangeKey = pulumi.StringPtr(*rangeKey)
        }
        if readCap != nil {
            args.ReadCapacity = readCap
        }
        if writeCap != nil {
            args.WriteCapacity = writeCap
        }

        result = append(result, args)
    }

    return result, nil
}

// BuildLocalSecondaryIndexArgs converts local secondary index definitions into
// Pulumi input structs.
func BuildLocalSecondaryIndexArgs(lsis []*awsdynamodbpb.LocalSecondaryIndex) (dynamodb.TableLocalSecondaryIndexArray, error) {
    var result dynamodb.TableLocalSecondaryIndexArray

    for _, lsi := range lsis {
        if lsi == nil {
            continue
        }

        // Extract RANGE key – LSIs share the HASH key with the table, so the
        // schema must contain exactly one RANGE entry.
        var rangeKey string
        for _, ks := range lsi.GetKeySchema() {
            if ks.GetKeyType() == awsdynamodbpb.KeyType_RANGE {
                rangeKey = ks.GetAttributeName()
                break
            }
        }
        if rangeKey == "" {
            return nil, fmt.Errorf("LSI %q is missing a RANGE key entry", lsi.GetIndexName())
        }

        // Projection settings
        projType, err := projectionTypeToString(lsi.GetProjection().GetProjectionType())
        if err != nil {
            return nil, fmt.Errorf("LSI %q: %w", lsi.GetIndexName(), err)
        }
        var nonKeyAttrs pulumi.StringArray
        for _, attr := range lsi.GetProjection().GetNonKeyAttributes() {
            nonKeyAttrs = append(nonKeyAttrs, pulumi.String(attr))
        }

        args := dynamodb.TableLocalSecondaryIndexArgs{
            Name:            pulumi.String(lsi.GetIndexName()),
            RangeKey:        pulumi.String(rangeKey),
            ProjectionType:  pulumi.String(projType),
            NonKeyAttributes: nonKeyAttrs,
        }

        result = append(result, args)
    }

    return result, nil
}

// ---------------------------------------------------------------------------
// Internal helpers
// ---------------------------------------------------------------------------

// projectionTypeToString converts the protobuf enum into the string literal
// expected by the AWS API / Pulumi provider.
func projectionTypeToString(pt awsdynamodbpb.ProjectionType) (string, error) {
    switch pt {
    case awsdynamodbpb.ProjectionType_ALL:
        return "ALL", nil
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY", nil
    case awsdynamodbpb.ProjectionType_INCLUDE:
        return "INCLUDE", nil
    default:
        return "", fmt.Errorf("unsupported projection type %v", pt)
    }
}
