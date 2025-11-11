package module

var vars = struct {
	Namespace              string
	HelmChartName          string
	HelmChartRepo          string
	DefaultStableVersion   string
	DefaultLatestVersion   string
	CrdManifestDownloadURL string
}{
	Namespace:     "solr-operator",
	HelmChartName: "solr-operator",
	HelmChartRepo: "https://solr.apache.org/charts",
	// keep these two in sync with upstream releases
	DefaultStableVersion:   "0.7.0",
	DefaultLatestVersion:   "0.8.1",
	CrdManifestDownloadURL: "https://solr.apache.org/operator/downloads/crds/v0.7.0/all-with-dependencies.yaml",
}
