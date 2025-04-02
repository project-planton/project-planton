package module

import (
	"fmt"
	"github.com/pkg/errors"
	awsvpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1/iac/pulumi/module/outputs"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/datatypes/stringmaps/convertstringmaps"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsvpcv1.AwsVpcStackInput) error {
	locals, err := initializeLocals(ctx, stackInput)
	if err != nil {
		return errors.Wrap(err, "failed to initialize locals")
	}

	awsCredential := stackInput.ProviderCredential

	//create aws provider using the credentials from the input
	awsProvider, err := aws.NewProvider(ctx,
		"classic-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	// create vpc
	createdVpc, err := ec2.NewVpc(ctx,
		locals.AwsVpc.Metadata.Name,
		&ec2.VpcArgs{
			CidrBlock:          pulumi.String(locals.AwsVpc.Spec.VpcCidr),
			EnableDnsSupport:   pulumi.Bool(locals.AwsVpc.Spec.IsDnsSupportEnabled),
			EnableDnsHostnames: pulumi.Bool(locals.AwsVpc.Spec.IsDnsHostnamesEnabled),
			Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
				stringmaps.AddEntry(locals.AwsTags, "Name", locals.AwsVpc.Metadata.Name)),
		}, pulumi.Provider(awsProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create vpc")
	}

	//add vpc id to outputs
	ctx.Export(outputs.VpcId, createdVpc.ID())

	// internet gateway for public subnets
	createdInternetGateway, err := ec2.NewInternetGateway(ctx,
		locals.AwsVpc.Metadata.Name,
		&ec2.InternetGatewayArgs{
			VpcId: createdVpc.ID(),
			Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
				stringmaps.AddEntry(locals.AwsTags, "Name", locals.AwsVpc.Metadata.Name)),
		}, pulumi.Parent(createdVpc))
	if err != nil {
		return errors.Wrap(err, "failed to create internet-gateway")
	}

	//add internet-gateway id to outputs
	ctx.Export(outputs.InternetGatewayId, createdInternetGateway.ID())

	// public route table for internet access
	createdPublicRouteTable, err := ec2.NewRouteTable(ctx,
		"public-route-table",
		&ec2.RouteTableArgs{
			VpcId: createdVpc.ID(),
			Tags: convertstringmaps.ConvertGoStringMapToPulumiStringMap(
				stringmaps.AddEntry(locals.AwsTags, "Name",
					fmt.Sprintf("%s-public", locals.AwsVpc.Metadata.Name))),
			Routes: ec2.RouteTableRouteArray{
				&ec2.RouteTableRouteArgs{
					CidrBlock: pulumi.String(vars.AllowAllCidrBlock),
					GatewayId: createdInternetGateway.ID(),
				},
			},
		}, pulumi.Parent(createdInternetGateway))
	if err != nil {
		return errors.Wrap(err, "failed to created route-table for public internet access")
	}

	if err := subnets(ctx, locals, createdVpc, createdPublicRouteTable); err != nil {
		return errors.Wrap(err, "failed to create subnets")
	}

	return nil
}
