package module

import (
	"github.com/pkg/errors"
	cloudfl "github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// route attaches the Worker script to a URL pattern (if provided) and exports it.
func route(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudfl.Provider,
	_ *cloudfl.WorkerScript, // underscores silence “unused” while enforcing call‑order.
) ([]pulumi.StringOutput, error) {

	if locals.CloudflareWorker.Spec.RoutePattern == "" {
		// No route requested.
		return nil, nil
	}

	routeArgs := &cloudfl.WorkerRouteArgs{
		Pattern:    pulumi.String(locals.CloudflareWorker.Spec.RoutePattern),
		ScriptName: pulumi.String(locals.CloudflareWorker.Spec.ScriptName),
	}

	createdWorkerRoute, err := cloudfl.NewWorkerRoute(
		ctx,
		"worker-route",
		routeArgs,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare worker route")
	}

	ctx.Export(OpRouteUrls, pulumi.StringArray{createdWorkerRoute.Pattern})
	return []pulumi.StringOutput{createdWorkerRoute.Pattern}, nil
}
