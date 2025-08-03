package module

import (
	civocertificatev1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civocertificate/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
// WARNING: looks like there is no terraform/pulumi support for Civo Certificate
// https://registry.terraform.io/providers/civo/civo/latest/docs
func Resources(
	ctx *pulumi.Context,
	stackInput *civocertificatev1.CivoCertificateStackInput,
) error {
	return nil
}
