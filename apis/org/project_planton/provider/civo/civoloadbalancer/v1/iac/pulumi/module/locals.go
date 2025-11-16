package module

import (
	"strconv"

	civoprovider "github.com/project-planton/project-planton/apis/org/project_planton/provider/civo"
	civoloadbalancerv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/civo/civoloadbalancer/v1"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Locals groups frequently-used values for the load balancer module.
type Locals struct {
	CivoProviderConfig *civoprovider.CivoProviderConfig
	CivoLoadBalancer   *civoloadbalancerv1.CivoLoadBalancer
	CivoLabels         map[string]string
}

// initializeLocals copies the stack-input into Locals and derives labels.
func initializeLocals(_ *pulumi.Context, stackInput *civoloadbalancerv1.CivoLoadBalancerStackInput) *Locals {
	locals := &Locals{}
	locals.CivoLoadBalancer = stackInput.Target
	locals.CivoProviderConfig = stackInput.ProviderConfig

	// Basic Planton labels (for consistency with other Civo resources)
	locals.CivoLabels = map[string]string{
		"planton:resource":      strconv.FormatBool(true),
		"planton:resource_name": locals.CivoLoadBalancer.Metadata.Name,
		"planton:resource_kind": cloudresourcekind.CloudResourceKind_CivoLoadBalancer.String(),
	}
	if locals.CivoLoadBalancer.Metadata.Org != "" {
		locals.CivoLabels["planton:organization"] = locals.CivoLoadBalancer.Metadata.Org
	}
	if locals.CivoLoadBalancer.Metadata.Env != "" {
		locals.CivoLabels["planton:environment"] = locals.CivoLoadBalancer.Metadata.Env
	}
	if locals.CivoLoadBalancer.Metadata.Id != "" {
		locals.CivoLabels["planton:resource_id"] = locals.CivoLoadBalancer.Metadata.Id
	}

	return locals
}

