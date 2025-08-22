package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
    "github.com/pulumi/pulumi-aws/sdk/v7/go/aws/rds"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// subnetGroup creates a DB Subnet Group when subnetIds are provided and dbSubnetGroupName is not set.
func subnetGroup(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) (*rds.SubnetGroup, error) {
    spec := locals.AwsRdsCluster.Spec
    if spec == nil {
        return nil, nil
    }

    if (spec.DbSubnetGroupName != nil && spec.DbSubnetGroupName.GetValue() != "") || len(spec.SubnetIds) == 0 {
        return nil, nil
    }

    var subnetIds pulumi.StringArray
    for _, s := range spec.SubnetIds {
        if s.GetValue() != "" {
            subnetIds = append(subnetIds, pulumi.String(s.GetValue()))
        }
    }
    if len(subnetIds) == 0 {
        return nil, nil
    }

    sg, err := rds.NewSubnetGroup(ctx, "cluster-subnet-group", &rds.SubnetGroupArgs{
        Name:      pulumi.String(locals.AwsRdsCluster.Metadata.Id),
        SubnetIds: subnetIds,
        Tags:      pulumi.ToStringMap(locals.Labels),
    }, pulumi.Provider(provider))
    if err != nil {
        return nil, errors.Wrap(err, "create subnet group")
    }
    return sg, nil
}


