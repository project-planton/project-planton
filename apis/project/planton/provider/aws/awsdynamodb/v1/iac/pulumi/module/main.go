package module

import (
    "github.com/pkg/errors"
    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    awsNative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
    awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates / updates every cloud resource that belongs to this Pulumi
// component. The function signature is fixed by Project Planton so external
// code can discover it via reflection.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
    // Build every local variable / tag that will be used later on.
    locals, err := initializeLocals(ctx, stackInput)
    if err != nil {
        return errors.Wrap(err, "initialising locals")
    }

    // ---------------------------------------------------------------------
    // Providers – classic and native. We need the classic one to actually
    // create the DynamoDB table (aws.dynamodb.Table) but having the native
    // one around lets other components piggy-back on it.
    // ---------------------------------------------------------------------

    awsCredential := stackInput.ProviderCredential
    var nativeProvider *awsNative.Provider
    var classicProvider *awsclassic.Provider

    if awsCredential == nil {
        nativeProvider, err = awsNative.NewProvider(ctx, "native-provider", &awsNative.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS native provider")
        }

        classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS classic provider")
        }
    } else {
        nativeProvider, err = awsNative.NewProvider(ctx, "native-provider", &awsNative.ProviderArgs{
            AccessKey: pulumi.String(awsCredential.AccessKeyId),
            SecretKey: pulumi.String(awsCredential.SecretAccessKey),
            Region:    pulumi.String(awsCredential.Region),
        })
        if err != nil {
            return errors.Wrap(err, "failed to create AWS native provider with credentials")
        }

        classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{
            AccessKey: pulumi.String(awsCredential.AccessKeyId),
            SecretKey: pulumi.String(awsCredential.SecretAccessKey),
            Region:    pulumi.String(awsCredential.Region),
        })
        if err != nil {
            return errors.Wrap(err, "failed to create AWS classic provider with credentials")
        }
    }

    // ---------------------------------------------------------------------
    // Main resource: DynamoDB table.
    // ---------------------------------------------------------------------

    tableArgs, err := buildTableArgs(locals)
    if err != nil {
        return errors.Wrap(err, "building DynamoDB TableArgs from spec")
    }

    table, err := dynamodb.NewTable(ctx, locals.AwsDynamodb.TableName, tableArgs, pulumi.Provider(classicProvider))
    if err != nil {
        return errors.Wrap(err, "creating aws_dynamodb table")
    }

    // ---------------------------------------------------------------------
    // Outputs – strictly follow the StackOutputs proto naming.
    // ---------------------------------------------------------------------
    ctx.Export(OpTableArn, table.Arn)
    ctx.Export(OpTableName, table.Name)
    ctx.Export(OpTableId, table.ID().ToStringOutput())

    // Stream outputs (might be empty when streams are disabled).
    ctx.Export(OpStreamStreamArn, table.StreamArn)
    ctx.Export(OpStreamStreamLabel, table.StreamLabel)

    // KMS CMK ARN when SSE = KMS.
    if locals.AwsDynamodb.SseSpecification != nil && locals.AwsDynamodb.SseSpecification.Enabled && locals.AwsDynamodb.SseSpecification.SseType == awsdynamodbv1.SSEType_KMS {
        ctx.Export(OpKmsKeyArn, pulumi.String(locals.AwsDynamodb.SseSpecification.KmsMasterKeyId))
    } else {
        ctx.Export(OpKmsKeyArn, pulumi.String(""))
    }

    // Secondary indexes.
    var gsiNames []string
    for _, g := range locals.AwsDynamodb.GlobalSecondaryIndexes {
        gsiNames = append(gsiNames, g.IndexName)
    }
    ctx.Export(OpGlobalSecondaryIndexNames, pulumi.ToStringArray(gsiNames))

    var lsiNames []string
    for _, l := range locals.AwsDynamodb.LocalSecondaryIndexes {
        lsiNames = append(lsiNames, l.IndexName)
    }
    ctx.Export(OpLocalSecondaryIndexNames, pulumi.ToStringArray(lsiNames))

    // Keeping the providers referenced to avoid being GC-ed by Pulumi while the
    // stack is being deployed.
    _ = nativeProvider

    return nil
}

