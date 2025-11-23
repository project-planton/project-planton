package module

import (
	"fmt"

	kubernetesexternaldnsv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes/kubernetesexternaldns/v1"
)

// getGkeServiceAccountEmail constructs the Google Service Account email for Workload Identity
// Format: <service-account-name>@<project-id>.iam.gserviceaccount.com
func getGkeServiceAccountEmail(serviceAccountName, projectId string) string {
	return fmt.Sprintf("%s@%s.iam.gserviceaccount.com", serviceAccountName, projectId)
}

// getCloudflareSecretName constructs the secret name for Cloudflare API token
// Format: cloudflare-api-token-<release-name>
func getCloudflareSecretName(releaseName string) string {
	return fmt.Sprintf("cloudflare-api-token-%s", releaseName)
}

// getProviderType returns a string representation of the configured DNS provider
func getProviderType(spec *kubernetesexternaldnsv1.KubernetesExternalDnsSpec) string {
	switch {
	case spec.GetGke() != nil:
		return "gke"
	case spec.GetEks() != nil:
		return "eks"
	case spec.GetAks() != nil:
		return "aks"
	case spec.GetCloudflare() != nil:
		return "cloudflare"
	default:
		return "unknown"
	}
}

// getNamespace returns the namespace from spec, with default if not specified
func getNamespace(spec *kubernetesexternaldnsv1.KubernetesExternalDnsSpec) string {
	namespace := spec.Namespace.GetValue()
	if namespace == "" {
		return "kubernetes-external-dns" // default
	}
	return namespace
}

// getHelmChartVersion returns the Helm chart version from spec, with default if not specified
func getHelmChartVersion(spec *kubernetesexternaldnsv1.KubernetesExternalDnsSpec) string {
	if spec.HelmChartVersion != nil && *spec.HelmChartVersion != "" {
		return *spec.HelmChartVersion
	}
	return "1.19.0" // default
}

// getExternalDnsVersion returns the ExternalDNS version from spec, with default if not specified
func getExternalDnsVersion(spec *kubernetesexternaldnsv1.KubernetesExternalDnsSpec) string {
	if spec.KubernetesExternalDnsVersion != nil && *spec.KubernetesExternalDnsVersion != "" {
		return *spec.KubernetesExternalDnsVersion
	}
	return "v0.19.0" // default
}
