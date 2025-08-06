package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func lambdaFunction(ctx *pulumi.Context,
	locals *Locals,
	awsProvider *aws.Provider,
	createdIamRole *iam.Role) (*lambda.Function, error) {

	functionArgs := &lambda.FunctionArgs{
		Architectures:                pulumi.ToStringArray([]string{"x86_64"}),
		Description:                  pulumi.String(locals.AwsLambda.Spec.Function.Description),
		Name:                         pulumi.String(locals.AwsLambda.Metadata.Id),
		KmsKeyArn:                    pulumi.String(locals.AwsLambda.Spec.Function.KmsKeyArn),
		Layers:                       pulumi.ToStringArray(locals.AwsLambda.Spec.Function.Layers),
		Publish:                      pulumi.Bool(locals.AwsLambda.Spec.Function.Publish),
		ReservedConcurrentExecutions: pulumi.Int(locals.AwsLambda.Spec.Function.ReservedConcurrentExecutions),
		Role:                         createdIamRole.Arn,
		SourceCodeHash:               pulumi.String(locals.AwsLambda.Spec.Function.SourceCodeHash),
		Tags:                         pulumi.ToStringMap(locals.AwsLambda.Metadata.Labels),
		Timeout:                      pulumi.Int(3),
		PackageType:                  pulumi.String("Zip"),
		MemorySize:                   pulumi.Int(128),
		Environment: &lambda.FunctionEnvironmentArgs{
			Variables: pulumi.ToStringMap(locals.AwsLambda.Spec.Function.Variables),
		},
		EphemeralStorage: &lambda.FunctionEphemeralStorageArgs{
			Size: pulumi.Int(512),
		},
	}

	if locals.AwsLambda.Spec.Function.Handler != "" {
		functionArgs.Handler = pulumi.String(locals.AwsLambda.Spec.Function.Handler)
	}

	if locals.AwsLambda.Spec.Function.PackageType != "" {
		functionArgs.PackageType = pulumi.String(locals.AwsLambda.Spec.Function.PackageType)
	}

	if locals.AwsLambda.Spec.Function.TracingConfigMode != "" {
		functionArgs.TracingConfig = &lambda.FunctionTracingConfigArgs{
			Mode: pulumi.String(locals.AwsLambda.Spec.Function.TracingConfigMode),
		}
	}

	if locals.AwsLambda.Spec.Function.Runtime != "" {
		functionArgs.Runtime = pulumi.String(locals.AwsLambda.Spec.Function.Runtime)
	}

	if locals.AwsLambda.Spec.Function.Architectures != nil &&
		len(locals.AwsLambda.Spec.Function.Architectures) > 0 {
		functionArgs.Architectures = pulumi.ToStringArray(locals.AwsLambda.Spec.Function.Architectures)
	}

	if locals.AwsLambda.Spec.Function.Timeout > 0 {
		functionArgs.Timeout = pulumi.Int(locals.AwsLambda.Spec.Function.Timeout)
	}

	if locals.AwsLambda.Spec.Function.MemorySize > 0 {
		functionArgs.MemorySize = pulumi.Int(locals.AwsLambda.Spec.Function.MemorySize)
	}

	if locals.AwsLambda.Spec.Function.EphemeralStorageSize > 0 {
		functionArgs.EphemeralStorage = &lambda.FunctionEphemeralStorageArgs{
			Size: pulumi.Int(locals.AwsLambda.Spec.Function.EphemeralStorageSize),
		}
	}

	if locals.AwsLambda.Spec.Function.ImageUri != "" &&
		locals.AwsLambda.Spec.Function.S3Bucket == "" &&
		locals.AwsLambda.Spec.Function.FileSystemConfig == nil {
		functionArgs.ImageUri = pulumi.String(locals.AwsLambda.Spec.Function.ImageUri)
	}

	if locals.AwsLambda.Spec.Function.S3Bucket != "" &&
		locals.AwsLambda.Spec.Function.ImageUri == "" &&
		locals.AwsLambda.Spec.Function.FileSystemConfig == nil {
		functionArgs.S3Bucket = pulumi.String(locals.AwsLambda.Spec.Function.S3Bucket)
		functionArgs.S3Key = pulumi.String(locals.AwsLambda.Spec.Function.S3Key)
		functionArgs.S3ObjectVersion = pulumi.String(locals.AwsLambda.Spec.Function.S3ObjectVersion)
	}

	if locals.AwsLambda.Spec.Function.FileSystemConfig != nil &&
		locals.AwsLambda.Spec.Function.S3Bucket == "" &&
		locals.AwsLambda.Spec.Function.ImageUri == "" {
		functionArgs.FileSystemConfig = &lambda.FunctionFileSystemConfigArgs{
			Arn:            pulumi.String(locals.AwsLambda.Spec.Function.FileSystemConfig.Arn),
			LocalMountPath: pulumi.String(locals.AwsLambda.Spec.Function.FileSystemConfig.LocalMountPath),
		}
	}

	if locals.AwsLambda.Spec.Function.VpcConfig != nil {
		functionArgs.VpcConfig = &lambda.FunctionVpcConfigArgs{
			SecurityGroupIds: pulumi.ToStringArray(locals.AwsLambda.Spec.Function.VpcConfig.SecurityGroupIds),
			SubnetIds:        pulumi.ToStringArray(locals.AwsLambda.Spec.Function.VpcConfig.SubnetIds),
			VpcId:            pulumi.String(locals.AwsLambda.Spec.Function.VpcConfig.VpcId),
		}
	}

	if locals.AwsLambda.Spec.Function.ImageConfig != nil {
		functionArgs.ImageConfig = &lambda.FunctionImageConfigArgs{
			Commands:         pulumi.ToStringArray(locals.AwsLambda.Spec.Function.ImageConfig.Commands),
			EntryPoints:      pulumi.ToStringArray(locals.AwsLambda.Spec.Function.ImageConfig.EntryPoints),
			WorkingDirectory: pulumi.String(locals.AwsLambda.Spec.Function.ImageConfig.WorkingDirectory),
		}
	}

	if locals.AwsLambda.Spec.Function.DeadLetterConfigTargetArn != "" {
		functionArgs.DeadLetterConfig = &lambda.FunctionDeadLetterConfigArgs{
			TargetArn: pulumi.String(locals.AwsLambda.Spec.Function.DeadLetterConfigTargetArn),
		}
	}

	createdLambdaFunction, err := lambda.NewFunction(ctx,
		locals.AwsLambda.Metadata.Id,
		functionArgs,
		pulumi.Provider(awsProvider), pulumi.DependsOn([]pulumi.Resource{createdIamRole}))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create lambda function")
	}
	return createdLambdaFunction, nil
}
