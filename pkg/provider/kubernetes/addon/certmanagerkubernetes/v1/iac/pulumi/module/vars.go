package module

var vars = struct {
	Namespace     string
	HelmChartName string
	HelmChartRepo string
	KsaName       string
}{
	Namespace:     "cert-manager",
	HelmChartName: "cert-manager",
	HelmChartRepo: "https://charts.jetstack.io",
	KsaName:       "cert-manager",
}
