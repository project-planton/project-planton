package module

var vars = struct {
	ClickhousePasswordKey string
	DefaultUsername       string
	ClickhouseHttpPort    int
	ClickhouseNativePort  int
	HelmChartName         string
	HelmChartRepoUrl      string
	HelmChartVersion      string
}{
	HelmChartName:         "clickhouse",
	HelmChartRepoUrl:      "https://charts.bitnami.com/bitnami",
	HelmChartVersion:      "6.2.15",
	ClickhousePasswordKey: "admin-password",
	DefaultUsername:       "default",
	ClickhouseHttpPort:    8123,
	ClickhouseNativePort:  9000,
}
