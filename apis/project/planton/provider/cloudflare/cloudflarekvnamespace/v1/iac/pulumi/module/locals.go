package module

import (
	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	cloudflarekvnamespacev1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarekvnamespace/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references used across the module.
type Locals struct {
	CloudflareCredentialSpec *cloudflarecredentialv1.CloudflareCredentialSpec
	CloudflareKvNamespace    *cloudflarekvnamespacev1.CloudflareKvNamespace
}

// initializeLocals copies stack‑input fields into the Locals struct.
func initializeLocals(_ *pulumi.Context, stackInput *cloudflarekvnamespacev1.CloudflareKvNamespaceStackInput) *Locals {
	locals := &Locals{}
	locals.CloudflareKvNamespace = stackInput.Target
	locals.CloudflareCredentialSpec = stackInput.ProviderCredential
	return locals
}
