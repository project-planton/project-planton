package outputs

const (
	ClusterEndpoint              = "cluster_endpoint"
	ClusterCaData                = "cluster_ca_data"
	ExternalNatIp                = "external_nat_ip"
	IngressInternalIp            = "ingress_internal_ip"
	IngressExternalIp            = "ingress_external_ip"
	CertManagerGsaEmail          = "cert_manager_gsa_email"
	ExternalSecretsGsaEmail      = "external_secrets_gsa_email"
	WorkloadDeployerGsaEmail     = "workload_deployer_gsa_email"
	WorkloadDeployerGsaKeyBase64 = "workload_deployer_gsa_key_base64"
	ExternalDnsGsaEmail          = "external_dns_gsa_email"
	GkeWebhooksFirewallSelfLink  = "gke_webhooks_firewall_self_link"
	NetworkSelfLink              = "network_self_link"
	SubNetworkSelfLink           = "sub_network_self_link"
	RouterNatName                = "router_nat_name"
	RouterSelfLink               = "router_self_link"
)
