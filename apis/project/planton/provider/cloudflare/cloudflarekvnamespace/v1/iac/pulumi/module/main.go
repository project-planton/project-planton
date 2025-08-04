package module

import (
	cloudflarekvnamespacev1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflarekvnamespace/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the single public entry‑point that Project Planton’s CLI invokes.
func Resources(
	ctx *pulumi.Context,
	stackInput *cloudflarekvnamespacev1.CloudflareKvNamespaceStackInput,
) error {
	return nil
}
