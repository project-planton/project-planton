package module

import (
	civovpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civovpc/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civovpcv1.CivoVpcStackInput,
) error {
	return nil
}
