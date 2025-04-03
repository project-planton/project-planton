package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/acm"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/route53"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// certManagerCert provisions an ACM certificate with DNS validation
// and creates corresponding CNAME records in Route53.
func certManagerCert(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	spec := locals.AwsCertManagerCert.Spec
	meta := locals.AwsCertManagerCert.Metadata

	// Combine primary + alternate domains into a single list.
	allDomains := append([]string{spec.PrimaryDomainName}, spec.AlternateDomainNames...)

	// Create the ACM certificate resource.
	cert, err := acm.NewCertificate(ctx, meta.Name+"-cert", &acm.CertificateArgs{
		DomainName:              pulumi.String(spec.PrimaryDomainName),
		SubjectAlternativeNames: pulumi.ToStringArray(spec.AlternateDomainNames),
		ValidationMethod:        pulumi.String("DNS"),
		Tags:                    pulumi.ToStringMap(locals.AwsTags),
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to create ACM certificate")
	}

	// For each DomainValidationOption, create the corresponding DNS record
	// and store its fqdn as a pulumi.StringInput. We do this in an ApplyT call
	// so the dynamic values are resolved at deploy time.
	recordFqdns := cert.DomainValidationOptions.ApplyT(func(domainValidationOptions []acm.CertificateDomainValidationOption) ([]pulumi.StringInput, error) {
		var fqdnOutputs []pulumi.StringInput
		for i, dvo := range domainValidationOptions {
			record, err := route53.NewRecord(ctx, fmt.Sprintf("%s-cname-%d", meta.Name, i), &route53.RecordArgs{
				Name: pulumi.String(*dvo.ResourceRecordName),
				Records: pulumi.StringArray{
					pulumi.String(*dvo.ResourceRecordValue),
				},
				Ttl:    pulumi.Int(300),
				Type:   pulumi.String(*dvo.ResourceRecordType),
				ZoneId: pulumi.String(spec.Route53HostedZoneId),
			}, pulumi.Provider(provider))
			if err != nil {
				return nil, errors.Wrapf(err, "failed to create DNS record for domain %s", allDomains[i])
			}
			// Store the FQDN output rather than a raw string
			fqdnOutputs = append(fqdnOutputs, record.Fqdn)
		}
		return fqdnOutputs, nil
	}).(pulumi.StringArrayOutput)

	// Validate the certificate by passing the list of DNS record FQDNs as pulumi.StringArrayInput.
	_, err = acm.NewCertificateValidation(ctx, meta.Name+"-validation", &acm.CertificateValidationArgs{
		CertificateArn:        cert.Arn,
		ValidationRecordFqdns: recordFqdns,
	}, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "failed to validate ACM certificate")
	}

	// Export the issued certificate ARN as a stack output.
	ctx.Export(OpCertArn, cert.Arn)
	return nil
}
