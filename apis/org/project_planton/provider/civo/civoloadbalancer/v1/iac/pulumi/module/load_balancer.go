package module

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	civoloadbalancerv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/civo/civoloadbalancer/v1"
	"github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// loadBalancer provisions the Civo Load Balancer and exports outputs.
func loadBalancer(
	ctx *pulumi.Context,
	locals *Locals,
	civoProvider *civo.Provider,
) (*civo.LoadBalancer, error) {

	spec := locals.CivoLoadBalancer.Spec

	// 1. Translate protocol enums to Civo-compatible strings
	translateProtocol := func(proto civoloadbalancerv1.CivoLoadBalancerProtocol) (string, error) {
		switch proto {
		case civoloadbalancerv1.CivoLoadBalancerProtocol_http:
			return "http", nil
		case civoloadbalancerv1.CivoLoadBalancerProtocol_https:
			return "https", nil
		case civoloadbalancerv1.CivoLoadBalancerProtocol_tcp:
			return "tcp", nil
		default:
			return "", errors.Errorf("unsupported protocol: %v", proto)
		}
	}

	// 2. Build backend configurations from instance IDs or tag
	var backends civo.LoadBalancerBackendArray

	if len(spec.InstanceIds) > 0 {
		// Use explicit instance IDs
		for _, instanceRef := range spec.InstanceIds {
			instanceId := instanceRef.GetValue()
			if instanceId == "" {
				continue
			}

			// For each forwarding rule, create a backend entry
			for _, rule := range spec.ForwardingRules {
				targetProtocol, err := translateProtocol(rule.TargetProtocol)
				if err != nil {
					return nil, err
				}

				backends = append(backends, civo.LoadBalancerBackendArgs{
					InstanceId: pulumi.String(instanceId),
					Protocol:   pulumi.String(targetProtocol),
					SourcePort: pulumi.Int(int(rule.EntryPort)),
					TargetPort: pulumi.Int(int(rule.TargetPort)),
				})
			}
		}
	} else if spec.InstanceTag != "" {
		// Use tag-based selection
		// For tag-based backends, Civo attaches instances automatically
		for _, rule := range spec.ForwardingRules {
			targetProtocol, err := translateProtocol(rule.TargetProtocol)
			if err != nil {
				return nil, err
			}

			backends = append(backends, civo.LoadBalancerBackendArgs{
				InstanceId: pulumi.String("tag:" + spec.InstanceTag),
				Protocol:   pulumi.String(targetProtocol),
				SourcePort: pulumi.Int(int(rule.EntryPort)),
				TargetPort: pulumi.Int(int(rule.TargetPort)),
			})
		}
	} else {
		return nil, errors.New("either instance_ids or instance_tag must be specified")
	}

	// 3. Build load balancer arguments
	lbArgs := &civo.LoadBalancerArgs{
		Name:      pulumi.String(spec.LoadBalancerName),
		Region:    pulumi.String(strings.ToLower(spec.Region.String())),
		NetworkId: pulumi.String(spec.Network.GetValue()),
		Backends:  backends,
		Algorithm: pulumi.String("round_robin"), // Default algorithm
	}

	// Optional: Enable sticky sessions via algorithm
	if spec.EnableStickySessions {
		lbArgs.Algorithm = pulumi.String("ip_hash")
	}

	// Optional: Reserved IP
	if spec.ReservedIpId != nil && spec.ReservedIpId.GetValue() != "" {
		lbArgs.ReservedIpId = pulumi.String(spec.ReservedIpId.GetValue())
	}

	// Optional: Health check configuration
	if spec.HealthCheck != nil {
		healthCheckProtocol, err := translateProtocol(spec.HealthCheck.Protocol)
		if err != nil {
			return nil, err
		}

		lbArgs.HealthCheckPath = pulumi.String(spec.HealthCheck.Path)

		// Note: Civo's LoadBalancer resource uses backends for health check configuration
		// The health check is applied to all backends automatically
		// If we need per-backend health checks, they're configured in the backend definition

		// Update first backend with health check info if needed
		if len(backends) > 0 {
			// The health check path is set at LB level
			// Protocol and port are inferred from backend configuration
			_ = healthCheckProtocol // Used for validation
		}
	}

	// 4. Create the Load Balancer
	createdLB, err := civo.NewLoadBalancer(
		ctx,
		"load-balancer",
		lbArgs,
		pulumi.Provider(civoProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Civo load balancer")
	}

	// 5. Export stack outputs
	ctx.Export(OpLoadBalancerId, createdLB.ID())
	ctx.Export(OpPublicIp, createdLB.PublicIp)
	ctx.Export(OpDnsName, createdLB.Hostname)
	ctx.Export(OpState, createdLB.State)

	return createdLB, nil
}
