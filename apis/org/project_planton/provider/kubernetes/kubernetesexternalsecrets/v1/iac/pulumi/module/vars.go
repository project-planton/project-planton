package module

var vars = struct {
	Namespace                     string
	HelmChartName                 string
	HelmChartRepo                 string
	DefaultStableVersion          string
	KsaName                       string
	DefaultSecretsPollIntervalSec int
}{
	Namespace:            "kubernetes-external-secrets",
	HelmChartName:        "kubernetes-external-secrets",
	HelmChartRepo:        "https://charts.kubernetes-external-secrets.io",
	DefaultStableVersion: "v0.9.20", // bump when you move the stable channel
	KsaName:              "kubernetes-external-secrets",
	// excessive polling can become expensive on some providers
	DefaultSecretsPollIntervalSec: 10,
}
