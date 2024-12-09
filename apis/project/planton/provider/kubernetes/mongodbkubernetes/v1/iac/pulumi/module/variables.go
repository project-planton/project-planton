package module

var vars = struct {
	MongodbRootPasswordKey string
	RootUsername           string
	MongoDbPort            int
	HelmChartName          string
	HelmChartRepoUrl       string
	HelmChartVersion       string
}{
	HelmChartName:          "mongodb",
	HelmChartRepoUrl:       "https://charts.bitnami.com/bitnami",
	HelmChartVersion:       "15.1.4",
	MongodbRootPasswordKey: "mongodb-root-password",
	RootUsername:           "root",
	MongoDbPort:            27017,
}
