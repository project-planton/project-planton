package provider

import "github.com/project-planton/project-planton/apis/project/planton/shared"

type KindName string

var ToKindMap = map[iac.KindProvider][]KindName{
	iac.KindProvider_kind_provider_atlas: {
		"MongodbAtlas",
	},

	iac.KindProvider_kind_provider_aws: {
		"AwsCloudFront",
		"AwsDynamodb",
		"AwsFargate",
		"AwsLambda",
		"AwsRdsCluster",
		"AwsRdsInstance",
		"AwsSecretsManager",
		"AwsStaticWebsite",
		"AwsVpc",
		"EksCluster",
		"ElasticContainerService",
		"Route53Zone",
		"S3Bucket",
	},

	iac.KindProvider_kind_provider_azure: {
		"AksCluster",
		"AzureKeyVault",
	},

	iac.KindProvider_kind_provider_confluent: {
		"ConfluentKafka",
	},

	iac.KindProvider_kind_provider_gcp: {
		"GcpArtifactRegistryRepo",
		"GcpCloudCdn",
		"GcpCloudFunction",
		"GcpCloudRun",
		"GcpCloudSql",
		"GcpDnsZone",
		"GcpSecretsManager",
		"GcpStaticWebsite",
		"GcsBucket",
		"GkeCluster",
	},

	iac.KindProvider_kind_provider_kubernetes: {
		"ArgocdKubernetes",
		"ElasticsearchKubernetes",
		"GitlabKubernetes",
		"GrafanaKubernetes",
		"HelmRelease",
		"JenkinsKubernetes",
		"KafkaKubernetes",
		"KeycloakKubernetes",
		"KubernetesHttpEndpoint",
		"LocustKubernetes",
		"MicroserviceKubernetes",
		"MongodbKubernetes",
		"Neo4jKubernetes",
		"OpenfgaKubernetes",
		"PostgresKubernetes",
		"PrometheusKubernetes",
		"RedisKubernetes",
		"SignozKubernetes",
		"SolrKubernetes",
		"StackJobRunnerKubernetes",
	},

	iac.KindProvider_kind_provider_snowflake: {
		"SnowflakeDatabase",
	},
}
