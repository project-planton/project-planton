package module

import (
    "strings"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
    awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the entry-point called by the Pulumi engine.  It converts the
// StackInput received from the control-plane into real AWS resources and
// exports the outputs declared in AwsDynamodbStackOutputs.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
    // ---------------------------------------------------------------------
    // Locals & provider initialisation
    // ---------------------------------------------------------------------
    locals, err := initializeLocals(stackInput)
    if err != nil {
        return errors.Wrap(err, "failed to initialise locals")
    }
    spec := locals.AwsDynamodb.GetSpec()

    awsCredential := stackInput.GetProviderCredential()
    var nativeProvider *aws.Provider
    var classicProvider *awsclassic.Provider

    if awsCredential == nil {
        nativeProvider, err = aws.NewProvider(ctx, "native-provider", &aws.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS provider")
        }
        classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS classic provider")
        }
    } else {
        nativeProvider, err = aws.NewProvider(ctx, "native-provider", &aws.ProviderArgs{
            AccessKey: pulumi.String(awsCredential.GetAccessKeyId()),
            SecretKey: pulumi.String(awsCredential.GetSecretAccessKey()),
            Region:    pulumi.String(awsCredential.GetRegion()),
        })
        if err != nil {
            return errors.Wrap(err, "failed to create AWS provider with custom credentials")
        }
        classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{
            AccessKey: pulumi.String(awsCredential.GetAccessKeyId()),
            SecretKey: pulumi.String(awsCredential.GetSecretAccessKey()),
            Region:    pulumi.String(awsCredential.GetRegion()),
        })
        if err != nil {
            return errors.Wrap(err, "failed to create AWS classic provider")
        }
    }

    // ---------------------------------------------------------------------
    // Build DynamoDB table arguments
    // ---------------------------------------------------------------------

    tblArgs := &dynamodb.TableArgs{}

    // Attributes
    var attrs dynamodb.TableAttributeArray
    for _, ad := range spec.GetAttributeDefinitions() {
        attrs = append(attrs, dynamodb.TableAttributeArgs{
            Name: pulumi.String(ad.GetAttributeName()),
            Type: pulumi.String(attributeTypeToAWS(ad.GetAttributeType())),
        })
    }
    tblArgs.Attributes = attrs

    // Key-schema (HASH + optional RANGE)
    var hashKey string
    var rangeKey *string
    for _, ks := range spec.GetKeySchema() {
        switch ks.GetKeyType() {
        case awsdynamodbv1.KeyType_HASH:
            hashKey = ks.GetAttributeName()
        case awsdynamodbv1.KeyType_RANGE:
            rk := ks.GetAttributeName()
            rangeKey = &rk
        }
    }
    if hashKey == "" {
        return errors.New("HASH key must be supplied in key_schema")
    }
    tblArgs.HashKey = pulumi.String(hashKey)
    if rangeKey != nil {
        tblArgs.RangeKey = pulumi.StringPtr(*rangeKey)
    }

    // Billing mode & capacity
    switch spec.GetBillingMode() {
    case awsdynamodbv1.BillingMode_PROVISIONED:
        tblArgs.BillingMode = pulumi.StringPtr("PROVISIONED")
        pt := spec.GetProvisionedThroughput()
        tblArgs.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
        tblArgs.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
    case awsdynamodbv1.BillingMode_PAY_PER_REQUEST:
        tblArgs.BillingMode = pulumi.StringPtr("PAY_PER_REQUEST")
    default:
        return errors.Errorf("unsupported billing mode %v", spec.GetBillingMode())
    }

    // Global secondary indexes
    for _, g := range spec.GetGlobalSecondaryIndexes() {
        gsi := dynamodb.TableGlobalSecondaryIndexArgs{
            Name:           pulumi.String(g.GetIndexName()),
            ProjectionType: pulumi.String(projectionTypeToAWS(g.GetProjection().GetProjectionType())),
        }

        // non-key attributes (optional)
        if len(g.GetProjection().GetNonKeyAttributes()) > 0 {
            var nka pulumi.StringArray
            for _, a := range g.GetProjection().GetNonKeyAttributes() {
                nka = append(nka, pulumi.String(a))
            }
            gsi.NonKeyAttributes = nka
        }

        // key-schema for the index
        for _, ks := range g.GetKeySchema() {
            switch ks.GetKeyType() {
            case awsdynamodbv1.KeyType_HASH:
                gsi.HashKey = pulumi.String(ks.GetAttributeName())
            case awsdynamodbv1.KeyType_RANGE:
                gsi.RangeKey = pulumi.StringPtr(ks.GetAttributeName())
            }
        }

        // Provisioned capacity (only when overall billing is PROVISIONED)
        if pt := g.GetProvisionedThroughput(); pt != nil {
            gsi.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
            gsi.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
        }

        tblArgs.GlobalSecondaryIndexes = append(tblArgs.GlobalSecondaryIndexes, gsi)
    }

    // Local secondary indexes
    for _, l := range spec.GetLocalSecondaryIndexes() {
        lsi := dynamodb.TableLocalSecondaryIndexArgs{
            Name:           pulumi.String(l.GetIndexName()),
            ProjectionType: pulumi.String(projectionTypeToAWS(l.GetProjection().GetProjectionType())),
        }
        if len(l.GetProjection().GetNonKeyAttributes()) > 0 {
            var nka pulumi.StringArray
            for _, a := range l.GetProjection().GetNonKeyAttributes() {
                nka = append(nka, pulumi.String(a))
            }
            lsi.NonKeyAttributes = nka
        }
        // range key is mandatory for LSI (table hash key is automatically reused)
        for _, ks := range l.GetKeySchema() {
            if ks.GetKeyType() == awsdynamodbv1.KeyType_RANGE {
                lsi.RangeKey = pulumi.String(ks.GetAttributeName())
                break
            }
        }
        tblArgs.LocalSecondaryIndexes = append(tblArgs.LocalSecondaryIndexes, lsi)
    }

    // Streams
    if ss := spec.GetStreamSpecification(); ss != nil {
        tblArgs.StreamEnabled = pulumi.BoolPtr(ss.GetStreamEnabled())
        if ss.GetStreamEnabled() {
            tblArgs.StreamViewType = pulumi.StringPtr(streamViewTypeToAWS(ss.GetStreamViewType()))
        }
    }

    // TTL
    if ttl := spec.GetTtlSpecification(); ttl != nil {
        tblArgs.Ttl = &dynamodb.TableTtlArgs{
            AttributeName: pulumi.String(ttl.GetAttributeName()),
            Enabled:       pulumi.Bool(ttl.GetTtlEnabled()),
        }
    }

    // Server-side encryption
    if sse := spec.GetSseSpecification(); sse != nil && sse.GetEnabled() {
        tblArgs.ServerSideEncryption = &dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
            KmsKeyArn: func() pulumi.StringPtrInput {
                if sse.GetSseType() == awsdynamodbv1.SSEType_KMS {
                    return pulumi.StringPtr(sse.GetKmsMasterKeyId())
                }
                return nil
            }(),
        }
    }

    // Point-in-time recovery
    if spec.GetPointInTimeRecoveryEnabled() {
        tblArgs.PointInTimeRecovery = &dynamodb.TablePointInTimeRecoveryArgs{
            Enabled: pulumi.Bool(true),
        }
    }

    // Tags
    tags := pulumi.StringMap{}
    for k, v := range locals.Tags {
        tags[k] = pulumi.String(v)
    }
    tblArgs.Tags = tags

    // ---------------------------------------------------------------------
    // Create the DynamoDB table
    // ---------------------------------------------------------------------
    table, err := dynamodb.NewTable(ctx, "ddb-table", tblArgs, pulumi.Provider(classicProvider))
    if err != nil {
        return errors.Wrap(err, "failed to create dynamodb.Table")
    }

    // ---------------------------------------------------------------------
    // Exports
    // ---------------------------------------------------------------------
    ctx.Export(OpTableArn, table.Arn)
    ctx.Export(OpTableName, table.Name)
    ctx.Export(OpTableID, table.ID().ToStringOutput())

    // Streams (may be empty when disabled)
    ctx.Export(OpStreamStreamArn, table.StreamArn)
    ctx.Export(OpStreamStreamLabel, table.StreamLabel)

    // KMS (output only when SSE uses CMK)
    kmsKeyArn := ""
    if sse := spec.GetSseSpecification(); sse != nil {
        kmsKeyArn = sse.GetKmsMasterKeyId()
    }
    ctx.Export(OpKmsKeyArn, pulumi.String(kmsKeyArn))

    // Index name lists – take from spec as they are deterministic.
    var gsiNames pulumi.StringArray
    for _, g := range spec.GetGlobalSecondaryIndexes() {
        gsiNames = append(gsiNames, pulumi.String(g.GetIndexName()))
    }
    ctx.Export(OpGlobalSecondaryIndexNames, gsiNames)

    var lsiNames pulumi.StringArray
    for _, l := range spec.GetLocalSecondaryIndexes() {
        lsiNames = append(lsiNames, pulumi.String(l.GetIndexName()))
    }
    ctx.Export(OpLocalSecondaryIndexNames, lsiNames)

    // Silence unused variable warning for nativeProvider (not required right
    // now, but kept in case native resources are added in the future).
    _ = nativeProvider

    return nil
}

