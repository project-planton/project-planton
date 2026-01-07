package module

// Helm chart constants for GHA Runner Scale Set.
var (
	// HelmRepoURL is the OCI repository for the ARC Helm charts.
	HelmRepoURL = "oci://ghcr.io/actions/actions-runner-controller-charts"

	// HelmChartName is the name of the runner scale set chart.
	HelmChartName = "gha-runner-scale-set"

	// DefaultChartVersion is the default version of the Helm chart.
	DefaultChartVersion = "0.13.1"
)

// Output keys for stack exports.
const (
	OpNamespace          = "namespace"
	OpReleaseName        = "release_name"
	OpChartVersion       = "chart_version"
	OpRunnerScaleSetName = "runner_scale_set_name"
	OpGitHubConfigURL    = "github_config_url"
	OpGitHubSecretName   = "github_secret_name"
	OpPvcNames           = "pvc_names"
	OpMinRunners         = "min_runners"
	OpMaxRunners         = "max_runners"
	OpContainerMode      = "container_mode"
)
