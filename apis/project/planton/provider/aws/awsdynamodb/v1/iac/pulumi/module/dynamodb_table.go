package module

import (
	awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pkg/errors"

	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// createDynamoDBTable creates (or imports) a DynamoDB table from the provided
// spec and returns the Pulumi resource together with any error that may have
// occurred. All cross-cutting concerns (provider injection, tags, etc.) are
// handled by the caller so that this function deals only with translating the
// protobuf spec into Pulumi-native arguments.
func createDynamoDBTable(
	ctx *pulumi.Context,
	locals *Locals,
	provider *awsclassic.Provider,
) (*dynamodb.Table, error) {
	if locals == nil || locals.AwsDynamodb == nil {
		return nil, errors.New("locals or AwsDynamodb spec cannot be nil")
	}
	sp := locals.AwsDynamodb

	// --- Attribute definitions -------------------------------------------------
	var attributes dynamodb.TableAttributeArray
	for _, attr := range sp.AttributeDefinitions {
		typeStr, err := convertAttributeType(attr.AttributeType)
		if err != nil {
			return nil, errors.Wrap(err, "invalid attribute type")
		}
		attributes = append(attributes, dynamodb.TableAttributeArgs{
			Name: pulumi.String(attr.AttributeName),
			Type: pulumi.String(typeStr),
		})
	}

	// --- Primary key (hash & optional range) ----------------------------------
	if len(sp.KeySchema) == 0 {
		return nil, errors.New("key_schema must contain at least the HASH key")
	}
	hashKey := ""
	var rangeKey *string
	for _, ks := range sp.KeySchema {
		if ks.KeyType == awsdynamodbv1.KeyType_HASH {
			hashKey = ks.AttributeName
		} else if ks.KeyType == awsdynamodbv1.KeyType_RANGE {
			v := ks.AttributeName
			rangeKey = &v
		}
	}
	if hashKey == "" {
		return nil, errors.New("HASH key missing in key_schema")
	}

	// --- Billing & capacity ----------------------------------------------------
	billingMode := convertBillingMode(sp.BillingMode)
	var readCap, writeCap *int
	if sp.BillingMode == awsdynamodbv1.BillingMode_PROVISIONED {
		if sp.ProvisionedThroughput == nil {
			return nil, errors.New("provisioned_throughput required when billing mode is PROVISIONED")
		}
		rc := int(sp.ProvisionedThroughput.ReadCapacityUnits)
		wc := int(sp.ProvisionedThroughput.WriteCapacityUnits)
		readCap = &rc
		writeCap = &wc
	}

	// --- Global Secondary Indexes (GSIs) --------------------------------------
	var gsis dynamodb.TableGlobalSecondaryIndexArray
	for _, g := range sp.GlobalSecondaryIndexes {
		gsiAttr := dynamodb.TableGlobalSecondaryIndexArgs{
			Name: pulumi.String(g.IndexName),
		}

		// Keys inside GSI
		for _, ks := range g.KeySchema {
			if ks.KeyType == awsdynamodbv1.KeyType_HASH {
				gsiAttr.HashKey = pulumi.String(ks.AttributeName)
			} else if ks.KeyType == awsdynamodbv1.KeyType_RANGE {
				gsiAttr.RangeKey = pulumi.StringPtr(ks.AttributeName)
			}
		}

		// Projection
		gsiAttr.ProjectionType = pulumi.String(convertProjectionType(g.Projection.ProjectionType))
		if len(g.Projection.NonKeyAttributes) > 0 {
			var nk pulumi.StringArray
			for _, n := range g.Projection.NonKeyAttributes {
				nk = append(nk, pulumi.String(n))
			}
			gsiAttr.NonKeyAttributes = nk
		}

		// Capacity only when PROVISIONED
		if g.ProvisionedThroughput != nil {
			rc := int(g.ProvisionedThroughput.ReadCapacityUnits)
			wc := int(g.ProvisionedThroughput.WriteCapacityUnits)
			gsiAttr.ReadCapacity = pulumi.IntPtr(rc)
			gsiAttr.WriteCapacity = pulumi.IntPtr(wc)
		}

		gsis = append(gsis, gsiAttr)
	}

	// --- Local Secondary Indexes (LSIs) ---------------------------------------
	var lsis dynamodb.TableLocalSecondaryIndexArray
	for _, l := range sp.LocalSecondaryIndexes {
		lsiAttr := dynamodb.TableLocalSecondaryIndexArgs{
			Name: pulumi.String(l.IndexName),
		}

		for _, ks := range l.KeySchema {
			if ks.KeyType == awsdynamodbv1.KeyType_HASH {
				lsiAttr.HashKey = pulumi.String(ks.AttributeName)
			} else if ks.KeyType == awsdynamodbv1.KeyType_RANGE {
				lsiAttr.RangeKey = pulumi.String(ks.AttributeName)
			}
		}

		lsiAttr.ProjectionType = pulumi.String(convertProjectionType(l.Projection.ProjectionType))
		if len(l.Projection.NonKeyAttributes) > 0 {
			var nk pulumi.StringArray
			for _, n := range l.Projection.NonKeyAttributes {
				nk = append(nk, pulumi.String(n))
			}
			lsiAttr.NonKeyAttributes = nk
		}

		lsis = append(lsis, lsiAttr)
	}

	// --- Streams --------------------------------------------------------------
	var streamEnabled *bool
	var streamViewType *string
	if sp.StreamSpecification != nil && sp.StreamSpecification.StreamEnabled {
		b := true
		streamEnabled = &b
		vt := convertStreamViewType(sp.StreamSpecification.StreamViewType)
		streamViewType = &vt
	}

	// --- TTL ------------------------------------------------------------------
	var ttl dynamodb.TableTtlArgsPtrInput
	if sp.TtlSpecification != nil {
		ttl = dynamodb.TableTtlArgs{
			Enabled:       pulumi.Bool(sp.TtlSpecification.TtlEnabled),
			AttributeName: pulumi.String(sp.TtlSpecification.AttributeName),
		}.ToTableTtlArgsPtrOutput().(dynamodb.TableTtlArgsPtrInput)
	}

	// --- SSE ------------------------------------------------------------------
	var sse dynamodb.TableServerSideEncryptionArgsPtrInput
	if sp.SseSpecification != nil && sp.SseSpecification.Enabled {
		sse = dynamodb.TableServerSideEncryptionArgs{
			Enabled:    pulumi.Bool(true),
			KmsKeyArn:  pulumi.StringPtr(sp.SseSpecification.KmsMasterKeyId),
		}.ToTableServerSideEncryptionArgsPtrOutput().(dynamodb.TableServerSideEncryptionArgsPtrInput)
	}

	// --- Point in time recovery ----------------------------------------------
	var pitr dynamodb.TablePointInTimeRecoveryArgsPtrInput
	if sp.PointInTimeRecoveryEnabled {
		pitr = dynamodb.TablePointInTimeRecoveryArgs{
			Enabled: pulumi.Bool(true),
		}.ToTablePointInTimeRecoveryArgsPtrOutput().(dynamodb.TablePointInTimeRecoveryArgsPtrInput)
	}

	// --- Tags -----------------------------------------------------------------
	// Convert map[string]string -> pulumi.StringMap
	pulumiTags := pulumi.StringMap{}
	for k, v := range locals.Tags {
		pulumiTags[k] = pulumi.String(v)
	}

	// --- Assemble final arguments & create the resource -----------------------
	table, err := dynamodb.NewTable(ctx, "dynamodb-table", &dynamodb.TableArgs{
		Name:                    pulumi.StringPtr(sp.TableName),
		Attributes:              attributes,
		HashKey:                 pulumi.String(hashKey),
		RangeKey:                pulumi.StringPtrPtr(rangeKey),
		BillingMode:             pulumi.StringPtr(billingMode),
		ReadCapacity:            pulumi.IntPtrPtr(readCap),
		WriteCapacity:           pulumi.IntPtrPtr(writeCap),
		GlobalSecondaryIndexes:  gsis,
		LocalSecondaryIndexes:   lsis,
		StreamEnabled:           pulumi.BoolPtrPtr(streamEnabled),
		StreamViewType:          pulumi.StringPtrPtr(streamViewType),
		Ttl:                     ttl,
		ServerSideEncryption:    sse,
		PointInTimeRecovery:     pitr,
		Tags:                    pulumiTags,
	}, pulumi.Provider(provider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create dynamodb table")
	}

	return table, nil
}

// ---------------------------------------------------------------------------
// Helper conversion functions
// ---------------------------------------------------------------------------

func convertAttributeType(t awsdynamodbv1.AttributeType) (string, error) {
	switch t {
	case awsdynamodbv1.AttributeType_STRING:
		return "S", nil
	case awsdynamodbv1.AttributeType_NUMBER:
		return "N", nil
	case awsdynamodbv1.AttributeType_BINARY:
		return "B", nil
	default:
		return "", errors.Errorf("unsupported AttributeType %v", t)
	}
}

func convertBillingMode(m awsdynamodbv1.BillingMode) string {
	if m == awsdynamodbv1.BillingMode_PROVISIONED {
		return "PROVISIONED"
	}
	return "PAY_PER_REQUEST"
}

func convertProjectionType(p awsdynamodbv1.ProjectionType) string {
	switch p {
	case awsdynamodbv1.ProjectionType_ALL:
		return "ALL"
	case awsdynamodbv1.ProjectionType_KEYS_ONLY:
		return "KEYS_ONLY"
	case awsdynamodbv1.ProjectionType_INCLUDE:
		return "INCLUDE"
	default:
		return "ALL" // fallback
	}
}

func convertStreamViewType(t awsdynamodbv1.StreamViewType) string {
	switch t {
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
