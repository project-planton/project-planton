package module

import (
	cloudflarecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/cloudflarecredential/v1"
	cloudflarer2bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarer2bucket/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals bundles handy references for the rest of the module.
type Locals struct {
	CloudflareCredentialSpec *cloudflarecredentialv1.CloudflareCredentialSpec
	CloudflareR2Bucket       *cloudflarer2bucketv1.CloudflareR2Bucket
}

// initializeLocals copies stack‑input fields into the Locals struct.
func initializeLocals(_ *pulumi.Context, stackInput *cloudflarer2bucketv1.CloudflareR2BucketStackInput) *Locals {
	locals := &Locals{}

	locals.CloudflareR2Bucket = stackInput.Target
	locals.CloudflareCredentialSpec = stackInput.ProviderCredential

	return locals
}
