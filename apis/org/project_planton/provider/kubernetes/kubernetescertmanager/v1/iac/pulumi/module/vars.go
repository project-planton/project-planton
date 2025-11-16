package module

var vars = struct {
	Namespace     string
	HelmChartName string
	HelmChartRepo string
	KsaName       string
}{
	Namespace:     "kubernetes-cert-manager",
	HelmChartName: "kubernetes-cert-manager",
	HelmChartRepo: "https://charts.jetstack.io",
	KsaName:       "kubernetes-cert-manager",
}
