package kubernetestemporalv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project_planton/shared/foreignkey/v1"
	"google.golang.org/protobuf/proto"
)

func TestKubernetesTemporal(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesTemporal Suite")
}

var _ = ginkgo.Describe("KubernetesTemporal Custom Validation Tests", func() {
	var input *KubernetesTemporal

	ginkgo.BeforeEach(func() {
		input = &KubernetesTemporal{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "KubernetesTemporal",
			Metadata: &shared.CloudResourceMetadata{
				Name: "temporal-demo",
			},
			Spec: &KubernetesTemporalSpec{
				TargetCluster: &kubernetes.KubernetesClusterSelector{
					ClusterKind: cloudresourcekind.CloudResourceKind_GcpGkeCluster,
					ClusterName: "test-cluster",
				},
				Namespace: &foreignkeyv1.StringValueOrRef{
					LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
						Value: "test-namespace",
					},
				},
				DisableWebUi: false,
				Ingress: &KubernetesTemporalIngress{
					Frontend: &KubernetesTemporalFrontendIngressEndpoint{
						Enabled:      true,
						GrpcHostname: "temporal-frontend.example.com",
					},
					WebUi: &KubernetesTemporalWebUiIngressEndpoint{
						Enabled:  true,
						Hostname: "temporal-ui.example.com",
					},
				},
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with cassandra external database and external elasticsearch", func() {
			ginkgo.BeforeEach(func() {
				input.Spec.Database = &KubernetesTemporalDatabaseConfig{
					Backend: KubernetesTemporalDatabaseBackend_cassandra,
					ExternalDatabase: &KubernetesTemporalExternalDatabase{
						Host:     "cassandra.example.com",
						Port:     9042,
						Username: "temporal_user",
						Password: "secret",
					},
					DatabaseName:           proto.String("temporal"),
					VisibilityName:         proto.String("temporal_visibility"),
					DisableAutoSchemaSetup: true,
				}
				input.Spec.ExternalElasticsearch = &KubernetesTemporalExternalElasticsearch{
					Host:     "es.example.com",
					Port:     9200,
					User:     "es_user",
					Password: "es_password",
				}
			})

			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with postgresql external database", func() {
			ginkgo.BeforeEach(func() {
				input.Spec.Database = &KubernetesTemporalDatabaseConfig{
					Backend: KubernetesTemporalDatabaseBackend_postgresql,
					ExternalDatabase: &KubernetesTemporalExternalDatabase{
						Host:     "postgres.example.com",
						Port:     5432,
						Username: "pg_user",
						Password: "pg_password",
					},
					DatabaseName:           proto.String("temporal_pg"),
					VisibilityName:         proto.String("temporal_visibility_pg"),
					DisableAutoSchemaSetup: false,
				}
			})

			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with mysql external database", func() {
			ginkgo.BeforeEach(func() {
				input.Spec.Database = &KubernetesTemporalDatabaseConfig{
					Backend: KubernetesTemporalDatabaseBackend_mysql,
					ExternalDatabase: &KubernetesTemporalExternalDatabase{
						Host:     "mysql.example.com",
						Port:     3306,
						Username: "mysql_user",
						Password: "mysql_password",
					},
					DatabaseName:   proto.String("temporal_mysql"),
					VisibilityName: proto.String("temporal_visibility_mysql"),
				}
			})

			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("Search Attribute Validation Tests", func() {
		ginkgo.BeforeEach(func() {
			// Set up a valid base configuration with database
			input.Spec.Database = &KubernetesTemporalDatabaseConfig{
				Backend: KubernetesTemporalDatabaseBackend_postgresql,
				ExternalDatabase: &KubernetesTemporalExternalDatabase{
					Host:     "postgres.example.com",
					Port:     5432,
					Username: "pg_user",
					Password: "pg_password",
				},
				DatabaseName:   proto.String("temporal_pg"),
				VisibilityName: proto.String("temporal_visibility_pg"),
			}
		})
	})
})
