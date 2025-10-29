package module

import (
	"github.com/pkg/errors"
	cloudfl "github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// route creates DNS record and attaches the Worker script to a URL pattern (if DNS is configured).
func route(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudfl.Provider,
	_ *cloudfl.WorkersScript, // underscores silence "unused" while enforcing call‑order.
) ([]pulumi.StringOutput, error) {

	// Check if DNS configuration is provided and enabled
	if locals.CloudflareWorker.Spec.Dns == nil || !locals.CloudflareWorker.Spec.Dns.Enabled {
		// No DNS/route configuration requested or explicitly disabled.
		return nil, nil
	}

	dns := locals.CloudflareWorker.Spec.Dns

	// Validate required DNS fields
	if dns.ZoneId == "" {
		return nil, errors.New("dns.zone_id is required when dns is enabled")
	}
	if dns.Hostname == "" {
		return nil, errors.New("dns.hostname is required when dns is enabled")
	}

	zoneId := pulumi.String(dns.ZoneId).ToStringOutput()

	// Create DNS record for the hostname
	createdDnsRecord, err := createDnsRecord(ctx, locals, cloudflareProvider, zoneId)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create dns record")
	}

	// Determine the route pattern: use specified pattern or default to "hostname/*"
	routePattern := dns.RoutePattern
	if routePattern == "" {
		routePattern = dns.Hostname + "/*"
	}

	routeArgs := &cloudfl.WorkersRouteArgs{
		ZoneId:  zoneId,
		Pattern: pulumi.String(routePattern),
		Script:  pulumi.String(locals.CloudflareWorker.Spec.Script.Name),
	}

	// Create the route, ensuring it depends on the DNS record
	var routeOptions []pulumi.ResourceOption
	routeOptions = append(routeOptions, pulumi.Provider(cloudflareProvider))
	if createdDnsRecord != nil {
		routeOptions = append(routeOptions, pulumi.DependsOn([]pulumi.Resource{createdDnsRecord}))
	}

	createdWorkerRoute, err := cloudfl.NewWorkersRoute(
		ctx,
		"workers-route",
		routeArgs,
		routeOptions...,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare workers route")
	}

	ctx.Export(OpRouteUrls, pulumi.StringArray{createdWorkerRoute.Pattern})
	return []pulumi.StringOutput{createdWorkerRoute.Pattern}, nil
}
