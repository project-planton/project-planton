package module

import (
	"github.com/pkg/errors"
	awseksnodegroupv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awseksnodegroup/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the main entry point for the aws_eks_node_group Pulumi module.
func Resources(ctx *pulumi.Context, stackInput *awseksnodegroupv1.AwsEksNodeGroupStackInput) error {
	// Provider
	var provider *aws.Provider
	var err error
	if stackInput.ProviderCredential == nil {
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{})
	} else {
		cred := stackInput.ProviderCredential
		provider, err = aws.NewProvider(ctx, "classic-provider", &aws.ProviderArgs{
			Region:    pulumi.String(cred.Region),
			AccessKey: pulumi.String(cred.AccessKeyId),
			SecretKey: pulumi.String(cred.SecretAccessKey),
		})
	}
	if err != nil {
		return errors.Wrap(err, "create AWS provider")
	}

	target := stackInput.Target
	spec := target.Spec

	// Inputs
	var subnetIds pulumi.StringArray
	for _, s := range spec.SubnetIds {
		subnetIds = append(subnetIds, pulumi.String(s.GetValue()))
	}

	scaling := &eks.NodeGroupScalingConfigArgs{
		MinSize:     pulumi.Int(int(spec.Scaling.MinSize)),
		MaxSize:     pulumi.Int(int(spec.Scaling.MaxSize)),
		DesiredSize: pulumi.Int(int(spec.Scaling.DesiredSize)),
	}

	capacityType := pulumi.String("ON_DEMAND")
	if spec.CapacityType == awseksnodegroupv1.AwsEksNodeGroupCapacityType_spot {
		capacityType = pulumi.String("SPOT")
	}

	args := &eks.NodeGroupArgs{
		ClusterName:   pulumi.String(spec.ClusterName.GetValue()),
		NodeRoleArn:   pulumi.String(spec.NodeRoleArn.GetValue()),
		SubnetIds:     subnetIds,
		InstanceTypes: pulumi.ToStringArray([]string{spec.InstanceType}),
		ScalingConfig: scaling,
		CapacityType:  capacityType,
		DiskSize:      pulumi.Int(int(spec.DiskSizeGb)),
		Labels:        pulumi.ToStringMap(spec.Labels),
	}

	if spec.SshKeyName != "" {
		args.RemoteAccess = &eks.NodeGroupRemoteAccessArgs{Ec2SshKey: pulumi.String(spec.SshKeyName)}
	}

	created, err := eks.NewNodeGroup(ctx, target.Metadata.Name, args, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "create EKS node group")
	}

	// Exports (align to AwsEksNodeGroupStackOutputs)
	ctx.Export(OpNodeGroupName, created.NodeGroupName)
	ctx.Export(OpAsgName, pulumi.String(""))
	if spec.SshKeyName != "" {
		ctx.Export(OpRemoteAccessSgId, pulumi.String(""))
	}
	ctx.Export(OpInstanceProfileArn, pulumi.String(""))

	return nil
}
