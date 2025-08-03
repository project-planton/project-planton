package module

import (
	civodatabasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civodatabase/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civodatabasev1.CivoDatabaseStackInput,
) error {
	return nil
}
