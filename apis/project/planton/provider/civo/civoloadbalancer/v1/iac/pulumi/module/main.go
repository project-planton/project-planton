package module

import (
	civoloadbalancerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civoloadbalancer/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry point.
func Resources(
	ctx *pulumi.Context,
	stackInput *civoloadbalancerv1.CivoLoadBalancerStackInput,
) error {
	return nil
}
