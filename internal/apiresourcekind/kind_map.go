package apiresourcekind

import (
	awscredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/awscredential/v1"
	azurecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/azurecredential/v1"
	confluentcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/confluentcredential/v1"
	gcpcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	kubernetesclustercredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1"
	mongodbatlascredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/mongodbatlascredential/v1"
	snowflakecredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/snowflakecredential/v1"
	mongodbatlasv1 "github.com/project-planton/project-planton/apis/project/planton/provider/atlas/mongodbatlas/v1"
	awscloudfrontv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscloudfront/v1"
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	awslambdav1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awslambda/v1"
	awsrdsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdscluster/v1"
	awsrdsinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdsinstance/v1"
	awssecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecretsmanager/v1"
	awssecuritygroupv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecuritygroup/v1"
	awsstaticwebsitev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsstaticwebsite/v1"
	awsvpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1"
	ecsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecscluster/v1"
	ecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ecsservice/v1"
	eksclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/ekscluster/v1"
	route53zonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/route53zone/v1"
	s3bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/s3bucket/v1"
	aksclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/azure/akscluster/v1"
	azurekeyvaultv1 "github.com/project-planton/project-planton/apis/project/planton/provider/azure/azurekeyvault/v1"
	confluentkafkav1 "github.com/project-planton/project-planton/apis/project/planton/provider/confluent/confluentkafka/v1"
	gcpartifactregistryrepov1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistryrepo/v1"
	gcpcloudcdnv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudcdn/v1"
	gcpcloudfunctionv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudfunction/v1"
	gcpcloudrunv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudrun/v1"
	gcpcloudsqlv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudsql/v1"
	gcpdnszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpdnszone/v1"
	gcpsecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpsecretsmanager/v1"
	gcpstaticwebsitev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpstaticwebsite/v1"
	gcsbucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcsbucket/v1"
	gkeaddonbundlev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkeaddonbundle/v1"
	gkeclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gkecluster/v1"
	argocdkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/argocdkubernetes/v1"
	cronjobkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/cronjobkubernetes/v1"
	elasticsearchkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/elasticsearchkubernetes/v1"
	gitlabkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/gitlabkubernetes/v1"
	grafanakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/grafanakubernetes/v1"
	helmreleasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/helmrelease/v1"
	jenkinskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/jenkinskubernetes/v1"
	kafkakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kafkakubernetes/v1"
	keycloakkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/keycloakkubernetes/v1"
	kuberneteshttpendpointv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/kuberneteshttpendpoint/v1"
	locustkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/locustkubernetes/v1"
	microservicekubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/microservicekubernetes/v1"
	mongodbkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/mongodbkubernetes/v1"
	neo4jkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/neo4jkubernetes/v1"
	openfgakubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/openfgakubernetes/v1"
	postgreskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/postgreskubernetes/v1"
	prometheuskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/prometheuskubernetes/v1"
	rediskubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/rediskubernetes/v1"
	signozkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/signozkubernetes/v1"
	solrkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/solrkubernetes/v1"
	stackjobrunnerkubernetesv1 "github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/stackjobrunnerkubernetes/v1"
	snowflakedatabasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/snowflake/snowflakedatabase/v1"
	"google.golang.org/protobuf/proto"
	"strings"
)

type ApiResourceKind string

