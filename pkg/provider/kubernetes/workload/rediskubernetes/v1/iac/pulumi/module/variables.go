package module

var vars = struct {
	HelmChartName           string
	HelmChartRepoUrl        string
	HelmChartVersion        string
	RedisPasswordSecretKey  string
	RedisPasswordSecretName string
	RedisPort               int
}{
	HelmChartName:           "redis",
	HelmChartRepoUrl:        "https://charts.bitnami.com/bitnami",
	HelmChartVersion:        "17.10.1",
	RedisPasswordSecretKey:  "password",
	RedisPasswordSecretName: "redis-password",
	RedisPort:               6379,
}
