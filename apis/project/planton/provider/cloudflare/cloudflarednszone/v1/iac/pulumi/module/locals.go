package module

import (
	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	cloudflarednszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarednszone/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles the bits we need everywhere else.
type Locals struct {
	CloudflareCredentialSpec *cloudflarecredentialv1.CloudflareCredentialSpec
	CloudflareDnsZone        *cloudflarednszonev1.CloudflareDnsZone
}

// initializeLocals copies fields from the stack‑input into Locals.
func initializeLocals(_ *pulumi.Context, stackInput *cloudflarednszonev1.CloudflareDnsZoneStackInput) *Locals {
	locals := &Locals{}

	locals.CloudflareDnsZone = stackInput.Target
	locals.CloudflareCredentialSpec = stackInput.ProviderCredential

	return locals
}
