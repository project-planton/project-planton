package module

import (
	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	cloudflared1databasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflared1database/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	CloudflareCredentialSpec *cloudflarecredentialv1.CloudflareCredentialSpec
	CloudflareD1Database     *cloudflared1databasev1.CloudflareD1Database
}

// initializeLocals copies stack‑input fields into the Locals struct.
// Mirrors the style used in other Project Planton modules.
func initializeLocals(_ *pulumi.Context, stackInput *cloudflared1databasev1.CloudflareD1DatabaseStackInput) *Locals {
	locals := &Locals{}
	locals.CloudflareD1Database = stackInput.Target
	locals.CloudflareCredentialSpec = stackInput.ProviderCredential
	return locals
}
