package module

var vars = struct {
	Namespace                     string
	HelmChartName                 string
	HelmChartRepo                 string
	DefaultStableVersion          string
	KsaName                       string
	DefaultSecretsPollIntervalSec int
}{
	Namespace:            "external-secrets",
	HelmChartName:        "external-secrets",
	HelmChartRepo:        "https://charts.external-secrets.io",
	DefaultStableVersion: "v0.9.20", // bump when you move the stable channel
	KsaName:              "external-secrets",
	// excessive polling can become expensive on some providers
	DefaultSecretsPollIntervalSec: 10,
}
