package module

// These constants define the output keys we export from our Pulumi module.
// They correspond to fields in EcsServiceStackOutputs, so that other parts
// of ProjectPlanton can consume them consistently.
const (
	OpEcsServiceName       = "ecs_service_name"
	OpEcsClusterName       = "ecs_cluster_name"
	OpLoadBalancerDnsName  = "load_balancer_dns_name"
	OpServiceUrl           = "service_url"
	OpServiceDiscoveryName = "service_discovery_name"
)
