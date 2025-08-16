package module

import (
	"github.com/pkg/errors"
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the primary entry point for the aws_dynamodb_table Pulumi module.
// It reads the AwsDynamodbStackInput, sets up AWS credentials if provided,
// and delegates to the dynamodbTable() function to create the resource.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *aws.Provider
	var err error
	awsCredential := stackInput.ProviderCredential

	// If no credential is provided, use the default AWS provider
	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		// Create a custom provider with explicit credentials
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Create the AWS DynamoDB table resource
	if err := dynamodbTable(ctx, locals, provider); err != nil {
		return errors.Wrap(err, "failed to create aws_dynamodb_table resource")
	}

	return nil
}

// dynamodbTable creates the DynamoDB table with the specified configuration
func dynamodbTable(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	spec := locals.AwsDynamodb.Spec

	// Build the key schema
	var keySchema []dynamodb.TableAttributeArgs
	keySchema = append(keySchema, dynamodb.TableAttributeArgs{
		Name: pulumi.String(spec.PartitionKeyName),
		Type: pulumi.String(attributeTypeToString(spec.PartitionKeyType)),
	})

	// Add sort key if specified
	if spec.SortKeyName != "" {
		keySchema = append(keySchema, dynamodb.TableAttributeArgs{
			Name: pulumi.String(spec.SortKeyName),
			Type: pulumi.String(attributeTypeToString(spec.SortKeyType)),
		})
	}

	// Build table args
	tableArgs := &dynamodb.TableArgs{
		Name:        pulumi.String(spec.TableName),
		HashKey:     pulumi.String(spec.PartitionKeyName),
		BillingMode: pulumi.String(billingModeToString(spec.BillingMode)),
		PointInTimeRecovery: &dynamodb.TablePointInTimeRecoveryArgs{
			Enabled: pulumi.Bool(spec.PointInTimeRecoveryEnabled),
		},
		ServerSideEncryption: &dynamodb.TableServerSideEncryptionArgs{
			Enabled: pulumi.Bool(spec.ServerSideEncryptionEnabled),
		},
	}

	// Add sort key if specified
	if spec.SortKeyName != "" {
		tableArgs.RangeKey = pulumi.String(spec.SortKeyName)
	}

	// Add capacity configuration for provisioned billing mode
	if spec.BillingMode == awsdynamodbv1.BillingMode_BILLING_MODE_PROVISIONED {
		tableArgs.ReadCapacity = pulumi.Int(spec.ReadCapacityUnits)
		tableArgs.WriteCapacity = pulumi.Int(spec.WriteCapacityUnits)
	}

	// Create the DynamoDB table
	createdTable, err := dynamodb.NewTable(ctx,
		locals.AwsDynamodb.Metadata.Name,
		tableArgs,
		pulumi.Provider(provider),
	)
	if err != nil {
		return errors.Wrap(err, "failed to create DynamoDB table")
	}

	// Export outputs
	ctx.Export(OpTableArn, createdTable.Arn)
	ctx.Export(OpTableId, createdTable.ID())
	ctx.Export(OpTableName, createdTable.Name)
	ctx.Export(OpAwsRegion, pulumi.String(spec.AwsRegion))

	// Export stream ARN if point-in-time recovery is enabled
	if spec.PointInTimeRecoveryEnabled {
		ctx.Export(OpStreamArn, createdTable.StreamArn)
	}

	return nil
}

// attributeTypeToString converts the protobuf AttributeType enum to DynamoDB string format
func attributeTypeToString(attrType awsdynamodbv1.AttributeType) string {
	switch attrType {
	case awsdynamodbv1.AttributeType_ATTRIBUTE_TYPE_STRING:
		return "S"
	case awsdynamodbv1.AttributeType_ATTRIBUTE_TYPE_NUMBER:
		return "N"
	case awsdynamodbv1.AttributeType_ATTRIBUTE_TYPE_BINARY:
		return "B"
	default:
		return "S" // Default to string
	}
}

// billingModeToString converts the protobuf BillingMode enum to DynamoDB string format
func billingModeToString(billingMode awsdynamodbv1.BillingMode) string {
	switch billingMode {
	case awsdynamodbv1.BillingMode_BILLING_MODE_PROVISIONED:
		return "PROVISIONED"
	case awsdynamodbv1.BillingMode_BILLING_MODE_PAY_PER_REQUEST:
		return "PAY_PER_REQUEST"
	default:
		return "PAY_PER_REQUEST" // Default to pay-per-request
	}
}
