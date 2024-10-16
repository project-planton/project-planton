package manifest

import (
	mongodbatlasv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/atlas/mongodbatlas/v1"
	awscloudfrontv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awscloudfront/v1"
	awsdynamodbv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awsdynamodb/v1"
	awsfargatev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awsfargate/v1"
	awslambdav1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awslambda/v1"
	awsrdsclusterv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awsrdscluster/v1"
	awsrdsinstancev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awsrdsinstance/v1"
	awssecretsmanagerv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awssecretsmanager/v1"
	awsstaticwebsitev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awsstaticwebsite/v1"
	awsvpcv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/awsvpc/v1"
	eksclusterv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/ekscluster/v1"
	elasticcontainerservicev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/elasticcontainerservice/v1"
	route53zonev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/route53zone/v1"
	s3bucketv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/aws/s3bucket/v1"
	aksclusterv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/azure/akscluster/v1"
	azurekeyvaultv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/azure/azurekeyvault/v1"
	confluentkafkav1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/confluent/confluentkafka/v1"
	gcpartifactregistryv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpartifactregistry/v1"
	gcpcloudcdnv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpcloudcdn/v1"
	gcpcloudfunctionv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpcloudfunction/v1"
	gcpcloudrunv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpcloudrun/v1"
	gcpcloudsqlv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpcloudsql/v1"
	gcpdnszonev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpdnszone/v1"
	gcpsecretsmanagerv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpsecretsmanager/v1"
	gcpstaticwebsitev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcpstaticwebsite/v1"
	gcsbucketv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gcsbucket/v1"
	gkeclusterv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/gcp/gkecluster/v1"
	argocdkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/argocdkubernetes/v1"
	elasticsearchkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/elasticsearchkubernetes/v1"
	gitlabkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/gitlabkubernetes/v1"
	grafanakubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/grafanakubernetes/v1"
	helmreleasev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/helmrelease/v1"
	jenkinskubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/jenkinskubernetes/v1"
	kafkakubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/kafkakubernetes/v1"
	keycloakkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/keycloakkubernetes/v1"
	kuberneteshttpendpointv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/kuberneteshttpendpoint/v1"
	locustkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/locustkubernetes/v1"
	microservicekubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/microservicekubernetes/v1"
	mongodbkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/mongodbkubernetes/v1"
	neo4jkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/neo4jkubernetes/v1"
	openfgakubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/openfgakubernetes/v1"
	postgreskubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/postgreskubernetes/v1"
	prometheuskubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/prometheuskubernetes/v1"
	rediskubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/rediskubernetes/v1"
	signozkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/signozkubernetes/v1"
	solrkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/solrkubernetes/v1"
	stackjobrunnerkubernetesv1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/kubernetes/stackjobrunnerkubernetes/v1"
	snowflakedatabasev1 "buf.build/gen/go/project-planton/apis/protocolbuffers/go/project/planton/provider/snowflake/snowflakedatabase/v1"
	"google.golang.org/protobuf/proto"
	"strings"
)

type DeploymentComponent string

var DeploymentComponentMap = merge(
	providerAtlasMap,
	providerAwsMap,
	providerAzureMap,
	providerConfluentMap,
	providerGcpMap,
	providerKubernetesMap,
	providerSnowflakeMap,
)

func merge(items ...map[DeploymentComponent]proto.Message) map[DeploymentComponent]proto.Message {
	resp := make(map[DeploymentComponent]proto.Message)
	for _, i := range items {
		for k, v := range i {
			resp[k] = v
		}
	}
	return resp
}

var providerAtlasMap = map[DeploymentComponent]proto.Message{
	"mongodb-atlas": &mongodbatlasv1.MongodbAtlas{},
}

var providerAwsMap = map[DeploymentComponent]proto.Message{
	"aws-cloud-front":           &awscloudfrontv1.AwsCloudFront{},
	"aws-dynamodb":              &awsdynamodbv1.AwsDynamodb{},
	"aws-fargate":               &awsfargatev1.AwsFargate{},
	"aws-lambda":                &awslambdav1.AwsLambda{},
	"aws-rds-cluster":           &awsrdsclusterv1.AwsRdsCluster{},
	"aws-rds-instance":          &awsrdsinstancev1.AwsRdsInstance{},
	"aws-secrets-manager":       &awssecretsmanagerv1.AwsSecretsManager{},
	"aws-static-website":        &awsstaticwebsitev1.AwsStaticWebsite{},
	"aws-vpc":                   &awsvpcv1.AwsVpc{},
	"eks-cluster":               &eksclusterv1.EksCluster{},
	"elastic-container-service": &elasticcontainerservicev1.ElasticContainerService{},
	"route53-zone":              &route53zonev1.Route53Zone{},
	"s3-bucket":                 &s3bucketv1.S3Bucket{},
}

