package module

var vars = struct {
	GatewayExternalLoadBalancerServiceHostname string
	GatewayIngressClassName                    string
	HelmChartName                              string
	HelmChartRepoUrl                           string
	HelmChartVersion                           string
	IstioIngressNamespace                      string
}{
	GatewayExternalLoadBalancerServiceHostname: "ingress-external.istio-ingress.svc.cluster.local",
	GatewayIngressClassName:                    "istio",
	HelmChartName:                              "openfga",
	HelmChartRepoUrl:                           "https://openfga.github.io/helm-charts",
	HelmChartVersion:                           "0.2.12",
	IstioIngressNamespace:                      "istio-ingress",
}
