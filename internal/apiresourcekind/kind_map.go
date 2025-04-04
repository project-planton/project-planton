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
	awsalbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsalb/v1"
	awscertmanagercertv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscertmanagercert/v1"
	awscloudfrontv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awscloudfront/v1"
	awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
	awsecrrepov1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecrrepo/v1"
	awsecsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecscluster/v1"
	awsecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecsservice/v1"
	awseksclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsekscluster/v1"
	awsiamrolev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsiamrole/v1"
	awslambdav1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awslambda/v1"
	awsrdsclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdscluster/v1"
	awsrdsinstancev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsrdsinstance/v1"
	awsroute53zonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsroute53zone/v1"
	awss3bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1"
	awssecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecretsmanager/v1"
	awssecuritygroupv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecuritygroup/v1"
	awsstaticwebsitev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsstaticwebsite/v1"
	awsvpcv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1"
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

	AwsAlbKind             ApiResourceKind = "aws-alb"
	AwsCertManagerCertKind ApiResourceKind = "aws-cert-manager-cert"
	AwsCloudFrontKind      ApiResourceKind = "aws-cloud-front"
	AwsDynamodbKind        ApiResourceKind = "aws-dynamodb"
	AwsEcrRepoKind         ApiResourceKind = "aws-ecr-repo"
	AwsEcsClusterKind      ApiResourceKind = "aws-ecs-cluster"
	AwsEcsServiceKind      ApiResourceKind = "aws-ecs-service"
	AwsEksClusterKind      ApiResourceKind = "aws-eks-cluster"
	AwsIamRoleKind         ApiResourceKind = "aws-iam-role"
	AwsLambdaKind          ApiResourceKind = "aws-lambda"
	AwsRdsClusterKind      ApiResourceKind = "aws-rds-cluster"
	AwsRdsInstanceKind     ApiResourceKind = "aws-rds-instance"
	AwsRoute53ZoneKind     ApiResourceKind = "aws-route53-zone"
	AwsS3BucketKind        ApiResourceKind = "aws-s3-bucket"
	AwsSecretsManagerKind  ApiResourceKind = "aws-secrets-manager"
	AwsSecurityGroupKind   ApiResourceKind = "aws-security-group"
	AwsStaticWebsiteKind   ApiResourceKind = "aws-static-website"
	AwsVpcKind             ApiResourceKind = "aws-vpc"

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
	AwsAlbKind:             &awsalbv1.AwsAlb{},
	AwsCertManagerCertKind: &awscertmanagercertv1.AwsCertManagerCert{},
	AwsCloudFrontKind:      &awscloudfrontv1.AwsCloudFront{},
	AwsDynamodbKind:        &awsdynamodbv1.AwsDynamodb{},
	AwsEcrRepoKind:         &awsecrrepov1.AwsEcrRepo{},
	AwsEcsClusterKind:      &awsecsclusterv1.AwsEcsCluster{},
	AwsEcsServiceKind:      &awsecsservicev1.AwsEcsService{},
	AwsEksClusterKind:      &awseksclusterv1.AwsEksCluster{},
	AwsIamRoleKind:         &awsiamrolev1.AwsIamRole{},
	AwsLambdaKind:          &awslambdav1.AwsLambda{},
	AwsRdsClusterKind:      &awsrdsclusterv1.AwsRdsCluster{},
	AwsRdsInstanceKind:     &awsrdsinstancev1.AwsRdsInstance{},
	AwsRoute53ZoneKind:     &awsroute53zonev1.AwsRoute53Zone{},
	AwsS3BucketKind:        &awss3bucketv1.AwsS3Bucket{},
	AwsSecretsManagerKind:  &awssecretsmanagerv1.AwsSecretsManager{},
	AwsSecurityGroupKind:   &awssecuritygroupv1.AwsSecurityGroup{},
	AwsStaticWebsiteKind:   &awsstaticwebsitev1.AwsStaticWebsite{},
	AwsVpcKind:             &awsvpcv1.AwsVpc{},
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
