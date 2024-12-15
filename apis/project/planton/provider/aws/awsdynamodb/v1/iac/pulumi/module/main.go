package module

import (
	"github.com/pkg/errors"
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsdynamodbv1.AwsDynamodbStackInput) (err error) {
	locals := initializeLocals(ctx, stackInput)

	awsCredential := stackInput.AwsCredential

	var provider *aws.Provider

	if awsCredential == nil {
		//create aws provider using the credentials from the input
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create aws native provider")
		}
	} else {
		//create aws provider using the credentials from the input
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create aws native provider")
		}
	}

	createdDynamodbTable, err := table(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "failed to create dynamo table resources")
	}

	if err = autoScale(ctx, locals, provider, createdDynamodbTable); err != nil {
		return errors.Wrap(err, "failed to create dynamo db auto scaling resources")
	}
	return nil
}
