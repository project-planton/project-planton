package module

var vars = struct {
	DefaultNamespace string
	HelmChartName    string
	HelmChartRepo    string
	HelmChartVersion string
}{
	DefaultNamespace: "percona-operator",
	HelmChartName:    "psmdb-operator",
	HelmChartRepo:    "https://percona.github.io/percona-helm-charts/",
	HelmChartVersion: "1.16.0",
}
