package module

import (
	"fmt"

	"github.com/pulumi/pulumi-cloudflare/sdk/v6/go/cloudflare"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// load_balancer provisions:
//  1. A Monitor (health check)
//  2. A Pool (collection of origins)
//  3. The Load Balancer itself
//
// It also exports stack outputs defined in outputs.go.
func load_balancer(
	ctx *pulumi.Context,
	locals *Locals,
	cloudflareProvider *cloudflare.Provider,
) (*cloudflare.LoadBalancer, error) {

	// ---------------------------------------------------------------------
	// 1. Monitor (health check) – uses the health_probe_path from the spec.
	// ---------------------------------------------------------------------
	createdMonitor, err := cloudflare.NewLoadBalancerMonitor(ctx, "monitor",
		&cloudflare.LoadBalancerMonitorArgs{
			Type:          pulumi.String("http"),
			Method:        pulumi.String("GET"),
			Path:          pulumi.String(locals.CloudflareLoadBalancer.Spec.HealthProbePath),
			Retries:       pulumi.Int(2),
			Timeout:       pulumi.Int(5),
			ExpectedCodes: pulumi.String("2xx"),
		}, pulumi.Provider(cloudflareProvider))
	if err != nil {
		return nil, fmt.Errorf("failed to create monitor: %w", err)
	}

	// ---------------------------------------------------------------------
	// 2. Pool – one pool containing all declared origins.
	// ---------------------------------------------------------------------
	var poolOrigins cloudflare.LoadBalancerPoolOriginArray
	for _, o := range locals.CloudflareLoadBalancer.Spec.Origins {
		poolOrigins = append(poolOrigins, cloudflare.LoadBalancerPoolOriginArgs{
			Name:    pulumi.String(o.Name),
			Address: pulumi.String(o.Address),
			//Weight:  pulumi.Float64Ptr(o.Weight),
		})
	}

	createdPool, err := cloudflare.NewLoadBalancerPool(ctx, "pool", &cloudflare.LoadBalancerPoolArgs{
		Name:    pulumi.String(locals.CloudflareLoadBalancer.Metadata.Name + "-pool"),
		Origins: poolOrigins,
		Monitor: createdMonitor.ID(),
		Enabled: pulumi.Bool(true),
	}, pulumi.Provider(cloudflareProvider))
	if err != nil {
		return nil, fmt.Errorf("failed to create pool: %w", err)
	}

	// ---------------------------------------------------------------------
	// 3. Load Balancer – wires everything together.
	// ---------------------------------------------------------------------
	// Map enum → string for steering policy.
	var steering pulumi.StringPtrInput
	switch locals.CloudflareLoadBalancer.Spec.SteeringPolicy {
	case 1: // STEERING_GEO
		steering = pulumi.StringPtr("geo")
	case 2: // STEERING_RANDOM
		steering = pulumi.StringPtr("random")
	default:
		steering = pulumi.StringPtr("off")
	}

	// Map enum → string for session affinity.
	var affinity pulumi.StringPtrInput
	switch locals.CloudflareLoadBalancer.Spec.SessionAffinity {
	case 1: // SESSION_AFFINITY_COOKIE
		affinity = pulumi.StringPtr("cookie")
	default:
		affinity = pulumi.StringPtr("none")
	}

	createdLoadBalancer, err := cloudflare.NewLoadBalancer(ctx, "load_balancer", &cloudflare.LoadBalancerArgs{
		ZoneId:          pulumi.String(locals.CloudflareLoadBalancer.Spec.ZoneId.GetValue()),
		Name:            pulumi.String(locals.CloudflareLoadBalancer.Spec.Hostname),
		DefaultPoolIds:  pulumi.StringArray{createdPool.ID()},
		FallbackPoolId:  createdPool.ID(),
		Proxied:         pulumi.BoolPtr(locals.CloudflareLoadBalancer.Spec.Proxied),
		SteeringPolicy:  steering,
		SessionAffinity: affinity,
	}, pulumi.Provider(cloudflareProvider))
	if err != nil {
		return nil, fmt.Errorf("failed to create load balancer: %w", err)
	}

	// ---------------------------------------------------------------------
	// 4. Stack outputs.
	// ---------------------------------------------------------------------
	ctx.Export(OpLoadBalancerId, createdLoadBalancer.ID())
	ctx.Export(OpLoadBalancerDnsRecordName, pulumi.String(locals.CloudflareLoadBalancer.Spec.Hostname))
	ctx.Export(OpLoadBalancerCnameTarget, createdLoadBalancer.Name)

	return createdLoadBalancer, nil
}
