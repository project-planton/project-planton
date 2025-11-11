package module

import (
	civoloadbalancerv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/civo/civoloadbalancer/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources can not be implemented since the civo provider currently does not support creating load balancers.
// https://registry.terraform.io/providers/civo/civo/latest/docs/data-sources/loadbalancer
func Resources(
	ctx *pulumi.Context,
	stackInput *civoloadbalancerv1.CivoLoadBalancerStackInput,
) error {
	return nil
}
