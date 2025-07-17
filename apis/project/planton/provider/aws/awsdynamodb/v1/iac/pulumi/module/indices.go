package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
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

// The remaining helper functions stay unchanged.