var providerConfluentMap = map[DeploymentComponent]proto.Message{
	"confluent-kafka": &confluentkafkav1.ConfluentKafka{},
}

var providerSnowflakeMap = map[DeploymentComponent]proto.Message{
	"snowflake-database": &snowflakedatabasev1.SnowflakeDatabase{},
}

var providerAzureMap = map[DeploymentComponent]proto.Message{
	"aks-cluster":     &aksclusterv1.AksCluster{},
	"azure-key-vault": &azurekeyvaultv1.AzureKeyVault{},
}

var providerGcpMap = map[DeploymentComponent]proto.Message{
	"gcp-artifact-registry": &gcpartifactregistryv1.GcpArtifactRegistry{},
	"gcp-cloud-cdn":         &gcpcloudcdnv1.GcpCloudCdn{},
	"gcp-cloud-function":    &gcpcloudfunctionv1.GcpCloudFunction{},
	"gcp-cloud-run":         &gcpcloudrunv1.GcpCloudRun{},
	"gcp-cloud-sql":         &gcpcloudsqlv1.GcpCloudSql{},
	"gcp-dns-zone":          &gcpdnszonev1.GcpDnsZone{},
	"gcp-secrets-manager":   &gcpsecretsmanagerv1.GcpSecretsManager{},
	"gcp-static-website":    &gcpstaticwebsitev1.GcpStaticWebsite{},
	"gcs-bucket":            &gcsbucketv1.GcsBucket{},
	"gke-cluster":           &gkeclusterv1.GkeCluster{},
}

var providerKubernetesMap = map[DeploymentComponent]proto.Message{
	"argocd-kubernetes":           &argocdkubernetesv1.ArgocdKubernetes{},
	"elasticsearch-kubernetes":    &elasticsearchkubernetesv1.ElasticsearchKubernetes{},
	"gitlab-kubernetes":           &gitlabkubernetesv1.GitlabKubernetes{},
	"grafana-kubernetes":          &grafanakubernetesv1.GrafanaKubernetes{},
	"helm-release":                &helmreleasev1.HelmRelease{},
	"jenkins-kubernetes":          &jenkinskubernetesv1.JenkinsKubernetes{},
	"kafka-kubernetes":            &kafkakubernetesv1.KafkaKubernetes{},
	"keycloak-kubernetes":         &keycloakkubernetesv1.KeycloakKubernetes{},
	"kubernetes-http-endpoint":    &kuberneteshttpendpointv1.KubernetesHttpEndpoint{},
	"locust-kubernetes":           &locustkubernetesv1.LocustKubernetes{},
	"microservice-kubernetes":     &microservicekubernetesv1.MicroserviceKubernetes{},
	"mongodb-kubernetes":          &mongodbkubernetesv1.MongodbKubernetes{},
	"neo4j-kubernetes":            &neo4jkubernetesv1.Neo4JKubernetes{},
	"openfga-kubernetes":          &openfgakubernetesv1.OpenfgaKubernetes{},
	"postgres-kubernetes":         &postgreskubernetesv1.PostgresKubernetes{},
	"prometheus-kubernetes":       &prometheuskubernetesv1.PrometheusKubernetes{},
	"redis-kubernetes":            &rediskubernetesv1.RedisKubernetes{},
	"signoz-kubernetes":           &signozkubernetesv1.SignozKubernetes{},
	"solr-kubernetes":             &solrkubernetesv1.SolrKubernetes{},
	"stack-job-runner-kubernetes": &stackjobrunnerkubernetesv1.StackJobRunnerKubernetes{},
}

// sanitizeString removes hyphens, spaces, and underscores, and converts the string to lowercase
func sanitizeString(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "_", "")
	return str
}

func FindMatchingComponent(input string) DeploymentComponent {
	sanitizedInput := sanitizeString(input)
	for key, _ := range DeploymentComponentMap {
		sanitizedKey := sanitizeString(string(key))
		if sanitizedKey == sanitizedInput {
			return key
		}
	}
	return ""
}
