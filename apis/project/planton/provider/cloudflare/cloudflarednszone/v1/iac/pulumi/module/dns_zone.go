package module

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// dnsZone provisions the Cloudflare zone and exports outputs.
func dnsZone(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudflare.Provider,
) (*cloudflare.Zone, error) {

	// 1. Convert proto enum -> provider plan string.
	var planValue string
	switch locals.CloudflareDnsZone.Spec.Plan {
	case 0: // FREE
		planValue = "free"
	case 1: // PRO
		planValue = "pro"
	case 2: // BUSINESS
		planValue = "business"
	case 3: // ENTERPRISE
		planValue = "enterprise"
	default:
		planValue = "free"
	}

	// 2. Build the arguments straight from proto fields.
	zoneArgs := &cloudflare.ZoneArgs{
		AccountId: pulumi.String(locals.CloudflareDnsZone.Spec.AccountId),
		Zone:      pulumi.String(locals.CloudflareDnsZone.Spec.ZoneName),
		Plan:      pulumi.StringPtr(planValue),
		Paused:    pulumi.BoolPtr(locals.CloudflareDnsZone.Spec.Paused),
		// NOTE: default_proxied isn't available at zone‑level in the provider,
		// so it is intentionally ignored here.
	}

	// 3. Create the zone.
	createdZone, err := cloudflare.NewZone(
		ctx,
		// Use metadata.name as the resource label to mimic Terraform naming.
		strings.ToLower(locals.CloudflareDnsZone.Metadata.Name),
		zoneArgs,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare zone")
	}

	// 4. Export required outputs.
	ctx.Export(OpZoneId, createdZone.ID())
	ctx.Export(OpNameservers, createdZone.NameServers)

	return createdZone, nil
}
