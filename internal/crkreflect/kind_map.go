package crkreflect

import (
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
	azureaksclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/azure/azureakscluster/v1"
	azurekeyvaultv1 "github.com/project-planton/project-planton/apis/project/planton/provider/azure/azurekeyvault/v1"
	confluentkafkav1 "github.com/project-planton/project-planton/apis/project/planton/provider/confluent/confluentkafka/v1"
	gcpartifactregistryrepov1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistryrepo/v1"
	gcpcloudcdnv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudcdn/v1"
	gcpcloudfunctionv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudfunction/v1"
	gcpcloudrunv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudrun/v1"
	gcpcloudsqlv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudsql/v1"
	gcpdnszonev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpdnszone/v1"
	gcpgcsbucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgcsbucket/v1"
	gcpgkeaddonbundlev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkeaddonbundle/v1"
	gcpgkeclusterv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpgkecluster/v1"
	gcpsecretsmanagerv1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpsecretsmanager/v1"
	gcpstaticwebsitev1 "github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpstaticwebsite/v1"
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
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"google.golang.org/protobuf/proto"
)

var ToMessageMap = merge(
	providerAtlasMap,
	providerAwsMap,
	providerAzureMap,
	providerConfluentMap,
	providerGcpMap,
	providerKubernetesMap,
	providerSnowflakeMap,
)

func merge(items ...map[cloudresourcekind.CloudResourceKind]proto.Message) map[cloudresourcekind.CloudResourceKind]proto.Message {
	resp := make(map[cloudresourcekind.CloudResourceKind]proto.Message)
	for _, i := range items {
		for k, v := range i {
			resp[k] = v
		}
	}
	return resp
}

var providerAtlasMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
	cloudresourcekind.CloudResourceKind_MongodbAtlas: &mongodbatlasv1.MongodbAtlas{},
}

var providerAwsMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
	cloudresourcekind.CloudResourceKind_AwsAlb:             &awsalbv1.AwsAlb{},
	cloudresourcekind.CloudResourceKind_AwsCertManagerCert: &awscertmanagercertv1.AwsCertManagerCert{},
	cloudresourcekind.CloudResourceKind_AwsCloudFront:      &awscloudfrontv1.AwsCloudFront{},
	cloudresourcekind.CloudResourceKind_AwsDynamodb:        &awsdynamodbv1.AwsDynamodb{},
	cloudresourcekind.CloudResourceKind_AwsEcrRepo:         &awsecrrepov1.AwsEcrRepo{},
	cloudresourcekind.CloudResourceKind_AwsEcsCluster:      &awsecsclusterv1.AwsEcsCluster{},
	cloudresourcekind.CloudResourceKind_AwsEcsService:      &awsecsservicev1.AwsEcsService{},
	cloudresourcekind.CloudResourceKind_AwsEksCluster:      &awseksclusterv1.AwsEksCluster{},
	cloudresourcekind.CloudResourceKind_AwsIamRole:         &awsiamrolev1.AwsIamRole{},
	cloudresourcekind.CloudResourceKind_AwsLambda:          &awslambdav1.AwsLambda{},
	cloudresourcekind.CloudResourceKind_AwsRdsCluster:      &awsrdsclusterv1.AwsRdsCluster{},
	cloudresourcekind.CloudResourceKind_AwsRdsInstance:     &awsrdsinstancev1.AwsRdsInstance{},
	cloudresourcekind.CloudResourceKind_AwsRoute53Zone:     &awsroute53zonev1.AwsRoute53Zone{},
	cloudresourcekind.CloudResourceKind_AwsS3Bucket:        &awss3bucketv1.AwsS3Bucket{},
	cloudresourcekind.CloudResourceKind_AwsSecretsManager:  &awssecretsmanagerv1.AwsSecretsManager{},
	cloudresourcekind.CloudResourceKind_AwsSecurityGroup:   &awssecuritygroupv1.AwsSecurityGroup{},
	cloudresourcekind.CloudResourceKind_AwsStaticWebsite:   &awsstaticwebsitev1.AwsStaticWebsite{},
	cloudresourcekind.CloudResourceKind_AwsVpc:             &awsvpcv1.AwsVpc{},
}

var providerConfluentMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
	cloudresourcekind.CloudResourceKind_ConfluentKafka: &confluentkafkav1.ConfluentKafka{},
}

var providerSnowflakeMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
	cloudresourcekind.CloudResourceKind_SnowflakeDatabase: &snowflakedatabasev1.SnowflakeDatabase{},
}

var providerAzureMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
	cloudresourcekind.CloudResourceKind_AzureAksCluster: &azureaksclusterv1.AzureAksCluster{},
	cloudresourcekind.CloudResourceKind_AzureKeyVault:   &azurekeyvaultv1.AzureKeyVault{},
}

var providerGcpMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
	cloudresourcekind.CloudResourceKind_GcpArtifactRegistryRepo: &gcpartifactregistryrepov1.GcpArtifactRegistryRepo{},
	cloudresourcekind.CloudResourceKind_GcpCloudCdn:             &gcpcloudcdnv1.GcpCloudCdn{},
	cloudresourcekind.CloudResourceKind_GcpCloudFunction:        &gcpcloudfunctionv1.GcpCloudFunction{},
	cloudresourcekind.CloudResourceKind_GcpCloudRun:             &gcpcloudrunv1.GcpCloudRun{},
	cloudresourcekind.CloudResourceKind_GcpCloudSql:             &gcpcloudsqlv1.GcpCloudSql{},
	cloudresourcekind.CloudResourceKind_GcpDnsZone:              &gcpdnszonev1.GcpDnsZone{},
	cloudresourcekind.CloudResourceKind_GcpGcsBucket:            &gcpgcsbucketv1.GcpGcsBucket{},
	cloudresourcekind.CloudResourceKind_GcpGkeAddonBundle:       &gcpgkeaddonbundlev1.GcpGkeAddonBundle{},
	cloudresourcekind.CloudResourceKind_GcpGkeCluster:           &gcpgkeclusterv1.GcpGkeCluster{},
	cloudresourcekind.CloudResourceKind_GcpSecretsManager:       &gcpsecretsmanagerv1.GcpSecretsManager{},
	cloudresourcekind.CloudResourceKind_GcpStaticWebsite:        &gcpstaticwebsitev1.GcpStaticWebsite{},
}

var providerKubernetesMap = map[cloudresourcekind.CloudResourceKind]proto.Message{
	cloudresourcekind.CloudResourceKind_ArgocdKubernetes:         &argocdkubernetesv1.ArgocdKubernetes{},
	cloudresourcekind.CloudResourceKind_CronJobKubernetes:        &cronjobkubernetesv1.CronJobKubernetes{},
	cloudresourcekind.CloudResourceKind_ElasticsearchKubernetes:  &elasticsearchkubernetesv1.ElasticsearchKubernetes{},
	cloudresourcekind.CloudResourceKind_GitlabKubernetes:         &gitlabkubernetesv1.GitlabKubernetes{},
	cloudresourcekind.CloudResourceKind_GrafanaKubernetes:        &grafanakubernetesv1.GrafanaKubernetes{},
	cloudresourcekind.CloudResourceKind_HelmRelease:              &helmreleasev1.HelmRelease{},
	cloudresourcekind.CloudResourceKind_JenkinsKubernetes:        &jenkinskubernetesv1.JenkinsKubernetes{},
	cloudresourcekind.CloudResourceKind_KafkaKubernetes:          &kafkakubernetesv1.KafkaKubernetes{},
	cloudresourcekind.CloudResourceKind_KeycloakKubernetes:       &keycloakkubernetesv1.KeycloakKubernetes{},
	cloudresourcekind.CloudResourceKind_KubernetesHttpEndpoint:   &kuberneteshttpendpointv1.KubernetesHttpEndpoint{},
	cloudresourcekind.CloudResourceKind_LocustKubernetes:         &locustkubernetesv1.LocustKubernetes{},
	cloudresourcekind.CloudResourceKind_MicroserviceKubernetes:   &microservicekubernetesv1.MicroserviceKubernetes{},
	cloudresourcekind.CloudResourceKind_MongodbKubernetes:        &mongodbkubernetesv1.MongodbKubernetes{},
	cloudresourcekind.CloudResourceKind_Neo4jKubernetes:          &neo4jkubernetesv1.Neo4JKubernetes{},
	cloudresourcekind.CloudResourceKind_OpenFgaKubernetes:        &openfgakubernetesv1.OpenFgaKubernetes{},
	cloudresourcekind.CloudResourceKind_PostgresKubernetes:       &postgreskubernetesv1.PostgresKubernetes{},
	cloudresourcekind.CloudResourceKind_PrometheusKubernetes:     &prometheuskubernetesv1.PrometheusKubernetes{},
	cloudresourcekind.CloudResourceKind_RedisKubernetes:          &rediskubernetesv1.RedisKubernetes{},
	cloudresourcekind.CloudResourceKind_SignozKubernetes:         &signozkubernetesv1.SignozKubernetes{},
	cloudresourcekind.CloudResourceKind_SolrKubernetes:           &solrkubernetesv1.SolrKubernetes{},
	cloudresourcekind.CloudResourceKind_StackJobRunnerKubernetes: &stackjobrunnerkubernetesv1.StackJobRunnerKubernetes{},
}
