package module

var vars = struct {
	DefaultNamespace string
	HelmChartName    string
	HelmChartRepo    string
	HelmChartVersion string
}{
	DefaultNamespace: "percona-postgresql-operator",
	HelmChartName:    "pg-operator",
	HelmChartRepo:    "https://percona.github.io/percona-helm-charts/",
	HelmChartVersion: "2.7.0",
}

