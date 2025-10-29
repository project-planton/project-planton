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

	// Build bindings array (unified in v6 API)
	// Includes both plain-text environment variables and KV namespace bindings
	var bindings cloudfl.WorkersScriptBindingArray

	// Add plain-text environment variable bindings
	// Note: env.secrets are uploaded separately via Cloudflare Secrets API
	if locals.CloudflareWorker.Spec.Env != nil {
		for k, v := range locals.CloudflareWorker.Spec.Env.Variables {
			bindings = append(bindings, cloudfl.WorkersScriptBindingArgs{
				Name: pulumi.String(k),
				Type: pulumi.String("plain_text"),
				Text: pulumi.String(v),
			})
		}
	}

	// Add KV namespace bindings
	for _, kvBinding := range locals.CloudflareWorker.Spec.KvBindings {
		bindings = append(bindings, cloudfl.WorkersScriptBindingArgs{
			Name:        pulumi.String(kvBinding.Name),
			Type:        pulumi.String("kv_namespace"),
			NamespaceId: pulumi.String(kvBinding.GetFieldPath()),
		})
	}

	// Build Worker script arguments
	// Use MainModule for module syntax workers (with import/export)
	// Content is only for service worker syntax (with addEventListener)
	scriptArgs := &cloudfl.WorkersScriptArgs{
		AccountId:          pulumi.String(locals.CloudflareWorker.Spec.AccountId),
		ScriptName:         pulumi.String(locals.CloudflareWorker.Spec.Script.Name),
		MainModule:         pulumi.String("index.js"), // Indicates this is a module worker
		Content:            scriptContent,             // The actual module code
		Bindings:           bindings,
		CompatibilityFlags: pulumi.StringArray{pulumi.String("nodejs_compat")}, // Enable Node.js compatibility
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
