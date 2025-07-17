package module

import (
    "github.com/pkg/errors"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createTable provisions the DynamoDB table itself together with all of the
// optional/related settings that can be expressed on the primary table
// resource (capacity, encryption, TTL, Streams, PITR, tags, GSIs & LSIs).
//
// The function intentionally receives only the values it really needs – the
// Pulumi context, a prepared Locals instance (which contains the parsed target
// resource) and the AWS provider.  All conversions from the proto-based spec
// types to the Pulumi SDK types happen in this layer so that callers do not
// need to care about the Pulumi representation.
func createTable(
    ctx *pulumi.Context,
    locals *Locals,
    provider *aws.Provider,
) (*dynamodb.Table, error) {
    // Convenience alias.
    spec := locals.Target.GetSpec()
    if spec == nil {
        return nil, errors.New("awsdynamodb table spec is nil")
    }

    // ---------------------------------------------------------------------
    //  Attribute definitions (schema for every referenced attribute)
    // ---------------------------------------------------------------------
    attrs := dynamodb.TableAttributeArray{}
    for _, ad := range spec.GetAttributeDefinitions() {
        var attrType string
        switch ad.GetAttributeType() {
        case awsdynamodbpb.AttributeType_STRING:
            attrType = "S"
        case awsdynamodbpb.AttributeType_NUMBER:
            attrType = "N"
        case awsdynamodbpb.AttributeType_BINARY:
            attrType = "B"
        default:
            attrType = "S" // Fallback so Pulumi won’t receive an empty string.
        }

        attrs = append(attrs, dynamodb.TableAttributeArgs{
            Name: pulumi.String(ad.GetAttributeName()),
            Type: pulumi.String(attrType),
        })
    }

    // ---------------------------------------------------------------------
    //  Primary key (hash + optional range key)
    // ---------------------------------------------------------------------
    var (
        hashKey  string
        rangeKey *string
    )
    for _, ks := range spec.GetKeySchema() {
        switch ks.GetKeyType() {
        case awsdynamodbpb.KeyType_HASH:
            hashKey = ks.GetAttributeName()
        case awsdynamodbpb.KeyType_RANGE:
            rk := ks.GetAttributeName()
            rangeKey = &rk
        }
    }
    if hashKey == "" {
        return nil, errors.New("no HASH key defined in key_schema")
    }

    // ---------------------------------------------------------------------
    //  Global secondary indexes (GSIs)
    // ---------------------------------------------------------------------
    gsiArgs := dynamodb.TableGlobalSecondaryIndexArray{}
    for _, g := range spec.GetGlobalSecondaryIndexes() {
        // Keys.
        var (
            gHash  string
            gRange *string
        )
        for _, k := range g.GetKeySchema() {
            switch k.GetKeyType() {
            case awsdynamodbpb.KeyType_HASH:
                gHash = k.GetAttributeName()
            case awsdynamodbpb.KeyType_RANGE:
                r := k.GetAttributeName()
                gRange = &r
            }
        }

        // Projection.
        proj := g.GetProjection()
        var projectionType string
        switch proj.GetProjectionType() {
        case awsdynamodbpb.ProjectionType_ALL:
            projectionType = "ALL"
        case awsdynamodbpb.ProjectionType_KEYS_ONLY:
            projectionType = "KEYS_ONLY"
        case awsdynamodbpb.ProjectionType_INCLUDE:
            projectionType = "INCLUDE"
        default:
            projectionType = "ALL"
        }

        gsi := dynamodb.TableGlobalSecondaryIndexArgs{
            Name:           pulumi.String(g.GetIndexName()),
            HashKey:        pulumi.String(gHash),
            ProjectionType: pulumi.StringPtr(projectionType),
        }
        if gRange != nil {
            gsi.RangeKey = pulumi.StringPtr(*gRange)
        }
        if projectionType == "INCLUDE" && len(proj.GetNonKeyAttributes()) > 0 {
            gsi.NonKeyAttributes = pulumi.ToStringArray(proj.GetNonKeyAttributes())
        }
        // Capacity settings are only valid when billing mode is PROVISIONED.
        if spec.GetBillingMode() == awsdynamodbpb.BillingMode_PROVISIONED {
            if pt := g.GetProvisionedThroughput(); pt != nil {
                gsi.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
                gsi.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
            }
        }

        gsiArgs = append(gsiArgs, gsi)
    }

    // ---------------------------------------------------------------------
    //  Local secondary indexes (LSIs)
    // ---------------------------------------------------------------------
    lsiArgs := dynamodb.TableLocalSecondaryIndexArray{}
    for _, l := range spec.GetLocalSecondaryIndexes() {
        var (
            lHash  string
            lRange string
        )
        for _, k := range l.GetKeySchema() {
            switch k.GetKeyType() {
            case awsdynamodbpb.KeyType_HASH:
                lHash = k.GetAttributeName()
            case awsdynamodbpb.KeyType_RANGE:
                lRange = k.GetAttributeName()
            }
        }

        proj := l.GetProjection()
        var projectionType string
        switch proj.GetProjectionType() {
        case awsdynamodbpb.ProjectionType_ALL:
            projectionType = "ALL"
        case awsdynamodbpb.ProjectionType_KEYS_ONLY:
            projectionType = "KEYS_ONLY"
        case awsdynamodbpb.ProjectionType_INCLUDE:
            projectionType = "INCLUDE"
        default:
            projectionType = "ALL"
        }

        lsi := dynamodb.TableLocalSecondaryIndexArgs{
            Name:           pulumi.String(l.GetIndexName()),
            RangeKey:       pulumi.String(lRange),
            ProjectionType: pulumi.StringPtr(projectionType),
        }
        if projectionType == "INCLUDE" && len(proj.GetNonKeyAttributes()) > 0 {
            lsi.NonKeyAttributes = pulumi.ToStringArray(proj.GetNonKeyAttributes())
        }
        // The HASH key of every LSI must be the same as the table’s, but the
        // Pulumi SDK does not explicitly ask for it.
        _ = lHash // retained for documentation / potential validation.

        lsiArgs = append(lsiArgs, lsi)
    }

    // ---------------------------------------------------------------------
    //  TTL specification
    // ---------------------------------------------------------------------
    var ttlArg *dynamodb.TableTtlArgs
    if ttl := spec.GetTtlSpecification(); ttl != nil && ttl.GetTtlEnabled() {
        ttlArg = &dynamodb.TableTtlArgs{
            Enabled:       pulumi.Bool(true),
            AttributeName: pulumi.String(ttl.GetAttributeName()),
        }
    }

    // ---------------------------------------------------------------------
    //  Server-side encryption (SSE)
    // ---------------------------------------------------------------------
    var sseArg *dynamodb.TableServerSideEncryptionArgs
    if sse := spec.GetSseSpecification(); sse != nil && sse.GetEnabled() {
        sseArg = &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if sse.GetKmsMasterKeyId() != "" {
            sseArg.KmsKeyArn = pulumi.StringPtr(sse.GetKmsMasterKeyId())
        }
    }

    // ---------------------------------------------------------------------
    //  Point-in-time recovery (PITR)
    // ---------------------------------------------------------------------
    var pitrArg *dynamodb.TablePointInTimeRecoveryArgs
    if spec.GetPointInTimeRecoveryEnabled() {
        pitrArg = &dynamodb.TablePointInTimeRecoveryArgs{
            Enabled: pulumi.Bool(true),
        }
    }

    // ---------------------------------------------------------------------
    //  Streams configuration
    // ---------------------------------------------------------------------
    var (
        streamEnabled   bool
        streamViewType  string
    )
    if ss := spec.GetStreamSpecification(); ss != nil && ss.GetStreamEnabled() {
        streamEnabled = true
        switch ss.GetStreamViewType() {
        case awsdynamodbpb.StreamViewType_NEW_IMAGE:
            streamViewType = "NEW_IMAGE"
        case awsdynamodbpb.StreamViewType_OLD_IMAGE:
            streamViewType = "OLD_IMAGE"
        case awsdynamodbpb.StreamViewType_NEW_AND_OLD_IMAGES:
            streamViewType = "NEW_AND_OLD_IMAGES"
        case awsdynamodbpb.StreamViewType_STREAM_KEYS_ONLY:
            streamViewType = "KEYS_ONLY"
        default:
            streamViewType = "NEW_AND_OLD_IMAGES"
        }
    }

    // ---------------------------------------------------------------------
    //  Billing mode & capacity
    // ---------------------------------------------------------------------
    var billingMode string
    switch spec.GetBillingMode() {
    case awsdynamodbpb.BillingMode_PROVISIONED:
        billingMode = "PROVISIONED"
    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        billingMode = "PAY_PER_REQUEST"
    default:
        billingMode = "PAY_PER_REQUEST"
    }

    // ---------------------------------------------------------------------
    //  Tags (merge resource-level tags with module-level labels)
    // ---------------------------------------------------------------------
    tags := pulumi.StringMap{}
    for k, v := range spec.GetTags() {
        tags[k] = pulumi.String(v)
    }
    for k, v := range locals.Labels {
        if _, exists := tags[k]; !exists {
            tags[k] = pulumi.String(v)
        }
    }

    // ---------------------------------------------------------------------
    //  Build the Pulumi Table arguments
    // ---------------------------------------------------------------------
    tableArgs := &dynamodb.TableArgs{
        Attributes:   attrs,
        BillingMode:  pulumi.StringPtr(billingMode),
        HashKey:      pulumi.String(hashKey),
        Name:         pulumi.StringPtr(spec.GetTableName()),
        Tags:         tags,
    }
    if rangeKey != nil {
        tableArgs.RangeKey = pulumi.StringPtr(*rangeKey)
    }
    if len(gsiArgs) > 0 {
        tableArgs.GlobalSecondaryIndexes = gsiArgs
    }
    if len(lsiArgs) > 0 {
        tableArgs.LocalSecondaryIndexes = lsiArgs
    }
    if billingMode == "PROVISIONED" {
        if pt := spec.GetProvisionedThroughput(); pt != nil {
            tableArgs.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
            tableArgs.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
        }
    }
    if streamEnabled {
        tableArgs.StreamEnabled = pulumi.BoolPtr(true)
        tableArgs.StreamViewType = pulumi.StringPtr(streamViewType)
    }
    if ttlArg != nil {
        tableArgs.Ttl = ttlArg
    }
    if sseArg != nil {
        tableArgs.ServerSideEncryption = sseArg
    }
    if pitrArg != nil {
        tableArgs.PointInTimeRecovery = pitrArg
    }

    // ---------------------------------------------------------------------
    //  Create the resource
    // ---------------------------------------------------------------------
    var opts []pulumi.ResourceOption
    if provider != nil {
        opts = append(opts, pulumi.Provider(provider))
    }

    table, err := dynamodb.NewTable(ctx, spec.GetTableName(), tableArgs, opts...)
    if err != nil {
        return nil, errors.Wrap(err, "creating DynamoDB table")
    }

    return table, nil
}
