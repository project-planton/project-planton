package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func clusterParameterGroup(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*rds.ClusterParameterGroup, error) {
	var clusterParameterGroupParameterArray = rds.ClusterParameterGroupParameterArray{}
	for _, parameter := range locals.AwsRdsCluster.Spec.ClusterParameters {
		clusterParameterGroupParameterArray = append(clusterParameterGroupParameterArray, &rds.ClusterParameterGroupParameterArgs{
			ApplyMethod: pulumi.String(parameter.ApplyMethod),
			Name:        pulumi.String(parameter.Name),
			Value:       pulumi.String(parameter.Value),
		})
	}

	clusterParameterGroupArgs := &rds.ClusterParameterGroupArgs{
		NamePrefix: pulumi.Sprintf("%s-", locals.AwsRdsCluster.Metadata.Id),
		Family:     pulumi.String(locals.AwsRdsCluster.Spec.ClusterFamily),
		Tags:       pulumi.ToStringMap(locals.Labels),
		Parameters: clusterParameterGroupParameterArray,
	}
	// Create rds cluster parameter group
	rdsClusterParameterGroup, err := rds.NewClusterParameterGroup(ctx, "rds-cluster-parameter-group", clusterParameterGroupArgs, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create rds cluster parameter group")
	}

	return rdsClusterParameterGroup, nil
}
