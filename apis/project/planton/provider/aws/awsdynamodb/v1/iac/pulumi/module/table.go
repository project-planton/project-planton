package module

import (
    "strings"

    "github.com/pkg/errors"
    awsprovider "github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// createTable provisions an Amazon DynamoDB table based on the supplied spec.
//
// The function is intended to be called from main.go; therefore it expects the
// computed locals and an (optional) explicit provider.  It translates the
// AwsDynamodbSpec proto message into the corresponding Pulumi
// dynamodb.TableArgs structure, handling all optional blocks and Pulumi-specific
// data-type conversions.
func createTable(
    ctx *pulumi.Context,
    locals *Locals,
    provider *awsprovider.Provider, // May be nil – Pulumi will use the default.
) (*dynamodb.Table, error) {
    if locals == nil || locals.Target == nil {
        return nil, errors.New("locals or locals.Target is nil – unable to create DynamoDB table")
    }

    spec := locals.Target.GetSpec()
    if spec == nil {
        return nil, errors.New("AwsDynamodbSpec is nil – invalid target resource")
    }

    // ---------------------------------------------------------------------
    // Attributes (AttributeDefinitions)
    // ---------------------------------------------------------------------
    attrInputs := make(dynamodb.TableAttributeArray, len(spec.AttributeDefinitions))
    for i, attr := range spec.AttributeDefinitions {
        attrType, err := mapAttributeType(attr.AttributeType)
        if err != nil {
            return nil, errors.Wrapf(err, "invalid attribute type for attribute %q", attr.AttributeName)
        }
        attrInputs[i] = dynamodb.TableAttributeArgs{
            Name: pulumi.String(attr.AttributeName),
            Type: pulumi.String(attrType),
        }
    }

    // ---------------------------------------------------------------------
    // Primary Key Schema – determine HashKey / RangeKey
    // ---------------------------------------------------------------------
    var hashKey, rangeKey string
    for _, ks := range spec.KeySchema {
        switch ks.KeyType {
        case awsdynamodbpb.KeyType_HASH:
            hashKey = ks.AttributeName
        case awsdynamodbpb.KeyType_RANGE:
            rangeKey = ks.AttributeName
        }
    }
    if hashKey == "" {
        return nil, errors.New("primary HASH key not provided in key_schema")
    }

    // ---------------------------------------------------------------------
    // Billing mode / Provisioned throughput
    // ---------------------------------------------------------------------
    var (
        billingMode               *string
        provisionedReadCapacity   *int
        provisionedWriteCapacity  *int
    )
    switch spec.BillingMode {
    case awsdynamodbpb.BillingMode_PROVISIONED:
        mode := "PROVISIONED"
        billingMode = &mode
        if spec.ProvisionedThroughput == nil {
            return nil, errors.New("billing_mode is PROVISIONED but provisioned_throughput is nil")
        }
        read := int(spec.ProvisionedThroughput.ReadCapacityUnits)
        write := int(spec.ProvisionedThroughput.WriteCapacityUnits)
        provisionedReadCapacity = &read
        provisionedWriteCapacity = &write
    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        mode := "PAY_PER_REQUEST"
        billingMode = &mode
    default:
        return nil, errors.Errorf("unsupported billing_mode: %v", spec.BillingMode)
    }

    // ---------------------------------------------------------------------
    // Global Secondary Indexes (GSIs)
    // ---------------------------------------------------------------------
    var gsiInputs dynamodb.TableGlobalSecondaryIndexArray
    for _, gsi := range spec.GlobalSecondaryIndexes {
        gHash, gRange, err := extractKeySchema(gsi.KeySchema)
        if err != nil {
            return nil, errors.Wrapf(err, "gsi %q", gsi.IndexName)
        }

        projType, err := mapProjectionType(gsi.Projection.ProjectionType)
        if err != nil {
            return nil, errors.Wrapf(err, "gsi %q projection", gsi.IndexName)
        }

        g := &dynamodb.TableGlobalSecondaryIndexArgs{
            Name:           pulumi.String(gsi.IndexName),
            HashKey:        pulumi.String(gHash),
            ProjectionType: pulumi.String(projType),
        }
        if gRange != "" {
            g.RangeKey = pulumi.StringPtr(gRange)
        }
        // Projection non-key attributes – only when INCLUDE.
        if projType == "INCLUDE" && len(gsi.Projection.NonKeyAttributes) > 0 {
            g.NonKeyAttributes = stringSliceToPulumiArray(gsi.Projection.NonKeyAttributes)
        }
        // Per-GSI capacity (only allowed when billing is PROVISIONED).
        if gsi.ProvisionedThroughput != nil {
            rc := int(gsi.ProvisionedThroughput.ReadCapacityUnits)
            wc := int(gsi.ProvisionedThroughput.WriteCapacityUnits)
            g.ReadCapacity = pulumi.IntPtr(rc)
            g.WriteCapacity = pulumi.IntPtr(wc)
        }

        gsiInputs = append(gsiInputs, g)
    }

    // ---------------------------------------------------------------------
    // Local Secondary Indexes (LSIs)
    // ---------------------------------------------------------------------
    var lsiInputs dynamodb.TableLocalSecondaryIndexArray
    for _, lsi := range spec.LocalSecondaryIndexes {
        lHash, lRange, err := extractKeySchema(lsi.KeySchema)
        if err != nil {
            return nil, errors.Wrapf(err, "lsi %q", lsi.IndexName)
        }
        // By definition, LSI shares HASH key with the table; validate if provided.
        if lHash != hashKey {
            return nil, errors.Errorf("lsi %q must use the same HASH key (%s) as the table", lsi.IndexName, hashKey)
        }

        projType, err := mapProjectionType(lsi.Projection.ProjectionType)
        if err != nil {
            return nil, errors.Wrapf(err, "lsi %q projection", lsi.IndexName)
        }

        l := &dynamodb.TableLocalSecondaryIndexArgs{
            Name:           pulumi.String(lsi.IndexName),
            RangeKey:       pulumi.String(lRange),
            ProjectionType: pulumi.String(projType),
        }
        if projType == "INCLUDE" && len(lsi.Projection.NonKeyAttributes) > 0 {
            l.NonKeyAttributes = stringSliceToPulumiArray(lsi.Projection.NonKeyAttributes)
        }
        lsiInputs = append(lsiInputs, l)
    }

    // ---------------------------------------------------------------------
    // Streams configuration
    // ---------------------------------------------------------------------
    var streamEnabled pulumi.BoolPtrInput
    var streamViewType pulumi.StringPtrInput
    if spec.StreamSpecification != nil && spec.StreamSpecification.StreamEnabled {
        streamEnabled = pulumi.BoolPtr(true)
        st, err := mapStreamViewType(spec.StreamSpecification.StreamViewType)
        if err != nil {
            return nil, errors.Wrap(err, "invalid stream_view_type")
        }
        streamViewType = pulumi.StringPtr(st)
    }

    // ---------------------------------------------------------------------
    // TTL configuration
    // ---------------------------------------------------------------------
    var ttlInput dynamodb.TableTtlPtrInput
    if spec.TtlSpecification != nil && spec.TtlSpecification.TtlEnabled {
        ttlInput = &dynamodb.TableTtlArgs{
            Enabled:       pulumi.Bool(true),
            AttributeName: pulumi.String(spec.TtlSpecification.AttributeName),
        }
    }

    // ---------------------------------------------------------------------
    // Server-side encryption (SSE)
    // ---------------------------------------------------------------------
    var sseInput dynamodb.TableServerSideEncryptionPtrInput
    if spec.SseSpecification != nil && spec.SseSpecification.Enabled {
        sse := &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if spec.SseSpecification.SseType == awsdynamodbpb.SSEType_KMS {
            // When KMS, the proto guarantees kms_master_key_id is non-empty.
            sse.KmsKeyArn = pulumi.StringPtr(spec.SseSpecification.KmsMasterKeyId)
        }
        sseInput = sse
    }

    // ---------------------------------------------------------------------
    // Aggregate tags (locals.Labels first, then spec.Tags to allow overrides).
    // ---------------------------------------------------------------------
    tagMap := map[string]string{}
    for k, v := range locals.Labels {
        tagMap[k] = v
    }
    for k, v := range spec.Tags {
        tagMap[k] = v
    }
    pulumiTags := pulumi.ToStringMap(tagMap)

    // ---------------------------------------------------------------------
    // Build TableArgs – only set optional fields when they are non-nil to avoid
    // Pulumi SDK complaining about mutually-exclusive settings.
    // ---------------------------------------------------------------------
    args := &dynamodb.TableArgs{
        Name:       pulumi.String(spec.TableName),
        Attributes: attrInputs,
        HashKey:    pulumi.String(hashKey),
        Tags:       pulumiTags,
    }
    // Range key is optional.
    if rangeKey != "" {
        args.RangeKey = pulumi.StringPtr(rangeKey)
    }
    if billingMode != nil {
        args.BillingMode = pulumi.StringPtr(*billingMode)
    }
    if provisionedReadCapacity != nil {
        args.ReadCapacity = pulumi.IntPtr(*provisionedReadCapacity)
    }
    if provisionedWriteCapacity != nil {
        args.WriteCapacity = pulumi.IntPtr(*provisionedWriteCapacity)
    }
    if len(gsiInputs) > 0 {
        args.GlobalSecondaryIndexes = gsiInputs
    }
    if len(lsiInputs) > 0 {
        args.LocalSecondaryIndexes = lsiInputs
    }
    if streamEnabled != nil {
        args.StreamEnabled = streamEnabled
    }
    if streamViewType != nil {
        args.StreamViewType = streamViewType
    }
    if ttlInput != nil {
        args.Ttl = ttlInput
    }
    if sseInput != nil {
        args.ServerSideEncryption = sseInput
    }
    if spec.PointInTimeRecoveryEnabled {
        args.PointInTimeRecovery = pulumi.BoolPtr(true)
    }

    // ---------------------------------------------------------------------
    // Finally, create the table.
    // ---------------------------------------------------------------------
    table, err := dynamodb.NewTable(ctx, spec.TableName, args, pulumi.Provider(provider))
    if err != nil {
        return nil, errors.Wrap(err, "creating aws.dynamodb.Table")
    }

    return table, nil
}

// -----------------------------------------------------------------------------
// Helper / mapping functions
// -----------------------------------------------------------------------------

func mapAttributeType(t awsdynamodbpb.AttributeType) (string, error) {
    switch t {
    case awsdynamodbpb.AttributeType_STRING:
        return "S", nil
    case awsdynamodbpb.AttributeType_NUMBER:
        return "N", nil
    case awsdynamodbpb.AttributeType_BINARY:
        return "B", nil
    default:
        return "", errors.Errorf("unsupported attribute type: %v", t)
    }
}

func mapProjectionType(t awsdynamodbpb.ProjectionType) (string, error) {
    switch t {
    case awsdynamodbpb.ProjectionType_ALL:
        return "ALL", nil
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY", nil
    case awsdynamodbpb.ProjectionType_INCLUDE:
        return "INCLUDE", nil
    default:
        return "", errors.Errorf("unsupported projection type: %v", t)
    }
}

func mapStreamViewType(t awsdynamodbpb.StreamViewType) (string, error) {
    switch t {
    case awsdynamodbpb.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE", nil
    case awsdynamodbpb.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE", nil
    case awsdynamodbpb.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES", nil
    case awsdynamodbpb.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY", nil
    default:
        return "", errors.Errorf("unsupported stream_view_type: %v", t)
    }
}

// extractKeySchema returns hashKey, rangeKey from the given repeated key schema.
func extractKeySchema(spec []*awsdynamodbpb.KeySchemaElement) (string, string, error) {
    var hashKey, rangeKey string
    for _, ks := range spec {
        switch ks.KeyType {
        case awsdynamodbpb.KeyType_HASH:
            hashKey = ks.AttributeName
        case awsdynamodbpb.KeyType_RANGE:
            rangeKey = ks.AttributeName
        }
    }
    if hashKey == "" {
        return "", "", errors.New("HASH key not found in key_schema")
    }
    return hashKey, rangeKey, nil
}

// stringSliceToPulumiArray converts a []string into pulumi.StringArray.
func stringSliceToPulumiArray(in []string) pulumi.StringArray {
    arr := make(pulumi.StringArray, len(in))
    for i, s := range in {
        arr[i] = pulumi.String(strings.TrimSpace(s))
    }
    return arr
}
