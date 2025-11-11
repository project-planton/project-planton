package module

var vars = struct {
	GatewayIngressClassName                    string
	GatewayInternalLoadBalancerServiceHostname string
	GatewayExternalLoadBalancerServiceHostname string
	IstioIngressNamespace                      string
	GcpSecretsManagerClusterSecretStoreName    string
}{
	GatewayIngressClassName:                    "istio",
	GatewayInternalLoadBalancerServiceHostname: "ingress-internal.istio-ingress.svc.cluster.local",
	GatewayExternalLoadBalancerServiceHostname: "ingress-external.istio-ingress.svc.cluster.local",
	IstioIngressNamespace:                      "istio-ingress",
	GcpSecretsManagerClusterSecretStoreName:    "gcp-secrets-manager",
}
