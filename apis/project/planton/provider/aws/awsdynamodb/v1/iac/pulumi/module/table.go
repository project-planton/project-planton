package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1/iac/pulumi/module/outputs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/dynamodb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func table(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*dynamodb.Table, error) {
	awsDynamodb := locals.AwsDynamodb

	// stream
	streamEnabled := awsDynamodb.Spec.EnableStreams
	streamViewType := ""
	if len(awsDynamodb.Spec.ReplicaRegionNames) > 0 {
		streamEnabled = true
	}

	if len(awsDynamodb.Spec.ReplicaRegionNames) > 0 || awsDynamodb.Spec.EnableStreams {
		streamViewType = awsDynamodb.Spec.StreamViewType
	}

	// capacity
	readCapacity := 0
	writeCapacity := 0
	if awsDynamodb.Spec.BillingMode == "PROVISIONED" {
		readCapacity = 5
		writeCapacity = 5
		if awsDynamodb.Spec.AutoScale != nil &&
			awsDynamodb.Spec.AutoScale.ReadCapacity != nil {
			readCapacity = int(awsDynamodb.Spec.AutoScale.ReadCapacity.MinCapacity)
		}
		if awsDynamodb.Spec.AutoScale != nil &&
			awsDynamodb.Spec.AutoScale.WriteCapacity != nil {
			writeCapacity = int(awsDynamodb.Spec.AutoScale.WriteCapacity.MinCapacity)
		}
	}

	// range key
	rangeKey := ""
	if awsDynamodb.Spec.RangeKey != nil {
		rangeKey = awsDynamodb.Spec.RangeKey.Name
	}

	// replicas
	var replicaArray = dynamodb.TableReplicaTypeArray{}
	for _, regionName := range awsDynamodb.Spec.ReplicaRegionNames {
		replicaArray = append(replicaArray, &dynamodb.TableReplicaTypeArgs{
			KmsKeyArn:           nil,
			PointInTimeRecovery: pulumi.Bool(false),
			PropagateTags:       pulumi.Bool(false),
			RegionName:          pulumi.String(regionName),
		})
	}

	// attributes
	var attributeArray = dynamodb.TableAttributeArray{}
	var attributeMap = make(map[string]bool)

	attributeMap[awsDynamodb.Spec.HashKey.Name] = true
	attributeArray = append(attributeArray, &dynamodb.TableAttributeArgs{
		Name: pulumi.String(awsDynamodb.Spec.HashKey.Name),
		Type: pulumi.String(awsDynamodb.Spec.HashKey.Type),
	})

	if awsDynamodb.Spec.RangeKey != nil && awsDynamodb.Spec.RangeKey.Name != "" {
		attributeMap[awsDynamodb.Spec.RangeKey.Name] = true
		attributeArray = append(attributeArray, &dynamodb.TableAttributeArgs{
			Name: pulumi.String(awsDynamodb.Spec.RangeKey.Name),
			Type: pulumi.String(awsDynamodb.Spec.RangeKey.Type),
		})
	}

	for _, attribute := range awsDynamodb.Spec.Attributes {
		_, exists := attributeMap[attribute.Name]
		if !exists {
			attributeMap[attribute.Name] = true
			attributeArray = append(attributeArray, &dynamodb.TableAttributeArgs{
				Name: pulumi.String(attribute.Name),
				Type: pulumi.String(attribute.Type),
			})
		}
	}

	// global secondary index
	var globalSecondaryIndexArray = dynamodb.TableGlobalSecondaryIndexArray{}
	for _, globalSecondaryIndex := range awsDynamodb.Spec.GlobalSecondaryIndexes {
		globalIndexReadCapacity := globalSecondaryIndex.ReadCapacity
		globalIndexWriteCapacity := globalSecondaryIndex.WriteCapacity
		if awsDynamodb.Spec.BillingMode == "PROVISIONED" && globalIndexReadCapacity == 0 {
			globalIndexReadCapacity = int32(readCapacity)
		}
		if awsDynamodb.Spec.BillingMode == "PROVISIONED" && globalIndexWriteCapacity == 0 {
			globalIndexWriteCapacity = int32(writeCapacity)
		}
		globalSecondaryIndexArray = append(globalSecondaryIndexArray, &dynamodb.TableGlobalSecondaryIndexArgs{
			Name:             pulumi.String(globalSecondaryIndex.Name),
			HashKey:          pulumi.String(globalSecondaryIndex.HashKey),
			RangeKey:         pulumi.String(globalSecondaryIndex.RangeKey),
			ReadCapacity:     pulumi.Int(globalIndexReadCapacity),
			WriteCapacity:    pulumi.Int(globalIndexWriteCapacity),
			ProjectionType:   pulumi.String(globalSecondaryIndex.ProjectionType),
			NonKeyAttributes: pulumi.ToStringArray(globalSecondaryIndex.NonKeyAttributes),
		})
	}

	// local secondary index
	var localSecondaryIndexArray = dynamodb.TableLocalSecondaryIndexArray{}
	for _, localSecondaryIndex := range awsDynamodb.Spec.LocalSecondaryIndexes {
		localSecondaryIndexArray = append(localSecondaryIndexArray, &dynamodb.TableLocalSecondaryIndexArgs{
			Name:             pulumi.String(localSecondaryIndex.Name),
			RangeKey:         pulumi.String(localSecondaryIndex.RangeKey),
			ProjectionType:   pulumi.String(localSecondaryIndex.ProjectionType),
			NonKeyAttributes: pulumi.ToStringArray(localSecondaryIndex.NonKeyAttributes),
		})
	}

	// server side encryption
	var serverSideEncryption *dynamodb.TableServerSideEncryptionArgs
	if awsDynamodb.Spec.ServerSideEncryption != nil {
		serverSideEncryption = &dynamodb.TableServerSideEncryptionArgs{
			Enabled:   pulumi.Bool(awsDynamodb.Spec.ServerSideEncryption.IsEnabled),
			KmsKeyArn: pulumi.StringPtr(awsDynamodb.Spec.ServerSideEncryption.KmsKeyArn),
		}
	}

	// point in time recovery
	var pointInTimeRecovery *dynamodb.TablePointInTimeRecoveryArgs
	if awsDynamodb.Spec.PointInTimeRecovery != nil {
		pointInTimeRecovery = &dynamodb.TablePointInTimeRecoveryArgs{
			Enabled: pulumi.Bool(awsDynamodb.Spec.PointInTimeRecovery.IsEnabled),
		}
	}

	// ttl
	var ttl *dynamodb.TableTtlArgs
	if awsDynamodb.Spec.Ttl != nil {
		ttl = &dynamodb.TableTtlArgs{
			Enabled:       pulumi.Bool(awsDynamodb.Spec.Ttl.IsEnabled),
			AttributeName: pulumi.String(awsDynamodb.Spec.Ttl.AttributeName),
		}
	}

	// import table
	var importTable *dynamodb.TableImportTableArgs
	if awsDynamodb.Spec.ImportTable != nil {
		inputFormatOptions := &dynamodb.TableImportTableInputFormatOptionsArgs{
			Csv: dynamodb.TableImportTableInputFormatOptionsCsvArgs{
				Delimiter:   pulumi.String(","),
				HeaderLists: pulumi.ToStringArray([]string{}),
			},
		}
		if awsDynamodb.Spec.ImportTable.InputFormatOptions != nil && awsDynamodb.Spec.ImportTable.InputFormatOptions.Csv != nil {
			inputFormatOptions = &dynamodb.TableImportTableInputFormatOptionsArgs{
				Csv: dynamodb.TableImportTableInputFormatOptionsCsvArgs{
					Delimiter:   pulumi.String(awsDynamodb.Spec.ImportTable.InputFormatOptions.Csv.Delimiter),
					HeaderLists: pulumi.ToStringArray(awsDynamodb.Spec.ImportTable.InputFormatOptions.Csv.Headers),
				},
			}
		}

		s3BucketSource := &dynamodb.TableImportTableS3BucketSourceArgs{}
		if awsDynamodb.Spec.ImportTable.S3BucketSource != nil {
			s3BucketSource = &dynamodb.TableImportTableS3BucketSourceArgs{
				Bucket:      pulumi.String(awsDynamodb.Spec.ImportTable.S3BucketSource.Bucket),
				BucketOwner: pulumi.String(awsDynamodb.Spec.ImportTable.S3BucketSource.BucketOwner),
				KeyPrefix:   pulumi.String(awsDynamodb.Spec.ImportTable.S3BucketSource.KeyPrefix),
			}
		}
		importTable = &dynamodb.TableImportTableArgs{
			InputCompressionType: pulumi.String(awsDynamodb.Spec.ImportTable.InputCompressionType),
			InputFormat:          pulumi.String(awsDynamodb.Spec.ImportTable.InputFormat),
			InputFormatOptions:   inputFormatOptions,
			S3BucketSource:       s3BucketSource,
		}
	}

	createdDynamodbTable, err := dynamodb.NewTable(ctx, awsDynamodb.Metadata.Name, &dynamodb.TableArgs{
		Name:                      pulumi.String(awsDynamodb.Spec.TableName),
		BillingMode:               pulumi.String(awsDynamodb.Spec.BillingMode),
		ReadCapacity:              pulumi.Int(readCapacity),
		WriteCapacity:             pulumi.Int(writeCapacity),
		HashKey:                   pulumi.String(awsDynamodb.Spec.HashKey.Name),
		RangeKey:                  pulumi.String(rangeKey),
		StreamEnabled:             pulumi.Bool(streamEnabled),
		StreamViewType:            pulumi.String(streamViewType),
		TableClass:                pulumi.String("STANDARD"),
		DeletionProtectionEnabled: pulumi.Bool(false),
		ServerSideEncryption:      serverSideEncryption,
		PointInTimeRecovery:       pointInTimeRecovery,
		Ttl:                       ttl,
		Tags:                      pulumi.ToStringMap(locals.Labels),
		Attributes:                attributeArray,
		GlobalSecondaryIndexes:    globalSecondaryIndexArray,
		LocalSecondaryIndexes:     localSecondaryIndexArray,
		Replicas:                  replicaArray,
		ImportTable:               importTable,
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create dynamo table resources")
	}

	ctx.Export(outputs.TableName, createdDynamodbTable.Name)
	ctx.Export(outputs.TableArn, createdDynamodbTable.Arn)
	ctx.Export(outputs.TableStreamArn, createdDynamodbTable.StreamArn)

	return createdDynamodbTable, nil
}
