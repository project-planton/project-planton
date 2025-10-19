package module

import (
	"github.com/pkg/errors"
	awsrdsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdscluster/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources orchestrates creation of AWS RDS Cluster related resources and exports outputs.
func Resources(ctx *pulumi.Context, stackInput *awsrdsclusterv1.AwsRdsClusterStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	var provider *aws.Provider
	var err error
	awsProviderConfig := stackInput.ProviderConfig

	if awsProviderConfig == nil {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			AccessKey: pulumi.String(awsProviderConfig.AccessKeyId),
			SecretKey: pulumi.String(awsProviderConfig.SecretAccessKey),
			Region:    pulumi.String(awsProviderConfig.GetRegion()),
			Token:     pulumi.StringPtr(awsProviderConfig.SessionToken),
		})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
	}

	// Security group (ingress from SGs and/or CIDRs)
	createdSg, err := securityGroup(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "security group")
	}

	// Subnet group (only when subnetIds provided and no name supplied)
	createdSubnetGroup, err := subnetGroup(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "subnet group")
	}

	// Cluster parameter group (when parameters provided or explicit family desired)
	createdParamGroup, err := clusterParameterGroup(ctx, locals, provider)
	if err != nil {
		return errors.Wrap(err, "cluster parameter group")
	}

	// Create the RDS Cluster
	cluster, err := rdsCluster(ctx, locals, provider, createdSg, createdSubnetGroup, createdParamGroup)
	if err != nil {
		return errors.Wrap(err, "rds cluster")
	}

	// Export outputs as defined in AwsRdsClusterStackOutputs
	ctx.Export(OpRdsClusterEndpoint, cluster.Endpoint)
	ctx.Export(OpRdsClusterReaderEndpoint, cluster.ReaderEndpoint)
	ctx.Export(OpRdsClusterId, cluster.ID())
	ctx.Export(OpRdsClusterArn, cluster.Arn)
	ctx.Export(OpRdsClusterEngine, cluster.Engine)
	ctx.Export(OpRdsClusterEngineVersion, cluster.EngineVersion)
	ctx.Export(OpRdsClusterPort, cluster.Port)
	if createdSubnetGroup != nil {
		ctx.Export(OpRdsSubnetGroup, createdSubnetGroup.Name)
	}
	if createdSg != nil {
		ctx.Export(OpRdsSecurityGroup, createdSg.Name)
	}
	if createdParamGroup != nil {
		ctx.Export(OpRdsClusterParameterGroup, createdParamGroup.Name)
	}
	return nil
}
