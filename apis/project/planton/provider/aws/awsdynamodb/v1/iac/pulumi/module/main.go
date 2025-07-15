package module

import (
	awsnative "github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pkg/errors"

	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Resources is the main entrypoint called by the Pulumi program. It reconciles
// the protobuf-based spec into actual cloud resources and exports the
// identifiers requested by the *StackOutputs proto.
func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) error {
	// 1. Initialise locals -------------------------------------------------------------------
	locals := initializeLocals(stackInput)
	if locals.AwsDynamodb == nil {
		return errors.New("aws_dynamodb spec cannot be nil in stack input")
	}

	// 2. Configure AWS providers -------------------------------------------------------------
	awsCredential := stackInput.GetProviderCredential()
	var nativeProvider *awsnative.Provider
	var classicProvider *awsclassic.Provider
	var err error

	if awsCredential == nil {
		nativeProvider, err = awsnative.NewProvider(ctx, "native-provider", &awsnative.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS native provider")
		}
		classicProvider, err = awsclassic.NewProvider(ctx, "classic-provider", &awsclassic.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS classic provider")
		}
	} else {
		nativeProvider, err = awsnative.NewProvider(ctx, "native-provider", &awsnative.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS native provider with custom credentials")
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

	// 3. Create the DynamoDB table -----------------------------------------------------------
	table, err := createDynamoDBTable(ctx, locals, classicProvider)
	if err != nil {
		return err
	}

	// 4. Export the outputs -----------------------------------------------------------------
	ctx.Export(OpTableArn, table.Arn)
	ctx.Export(OpTableName, table.Name)
	ctx.Export(OpTableId, table.ID().ToStringOutput())

	// Streams
	ctx.Export(OpStreamStreamArn, table.StreamArn)
	ctx.Export(OpStreamStreamLabel, table.StreamLabel)

	// KMS Key ARN (if SSE with CMK)
	if locals.AwsDynamodb.SseSpecification != nil {
		ctx.Export(OpKmsKeyArn, pulumi.String(locals.AwsDynamodb.SseSpecification.KmsMasterKeyId))
	}

	// GSI / LSI names (plain string arrays)
	ctx.Export(OpGlobalSecondaryIndexNames, table.GlobalSecondaryIndexes.ApplyT(func(arr []dynamodb.TableGlobalSecondaryIndex) []string {
		var names []string
		for _, g := range arr {
			if g.Name != nil {
				names = append(names, *g.Name)
			}
		}
		return names
	}).(pulumi.StringArrayOutput))

	ctx.Export(OpLocalSecondaryIndexNames, table.LocalSecondaryIndexes.ApplyT(func(arr []dynamodb.TableLocalSecondaryIndex) []string {
		var names []string
		for _, l := range arr {
			if l.Name != nil {
				names = append(names, *l.Name)
			}
		}
		return names
	}).(pulumi.StringArrayOutput))

	return nil
}
