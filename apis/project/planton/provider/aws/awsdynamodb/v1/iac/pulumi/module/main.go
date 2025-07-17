// main.go
//
// Pulumi program that materialises an Amazon DynamoDB table from an
// awsdynamodbpb.AwsDynamodbSpec specification and exports stack outputs
// matching awsdynamodbpb.AwsDynamodbStackOutputs.
//
// The program expects a Pulumi configuration key called "spec" that
// contains the spec in JSON-encoded form (generated from the protobuf
// definition).  A higher-level orchestration layer can serialise the
// protobuf message to JSON and pass it through Pulumi configuration, but
// the Resources function can just as well be called directly when the
// spec is already in memory.
package main

import (
	"encoding/json"
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/encoding/protojson"

	awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// The spec is supplied via Pulumi configuration (JSON encoded).
		cfg := pulumi.Config(ctx)
		specJSON := cfg.Get("spec")
		if specJSON == "" {
			return fmt.Errorf("pulumi config key \"spec\" must be provided and contain AwsDynamodbSpec as JSON")
		}

		var spec awsdynamodbpb.AwsDynamodbSpec
		// Allow unknown fields so that newer controller/CLI versions can still
		// talk to older programs.
		unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
		if err := unmarshaler.Unmarshal([]byte(specJSON), &spec); err != nil {
			return fmt.Errorf("failed to unmarshal AwsDynamodbSpec JSON: %w", err)
		}

		return Resources(ctx, &spec)
	})
}

