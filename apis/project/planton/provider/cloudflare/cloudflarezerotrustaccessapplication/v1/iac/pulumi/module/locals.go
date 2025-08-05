package module

import (
	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	cloudflarezerotrustaccessapplicationv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarezerotrustaccessapplication/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles convenient shortcuts for the rest of the module.
type Locals struct {
	CloudflareCredentialSpec             *cloudflarecredentialv1.CloudflareCredentialSpec
	CloudflareZeroTrustAccessApplication *cloudflarezerotrustaccessapplicationv1.CloudflareZeroTrustAccessApplication
}

// initializeLocals copies stack‑input fields into Locals.
func initializeLocals(
	_ *pulumi.Context,
	stackInput *cloudflarezerotrustaccessapplicationv1.CloudflareZeroTrustAccessApplicationStackInput,
) *Locals {
	locals := &Locals{}
	locals.CloudflareZeroTrustAccessApplication = stackInput.Target
	locals.CloudflareCredentialSpec = stackInput.ProviderCredential
	return locals
}
