package module

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"

	"github.com/pkg/errors"
	awseksclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsekscluster/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awseksclusterv1.AwsEksClusterStackInput) (err error) {
	// Create a variable with descriptive name for the API resource in the input
	awsEksCluster := stackInput.Target

	awsCredential := stackInput.ProviderCredential

	var provider *aws.Provider

	//create aws provider using the credentials from the input
	if awsCredential == nil {
		//create aws provider using the credentials from the input
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create aws native provider")
		}
	} else {
		//create aws provider using the credentials from the input
		provider, err = aws.NewProvider(ctx,
			"classic-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create aws native provider")
		}
	}

	// Default tags
	defaultTags := pulumi.StringMap{
		"module_name": pulumi.String("planton/aws-eks-cluster-pulumi-module"),
	}

	var userTags pulumi.StringMap
	if awsEksCluster.Spec.Tags != nil {
		userTags = pulumi.StringMap{}
		for key, value := range awsEksCluster.Spec.Tags {
			userTags[key] = pulumi.String(value)
		}
	} else {
		userTags = pulumi.StringMap{}
	}

	// Tags to be applied to both roles
	tags := pulumi.StringMap{}

	for key, value := range defaultTags {
		tags[key] = value
	}

	for key, value := range userTags {
		tags[key] = value
	}

	fmt.Printf("awsEksCluster.Spec.Tags: %+v\n", awsEksCluster.Spec.Tags)

	// Prepare SubnetIds and SecurityGroupIds as pulumi.StringArray
	subnetIds := pulumi.ToStringArray(awsEksCluster.Spec.Subnets)
	securityGroupIds := pulumi.ToStringArray(awsEksCluster.Spec.SecurityGroups)

	var awsEksClusterRoleArn pulumi.StringInput
	var awsEksClusterNodeRoleArn pulumi.StringInput

	// **Declare the IAM roles outside the if blocks for scope**
	var awsEksClusterRole *iam.Role   // Ensure this is declared here
	var eksNodeGroupRole *iam.Role // Ensure this is declared here

	// Create EKS cluster role and attach policies
	if awsEksCluster.Spec.RoleArn != "" {
		awsEksClusterRoleArn = pulumi.String(awsEksCluster.Spec.RoleArn)
	} else {
		awsEksClusterRole, err := iam.NewRole(ctx, "awsEksClusterRole", &iam.RoleArgs{
			Name: pulumi.String(fmt.Sprintf("%s-cluster-role", awsEksCluster.Metadata.Name)),
			AssumeRolePolicy: pulumi.String(`{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "eks.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }`),
			Tags: tags,
		})
		if err != nil {
			return err
		}

		// Attach policies to the cluster role
		_, err = iam.NewRolePolicyAttachment(ctx, "awsEksClusterRole-AmazonEKSClusterPolicy", &iam.RolePolicyAttachmentArgs{
			Role:      awsEksClusterRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
		})
		if err != nil {
			return err
		}

		_, err = iam.NewRolePolicyAttachment(ctx, "clusterRole-AmazonEKSServicePolicy", &iam.RolePolicyAttachmentArgs{
			Role:      awsEksClusterRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSServicePolicy"),
		})
		if err != nil {
			return err
		}

		// Set roleArnInput to the ARN of the created cluster role
		awsEksClusterRoleArn = awsEksClusterRole.Arn

	}

	// Create EKS node group role and attach policies
	if awsEksCluster.Spec.NodeRoleArn != "" {
		awsEksClusterNodeRoleArn = pulumi.String(awsEksCluster.Spec.NodeRoleArn)
	} else {
		eksNodeGroupRole, err := iam.NewRole(ctx, "eksNodeGroupRole", &iam.RoleArgs{
			Name: pulumi.Sprintf("%s-node-group-role", awsEksCluster.Metadata.Name),
			AssumeRolePolicy: pulumi.String(`{
    "Version": "2012-10-17",
    "Statement": [
      {
        "Effect": "Allow",
        "Principal": {
          "Service": "ec2.amazonaws.com"
        },
        "Action": "sts:AssumeRole"
      }
    ]
  }`),
			Tags: tags,
		})
		if err != nil {
			return err
		}

		// Attach policies to the node group role
		// Attach policies to the node role
		_, err = iam.NewRolePolicyAttachment(ctx, "eksNodeGroupRole-AmazonEC2ContainerRegistryReadOnly", &iam.RolePolicyAttachmentArgs{
			Role:      eksNodeGroupRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"),
		})
		if err != nil {
			return err
		}

		_, err = iam.NewRolePolicyAttachment(ctx, "nodeRole-AmazonEKS_CNI_Policy", &iam.RolePolicyAttachmentArgs{
			Role:      eksNodeGroupRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"),
		})
		if err != nil {
			return err
		}

		_, err = iam.NewRolePolicyAttachment(ctx, "nodeRole-AmazonEKSWorkerNodePolicy", &iam.RolePolicyAttachmentArgs{
			Role:      eksNodeGroupRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"),
		})
		if err != nil {
			return err
		}

		awsEksClusterNodeRoleArn = eksNodeGroupRole.Arn
	}

	// Dependencies for the EKS cluster
	var clusterDependsOn []pulumi.Resource
	if awsEksClusterRole != nil {
		clusterDependsOn = append(clusterDependsOn, awsEksClusterRole)
	}

	// Create EKS cluster
	awsEksClusterResource, err := eks.NewCluster(ctx,
		"awsEksCluster",
		&eks.ClusterArgs{
			Name:    pulumi.String(awsEksCluster.Metadata.Name),
			RoleArn: awsEksClusterRoleArn,
			VpcConfig: &eks.ClusterVpcConfigArgs{
				SubnetIds:        subnetIds,
				SecurityGroupIds: securityGroupIds,
			},
			Tags: tags,
		},
		pulumi.DependsOn(clusterDependsOn),
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create EKS cluster")
	}

	// Prepare dependencies for the managed node group
	var nodeGroupDependsOn []pulumi.Resource
	nodeGroupDependsOn = append(nodeGroupDependsOn, awsEksClusterResource)
	if eksNodeGroupRole != nil {
		nodeGroupDependsOn = append(nodeGroupDependsOn, eksNodeGroupRole)
	}

	// Create managed node group
	managedNodeGroup, err := eks.NewNodeGroup(ctx,
		"eksManagedNodeGroup",
		&eks.NodeGroupArgs{
			ClusterName:   awsEksClusterResource.Name,
			NodeGroupName: pulumi.String(fmt.Sprintf("%s-node-group", awsEksCluster.Metadata.Name)),
			NodeRoleArn:   awsEksClusterNodeRoleArn,
			SubnetIds:     pulumi.ToStringArray(awsEksCluster.Spec.Subnets),
			InstanceTypes: pulumi.StringArray{
				pulumi.String(awsEksCluster.Spec.InstanceType),
			},
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: pulumi.Int(awsEksCluster.Spec.DesiredSize),
				MaxSize:     pulumi.Int(awsEksCluster.Spec.MaxSize),
				MinSize:     pulumi.Int(awsEksCluster.Spec.MinSize),
			},
		},
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create EKS node group")
	}

	// Export outputs
	ctx.Export("awsEksClusterName", awsEksClusterResource.Name)
	ctx.Export("eksNodeGroupName", managedNodeGroup.NodeGroupName)

	return nil
}
