package module

import (
    "fmt"

    "github.com/pkg/errors"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    // AWS providers
    awsNative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
    awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"

    // Proto APIs
    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Resources is the single entry-point that Project Planton’s provisioning
// engine invokes when it needs to create/update/destroy an AWS DynamoDB table
// using Pulumi.  The function is intentionally designed to be deterministic and
// as side-effect-free as possible (aside from calling the cloud provider).
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
    // ---------------------------------------------------------------------
    // 1. Parse the high level input into convenient local objects
    // ---------------------------------------------------------------------
    locals, err := initializeLocals(stackInput)
    if err != nil {
        return errors.Wrap(err, "failed to initialise locals")
    }

    // Shorthand variables used through the rest of the function.
    target := locals.AwsDynamodb
    spec := target.GetSpec()

    // ---------------------------------------------------------------------
    // 2. Configure providers (native & classic)
    // ---------------------------------------------------------------------
    awsCredential := stackInput.GetProviderCredential()
    var nativeProvider *awsNative.Provider
    var classicProvider *awsclassic.Provider

    if awsCredential == nil {
        nativeProvider, err = awsNative.NewProvider(ctx, "native-provider", &awsNative.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS provider")
        }
        classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{})
        if err != nil {
            return errors.Wrap(err, "failed to create default AWS classic provider")
        }
    } else {
        nativeProvider, err = awsNative.NewProvider(ctx, "native-provider", &awsNative.ProviderArgs{
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
            return errors.Wrap(err, "failed to create AWS classic provider with custom credentials")
        }
    }

    // nativeProvider is currently not used for resource creation in this
    // example module. Adding the following assignment ensures the variable is
    // considered "used" by the compiler while preserving the opportunity to
    // leverage it in the future without further refactoring.
    _ = nativeProvider

    // ---------------------------------------------------------------------
    // 3. Translate the protobuf spec to Pulumi resource arguments
    // ---------------------------------------------------------------------

    tableArgs := &dynamodb.TableArgs{
        Name: pulumi.String(spec.GetTableName()),
        Tags: pulumi.ToStringMap(locals.Tags),
    }

    // 3.1  Attribute definitions
    var attributes dynamodb.TableAttributeArray
    for _, ad := range spec.GetAttributeDefinitions() {
        attributes = append(attributes, dynamodb.TableAttributeArgs{
            Name: pulumi.String(ad.GetAttributeName()),
            Type: pulumi.String(convertAttributeType(ad.GetAttributeType())),
        })
    }
    tableArgs.Attributes = attributes

    // 3.2  Primary key (partition + optional sort)
    var hashKey, rangeKey string
    for _, ks := range spec.GetKeySchema() {
        if ks.GetKeyType() == awsdynamodbv1.KeyType_HASH {
            hashKey = ks.GetAttributeName()
        } else if ks.GetKeyType() == awsdynamodbv1.KeyType_RANGE {
            rangeKey = ks.GetAttributeName()
        }
    }
    if hashKey == "" {
        return errors.New("primary HASH key is required but not found in key_schema")
    }
    tableArgs.HashKey = pulumi.String(hashKey)
    if rangeKey != "" {
        tableArgs.RangeKey = pulumi.StringPtr(rangeKey)
    }

    // 3.3  Billing mode & capacity units
    switch spec.GetBillingMode() {
    case awsdynamodbv1.BillingMode_PAY_PER_REQUEST:
        tableArgs.BillingMode = pulumi.StringPtr("PAY_PER_REQUEST")
    case awsdynamodbv1.BillingMode_PROVISIONED:
        pt := spec.GetProvisionedThroughput()
        if pt == nil {
            return errors.New("provisioned_throughput must be provided when billing_mode = PROVISIONED")
        }
        tableArgs.BillingMode = pulumi.StringPtr("PROVISIONED")
        tableArgs.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
        tableArgs.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
    default:
        return errors.Errorf("unknown billing mode %v", spec.GetBillingMode())
    }

    // 3.4  Global Secondary Indexes (GSIs)
    var gsiNames pulumi.StringArray
    if len(spec.GetGlobalSecondaryIndexes()) > 0 {
        var gsis dynamodb.TableGlobalSecondaryIndexArray
        for _, g := range spec.GetGlobalSecondaryIndexes() {
            gsiArg := dynamodb.TableGlobalSecondaryIndexArgs{
                Name:           pulumi.String(g.GetIndexName()),
                ProjectionType: pulumi.String(convertProjectionType(g.GetProjection().GetProjectionType())),
            }

            // Key schema for the GSI
            var gHash, gRange string
            for _, ks := range g.GetKeySchema() {
                if ks.GetKeyType() == awsdynamodbv1.KeyType_HASH {
                    gHash = ks.GetAttributeName()
                } else if ks.GetKeyType() == awsdynamodbv1.KeyType_RANGE {
                    gRange = ks.GetAttributeName()
                }
            }
            gsiArg.HashKey = pulumi.String(gHash)
            if gRange != "" {
                gsiArg.RangeKey = pulumi.StringPtr(gRange)
            }

            // Capacity for the index (only in PROVISIONED mode)
            if spec.GetBillingMode() == awsdynamodbv1.BillingMode_PROVISIONED {
                pt := g.GetProvisionedThroughput()
                if pt == nil {
                    return errors.Errorf("gsi %s is missing provisioned_throughput while table is PROVISIONED", g.GetIndexName())
                }
                gsiArg.ReadCapacity = pulumi.IntPtr(int(pt.GetReadCapacityUnits()))
                gsiArg.WriteCapacity = pulumi.IntPtr(int(pt.GetWriteCapacityUnits()))
            }

            gsis = append(gsis, gsiArg)
            gsiNames = append(gsiNames, pulumi.String(g.GetIndexName()))
        }
        tableArgs.GlobalSecondaryIndexes = gsis
    }

    // 3.5  Local Secondary Indexes (LSIs)
    var lsiNames pulumi.StringArray
    if len(spec.GetLocalSecondaryIndexes()) > 0 {
        var lsis dynamodb.TableLocalSecondaryIndexArray
        for _, l := range spec.GetLocalSecondaryIndexes() {
            lsiArg := dynamodb.TableLocalSecondaryIndexArgs{
                Name:           pulumi.String(l.GetIndexName()),
                ProjectionType: pulumi.String(convertProjectionType(l.GetProjection().GetProjectionType())),
            }
            // Key schema – for LSIs HASH key is identical to table's, we only
            // need the RANGE key.
            var rangeKey string
            for _, ks := range l.GetKeySchema() {
                if ks.GetKeyType() == awsdynamodbv1.KeyType_RANGE {
                    rangeKey = ks.GetAttributeName()
                }
            }
            if rangeKey == "" {
                return errors.Errorf("lsi %s must define a RANGE key", l.GetIndexName())
            }
            lsiArg.RangeKey = pulumi.String(rangeKey)
            lsis = append(lsis, lsiArg)
            lsiNames = append(lsiNames, pulumi.String(l.GetIndexName()))
        }
        tableArgs.LocalSecondaryIndexes = lsis
    }

    // 3.6  Streams configuration
    if spec.GetStreamSpecification() != nil && spec.GetStreamSpecification().GetStreamEnabled() {
        tableArgs.StreamEnabled = pulumi.BoolPtr(true)
        tableArgs.StreamViewType = pulumi.StringPtr(convertStreamViewType(spec.GetStreamSpecification().GetStreamViewType()))
    }

    // 3.7  Time-to-Live
    if spec.GetTtlSpecification() != nil && spec.GetTtlSpecification().GetTtlEnabled() {
        tableArgs.Ttl = dynamodb.TableTtlArgs{
            AttributeName: pulumi.String(spec.GetTtlSpecification().GetAttributeName()),
            Enabled:       pulumi.Bool(true),
        }
    }

    // 3.8  Point-in-time recovery
    if spec.GetPointInTimeRecoveryEnabled() {
        tableArgs.PointInTimeRecovery = dynamodb.TablePointInTimeRecoveryArgs{
            Enabled: pulumi.Bool(true),
        }
    }

    // 3.9  Server-side encryption
    if spec.GetSseSpecification() != nil && spec.GetSseSpecification().GetEnabled() {
        sse := dynamodb.TableServerSideEncryptionArgs{
            Enabled: pulumi.Bool(true),
        }
        if spec.GetSseSpecification().GetSseType() == awsdynamodbv1.SSEType_KMS {
            sse.KmsKeyArn = pulumi.StringPtr(spec.GetSseSpecification().GetKmsMasterKeyId())
        }
        tableArgs.ServerSideEncryption = sse
    }

    // ---------------------------------------------------------------------
    // 4.  Create the DynamoDB Table resource
    // ---------------------------------------------------------------------
    table, err := dynamodb.NewTable(ctx, fmt.Sprintf("%s-table", spec.GetTableName()), tableArgs, pulumi.Provider(classicProvider))
    if err != nil {
        return errors.Wrap(err, "failed to create aws_dynamodb table")
    }

    // ---------------------------------------------------------------------
    // 5.  Export stack outputs as defined in the StackOutputs proto
    // ---------------------------------------------------------------------

    ctx.Export(OpTableArn, table.Arn)
    ctx.Export(OpTableName, table.Name)
    ctx.Export(OpTableId, table.ID().ToStringOutput())

    // Streams (only present when enabled)
    ctx.Export(OpStreamStreamArn, table.StreamArn)
    ctx.Export(OpStreamStreamLabel, table.StreamLabel)

    // KMS CMK ARN (when SSE uses KMS)
    if spec.GetSseSpecification() != nil && spec.GetSseSpecification().GetEnabled() && spec.GetSseSpecification().GetSseType() == awsdynamodbv1.SSEType_KMS {
        ctx.Export(OpKmsKeyArn, pulumi.String(spec.GetSseSpecification().GetKmsMasterKeyId()))
    }

    // Index names – these are static & known ahead of time
    if len(gsiNames) > 0 {
        ctx.Export(OpGlobalSecondaryIndexNames, gsiNames)
    }
    if len(lsiNames) > 0 {
        ctx.Export(OpLocalSecondaryIndexNames, lsiNames)
    }

    return nil
}

// ---------------------------------------------------------------------------
// Helper conversion utilities (proto → AWS string constants)
// ---------------------------------------------------------------------------

func convertAttributeType(t awsdynamodbv1.AttributeType) string {
    switch t {
    case awsdynamodbv1.AttributeType_STRING:
        return "S"
    case awsdynamodbv1.AttributeType_NUMBER:
        return "N"
    case awsdynamodbv1.AttributeType_BINARY:
        return "B"
    default:
        return "S" // fall-back, validation should have rejected unsupported values already
    }
}

func convertProjectionType(pt awsdynamodbv1.ProjectionType) string {
    switch pt {
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

func convertStreamViewType(vt awsdynamodbv1.StreamViewType) string {
    switch vt {
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
