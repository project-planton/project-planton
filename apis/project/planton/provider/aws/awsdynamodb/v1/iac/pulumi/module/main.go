package module

import (
	"github.com/pkg/errors"
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources orchestrates DynamoDB table creation and exports outputs.
func Resources(ctx *pulumi.Context, in *awsdynamodbv1.AwsDynamodbStackInput) error {
	locals := initializeLocals(ctx, in)

	var provider *aws.Provider
	var err error

	if in.ProviderCredential == nil {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "create default AWS provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(in.ProviderCredential.AccessKeyId),
			SecretKey: pulumi.String(in.ProviderCredential.SecretAccessKey),
			Region:    pulumi.String(in.ProviderCredential.Region),
		})
		if err != nil {
			return errors.Wrap(err, "create AWS provider with custom credentials")
		}
	}

	tbl, err := createTable(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "dynamodb table")
	}

	// Export outputs mapping to AwsDynamodbStackOutputs
	ctx.Export(OpTableName, tbl.Table.Name)
	ctx.Export(OpTableArn, tbl.Table.Arn)
	ctx.Export(OpTableId, tbl.Table.ID())
	ctx.Export(OpStreamArn, tbl.Table.StreamArn)
	ctx.Export(OpStreamLabel, tbl.Table.StreamLabel)

	return nil
}
