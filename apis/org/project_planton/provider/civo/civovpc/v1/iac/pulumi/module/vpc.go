package module

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// vpc provisions the Civo network and exports its outputs.
func vpc(
	ctx *pulumi.Context,
	locals *Locals,
	civoProvider *civo.Provider,
) (*civo.Network, error) {

	// 1. Build resource arguments from proto fields.
	networkArgs := &civo.NetworkArgs{
		Label:  pulumi.String(locals.CivoVpc.Spec.NetworkName),
		Region: pulumi.String(locals.CivoVpc.Spec.Region),
	}

	// Optional: IPv4 CIDR block (if not specified, Civo auto-allocates)
	if locals.CivoVpc.Spec.IpRangeCidr != "" {
		networkArgs.CidrV4 = pulumi.String(locals.CivoVpc.Spec.IpRangeCidr)
	}

	// Optional: Set as default network for the region
	// Note: Only one network per region can be set as default
	if locals.CivoVpc.Spec.IsDefaultForRegion {
		networkArgs.Default = pulumi.Bool(true)
		ctx.Log.Info(fmt.Sprintf(
			"Network '%s' will be set as the default network for region '%s'. "+
				"Note: Only one default network is allowed per region.",
			locals.CivoVpc.Spec.NetworkName,
			locals.CivoVpc.Spec.Region,
		), nil)
	}

	// Optional: Description
	// Note: The Civo provider's Network resource doesn't currently expose a description field.
	// This field is recorded in metadata but not applied to the Civo resource.
	if locals.CivoVpc.Spec.Description != "" {
		ctx.Log.Info(fmt.Sprintf(
			"Description specified for network '%s': '%s'. "+
				"Note: The Civo Network provider doesn't currently support description field. "+
				"This is recorded in Project Planton metadata only.",
			locals.CivoVpc.Spec.NetworkName,
			locals.CivoVpc.Spec.Description,
		), nil)
	}

	// 2. Create the network.
	createdNetwork, err := civo.NewNetwork(
		ctx,
		"network",
		networkArgs,
		pulumi.Provider(civoProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create civo network")
	}

	// 3. Export stack outputs.
	ctx.Export(OpNetworkId, createdNetwork.ID())
	ctx.Export(OpCidrBlock, createdNetwork.CidrV4)
	// Note: created_at_rfc3339 is not exposed by the Civo provider as an attribute.
	// The network exists after creation, but the timestamp is not available via the provider.
	// This can be added manually if needed using time.Now() at creation time.

	return createdNetwork, nil
}
