package module

// vars holds all tunables for this add‑on.
var vars = struct {
	Namespace           string
	HelmChartName       string
	HelmChartRepo       string
	DefaultChartVersion string
	KsaName             string
}{
	Namespace:     "external-dns",
	HelmChartName: "external-dns",
	HelmChartRepo: "https://kubernetes-sigs.github.io/external-dns/",
	// Update when a newer stable version is promoted.
	DefaultChartVersion: "1.14.4",
	KsaName:             "external-dns",
}
