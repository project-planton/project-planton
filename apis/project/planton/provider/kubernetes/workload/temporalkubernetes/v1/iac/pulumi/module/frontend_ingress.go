package module

import (
	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	kubernetesmetav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// frontendIngress installs a single external LoadBalancer Service that
// exposes the Temporal gRPC frontend.  HTTP traffic is handled separately in
// frontendHttpIngress.go via Gateway-API / Istio.
func frontendIngress(ctx *pulumi.Context, locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	if locals.TemporalKubernetes.Spec.Ingress == nil ||
		locals.TemporalKubernetes.Spec.Ingress.Frontend == nil ||
		!locals.TemporalKubernetes.Spec.Ingress.Frontend.Enabled ||
		locals.TemporalKubernetes.Spec.Ingress.Frontend.GrpcHostname == "" {
		// frontend gRPC ingress disabled – nothing to provision
		return nil
	}

	selector := map[string]string{
		"app.kubernetes.io/instance": locals.TemporalKubernetes.Metadata.Name,
		// Temporal Helm labels its pods with workload-name = “temporal-frontend”
		"app.kubernetes.io/component": "frontend",
	}

	_, err := kubernetescorev1.NewService(ctx,
		"frontend-external-lb",
		&kubernetescorev1.ServiceArgs{
			Metadata: &kubernetesmetav1.ObjectMetaArgs{
				Name:      pulumi.String("frontend-external-lb"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    createdNamespace.Metadata.Labels(),
				Annotations: pulumi.StringMap{
					"external-dns.alpha.kubernetes.io/hostname": pulumi.String(locals.IngressFrontendGrpcHostname),
				},
			},
			Spec: &kubernetescorev1.ServiceSpecArgs{
				Type: pulumi.String("LoadBalancer"),
				Ports: kubernetescorev1.ServicePortArray{
					&kubernetescorev1.ServicePortArgs{
						Name:       pulumi.String("grpc-frontend"),
						Port:       pulumi.Int(vars.FrontendGrpcPort),
						Protocol:   pulumi.String("TCP"),
						TargetPort: pulumi.Int(vars.FrontendGrpcPort),
					},
				},
				Selector: pulumi.ToStringMap(selector),
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "failed to create frontend gRPC load balancer service")
	}

	return nil
}
