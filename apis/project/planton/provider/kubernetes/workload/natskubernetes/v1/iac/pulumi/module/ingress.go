// ingress.go
package module

import (
	"fmt"

	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmeta "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ingress creates an external LoadBalancer Service when spec.ingress.enabled is true.
// A previous version used label selector {app: nats,…} but the Helm chart labels
// pods as app.kubernetes.io/name=nats and app.kubernetes.io/component=nats.  The
// mismatch meant the Service had **zero endpoints**, so the LB IP could not reach
// any pod.  We now align the selector with the chart.
func ingress(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider pulumi.ProviderResource,
	createdNamespace *kubernetescorev1.Namespace) error {

	if locals.NatsKubernetes.Spec.Ingress == nil ||
		!locals.NatsKubernetes.Spec.Ingress.Enabled {
		return nil // no external exposure requested
	}

	svcName := "nats-external-lb"

	selector := map[string]string{
		// Matches labels applied by the official nats Helm chart v1.x
		"app.kubernetes.io/name":      "nats",
		"app.kubernetes.io/component": "nats",
		"app.kubernetes.io/instance":  locals.NatsKubernetes.Metadata.Name,
	}

	annotations := pulumi.StringMap{}
	if locals.NatsKubernetes.Spec.Ingress.Hostname != "" {
		annotations["external-dns.alpha.kubernetes.io/hostname"] = pulumi.String(locals.NatsKubernetes.Spec.Ingress.Hostname)
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
		return errors.Wrap(err, "failed to create external LoadBalancer ingress")
	}

	// Export external client URL if not already set via DNS hostname.
	if locals.ClientURLExternal == "" {
		endpoint := createdLoadBalancerService.Status.ApplyT(func(st kubernetescorev1.ServiceStatus) string {
			if st.LoadBalancer == nil || len(st.LoadBalancer.Ingress) == 0 {
				return "" // pending
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
