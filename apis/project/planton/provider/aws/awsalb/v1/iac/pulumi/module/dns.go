package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// dns configures Route 53 alias records for the ALB if dns_config.enabled is true.
func dns(
	ctx *pulumi.Context,
	locals *Locals,
	provider *aws.Provider,
	albResource *lb.LoadBalancer,
) error {
	dnsCfg := locals.AwsAlb.Spec.Dns

	if dnsCfg.Route53ZoneId == "" {
		return errors.New("dns_config.enabled is true but route53_zone_id is not provided")
	}

	if len(dnsCfg.Hostnames) == 0 {
		return errors.New("dns_config.enabled is true but no hostnames provided")
	}

	for i, hostname := range dnsCfg.Hostnames {
		recordName := fmt.Sprintf("%s-dns-%d", locals.AwsAlb.Metadata.Name, i)

		_, err := route53.NewRecord(ctx, recordName, &route53.RecordArgs{
			Name:   pulumi.String(hostname),
			Type:   pulumi.String("A"),
			ZoneId: pulumi.String(dnsCfg.Route53ZoneId),
			Aliases: route53.RecordAliasArray{
				route53.RecordAliasArgs{
					Name:                 albResource.DnsName,
					ZoneId:               albResource.ZoneId,
					EvaluateTargetHealth: pulumi.Bool(true),
				},
			},
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrapf(err, "failed to create Route53 record for %s", hostname)
		}
	}

	return nil
}
