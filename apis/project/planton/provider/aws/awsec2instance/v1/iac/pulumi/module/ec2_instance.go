package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/valuefrom"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	awsec2instancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsec2instance/v1"
)

// ec2Instance creates a single AWS EC2 instance in a private subnet.
// It keeps logic linear and “Terraform‑like” for readability.
func ec2Instance(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	// --------‑‑‑ Validate connection method requirements ----------------
	switch locals.AwsEc2Instance.Spec.ConnectionMethod {
	case awsec2instancev1.AwsEc2InstanceConnectionMethod_SSM:
		if locals.AwsEc2Instance.Spec.IamInstanceProfileArn == nil ||
			locals.AwsEc2Instance.Spec.IamInstanceProfileArn.GetValue() == "" {
			return errors.New("iam_instance_profile_arn is required when connection_method = SSM")
		}
	case awsec2instancev1.AwsEc2InstanceConnectionMethod_BASTION,
		awsec2instancev1.AwsEc2InstanceConnectionMethod_INSTANCE_CONNECT:
		if locals.AwsEc2Instance.Spec.KeyName == "" {
			return errors.New("key_name is required when connection_method = BASTION or INSTANCE_CONNECT")
		}
	}

	// --------‑‑‑ Convert security‑group refs ---------------------------
	sgIDs := pulumi.ToStringArray(valuefrom.ToStringArray(locals.AwsEc2Instance.Spec.SecurityGroupIds))

	// --------‑‑‑ Root block device (always include – defaults to 30 GiB) -
	rootBlock := &ec2.InstanceRootBlockDeviceArgs{
		VolumeSize: pulumi.Int(int(locals.AwsEc2Instance.Spec.RootVolumeSizeGb)),
	}

	// --------‑‑‑ Assemble EC2 arguments --------------------------------
	instanceArgs := &ec2.InstanceArgs{
		Ami:                   pulumi.String(locals.AwsEc2Instance.Spec.AmiId),
		InstanceType:          pulumi.String(locals.AwsEc2Instance.Spec.InstanceType),
		SubnetId:              pulumi.String(locals.AwsEc2Instance.Spec.SubnetId.GetValue()),
		VpcSecurityGroupIds:   sgIDs,
		EbsOptimized:          pulumi.Bool(locals.AwsEc2Instance.Spec.EbsOptimized),
		DisableApiTermination: pulumi.Bool(locals.AwsEc2Instance.Spec.DisableApiTermination),
		RootBlockDevice:       rootBlock,
		Tags:                  pulumi.ToStringMap(locals.AwsTags),
	}

	// Only set optional fields when supplied
	if locals.AwsEc2Instance.Spec.KeyName != "" {
		instanceArgs.KeyName = pulumi.String(locals.AwsEc2Instance.Spec.KeyName)
	}

	if locals.AwsEc2Instance.Spec.IamInstanceProfileArn != nil &&
		locals.AwsEc2Instance.Spec.IamInstanceProfileArn.GetValue() != "" {
		instanceArgs.IamInstanceProfile = pulumi.String(locals.AwsEc2Instance.Spec.IamInstanceProfileArn.GetValue())
	}

	if locals.AwsEc2Instance.Spec.UserData != "" {
		instanceArgs.UserData = pulumi.String(locals.AwsEc2Instance.Spec.UserData)
	}

	// --------‑‑‑ Create the instance -----------------------------------
	createdInstance, err := ec2.NewInstance(ctx,
		locals.AwsEc2Instance.Metadata.Name,
		instanceArgs,
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "create aws ec2 instance")
	}

	// --------‑‑‑ Export stack outputs ----------------------------------
	ctx.Export(OpInstanceId, createdInstance.ID())
	ctx.Export(OpPrivateIp, createdInstance.PrivateIp)
	ctx.Export(OpPrivateDnsName, createdInstance.PrivateDns)
	ctx.Export(OpAvailabilityZone, createdInstance.AvailabilityZone)
	ctx.Export(OpInstanceProfileArn, createdInstance.IamInstanceProfile)

	return nil
}
