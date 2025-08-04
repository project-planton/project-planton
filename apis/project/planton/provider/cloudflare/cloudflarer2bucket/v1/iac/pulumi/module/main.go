package module

import (
	cloudflarer2bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarer2bucket/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the single public entry‑point that Project Planton’s CLI invokes.
func Resources(
	ctx *pulumi.Context,
	stackInput *cloudflarer2bucketv1.CloudflareR2BucketStackInput,
) error {
	return nil
}
