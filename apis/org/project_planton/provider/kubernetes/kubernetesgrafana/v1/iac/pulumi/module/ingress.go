package module

import (
	"fmt"

	"github.com/pkg/errors"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	kubernetesnetworkingv1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/networking/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ingress(ctx *pulumi.Context,
	locals *Locals,
	createdNamespace *kubernetescorev1.Namespace) error {

	if locals.KubernetesGrafana.Spec.Ingress == nil ||
		!locals.KubernetesGrafana.Spec.Ingress.Enabled {
		return nil
	}

	// Extract external hostname without https:// prefix
	externalHost := ""
	if locals.IngressExternalHostname != "" {
		// Remove https:// prefix if present
		externalHost = locals.KubernetesGrafana.Spec.Ingress.Hostname
	}

	// Extract internal hostname without https:// prefix
	internalHost := ""
	if locals.IngressInternalHostname != "" {
		internalHost = fmt.Sprintf("internal-%s", locals.KubernetesGrafana.Spec.Ingress.Hostname)
	}

	pathType := "Prefix"

	// Create external ingress if hostname is provided
	if externalHost != "" {
		_, err := kubernetesnetworkingv1.NewIngress(ctx,
			fmt.Sprintf("%s-external", locals.KubernetesGrafana.Metadata.Name),
			&kubernetesnetworkingv1.IngressArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name:      pulumi.String(fmt.Sprintf("%s-external", locals.KubernetesGrafana.Metadata.Name)),
					Namespace: createdNamespace.Metadata.Name(),
					Annotations: pulumi.StringMap{
						"kubernetes.io/ingress.class": pulumi.String("nginx"),
					},
				},
				Spec: &kubernetesnetworkingv1.IngressSpecArgs{
					Rules: kubernetesnetworkingv1.IngressRuleArray{
						&kubernetesnetworkingv1.IngressRuleArgs{
							Host: pulumi.String(externalHost),
							Http: &kubernetesnetworkingv1.HTTPIngressRuleValueArgs{
								Paths: kubernetesnetworkingv1.HTTPIngressPathArray{
									&kubernetesnetworkingv1.HTTPIngressPathArgs{
										Path:     pulumi.String("/"),
										PathType: pulumi.String(pathType),
										Backend: &kubernetesnetworkingv1.IngressBackendArgs{
											Service: &kubernetesnetworkingv1.IngressServiceBackendArgs{
												Name: pulumi.String(locals.KubeServiceName),
												Port: &kubernetesnetworkingv1.ServiceBackendPortArgs{
													Number: pulumi.Int(80),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "failed to create external ingress")
		}
	}

	// Create internal ingress if hostname is provided
	if internalHost != "" {
		_, err := kubernetesnetworkingv1.NewIngress(ctx,
			fmt.Sprintf("%s-internal", locals.KubernetesGrafana.Metadata.Name),
			&kubernetesnetworkingv1.IngressArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name:      pulumi.String(fmt.Sprintf("%s-internal", locals.KubernetesGrafana.Metadata.Name)),
					Namespace: createdNamespace.Metadata.Name(),
					Annotations: pulumi.StringMap{
						"kubernetes.io/ingress.class": pulumi.String("nginx-internal"),
					},
				},
				Spec: &kubernetesnetworkingv1.IngressSpecArgs{
					Rules: kubernetesnetworkingv1.IngressRuleArray{
						&kubernetesnetworkingv1.IngressRuleArgs{
							Host: pulumi.String(internalHost),
							Http: &kubernetesnetworkingv1.HTTPIngressRuleValueArgs{
								Paths: kubernetesnetworkingv1.HTTPIngressPathArray{
									&kubernetesnetworkingv1.HTTPIngressPathArgs{
										Path:     pulumi.String("/"),
										PathType: pulumi.String(pathType),
										Backend: &kubernetesnetworkingv1.IngressBackendArgs{
											Service: &kubernetesnetworkingv1.IngressServiceBackendArgs{
												Name: pulumi.String(locals.KubeServiceName),
												Port: &kubernetesnetworkingv1.ServiceBackendPortArgs{
													Number: pulumi.Int(80),
												},
											},
										},
									},
								},
							},
						},
					},
				},
			}, pulumi.Parent(createdNamespace))
		if err != nil {
			return errors.Wrap(err, "failed to create internal ingress")
		}
	}

	return nil
}
