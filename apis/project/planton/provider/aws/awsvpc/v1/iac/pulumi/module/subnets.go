package module

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/convertstringmaps"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func subnets(ctx *pulumi.Context, locals *Locals, createdVpc *ec2.Vpc,
	createdPublicRouteTable *ec2.RouteTable) error {
	// iterate through azs and create the configured number of public and private subnets per az
	sortedPrivateAzKeys := getSortedAzKeys(locals.PrivateAzSubnetMap)
	// create private subnets
	for _, availabilityZone := range sortedPrivateAzKeys {
		azSubnetMap := locals.PrivateAzSubnetMap[AvailabilityZone(availabilityZone)]
		sortedSubnetNames := getSortedSubnetNameKeys(azSubnetMap)
		for i, subnetName := range sortedSubnetNames {
			// create private subnet in az
			createdSubnet, err := ec2.NewSubnet(ctx,
				subnetName,
				&ec2.SubnetArgs{
					VpcId:            createdVpc.ID(),
					CidrBlock:        pulumi.String(azSubnetMap[SubnetName(subnetName)]),
					AvailabilityZone: pulumi.String(availabilityZone),
					Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
						stringmaps.AddEntry(locals.AwsTags, "Name", subnetName)),
				}, pulumi.Parent(createdVpc))
			if err != nil {
				return errors.Wrapf(err, "error creating private subnet %s", subnetName)
			}
			ctx.Export(fmt.Sprintf("%s.%d", outputs.PrivateSubnetsName, i), pulumi.String(subnetName))
			ctx.Export(fmt.Sprintf("%s.%d", outputs.PrivateSubnetsId, i), createdSubnet.ID())
			ctx.Export(fmt.Sprintf("%s.%d", outputs.PrivateSubnetsCidr, i), createdSubnet.CidrBlock)

			if locals.AwsVpc.Spec.IsNatGatewayEnabled {
				createdNatGateway, err := natGateway(ctx, locals, createdVpc, subnetName, createdSubnet)
				if err != nil {
					return errors.Wrapf(err, "failed to create nat-gateway for %s subnet", subnetName)
				}
				ctx.Export(fmt.Sprintf("%s.%d", outputs.PrivateSubnetsNatGatewayId, i), createdNatGateway.ID())
				ctx.Export(fmt.Sprintf("%s.%d", outputs.PrivateSubnetsNatGatewayPublicIp, i), createdNatGateway.PublicIp)
				ctx.Export(fmt.Sprintf("%s.%d", outputs.PrivateSubnetsNatGatewayPrivateIp, i), createdNatGateway.PrivateIp)
			}
		}
	}

	sortedPublicAzKeys := getSortedAzKeys(locals.PublicAzSubnetMap)
	// create public subnets
	for _, availabilityZone := range sortedPublicAzKeys {
		azSubnetMap := locals.PublicAzSubnetMap[AvailabilityZone(availabilityZone)]
		sortedSubnetNames := getSortedSubnetNameKeys(azSubnetMap)
		for i, subnetName := range sortedSubnetNames {
			// create public subnet in az
			createdSubnet, err := ec2.NewSubnet(ctx,
				subnetName,
				&ec2.SubnetArgs{
					VpcId:            createdVpc.ID(),
					CidrBlock:        pulumi.String(azSubnetMap[SubnetName(subnetName)]),
					AvailabilityZone: pulumi.String(availabilityZone),
					//required for public subnets
					MapPublicIpOnLaunch: pulumi.Bool(true),
					Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
						stringmaps.AddEntry(locals.AwsTags, "Name", subnetName)),
				}, pulumi.Parent(createdVpc))
			if err != nil {
				return errors.Wrapf(err, "error creating public subnet %s", subnetName)
			}

			ctx.Export(fmt.Sprintf("%s.%d", outputs.PublicSubnetsName, i), pulumi.String(subnetName))
			ctx.Export(fmt.Sprintf("%s.%d", outputs.PublicSubnetsId, i), createdSubnet.ID())
			ctx.Export(fmt.Sprintf("%s.%d", outputs.PublicSubnetsCidr, i), createdSubnet.CidrBlock)

			_, err = ec2.NewRouteTableAssociation(ctx,
				fmt.Sprintf("public-route-assoc-%s", subnetName),
				&ec2.RouteTableAssociationArgs{
					RouteTableId: createdPublicRouteTable.ID(),
					SubnetId:     createdSubnet.ID(),
				}, pulumi.Parent(createdPublicRouteTable))
			if err != nil {
				return errors.Wrap(err, "error associating route table with public subnet")
			}
		}
	}
	return nil
}
