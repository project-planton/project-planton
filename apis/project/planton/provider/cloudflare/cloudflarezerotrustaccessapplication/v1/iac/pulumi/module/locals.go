package module

import (
	cloudflarezerotrustaccessapplicationv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarezerotrustaccessapplication/v1"
	cloudflareprovider "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles convenient shortcuts for the rest of the module.
type Locals struct {
	CloudflareProviderConfig             *cloudflareprovider.CloudflareProviderConfig
	CloudflareZeroTrustAccessApplication *cloudflarezerotrustaccessapplicationv1.CloudflareZeroTrustAccessApplication
}

// initializeLocals copies stack‑input fields into Locals.
func initializeLocals(
	_ *pulumi.Context,
	stackInput *cloudflarezerotrustaccessapplicationv1.CloudflareZeroTrustAccessApplicationStackInput,
) *Locals {
	locals := &Locals{}
	locals.CloudflareZeroTrustAccessApplication = stackInput.Target
	locals.CloudflareProviderConfig = stackInput.ProviderConfig
	return locals
}
