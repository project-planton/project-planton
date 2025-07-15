package module

import (
    "strings"

    "github.com/pkg/errors"
    awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
    aws "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Resources is the entry-point invoked by the Project Planton engine when the
// stack needs to be provisioned. The function must be idempotent – running it
// multiple times with the same input should yield the exact same cloud state.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
    // ────────────────────────────────────────────────────────────────────────────
    // 1. Build locals (validated + enriched configuration)
    // ────────────────────────────────────────────────────────────────────────────
    locals, err := initializeLocals(ctx, stackInput)
    if err != nil {
        return errors.Wrap(err, "failed to initialise locals")
    }

    // ────────────────────────────────────────────────────────────────────────────
    // 2. Configure providers (AWS Native + Classic)
    // ────────────────────────────────────────────────────────────────────────────
    awsCredential := stackInput.ProviderCredential

    var provider *aws.Provider
    var classicProvider *awsclassic.Provider

    if awsCredential == nil {
        provider, err = aws.NewProvider(ctx, "native-provider", &aws.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS provider")
        }
        classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS classic provider")
        }
    } else {
        provider, err = aws.NewProvider(ctx, "native-provider", &aws.ProviderArgs{
            AccessKey: pulumi.String(awsCredential.AccessKeyId),
            SecretKey: pulumi.String(awsCredential.SecretAccessKey),
            Region:    pulumi.String(awsCredential.Region),
        })
        if err != nil {
            return errors.Wrap(err, "failed to create AWS provider with custom credentials")
        }
        classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{
            AccessKey: pulumi.String(awsCredential.AccessKeyId),
            SecretKey: pulumi.String(awsCredential.SecretAccessKey),
            Region:    pulumi.String(awsCredential.Region),
        })
        if err != nil {
            return errors.Wrap(err, "failed to create AWS classic provider with custom credentials")
        }
    }

    // ────────────────────────────────────────────────────────────────────────────
    // 3. Create the DynamoDB table (main resource)
    // ────────────────────────────────────────────────────────────────────────────

    spec := locals.AwsDynamodb

    // Attributes conversion ----------------------------------------------------
    attrs := make([]dynamodb.TableAttributeArgs, len(spec.AttributeDefinitions))
    for i, a := range spec.AttributeDefinitions {
        var attrType string
        switch a.AttributeType {
        case awsdynamodbv1.AttributeType_STRING:
            attrType = "S"
        case awsdynamodbv1.AttributeType_NUMBER:
            attrType = "N"
        case awsdynamodbv1.AttributeType_BINARY:
            attrType = "B"
        default:
            return errors.Errorf("unsupported attribute type %v", a.AttributeType)
        }

        attrs[i] = dynamodb.TableAttributeArgs{
            Name: pulumi.String(a.AttributeName),
            Type: pulumi.String(attrType),
        }
    }

    // Key schema ---------------------------------------------------------------
    var hashKey, rangeKey *string
    for _, k := range spec.KeySchema {
        switch k.KeyType {
        case awsdynamodbv1.KeyType_HASH:
            v := k.AttributeName
            hashKey = &v
        case awsdynamodbv1.KeyType_RANGE:
            v := k.AttributeName
            rangeKey = &v
        default:
            return errors.Errorf("unsupported key type %v", k.KeyType)
        }
    }
    if hashKey == nil {
        return errors.New("partition (HASH) key must be defined")
    }

    // Billing / capacity -------------------------------------------------------
    var billingMode *string
    var readCap, writeCap *int
    switch spec.BillingMode {
    case awsdynamodbv1.BillingMode_PROVISIONED:
        bm := "PROVISIONED"
        billingMode = &bm
        rc := int(spec.ProvisionedThroughput.GetReadCapacityUnits())
        wc := int(spec.ProvisionedThroughput.GetWriteCapacityUnits())
        readCap = &rc
        writeCap = &wc
    case awsdynamodbv1.BillingMode_PAY_PER_REQUEST:
        bm := "PAY_PER_REQUEST"
        billingMode = &bm
    default:
        return errors.Errorf("unsupported billing mode %v", spec.BillingMode)
    }

    // Streams ------------------------------------------------------------------
    var streamEnabled pulumi.BoolPtrInput
    var streamViewType *string
    if spec.StreamSpecification != nil && spec.StreamSpecification.StreamEnabled {
        streamEnabled = pulumi.Bool(true)
        switch spec.StreamSpecification.StreamViewType {
        case awsdynamodbv1.StreamViewType_NEW_IMAGE:
            v := "NEW_IMAGE"; streamViewType = &v
        case awsdynamodbv1.StreamViewType_OLD_IMAGE:
            v := "OLD_IMAGE"; streamViewType = &v
        case awsdynamodbv1.StreamViewType_NEW_AND_OLD_IMAGES:
            v := "NEW_AND_OLD_IMAGES"; streamViewType = &v
        case awsdynamodbv1.StreamViewType_STREAM_KEYS_ONLY:
            v := "KEYS_ONLY"; streamViewType = &v
        default:
            return errors.Errorf("unsupported stream view type %v", spec.StreamSpecification.StreamViewType)
        }
    }

    // Point-in-time recovery ----------------------------------------------------
    var pointInTimeRecovery pulumi.BoolPtrInput
    if spec.PointInTimeRecoveryEnabled {
        pointInTimeRecovery = pulumi.Bool(true)
    }

    // SSE ----------------------------------------------------------------------
    var sseArgs *dynamodb.TableServerSideEncryptionArgs
    if spec.SseSpecification != nil && spec.SseSpecification.Enabled {
        sseArgs = &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if spec.SseSpecification.SseType == awsdynamodbv1.SSEType_KMS {
            sseArgs.KmsKeyArn = pulumi.StringPtr(spec.SseSpecification.KmsMasterKeyId)
        }
    }

    // Tags conversion ----------------------------------------------------------
    tags := pulumi.StringMap{}
    for k, v := range locals.Tags {
        tags[k] = pulumi.String(v)
    }

    // Global secondary indexes -------------------------------------------------
    gsiArgs := make([]dynamodb.TableGlobalSecondaryIndexArgs, len(spec.GlobalSecondaryIndexes))
    for i, g := range spec.GlobalSecondaryIndexes {
        var gHashKey, gRangeKey *string
        for _, k := range g.KeySchema {
            switch k.KeyType {
            case awsdynamodbv1.KeyType_HASH:
                v := k.AttributeName
                gHashKey = &v
            case awsdynamodbv1.KeyType_RANGE:
                v := k.AttributeName
                gRangeKey = &v
            }
        }
        if gHashKey == nil {
            return errors.Errorf("GSI %s is missing HASH key", g.IndexName)
        }
        projectionType := "ALL"
        if g.Projection != nil {
            switch g.Projection.ProjectionType {
            case awsdynamodbv1.ProjectionType_ALL:
                projectionType = "ALL"
            case awsdynamodbv1.ProjectionType_KEYS_ONLY:
                projectionType = "KEYS_ONLY"
            case awsdynamodbv1.ProjectionType_INCLUDE:
                projectionType = "INCLUDE"
            }
        }
        var nonKey pulumi.StringArray
        if g.Projection != nil {
            for _, n := range g.Projection.NonKeyAttributes {
                nonKey = append(nonKey, pulumi.String(n))
            }
        }
        var rc, wc *int
        if g.ProvisionedThroughput != nil {
            r := int(g.ProvisionedThroughput.ReadCapacityUnits)
            w := int(g.ProvisionedThroughput.WriteCapacityUnits)
            rc, wc = &r, &w
        }
        gsiArgs[i] = dynamodb.TableGlobalSecondaryIndexArgs{
            Name:               pulumi.String(g.IndexName),
            HashKey:            pulumi.String(*gHashKey),
            ProjectionType:     pulumi.String(projectionType),
            NonKeyAttributes:   nonKey,
            ReadCapacity:       pulumi.IntPtr(rc),
            WriteCapacity:      pulumi.IntPtr(wc),
        }
        if gRangeKey != nil {
            gsiArgs[i].RangeKey = pulumi.StringPtr(*gRangeKey)
        }
    }

    // Local secondary indexes --------------------------------------------------
    lsiArgs := make([]dynamodb.TableLocalSecondaryIndexArgs, len(spec.LocalSecondaryIndexes))
    for i, l := range spec.LocalSecondaryIndexes {
        // For an LSI the HASH key must be the same as table HASH key – AWS will
        // validate that for us so we don’t re-check here.
        var lRangeKey *string
        for _, k := range l.KeySchema {
            if k.KeyType == awsdynamodbv1.KeyType_RANGE {
                v := k.AttributeName
                lRangeKey = &v
            }
        }
        if lRangeKey == nil {
            return errors.Errorf("LSI %s is missing RANGE key", l.IndexName)
        }
        projectionType := "ALL"
        if l.Projection != nil {
            switch l.Projection.ProjectionType {
            case awsdynamodbv1.ProjectionType_ALL:
                projectionType = "ALL"
            case awsdynamodbv1.ProjectionType_KEYS_ONLY:
                projectionType = "KEYS_ONLY"
            case awsdynamodbv1.ProjectionType_INCLUDE:
                projectionType = "INCLUDE"
            }
        }
        var nonKey pulumi.StringArray
        if l.Projection != nil {
            for _, n := range l.Projection.NonKeyAttributes {
                nonKey = append(nonKey, pulumi.String(n))
            }
        }
        lsiArgs[i] = dynamodb.TableLocalSecondaryIndexArgs{
            Name:             pulumi.String(l.IndexName),
            RangeKey:         pulumi.String(*lRangeKey),
            ProjectionType:   pulumi.String(projectionType),
            NonKeyAttributes: nonKey,
        }
    }

    // TTL ----------------------------------------------------------------------
    var ttl *dynamodb.TableTtlArgs
    if spec.TtlSpecification != nil && spec.TtlSpecification.TtlEnabled {
        ttl = &dynamodb.TableTtlArgs{
            Enabled:     pulumi.Bool(true),
            AttributeName: pulumi.String(spec.TtlSpecification.AttributeName),
        }
    }

    // Finally create the table --------------------------------------------------
    table, err := dynamodb.NewTable(ctx, spec.TableName, &dynamodb.TableArgs{
        Attributes:            dynamodb.TableAttributeArray(attrs),
        HashKey:               pulumi.StringPtr(*hashKey),
        RangeKey:              pulumi.StringPtrRange(rangeKey),
        BillingMode:           pulumi.StringPtr(strings.ToUpper(*billingMode)),
        ReadCapacity:          pulumi.IntPtr(readCap),
        WriteCapacity:         pulumi.IntPtr(writeCap),
        StreamEnabled:         streamEnabled,
        StreamViewType:        pulumi.StringPtr(streamViewType),
        PointInTimeRecovery:   pointInTimeRecovery,
        ServerSideEncryption:  sseArgs,
        GlobalSecondaryIndexes: dynamodb.TableGlobalSecondaryIndexArray(gsiArgs),
        LocalSecondaryIndexes:  dynamodb.TableLocalSecondaryIndexArray(lsiArgs),
        Tags:                 tags,
        Ttl:                  ttl,
    }, pulumi.Provider(classicProvider))
    if err != nil {
        return errors.Wrap(err, "failed to create dynamodb table")
    }

    // ────────────────────────────────────────────────────────────────────────────
    // 4. Export outputs so that other stacks / systems can reference them.
    // ────────────────────────────────────────────────────────────────────────────

    ctx.Export(OpTableArn, table.Arn)
    ctx.Export(OpTableName, table.Name)
    ctx.Export(OpTableId, table.ID())

    if spec.StreamSpecification != nil && spec.StreamSpecification.StreamEnabled {
        ctx.Export(OpStreamStreamArn, table.StreamArn)
        ctx.Export(OpStreamStreamLabel, table.StreamLabel)
    }

    if sseArgs != nil && sseArgs.KmsKeyArn != nil {
        ctx.Export(OpKmsKeyArn, sseArgs.KmsKeyArn)
    }

    // GSI / LSI names – we rely on the spec because it’s guaranteed to match
    // the provisioned indexes’ names (Pulumi does not expose GSI/LSI names as
    // separate outputs).
    var gsiNames pulumi.StringArray
    for _, g := range spec.GlobalSecondaryIndexes {
        gsiNames = append(gsiNames, pulumi.String(g.IndexName))
    }
    ctx.Export(OpGlobalSecondaryIndexNames, gsiNames)

    var lsiNames pulumi.StringArray
    for _, l := range spec.LocalSecondaryIndexes {
        lsiNames = append(lsiNames, pulumi.String(l.IndexName))
    }
    ctx.Export(OpLocalSecondaryIndexNames, lsiNames)

    return nil
}
