package module

import (
	"fmt"

	"github.com/pkg/errors"
	certmanagerv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/certmanager/kubernetes/cert_manager/v1"
	gatewayv1 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/gatewayapis/kubernetes/gateway/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// createOtelCollectorIngress creates Kubernetes Gateway API resources for OpenTelemetry Collector external access
// This includes Certificate, Gateway, and HTTPRoute resources for HTTP endpoint
func createOtelCollectorIngress(ctx *pulumi.Context, locals *Locals,
	kubernetesProvider pulumi.ProviderResource) error {

	// Skip if ingress is not enabled
	if locals.KubernetesSignoz.Spec.Ingress == nil ||
		locals.KubernetesSignoz.Spec.Ingress.OtelCollector == nil ||
		!locals.KubernetesSignoz.Spec.Ingress.OtelCollector.Enabled ||
		locals.KubernetesSignoz.Spec.Ingress.OtelCollector.Hostname == "" {
		return nil
	}

	// Certificate for OTEL Collector HTTP endpoint
	otelCertName := fmt.Sprintf("cert-%s-otel-http", locals.Namespace)
	otelHostnames := []string{
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

	// Gateway for OTEL Collector HTTP endpoint
	createdOtelGateway, err := gatewayv1.NewGateway(ctx,
		"otel-http-external",
		&gatewayv1.GatewayArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.Sprintf("%s-otel-http-external", locals.Namespace),
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
		return errors.Wrap(err, "error creating OTEL Collector HTTP gateway")
	}

	// HTTPRoute for HTTP endpoint (routes to OTEL Collector HTTP port 4318)
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"https-otel-http",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("https-otel-http"),
				Namespace: pulumi.String(locals.Namespace),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.OtelCollectorExternalHttpHostname)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s-otel-http-external", locals.Namespace),
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
								Name:      pulumi.Sprintf("%s-otel-collector", locals.KubernetesSignoz.Metadata.Name),
								Namespace: pulumi.String(locals.Namespace),
								Port:      pulumi.Int(vars.OtelHttpPort),
							},
						},
					},
				},
			},
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "error creating HTTPRoute for OTEL Collector HTTP endpoint")
	}

	return nil
}
