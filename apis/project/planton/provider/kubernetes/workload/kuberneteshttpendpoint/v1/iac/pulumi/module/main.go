package module

import (
	"github.com/pkg/errors"
	kuberneteshttpendpointv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/kuberneteshttpendpoint/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	certmanagerv1 "github.com/project-planton/project-planton/pkg/kubernetestypes/certmanager/kubernetes/cert_manager/v1"
	gatewayv1 "github.com/project-planton/project-planton/pkg/kubernetestypes/gatewayapis/kubernetes/gateway/v1"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func Resources(ctx *pulumi.Context, stackInput *kuberneteshttpendpointv1.KubernetesHttpEndpointStackInput) error {
	locals := initializeLocals(ctx, stackInput)

	//create kubernetes-provider from the credential in the stack-input
	kubernetesProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(ctx,
		stackInput.ProviderCredential, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to create kubernetes provider")
	}

	listenersArray := gatewayv1.GatewaySpecListenersArray{
		&gatewayv1.GatewaySpecListenersArgs{
			Name:     pulumi.String("http-external"),
			Hostname: pulumi.String(locals.EndpointDomainName),
			Port:     pulumi.Int(80),
			Protocol: pulumi.String("HTTP"),
			AllowedRoutes: gatewayv1.GatewaySpecListenersAllowedRoutesArgs{
				Namespaces: gatewayv1.GatewaySpecListenersAllowedRoutesNamespacesArgs{
					From: pulumi.String("All"),
				},
			},
		},
	}

	if locals.KubernetesHttpEndpoint.Spec.IsTlsEnabled {
		// Create new certificate
		_, err := certmanagerv1.NewCertificate(ctx,
			"ingress-certificate",
			&certmanagerv1.CertificateArgs{
				Metadata: metav1.ObjectMetaArgs{
					Name:      pulumi.String(locals.KubernetesHttpEndpoint.Metadata.Id),
					Namespace: pulumi.String(vars.IstioIngressNamespace),
					Labels:    pulumi.ToStringMap(locals.Labels),
				},
				Spec: certmanagerv1.CertificateSpecArgs{
					DnsNames:   pulumi.ToStringArray([]string{locals.KubernetesHttpEndpoint.Metadata.Name}),
					SecretName: pulumi.String(locals.IngressCertSecretName),
					IssuerRef: certmanagerv1.CertificateSpecIssuerRefArgs{
						Kind: pulumi.String("ClusterIssuer"),
						Name: pulumi.String(locals.KubernetesHttpEndpoint.Spec.CertClusterIssuerName),
					},
				},
			}, pulumi.Provider(kubernetesProvider))
		if err != nil {
			return errors.Wrap(err, "error creating certificate")
		}

		listenersArray = append(listenersArray, &gatewayv1.GatewaySpecListenersArgs{
			Name:     pulumi.String("https-external"),
			Hostname: pulumi.String(locals.EndpointDomainName),
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
		})
	}

	// Create external gateway
	createdGateway, err := gatewayv1.NewGateway(ctx,
		"gateway",
		&gatewayv1.GatewayArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name: pulumi.String(locals.KubernetesHttpEndpoint.Metadata.Id),
				// All gateway resources should be created in the ingress deployment namespace
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
				Listeners: listenersArray,
			},
		}, pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrap(err, "error creating gateway")
	}

	if locals.KubernetesHttpEndpoint.Spec.IsTlsEnabled {
		//create http-route for setting up https-redirect for external-hostname
		_, err = gatewayv1.NewHTTPRoute(ctx,
			"http-external-redirect",
			&gatewayv1.HTTPRouteArgs{
				Metadata: metav1.ObjectMetaArgs{
					Name:      pulumi.Sprintf("%s-http-external-redirect", locals.KubernetesHttpEndpoint.Metadata.Id),
					Namespace: pulumi.String(vars.IstioIngressNamespace),
					Labels:    pulumi.ToStringMap(locals.Labels),
				},
				Spec: gatewayv1.HTTPRouteSpecArgs{
					Hostnames: pulumi.StringArray{pulumi.String(locals.EndpointDomainName)},
					ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
						gatewayv1.HTTPRouteSpecParentRefsArgs{
							Name:        pulumi.String(locals.KubernetesHttpEndpoint.Metadata.Id),
							Namespace:   createdGateway.Metadata.Namespace(),
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
			}, pulumi.Provider(kubernetesProvider))
	}

	httpRulesArray := gatewayv1.HTTPRouteSpecRulesArray{}

	//build http-rules based on the routes configured in the input
	for _, routingRule := range locals.KubernetesHttpEndpoint.Spec.RoutingRules {
		httpRulesArray = append(httpRulesArray,
			gatewayv1.HTTPRouteSpecRulesArgs{
				BackendRefs: gatewayv1.HTTPRouteSpecRulesBackendRefsArray{
					gatewayv1.HTTPRouteSpecRulesBackendRefsArgs{
						Name:      pulumi.String(routingRule.BackendService.Name),
						Namespace: pulumi.String(routingRule.BackendService.Namespace),
						Port:      pulumi.Int(routingRule.BackendService.Port),
					},
				},
				Filters: nil,
				Matches: gatewayv1.HTTPRouteSpecRulesMatchesArray{
					gatewayv1.HTTPRouteSpecRulesMatchesArgs{
						Path: gatewayv1.HTTPRouteSpecRulesMatchesPathArgs{
							Type:  pulumi.String("PathPrefix"),
							Value: pulumi.String(routingRule.UrlPathPrefix),
						},
					},
				},
			})
	}

	// Create HTTP route with routing rules for http listener
	_, err = gatewayv1.NewHTTPRoute(ctx,
		"https",
		&gatewayv1.HTTPRouteArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String(locals.KubernetesHttpEndpoint.Metadata.Id),
				Namespace: pulumi.String(vars.IstioIngressNamespace),
				Labels:    pulumi.ToStringMap(locals.Labels),
			},
			Spec: gatewayv1.HTTPRouteSpecArgs{
				Hostnames: pulumi.StringArray{pulumi.String(locals.EndpointDomainName)},
				ParentRefs: gatewayv1.HTTPRouteSpecParentRefsArray{
					gatewayv1.HTTPRouteSpecParentRefsArgs{
						Name:        pulumi.Sprintf("%s", createdGateway.Metadata.Name()),
						Namespace:   createdGateway.Metadata.Namespace(),
						SectionName: pulumi.String("https-external"),
					},
				},
				Rules: httpRulesArray,
			},
		}, pulumi.Parent(createdGateway))

	if err != nil {
		return errors.Wrap(err, "error creating HTTP route")
	}
	return nil
}
