package module

var vars = struct {
	// Helm
	HelmChartName    string
	HelmChartRepoUrl string
	HelmChartVersion string
	// Secret for external DB password
	DatabasePasswordSecretKey  string
	DatabasePasswordSecretName string
	// Service ports
	FrontendGrpcPort int
	FrontendHttpPort int
	UIPort           int
	// Istio / Gateway-API constants for UI ingress
	IstioIngressNamespace                      string
	GatewayIngressClassName                    string
	GatewayExternalLoadBalancerServiceHostname string
	// Temporal defaults and common Helm keys
	DefaultCassandraReplicas     int
	HelmKeyPrometheusEnabled     string
	HelmKeyGrafanaEnabled        string
	HelmKeyElasticsearchEnabled  string
	HelmKeyCassandraReplicaCount string
}{
	HelmChartName:    "temporal",
	HelmChartRepoUrl: "https://go.temporal.io/helm-charts",
	HelmChartVersion: "0.62.0",

	DatabasePasswordSecretKey:  "password",
	DatabasePasswordSecretName: "temporal-db-password",

	FrontendGrpcPort: 7233,
	FrontendHttpPort: 7243,
	UIPort:           8080,

	IstioIngressNamespace:                      "istio-ingress",
	GatewayIngressClassName:                    "istio",
	GatewayExternalLoadBalancerServiceHostname: "ingress-external.istio-ingress.svc.cluster.local",

	DefaultCassandraReplicas: 1,

	HelmKeyPrometheusEnabled:     "prometheus.enabled",
	HelmKeyGrafanaEnabled:        "grafana.enabled",
	HelmKeyElasticsearchEnabled:  "elasticsearch.enabled",
	HelmKeyCassandraReplicaCount: "cassandra.replicaCount",
}
