package module

import (
	civoipaddressv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civoipaddress/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civoipaddressv1.CivoIpAddressStackInput,
) error {
	return nil
}
