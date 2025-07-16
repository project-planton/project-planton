package module

import (
	"github.com/pkg/errors"
	digitaloceandnszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceandnszone/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/digitalocean/pulumidigitaloceanprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the entry point, mirroring digital_ocean_vpc.
func Resources(
	ctx *pulumi.Context,
	stackInput *digitaloceandnszonev1.DigitalOceanDnsZoneStackInput,
) error {
	// 1. Collate locals.
	locals := initializeLocals(ctx, stackInput)

	// 2. Create DO provider from credential.
	doProvider, err := pulumidigitaloceanprovider.Get(
		ctx,
		stackInput.ProviderCredential,
	)
	if err != nil {
		return errors.Wrap(err, "failed to setup digitalocean provider")
	}

	// 3. Provision the DNS zone (domain + records).
	if _, err := dnsZone(ctx, locals, doProvider); err != nil {
		return errors.Wrap(err, "failed to create dns zone")
	}

	return nil
}
