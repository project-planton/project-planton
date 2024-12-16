package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awslambda/v1/iac/pulumi/module/outputs"
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
	ctx.Export(outputs.CLOUDWATCH_LOG_GROUP_NAME, createdCloudwatchLogGroup.Name)
	return createdCloudwatchLogGroup, nil
}
