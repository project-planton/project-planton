package module

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v7/go/aws/s3"
	cloudfl "github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createWorkerScript creates the Worker script resource with content from R2 bundle.
// Uses AWS S3 provider for IaC-native private R2 bucket access.
func createWorkerScript(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudfl.Provider,
	r2Provider *aws.Provider,
) (*cloudfl.WorkersScript, error) {

	// Fetch script content from R2 bundle
	script := locals.CloudflareWorker.Spec.Script
	bundle := script.Bundle

	// Use GetObject instead of deprecated LookupBucketObject
	scriptObject := s3.GetObjectOutput(ctx, s3.GetObjectOutputArgs{
		Bucket: pulumi.String(bundle.Bucket),
		Key:    pulumi.String(bundle.Path),
	}, pulumi.Provider(r2Provider))

	scriptContent := scriptObject.Body()

	// Build plain-text environment variable bindings from env.variables
	// Note: env.secrets are uploaded separately via Cloudflare Secrets API
	var plainTextBindings cloudfl.WorkersScriptPlainTextBindingArray
	if locals.CloudflareWorker.Spec.Env != nil {
		for k, v := range locals.CloudflareWorker.Spec.Env.Variables {
			plainTextBindings = append(plainTextBindings, cloudfl.WorkersScriptPlainTextBindingArgs{
				Name: pulumi.String(k),
				Text: pulumi.String(v),
			})
		}
	}

	// Build KV namespace bindings (if any)
	var kvBindings cloudfl.WorkersScriptKvNamespaceBindingArray
	for _, kvBinding := range locals.CloudflareWorker.Spec.KvBindings {
		kvBindings = append(kvBindings, cloudfl.WorkersScriptKvNamespaceBindingArgs{
			Name:        pulumi.String(kvBinding.Name),
			NamespaceId: pulumi.String(kvBinding.GetFieldPath()),
		})
	}

	// Build Worker script arguments
	scriptArgs := &cloudfl.WorkersScriptArgs{
		AccountId:           pulumi.String(locals.CloudflareWorker.Spec.AccountId),
		Name:                pulumi.String(locals.CloudflareWorker.Spec.Script.Name),
		Content:             scriptContent,
		PlainTextBindings:   plainTextBindings,
		KvNamespaceBindings: kvBindings,
	}

	if locals.CloudflareWorker.Spec.CompatibilityDate != "" {
		scriptArgs.CompatibilityDate = pulumi.StringPtr(locals.CloudflareWorker.Spec.CompatibilityDate)
	}

	// Create the Worker script using new WorkersScript resource
	createdWorkerScript, err := cloudfl.NewWorkersScript(
		ctx,
		"workers-script",
		scriptArgs,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare workers script")
	}

	// Export stack output
	ctx.Export(OpScriptId, createdWorkerScript.ID())

	return createdWorkerScript, nil
}
