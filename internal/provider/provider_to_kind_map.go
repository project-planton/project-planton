package provider

import "github.com/project-planton/project-planton/apis/project/planton/shared"

type KindName string

var ToKindMap = map[shared.KindProvider][]KindName{
	shared.KindProvider_kind_provider_atlas: {
		"MongodbAtlas",
	},

	shared.KindProvider_kind_provider_aws: {
		"AwsAlb",
		"AwsCertManagerCert",
		"AwsCloudFront",
		"AwsDynamodb",
		"AwsEcrRepo",
		"AwsEcsCluster",
		"AwsEcsService",
		"AwsEksCluster",
		"AwsIamRole",
		"AwsLambda",
		"AwsRdsCluster",
		"AwsRdsInstance",
		"AwsRoute53Zone",
		"AwsS3Bucket",
		"AwsSecretsManager",
		"AwsSecurityGroup",
		"AwsStaticWebsite",
		"AwsVpc",
	},

	shared.KindProvider_kind_provider_azure: {
		"AzureAksCluster",
		"AzureKeyVault",
	},

	shared.KindProvider_kind_provider_confluent: {
		"ConfluentKafka",
	},

	shared.KindProvider_kind_provider_gcp: {
		"GcpArtifactRegistryRepo",
		"GcpCloudCdn",
		"GcpCloudFunction",
		"GcpCloudRun",
		"GcpCloudSql",
		"GcpDnsZone",
		"GcpGcsBucket",
		"GcpGkeAddonBundle",
		"GcpGkeCluster",
		"GcpSecretsManager",
		"GcpStaticWebsite",
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
		"TemporalKubernetes",
	},

	shared.KindProvider_kind_provider_snowflake: {
		"SnowflakeDatabase",
	},
}
