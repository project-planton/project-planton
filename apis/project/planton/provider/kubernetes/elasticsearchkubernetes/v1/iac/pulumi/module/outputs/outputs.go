package outputs

const (
	Namespace                       = "namespace"
	ElasticsearchService            = "elasticsearch.service"
	ElasticsearchPortForwardCommand = "elasticsearch.port_forward_command"
	ElasticsearchKubeEndpoint       = "elasticsearch.kube_endpoint"
	ElasticsearchExternalHostname   = "elasticsearch.external_hostname"
	ElasticsearchInternalHostname   = "elasticsearch.internal_hostname"
	ElasticsearchUsername           = "elasticsearch.username"
	ElasticsearchPasswordSecretName = "elasticsearch.password_secret.name"
	ElasticsearchPasswordSecretKey  = "elasticsearch.password_secret.key"
	KibanaService                   = "kibana.service"
	KibanaPortForwardCommand        = "kibana.port_forward_command"
	KibanaKubeEndpoint              = "kibana.kube_endpoint"
	KibanaExternalHostname          = "kibana.external_hostname"
	KibanaInternalHostname          = "kibana.internal_hostname"
)