// buildTableArgs converts the Business level spec into the low-level provider
// arguments expected by pulumi-aws classic.
func buildTableArgs(locals *Locals) (*dynamodb.TableArgs, error) {
    spec := locals.AwsDynamodb

    // ------------------------------------------------------------------
    // Basic attributes and key-schema.
    // ------------------------------------------------------------------
    var attrs dynamodb.TableAttributeArray
    for _, ad := range spec.AttributeDefinitions {
        attrs = append(attrs, dynamodb.TableAttributeArgs{
            Name: pulumi.String(ad.AttributeName),
            Type: pulumi.String(attributeTypeToString(ad.AttributeType)),
        })
    }

    // Figure out HASH / RANGE keys.
    var hashKey, rangeKey *string
    for _, ks := range spec.KeySchema {
        switch ks.KeyType {
        case awsdynamodbv1.KeyType_HASH:
            h := ks.AttributeName
            hashKey = &h
        case awsdynamodbv1.KeyType_RANGE:
            r := ks.AttributeName
            rangeKey = &r
        }
    }
    if hashKey == nil {
        return nil, errors.New("HASH key not found in key_schema – this should have been validated at API level")
    }

    // ------------------------------------------------------------------
    // Billing / capacity.
    // ------------------------------------------------------------------
    var billingMode pulumi.StringPtrInput
    var readCap, writeCap pulumi.IntPtrInput

    switch spec.BillingMode {
    case awsdynamodbv1.BillingMode_PROVISIONED:
        billingMode = pulumi.String("PROVISIONED")
        readCap = pulumi.Int(int(spec.ProvisionedThroughput.ReadCapacityUnits))
        writeCap = pulumi.Int(int(spec.ProvisionedThroughput.WriteCapacityUnits))
    case awsdynamodbv1.BillingMode_PAY_PER_REQUEST:
        billingMode = pulumi.String("PAY_PER_REQUEST")
    default:
        billingMode = pulumi.String("PAY_PER_REQUEST") // safe default
    }

    // ------------------------------------------------------------------
    // Global Secondary Indexes.
    // ------------------------------------------------------------------
    var gsiArray dynamodb.TableGlobalSecondaryIndexArray
    for _, g := range spec.GlobalSecondaryIndexes {
        // Build key-schema for the index.
        var gHashKey, gRangeKey string
        for _, ks := range g.KeySchema {
            switch ks.KeyType {
            case awsdynamodbv1.KeyType_HASH:
                gHashKey = ks.AttributeName
            case awsdynamodbv1.KeyType_RANGE:
                gRangeKey = ks.AttributeName
            }
        }

        gsi := dynamodb.TableGlobalSecondaryIndexArgs{
            Name:           pulumi.String(g.IndexName),
            HashKey:        pulumi.String(gHashKey),
            ProjectionType: pulumi.String(projectionTypeToString(g.Projection.ProjectionType)),
        }
        if gRangeKey != "" {
            gsi.RangeKey = pulumi.StringPtr(gRangeKey)
        }

        // When billing mode is PROVISIONED we also need capacity numbers.
        if spec.BillingMode == awsdynamodbv1.BillingMode_PROVISIONED {
            gsi.ReadCapacity  = pulumi.IntPtr(int(g.ProvisionedThroughput.ReadCapacityUnits))
            gsi.WriteCapacity = pulumi.IntPtr(int(g.ProvisionedThroughput.WriteCapacityUnits))
        }

        // Projection INCLUDE requires non-key attributes list.
        if g.Projection.ProjectionType == awsdynamodbv1.ProjectionType_INCLUDE {
            var nks pulumi.StringArray
            for _, attr := range g.Projection.NonKeyAttributes {
                nks = append(nks, pulumi.String(attr))
            }
            gsi.NonKeyAttributes = nks
        }
        gsiArray = append(gsiArray, gsi)
    }

    // ------------------------------------------------------------------
    // Local Secondary Indexes.
    // ------------------------------------------------------------------
    var lsiArray dynamodb.TableLocalSecondaryIndexArray
    for _, l := range spec.LocalSecondaryIndexes {
        // Key schema for LSI (HASH must match table; RANGE is mandatory)
        var lHashKey, lRangeKey string
        for _, ks := range l.KeySchema {
            switch ks.KeyType {
            case awsdynamodbv1.KeyType_HASH:
                lHashKey = ks.AttributeName
            case awsdynamodbv1.KeyType_RANGE:
                lRangeKey = ks.AttributeName
            }
        }

        lsi := dynamodb.TableLocalSecondaryIndexArgs{
            Name:           pulumi.String(l.IndexName),
            HashKey:        pulumi.String(lHashKey),
            RangeKey:       pulumi.String(lRangeKey),
            ProjectionType: pulumi.String(projectionTypeToString(l.Projection.ProjectionType)),
        }
        if l.Projection.ProjectionType == awsdynamodbv1.ProjectionType_INCLUDE {
            var nks pulumi.StringArray
            for _, attr := range l.Projection.NonKeyAttributes {
                nks = append(nks, pulumi.String(attr))
            }
            lsi.NonKeyAttributes = nks
        }
        lsiArray = append(lsiArray, lsi)
    }

    // ------------------------------------------------------------------
    // Stream specification.
    // ------------------------------------------------------------------
    var streamEnabled pulumi.BoolPtrInput
    var streamView pulumi.StringPtrInput
    if spec.StreamSpecification.StreamEnabled {
        streamEnabled = pulumi.Bool(true)
        streamView = pulumi.String(streamViewTypeToString(spec.StreamSpecification.StreamViewType))
    }

    // ------------------------------------------------------------------
    // TTL.
    // ------------------------------------------------------------------
    var ttl *dynamodb.TableTtlArgs
    if spec.TtlSpecification.TtlEnabled {
        ttl = &dynamodb.TableTtlArgs{
            Enabled: pulumi.Bool(true),
            AttributeName: pulumi.String(spec.TtlSpecification.AttributeName),
        }
    }

    // ------------------------------------------------------------------
    // Server-side encryption.
    // ------------------------------------------------------------------
    var sse *dynamodb.TableServerSideEncryptionArgs
    if spec.SseSpecification.Enabled {
        sse = &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if spec.SseSpecification.SseType == awsdynamodbv1.SSEType_KMS {
            sse.KmsKeyArn = pulumi.StringPtr(spec.SseSpecification.KmsMasterKeyId)
        }
    }

    // ------------------------------------------------------------------
    // Point-in-time recovery.
    // ------------------------------------------------------------------
    var pitr *dynamodb.TablePointInTimeRecoveryArgs
    if spec.PointInTimeRecoveryEnabled {
        pitr = &dynamodb.TablePointInTimeRecoveryArgs{Enabled: pulumi.Bool(true)}
    }

    // ------------------------------------------------------------------
    // Build the final arguments structure.
    // ------------------------------------------------------------------
    args := &dynamodb.TableArgs{
        Attributes: attrs,
        HashKey:    pulumi.String(*hashKey),
        Tags:       pulumi.ToStringMap(locals.Tags),
    }

    if rangeKey != nil {
        args.RangeKey = pulumi.StringPtr(*rangeKey)
    }

    if billingMode != nil {
        args.BillingMode = billingMode
    }

    if readCap != nil {
        args.ReadCapacity = readCap
    }
    if writeCap != nil {
        args.WriteCapacity = writeCap
    }

    if len(gsiArray) > 0 {
        args.GlobalSecondaryIndexes = gsiArray
    }
    if len(lsiArray) > 0 {
        args.LocalSecondaryIndexes = lsiArray
    }

    if streamEnabled != nil {
        args.StreamEnabled = streamEnabled
        args.StreamViewType = streamView
    }

    if ttl != nil {
        args.Ttl = ttl
    }

    if sse != nil {
        args.ServerSideEncryption = sse
    }

    if pitr != nil {
        args.PointInTimeRecovery = pitr
    }

    return args, nil
}
