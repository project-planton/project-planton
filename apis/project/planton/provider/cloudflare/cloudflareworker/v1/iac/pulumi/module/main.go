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

	// 3. Create the Worker script with content from inline or R2 URL.
	createdWorkerScript, err := createWorkerScript(ctx, locals, createdProvider)
	if err != nil {
		return errors.Wrap(err, "failed to create worker script")
	}

	// 4. Upload secrets via Cloudflare Secrets API (if any).
	// This must happen AFTER the worker script exists.
	if locals.CloudflareWorker.Spec.Env != nil && len(locals.CloudflareWorker.Spec.Env.Secrets) > 0 {
		// Use Apply to ensure secrets upload happens after worker is created
		createdWorkerScript.ID().ApplyT(func(_ pulumi.ID) error {
			return uploadWorkerSecrets(
				ctx,
				locals,
				locals.CloudflareWorker.Spec.ScriptName,
				locals.CloudflareWorker.Spec.Env.Secrets,
			)
		})
	}

	// 5. Optionally attach to a route.
	if _, err := route(ctx, locals, createdProvider, createdWorkerScript); err != nil {
		return errors.Wrap(err, "failed to create cloudflare worker route")
	}

	return nil
}