const (
	AwsCredentialKind               ApiResourceKind = "aws-credential"
	AzureCredentialKind             ApiResourceKind = "azure-credential"
	ConfluentCredentialKind         ApiResourceKind = "confluent-credential"
	GcpCredentialKind               ApiResourceKind = "gcp-credential"
	KubernetesClusterCredentialKind ApiResourceKind = "kubernetes-cluster-credential"
	MongodbAtlasCredentialKind      ApiResourceKind = "mongodb-atlas-credential"
	SnowflakeCredentialKind         ApiResourceKind = "snowflake-credential"

	MongodbAtlasKind ApiResourceKind = "mongodb-atlas"

	AwsCloudFrontKind     ApiResourceKind = "aws-cloud-front"
	AwsDynamodbKind       ApiResourceKind = "aws-dynamodb"
	AwsLambdaKind         ApiResourceKind = "aws-lambda"
	AwsRdsClusterKind     ApiResourceKind = "aws-rds-cluster"
	AwsRdsInstanceKind    ApiResourceKind = "aws-rds-instance"
	AwsSecretsManagerKind ApiResourceKind = "aws-secrets-manager"
	AwsSecurityGroupKind  ApiResourceKind = "aws-security-group"
	AwsStaticWebsiteKind  ApiResourceKind = "aws-static-website"
	AwsVpcKind            ApiResourceKind = "aws-vpc"
	EcsClusterKind        ApiResourceKind = "ecs-cluster"
	EcsServiceKind        ApiResourceKind = "ecs-service"
	EksClusterKind        ApiResourceKind = "eks-cluster"
	Route53ZoneKind       ApiResourceKind = "route53-zone"
	S3BucketKind          ApiResourceKind = "s3-bucket"

	AksClusterKind    ApiResourceKind = "aks-cluster"
	AzureKeyVaultKind ApiResourceKind = "azure-key-vault"

	ConfluentKafkaKind ApiResourceKind = "confluent-kafka"

	GcpArtifactRegistryRepoKind ApiResourceKind = "gcp-artifact-registry-repo"
	GcpCloudCdnKind             ApiResourceKind = "gcp-cloud-cdn"
	GcpCloudFunctionKind        ApiResourceKind = "gcp-cloud-function"
	GcpCloudRunKind             ApiResourceKind = "gcp-cloud-run"
	GcpCloudSqlKind             ApiResourceKind = "gcp-cloud-sql"
	GcpDnsZoneKind              ApiResourceKind = "gcp-dns-zone"
	GcpSecretsManagerKind       ApiResourceKind = "gcp-secrets-manager"
	GcpStaticWebsiteKind        ApiResourceKind = "gcp-static-website"
	GcsBucketKind               ApiResourceKind = "gcs-bucket"
	GkeClusterKind              ApiResourceKind = "gke-cluster"
	GkeAddonBundleKind          ApiResourceKind = "gke-addon-bundle"

	ArgocdKubernetesKind         ApiResourceKind = "argocd-kubernetes"
	ElasticsearchKubernetesKind  ApiResourceKind = "elasticsearch-kubernetes"
	GitlabKubernetesKind         ApiResourceKind = "gitlab-kubernetes"
	GrafanaKubernetesKind        ApiResourceKind = "grafana-kubernetes"
	HelmReleaseKind              ApiResourceKind = "helm-release"
	JenkinsKubernetesKind        ApiResourceKind = "jenkins-kubernetes"
	KafkaKubernetesKind          ApiResourceKind = "kafka-kubernetes"
	KeycloakKubernetesKind       ApiResourceKind = "keycloak-kubernetes"
	KubernetesHttpEndpointKind   ApiResourceKind = "kubernetes-http-endpoint"
	LocustKubernetesKind         ApiResourceKind = "locust-kubernetes"
	MicroserviceKubernetesKind   ApiResourceKind = "microservice-kubernetes"
	CronJobKubernetesKind        ApiResourceKind = "cron-job-kubernetes"
	MongodbKubernetesKind        ApiResourceKind = "mongodb-kubernetes"
	Neo4JKubernetesKind          ApiResourceKind = "neo4j-kubernetes"
	OpenfgaKubernetesKind        ApiResourceKind = "openfga-kubernetes"
	PostgresKubernetesKind       ApiResourceKind = "postgres-kubernetes"
	PrometheusKubernetesKind     ApiResourceKind = "prometheus-kubernetes"
	RedisKubernetesKind          ApiResourceKind = "redis-kubernetes"
	SignozKubernetesKind         ApiResourceKind = "signoz-kubernetes"
	SolrKubernetesKind           ApiResourceKind = "solr-kubernetes"
	StackJobRunnerKubernetesKind ApiResourceKind = "stack-job-runner-kubernetes"

	SnowflakeDatabaseKind ApiResourceKind = "snowflake-database"
)

var ToMessageMap = merge(
	credentialMap,
	providerAtlasMap,
	providerAwsMap,
	providerAzureMap,
	providerConfluentMap,
	providerGcpMap,
	providerKubernetesMap,
	providerSnowflakeMap,
)

func merge(items ...map[ApiResourceKind]proto.Message) map[ApiResourceKind]proto.Message {
	resp := make(map[ApiResourceKind]proto.Message)
	for _, i := range items {
		for k, v := range i {
			resp[k] = v
		}
	}
	return resp
}

var credentialMap = map[ApiResourceKind]proto.Message{
	AwsCredentialKind:               &awscredentialv1.AwsCredential{},
	AzureCredentialKind:             &azurecredentialv1.AzureCredential{},
	ConfluentCredentialKind:         &confluentcredentialv1.ConfluentCredential{},
	GcpCredentialKind:               &gcpcredentialv1.GcpCredential{},
	KubernetesClusterCredentialKind: &kubernetesclustercredentialv1.KubernetesClusterCredential{},
	MongodbAtlasCredentialKind:      &mongodbatlascredentialv1.MongodbAtlasCredential{},
	SnowflakeCredentialKind:         &snowflakecredentialv1.SnowflakeCredential{},
}

var providerAtlasMap = map[ApiResourceKind]proto.Message{
	MongodbAtlasKind: &mongodbatlasv1.MongodbAtlas{},
}

