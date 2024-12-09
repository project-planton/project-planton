package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/cloudwatch"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func cloudwatchLogGroup(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*cloudwatch.LogGroup, error) {

	createdCloudwatchLogGroup, err := cloudwatch.NewLogGroup(ctx,
		"cloudwatch-log-group",
		&cloudwatch.LogGroupArgs{
			Name:            pulumi.Sprintf("/aws/lambda/%s", locals.AwsLambda.Metadata.Id),
			RetentionInDays: pulumi.Int(locals.AwsLambda.Spec.CloudwatchLogGroup.RetentionInDays),
		}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudwatch log group")
	}
	ctx.Export("cloudwatch-log-group-name", createdCloudwatchLogGroup.Name)
	return createdCloudwatchLogGroup, nil
}
