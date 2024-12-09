package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func invokeFunctionPermissions(ctx *pulumi.Context,
	locals *Locals,
	awsProvider *aws.Provider,
	createdLambdaFunction *lambda.Function) error {
	if locals.AwsLambda.Spec.InvokeFunctionPermissions == nil {
		return nil
	}

	for i, invokeFunctionPermission := range locals.AwsLambda.Spec.InvokeFunctionPermissions {
		_, err := lambda.NewPermission(ctx,
			fmt.Sprintf("invoke-permission-%d", i),
			&lambda.PermissionArgs{
				Action:    pulumi.String("lambda:InvokeFunction"),
				Function:  createdLambdaFunction.Name,
				Principal: pulumi.String(invokeFunctionPermission.Principal),
				SourceArn: pulumi.String(invokeFunctionPermission.SourceArn),
			}, pulumi.Provider(awsProvider), pulumi.Parent(createdLambdaFunction))
		if err != nil {
			return errors.Wrap(err, "failed to create new permission")
		}
	}
	return nil
}
