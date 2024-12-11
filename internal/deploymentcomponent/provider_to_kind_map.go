package deploymentcomponent

import "github.com/project-planton/project-planton/apis/project/planton/shared"

type KindName string

var ProviderToKindMap = map[shared.KindProvider_KindProvider][]KindName{
	shared.KindProvider_atlas: {
		"MongodbAtlas",
	},

	shared.KindProvider_aws: {
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

	shared.KindProvider_azure: {
		"AksCluster",
		"AzureKeyVault",
	},

	shared.KindProvider_confluent: {
		"ConfluentKafka",
	},

	shared.KindProvider_gcp: {
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

	shared.KindProvider_kubernetes: {
		"ArgoCdKubernetes",
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

	shared.KindProvider_snowflake: {
		"SnowflakeDatabase",
	},
}
