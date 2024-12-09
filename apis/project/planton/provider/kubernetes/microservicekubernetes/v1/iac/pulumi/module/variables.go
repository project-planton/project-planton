package module

var vars = struct {
	GatewayIngressClassName                    string
	GatewayInternalLoadBalancerServiceHostname string
	GatewayExternalLoadBalancerServiceHostname string
	IstioIngressNamespace                      string
	MainPyConfigMapName                        string
	LibFilesConfigMapName                      string
	GcpSecretsManagerClusterSecretStoreName    string
}{
	GatewayIngressClassName:                    "istio",
	GatewayInternalLoadBalancerServiceHostname: "ingress-internal.istio-ingress.svc.cluster.local",
	GatewayExternalLoadBalancerServiceHostname: "ingress-external.istio-ingress.svc.cluster.local",
	IstioIngressNamespace:                      "istio-ingress",
	MainPyConfigMapName:                        "main-py",
	LibFilesConfigMapName:                      "lib-files",
	GcpSecretsManagerClusterSecretStoreName:    "gcp-secrets-manager",
}
