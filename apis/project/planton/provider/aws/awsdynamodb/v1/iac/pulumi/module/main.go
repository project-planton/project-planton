package main

import (
    "encoding/json"
    "fmt"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// Helpers ---------------------------------------------------------------------------------------
// Converts proto enum AttributeType to the string expected by AWS.
func attributeTypeToString(t awsdynamodbv1.AttributeType) (string, error) {
    switch t {
    case awsdynamodbv1.AttributeType_STRING:
        return "S", nil
    case awsdynamodbv1.AttributeType_NUMBER:
        return "N", nil
    case awsdynamodbv1.AttributeType_BINARY:
        return "B", nil
    default:
        return "", fmt.Errorf("unsupported attribute type %v", t)
    }
}

// Converts proto enum StreamViewType to AWS string constant.
func streamViewTypeToString(v awsdynamodbv1.StreamViewType) (string, error) {
    switch v {
    case awsdynamodbv1.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE", nil
    case awsdynamodbv1.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE", nil
    case awsdynamodbv1.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES", nil
    case awsdynamodbv1.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY", nil // provider calls it KEYS_ONLY
    default:
        return "", fmt.Errorf("unsupported stream view type %v", v)
    }
}

// Converts proto enum ProjectionType to AWS string constant.
func projectionTypeToString(p awsdynamodbv1.ProjectionType) (string, error) {
    switch p {
    case awsdynamodbv1.ProjectionType_ALL:
        return "ALL", nil
    case awsdynamodbv1.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY", nil
    case awsdynamodbv1.ProjectionType_INCLUDE:
        return "INCLUDE", nil
    default:
        return "", fmt.Errorf("unsupported projection type %v", p)
    }
}

// Extracts HASH and (optional) RANGE keys from a key schema list.
func resolveKeySchema(schema []*awsdynamodbv1.KeySchemaElement) (hashKey string, rangeKey *string, err error) {
    for _, ks := range schema {
        switch ks.KeyType {
        case awsdynamodbv1.KeyType_HASH:
            hashKey = ks.AttributeName
        case awsdynamodbv1.KeyType_RANGE:
            rk := ks.AttributeName
            rangeKey = &rk
        default:
            return "", nil, fmt.Errorf("unsupported key type %v", ks.KeyType)
        }
    }
    if hashKey == "" {
        err = fmt.Errorf("hash (partition) key not found in key schema")
    }
    return
}

func main() {
    pulumi.Run(func(ctx *pulumi.Context) error {
        // -------------------------------------------------------------------------------------
        // 1. Read the spec from stack configuration.
        // -------------------------------------------------------------------------------------
        cfg := config.New(ctx, "")
        var spec awsdynamodbv1.AwsDynamodbSpec
        {
            raw := cfg.Require("spec") // required – stack must define it.
            if err := json.Unmarshal([]byte(raw), &spec); err != nil {
                return fmt.Errorf("failed to unmarshal spec: %w", err)
            }
        }

        // -------------------------------------------------------------------------------------
        // 2. Translate proto spec into Pulumi/AWS args.
        // -------------------------------------------------------------------------------------
        // 2.1 Attributes ---------------------------------------------------------------------
        attrArgs := make(dynamodb.TableAttributeArray, 0, len(spec.AttributeDefinitions))
        for _, ad := range spec.AttributeDefinitions {
            attrType, err := attributeTypeToString(ad.AttributeType)
            if err != nil {
                return err
            }
            attrArgs = append(attrArgs, &dynamodb.TableAttributeArgs{
                Name: pulumi.String(ad.AttributeName),
                Type: pulumi.String(attrType),
            })
        }

        // 2.2 Primary key --------------------------------------------------------------------
        hashKey, rangeKey, err := resolveKeySchema(spec.KeySchema)
        if err != nil {
            return err
        }

        // 2.3 Billing mode & capacity ---------------------------------------------------------
        var billingMode string
        var readCap, writeCap pulumi.IntPtrInput
        switch spec.BillingMode {
        case awsdynamodbv1.BillingMode_PROVISIONED:
            billingMode = "PROVISIONED"
            readCap = pulumi.IntPtr(int(spec.ProvisionedThroughput.ReadCapacityUnits))
            writeCap = pulumi.IntPtr(int(spec.ProvisionedThroughput.WriteCapacityUnits))
        case awsdynamodbv1.BillingMode_PAY_PER_REQUEST:
            billingMode = "PAY_PER_REQUEST"
        default:
            return fmt.Errorf("unsupported billing mode %v", spec.BillingMode)
        }

        // 2.4 Global Secondary Indexes --------------------------------------------------------
        gsiArgs := make(dynamodb.TableGlobalSecondaryIndexArray, 0, len(spec.GlobalSecondaryIndexes))
        for _, g := range spec.GlobalSecondaryIndexes {
            gHash, gRange, err := resolveKeySchema(g.KeySchema)
            if err != nil {
                return err
            }
            projType, err := projectionTypeToString(g.Projection.ProjectionType)
            if err != nil {
                return err
            }
            gsi := &dynamodb.TableGlobalSecondaryIndexArgs{
                Name:           pulumi.String(g.IndexName),
                HashKey:        pulumi.String(gHash),
                ProjectionType: pulumi.String(projType),
            }
            if gRange != nil {
                gsi.RangeKey = pulumi.StringPtr(*gRange)
            }
            // Non-key attributes only for INCLUDE projections.
            if len(g.Projection.NonKeyAttributes) > 0 {
                nks := make(pulumi.StringArray, 0, len(g.Projection.NonKeyAttributes))
                for _, a := range g.Projection.NonKeyAttributes {
                    nks = append(nks, pulumi.String(a))
                }
                gsi.NonKeyAttributes = nks
            }
            // Capacity (only when PROVISIONED).
            if spec.BillingMode == awsdynamodbv1.BillingMode_PROVISIONED {
                gsi.ReadCapacity = pulumi.IntPtr(int(g.ProvisionedThroughput.ReadCapacityUnits))
                gsi.WriteCapacity = pulumi.IntPtr(int(g.ProvisionedThroughput.WriteCapacityUnits))
            }
            gsiArgs = append(gsiArgs, gsi)
        }

        // 2.5 Local Secondary Indexes ---------------------------------------------------------
        lsiArgs := make(dynamodb.TableLocalSecondaryIndexArray, 0, len(spec.LocalSecondaryIndexes))
        for _, l := range spec.LocalSecondaryIndexes {
            _, lRange, err := resolveKeySchema(l.KeySchema)
            if err != nil {
                return err
            }
            if lRange == nil {
                return fmt.Errorf("local secondary index %s must define a RANGE key", l.IndexName)
            }
            projType, err := projectionTypeToString(l.Projection.ProjectionType)
            if err != nil {
                return err
            }
            lsi := &dynamodb.TableLocalSecondaryIndexArgs{
                Name:           pulumi.String(l.IndexName),
                RangeKey:       pulumi.String(*lRange),
                ProjectionType: pulumi.String(projType),
            }
            if len(l.Projection.NonKeyAttributes) > 0 {
                nks := make(pulumi.StringArray, 0, len(l.Projection.NonKeyAttributes))
                for _, a := range l.Projection.NonKeyAttributes {
                    nks = append(nks, pulumi.String(a))
                }
                lsi.NonKeyAttributes = nks
            }
            lsiArgs = append(lsiArgs, lsi)
        }

        // 2.6 Streams ------------------------------------------------------------------------
        var streamEnabled pulumi.BoolPtrInput
        var streamViewType pulumi.StringPtrInput
        if spec.StreamSpecification.StreamEnabled {
            svt, err := streamViewTypeToString(spec.StreamSpecification.StreamViewType)
            if err != nil {
                return err
            }
            streamEnabled = pulumi.BoolPtr(true)
            streamViewType = pulumi.StringPtr(svt)
        }

        // 2.7 TTL -----------------------------------------------------------------------------
        var ttlArg *dynamodb.TableTtlArgs
        if spec.TtlSpecification.TtlEnabled {
            ttlArg = &dynamodb.TableTtlArgs{
                Enabled:       pulumi.Bool(true),
                AttributeName: pulumi.String(spec.TtlSpecification.AttributeName),
            }
        }

        // 2.8 Server-side encryption ----------------------------------------------------------
        var sseArg *dynamodb.TableServerSideEncryptionArgs
        if spec.SseSpecification.Enabled {
            sseArg = &dynamodb.TableServerSideEncryptionArgs{
                Enabled: pulumi.Bool(true),
            }
            if spec.SseSpecification.SseType == awsdynamodbv1.SSEType_KMS && spec.SseSpecification.KmsMasterKeyId != "" {
                sseArg.KmsKeyArn = pulumi.StringPtr(spec.SseSpecification.KmsMasterKeyId)
            }
        }

        // 2.9 Point-in-time recovery ----------------------------------------------------------
        var pitrArg *dynamodb.TablePointInTimeRecoveryArgs
        if spec.PointInTimeRecoveryEnabled {
            pitrArg = &dynamodb.TablePointInTimeRecoveryArgs{Enabled: pulumi.Bool(true)}
        }

        // 2.10 Tags ---------------------------------------------------------------------------
        tags := make(pulumi.StringMap)
        for k, v := range spec.Tags {
            tags[k] = pulumi.String(v)
        }

        // -------------------------------------------------------------------------------------
        // 3. Create DynamoDB Table resource.
        // -------------------------------------------------------------------------------------
        tableArgs := &dynamodb.TableArgs{
            Attributes: attrArgs,
            BillingMode: pulumi.String(billingMode),
            HashKey:     pulumi.String(hashKey),
            Tags:        tags,
        }
        if rangeKey != nil {
            tableArgs.RangeKey = pulumi.StringPtr(*rangeKey)
        }
        if readCap != nil {
            tableArgs.ReadCapacity = readCap
        }
        if writeCap != nil {
            tableArgs.WriteCapacity = writeCap
        }
        if len(gsiArgs) > 0 {
            tableArgs.GlobalSecondaryIndexes = gsiArgs
        }
        if len(lsiArgs) > 0 {
            tableArgs.LocalSecondaryIndexes = lsiArgs
        }
        if streamEnabled != nil {
            tableArgs.StreamEnabled = streamEnabled
            tableArgs.StreamViewType = streamViewType
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

        table, err := dynamodb.NewTable(ctx, spec.TableName, tableArgs)
        if err != nil {
            return err
        }

        // -------------------------------------------------------------------------------------
        // 4. Export outputs matching AwsDynamodbStackOutputs proto.
        // -------------------------------------------------------------------------------------
        ctx.Export("table_arn", table.Arn)
        ctx.Export("table_name", table.Name)
        ctx.Export("table_id", table.ID().ToStringOutput())

        if spec.StreamSpecification.StreamEnabled {
            stream := pulumi.All(table.StreamArn, table.StreamLabel).ApplyT(func(vs []interface{}) map[string]interface{} {
                return map[string]interface{}{
                    "stream_arn":   vs[0],
                    "stream_label": vs[1],
                }
            }).(pulumi.MapOutput)
            ctx.Export("stream", stream)
        }

        if sseArg != nil && sseArg.KmsKeyArn != nil {
            ctx.Export("kms_key_arn", sseArg.KmsKeyArn)
        }

        // Index names as plain, known-at-deploy constants -----------------------------------
        if len(spec.GlobalSecondaryIndexes) > 0 {
            arr := make(pulumi.StringArray, 0, len(spec.GlobalSecondaryIndexes))
            for _, g := range spec.GlobalSecondaryIndexes {
                arr = append(arr, pulumi.String(g.IndexName))
            }
            ctx.Export("global_secondary_index_names", arr)
        }
        if len(spec.LocalSecondaryIndexes) > 0 {
            arr := make(pulumi.StringArray, 0, len(spec.LocalSecondaryIndexes))
            for _, l := range spec.LocalSecondaryIndexes {
                arr = append(arr, pulumi.String(l.IndexName))
            }
            ctx.Export("local_secondary_index_names", arr)
        }

        return nil
    })
}
