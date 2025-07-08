package module

import (
	"github.com/pkg/errors"
	ingressnginxkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/addon/ingressnginxkubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates all Pulumi resources for the Ingress‑Nginx Kubernetes add‑on.
func Resources(ctx *pulumi.Context,
	stackInput *ingressnginxkubernetesv1.IngressNginxKubernetesStackInput) error {

	// Build Kubernetes provider from the target‑cluster credential
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(
		ctx, stackInput.ProviderCredential, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	spec := stackInput.Target.Spec

	// Determine chart version – explicit > default
	chartVersion := spec.ChartVersion
	if chartVersion == "" {
		chartVersion = vars.DefaultChartVersion
	}

	// Build service annotations based on cluster flavour + internal/external flag
	var serviceAnnotations map[string]string
	if gke := spec.GetGke(); gke != nil {
		if spec.Internal {
			serviceAnnotations = map[string]string{
				"cloud.google.com/load-balancer-type": "internal",
			}
		} else {
			serviceAnnotations = map[string]string{
				"cloud.google.com/load-balancer-type": "external",
			}
		}
	} else if eks := spec.GetEks(); eks != nil {
		if spec.Internal {
			serviceAnnotations = map[string]string{
				"service.beta.kubernetes.io/aws-load-balancer-scheme": "internal",
			}
		} else {
			serviceAnnotations = map[string]string{
				"service.beta.kubernetes.io/aws-load-balancer-scheme": "internet-facing",
			}
		}
		_ = eks // mute unused‑var if eks has no fields accessed here
	} else if aks := spec.GetAks(); aks != nil {
		if spec.Internal {
			serviceAnnotations = map[string]string{
				"service.beta.kubernetes.io/azure-load-balancer-internal": "true",
			}
		}
		_ = aks
	}

	// ---------------------------------------------------------------------
	// Namespace
	// ---------------------------------------------------------------------
	ns, err := corev1.NewNamespace(ctx, vars.Namespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(vars.Namespace),
			},
		},
		pulumi.Provider(kubeProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// ---------------------------------------------------------------------
	// Helm chart
	// ---------------------------------------------------------------------
	// Build values dynamically so that annotations are only set when needed
	serviceVals := pulumi.Map{
		"type": pulumi.String("LoadBalancer"),
	}
	if len(serviceAnnotations) > 0 {
		serviceVals["annotations"] = pulumi.ToStringMap(serviceAnnotations)
	}

	values := pulumi.Map{
		"controller": pulumi.Map{
			"service":                  serviceVals,
			"ingressClassResource":     pulumi.Map{"default": pulumi.Bool(true)},
			"watchIngressWithoutClass": pulumi.Bool(true),
		},
	}

	_, err = helm.NewRelease(ctx, vars.HelmChartName,
		&helm.ReleaseArgs{
			Name:            pulumi.String(vars.HelmChartName), // release name
			Namespace:       ns.Metadata.Name(),
			Chart:           pulumi.String(vars.HelmChartName),
			Version:         pulumi.String(chartVersion),
			CreateNamespace: pulumi.Bool(false),
			Atomic:          pulumi.Bool(true),
			CleanupOnFail:   pulumi.Bool(true),
			WaitForJobs:     pulumi.Bool(true),
			Timeout:         pulumi.Int(180),
			Values:          values,
			RepositoryOpts: helm.RepositoryOptsArgs{
				Repo: pulumi.String(vars.HelmChartRepo),
			},
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns))
	if err != nil {
		return errors.Wrap(err, "failed to install ingress-nginx helm release")
	}

	// ---------------------------------------------------------------------
	// Stack outputs
	// ---------------------------------------------------------------------
	ctx.Export(OpNamespace, ns.Metadata.Name())
	ctx.Export(OpReleaseName, pulumi.String(vars.HelmChartName))
	ctx.Export(OpServiceName,
		pulumi.Sprintf("%s-controller", vars.HelmChartName)) // helm default
	ctx.Export(OpServiceType, pulumi.String("LoadBalancer"))

	return nil
}
