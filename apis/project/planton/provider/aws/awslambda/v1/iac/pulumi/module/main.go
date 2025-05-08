package module

import (
	"github.com/pkg/errors"
	awslambdav1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awslambda/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awslambdav1.AwsLambdaStackInput) error {
	locals := initializeLocals(ctx, stackInput)
	awsCredential := stackInput.ProviderCredential

	//create aws provider using the credentials from the input
	awsProvider, err := aws.NewProvider(ctx,
		"classic-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	createdIamRole, err := iamRole(ctx, locals, awsProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create iam role")
	}

	if stackInput.Target.Spec.CloudwatchLogGroup != nil {
		_, err := cloudwatchLogGroup(ctx, locals, awsProvider)
		if err != nil {
			return errors.Wrap(err, "failed to create cloud watch log group")
		}
	}

	createdLambdaFunction, err := lambdaFunction(ctx, locals, awsProvider, createdIamRole)
	if err != nil {
		return errors.Wrap(err, "failed to create lambda function")
	}

	err = invokeFunctionPermissions(ctx, locals, awsProvider, createdLambdaFunction)
	if err != nil {
		return errors.Wrap(err, "failed to create invoke function permissions")
	}

	ctx.Export(OpLambdaFunctionArn, createdLambdaFunction.Arn)
	ctx.Export(OpLambdaFunctionName, createdLambdaFunction.Name)
	ctx.Export(OpIamRoleName, createdIamRole.Name)

	return nil
}
