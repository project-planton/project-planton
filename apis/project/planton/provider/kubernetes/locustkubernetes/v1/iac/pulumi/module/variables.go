package module

var vars = struct {
	GatewayExternalLoadBalancerServiceHostname string
	GatewayIngressClassName                    string
	IstioIngressNamespace                      string
	LibFilesConfigMapName                      string
	MainPyConfigMapName                        string
}{
	GatewayExternalLoadBalancerServiceHostname: "ingress-external.istio-ingress.svc.cluster.local",
	GatewayIngressClassName:                    "istio",
	IstioIngressNamespace:                      "istio-ingress",
	LibFilesConfigMapName:                      "lib-files",
	MainPyConfigMapName:                        "main-py",
}
