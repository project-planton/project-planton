package main

import (
    "fmt"

    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// createAwsDynamoDBTable provisions an Amazon DynamoDB table based on the
// AwsDynamodbSpec protobuf definition and returns the created Pulumi resource.
//
// The helper also exports the most relevant identifiers to the Pulumi stack so
// that they can be referenced from the CLI / console or from other stacks.
func createAwsDynamoDBTable(
    ctx *pulumi.Context,
    name string,
    spec *awsdynamodbpb.AwsDynamodbSpec,
    opts ...pulumi.ResourceOption,
) (*dynamodb.Table, error) {
    // ---------------------------------------------------------------------
    // 1.  Attribute definitions (required by every table / index reference)
    // ---------------------------------------------------------------------
    var attributes dynamodb.TableAttributeArray
    for _, attr := range spec.AttributeDefinitions {
        attributes = append(attributes, &dynamodb.TableAttributeArgs{
            Name: pulumi.String(attr.AttributeName),
            Type: pulumi.String(attributeTypeToString(attr.AttributeType)),
        })
    }

    // ---------------------------------------------------------------------
    // 2.  Key schema (Hash & optional Range keys) for the table itself
    // ---------------------------------------------------------------------
    var (
        hashKey  string
        rangeKey string
    )
    for _, ks := range spec.KeySchema {
        switch ks.KeyType {
        case awsdynamodbpb.KeyType_HASH:
            hashKey = ks.AttributeName
        case awsdynamodbpb.KeyType_RANGE:
            rangeKey = ks.AttributeName
        }
    }
    if hashKey == "" {
        return nil, fmt.Errorf("table must have a HASH key defined in key_schema")
    }

    // ---------------------------------------------------------------------
    // 3.  Global secondary indexes (GSIs)
    // ---------------------------------------------------------------------
    var gsis dynamodb.TableGlobalSecondaryIndexArray
    for _, g := range spec.GlobalSecondaryIndexes {
        var gsiHashKey, gsiRangeKey string
        for _, ks := range g.KeySchema {
            switch ks.KeyType {
            case awsdynamodbpb.KeyType_HASH:
                gsiHashKey = ks.AttributeName
            case awsdynamodbpb.KeyType_RANGE:
                gsiRangeKey = ks.AttributeName
            }
        }
        if gsiHashKey == "" {
            return nil, fmt.Errorf("global secondary index %q has no HASH key", g.IndexName)
        }

        gsi := &dynamodb.TableGlobalSecondaryIndexArgs{
            Name:           pulumi.String(g.IndexName),
            HashKey:        pulumi.String(gsiHashKey),
            ProjectionType: pulumi.String(projectionTypeToString(g.Projection.ProjectionType)),
        }
        if gsiRangeKey != "" {
            gsi.RangeKey = pulumi.StringPtr(gsiRangeKey)
        }
        if len(g.Projection.NonKeyAttributes) > 0 {
            var nks pulumi.StringArray
            for _, a := range g.Projection.NonKeyAttributes {
                nks = append(nks, pulumi.String(a))
            }
            gsi.NonKeyAttributes = nks
        }
        // Capacity settings only allowed in PROVISIONED mode
        if spec.BillingMode == awsdynamodbpb.BillingMode_PROVISIONED && g.ProvisionedThroughput != nil {
            gsi.ReadCapacity = pulumi.IntPtr(int(g.ProvisionedThroughput.ReadCapacityUnits))
            gsi.WriteCapacity = pulumi.IntPtr(int(g.ProvisionedThroughput.WriteCapacityUnits))
        }
        gsis = append(gsis, gsi)
    }

    // ---------------------------------------------------------------------
    // 4.  Local secondary indexes (LSIs)
    // ---------------------------------------------------------------------
    var lsis dynamodb.TableLocalSecondaryIndexArray
    for _, l := range spec.LocalSecondaryIndexes {
        var lsiRangeKey string
        for _, ks := range l.KeySchema {
            if ks.KeyType == awsdynamodbpb.KeyType_RANGE {
                lsiRangeKey = ks.AttributeName
            }
        }
        if lsiRangeKey == "" {
            return nil, fmt.Errorf("local secondary index %q has no RANGE key", l.IndexName)
        }
        lsi := &dynamodb.TableLocalSecondaryIndexArgs{
            Name:           pulumi.String(l.IndexName),
            RangeKey:       pulumi.String(lsiRangeKey),
            ProjectionType: pulumi.String(projectionTypeToString(l.Projection.ProjectionType)),
        }
        if len(l.Projection.NonKeyAttributes) > 0 {
            var nks pulumi.StringArray
            for _, a := range l.Projection.NonKeyAttributes {
                nks = append(nks, pulumi.String(a))
            }
            lsi.NonKeyAttributes = nks
        }
        lsis = append(lsis, lsi)
    }

    // ---------------------------------------------------------------------
    // 5.  TTL specification (optional)
    // ---------------------------------------------------------------------
    var ttl *dynamodb.TableTtlArgs
    if spec.TtlSpecification != nil && spec.TtlSpecification.TtlEnabled {
        ttl = &dynamodb.TableTtlArgs{
            Enabled:       pulumi.Bool(true),
            AttributeName: pulumi.String(spec.TtlSpecification.AttributeName),
        }
    }

    // ---------------------------------------------------------------------
    // 6.  Server-side encryption (optional)
    // ---------------------------------------------------------------------
    var sse *dynamodb.TableServerSideEncryptionArgs
    if spec.SseSpecification != nil && spec.SseSpecification.Enabled {
        sse = &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if spec.SseSpecification.SseType == awsdynamodbpb.SSEType_KMS && spec.SseSpecification.KmsMasterKeyId != "" {
            sse.KmsKeyArn = pulumi.StringPtr(spec.SseSpecification.KmsMasterKeyId)
        }
    }

    // ---------------------------------------------------------------------
    // 7.  Point-in-time recovery (continuous backups, optional)
    // ---------------------------------------------------------------------
    var pit *dynamodb.TablePointInTimeRecoveryArgs
    if spec.PointInTimeRecoveryEnabled {
        pit = &dynamodb.TablePointInTimeRecoveryArgs{
            Enabled: pulumi.Bool(true),
        }
    }

    // ---------------------------------------------------------------------
    // 8.  Tags
    // ---------------------------------------------------------------------
    tags := pulumi.StringMap{}
    for k, v := range spec.Tags {
        tags[k] = pulumi.String(v)
    }

    // ---------------------------------------------------------------------
    // 9.  Build the TableArgs object
    // ---------------------------------------------------------------------
    args := &dynamodb.TableArgs{
        Attributes:             attributes,
        BillingMode:            pulumi.String(billingModeToString(spec.BillingMode)),
        HashKey:                pulumi.String(hashKey),
        GlobalSecondaryIndexes: gsis,
        LocalSecondaryIndexes:  lsis,
        Tags:                   tags,
    }

    if rangeKey != "" {
        args.RangeKey = pulumi.StringPtr(rangeKey)
    }

    // Provisioned capacity only when billing mode is PROVISIONED.
    if spec.BillingMode == awsdynamodbpb.BillingMode_PROVISIONED && spec.ProvisionedThroughput != nil {
        args.ReadCapacity = pulumi.IntPtr(int(spec.ProvisionedThroughput.ReadCapacityUnits))
        args.WriteCapacity = pulumi.IntPtr(int(spec.ProvisionedThroughput.WriteCapacityUnits))
    }

    // Streams configuration
    if spec.StreamSpecification != nil && spec.StreamSpecification.StreamEnabled {
        args.StreamEnabled = pulumi.BoolPtr(true)
        args.StreamViewType = pulumi.StringPtr(streamViewTypeToString(spec.StreamSpecification.StreamViewType))
    }

    // TTL, SSE, PIT
    if ttl != nil {
        args.Ttl = ttl
    }
    if sse != nil {
        args.ServerSideEncryption = sse
    }
    if pit != nil {
        args.PointInTimeRecovery = pit
    }

    // Use the user-supplied name when provided; otherwise Pulumi auto-names.
    if spec.TableName != "" {
        args.Name = pulumi.String(spec.TableName)
    }

    // ---------------------------------------------------------------------
    // 10.  Create the resource
    // ---------------------------------------------------------------------
    table, err := dynamodb.NewTable(ctx, name, args, opts...)
    if err != nil {
        return nil, err
    }

    // ---------------------------------------------------------------------
    // 11.  Export commonly-used identifiers to the stack outputs
    // ---------------------------------------------------------------------
    ctx.Export("table_arn", table.Arn)
    ctx.Export("table_name", table.Name)
    ctx.Export("table_id", table.ID().ToStringOutput())

    // Streams (only present when enabled)
    ctx.Export("stream_arn", table.StreamArn)
    ctx.Export("stream_label", table.StreamLabel)

    // KMS key (when customer-managed key is configured)
    if sse != nil && sse.KmsKeyArn != nil {
        ctx.Export("kms_key_arn", sse.KmsKeyArn.(pulumi.StringPtrInput))
    }

    // Index names
    var gsiNames []pulumi.StringInput
    for _, g := range spec.GlobalSecondaryIndexes {
        gsiNames = append(gsiNames, pulumi.String(g.IndexName))
    }
    if len(gsiNames) > 0 {
        ctx.Export("global_secondary_index_names", pulumi.All(gsiNames...).ApplyT(func(_ []interface{}) []string {
            names := make([]string, len(gsiNames))
            for i, n := range gsiNames {
                names[i] = n.(pulumi.StringInput).StringValue().(string)
            }
            return names
        }).(pulumi.StringArrayOutput))
    }

    var lsiNames []pulumi.StringInput
    for _, l := range spec.LocalSecondaryIndexes {
        lsiNames = append(lsiNames, pulumi.String(l.IndexName))
    }
    if len(lsiNames) > 0 {
        ctx.Export("local_secondary_index_names", pulumi.All(lsiNames...).ApplyT(func(_ []interface{}) []string {
            names := make([]string, len(lsiNames))
            for i, n := range lsiNames {
                names[i] = n.(pulumi.StringInput).StringValue().(string)
            }
            return names
        }).(pulumi.StringArrayOutput))
    }

    return table, nil
}

// -----------------------------------------------------------------------------
// Helper conversion utilities
// -----------------------------------------------------------------------------

func attributeTypeToString(t awsdynamodbpb.AttributeType) string {
    switch t {
    case awsdynamodbpb.AttributeType_STRING:
        return "S"
    case awsdynamodbpb.AttributeType_NUMBER:
        return "N"
    case awsdynamodbpb.AttributeType_BINARY:
        return "B"
    default:
        return "" // Validation already ensures we never hit this.
    }
}

func billingModeToString(m awsdynamodbpb.BillingMode) string {
    switch m {
    case awsdynamodbpb.BillingMode_PROVISIONED:
        return "PROVISIONED"
    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        return "PAY_PER_REQUEST"
    default:
        return "PROVISIONED" // Sensible AWS default
    }
}

func projectionTypeToString(p awsdynamodbpb.ProjectionType) string {
    switch p {
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

func streamViewTypeToString(v awsdynamodbpb.StreamViewType) string {
    switch v {
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
