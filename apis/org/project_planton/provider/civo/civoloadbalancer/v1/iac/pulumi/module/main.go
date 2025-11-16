package module

import (
	"github.com/pkg/errors"
	civoloadbalancerv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/civo/civoloadbalancer/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/civo/pulumicivoprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the Pulumi entry-point that ProjectPlanton invokes.
func Resources(
	ctx *pulumi.Context,
	stackInput *civoloadbalancerv1.CivoLoadBalancerStackInput,
) error {
	// 1. Collect useful references from stack-input.
	locals := initializeLocals(ctx, stackInput)

	// 2. Build a Civo provider from the supplied credential.
	civoProvider, err := pulumicivoprovider.Get(ctx, stackInput.ProviderConfig)
	if err != nil {
		return errors.Wrap(err, "failed to setup Civo provider")
	}

	// 3. Provision the load balancer.
	if _, err := loadBalancer(ctx, locals, civoProvider); err != nil {
		return errors.Wrap(err, "failed to create Civo load balancer")
	}

	return nil
}
