package module

import (
	civovolumev1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civovolume/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civovolumev1.CivoVolumeStackInput,
) error {
	return nil
}
