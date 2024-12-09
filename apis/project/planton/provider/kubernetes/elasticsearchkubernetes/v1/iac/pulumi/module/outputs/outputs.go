package outputs

const (
	Namespace                            = "namespace"
	ElasticUsername                      = "username"
	ElasticPasswordSecretName            = "password-secret-name"
	ElasticPasswordSecretKey             = "password-secret-key"
	ElasticsearchService                 = "elasticsearch-service"
	ElasticsearchPortForwardCommand      = "elasticsearch-port-forward-command"
	ElasticsearchKubeEndpoint            = "elasticsearch-kube-endpoint"
	ElasticsearchIngressExternalHostname = "elasticsearch-ingress-external-hostname"
	ElasticsearchIngressInternalHostname = "elasticsearch-ingress-internal-hostname"

	KibanaService                 = "kibana-service"
	KibanaPortForwardCommand      = "kibana-port-forward-command"
	KibanaKubeEndpoint            = "kibana-kube-endpoint"
	KibanaIngressExternalHostname = "kibana-ingress-external-hostname"
	KibanaIngressInternalHostname = "kibana-ingress-internal-hostname"
)
