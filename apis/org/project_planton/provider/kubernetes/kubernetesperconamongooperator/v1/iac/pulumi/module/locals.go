package module

var locals = struct {
	HelmChartName    string
	HelmChartRepo    string
	HelmChartVersion string
}{
	HelmChartName:    "psmdb-operator",
	HelmChartRepo:    "https://percona.github.io/percona-helm-charts/",
	HelmChartVersion: "1.20.1",
}