var providerAwsMap = map[ApiResourceKind]proto.Message{
	AwsCloudFrontKind:     &awscloudfrontv1.AwsCloudFront{},
	AwsDynamodbKind:       &awsdynamodbv1.AwsDynamodb{},
	AwsLambdaKind:         &awslambdav1.AwsLambda{},
	AwsRdsClusterKind:     &awsrdsclusterv1.AwsRdsCluster{},
	AwsRdsInstanceKind:    &awsrdsinstancev1.AwsRdsInstance{},
	AwsSecretsManagerKind: &awssecretsmanagerv1.AwsSecretsManager{},
	AwsSecurityGroupKind:  &awssecuritygroupv1.AwsSecurityGroup{},
	AwsStaticWebsiteKind:  &awsstaticwebsitev1.AwsStaticWebsite{},
	AwsVpcKind:            &awsvpcv1.AwsVpc{},
	EcsClusterKind:        &ecsclusterv1.EcsCluster{},
	EcsServiceKind:        &ecsservicev1.EcsService{},
	EksClusterKind:        &eksclusterv1.EksCluster{},
	Route53ZoneKind:       &route53zonev1.Route53Zone{},
	S3BucketKind:          &s3bucketv1.S3Bucket{},
}

var providerConfluentMap = map[ApiResourceKind]proto.Message{
	ConfluentKafkaKind: &confluentkafkav1.ConfluentKafka{},
}

var providerSnowflakeMap = map[ApiResourceKind]proto.Message{
	SnowflakeDatabaseKind: &snowflakedatabasev1.SnowflakeDatabase{},
}

var providerAzureMap = map[ApiResourceKind]proto.Message{
	AksClusterKind:    &aksclusterv1.AksCluster{},
	AzureKeyVaultKind: &azurekeyvaultv1.AzureKeyVault{},
}

var providerGcpMap = map[ApiResourceKind]proto.Message{
	GcpArtifactRegistryRepoKind: &gcpartifactregistryrepov1.GcpArtifactRegistryRepo{},
	GcpCloudCdnKind:             &gcpcloudcdnv1.GcpCloudCdn{},
	GcpCloudFunctionKind:        &gcpcloudfunctionv1.GcpCloudFunction{},
	GcpCloudRunKind:             &gcpcloudrunv1.GcpCloudRun{},
	GcpCloudSqlKind:             &gcpcloudsqlv1.GcpCloudSql{},
	GcpDnsZoneKind:              &gcpdnszonev1.GcpDnsZone{},
	GcpSecretsManagerKind:       &gcpsecretsmanagerv1.GcpSecretsManager{},
	GcpStaticWebsiteKind:        &gcpstaticwebsitev1.GcpStaticWebsite{},
	GcsBucketKind:               &gcsbucketv1.GcsBucket{},
	GkeClusterKind:              &gkeclusterv1.GkeCluster{},
	GkeAddonBundleKind:          &gkeaddonbundlev1.GkeAddonBundle{},
}

var providerKubernetesMap = map[ApiResourceKind]proto.Message{
	ArgocdKubernetesKind:         &argocdkubernetesv1.ArgocdKubernetes{},
	ElasticsearchKubernetesKind:  &elasticsearchkubernetesv1.ElasticsearchKubernetes{},
	GitlabKubernetesKind:         &gitlabkubernetesv1.GitlabKubernetes{},
	GrafanaKubernetesKind:        &grafanakubernetesv1.GrafanaKubernetes{},
	HelmReleaseKind:              &helmreleasev1.HelmRelease{},
	JenkinsKubernetesKind:        &jenkinskubernetesv1.JenkinsKubernetes{},
	KafkaKubernetesKind:          &kafkakubernetesv1.KafkaKubernetes{},
	KeycloakKubernetesKind:       &keycloakkubernetesv1.KeycloakKubernetes{},
	KubernetesHttpEndpointKind:   &kuberneteshttpendpointv1.KubernetesHttpEndpoint{},
	LocustKubernetesKind:         &locustkubernetesv1.LocustKubernetes{},
	MicroserviceKubernetesKind:   &microservicekubernetesv1.MicroserviceKubernetes{},
	CronJobKubernetesKind:        &cronjobkubernetesv1.CronJobKubernetes{},
	MongodbKubernetesKind:        &mongodbkubernetesv1.MongodbKubernetes{},
	Neo4JKubernetesKind:          &neo4jkubernetesv1.Neo4JKubernetes{},
	OpenfgaKubernetesKind:        &openfgakubernetesv1.OpenfgaKubernetes{},
	PostgresKubernetesKind:       &postgreskubernetesv1.PostgresKubernetes{},
	PrometheusKubernetesKind:     &prometheuskubernetesv1.PrometheusKubernetes{},
	RedisKubernetesKind:          &rediskubernetesv1.RedisKubernetes{},
	SignozKubernetesKind:         &signozkubernetesv1.SignozKubernetes{},
	SolrKubernetesKind:           &solrkubernetesv1.SolrKubernetes{},
	StackJobRunnerKubernetesKind: &stackjobrunnerkubernetesv1.StackJobRunnerKubernetes{},
}

// sanitizeString removes hyphens, spaces, and underscores, and converts the string to lowercase
func sanitizeString(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "_", "")
	return str
}

func FindMatchingComponent(input string) ApiResourceKind {
	sanitizedInput := sanitizeString(input)
	for key, _ := range ToMessageMap {
		sanitizedKey := sanitizeString(string(key))
		if sanitizedKey == sanitizedInput {
			return key
		}
	}
	return ""
}