// Resources turns an AwsDynamodbSpec into concrete AWS resources and exports
// an AwsDynamodbStackOutputs-compatible set of stack outputs.
func Resources(ctx *pulumi.Context, in *awsdynamodbpb.AwsDynamodbSpec) error {
	// --- Attribute definitions -------------------------------------------------
	attributes := dynamodb.TableAttributeArray{}
	for _, ad := range in.AttributeDefinitions {
		attributes = append(attributes, dynamodb.TableAttributeArgs{
			Name: pulumi.String(ad.AttributeName),
			Type: pulumi.String(attributeTypeToString(ad.AttributeType)),
		})
	}

	// --- Primary key (hash/range) ---------------------------------------------
	hashKey, rangeKey, err := translateKeySchema(in.KeySchema)
	if err != nil {
		return err
	}

	// --- Billing / capacity ----------------------------------------------------
	var (
		billingMode                       *string
		readCap, writeCap                pulumi.IntPtrInput
		usingProvisionedThroughput       bool
	)

	switch in.BillingMode {
	case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
		bm := "PAY_PER_REQUEST"
		billingMode = &bm
		usingProvisionedThroughput = false
	case awsdynamodbpb.BillingMode_PROVISIONED:
		bm := "PROVISIONED"
		billingMode = &bm
		usingProvisionedThroughput = true
		if in.ProvisionedThroughput == nil {
			return fmt.Errorf("billing_mode is PROVISIONED but provisioned_throughput is nil – spec validation should have caught this")
		}
		readCap = pulumi.IntPtr(in.ProvisionedThroughput.ReadCapacityUnits)
		writeCap = pulumi.IntPtr(in.ProvisionedThroughput.WriteCapacityUnits)
	default:
		return fmt.Errorf("unsupported billing_mode %v", in.BillingMode)
	}

	// --- Global secondary indexes ---------------------------------------------
	gsiArray := dynamodb.TableGlobalSecondaryIndexArray{}
	for _, gsi := range in.GlobalSecondaryIndexes {
		gHash, gRange, err := translateKeySchema(gsi.KeySchema)
		if err != nil {
			return fmt.Errorf("in GSI %q: %w", gsi.IndexName, err)
		}

		gsiArgs := dynamodb.TableGlobalSecondaryIndexArgs{
			Name:            pulumi.String(gsi.IndexName),
			HashKey:         pulumi.String(gHash),
			ProjectionType:  pulumi.String(projectionTypeToString(gsi.Projection.ProjectionType)),
			NonKeyAttributes: stringSliceToPulumiStringArray(gsi.Projection.NonKeyAttributes),
		}
		if gRange != nil {
			gsiArgs.RangeKey = pulumi.StringPtr(*gRange)
		}

		if usingProvisionedThroughput && gsi.ProvisionedThroughput != nil {
			gsiArgs.ReadCapacity = pulumi.IntPtr(gsi.ProvisionedThroughput.ReadCapacityUnits)
			gsiArgs.WriteCapacity = pulumi.IntPtr(gsi.ProvisionedThroughput.WriteCapacityUnits)
		}

		gsiArray = append(gsiArray, gsiArgs)
	}

	// --- Local secondary indexes ----------------------------------------------
	lsiArray := dynamodb.TableLocalSecondaryIndexArray{}
	for _, lsi := range in.LocalSecondaryIndexes {
		// LSIs must share HASH with the base table; translateKeySchema will still
		// validate presence of HASH, but we ignore it afterwards.
		_, lRange, err := translateKeySchema(lsi.KeySchema)
		if err != nil {
			return fmt.Errorf("in LSI %q: %w", lsi.IndexName, err)
		}
		if lRange == nil {
			return fmt.Errorf("LSI %q must have a RANGE key (spec validation bug)", lsi.IndexName)
		}

		lsiArray = append(lsiArray, dynamodb.TableLocalSecondaryIndexArgs{
			Name:            pulumi.String(lsi.IndexName),
			RangeKey:        pulumi.StringPtr(*lRange),
			ProjectionType:  pulumi.String(projectionTypeToString(lsi.Projection.ProjectionType)),
			NonKeyAttributes: stringSliceToPulumiStringArray(lsi.Projection.NonKeyAttributes),
		})
	}

	// --- TTL -------------------------------------------------------------------
	var ttlArgs *dynamodb.TableTtlArgs
	if in.TtlSpecification != nil && in.TtlSpecification.TtlEnabled {
		ttlArgs = &dynamodb.TableTtlArgs{
			Enabled:       pulumi.Bool(true),
			AttributeName: pulumi.String(in.TtlSpecification.AttributeName),
		}
	}

	// --- Streams ---------------------------------------------------------------
	var (
		streamEnabled   pulumi.BoolPtrInput
		streamViewType  pulumi.StringPtrInput
	)
	if in.StreamSpecification != nil && in.StreamSpecification.StreamEnabled {
		streamEnabled = pulumi.BoolPtr(true)
		streamViewType = pulumi.StringPtr(streamViewTypeToString(in.StreamSpecification.StreamViewType))
	}

	// --- Server-side encryption -------------------------------------------------
	var sseArgs *dynamodb.TableServerSideEncryptionArgs
	if in.SseSpecification != nil && in.SseSpecification.Enabled {
		sseArgs = &dynamodb.TableServerSideEncryptionArgs{
			Enabled:  pulumi.BoolPtr(true),
			SseType:  pulumi.StringPtr(sseTypeToString(in.SseSpecification.SseType)),
		}
		if in.SseSpecification.SseType == awsdynamodbpb.SSEType_KMS {
			sseArgs.KmsKeyArn = pulumi.StringPtr(in.SseSpecification.KmsMasterKeyId)
		}
	}

	// --- Tags ------------------------------------------------------------------
	tags := pulumi.StringMap{}
	for k, v := range in.Tags {
		tags[k] = pulumi.String(v)
	}

	// --- Create the table ------------------------------------------------------
	table, err := dynamodb.NewTable(ctx, in.TableName, &dynamodb.TableArgs{
		Attributes:                 attributes,
		HashKey:                    pulumi.String(hashKey),
		RangeKey:                   stringPtrInput(rangeKey),
		BillingMode:                pulumi.StringPtrInput(pulumi.StringPtr(*billingMode)),
		ReadCapacity:               readCap,
		WriteCapacity:              writeCap,
		GlobalSecondaryIndexes:     gsiArray,
		LocalSecondaryIndexes:      lsiArray,
		Ttl:                        ttlArgs,
		Tags:                       tags,
		StreamEnabled:              streamEnabled,
		StreamViewType:             streamViewType,
		ServerSideEncryption:       sseArgs,
		PointInTimeRecoveryEnabled: pulumi.BoolPtr(in.PointInTimeRecoveryEnabled),
	})
	if err != nil {
		return err
	}

	// --- Stack exports ---------------------------------------------------------
	ctx.Export("table_arn", table.Arn)
	ctx.Export("table_name", table.Name)
	ctx.Export("table_id", table.ID()) // Pulumi ID (same as table ID in AWS).

	ctx.Export("global_secondary_index_names", table.GlobalSecondaryIndexes.ApplyT(func(gs interface{}) []string {
		// The provider returns []dynamodb.TableGlobalSecondaryIndexState – we only need names.
		var names []string
		if arr, ok := gs.([]dynamodb.TableGlobalSecondaryIndexState); ok {
			for _, g := range arr {
				if g.Name != nil {
					names = append(names, *g.Name)
				}
			}
		}
		return names
	}).(pulumi.StringArrayOutput))

	ctx.Export("local_secondary_index_names", table.LocalSecondaryIndexes.ApplyT(func(ls interface{}) []string {
		var names []string
		if arr, ok := ls.([]dynamodb.TableLocalSecondaryIndexState); ok {
			for _, l := range arr {
				if l.Name != nil {
					names = append(names, *l.Name)
				}
			}
		}
		return names
	}).(pulumi.StringArrayOutput))

	if in.StreamSpecification != nil && in.StreamSpecification.StreamEnabled {
		ctx.Export("stream", pulumi.All(table.StreamArn, table.StreamLabel).ApplyT(func(vals []interface{}) map[string]string {
			return map[string]string{
				"stream_arn":   vals[0].(string),
				"stream_label": vals[1].(string),
			}
		}))
	}

	if in.SseSpecification != nil && in.SseSpecification.Enabled && in.SseSpecification.SseType == awsdynamodbpb.SSEType_KMS {
		ctx.Export("kms_key_arn", pulumi.String(in.SseSpecification.KmsMasterKeyId))
	}

	return nil
}

