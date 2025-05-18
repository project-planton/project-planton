package module

import (
	"fmt"
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmeta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// service creates an external LoadBalancer Service when spec.ingress.enabled is true.
// This is required because NATS clients use the NATS (TCP) protocol on port 4222,
// which cannot be terminated by standard HTTP ingress. A LoadBalancer keeps the
// architecture simple and mirrors the pattern from Terraform.
func service(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider pulumi.ProviderResource,
	createdNamespace *kubernetescorev1.Namespace) error {

	if locals.NatsKubernetes.Spec.Ingress == nil ||
		!locals.NatsKubernetes.Spec.Ingress.Enabled {
		// No external exposure requested.
		return nil
	}

	svcName := "nats-external-lb"

	selector := map[string]string{
		"app.kubernetes.io/instance": locals.NatsKubernetes.Metadata.Name,
		"app":                        "nats",
	}

	annotations := pulumi.StringMap{}
	if locals.NatsKubernetes.Spec.Ingress.DnsDomain != "" {
		host := fmt.Sprintf("%s.%s", locals.Namespace, locals.NatsKubernetes.Spec.Ingress.DnsDomain)
		annotations["external-dns.alpha.kubernetes.io/hostname"] = pulumi.String(host)
	}

	createdLoadBalancerService, err := kubernetescorev1.NewService(ctx,
		svcName,
		&kubernetescorev1.ServiceArgs{
			Metadata: &kubernetesmeta.ObjectMetaArgs{
				Name:        pulumi.String(svcName),
				Namespace:   createdNamespace.Metadata.Name(),
				Labels:      pulumi.ToStringMap(locals.Labels),
				Annotations: annotations,
			},
			Spec: &kubernetescorev1.ServiceSpecArgs{
				Type: pulumi.String("LoadBalancer"),
				Ports: kubernetescorev1.ServicePortArray{
					&kubernetescorev1.ServicePortArgs{
						Name:       pulumi.String("client"),
						Port:       pulumi.Int(vars.NatsClientPort),
						Protocol:   pulumi.String("TCP"),
						TargetPort: pulumi.Int(vars.NatsClientPort),
					},
				},
				Selector: pulumi.ToStringMap(selector),
			},
		}, pulumi.Provider(kubernetesProvider), pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create external LoadBalancer service")
	}

	// Export external client URL if not already set via DNS hostname.
	if locals.ClientURLExternal == "" {
		endpoint := createdLoadBalancerService.Status.ApplyT(func(st kubernetescorev1.ServiceStatus) string {
			if st.LoadBalancer == nil || len(st.LoadBalancer.Ingress) == 0 {
				return "" // unknown until LB allocates an IP/hostname
			}
			ingress := st.LoadBalancer.Ingress[0]
			host := ingress.Hostname
			if host == nil || *host == "" {
				host = ingress.Ip
			}
			if host == nil {
				return "" // still pending
			}
			return fmt.Sprintf("nats://%s:%d", *host, vars.NatsClientPort)
		}).(pulumi.StringOutput)

		ctx.Export(OpClientUrlExternal, endpoint)
	}

	return nil
}
