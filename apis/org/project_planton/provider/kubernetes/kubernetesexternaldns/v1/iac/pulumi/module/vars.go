package module

// vars holds all tunables for this add‑on.
var vars = struct {
	HelmChartName string
	HelmChartRepo string
}{
	HelmChartName: "kubernetes-external-dns",
	HelmChartRepo: "https://kubernetes-sigs.github.io/kubernetes-external-dns/",
}
