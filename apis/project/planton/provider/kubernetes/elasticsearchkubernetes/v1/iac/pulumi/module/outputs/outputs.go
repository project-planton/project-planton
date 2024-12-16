package outputs

const (
	NAMESPACE                          = "namespace"
	ELASTICSEARCH_SERVICE              = "elasticsearch.service"
	ELASTICSEARCH_PORT_FORWARD_COMMAND = "elasticsearch.port_forward_command"
	ELASTICSEARCH_KUBE_ENDPOINT        = "elasticsearch.kube_endpoint"
	ELASTICSEARCH_EXTERNAL_HOSTNAME    = "elasticsearch.external_hostname"
	ELASTICSEARCH_INTERNAL_HOSTNAME    = "elasticsearch.internal_hostname"
	ELASTICSEARCH_USERNAME             = "elasticsearch.username"
	ELASTICSEARCH_PASSWORD_SECRET_NAME = "elasticsearch.password_secret.name"
	ELASTICSEARCH_PASSWORD_SECRET_KEY  = "elasticsearch.password_secret.key"
	KIBANA_SERVICE                     = "kibana.service"
	KIBANA_PORT_FORWARD_COMMAND        = "kibana.port_forward_command"
	KIBANA_KUBE_ENDPOINT               = "kibana.kube_endpoint"
	KIBANA_EXTERNAL_HOSTNAME           = "kibana.external_hostname"
	KIBANA_INTERNAL_HOSTNAME           = "kibana.internal_hostname"
)
