package module

import (
	"github.com/pkg/errors"
	digitaloceancertificatev1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceancertificate/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/pulumidigitaloceanprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *digitaloceancertificatev1.DigitalOceanCertificateStackInput,
) error {
	// 1. Prepare locals (metadata, labels, credentials, etc.).
	locals := initializeLocals(ctx, stackInput)

	// 2. Create a DigitalOcean provider from the supplied credential.
	digitalOceanProvider, err := pulumidigitaloceanprovider.Get(
		ctx,
		stackInput.ProviderCredential,
	)
	if err != nil {
		return errors.Wrap(err, "failed to setup digitalocean provider")
	}

	// 3. Create the certificate.
	if _, err := certificate(ctx, locals, digitalOceanProvider); err != nil {
		return errors.Wrap(err, "failed to create certificate")
	}

	return nil
}
