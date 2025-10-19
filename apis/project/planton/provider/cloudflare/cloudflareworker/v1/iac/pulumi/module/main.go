package module

import (
	"github.com/pkg/errors"
	cloudflareworkerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflareworker/v1"
	"github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the entry‑point expected by the Project Planton CLI.
func Resources(
	ctx *pulumi.Context,
	stackInput *cloudflareworkerv1.CloudflareWorkerStackInput,
) error {
	// 1. Prepare handy references.
	locals := initializeLocals(ctx, stackInput)

	// 2. Stand‑up a Cloudflare provider from the supplied credential.
	createdProvider, err := cloudflare.NewProvider(ctx, "cloudflare-provider",
		&cloudflare.ProviderArgs{
			ApiToken: pulumi.String(locals.CloudflareProviderConfig.ApiToken),
		})
	if err != nil {
		return errors.Wrap(err, "failed to set up cloudflare provider")
	}

	// 3. Create (or update) the Worker script.
	createdWorkerScript, err := worker_script(ctx, locals, createdProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create cloudflare worker script")
	}

	// 4. Optionally attach the script to a route.
	if _, err := route(ctx, locals, createdProvider, createdWorkerScript); err != nil {
		return errors.Wrap(err, "failed to create cloudflare worker route")
	}

	return nil
}
