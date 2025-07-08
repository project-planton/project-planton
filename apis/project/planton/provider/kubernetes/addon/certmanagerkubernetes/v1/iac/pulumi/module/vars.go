package module

var vars = struct {
	Namespace            string
	HelmChartName        string
	HelmChartRepo        string
	KsaName              string
	DefaultStableVersion string
	DefaultLatestVersion string
}{
	Namespace:            "cert-manager",
	HelmChartName:        "cert-manager",
	HelmChartRepo:        "https://charts.jetstack.io",
	KsaName:              "cert-manager",
	DefaultStableVersion: "v1.15.2", // update when you move the stable channel
	DefaultLatestVersion: "v1.16.0", // update when Jetstack tags a new version
}
