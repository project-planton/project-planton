package module

import (
	"fmt"

	"github.com/pkg/errors"
	externaldnsv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/addon/externaldnskubernetes/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/kubernetes/pulumikubernetesprovider"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v4/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Resources creates all Pulumi resources for the ExternalDNS Kubernetes add‑on.
func Resources(ctx *pulumi.Context, stackInput *externaldnsv1.ExternalDnsKubernetesStackInput) error {
	// Set up the Kubernetes provider from the supplied cluster credential.
	kubeProvider, err := pulumikubernetesprovider.GetWithKubernetesClusterCredential(
		ctx, stackInput.ProviderCredential, "kubernetes")
	if err != nil {
		return errors.Wrap(err, "failed to set up kubernetes provider")
	}

	spec := stackInput.Target.Spec

	// Namespace and chart version now come from spec with defaults in proto
	namespace := spec.Namespace
	chartVersion := spec.HelmChartVersion

	// Use the resource name as the Helm release name (e.g., "external-dns-planton-cloud")
	// This allows multiple ExternalDNS instances for different domains in the same namespace
	releaseName := stackInput.Target.Metadata.Name
	ksaName := releaseName // ServiceAccount name matches the release name

	// Create the namespace.
	ns, err := corev1.NewNamespace(ctx, namespace,
		&corev1.NamespaceArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name: pulumi.String(namespace),
			},
		},
		pulumi.Provider(kubeProvider))
	if err != nil {
		return errors.Wrap(err, "failed to create namespace")
	}

	// Build ServiceAccount annotations and Helm values according to provider.
	annotations := pulumi.StringMap{}
	values := pulumi.Map{
		"serviceAccount": pulumi.Map{
			"create": pulumi.Bool(false),
			"name":   pulumi.String(ksaName),
		},
	}

	switch {
	case spec.GetGke() != nil:
		gke := spec.GetGke()
		values["provider"] = pulumi.String("google")
		values["google"] = pulumi.Map{
			"project": pulumi.String(gke.ProjectId.GetValue()),
		}
		// Zone‑filter keeps ExternalDNS scoped to the desired zone.
		values["zoneIdFilters"] = pulumi.StringArray{
			pulumi.String(gke.DnsZoneId.GetValue()),
		}
		// Best‑effort GSA e‑mail derivation; users can override by patching the SA.
		gsaEmail := fmt.Sprintf("%s@%s.iam.gserviceaccount.com", ksaName, gke.ProjectId.GetValue())
		annotations["iam.gke.io/gcp-service-account"] = pulumi.String(gsaEmail)

	case spec.GetEks() != nil:
		eks := spec.GetEks()
		values["provider"] = pulumi.String("aws")
		values["zoneIdFilters"] = pulumi.StringArray{
			pulumi.String(eks.Route53ZoneId.GetValue()),
		}
		if eks.IrsaRoleArnOverride != "" {
			annotations["eks.amazonaws.com/role-arn"] = pulumi.String(eks.IrsaRoleArnOverride)
		}

	case spec.GetAks() != nil:
		aks := spec.GetAks()
		values["provider"] = pulumi.String("azure")
		if aks.DnsZoneId != "" {
			values["domainFilters"] = pulumi.StringArray{pulumi.String(aks.DnsZoneId)}
		}
		if aks.ManagedIdentityClientId != "" {
			annotations["azure.workload.identity/client-id"] =
				pulumi.String(aks.ManagedIdentityClientId)
		}

	case spec.GetCloudflare() != nil:
		cf := spec.GetCloudflare()

		// Create secret for Cloudflare API token with domain-specific name
		// Note: The Helm chart expects the key to be "apiKey" even for tokens
		secretName := fmt.Sprintf("cloudflare-api-token-%s", releaseName)
		secret, err := corev1.NewSecret(ctx, secretName,
			&corev1.SecretArgs{
				Metadata: &metav1.ObjectMetaArgs{
					Name:      pulumi.String(secretName),
					Namespace: ns.Metadata.Name(),
				},
				StringData: pulumi.StringMap{
					"apiKey": pulumi.String(cf.ApiToken),
				},
			},
			pulumi.Provider(kubeProvider),
			pulumi.Parent(ns))
		if err != nil {
			return errors.Wrap(err, "failed to create cloudflare api token secret")
		}

		// Configure provider
		values["provider"] = pulumi.String("cloudflare")

		// Mount the API token secret as environment variable
		values["env"] = pulumi.Array{
			pulumi.Map{
				"name": pulumi.String("CF_API_TOKEN"),
				"valueFrom": pulumi.Map{
					"secretKeyRef": pulumi.Map{
						"name": secret.Metadata.Name(),
						"key":  pulumi.String("apiKey"),
					},
				},
			},
		}

		// Configure extra args for Cloudflare-specific features
		extraArgs := pulumi.StringArray{
			pulumi.String("--cloudflare-dns-records-per-page=5000"),
			pulumi.String(fmt.Sprintf("--zone-id-filter=%s", cf.DnsZoneId)),
		}

		// Add proxy flag if enabled
		if cf.IsProxied {
			extraArgs = append(extraArgs, pulumi.String("--cloudflare-proxied"))
		}

		values["extraArgs"] = extraArgs

	default:
		return errors.New("spec.provider_config must be set (gke, eks, aks, or cloudflare)")
	}

	// Honor an optional custom ExternalDNS version.
	if spec.ExternalDnsVersion != "" {
		values["image"] = pulumi.Map{
			"tag": pulumi.String(spec.ExternalDnsVersion),
		}
	}

	// Create the ServiceAccount.
	_, err = corev1.NewServiceAccount(ctx, ksaName,
		&corev1.ServiceAccountArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:        pulumi.String(ksaName),
				Namespace:   ns.Metadata.Name(),
				Annotations: annotations,
			},
		},
		pulumi.Provider(kubeProvider),
		pulumi.Parent(ns))
	if err != nil {
		return errors.Wrap(err, "failed to create service account")
	}

	// Deploy the Helm release.
	_, err = helm.NewRelease(ctx, releaseName,
		&helm.ReleaseArgs{
			Name:            pulumi.String(releaseName),
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
		return errors.Wrap(err, "failed to install external-dns helm release")
	}

	// Export stack outputs.
	ctx.Export(OpNamespace, ns.Metadata.Name())
	ctx.Export(OpReleaseName, pulumi.String(releaseName))
	ctx.Export(OpSolverSa, pulumi.String(ksaName))

	return nil
}