// ---------------------------------------------------------------------------
// Helper conversion utilities
// ---------------------------------------------------------------------------

func attributeTypeToAWS(t awsdynamodbv1.AttributeType) string {
    switch t {
    case awsdynamodbv1.AttributeType_STRING:
        return "S"
    case awsdynamodbv1.AttributeType_NUMBER:
        return "N"
    case awsdynamodbv1.AttributeType_BINARY:
        return "B"
    default:
        return "S"
    }
}

func projectionTypeToAWS(p awsdynamodbv1.ProjectionType) string {
    switch p {
    case awsdynamodbv1.ProjectionType_ALL:
        return "ALL"
    case awsdynamodbv1.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY"
    case awsdynamodbv1.ProjectionType_INCLUDE:
        return "INCLUDE"
    default:
        return "ALL"
    }
}

func streamViewTypeToAWS(s awsdynamodbv1.StreamViewType) string {
    switch s {
    case awsdynamodbv1.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE"
    case awsdynamodbv1.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE"
    case awsdynamodbv1.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES"
    case awsdynamodbv1.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY"
    default:
        return "NEW_AND_OLD_IMAGES"
    }
}

// Pulumi dislikes empty strings in optional fields that expect pointers. This
// helper converts "" into nil.
func strPtrIfNotEmpty(s string) *string {
    if strings.TrimSpace(s) == "" {
        return nil
    }
    return &s
}
