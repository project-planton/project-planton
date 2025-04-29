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
	FrontendPort int
	UIPort       int
	// Istio / Gateway-API constants for UI ingress
	IstioIngressNamespace                      string
	GatewayIngressClassName                    string
	GatewayExternalLoadBalancerServiceHostname string
	// Certificate details
	IngressCertClusterIssuerName string
}{
	HelmChartName:    "temporal",
	HelmChartRepoUrl: "https://go.temporal.io/helm-charts",
	HelmChartVersion: "0.62.0",

	DatabasePasswordSecretKey:  "password",
	DatabasePasswordSecretName: "temporal-db-password",

	FrontendPort: 7233,
	UIPort:       8080,

	IstioIngressNamespace:                      "istio-ingress",
	GatewayIngressClassName:                    "istio",
	GatewayExternalLoadBalancerServiceHostname: "istio-ingress.example.com",

	IngressCertClusterIssuerName: "letsencrypt-prod",
}
