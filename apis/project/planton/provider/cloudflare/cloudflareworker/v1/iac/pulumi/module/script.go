package module

import (
	"github.com/pkg/errors"
	cloudfl "github.com/pulumi/pulumi-cloudflare/sdk/v5/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// worker_script provisions the Cloudflare Worker script and exports its ID.
func worker_script(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudfl.Provider,
) (*cloudfl.WorkerScript, error) {

	//// 1. KV namespace bindings (if any).
	//var kvBindings cloudfl.WorkerScriptKvNamespaceBindingArray
	//for _, _ := range locals.CloudflareWorker.Spec.KvBindings {
	//	kvBindings = append(kvBindings, cloudfl.WorkerScriptKvNamespaceBindingArgs{
	//		//NamespaceId:  pulumi.String(bindingRef.),
	//	})
	//}

	// 2. Plain‑text environment variables.
	plainTextBindings := pulumi.StringMap{}
	for k, v := range locals.CloudflareWorker.Spec.EnvVars {
		plainTextBindings[k] = pulumi.String(v)
	}

	// 3. Usage model (bundled vs unbound).

	// 4. Build arguments directly from proto fields.
	scriptArgs := &cloudfl.WorkerScriptArgs{
		AccountId: pulumi.String(locals.CloudflareCredentialSpec.AccountId),
		Name:      pulumi.String(locals.CloudflareWorker.Spec.ScriptName),
		Content:   pulumi.String(locals.CloudflareWorker.Spec.ScriptSource.GetValue()),
	}
	if locals.CloudflareWorker.Spec.CompatibilityDate != "" {
		scriptArgs.CompatibilityDate = pulumi.StringPtr(locals.CloudflareWorker.Spec.CompatibilityDate)
	}

	// 5. Create the Worker script.
	createdWorkerScript, err := cloudfl.NewWorkerScript(
		ctx,
		"worker-script",
		scriptArgs,
		pulumi.Provider(cloudflareProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create cloudflare worker script")
	}

	// 6. Export stack output.
	ctx.Export(OpScriptId, createdWorkerScript.ID())

	return createdWorkerScript, nil
}
