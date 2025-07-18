package module

import (
    "github.com/pkg/errors"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"

    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// table creates the DynamoDB table together with every feature that can be
// enabled via AwsDynamodbSpec: TTL, Streams, (P)ITR, SSE, GSIs and LSIs.
//
// It converts the protobuf-defined spec into the Pulumi AWS provider shapes.
// The caller must provide the already-configured AWS provider so the resources
// use the correct credentials/region.
func table(
    ctx *pulumi.Context,
    locals *Locals,
    provider *aws.Provider,
) (*dynamodb.Table, error) {
    if locals == nil || locals.Target == nil {
        return nil, errors.New("locals or target cannot be nil")
    }

    spec := locals.Target.GetSpec()
    if spec == nil {
        return nil, errors.New("target.spec is required")
    }

    // ---------------------------------------------------------------------
    // Attribute definitions ------------------------------------------------
    // ---------------------------------------------------------------------
    var attrDefs dynamodb.TableAttributeArray
    for _, a := range spec.GetAttributeDefinitions() {
        attrDefs = append(attrDefs, dynamodb.TableAttributeArgs{
            Name: pulumi.String(a.GetAttributeName()),
            Type: pulumi.String(attributeTypeToString(a.GetAttributeType())),
        })
    }

    // ---------------------------------------------------------------------
    // Primary key (hash + optional range) ----------------------------------
    // ---------------------------------------------------------------------
    hashKey, _ := findKey(spec.GetKeySchema(), awsdynamodbpb.KeyType_HASH)
    rangeKey, hasRange := findKey(spec.GetKeySchema(), awsdynamodbpb.KeyType_RANGE)

    // ---------------------------------------------------------------------
    // Billing ‑ provisioned vs. on-demand ----------------------------------
    // ---------------------------------------------------------------------
    var (
        billingMode               string
        readCapacity, writeCapacity pulumi.IntPtrInput
    )
    switch spec.GetBillingMode() {
    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        billingMode = "PAY_PER_REQUEST"
    default: // PROVISIONED is already validated in the spec.
        billingMode = "PROVISIONED"
        if pt := spec.GetProvisionedThroughput(); pt != nil {
            readCapacity = pulumi.Int(int(pt.GetReadCapacityUnits()))
            writeCapacity = pulumi.Int(int(pt.GetWriteCapacityUnits()))
        }
    }

    // ---------------------------------------------------------------------
    // Global secondary indexes --------------------------------------------
    // ---------------------------------------------------------------------
    var gsis dynamodb.TableGlobalSecondaryIndexArray
    for _, g := range spec.GetGlobalSecondaryIndexes() {
        gsi := dynamodb.TableGlobalSecondaryIndexArgs{
            Name:           pulumi.String(g.GetIndexName()),
            HashKey:        pulumi.String(mustHashKey(g.GetKeySchema())),
            ProjectionType: pulumi.String(projectionTypeToString(g.GetProjection().GetProjectionType())),
        }
        if rk, ok := findKey(g.GetKeySchema(), awsdynamodbpb.KeyType_RANGE); ok {
            gsi.RangeKey = pulumi.String(rk)
        }
        if g.GetProjection().GetProjectionType() == awsdynamodbpb.ProjectionType_INCLUDE {
            for _, attr := range g.GetProjection().GetNonKeyAttributes() {
                gsi.NonKeyAttributes = append(gsi.NonKeyAttributes, pulumi.String(attr))
            }
        }
        // Capacity only when the table is provisioned.
        if spec.GetBillingMode() == awsdynamodbpb.BillingMode_PROVISIONED {
            if pt := g.GetProvisionedThroughput(); pt != nil {
                gsi.ReadCapacity = pulumi.Int(int(pt.GetReadCapacityUnits()))
                gsi.WriteCapacity = pulumi.Int(int(pt.GetWriteCapacityUnits()))
            }
        }
        gsis = append(gsis, gsi)
    }

    // ---------------------------------------------------------------------
    // Local secondary indexes ---------------------------------------------
    // ---------------------------------------------------------------------
    var lsis dynamodb.TableLocalSecondaryIndexArray
    for _, l := range spec.GetLocalSecondaryIndexes() {
        lsi := dynamodb.TableLocalSecondaryIndexArgs{
            Name:           pulumi.String(l.GetIndexName()),
            RangeKey:       pulumi.String(mustRangeKey(l.GetKeySchema())),
            ProjectionType: pulumi.String(projectionTypeToString(l.GetProjection().GetProjectionType())),
        }
        if l.GetProjection().GetProjectionType() == awsdynamodbpb.ProjectionType_INCLUDE {
            for _, attr := range l.GetProjection().GetNonKeyAttributes() {
                lsi.NonKeyAttributes = append(lsi.NonKeyAttributes, pulumi.String(attr))
            }
        }
        lsis = append(lsis, lsi)
    }

    // ---------------------------------------------------------------------
    // Time-to-live (TTL) ----------------------------------------------------
    // ---------------------------------------------------------------------
    var ttl *dynamodb.TableTtlArgs
    if ttlSpec := spec.GetTtlSpecification(); ttlSpec != nil && ttlSpec.GetTtlEnabled() {
        ttl = &dynamodb.TableTtlArgs{
            Enabled:       pulumi.Bool(true),
            AttributeName: pulumi.String(ttlSpec.GetAttributeName()),
        }
    }

    // ---------------------------------------------------------------------
    // Server-side encryption ----------------------------------------------
    // ---------------------------------------------------------------------
    var sse *dynamodb.TableServerSideEncryptionArgs
    if sseSpec := spec.GetSseSpecification(); sseSpec != nil && sseSpec.GetEnabled() {
        sse = &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if sseSpec.GetSseType() == awsdynamodbpb.SSEType_KMS {
            sse.KmsKeyArn = pulumi.String(sseSpec.GetKmsMasterKeyId())
        }
    }

    // ---------------------------------------------------------------------
    // Point-in-time recovery ----------------------------------------------
    // ---------------------------------------------------------------------
    var pitr *dynamodb.TablePointInTimeRecoveryArgs
    if spec.GetPointInTimeRecoveryEnabled() {
        pitr = &dynamodb.TablePointInTimeRecoveryArgs{
            Enabled: pulumi.Bool(true),
        }
    }

    // ---------------------------------------------------------------------
    // Streams --------------------------------------------------------------
    // ---------------------------------------------------------------------
    var streamEnabled pulumi.BoolPtrInput
    var streamViewType pulumi.StringPtrInput
    if s := spec.GetStreamSpecification(); s != nil && s.GetStreamEnabled() {
        streamEnabled = pulumi.Bool(true)
        streamViewType = pulumi.String(streamViewTypeToString(s.GetStreamViewType()))
    }

    // ---------------------------------------------------------------------
    // Tags -----------------------------------------------------------------
    // ---------------------------------------------------------------------
    tags := pulumi.StringMap{}
    for k, v := range spec.GetTags() {
        tags[k] = pulumi.String(v)
    }
    for k, v := range locals.Labels {
        // Do not overwrite user-supplied tags
        if _, ok := tags[k]; !ok {
            tags[k] = pulumi.String(v)
        }
    }

    // ---------------------------------------------------------------------
    // Assemble the final TableArgs ----------------------------------------
    // ---------------------------------------------------------------------
    args := &dynamodb.TableArgs{
        Name:                   pulumi.String(spec.GetTableName()),
        Attributes:             attrDefs,
        HashKey:                pulumi.String(hashKey),
        BillingMode:            pulumi.StringPtr(billingMode),
        StreamEnabled:          streamEnabled,
        StreamViewType:         streamViewType,
        Tags:                   tags,
        GlobalSecondaryIndexes: gsis,
        LocalSecondaryIndexes:  lsis,
        Ttl:                    ttl,
        ServerSideEncryption:   sse,
        PointInTimeRecovery:    pitr,
    }
    if hasRange {
        args.RangeKey = pulumi.String(rangeKey)
    }
    if readCapacity != nil {
        args.ReadCapacity = readCapacity
    }
    if writeCapacity != nil {
        args.WriteCapacity = writeCapacity
    }

    // ---------------------------------------------------------------------
    // Create the table -----------------------------------------------------
    // ---------------------------------------------------------------------
    tbl, err := dynamodb.NewTable(
        ctx,
        locals.Names("dynamodb-table", spec.GetTableName()), // helper adds stack-unique suffixes
        args,
        pulumi.Provider(provider),
    )
    if err != nil {
        return nil, errors.Wrap(err, "creating DynamoDB table")
    }

    return tbl, nil
}

// -------------------------------------------------------------------------
// Helper functions ---------------------------------------------------------
// -------------------------------------------------------------------------

func attributeTypeToString(t awsdynamodbpb.AttributeType) string {
    switch t {
    case awsdynamodbpb.AttributeType_STRING:
        return "S"
    case awsdynamodbpb.AttributeType_NUMBER:
        return "N"
    case awsdynamodbpb.AttributeType_BINARY:
        return "B"
    default:
        return "S"
    }
}

func projectionTypeToString(t awsdynamodbpb.ProjectionType) string {
    switch t {
    case awsdynamodbpb.ProjectionType_ALL:
        return "ALL"
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY"
    case awsdynamodbpb.ProjectionType_INCLUDE:
        return "INCLUDE"
    default:
        return "ALL"
    }
}

func streamViewTypeToString(t awsdynamodbpb.StreamViewType) string {
    switch t {
    case awsdynamodbpb.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE"
    case awsdynamodbpb.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE"
    case awsdynamodbpb.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES"
    case awsdynamodbpb.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY"
    default:
        return "NEW_AND_OLD_IMAGES"
    }
}

// findKey searches a key schema slice for the requested key type, returning
// the attribute name and whether it was found.
func findKey(schema []*awsdynamodbpb.KeySchemaElement, keyType awsdynamodbpb.KeyType) (string, bool) {
    for _, ks := range schema {
        if ks.GetKeyType() == keyType {
            return ks.GetAttributeName(), true
        }
    }
    return "", false
}

// mustHashKey returns the HASH key of a schema or panics – only used for
// sources that are validated by protobuf constraints.
func mustHashKey(schema []*awsdynamodbpb.KeySchemaElement) string {
    if k, ok := findKey(schema, awsdynamodbpb.KeyType_HASH); ok {
        return k
    }
    panic("HASH key not found in key schema – validation should have caught this")
}

// mustRangeKey returns the RANGE key of a schema or panics – only used when
// spec validation guarantees its presence.
func mustRangeKey(schema []*awsdynamodbpb.KeySchemaElement) string {
    if k, ok := findKey(schema, awsdynamodbpb.KeyType_RANGE); ok {
        return k
    }
    panic("RANGE key not found in key schema – validation should have caught this")
}
