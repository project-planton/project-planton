package module

import (
	"github.com/pkg/errors"
	cloudfl "github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createWorkerScript creates the Worker script resource with content from inline or R2 URL.
// Uses Pulumi's native RemoteAsset for R2 downloads (IaC-native, works in Terraform too).
func createWorkerScript(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudfl.Provider,
) (*cloudfl.WorkerScript, error) {

	// Resolve script content (inline string or R2 URL via RemoteAsset)
	scriptContent := resolveScriptContent(ctx, locals)

	// Build plain-text environment variable bindings from env.variables
	// Note: env.secrets are uploaded separately via Cloudflare Secrets API
	var plainTextBindings cloudfl.WorkerScriptPlainTextBindingArray
	if locals.CloudflareWorker.Spec.Env != nil {
		for k, v := range locals.CloudflareWorker.Spec.Env.Variables {
			plainTextBindings = append(plainTextBindings, cloudfl.WorkerScriptPlainTextBindingArgs{
				Name: pulumi.String(k),
				Text: pulumi.String(v),
			})
		}
	}

	// Build KV namespace bindings (if any)
	var kvBindings cloudfl.WorkerScriptKvNamespaceBindingArray
	for _, kvBinding := range locals.CloudflareWorker.Spec.KvBindings {
		kvBindings = append(kvBindings, cloudfl.WorkerScriptKvNamespaceBindingArgs{
			Name:        pulumi.String(kvBinding.Name),
			NamespaceId: pulumi.String(kvBinding.GetFieldPath()),
		})
	}

	// Build Worker script arguments
	scriptArgs := &cloudfl.WorkerScriptArgs{
		AccountId:             pulumi.String(locals.CloudflareWorker.Spec.AccountId),
		Name:                  pulumi.String(locals.CloudflareWorker.Spec.ScriptName),
		Content:               scriptContent,
		PlainTextBindings:     plainTextBindings,
		KvNamespaceBindings:   kvBindings,
	}

	if locals.CloudflareWorker.Spec.CompatibilityDate != "" {
		scriptArgs.CompatibilityDate = pulumi.StringPtr(locals.CloudflareWorker.Spec.CompatibilityDate)
	}

	// Create the Worker script
	createdWorkerScript, err := cloudfl.NewWorkerScript(
		ctx,
		"worker-script",
		scriptArgs,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare worker script")
	}

	// Export stack output
	ctx.Export(OpScriptId, createdWorkerScript.ID())

	return createdWorkerScript, nil
}

