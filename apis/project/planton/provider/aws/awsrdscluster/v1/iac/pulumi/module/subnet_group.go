package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func subnetGroup(ctx *pulumi.Context, locals *Locals, awsProvider *aws.Provider) (*rds.SubnetGroup, error) {
	subnetGroup, err := rds.NewSubnetGroup(ctx, "default", &rds.SubnetGroupArgs{
		Name:      pulumi.String(locals.AwsRdsCluster.Metadata.Id),
		SubnetIds: pulumi.ToStringArray(locals.AwsRdsCluster.Spec.SubnetIds),
		Tags:      pulumi.ToStringMap(locals.Labels),
	}, pulumi.Provider(awsProvider))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create subnet group")
	}

	return subnetGroup, nil
}
