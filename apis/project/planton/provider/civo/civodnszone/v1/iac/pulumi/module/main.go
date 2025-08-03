package module

import (
	civodnszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civodnszone/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civodnszonev1.CivoDnsZoneStackInput,
) error {
	return nil
}
