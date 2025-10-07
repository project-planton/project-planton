package module

var vars = struct {
	DefaultNamespace string
	HelmChartName    string
	HelmChartRepo    string
	HelmChartVersion string
}{
	DefaultNamespace: "altinity-operator",
	HelmChartName:    "altinity-clickhouse-operator",
	HelmChartRepo:    "https://docs.altinity.com/clickhouse-operator/",
	HelmChartVersion: "0.23.6", // Check for latest stable version
}
