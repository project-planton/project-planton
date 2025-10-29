package module

import (
	"github.com/pkg/errors"
	cloudfl "github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// route attaches the Worker script to a URL pattern (if provided) and exports it.
func route(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudfl.Provider,
	_ *cloudfl.WorkersScript, // underscores silence "unused" while enforcing call‑order.
) ([]pulumi.StringOutput, error) {

	if locals.CloudflareWorker.Spec.RoutePattern == "" {
		// No route requested.
		return nil, nil
	}

	// Validate that zone_id is provided when route_pattern is specified
	if locals.CloudflareWorker.Spec.ZoneId == "" {
		return nil, errors.New("zone_id is required when route_pattern is specified")
	}

	routeArgs := &cloudfl.WorkersRouteArgs{
		ZoneId:     pulumi.String(locals.CloudflareWorker.Spec.ZoneId),
		Pattern:    pulumi.String(locals.CloudflareWorker.Spec.RoutePattern),
		ScriptName: pulumi.String(locals.CloudflareWorker.Spec.Script.Name),
	}

	createdWorkerRoute, err := cloudfl.NewWorkersRoute(
		ctx,
		"workers-route",
		routeArgs,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare workers route")
	}

	ctx.Export(OpRouteUrls, pulumi.StringArray{createdWorkerRoute.Pattern})
	return []pulumi.StringOutput{createdWorkerRoute.Pattern}, nil
}
