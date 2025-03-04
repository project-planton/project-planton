package module

import (
	"github.com/pkg/errors"
	certmanagerv1 "github.com/project-planton/project-planton/pkg/kubernetestypes/certmanager/kubernetes/cert_manager/v1"
	gatewayv1 "github.com/project-planton/project-planton/pkg/kubernetestypes/gatewayapis/kubernetes/gateway/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	kubernetescorev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ingress(ctx *pulumi.Context, locals *Locals, kubernetesProvider *kubernetes.Provider,
	createdNamespace *kubernetescorev1.Namespace) error {
	//crate new certificate
	addedCertificate, err := certmanagerv1.NewCertificate(ctx,
		"ingress-certificate",
		&certmanagerv1.CertificateArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.Namespace),
				Namespace: pulumi.String(vars.IstioIngressNamespace),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: certmanagerv1.CertificateSpecArgs{
				DnsNames:   pulumi.ToStringArray(locals.IngressHostnames),
				SecretName: pulumi.String(locals.IngressCertSecretName),
				IssuerRef: certmanagerv1.CertificateSpecIssuerRefArgs{
					Kind: pulumi.String("ClusterIssuer"),
					Name: pulumi.String(locals.IngressCertClusterIssuerName),
				},
			},
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "error creating certificate")
	}

	//create gateway for ingress from external(outside vpc) clients
	createdExternalGateway, err := gatewayv1.NewGateway(ctx,
		"external",
		&gatewayv1.GatewayArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name: pulumi.Sprintf("%s-external", locals.Namespace),
				//all gateway resources should be created in the istio-ingress deployment namespace
				Namespace: pulumi.String(vars.IstioIngressNamespace),
				Labels:    pulumi.ToStringMap(locals.Labels),
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
					&gatewayv1.GatewaySpecListenersArgs{
						Name:     pulumi.String("https-external"),
						Hostname: pulumi.String(locals.IngressExternalHostname),
						Port:     pulumi.Int(443),
						Protocol: pulumi.String("HTTPS"),
						Tls: &gatewayv1.GatewaySpecListenersTlsArgs{
							Mode: pulumi.String("Terminate"),
							CertificateRefs: gatewayv1.GatewaySpecListenersTlsCertificateRefsArray{
								gatewayv1.GatewaySpecListenersTlsCertificateRefsArgs{
									Name: pulumi.String(locals.IngressCertSecretName),
								},
							},
						},
						AllowedRoutes: gatewayv1.GatewaySpecListenersAllowedRoutesArgs{
							Namespaces: gatewayv1.GatewaySpecListenersAllowedRoutesNamespacesArgs{
								From: pulumi.String("All"),
							},
						},
					},
					&gatewayv1.GatewaySpecListenersArgs{
						Name:     pulumi.String("http-external"),
						Hostname: pulumi.String(locals.IngressExternalHostname),
						Port:     pulumi.Int(80),
						Protocol: pulumi.String("HTTP"),
						AllowedRoutes: gatewayv1.GatewaySpecListenersAllowedRoutesArgs{
							Namespaces: gatewayv1.GatewaySpecListenersAllowedRoutesNamespacesArgs{
								From: pulumi.String("All"),
							},
						},
					},
				},
			},
		}, pulumi.Provider(kubernetesProvider), pulumi.DependsOn([]pulumi.Resource{addedCertificate}))
	if err != nil {
		return errors.Wrap(err, "error creating gateway for ingress from external clients")
	}

	//create gateway for ingress from external(outside vpc) clients
	createdInternalGateway, err := gatewayv1.NewGateway(ctx,
		"internal",
		&gatewayv1.GatewayArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name: pulumi.Sprintf("%s-internal", locals.Namespace),
				//all gateway resources should be created in the istio-ingress deployment namespace
				Namespace: pulumi.String(vars.IstioIngressNamespace),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: gatewayv1.GatewaySpecArgs{
				GatewayClassName: pulumi.String(vars.GatewayIngressClassName),
				Addresses: gatewayv1.GatewaySpecAddressesArray{
					gatewayv1.GatewaySpecAddressesArgs{
						Type:  pulumi.String("Hostname"),
						Value: pulumi.String(vars.GatewayInternalLoadBalancerServiceHostname),
					},
				},
				Listeners: gatewayv1.GatewaySpecListenersArray{
					&gatewayv1.GatewaySpecListenersArgs{
						Name:     pulumi.String("https-internal"),
						Hostname: pulumi.String(locals.IngressInternalHostname),
						Port:     pulumi.Int(443),
						Protocol: pulumi.String("HTTPS"),
						Tls: &gatewayv1.GatewaySpecListenersTlsArgs{
							Mode: pulumi.String("Terminate"),
							CertificateRefs: gatewayv1.GatewaySpecListenersTlsCertificateRefsArray{
								gatewayv1.GatewaySpecListenersTlsCertificateRefsArgs{
									Name: pulumi.String(locals.IngressCertSecretName),
								},
							},
						},
						AllowedRoutes: gatewayv1.GatewaySpecListenersAllowedRoutesArgs{
							Namespaces: gatewayv1.GatewaySpecListenersAllowedRoutesNamespacesArgs{
								From: pulumi.String("All"),
							},
						},
					},
					&gatewayv1.GatewaySpecListenersArgs{
						Name:     pulumi.String("http-internal"),
						Hostname: pulumi.String(locals.IngressInternalHostname),
						Port:     pulumi.Int(80),
						Protocol: pulumi.String("HTTP"),
						AllowedRoutes: gatewayv1.GatewaySpecListenersAllowedRoutesArgs{
							Namespaces: gatewayv1.GatewaySpecListenersAllowedRoutesNamespacesArgs{
								From: pulumi.String("All"),
							},
						},
					},
				},
			},
		}, pulumi.Provider(kubernetesProvider), pulumi.DependsOn([]pulumi.Resource{addedCertificate}))
	if err != nil {
		return errors.Wrap(err, "error creating gateway for ingress from external clients")
	}

	var destinationServicePort = pulumi.Int(80)
	for _, p := range locals.MicroserviceKubernetes.Spec.Container.App.Ports {
		if p.IsIngressPort {
			destinationServicePort = pulumi.Int(p.ServicePort)
		}
	}

	//create http-route for setting up https-redirect for external-hostname
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"http-external-redirect",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("http-external-redirect"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.IngressExternalHostname)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s-external", locals.Namespace),
						Namespace:   createdExternalGateway.Metadata.Namespace(),
						SectionName: pulumi.String("http-external"),
					},
				},
				Rules: gatewayv1.HTTPRouteSpecRulesArray{
					gatewayv1.HTTPRouteSpecRulesArgs{
						Filters: gatewayv1.HTTPRouteSpecRulesFiltersArray{
							gatewayv1.HTTPRouteSpecRulesFiltersArgs{
								RequestRedirect: gatewayv1.HTTPRouteSpecRulesFiltersRequestRedirectArgs{
									Scheme:     pulumi.String("https"),
									StatusCode: pulumi.Int(301),
								},
								Type: pulumi.String("RequestRedirect"),
							},
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace))

	//create http-route for external-hostname with https listener
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"https-external",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("https-external"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.IngressExternalHostname)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s-external", locals.Namespace),
						Namespace:   createdExternalGateway.Metadata.Namespace(),
						SectionName: pulumi.String("https-external"),
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
								Name:      pulumi.String(locals.KubeServiceName),
								Namespace: createdNamespace.Metadata.Name(),
								Port:      destinationServicePort,
							},
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace))

	//create http-route for setting up https-redirect for internal-hostname
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"http-internal-redirect",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("http-internal-redirect"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.IngressInternalHostname)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s-internal", locals.Namespace),
						Namespace:   createdInternalGateway.Metadata.Namespace(),
						SectionName: pulumi.String("http-external"),
					},
				},
				Rules: gatewayv1.HTTPRouteSpecRulesArray{
					gatewayv1.HTTPRouteSpecRulesArgs{
						Filters: gatewayv1.HTTPRouteSpecRulesFiltersArray{
							gatewayv1.HTTPRouteSpecRulesFiltersArgs{
								RequestRedirect: gatewayv1.HTTPRouteSpecRulesFiltersRequestRedirectArgs{
									Scheme:     pulumi.String("https"),
									StatusCode: pulumi.Int(301),
								},
								Type: pulumi.String("RequestRedirect"),
							},
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace))

	//create http-route for internal-hostname with https listener
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"https-internal",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("https-internal"),
				Namespace: createdNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.IngressInternalHostname)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s-internal", locals.Namespace),
						Namespace:   createdInternalGateway.Metadata.Namespace(),
						SectionName: pulumi.String("https-internal"),
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
								Name:      pulumi.String(locals.KubeServiceName),
								Namespace: createdNamespace.Metadata.Name(),
								Port:      destinationServicePort,
							},
						},
					},
				},
			},
		}, pulumi.Parent(createdNamespace))
	return nil
}
