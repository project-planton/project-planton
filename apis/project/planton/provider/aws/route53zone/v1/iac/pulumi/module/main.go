package module

import (
	"fmt"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws"
	"strings"

	"github.com/pkg/errors"
	route53zonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/route53zone/v1"
	"github.com/project-planton/project-planton/apis/project/planton/provider/aws/route53zone/v1/iac/pulumi/module/outputs"
	"github.com/pulumi/pulumi-aws-native/sdk/go/aws/route53"
	awsclassicroute53 "github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *route53zonev1.Route53ZoneStackInput) error {
	//create a variable with descriptive name for the api-resource in the input
	route53Zone := stackInput.Target

	awsCredential := stackInput.AwsCredential

	//create aws provider using the credentials from the input
	awsNativeProvider, err := aws.NewProvider(ctx,
		"native-provider",
		&aws.ProviderArgs{
			AccessKey: pulumi.String(awsCredential.AccessKeyId),
			SecretKey: pulumi.String(awsCredential.SecretAccessKey),
			Region:    pulumi.String(awsCredential.Region),
		})
	if err != nil {
		return errors.Wrap(err, "failed to create aws provider")
	}

	//replace dots with hyphens to create valid managed-zone name
	managedZoneName := strings.ReplaceAll(route53Zone.Metadata.Name, ".", "-")

	//create new hosted-zone
	createdHostedZone, err := route53.NewHostedZone(ctx,
		managedZoneName,
		&route53.HostedZoneArgs{
			Name: pulumi.String(route53Zone.Metadata.Name),
			//HostedZoneTags: convertLabelsToTags(input.Labels),
		}, pulumi.Provider(awsNativeProvider))

	if err != nil {
		return errors.Wrapf(err, "failed to create hosted-zone for %s domain",
			route53Zone.Metadata.Name)
	}

	//export important information about created hosted-zone as outputs
	ctx.Export(outputs.HostedZoneName, createdHostedZone.Name)
	// todo: add nameservers export

	//for each dns-record in the input spec, insert the record in the created hosted-zone
	for index, dnsRecord := range route53Zone.Spec.Records {
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
			}, pulumi.Provider(awsNativeProvider))
		if err != nil {
			return errors.Wrapf(err, "failed to add %s rec", dnsRecord)
		}
	}
	return nil
}
