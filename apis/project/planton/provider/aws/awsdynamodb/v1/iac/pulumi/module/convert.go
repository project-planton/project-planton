package module

import (
    "fmt"

    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    pb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// -----------------------------------------------------------------------------
// Primitive helpers
// -----------------------------------------------------------------------------

func intPtr(v int) *int { return &v }

// -----------------------------------------------------------------------------
// Enum conversions (helper variants local to this file)
// -----------------------------------------------------------------------------

// attrTypeToString converts AttributeType to the short form ("S" | "N" |
// "B") required by aws.dynamodb.TableAttribute.
func attrTypeToString(t pb.AttributeType) (string, error) {
    switch t {
    case pb.AttributeType_STRING:
        return "S", nil
    case pb.AttributeType_NUMBER:
        return "N", nil
    case pb.AttributeType_BINARY:
        return "B", nil
    default:
        return "", errors.Errorf("unsupported attribute type: %v", t)
    }
}

// projTypeToString converts ProjectionType to the string used by the Pulumi
// provider.
func projTypeToString(t pb.ProjectionType) (string, error) {
    switch t {
    case pb.ProjectionType_ALL:
        return "ALL", nil
    case pb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY", nil
    case pb.ProjectionType_INCLUDE:
        return "INCLUDE", nil
    default:
        return "", errors.Errorf("unsupported projection type: %v", t)
    }
}

// streamViewTypeToStringConv converts StreamViewType to the string used by Pulumi.
func streamViewTypeToStringConv(t pb.StreamViewType) (string, error) {
    switch t {
    case pb.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE", nil
    case pb.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE", nil
    case pb.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES", nil
    case pb.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY", nil
    default:
        return "", errors.Errorf("unsupported stream view type: %v", t)
    }
}

// -----------------------------------------------------------------------------
// Complex structure conversions
// -----------------------------------------------------------------------------

// toTableAttributes converts proto attribute definitions to Pulumi
// TableAttribute objects.
func toTableAttributes(in []*pb.AttributeDefinition) ([]dynamodb.TableAttribute, error) {
    attrs := make([]dynamodb.TableAttribute, len(in))
    for i, a := range in {
        typ, err := attrTypeToString(a.GetAttributeType())
        if err != nil {
            return nil, errors.Wrapf(err, "attribute[%d] (%s)", i, a.GetAttributeName())
        }
        attrs[i] = dynamodb.TableAttribute{
            Name: a.GetAttributeName(),
            Type: typ,
        }
    }
    return attrs, nil
}

// extractKeys returns the HASH (partition) and RANGE (sort) keys from a key
// schema definition.
func extractKeys(schema []*pb.KeySchemaElement) (partitionKey, sortKey string, err error) {
    for _, el := range schema {
        switch el.KeyType {
        case pb.KeyType_HASH:
            if partitionKey != "" {
                return "", "", fmt.Errorf("duplicate HASH key in schema")
            }
            partitionKey = el.AttributeName
        case pb.KeyType_RANGE:
            if sortKey != "" {
                return "", "", fmt.Errorf("duplicate RANGE key in schema")
            }
            sortKey = el.AttributeName
        default:
            return "", "", fmt.Errorf("unsupported key type %v", el.KeyType)
        }
    }

    if partitionKey == "" {
        return "", "", fmt.Errorf("partition (HASH) key missing in key schema")
    }

    return partitionKey, sortKey, nil
}

// toProvisionedThroughput converts the optional proto definition into read &
// write capacity pointers that the Pulumi provider expects. When the proto
// value is nil, both pointers are nil indicating on-demand billing.
func toProvisionedThroughput(pt *pb.ProvisionedThroughput) (read, write *int) {
    if pt == nil {
        return nil, nil
    }
    r := int(pt.GetReadCapacityUnits())
    w := int(pt.GetWriteCapacityUnits())
    return intPtr(r), intPtr(w)
}

// toGlobalSecondaryIndexes converts proto GSIs into Pulumi structures.
func toGlobalSecondaryIndexes(in []*pb.GlobalSecondaryIndex) ([]dynamodb.TableGlobalSecondaryIndex, error) {
    gsis := make([]dynamodb.TableGlobalSecondaryIndex, len(in))

    for i, g := range in {
        pk, sk, err := extractKeys(g.KeySchema)
        if err != nil {
            return nil, errors.Wrapf(err, "gsi[%d] (%s) invalid key schema", i, g.IndexName)
        }

        projType, err := projTypeToString(g.GetProjection().GetProjectionType())
        if err != nil {
            return nil, errors.Wrapf(err, "gsi[%d] (%s) projection type", i, g.IndexName)
        }

        read, write := toProvisionedThroughput(g.ProvisionedThroughput)

        gsis[i] = dynamodb.TableGlobalSecondaryIndex{
            Name:             g.GetIndexName(),
            HashKey:          pk,
            ProjectionType:   projType,
            NonKeyAttributes: g.GetProjection().GetNonKeyAttributes(),
            ReadCapacity:     read,
            WriteCapacity:    write,
        }

        if sk != "" {
            gsis[i].RangeKey = &sk
        }
    }

    return gsis, nil
}

// toLocalSecondaryIndexes converts proto LSIs into Pulumi structures.
func toLocalSecondaryIndexes(in []*pb.LocalSecondaryIndex) ([]dynamodb.TableLocalSecondaryIndex, error) {
    lsis := make([]dynamodb.TableLocalSecondaryIndex, len(in))

    for i, l := range in {
        _, sk, err := extractKeys(l.KeySchema)
        if err != nil {
            return nil, errors.Wrapf(err, "lsi[%d] (%s) invalid key schema", i, l.IndexName)
        }
        if sk == "" {
            return nil, errors.Errorf("lsi[%d] (%s) must define a RANGE key", i, l.IndexName)
        }

        projType, err := projTypeToString(l.GetProjection().GetProjectionType())
        if err != nil {
            return nil, errors.Wrapf(err, "lsi[%d] (%s) projection type", i, l.IndexName)
        }

        lsis[i] = dynamodb.TableLocalSecondaryIndex{
            Name:             l.GetIndexName(),
            RangeKey:         sk,
            ProjectionType:   projType,
            NonKeyAttributes: l.GetProjection().GetNonKeyAttributes(),
        }
    }

    return lsis, nil
}

// toTTL converts the proto TimeToLiveSpecification into the Pulumi
// TableTtlArgs structure. When TTL is nil or disabled, the function returns nil
// so that the field can be omitted from resource options.
func toTTL(ttl *pb.TimeToLiveSpecification) *dynamodb.TableTtlArgs {
    if ttl == nil || !ttl.GetTtlEnabled() {
        return nil
    }
    return &dynamodb.TableTtlArgs{
        Enabled:       pulumi.Bool(ttl.GetTtlEnabled()),
        AttributeName: pulumi.String(ttl.GetAttributeName()),
    }
}
