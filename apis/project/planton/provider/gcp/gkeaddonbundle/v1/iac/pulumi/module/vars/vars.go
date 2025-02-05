package vars

var (
	WorkloadIdentityKubeAnnotationKey = "iam.gke.io/gcp-service-account"
	IstioPilotWebhookPort             = "15017"

	GatewayApis = struct {
		CrdDownloadBaseUrl string
		CrdFiles           []string
	}{
		//version is included in the base-url
		CrdDownloadBaseUrl: "https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/v1.1.0/config/crd/standard",
		CrdFiles: []string{
			"gateway.networking.k8s.io_gatewayclasses.yaml",
			"gateway.networking.k8s.io_gateways.yaml",
			"gateway.networking.k8s.io_grpcroutes.yaml",
			"gateway.networking.k8s.io_httproutes.yaml",
			"gateway.networking.k8s.io_referencegrants.yaml",
		},
	}

	CertManager = struct {
		Namespace                          string
		HelmChartName                      string
		HelmChartRepo                      string
		HelmChartVersion                   string
		KsaName                            string
		SelfSignedIssuerName               string
		LetsEncryptServer                  string
		LetsEncryptClusterIssuerSecretName string
		Http01ChallengeSolverIngressClass  string
	}{
		Namespace:                          "cert-manager",
		HelmChartName:                      "cert-manager",
		HelmChartRepo:                      "https://charts.jetstack.io",
		HelmChartVersion:                   "v1.15.2",
		KsaName:                            "cert-manager",
		SelfSignedIssuerName:               "self-signed",
		LetsEncryptServer:                  "https://acme-v02.api.letsencrypt.org/directory",
		LetsEncryptClusterIssuerSecretName: "letsencrypt-production",
		Http01ChallengeSolverIngressClass:  "istio",
	}

	ExternalDns = struct {
		Namespace               string
		HelmChartName           string
		HelmChartRepo           string
		HelmChartVersion        string
		KsaName                 string
		GcpCloudDnsProviderName string
	}{
		Namespace:               "external-dns",
		HelmChartName:           "external-dns",
		HelmChartRepo:           "https://kubernetes-sigs.github.io/external-dns/",
		HelmChartVersion:        "1.14.4", //https://github.com/kubernetes-sigs/external-dns/blob/v0.14.2/charts/external-dns/Chart.yaml#L5
		KsaName:                 "external-dns",
		GcpCloudDnsProviderName: "google",
	}

	ExternalSecrets = struct {
		Namespace                               string
		HelmChartName                           string
		HelmChartRepo                           string
		HelmChartVersion                        string
		KsaName                                 string
		SecretsPollingIntervalSeconds           int
		GcpSecretsManagerClusterSecretStoreName string
	}{
		Namespace:        "external-secrets",
		HelmChartName:    "external-secrets",
		HelmChartRepo:    "https://charts.external-secrets.io",
		HelmChartVersion: "v0.9.20",
		KsaName:          "external-secrets",
		//caution: polling interval frequency may have effect on provider costs on some platforms
		SecretsPollingIntervalSeconds:           10,
		GcpSecretsManagerClusterSecretStoreName: "gcp-secrets-manager",
	}

	IngressNginx = struct {
		Namespace        string
		HelmChartName    string
		HelmChartRepo    string
		HelmChartVersion string
	}{
		Namespace:     "ingress-nginx",
		HelmChartName: "ingress-nginx",
		HelmChartRepo: "https://kubernetes.github.io/ingress-nginx",
		//https://github.com/kubernetes/ingress-nginx/blob/main/charts/ingress-nginx/Chart.yaml#L26C9-L26C14
		HelmChartVersion: "4.11.1",
	}

	ZalandoPostgresOperator = struct {
		Namespace        string
		HelmChartName    string
		HelmChartRepo    string
		HelmChartVersion string
	}{
		Namespace:     "postgres-operator",
		HelmChartName: "postgres-operator",
		HelmChartRepo: "https://opensource.zalando.com/postgres-operator/charts/postgres-operator",
		//https://github.com/zalando/postgres-operator/blob/v1.12.2/charts/postgres-operator/Chart.yaml#L3
		HelmChartVersion: "1.12.2",
	}

	SolrOperator = struct {
		Namespace              string
		CrdManifestDownloadUrl string
		HelmChartName          string
		HelmChartRepo          string
		HelmChartVersion       string
	}{
		Namespace: "solr-operator",
		//version in the url should match the helm-chart version and should be prefixed with 'v'
		CrdManifestDownloadUrl: "https://solr.apache.org/operator/downloads/crds/v0.7.0/all-with-dependencies.yaml",
		HelmChartName:          "solr-operator",
		HelmChartRepo:          "https://solr.apache.org/charts",
		//https://github.com/apache/solr-operator/blob/v0.8.1/helm/solr-operator/Chart.yaml#L18
		HelmChartVersion: "0.7.0",
	}

	StrimziKafkaOperator = struct {
		Namespace        string
		HelmChartName    string
		HelmChartRepo    string
		HelmChartVersion string
	}{
		Namespace:     "strimzi-operator",
		HelmChartName: "strimzi-kafka-operator",
		//https://artifacthub.io/packages/helm/strimzi/strimzi-kafka-operator
		HelmChartRepo: "https://strimzi.io/charts/",
		//check artifact-hub for the latest version
		HelmChartVersion: "0.42.0",
	}

	Istio = struct {
		SystemNamespace                        string
		GatewayNamespace                       string
		HelmChartsRepo                         string
		HelmChartsVersion                      string
		BaseHelmChartName                      string
		IstiodHelmChartName                    string
		GatewayHelmChartName                   string
		IngressInternalLoadBalancerServiceName string
		IngressExternalLoadBalancerServiceName string
		IngressExternalServiceAnnotations      map[string]string
		IngressInternalServiceAnnotations      map[string]string
		SelectorLabels                         map[string]string
		HttpPort                               int
		HttpsPort                              int
		IstiodStatusPort                       int
	}{
		SystemNamespace:  "istio-system",
		GatewayNamespace: "istio-ingress",
		HelmChartsRepo:   "https://istio-release.storage.googleapis.com/charts",
		//all three charts are versioned separately but consistently. so we use the same version for all three charts.
		HelmChartsVersion: "1.22.3",
		//https://artifacthub.io/packages/helm/istio-official/base
		BaseHelmChartName: "base",
		//https://artifacthub.io/packages/helm/istio-official/istiod
		IstiodHelmChartName: "istiod",
		//https://artifacthub.io/packages/helm/istio-official/gateway
		GatewayHelmChartName:                   "gateway",
		IngressExternalLoadBalancerServiceName: "ingress-external",
		IngressInternalLoadBalancerServiceName: "ingress-internal",
		IngressExternalServiceAnnotations: map[string]string{
			"cloud.google.com/load-balancer-type": "external",
		},
		IngressInternalServiceAnnotations: map[string]string{
			"cloud.google.com/load-balancer-type": "internal",
		},
		SelectorLabels: map[string]string{
			"app":   "gateway",
			"istio": "gateway",
		},
		HttpPort:         80,
		HttpsPort:        443,
		IstiodStatusPort: 15021,
	}

	ElasticOperator = struct {
		Namespace        string
		HelmChartName    string
		HelmChartRepo    string
		HelmChartVersion string
	}{
		Namespace:     "elastic-system",
		HelmChartName: "eck-operator",
		HelmChartRepo: "https://helm.elastic.co",
		//https://github.com/elastic/cloud-on-k8s/blob/main/deploy/eck-operator/values.yaml
		HelmChartVersion: "2.14.0",
	}
)
