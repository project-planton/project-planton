package module

import (
	civofirewallv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civofirewall/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civofirewallv1.CivoFirewallStackInput,
) error {
	return nil
}
