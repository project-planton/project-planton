package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// buildGlobalSecondaryIndexes converts the spec-defined Global Secondary Indexes
// into the Pulumi representation expected by aws.dynamodb.Table.
func buildGlobalSecondaryIndexes(spec *awsdynamodbpb.AwsDynamodbSpec) (dynamodb.TableGlobalSecondaryIndexArray, error) {
    if spec == nil || len(spec.GetGlobalSecondaryIndexes()) == 0 {
        return nil, nil
    }

    indices := make(dynamodb.TableGlobalSecondaryIndexArray, 0, len(spec.GetGlobalSecondaryIndexes()))

    for _, g := range spec.GetGlobalSecondaryIndexes() {
        if g == nil {
            continue
        }

        hashKey, rangeKey, err := extractKeys(g.GetKeySchema(), g.GetIndexName())
        if err != nil {
            return nil, errors.Wrap(err, "extracting key schema for GSI")
        }

        projType, err := projectionTypeToString(g.GetProjection().GetProjectionType())
        if err != nil {
            return nil, errors.Wrapf(err, "parsing projection type for GSI %s", g.GetIndexName())
        }

        gs := dynamodb.TableGlobalSecondaryIndexArgs{
            Name:            pulumi.String(g.GetIndexName()),
            HashKey:         pulumi.String(hashKey),
            ProjectionType:  pulumi.String(projType),
            NonKeyAttributes: toPulumiStringArray(g.GetProjection().GetNonKeyAttributes()),
        }

        if rangeKey != "" {
            gs.RangeKey = pulumi.StringPtr(rangeKey)
        }

        // Provisioned throughput is optional (only when billing mode is PROVISIONED).
        if pt := g.GetProvisionedThroughput(); pt != nil {
            if pt.GetReadCapacityUnits() > 0 {
                gs.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
            }
            if pt.GetWriteCapacityUnits() > 0 {
                gs.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
            }
        }

        indices = append(indices, gs)
    }

    return indices, nil
}

// buildLocalSecondaryIndexes converts the spec-defined Local Secondary Indexes
// into the Pulumi representation expected by aws.dynamodb.Table.
func buildLocalSecondaryIndexes(spec *awsdynamodbpb.AwsDynamodbSpec) (dynamodb.TableLocalSecondaryIndexArray, error) {
    if spec == nil || len(spec.GetLocalSecondaryIndexes()) == 0 {
        return nil, nil
    }

    indices := make(dynamodb.TableLocalSecondaryIndexArray, 0, len(spec.GetLocalSecondaryIndexes()))

    for _, l := range spec.GetLocalSecondaryIndexes() {
        if l == nil {
            continue
        }

        // The HASH key is inherited from the table; we only need the RANGE key for a LSI.
        var rangeKey string
        for _, k := range l.GetKeySchema() {
            if k.GetKeyType() == awsdynamodbpb.KeyType_RANGE {
                rangeKey = k.GetAttributeName()
                break
            }
        }
        if rangeKey == "" {
            return nil, errors.Errorf("LSI %s missing RANGE key in key_schema", l.GetIndexName())
        }

        projType, err := projectionTypeToString(l.GetProjection().GetProjectionType())
        if err != nil {
            return nil, errors.Wrapf(err, "parsing projection type for LSI %s", l.GetIndexName())
        }

        ls := dynamodb.TableLocalSecondaryIndexArgs{
            Name:            pulumi.String(l.GetIndexName()),
            RangeKey:        pulumi.String(rangeKey),
            ProjectionType:  pulumi.String(projType),
            NonKeyAttributes: toPulumiStringArray(l.GetProjection().GetNonKeyAttributes()),
        }

        indices = append(indices, ls)
    }

    return indices, nil
}

// extractKeys returns the HASH and (optional) RANGE key attribute names from the
// provided key schema.
func extractKeys(schema []*awsdynamodbpb.KeySchemaElement, indexName string) (hashKey string, rangeKey string, err error) {
    for _, k := range schema {
        switch k.GetKeyType() {
        case awsdynamodbpb.KeyType_HASH:
            hashKey = k.GetAttributeName()
        case awsdynamodbpb.KeyType_RANGE:
            rangeKey = k.GetAttributeName()
        default:
            return "", "", errors.Errorf("unknown key_type in key schema for index %s", indexName)
        }
    }

    if hashKey == "" {
        return "", "", errors.Errorf("index %s missing HASH key in key_schema", indexName)
    }

    return hashKey, rangeKey, nil
}

// projectionTypeToString maps the protobuf enum to the AWS/Pulumi string literal.
func projectionTypeToString(pt awsdynamodbpb.ProjectionType) (string, error) {
    switch pt {
    case awsdynamodbpb.ProjectionType_ALL:
        return "ALL", nil
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY", nil
    case awsdynamodbpb.ProjectionType_INCLUDE:
        return "INCLUDE", nil
    default:
        return "", errors.Errorf("unsupported projection_type %v", pt)
    }
}

// toPulumiStringArray converts a slice of strings into Pulumi's StringArray type.
func toPulumiStringArray(values []string) pulumi.StringArray {
    arr := make(pulumi.StringArray, len(values))
    for i, v := range values {
        arr[i] = pulumi.String(v)
    }
    return arr
}
