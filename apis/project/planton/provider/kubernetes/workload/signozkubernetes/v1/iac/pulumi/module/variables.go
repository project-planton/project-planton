package module

var vars = struct {
	HelmChartName    string
	HelmChartRepoUrl string
	HelmChartVersion string
	SignozUIPort     int
	OtelGrpcPort     int
	OtelHttpPort     int
}{
	HelmChartName:    "signoz",
	HelmChartRepoUrl: "https://charts.signoz.io",
	HelmChartVersion: "0.52.0",
	SignozUIPort:     8080,
	OtelGrpcPort:     4317,
	OtelHttpPort:     4318,
}
