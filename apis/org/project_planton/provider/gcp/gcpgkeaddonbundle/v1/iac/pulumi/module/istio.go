package module

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp/gcpgkeaddonbundle/v1/iac/pulumi/module/vars"
	istiov1alpha3 "github.com/project-planton/project-planton/pkg/kubernetes/kubernetestypes/istio/kubernetes/networking/v1alpha3"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp"
	"github.com/pulumi/pulumi-gcp/sdk/v8/go/gcp/compute"
	pulumikubernetes "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// istio installs the Istio service mesh in the Kubernetes cluster using Helm. It creates the necessary namespaces,
// installs the Helm charts for Istio base, Istiod, and gateway components, and sets up load balancers for ingress.
//
// Parameters:
// - ctx: The Pulumi context used for defining cloud resources.
// - locals: A struct containing local configuration and metadata.
// - createdCluster: The GKE cluster where Istio will be installed.
// - gcpProvider: The GCP provider for Pulumi.
// - kubernetesProvider: The Kubernetes provider for Pulumi.
//
// Returns:
// - error: An error object if there is any issue during the installation.
//
// The function performs the following steps:
// 1. Creates the `istio-system` namespace and labels it with metadata from locals.
// 2. Deploys the Istio base Helm chart into the `istio-system` namespace.
// 3. Deploys the Istiod Helm chart into the `istio-system` namespace with specific mesh configuration.
// 4. Creates the Istio gateway namespace and labels it with metadata from locals.
// 5. Deploys the Istio gateway Helm chart into the gateway namespace, configuring service ports for HTTP, HTTPS, and other protocols.
// 6. Creates a compute IP address for the internal load balancer and exports its address.
// 7. Creates a Kubernetes service for the internal load balancer using the created IP address and service port configurations.
// 8. Creates a compute IP address for the external load balancer and exports its address.
// 9. Creates a Kubernetes service for the external load balancer using the created IP address and service port configurations.
// 10. Handles errors and returns any errors encountered during the namespace creation, Helm release deployment, or service setup.
func istio(ctx *pulumi.Context, locals *Locals,
	gcpProvider *gcp.Provider,
	kubernetesProvider *pulumikubernetes.Provider) error {
	//create istio-system namespace resource
	createdIstioSystemNamespace, err := corev1.NewNamespace(ctx,
		vars.Istio.SystemNamespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.Istio.SystemNamespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create istio-system namespace")
	}

	//create istio-base helm-release
	_, err = helm.NewRelease(ctx, "istio-base",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.Istio.BaseHelmChartName),
			Namespace:       createdIstioSystemNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.Istio.BaseHelmChartName),
			Version:         pulumi.String(vars.Istio.HelmChartsVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values:          pulumi.Map{},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.Istio.HelmChartsRepo),
			},
		}, pulumi.Parent(createdIstioSystemNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create istio-base helm release")
	}

	//create istiod helm-release
	_, err = helm.NewRelease(ctx, "istiod",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.Istio.IstiodHelmChartName),
			Namespace:       createdIstioSystemNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.Istio.IstiodHelmChartName),
			Version:         pulumi.String(vars.Istio.HelmChartsVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"meshConfig": pulumi.StringMap{
					"ingressClass":          pulumi.String("istio"),
					"ingressControllerMode": pulumi.String("STRICT"),
					"ingressService":        pulumi.String("ingress-external"),
					"ingressSelector":       pulumi.String("ingress"),
				},
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.Istio.HelmChartsRepo),
			},
		}, pulumi.Parent(createdIstioSystemNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create istiod helm release")
	}

	//create istio-gateway namespace resource
	createdIstioGatewayNamespace, err := corev1.NewNamespace(ctx,
		vars.Istio.GatewayNamespace,
		&corev1.NamespaceArgs{
			Metadata: metav1.ObjectMetaPtrInput(
				&metav1.ObjectMetaArgs{
					Name:   pulumi.String(vars.Istio.GatewayNamespace),
					Labels: pulumi.ToStringMap(locals.KubernetesLabels),
				}),
		},
		pulumi.Provider(kubernetesProvider))
	if err != nil {
		return errors.Wrapf(err, "failed to create istio-system namespace")
	}

	//create istio-gateway helm-release
	createdIstioGatewayHelmRelease, err := helm.NewRelease(ctx,
		"istio-gateway",
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.Istio.GatewayHelmChartName),
			Namespace:       createdIstioGatewayNamespace.Metadata.Name(),
			Chart:           pulumi.String(vars.Istio.GatewayHelmChartName),
			Version:         pulumi.String(vars.Istio.HelmChartsVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(false),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values: pulumi.Map{
				"service": pulumi.Map{
					"type": pulumi.String("ClusterIP"),
					"ports": pulumi.MapArray{
						pulumi.Map{
							"name":       pulumi.String("status-port"),
							"protocol":   pulumi.String("TCP"),
							"port":       pulumi.Int(15021),
							"targetPort": pulumi.Int(15021),
						},
						pulumi.Map{
							"name":       pulumi.String("http2"),
							"protocol":   pulumi.String("TCP"),
							"port":       pulumi.Int(80),
							"targetPort": pulumi.Int(80),
						},
						pulumi.Map{
							"name":       pulumi.String("https"),
							"protocol":   pulumi.String("TCP"),
							"port":       pulumi.Int(443),
							"targetPort": pulumi.Int(443),
						},
						pulumi.Map{
							"name":       pulumi.String("debug"),
							"protocol":   pulumi.String("TCP"),
							"port":       pulumi.Int(5005),
							"targetPort": pulumi.Int(5005),
						},
					},
				},
			},
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.Istio.HelmChartsRepo),
			},
		}, pulumi.Parent(createdIstioGatewayNamespace),
		pulumi.IgnoreChanges([]string{"status", "description", "resourceNames"}))
	if err != nil {
		return errors.Wrap(err, "failed to create istio-gateway helm release")
	}

	//create grpc-web envoy filter to support grpc-web clients
	_, err = istiov1alpha3.NewEnvoyFilter(ctx,
		"grpc-web",
		&istiov1alpha3.EnvoyFilterArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:      pulumi.String("grpc-web"),
				Namespace: createdIstioGatewayNamespace.Metadata.Name(),
				Labels:    pulumi.ToStringMap(locals.KubernetesLabels),
			},
			Spec: istiov1alpha3.EnvoyFilterSpecArgs{
				ConfigPatches: istiov1alpha3.EnvoyFilterSpecConfigPatchesArray{
					istiov1alpha3.EnvoyFilterSpecConfigPatchesArgs{
						ApplyTo: pulumi.String("HTTP_FILTER"),
						Match: istiov1alpha3.EnvoyFilterSpecConfigPatchesMatchArgs{
							Context: pulumi.String("GATEWAY"),
							Listener: istiov1alpha3.EnvoyFilterSpecConfigPatchesMatchListenerArgs{
								FilterChain: istiov1alpha3.EnvoyFilterSpecConfigPatchesMatchListenerFilterChainArgs{
									Filter: istiov1alpha3.EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterArgs{
										Name: pulumi.String("envoy.filters.network.http_connection_manager"),
										SubFilter: istiov1alpha3.EnvoyFilterSpecConfigPatchesMatchListenerFilterChainFilterSubFilterArgs{
											Name: pulumi.String("envoy.filters.http.cors"),
										},
									},
								},
							},
						},
						Patch: istiov1alpha3.EnvoyFilterSpecConfigPatchesPatchArgs{
							Operation: pulumi.String("INSERT_BEFORE"),
							Value: pulumi.Map{
								"name": pulumi.String("envoy.filters.http.grpc_web"),
								"typed_config": pulumi.Map{
									"@type": pulumi.String("type.googleapis.com/envoy.extensions.filters.http.grpc_web.v3.GrpcWeb"),
								},
							},
						},
					},
				},
				WorkloadSelector: &istiov1alpha3.EnvoyFilterSpecWorkloadSelectorArgs{
					Labels: pulumi.ToStringMap(vars.Istio.SelectorLabels),
				},
			},
		}, pulumi.Parent(createdIstioGatewayNamespace),
		pulumi.DependsOn([]pulumi.Resource{createdIstioGatewayHelmRelease}))

	//define array of ports to be configured for both internal and external ingress services
	loadBalancerServicePortArray := corev1.ServicePortArray{
		&corev1.ServicePortArgs{
			Name:       pulumi.String("status-port"),
			Protocol:   pulumi.String("TCP"),
			Port:       pulumi.Int(vars.Istio.IstiodStatusPort),
			TargetPort: pulumi.Int(vars.Istio.IstiodStatusPort),
		},
		&corev1.ServicePortArgs{
			Name:       pulumi.String("http2"),
			Protocol:   pulumi.String("TCP"),
			Port:       pulumi.Int(vars.Istio.HttpPort),
			TargetPort: pulumi.Int(vars.Istio.HttpPort),
		},
		&corev1.ServicePortArgs{
			Name:       pulumi.String("https"),
			Protocol:   pulumi.String("TCP"),
			Port:       pulumi.Int(vars.Istio.HttpsPort),
			TargetPort: pulumi.Int(vars.Istio.HttpsPort),
		},
	}

	//create compute ip address for internal load-balancer
	createdIngressInternalLoadBalancerIp, err := compute.NewAddress(ctx,
		vars.Istio.IngressInternalLoadBalancerServiceName,
		&compute.AddressArgs{
			Name:        pulumi.Sprintf("gke-%s-ingress-internal", locals.GcpGkeAddonBundle.Metadata.Name),
			Project:     pulumi.String(locals.GcpGkeAddonBundle.Spec.ClusterProjectId),
			Region:      pulumi.String(locals.GcpGkeAddonBundle.Spec.Istio.ClusterRegion),
			AddressType: pulumi.String("INTERNAL"),
			Labels:      pulumi.ToStringMap(locals.GcpLabels),
			Subnetwork:  pulumi.String(locals.GcpGkeAddonBundle.Spec.Istio.SubNetworkSelfLink),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create ip address for ingress-internal load-balancer")
	}

	//export ingress-internal ip
	ctx.Export(OpIngressInternalIp, createdIngressInternalLoadBalancerIp.Address)

	//create load-balancer service for internal load-balancer
	_, err = corev1.NewService(ctx,
		vars.Istio.IngressInternalLoadBalancerServiceName,
		&corev1.ServiceArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:        pulumi.String(vars.Istio.IngressInternalLoadBalancerServiceName),
				Namespace:   createdIstioGatewayNamespace.Metadata.Name(),
				Annotations: pulumi.ToStringMap(vars.Istio.IngressInternalServiceAnnotations),
				Labels:      pulumi.ToStringMap(vars.Istio.SelectorLabels),
			},
			Spec: &corev1.ServiceSpecArgs{
				Type:           pulumi.String("LoadBalancer"),
				Selector:       pulumi.ToStringMap(vars.Istio.SelectorLabels),
				LoadBalancerIP: createdIngressInternalLoadBalancerIp.Address,
				Ports:          loadBalancerServicePortArray,
			},
		}, pulumi.Parent(createdIstioGatewayNamespace))
	if err != nil {
		return errors.Wrapf(err, "failed to create ingress-external kubernetes service")
	}

	//create compute ip address for external load-balancer
	createdIngressExternalLoadBalancerIp, err := compute.NewAddress(ctx,
		vars.Istio.IngressExternalLoadBalancerServiceName,
		&compute.AddressArgs{
			Name:        pulumi.Sprintf("gke-%s-ingress-external", locals.GcpGkeAddonBundle.Metadata.Name),
			Project:     pulumi.String(locals.GcpGkeAddonBundle.Spec.ClusterProjectId),
			Region:      pulumi.String(locals.GcpGkeAddonBundle.Spec.Istio.ClusterRegion),
			AddressType: pulumi.String("EXTERNAL"),
			Labels:      pulumi.ToStringMap(locals.GcpLabels),
		}, pulumi.Provider(gcpProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create ip address for ingress-internal load-balancer")
	}

	//export ingress-external ip
	ctx.Export(OpIngressExternalIp, createdIngressExternalLoadBalancerIp.Address)

	//create load-balancer service for external load-balancer
	_, err = corev1.NewService(ctx,
		vars.Istio.IngressExternalLoadBalancerServiceName,
		&corev1.ServiceArgs{
			Metadata: metav1.ObjectMetaArgs{
				Name:        pulumi.String(vars.Istio.IngressExternalLoadBalancerServiceName),
				Namespace:   createdIstioGatewayNamespace.Metadata.Name(),
				Annotations: pulumi.ToStringMap(vars.Istio.IngressExternalServiceAnnotations),
				Labels:      pulumi.ToStringMap(vars.Istio.SelectorLabels),
			},
			Spec: &corev1.ServiceSpecArgs{
				Type:           pulumi.String("LoadBalancer"),
				Selector:       pulumi.ToStringMap(vars.Istio.SelectorLabels),
				LoadBalancerIP: createdIngressExternalLoadBalancerIp.Address,
				Ports:          loadBalancerServicePortArray,
			},
		}, pulumi.Parent(createdIstioGatewayNamespace))
	if err != nil {
		return errors.Wrapf(err, "failed to create ingress-external kubernetes service")
	}

	return nil
}
