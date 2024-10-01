package pulumimodule

type CloneUrl string

var DefaultGitRepoMap = merge(
	atlasDefaultGitRepoMap,
	awsDefaultGitRepoMap,
	confluentDefaultGitRepoMap,
	gcpDefaultGitRepoMap,
	kubernetesDefaultGitRepoMap,
	snowflakeDefaultGitRepoMap,
)

func merge(items ...map[string]CloneUrl) map[string]CloneUrl {
	resp := make(map[string]CloneUrl)
	for _, i := range items {
		for k, v := range i {
			resp[k] = v
		}
	}
	return resp
}

var atlasDefaultGitRepoMap = map[string]CloneUrl{
	"mongodb-atlas": "https://github.com/plantoncloud/mongodb-atlas-pulumi-module.git",
}

var awsDefaultGitRepoMap = map[string]CloneUrl{
	"aws-cloud-front":     "https://github.com/plantoncloud/aws-cloud-front-pulumi-module.git",
	"aws-dynamodb":        "https://github.com/plantoncloud/aws-dynamodb-pulumi-module.git",
	"aws-fargate":         "https://github.com/plantoncloud/aws-fargate-pulumi-module.git",
	"aws-lambda":          "https://github.com/plantoncloud/aws-lambda-pulumi-module.git",
	"aws-rds-cluster":     "https://github.com/plantoncloud/aws-rds-cluster-pulumi-module.git",
	"aws-rds-instance":    "https://github.com/plantoncloud/aws-rds-instance-pulumi-module.git",
	"aws-secrets-manager": "https://github.com/plantoncloud/aws-secrets-manager-pulumi-module.git",
	"aws-static-website":  "https://github.com/plantoncloud/aws-static-website-pulumi-module.git",
	"aws-vpc":             "https://github.com/plantoncloud/aws-vpc-pulumi-module.git",
	"eks-cluster":         "https://github.com/plantoncloud/eks-cluster-pulumi-module.git",
	"route53-zone":        "https://github.com/plantoncloud/route53-zone-pulumi-module.git",
}

var confluentDefaultGitRepoMap = map[string]CloneUrl{
	"confluent-kafka": "https://github.com/plantoncloud/confluent-kafka-pulumi-module.git",
}

var gcpDefaultGitRepoMap = map[string]CloneUrl{
	"gcp-artifact-registry": "https://github.com/plantoncloud/gcp-artifact-registry-pulumi-module.git",
	"gcp-cloud-cdn":         "https://github.com/plantoncloud/gcp-cloud-cdn-pulumi-module.git",
	"gcp-cloud-function":    "https://github.com/plantoncloud/gcp-cloud-function-pulumi-module.git",
	"gcp-cloud-run":         "https://github.com/plantoncloud/gcp-cloud-run-pulumi-module.git",
	"gcp-cloud-sql":         "https://github.com/plantoncloud/gcp-cloud-sql-pulumi-module.git",
	"gcp-dns-zone":          "https://github.com/plantoncloud/gcp-dns-zone-pulumi-module.git",
	"gcp-secrets-manager":   "https://github.com/plantoncloud/gcp-secrets-manager-pulumi-module.git",
	"gcp-static-website":    "https://github.com/plantoncloud/gcp-static-website-pulumi-module.git",
	"gcs-bucket":            "https://github.com/plantoncloud/gcs-bucket-pulumi-module.git",
	"gke-cluster":           "https://github.com/plantoncloud/gke-cluster-pulumi-module.git",
}

var kubernetesDefaultGitRepoMap = map[string]CloneUrl{
	"argocd-kubernetes":           "https://github.com/plantoncloud/argocd-kubernetes-pulumi-module.git",
	"elasticsearch-kubernetes":    "https://github.com/plantoncloud/elasticsearch-kubernetes-pulumi-module.git",
	"gitlab-kubernetes":           "https://github.com/plantoncloud/gitlab-kubernetes-pulumi-module.git",
	"grafana-kubernetes":          "https://github.com/plantoncloud/grafana-kubernetes-pulumi-module.git",
	"helm-release":                "https://github.com/plantoncloud/helm-release-pulumi-module.git",
	"jenkins-kubernetes":          "https://github.com/plantoncloud/jenkins-kubernetes-pulumi-module.git",
	"kafka-kubernetes":            "https://github.com/plantoncloud/kafka-kubernetes-pulumi-module.git",
	"keycloak-kubernetes":         "https://github.com/plantoncloud/keycloak-kubernetes-pulumi-module.git",
	"kubernetes-http-endpoint":    "https://github.com/plantoncloud/kubernetes-http-endpoint-pulumi-module.git",
	"locust-kubernetes":           "https://github.com/plantoncloud/locust-kubernetes-pulumi-module.git",
	"microservice-kubernetes":     "https://github.com/plantoncloud/microservice-kubernetes-pulumi-module.git",
	"mongodb-kubernetes":          "https://github.com/plantoncloud/mongodb-kubernetes-pulumi-module.git",
	"neo4j-kubernetes":            "https://github.com/plantoncloud/neo4j-kubernetes-pulumi-module.git",
	"openfga-kubernetes":          "https://github.com/plantoncloud/openfga-kubernetes-pulumi-module.git",
	"postgres-kubernetes":         "https://github.com/plantoncloud/postgres-kubernetes-pulumi-module.git",
	"prometheus-kubernetes":       "https://github.com/plantoncloud/prometheus-kubernetes-pulumi-module.git",
	"redis-kubernetes":            "https://github.com/plantoncloud/redis-kubernetes-pulumi-module.git",
	"signoz-kubernetes":           "https://github.com/plantoncloud/signoz-kubernetes-pulumi-module.git",
	"solr-kubernetes":             "https://github.com/plantoncloud/solr-kubernetes-pulumi-module.git",
	"stack-job-runner-kubernetes": "https://github.com/plantoncloud/stack-job-runner-kubernetes-pulumi-module.git",
}

var snowflakeDefaultGitRepoMap = map[string]CloneUrl{
	"snowflake-database": "https://github.com/plantoncloud/snowflake-database-pulumi-module.git",
}
