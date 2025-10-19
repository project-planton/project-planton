package module

import (
	cloudflareworkerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflareworker/v1"
	cloudflareprovider "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles quick references copied from the stack‑input.
type Locals struct {
	CloudflareProviderConfig *cloudflareprovider.CloudflareProviderConfig
	CloudflareWorker         *cloudflareworkerv1.CloudflareWorker
}

// initializeLocals mirrors the pattern used in existing modules.
func initializeLocals(_ *pulumi.Context, stackInput *cloudflareworkerv1.CloudflareWorkerStackInput) *Locals {
	locals := &Locals{}
	locals.CloudflareWorker = stackInput.Target
	locals.CloudflareProviderConfig = stackInput.ProviderConfig
	return locals
}
