package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func parameterGroup(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*rds.ParameterGroup, error) {
	var parameterGroupParameterArray = rds.ParameterGroupParameterArray{}
	for _, parameter := range locals.AwsRdsInstance.Spec.Parameters {
		parameterGroupParameterArray = append(parameterGroupParameterArray, &rds.ParameterGroupParameterArgs{
			ApplyMethod: pulumi.String(parameter.ApplyMethod),
			Name:        pulumi.String(parameter.Name),
			Value:       pulumi.String(parameter.Value),
		})

	}

	parameterGroupArgs := &rds.ParameterGroupArgs{
		NamePrefix: pulumi.Sprintf("%s-", locals.AwsRdsInstance.Metadata.Id),
		Family:     pulumi.String(locals.AwsRdsInstance.Spec.DbParameterGroup),
		Tags:       pulumi.ToStringMap(locals.Labels),
		Parameters: parameterGroupParameterArray,
	}
	// Create RDS Parameter Group
	rdsParameterGroup, err := rds.NewParameterGroup(ctx, "rds-parameter-group", parameterGroupArgs, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rds parameter group")
	}

	return rdsParameterGroup, nil
}
