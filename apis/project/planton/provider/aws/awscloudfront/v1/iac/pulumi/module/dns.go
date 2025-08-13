package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// dns creates Route53 alias records to the CloudFront distribution when enabled.
func dns(
	ctx *pulumi.Context,
	locals *Locals,
	provider *aws.Provider,
	dist *DistributionRefs,
) error {
	zoneId := locals.Spec.GetDns().GetRoute53ZoneId()
	if zoneId == "" {
		return errors.New("dns.enabled is true but route53_zone_id is not provided")
	}
	if len(locals.Spec.Aliases) == 0 {
		return errors.New("dns.enabled is true but no aliases provided")
	}

	for i, hostname := range locals.Spec.Aliases {
		recordName := fmt.Sprintf("%s-alias-%d", locals.AwsCloudFront.Metadata.Name, i)
		_, err := route53.NewRecord(ctx, recordName, &route53.RecordArgs{
			Name:   pulumi.String(hostname),
			Type:   pulumi.String("A"),
			ZoneId: pulumi.String(zoneId),
			Aliases: route53.RecordAliasArray{
				route53.RecordAliasArgs{
					EvaluateTargetHealth: pulumi.Bool(true),
					Name:                 dist.DomainName,
					ZoneId:               dist.HostedZoneId,
				},
			},
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrapf(err, "failed to create Route53 record for %s", hostname)
		}
	}
	return nil
}
