package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/convertstringmaps"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func subnets(
	ctx *pulumi.Context,
	locals *Locals,
	createdVpc *ec2.Vpc,
	createdPublicRouteTable *ec2.RouteTable,
) error {

	// We'll store a NAT Gateway for each AZ here, so private subnets can reference it.
	natGatewayPerAz := make(map[string]*ec2.NatGateway)

	// 1) Create all PUBLIC subnets first.
	sortedPublicAzKeys := getSortedAzKeys(locals.PublicAzSubnetMap)
	for _, availabilityZone := range sortedPublicAzKeys {
		azSubnetMap := locals.PublicAzSubnetMap[AvailabilityZone(availabilityZone)]
		sortedSubnetNames := getSortedSubnetNameKeys(azSubnetMap)

		for i, subnetName := range sortedSubnetNames {
			// Create the public subnet in this AZ
			createdSubnet, err := ec2.NewSubnet(ctx,
				subnetName,
				&ec2.SubnetArgs{
					VpcId:            createdVpc.ID(),
					CidrBlock:        pulumi.String(azSubnetMap[SubnetName(subnetName)]),
					AvailabilityZone: pulumi.String(availabilityZone),
					// required for a typical public subnet
					MapPublicIpOnLaunch: pulumi.Bool(true),
					Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
						stringmaps.AddEntry(locals.AwsTags, "Name", subnetName)),
				}, pulumi.Parent(createdVpc))
			if err != nil {
				return errors.Wrapf(err, "error creating public subnet %s", subnetName)
			}

			ctx.Export(fmt.Sprintf("%s.%d.%s", OpPublicSubnets, i, OpSubnetName), pulumi.String(subnetName))
			ctx.Export(fmt.Sprintf("%s.%d.%s", OpPublicSubnets, i, OpSubnetId), createdSubnet.ID())
			ctx.Export(fmt.Sprintf("%s.%d.%s", OpPublicSubnets, i, OpSubnetCidr), createdSubnet.CidrBlock)

			// Associate this public subnet with the public route table
			_, err = ec2.NewRouteTableAssociation(ctx,
				fmt.Sprintf("public-route-assoc-%s", subnetName),
				&ec2.RouteTableAssociationArgs{
					RouteTableId: createdPublicRouteTable.ID(),
					SubnetId:     createdSubnet.ID(),
				}, pulumi.Parent(createdPublicRouteTable))
			if err != nil {
				return errors.Wrap(err, "error associating route table with public subnet")
			}

			// If NAT is enabled, create one NAT Gateway in the first public subnet per AZ
			if locals.AwsVpc.Spec.IsNatGatewayEnabled && i == 0 {
				createdElasticIp, err := ec2.NewEip(ctx,
					fmt.Sprintf("nat-eip-%s", subnetName),
					&ec2.EipArgs{
						Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
							stringmaps.AddEntry(locals.AwsTags, "Name",
								fmt.Sprintf("%s-nat-eip", subnetName))),
					}, pulumi.Parent(createdSubnet))
				if err != nil {
					return errors.Wrap(err, "error creating elastic ip for nat gateway")
				}

				natGw, err := ec2.NewNatGateway(ctx,
					fmt.Sprintf("nat-gateway-%s", subnetName),
					&ec2.NatGatewayArgs{
						SubnetId:     createdSubnet.ID(),
						AllocationId: createdElasticIp.ID(),
						Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
							stringmaps.AddEntry(locals.AwsTags, "Name", fmt.Sprintf("%s-nat", subnetName))),
					}, pulumi.Parent(createdSubnet))
				if err != nil {
					return errors.Wrap(err, "error creating nat gateway in public subnet")
				}

				natGatewayPerAz[availabilityZone] = natGw
			}
		}
	}

	// 2) Create all PRIVATE subnets next.
	sortedPrivateAzKeys := getSortedAzKeys(locals.PrivateAzSubnetMap)
	for _, availabilityZone := range sortedPrivateAzKeys {
		azSubnetMap := locals.PrivateAzSubnetMap[AvailabilityZone(availabilityZone)]
		sortedSubnetNames := getSortedSubnetNameKeys(azSubnetMap)

		for i, subnetName := range sortedSubnetNames {
			// Create the private subnet in this AZ
			createdSubnet, err := ec2.NewSubnet(ctx,
				subnetName,
				&ec2.SubnetArgs{
					VpcId:               createdVpc.ID(),
					CidrBlock:           pulumi.String(azSubnetMap[SubnetName(subnetName)]),
					AvailabilityZone:    pulumi.String(availabilityZone),
					MapPublicIpOnLaunch: pulumi.Bool(false),
					Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
						stringmaps.AddEntry(locals.AwsTags, "Name", subnetName)),
				}, pulumi.Parent(createdVpc))
			if err != nil {
				return errors.Wrapf(err, "error creating private subnet %s", subnetName)
			}

			ctx.Export(fmt.Sprintf("%s.%d.%s", OpPrivateSubnets, i, OpSubnetName), pulumi.String(subnetName))
			ctx.Export(fmt.Sprintf("%s.%d.%s", OpPrivateSubnets, i, OpSubnetId), createdSubnet.ID())
			ctx.Export(fmt.Sprintf("%s.%d.%s", OpPrivateSubnets, i, OpSubnetCidr), createdSubnet.CidrBlock)

			// If NAT is enabled, create a route table that routes to that NAT
			// NOTE: We only do this if a NAT gateway exists in this AZ.
			if locals.AwsVpc.Spec.IsNatGatewayEnabled {
				natGw, hasNat := natGatewayPerAz[availabilityZone]
				if hasNat {
					privateRouteTable, err := ec2.NewRouteTable(ctx,
						fmt.Sprintf("private-route-table-%s", subnetName),
						&ec2.RouteTableArgs{
							VpcId: createdVpc.ID(),
							Routes: ec2.RouteTableRouteArray{
								&ec2.RouteTableRouteArgs{
									CidrBlock:    pulumi.String("0.0.0.0/0"),
									NatGatewayId: natGw.ID(),
								},
							},
							Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
								stringmaps.AddEntry(locals.AwsTags, "Name", fmt.Sprintf("%s-private-rt", subnetName))),
						}, pulumi.Parent(createdSubnet))
					if err != nil {
						return errors.Wrap(err, "error creating private route table for NAT")
					}

					_, err = ec2.NewRouteTableAssociation(ctx,
						fmt.Sprintf("private-route-assoc-%s", subnetName),
						&ec2.RouteTableAssociationArgs{
							RouteTableId: privateRouteTable.ID(),
							SubnetId:     createdSubnet.ID(),
						}, pulumi.Parent(privateRouteTable))
					if err != nil {
						return errors.Wrap(err, "error associating private route table")
					}
				}
			}
		}
	}

	return nil
}
