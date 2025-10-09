package module

import (
	"fmt"

	"github.com/pkg/errors"
	certmanagerv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/certmanager/kubernetes/cert_manager/v1"
	gatewayv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/gatewayapis/kubernetes/gateway/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createOtelCollectorIngress creates Kubernetes Gateway API resources for OpenTelemetry Collector external access
// This includes Certificate, Gateway, and HTTPRoute resources for both gRPC and HTTP endpoints
func createOtelCollectorIngress(ctx *pulumi.Context, locals *Locals, kubernetesProvider *kubernetes.Provider,
	createdNamespace *kubernetescorev1.Namespace) error {

	// Skip if ingress is not enabled
	if locals.SignozKubernetes.Spec.OtelCollectorIngress == nil ||
		!locals.SignozKubernetes.Spec.OtelCollectorIngress.Enabled ||
		locals.SignozKubernetes.Spec.OtelCollectorIngress.DnsDomain == "" {
		return nil
	}

	// Certificate for OTEL Collector endpoints (both gRPC and HTTP)
	otelCertName := fmt.Sprintf("cert-%s-otel", locals.Namespace)
	otelHostnames := []string{
		locals.OtelCollectorExternalGrpcHostname,
		locals.OtelCollectorExternalHttpHostname,
	}

	addedOtelCertificate, err := certmanagerv1.NewCertificate(ctx,
		"otel-ingress-certificate",
		&certmanagerv1.CertificateArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String(otelCertName),
				Namespace: pulumi.String(vars.IstioIngressNamespace),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: certmanagerv1.CertificateSpecArgs{
				DnsNames:   pulumi.ToStringArray(otelHostnames),
				SecretName: pulumi.String(otelCertName),
				IssuerRef: certmanagerv1.CertificateSpecIssuerRefArgs{
					Kind: pulumi.String("ClusterIssuer"),
					Name: pulumi.String(locals.IngressCertClusterIssuerName),
				},
			},
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "error creating OTEL Collector certificate")
	}

	// Gateway for OTEL Collector endpoints
	createdOtelGateway, err := gatewayv1.NewGateway(ctx,
		"otel-external",
		&gatewayv1.GatewayArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.Sprintf("%s-otel-external", locals.Namespace),
				Namespace: pulumi.String(vars.IstioIngressNamespace),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: gatewayv1.GatewaySpecArgs{
				GatewayClassName: pulumi.String(vars.GatewayIngressClassName),
				Addresses: gatewayv1.GatewaySpecAddressesArray{
					gatewayv1.GatewaySpecAddressesArgs{
						Type:  pulumi.String("Hostname"),
						Value: pulumi.String(vars.GatewayExternalLoadBalancerServiceHostname),
					},
				},
				Listeners: gatewayv1.GatewaySpecListenersArray{
					// HTTPS listener for gRPC endpoint (port 443, gRPC over TLS)
					&gatewayv1.GatewaySpecListenersArgs{
						Name:     pulumi.String("https-otel-grpc"),
						Hostname: pulumi.String(locals.OtelCollectorExternalGrpcHostname),
						Port:     pulumi.Int(443),
						Protocol: pulumi.String("HTTPS"),
						Tls: &gatewayv1.GatewaySpecListenersTlsArgs{
							Mode: pulumi.String("Terminate"),
							CertificateRefs: gatewayv1.GatewaySpecListenersTlsCertificateRefsArray{
								gatewayv1.GatewaySpecListenersTlsCertificateRefsArgs{
									Name: pulumi.String(otelCertName),
								},
							},
						},
						AllowedRoutes: gatewayv1.GatewaySpecListenersAllowedRoutesArgs{
							Namespaces: gatewayv1.GatewaySpecListenersAllowedRoutesNamespacesArgs{
								From: pulumi.String("All"),
							},
						},
					},
					// HTTPS listener for HTTP endpoint (port 443)
					&gatewayv1.GatewaySpecListenersArgs{
						Name:     pulumi.String("https-otel-http"),
						Hostname: pulumi.String(locals.OtelCollectorExternalHttpHostname),
						Port:     pulumi.Int(443),
						Protocol: pulumi.String("HTTPS"),
						Tls: &gatewayv1.GatewaySpecListenersTlsArgs{
							Mode: pulumi.String("Terminate"),
							CertificateRefs: gatewayv1.GatewaySpecListenersTlsCertificateRefsArray{
								gatewayv1.GatewaySpecListenersTlsCertificateRefsArgs{
									Name: pulumi.String(otelCertName),
								},
							},
						},
						AllowedRoutes: gatewayv1.GatewaySpecListenersAllowedRoutesArgs{
							Namespaces: gatewayv1.GatewaySpecListenersAllowedRoutesNamespacesArgs{
								From: pulumi.String("All"),
							},
						},
					},
				},
			},
		}, pulumi.Provider(kubernetesProvider), pulumi.DependsOn([]pulumi.Resource{addedOtelCertificate}))
	if err != nil {
		return errors.Wrap(err, "error creating OTEL Collector gateway")
	}

	// HTTPRoute for gRPC endpoint (routes to OTEL Collector gRPC port 4317)
	// Note: gRPC works over HTTP/2, so HTTPRoute is appropriate here
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"https-otel-grpc",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("https-otel-grpc"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.OtelCollectorExternalGrpcHostname)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s-otel-external", locals.Namespace),
						Namespace:   createdOtelGateway.Metadata.Namespace(),
						SectionName: pulumi.String("https-otel-grpc"),
					},
				},
				Rules: gatewayv1.HTTPRouteSpecRulesArray{
					gatewayv1.HTTPRouteSpecRulesArgs{
						Matches: gatewayv1.HTTPRouteSpecRulesMatchesArray{
							gatewayv1.HTTPRouteSpecRulesMatchesArgs{
								Path: gatewayv1.HTTPRouteSpecRulesMatchesPathArgs{
									Type:  pulumi.String("PathPrefix"),
									Value: pulumi.String("/"),
								},
							},
						},
						BackendRefs: gatewayv1.HTTPRouteSpecRulesBackendRefsArray{
							gatewayv1.HTTPRouteSpecRulesBackendRefsArgs{
								// Route to OTEL Collector gRPC port (4317)
								Name:      pulumi.Sprintf("%s-otel-collector", locals.SignozKubernetes.Metadata.Name),
								Namespace: createdNamespace.Metadata.Name(),
								Port:      pulumi.Int(vars.OtelGrpcPort),
							},
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "error creating HTTPRoute for OTEL Collector gRPC endpoint")
	}

	// HTTPRoute for HTTP endpoint (routes to OTEL Collector HTTP port 4318)
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"https-otel-http",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("https-otel-http"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.OtelCollectorExternalHttpHostname)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s-otel-external", locals.Namespace),
						Namespace:   createdOtelGateway.Metadata.Namespace(),
						SectionName: pulumi.String("https-otel-http"),
					},
				},
				Rules: gatewayv1.HTTPRouteSpecRulesArray{
					gatewayv1.HTTPRouteSpecRulesArgs{
						Matches: gatewayv1.HTTPRouteSpecRulesMatchesArray{
							gatewayv1.HTTPRouteSpecRulesMatchesArgs{
								Path: gatewayv1.HTTPRouteSpecRulesMatchesPathArgs{
									Type:  pulumi.String("PathPrefix"),
									Value: pulumi.String("/"),
								},
							},
						},
						BackendRefs: gatewayv1.HTTPRouteSpecRulesBackendRefsArray{
							gatewayv1.HTTPRouteSpecRulesBackendRefsArgs{
								// Route to OTEL Collector HTTP port (4318)
								Name:      pulumi.Sprintf("%s-otel-collector", locals.SignozKubernetes.Metadata.Name),
								Namespace: createdNamespace.Metadata.Name(),
								Port:      pulumi.Int(vars.OtelHttpPort),
							},
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace))
	if err != nil {
		return errors.Wrap(err, "error creating HTTPRoute for OTEL Collector HTTP endpoint")
	}

	return nil
}
