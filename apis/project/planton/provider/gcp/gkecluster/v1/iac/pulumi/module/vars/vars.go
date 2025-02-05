package vars

var (
	GkeReleaseChannel = "STABLE"
	// GoogleFolderAndProjectPlantonCloudPrefix will be prefixed for all ids of folders and projects
	//created by this module for easy identification
	GoogleFolderAndProjectPlantonCloudPrefix = "gke"
	NetworkProjectApis                       = []string{
		"compute.googleapis.com",
		"container.googleapis.com",
		"dns.googleapis.com",
	}

	ContainerClusterProjectApis = []string{
		"compute.googleapis.com",
		"container.googleapis.com",
		"secretmanager.googleapis.com",
		"dns.googleapis.com",
	}

	WorkloadIdentityKubeAnnotationKey = "iam.gke.io/gcp-service-account"

	// SubNetworkCidr 10.0.0.0/14
	// this subnet will be divided into two equal halves for pod-secondary-ip-range and service-secondary-ip-range
	//https://jodies.de/ipcalc?host=10.0.0.0&mask1=14&mask2=15
	SubNetworkCidr = "10.0.0.0/14"

	// KubernetesPodSecondaryIpRange https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#cluster_sizing_secondary_range_pods
	KubernetesPodSecondaryIpRange = "10.4.0.0/16"
	// KubernetesServiceSecondaryIpRange https://cloud.google.com/kubernetes-engine/docs/concepts/alias-ips#cluster_sizing_secondary_range_svcs
	KubernetesServiceSecondaryIpRange = "10.5.0.0/16"

	ApiServerIpCidr                                     = "172.16.0.0/28"
	ClusterMasterAuthorizedNetworksCidrBlock            = "0.0.0.0/0"
	ClusterMasterAuthorizedNetworksCidrBlockDescription = "kubectl-from-anywhere"
	ApiServerWebhookPort                                = "8443"
	IstioPilotWebhookPort                               = "15017"

	// WorkloadDeployServiceAccountName name of the google service account to
	//be used for deploying workloads to the gke cluster.
	WorkloadDeployServiceAccountName = "workload-deployer"
)
