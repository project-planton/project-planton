package module

import (
	"github.com/pkg/errors"
	awseksclusterv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/aws/awsekscluster/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awseksclusterv1.AwsEksClusterStackInput) (err error) {
	var provider *aws.Provider
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

	// Inputs
	target := stackInput.Target
	spec := target.Spec

	// Build subnet IDs input
	var subnetIds pulumi.StringArray
	for _, s := range spec.SubnetIds {
		subnetIds = append(subnetIds, pulumi.String(s.GetValue()))
	}

	// Control plane logs
	var logTypes pulumi.StringArray
	if spec.EnableControlPlaneLogs {
		logTypes = pulumi.StringArray{
			pulumi.String("api"), pulumi.String("audit"), pulumi.String("authenticator"), pulumi.String("controllerManager"), pulumi.String("scheduler"),
		}
	}

	clusterArgs := &eks.ClusterArgs{
		Name:    pulumi.String(target.Metadata.Name),
		RoleArn: pulumi.String(spec.ClusterRoleArn.GetValue()),
		Version: pulumi.String(spec.Version),
		VpcConfig: &eks.ClusterVpcConfigArgs{
			SubnetIds:            subnetIds,
			EndpointPublicAccess: pulumi.Bool(!spec.DisablePublicEndpoint),
			PublicAccessCidrs:    pulumi.ToStringArray(spec.PublicAccessCidrs),
		},
		EnabledClusterLogTypes: logTypes,
	}

	createdCluster, err := eks.NewCluster(ctx, target.Metadata.Name, clusterArgs, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "create EKS cluster")
	}

	// Export outputs aligned to AwsEksClusterStackOutputs
	ctx.Export(OpEndpoint, createdCluster.Endpoint)
	ctx.Export(OpClusterCaCertificate, createdCluster.CertificateAuthority.Data().Elem())
	ctx.Export(OpClusterSecurityGroupId, pulumi.String(""))
	ctx.Export(OpOidcIssuerUrl, pulumi.String(""))
	ctx.Export(OpClusterArn, createdCluster.Arn)
	ctx.Export(OpName, createdCluster.Name)

	return nil
}
