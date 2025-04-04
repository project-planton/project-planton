package module

import (
	"fmt"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	awsclassic "github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"strings"

	"github.com/pkg/errors"
	awsroute53zonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsroute53zone/v1"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/route53"
	awsclassicroute53 "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *awsroute53zonev1.AwsRoute53ZoneStackInput) error {
	//create a variable with descriptive name for the api-resource in the input
	awsRoute53Zone := stackInput.Target

	awsCredential := stackInput.ProviderCredential
	var provider *aws.Provider
	var classicProvider *awsclassic.Provider
	var err error

	// If the user didn't provide AWS credentials, create a default provider.
	// Otherwise, inject custom credentials for the region, access key, etc.
	if awsCredential == nil {
		provider, err = aws.NewProvider(ctx,
			"native-provider",
			&aws.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS provider")
		}
		classicProvider, err = awsclassic.NewProvider(ctx,
			"classic-provider",
			&awsclassic.ProviderArgs{})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS classic provider")
		}
	} else {
		provider, err = aws.NewProvider(ctx,
			"native-provider",
			&aws.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create AWS provider with custom credentials")
		}
		classicProvider, err = awsclassic.NewProvider(ctx,
			"classic-provider",
			&awsclassic.ProviderArgs{
				AccessKey: pulumi.String(awsCredential.AccessKeyId),
				SecretKey: pulumi.String(awsCredential.SecretAccessKey),
				Region:    pulumi.String(awsCredential.Region),
			})
		if err != nil {
			return errors.Wrap(err, "failed to create default AWS classic provider")
		}
	}

	//replace dots with hyphens to create valid managed-zone name
	managedZoneName := strings.ReplaceAll(awsRoute53Zone.Metadata.Name, ".", "-")

	//create new hosted-zone
	createdHostedZone, err := route53.NewHostedZone(ctx,
		managedZoneName,
		&route53.HostedZoneArgs{
			Name: pulumi.String(awsRoute53Zone.Metadata.Name),
			//HostedZoneTags: convertLabelsToTags(input.KubernetesLabels),
		}, pulumi.Provider(provider))

	if err != nil {
		return errors.Wrapf(err, "failed to create hosted-zone for %s domain",
			awsRoute53Zone.Metadata.Name)
	}

	//export important information about created hosted-zone as outputs
	ctx.Export(OpZoneName, createdHostedZone.Name)
	ctx.Export(OpZoneId, createdHostedZone.ID())
	ctx.Export(OpNameservers, createdHostedZone.NameServers)

	//for each dns-record in the input spec, insert the record in the created hosted-zone
	for index, dnsRecord := range awsRoute53Zone.Spec.Records {
		TtlSeconds := dnsRecord.TtlSeconds
		if TtlSeconds == 0 {
			TtlSeconds = 300 // Set Default TTL to 300 Seconds
		}
		_, err := awsclassicroute53.NewRecord(ctx,
			fmt.Sprintf("dns-record-%d", index),
			&awsclassicroute53.RecordArgs{
				ZoneId:  createdHostedZone.ID(),
				Name:    pulumi.String(dnsRecord.Name),
				Ttl:     pulumi.IntPtr(int(TtlSeconds)),
				Type:    pulumi.String(dnsRecord.RecordType.String()),
				Records: pulumi.ToStringArray(dnsRecord.Values),
			}, pulumi.Provider(classicProvider))
		if err != nil {
			return errors.Wrapf(err, "failed to add %s rec", dnsRecord)
		}
	}
	return nil
}
