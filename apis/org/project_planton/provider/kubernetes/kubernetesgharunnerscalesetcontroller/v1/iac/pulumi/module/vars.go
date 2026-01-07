package module

// vars groups all operator-level constants so version bumps or repo changes
// remain one-liner edits.
var vars = struct {
	// Helm chart configuration
	HelmRepoURL        string
	HelmChartName      string
	DefaultReleaseName string

	// Default Helm chart version (can be overridden via spec)
	DefaultChartVersion string
}{
	HelmRepoURL:         "oci://ghcr.io/actions/actions-runner-controller-charts",
	HelmChartName:       "gha-runner-scale-set-controller",
	DefaultReleaseName:  "arc",
	DefaultChartVersion: "0.13.1",
}
