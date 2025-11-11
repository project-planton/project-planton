package module

import (
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

	// 1. Build resource arguments straight from proto fields.
	networkArgs := &civo.NetworkArgs{
		Label:  pulumi.String(locals.CivoVpc.Spec.NetworkName),
		Region: pulumi.String(locals.CivoVpc.Spec.Region),
	}

	if locals.CivoVpc.Spec.IpRangeCidr != "" {
		networkArgs.CidrV4 = pulumi.String(locals.CivoVpc.Spec.IpRangeCidr)
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

	return createdNetwork, nil
}
