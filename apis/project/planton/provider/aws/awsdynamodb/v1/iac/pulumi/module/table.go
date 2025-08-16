package module

import (
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type tableResult struct {
	Table *dynamodb.Table
}

func createTable(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*tableResult, error) {
	// Map billing mode
	var billingMode pulumi.StringPtrInput
	switch locals.Spec.BillingMode {
	case awsdynamodbv1.AwsDynamodbSpec_BILLING_MODE_PROVISIONED:
		billingMode = pulumi.StringPtr("PROVISIONED")
	case awsdynamodbv1.AwsDynamodbSpec_BILLING_MODE_PAY_PER_REQUEST:
		billingMode = pulumi.StringPtr("PAY_PER_REQUEST")
	}

	// Attributes
	var attrDefs dynamodb.TableAttributeArray
	for _, a := range locals.Spec.AttributeDefinitions {
		attrType := pulumi.String("S")
		if a.Type == awsdynamodbv1.AwsDynamodbSpec_ATTRIBUTE_TYPE_N {
			attrType = pulumi.String("N")
		}
		if a.Type == awsdynamodbv1.AwsDynamodbSpec_ATTRIBUTE_TYPE_B {
			attrType = pulumi.String("B")
		}
		attrDefs = append(attrDefs, dynamodb.TableAttributeArgs{
			Name: pulumi.String(a.Name),
			Type: attrType,
		})
	}

	// Table keys
	var tableHashKey, tableRangeKey string
	for _, k := range locals.Spec.KeySchema {
		if k.KeyType == awsdynamodbv1.AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_HASH {
			tableHashKey = k.AttributeName
		}
		if k.KeyType == awsdynamodbv1.AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_RANGE {
			tableRangeKey = k.AttributeName
		}
	}

	// LSI
	var lsiArgs dynamodb.TableLocalSecondaryIndexArray
	for _, l := range locals.Spec.LocalSecondaryIndexes {
		var lsiRangeKey string
		for _, lk := range l.KeySchema {
			if lk.KeyType == awsdynamodbv1.AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_RANGE {
				lsiRangeKey = lk.AttributeName
			}
		}
		lsi := dynamodb.TableLocalSecondaryIndexArgs{
			Name:     pulumi.String(l.Name),
			RangeKey: pulumi.String(lsiRangeKey),
		}
		// Default projection ALL; override per spec
		lsi.ProjectionType = pulumi.String("ALL")
		if l.Projection != nil {
			switch l.Projection.Type {
			case awsdynamodbv1.AwsDynamodbSpec_PROJECTION_TYPE_KEYS_ONLY:
				lsi.ProjectionType = pulumi.String("KEYS_ONLY")
			case awsdynamodbv1.AwsDynamodbSpec_PROJECTION_TYPE_INCLUDE:
				lsi.ProjectionType = pulumi.String("INCLUDE")
				var nonKeys pulumi.StringArray
				for _, n := range l.Projection.NonKeyAttributes {
					nonKeys = append(nonKeys, pulumi.String(n))
				}
				lsi.NonKeyAttributes = nonKeys
			default:
				lsi.ProjectionType = pulumi.String("ALL")
			}
		}
		lsiArgs = append(lsiArgs, lsi)
	}

	// GSI
	var gsiArgs dynamodb.TableGlobalSecondaryIndexArray
	for _, g := range locals.Spec.GlobalSecondaryIndexes {
		var gsiHashKey, gsiRangeKey string
		for _, gk := range g.KeySchema {
			if gk.KeyType == awsdynamodbv1.AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_HASH {
				gsiHashKey = gk.AttributeName
			}
			if gk.KeyType == awsdynamodbv1.AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_RANGE {
				gsiRangeKey = gk.AttributeName
			}
		}
		gsi := dynamodb.TableGlobalSecondaryIndexArgs{
			Name:    pulumi.String(g.Name),
			HashKey: pulumi.String(gsiHashKey),
		}
		if gsiRangeKey != "" {
			gsi.RangeKey = pulumi.StringPtr(gsiRangeKey)
		}
		// Default projection ALL; override per spec
		gsi.ProjectionType = pulumi.String("ALL")
		if g.Projection != nil {
			switch g.Projection.Type {
			case awsdynamodbv1.AwsDynamodbSpec_PROJECTION_TYPE_KEYS_ONLY:
				gsi.ProjectionType = pulumi.String("KEYS_ONLY")
			case awsdynamodbv1.AwsDynamodbSpec_PROJECTION_TYPE_INCLUDE:
				gsi.ProjectionType = pulumi.String("INCLUDE")
				var nonKeys pulumi.StringArray
				for _, n := range g.Projection.NonKeyAttributes {
					nonKeys = append(nonKeys, pulumi.String(n))
				}
				gsi.NonKeyAttributes = nonKeys
			default:
				gsi.ProjectionType = pulumi.String("ALL")
			}
		}
		if g.ProvisionedThroughput != nil && locals.Spec.BillingMode == awsdynamodbv1.AwsDynamodbSpec_BILLING_MODE_PROVISIONED {
			gsi.ReadCapacity = pulumi.IntPtr(int(g.ProvisionedThroughput.ReadCapacityUnits))
			gsi.WriteCapacity = pulumi.IntPtr(int(g.ProvisionedThroughput.WriteCapacityUnits))
		}
		gsiArgs = append(gsiArgs, gsi)
	}

	// SSE
	var sseArgs *dynamodb.TableServerSideEncryptionArgs
	if locals.Spec.ServerSideEncryption != nil && locals.Spec.ServerSideEncryption.Enabled {
		sseArgs = &dynamodb.TableServerSideEncryptionArgs{
			Enabled:   pulumi.Bool(true),
			KmsKeyArn: pulumi.StringPtr(locals.Spec.ServerSideEncryption.KmsKeyArn),
		}
	}

	args := &dynamodb.TableArgs{
		Attributes:             attrDefs,
		HashKey:                pulumi.String(tableHashKey),
		BillingMode:            billingMode,
		LocalSecondaryIndexes:  lsiArgs,
		GlobalSecondaryIndexes: gsiArgs,
		ServerSideEncryption:   sseArgs,
		Tags:                   pulumi.ToStringMap(locals.AwsTags),
		StreamEnabled:          pulumi.BoolPtr(locals.Spec.StreamEnabled),
	}
	if tableRangeKey != "" {
		args.RangeKey = pulumi.StringPtr(tableRangeKey)
	}

	if locals.Spec.StreamEnabled {
		switch locals.Spec.StreamViewType {
		case awsdynamodbv1.AwsDynamodbSpec_STREAM_VIEW_TYPE_KEYS_ONLY:
			args.StreamViewType = pulumi.StringPtr("KEYS_ONLY")
		case awsdynamodbv1.AwsDynamodbSpec_STREAM_VIEW_TYPE_NEW_IMAGE:
			args.StreamViewType = pulumi.StringPtr("NEW_IMAGE")
		case awsdynamodbv1.AwsDynamodbSpec_STREAM_VIEW_TYPE_OLD_IMAGE:
			args.StreamViewType = pulumi.StringPtr("OLD_IMAGE")
		case awsdynamodbv1.AwsDynamodbSpec_STREAM_VIEW_TYPE_NEW_AND_OLD_IMAGES:
			args.StreamViewType = pulumi.StringPtr("NEW_AND_OLD_IMAGES")
		}
	}

	if locals.Spec.TableClass == awsdynamodbv1.AwsDynamodbSpec_TABLE_CLASS_STANDARD_INFREQUENT_ACCESS {
		args.TableClass = pulumi.StringPtr("STANDARD_INFREQUENT_ACCESS")
	}
	if locals.Spec.PointInTimeRecoveryEnabled {
		args.PointInTimeRecovery = &dynamodb.TablePointInTimeRecoveryArgs{Enabled: pulumi.Bool(true)}
	}
	if locals.Spec.BillingMode == awsdynamodbv1.AwsDynamodbSpec_BILLING_MODE_PROVISIONED && locals.Spec.ProvisionedThroughput != nil {
		args.ReadCapacity = pulumi.IntPtr(int(locals.Spec.ProvisionedThroughput.ReadCapacityUnits))
		args.WriteCapacity = pulumi.IntPtr(int(locals.Spec.ProvisionedThroughput.WriteCapacityUnits))
	}

	table, err := dynamodb.NewTable(ctx, locals.Target.Metadata.Name, args, pulumi.Provider(provider))
	if err != nil {
		return nil, err
	}

	return &tableResult{Table: table}, nil
}
