package module

import (
	"github.com/pkg/errors"
	cloudflareloadbalancerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/cloudflare/cloudflareloadbalancer/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/cloudflare/pulumicloudflareprovider"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources is the module entry‑point. It prepares locals, sets up the provider,
// then provisions the load balancer.
func Resources(
	ctx *pulumi.Context,
	stackInput *cloudflareloadbalancerv1.CloudflareLoadBalancerStackInput,
) error {
	// 1. Gather handy references.
	locals := initializeLocals(ctx, stackInput)

	// 2. Create a Cloudflare provider from the supplied credential.
	cloudflareProvider, err := pulumicloudflareprovider.Get(
		ctx,
		stackInput.ProviderConfig,
	)
	if err != nil {
		return errors.Wrap(err, "failed to setup cloudflare provider")
	}

	// 3. Provision the load balancer & its pool / monitor.
	if _, err := load_balancer(ctx, locals, cloudflareProvider); err != nil {
		return errors.Wrap(err, "failed to create cloudflare load balancer")
	}

	return nil
}
