package awsdynamodb

import (
    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ConvertGlobalSecondaryIndexes turns the protobuf representation of a list of
// global secondary indexes into the Pulumi input slice expected by
// aws.dynamodb.Table.
func ConvertGlobalSecondaryIndexes(gs []*awsdynamodbpb.GlobalSecondaryIndex) dynamodb.TableGlobalSecondaryIndexArray {
    if len(gs) == 0 {
        return nil
    }

    var result dynamodb.TableGlobalSecondaryIndexArray
    for _, g := range gs {
        if g == nil {
            continue
        }

        // Extract HASH and RANGE keys from the key schema.
        var hashKey, rangeKey string
        for _, ks := range g.KeySchema {
            switch ks.KeyType {
            case awsdynamodbpb.KeyType_HASH:
                hashKey = ks.AttributeName
            case awsdynamodbpb.KeyType_RANGE:
                rangeKey = ks.AttributeName
            }
        }

        gsi := &dynamodb.TableGlobalSecondaryIndexArgs{
            Name:           pulumi.String(g.IndexName),
            ProjectionType: pulumi.StringPtr(projectionTypeToString(g.Projection.ProjectionType)),
        }

        if hashKey != "" {
            gsi.HashKey = pulumi.StringPtr(hashKey)
        }
        if rangeKey != "" {
            gsi.RangeKey = pulumi.StringPtr(rangeKey)
        }
        if len(g.Projection.NonKeyAttributes) > 0 {
            gsi.NonKeyAttributes = toPulumiStringArray(g.Projection.NonKeyAttributes)
        }
        if pt := g.ProvisionedThroughput; pt != nil {
            if pt.ReadCapacityUnits > 0 {
                gsi.ReadCapacity = pulumi.IntPtr(int(pt.ReadCapacityUnits))
            }
            if pt.WriteCapacityUnits > 0 {
                gsi.WriteCapacity = pulumi.IntPtr(int(pt.WriteCapacityUnits))
            }
        }

        result = append(result, gsi)
    }

    return result
}

// ConvertLocalSecondaryIndexes turns the protobuf representation of a list of
// local secondary indexes into the Pulumi input slice expected by
// aws.dynamodb.Table.
func ConvertLocalSecondaryIndexes(ls []*awsdynamodbpb.LocalSecondaryIndex) dynamodb.TableLocalSecondaryIndexArray {
    if len(ls) == 0 {
        return nil
    }

    var result dynamodb.TableLocalSecondaryIndexArray
    for _, l := range ls {
        if l == nil {
            continue
        }

        // Extract RANGE key (the HASH key is inherited from the table and is not
        // provided to the provider).
        var rangeKey string
        for _, ks := range l.KeySchema {
            if ks.KeyType == awsdynamodbpb.KeyType_RANGE {
                rangeKey = ks.AttributeName
                break
            }
        }

        lsi := &dynamodb.TableLocalSecondaryIndexArgs{
            Name:           pulumi.String(l.IndexName),
            ProjectionType: pulumi.StringPtr(projectionTypeToString(l.Projection.ProjectionType)),
        }

        if rangeKey != "" {
            lsi.RangeKey = pulumi.StringPtr(rangeKey)
        }
        if len(l.Projection.NonKeyAttributes) > 0 {
            lsi.NonKeyAttributes = toPulumiStringArray(l.Projection.NonKeyAttributes)
        }

        result = append(result, lsi)
    }

    return result
}

// projectionTypeToString converts the protobuf ProjectionType enum into the
// literal strings required by the AWS API.
func projectionTypeToString(pt awsdynamodbpb.ProjectionType) string {
    switch pt {
    case awsdynamodbpb.ProjectionType_ALL:
        return "ALL"
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY"
    case awsdynamodbpb.ProjectionType_INCLUDE:
        return "INCLUDE"
    default:
        return ""
    }
}

// toPulumiStringArray converts a slice of string primitives into Pulumi
// StringArray, which implements pulumi.StringArrayInput.
func toPulumiStringArray(values []string) pulumi.StringArray {
    var arr pulumi.StringArray
    for _, v := range values {
        arr = append(arr, pulumi.String(v))
    }
    return arr
}
