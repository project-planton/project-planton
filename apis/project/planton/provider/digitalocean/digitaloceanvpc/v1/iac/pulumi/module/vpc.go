package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-digitalocean/sdk/v4/go/digitalocean"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// vpc provisions the VPC itself and exports its ID.
func vpc(
	ctx *pulumi.Context,
	locals *Locals,
	doProvider *digitalocean.Provider,
) (*digitalocean.Vpc, error) {

	// 1. Build the resource arguments straight from the proto fields.
	vpcArgs := &digitalocean.VpcArgs{
		Description: pulumi.String(locals.DigitalOceanVpc.Spec.Description),
		IpRange:     pulumi.String(locals.DigitalOceanVpc.Spec.IpRangeCidr),
		Name:        pulumi.String(locals.DigitalOceanVpc.Metadata.Name),
		Region:      pulumi.String(locals.DigitalOceanVpc.Spec.Region.String()),
	}

	// 2. Create the VPC.
	createdVpc, err := digitalocean.NewVpc(
		ctx,
		"vpc",
		vpcArgs,
		pulumi.Provider(doProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create digitalocean vpc")
	}

	// 3. Export stack output.
	ctx.Export(OpVpcId, createdVpc.ID())

	return createdVpc, nil
}
