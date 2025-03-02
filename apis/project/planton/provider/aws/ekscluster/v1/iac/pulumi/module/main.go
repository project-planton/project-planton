package module

import (
	"fmt"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"

	"github.com/pkg/errors"
	eksclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ekscluster/v1"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/eks"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *eksclusterv1.EksClusterStackInput) (err error) {
	// Create a variable with descriptive name for the API resource in the input
	eksCluster := stackInput.Target

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
		"module_name": pulumi.String("planton/eks-cluster-pulumi-module"),
	}

	var userTags pulumi.StringMap
	if eksCluster.Spec.Tags != nil {
		userTags = pulumi.StringMap{}
		for key, value := range eksCluster.Spec.Tags {
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

	fmt.Printf("eksCluster.Spec.Tags: %+v\n", eksCluster.Spec.Tags)

	// Prepare SubnetIds and SecurityGroupIds as pulumi.StringArray
	subnetIds := pulumi.ToStringArray(eksCluster.Spec.Subnets)
	securityGroupIds := pulumi.ToStringArray(eksCluster.Spec.SecurityGroups)

	var eksClusterRoleArn pulumi.StringInput
	var eksClusterNodeRoleArn pulumi.StringInput

	// **Declare the IAM roles outside the if blocks for scope**
	var eksClusterRole *iam.Role   // Ensure this is declared here
	var eksNodeGroupRole *iam.Role // Ensure this is declared here

	// Create EKS cluster role and attach policies
	if eksCluster.Spec.RoleArn != "" {
		eksClusterRoleArn = pulumi.String(eksCluster.Spec.RoleArn)
	} else {
		eksClusterRole, err := iam.NewRole(ctx, "eksClusterRole", &iam.RoleArgs{
			Name: pulumi.String(fmt.Sprintf("%s-cluster-role", eksCluster.Metadata.Name)),
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
		_, err = iam.NewRolePolicyAttachment(ctx, "eksClusterRole-AmazonEKSClusterPolicy", &iam.RolePolicyAttachmentArgs{
			Role:      eksClusterRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"),
		})
		if err != nil {
			return err
		}

		_, err = iam.NewRolePolicyAttachment(ctx, "clusterRole-AmazonEKSServicePolicy", &iam.RolePolicyAttachmentArgs{
			Role:      eksClusterRole.Name,
			PolicyArn: pulumi.String("arn:aws:iam::aws:policy/AmazonEKSServicePolicy"),
		})
		if err != nil {
			return err
		}

		// Set roleArnInput to the ARN of the created cluster role
		eksClusterRoleArn = eksClusterRole.Arn

	}

	// Create EKS node group role and attach policies
	if eksCluster.Spec.NodeRoleArn != "" {
		eksClusterNodeRoleArn = pulumi.String(eksCluster.Spec.NodeRoleArn)
	} else {
		eksNodeGroupRole, err := iam.NewRole(ctx, "eksNodeGroupRole", &iam.RoleArgs{
			Name: pulumi.Sprintf("%s-node-group-role", eksCluster.Metadata.Name),
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

		eksClusterNodeRoleArn = eksNodeGroupRole.Arn
	}

	// Dependencies for the EKS cluster
	var clusterDependsOn []pulumi.Resource
	if eksClusterRole != nil {
		clusterDependsOn = append(clusterDependsOn, eksClusterRole)
	}

	// Create EKS cluster
	eksClusterResource, err := eks.NewCluster(ctx,
		"eksCluster",
		&eks.ClusterArgs{
			Name:    pulumi.String(eksCluster.Metadata.Name),
			RoleArn: eksClusterRoleArn,
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
	nodeGroupDependsOn = append(nodeGroupDependsOn, eksClusterResource)
	if eksNodeGroupRole != nil {
		nodeGroupDependsOn = append(nodeGroupDependsOn, eksNodeGroupRole)
	}

	// Create managed node group
	managedNodeGroup, err := eks.NewNodeGroup(ctx,
		"eksManagedNodeGroup",
		&eks.NodeGroupArgs{
			ClusterName:   eksClusterResource.Name,
			NodeGroupName: pulumi.String(fmt.Sprintf("%s-node-group", eksCluster.Metadata.Name)),
			NodeRoleArn:   eksClusterNodeRoleArn,
			SubnetIds:     pulumi.ToStringArray(eksCluster.Spec.Subnets),
			InstanceTypes: pulumi.StringArray{
				pulumi.String(eksCluster.Spec.InstanceType),
			},
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: pulumi.Int(eksCluster.Spec.DesiredSize),
				MaxSize:     pulumi.Int(eksCluster.Spec.MaxSize),
				MinSize:     pulumi.Int(eksCluster.Spec.MinSize),
			},
		},
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create EKS node group")
	}

	// Export outputs
	ctx.Export("eksClusterName", eksClusterResource.Name)
	ctx.Export("eksNodeGroupName", managedNodeGroup.NodeGroupName)

	return nil
}