// -------------------------- helper functions ---------------------------------

func attributeTypeToString(t awsdynamodbpb.AttributeType) string {
	switch t {
	case awsdynamodbpb.AttributeType_STRING:
		return "S"
	case awsdynamodbpb.AttributeType_NUMBER:
		return "N"
	case awsdynamodbpb.AttributeType_BINARY:
		return "B"
	default:
		return "S" // default to string; spec validation prevents UNKNOWN.
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

func streamViewTypeToString(s awsdynamodbpb.StreamViewType) string {
	switch s {
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

func sseTypeToString(t awsdynamodbpb.SSEType) string {
	switch t {
	case awsdynamodbpb.SSEType_AES256:
		return "AES256"
	case awsdynamodbpb.SSEType_KMS:
		return "KMS"
	default:
		return "AES256"
	}
}

// translateKeySchema returns the HASH key (mandatory) and an optional RANGE
// key (may be nil). It also validates that only one HASH and at most one RANGE
// are present.
func translateKeySchema(schema []*awsdynamodbpb.KeySchemaElement) (string, *string, error) {
	var hash string
	var rangeKey *string
	for _, ks := range schema {
		if ks.KeyType == awsdynamodbpb.KeyType_HASH {
			if hash != "" {
				return "", nil, fmt.Errorf("multiple HASH keys specified in key_schema")
			}
			hash = ks.AttributeName
		} else if ks.KeyType == awsdynamodbpb.KeyType_RANGE {
			if rangeKey != nil {
				return "", nil, fmt.Errorf("multiple RANGE keys specified in key_schema")
			}
			v := ks.AttributeName
			rangeKey = &v
		}
	}
	if hash == "" {
		return "", nil, fmt.Errorf("no HASH key defined in key_schema")
	}
	return hash, rangeKey, nil
}

func stringSliceToPulumiStringArray(ss []string) pulumi.StringArray {
	arr := pulumi.StringArray{}
	for _, s := range ss {
		arr = append(arr, pulumi.String(s))
	}
	return arr
}

// stringPtrInput converts *string into pulumi.StringPtrInput.
func stringPtrInput(s *string) pulumi.StringPtrInput {
	if s == nil {
		return nil
	}
	return pulumi.StringPtr(*s)
}

// The json.Marshal helpers come in handy when inspecting stack exports during
// development/debugging but are not required at runtime. They are kept here
// commented out for quick use when needed.
// func debugExport(ctx *pulumi.Context, name string, v interface{}) {
//     bs, _ := json.MarshalIndent(v, "", "  ")
//     ctx.Export(name, pulumi.String(string(bs)))
// }
