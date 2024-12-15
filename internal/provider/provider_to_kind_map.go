package provider

import "github.com/project-planton/project-planton/apis/project/planton/shared"

type KindName string

var ToKindMap = map[shared.KindProvider][]KindName{
	shared.KindProvider_kind_provider_atlas: {
		"MongodbAtlas",
	},

	shared.KindProvider_kind_provider_aws: {
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

	shared.KindProvider_kind_provider_azure: {
		"AksCluster",
		"AzureKeyVault",
	},

	shared.KindProvider_kind_provider_confluent: {
		"ConfluentKafka",
	},

	shared.KindProvider_kind_provider_gcp: {
		"GcpArtifactRegistry",
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

	shared.KindProvider_kind_provider_kubernetes: {
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

	shared.KindProvider_kind_provider_snowflake: {
		"SnowflakeDatabase",
	},
}
