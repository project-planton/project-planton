package module

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// loadBalancerIngress creates external and internal load balancer Services
// for exposing Neo4j via HTTP (7474) and Bolt (7687).
func loadBalancerIngress(
	ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace,
) error {
	// External LB service
	_, err := kubernetescorev1.NewService(ctx,
		"ingress-external-lb",
		&kubernetescorev1.ServiceArgs{
			Metadata: &kubernetesmetav1.ObjectMetaArgs{
				Name:      pulumi.String("ingress-external-lb"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    createdNamespace.Metadata.Labels(),
				Annotations: pulumi.StringMap{
					"external-dns.alpha.kubernetes.io/hostname": pulumi.String(locals.IngressExternalHostname),
				},
			},
			Spec: &kubernetescorev1.ServiceSpecArgs{
				Type: pulumi.String("LoadBalancer"),
				Ports: kubernetescorev1.ServicePortArray{
					// HTTP (Neo4j Browser)
					&kubernetescorev1.ServicePortArgs{
						Name:       pulumi.String("http-neo4j"),
						Port:       pulumi.Int(7474),
						Protocol:   pulumi.String("TCP"),
						TargetPort: pulumi.String("http"),
					},
					// Bolt
					&kubernetescorev1.ServicePortArgs{
						Name:       pulumi.String("bolt-neo4j"),
						Port:       pulumi.Int(7687),
						Protocol:   pulumi.String("TCP"),
						TargetPort: pulumi.String("bolt"),
					},
				},
				Selector: pulumi.ToStringMap(locals.Neo4jPodSelectorLabels),
			},
		},
		pulumi.Parent(createdNamespace),
	)
	if err != nil {
		return errors.Wrapf(err, "failed to create external load balancer service")
	}
	return nil
}
