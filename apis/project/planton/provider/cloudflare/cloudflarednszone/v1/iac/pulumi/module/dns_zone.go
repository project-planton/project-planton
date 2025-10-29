package module

import (
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// dnsZone provisions the Cloudflare zone and exports outputs.
func dnsZone(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudflare.Provider,
) (*cloudflare.Zone, error) {

	// 1. Build the arguments straight from proto fields.
	// Note: Plan field was removed in v6 and is now set via a separate API call or zone settings
	zoneArgs := &cloudflare.ZoneArgs{
		Account: cloudflare.ZoneAccountArgs{
			Id: pulumi.String(locals.CloudflareDnsZone.Spec.AccountId),
		},
		Name:   pulumi.String(locals.CloudflareDnsZone.Spec.ZoneName),
		Paused: pulumi.BoolPtr(locals.CloudflareDnsZone.Spec.Paused),
		// NOTE: default_proxied and plan are not available at zone‑level in the v6 provider.
		// Plan is now managed separately via zone settings or account configuration.
	}

	// 2. Create the zone.
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

	// 3. Export required outputs.
	ctx.Export(OpZoneId, createdZone.ID())
	ctx.Export(OpNameservers, createdZone.NameServers)

	return createdZone, nil
}
