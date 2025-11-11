package module

import (
	"github.com/pkg/errors"
	digitaloceancertificatev1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/digitalocean/digitaloceancertificate/v1"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// certificate provisions the DigitalOcean SSL certificate and exports its outputs.
//
// NOTE: The DigitalOcean Pulumi provider currently lacks fields for tags
// and automatic‑renew configuration, so spec.tags and disable_auto_renew are ignored.
func certificate(
	ctx *pulumi.Context,
	locals *Locals,
	digitalOceanProvider *digitalocean.Provider,
) (*digitalocean.Certificate, error) {
	var domains pulumi.StringArray

	if locals.DigitalOceanCertificate.Spec.GetLetsEncrypt() != nil {
		for _, d := range locals.DigitalOceanCertificate.Spec.GetLetsEncrypt().Domains {
			domains = append(domains, pulumi.String(d))
		}
	}

	certArgs := &digitalocean.CertificateArgs{
		Name: pulumi.String(locals.DigitalOceanCertificate.Spec.CertificateName),
		Type: pulumi.String(locals.DigitalOceanCertificate.Spec.Type.String()),
	}
	if locals.DigitalOceanCertificate.Spec.Type == digitaloceancertificatev1.DigitalOceanCertificateType_letsEncrypt {
		certArgs.Domains = domains
	}

	if locals.DigitalOceanCertificate.Spec.Type == digitaloceancertificatev1.DigitalOceanCertificateType_custom {
		certArgs.LeafCertificate = pulumi.String(locals.DigitalOceanCertificate.Spec.GetCustom().LeafCertificate)
		certArgs.PrivateKey = pulumi.String(locals.DigitalOceanCertificate.Spec.GetCustom().PrivateKey)
		if locals.DigitalOceanCertificate.Spec.GetCustom().CertificateChain != "" {
			certArgs.CertificateChain = pulumi.StringPtr(locals.DigitalOceanCertificate.Spec.GetCustom().CertificateChain)
		}
	}

	// 4. Create the certificate.
	createdCertificate, err := digitalocean.NewCertificate(
		ctx,
		"certificate",
		certArgs,
		pulumi.Provider(digitalOceanProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean certificate")
	}

	// 5. Export required stack outputs.
	ctx.Export(OpCertificateId, createdCertificate.ID())
	ctx.Export(OpExpiryRfc3339, createdCertificate.NotAfter)

	return createdCertificate, nil
}
