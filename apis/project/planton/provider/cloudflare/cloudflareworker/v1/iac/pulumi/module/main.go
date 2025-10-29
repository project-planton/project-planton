package module

import (
	"fmt"

	"github.com/pkg/errors"
	cloudflareworkerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflareworker/v1"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
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

	// 3. Create AWS provider for R2 (only if R2 credentials provided).
	var r2Provider *aws.Provider
	if locals.CloudflareProviderConfig.R2 != nil {
		r2Provider, err = createR2Provider(ctx, locals)
		if err != nil {
			return errors.Wrap(err, "failed to create R2 provider")
		}
	}

	// 4. Create the Worker script with content from inline or R2 URL.
	createdWorkerScript, err := createWorkerScript(ctx, locals, createdProvider, r2Provider)
	if err != nil {
		return errors.Wrap(err, "failed to create worker script")
	}

	// 5. Upload secrets via Cloudflare Secrets API (if any).
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

	// 6. Optionally attach to a route.
	if _, err := route(ctx, locals, createdProvider, createdWorkerScript); err != nil {
		return errors.Wrap(err, "failed to create cloudflare worker route")
	}

	return nil
}

// createR2Provider creates AWS provider configured for Cloudflare R2.
// Uses credentials from CloudflareProviderConfig.r2 (already exists in proto).
func createR2Provider(
	ctx *pulumi.Context,
	locals *Locals,
) (*aws.Provider, error) {

	r2Creds := locals.CloudflareProviderConfig.R2

	// Determine R2 endpoint (use custom if provided, otherwise derive from account ID)
	var endpoint string
	if r2Creds.Endpoint != "" {
		endpoint = r2Creds.Endpoint
	} else {
		// Default R2 endpoint format
		endpoint = fmt.Sprintf("https://%s.r2.cloudflarestorage.com", locals.CloudflareWorker.Spec.AccountId)
	}

	r2Provider, err := aws.NewProvider(ctx, "r2-provider", &aws.ProviderArgs{
		Region:                    pulumi.String("auto"),
		AccessKey:                 pulumi.String(r2Creds.AccessKeyId),
		SecretKey:                 pulumi.String(r2Creds.SecretAccessKey),
		S3UsePathStyle:            pulumi.Bool(true),
		SkipCredentialsValidation: pulumi.Bool(true),
		SkipMetadataApiCheck:      pulumi.Bool(true),
		SkipRegionValidation:      pulumi.Bool(true),
		SkipRequestingAccountId:   pulumi.Bool(true),
		Endpoints: aws.ProviderEndpointArray{
			aws.ProviderEndpointArgs{
				S3: pulumi.String(endpoint),
			},
		},
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to create AWS provider for R2")
	}

	return r2Provider, nil
}
