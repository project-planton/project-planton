package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/civo/pulumicivoprovider"
	civoipaddressv1 "github.com/project-planton/project-planton/pkg/provider/civo/civoipaddress/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the single entry‑point exposed to the Pulumi engine.
func Resources(
	ctx *pulumi.Context,
	stackInput *civoipaddressv1.CivoIpAddressStackInput,
) error {
	// 1. Setup locals.
	locals := initializeLocals(ctx, stackInput)

	// 2. Instantiate provider using the credential spec.
	civoProvider, err := pulumicivoprovider.Get(
		ctx,
		stackInput.ProviderConfig,
	)
	if err != nil {
		return errors.Wrap(err, "failed to set up Civo provider")
	}

	// 3. Provision the reserved IP address.
	if _, err := ipAddress(ctx, locals, civoProvider); err != nil {
		return errors.Wrap(err, "failed to create Civo reserved IP")
	}

	return nil
}
