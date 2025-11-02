package temporalkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"google.golang.org/protobuf/proto"
)

func TestTemporalKubernetes(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "TemporalKubernetes Suite")
}

var _ = ginkgo.Describe("TemporalKubernetes Custom Validation Tests", func() {
	var input *TemporalKubernetes

	ginkgo.BeforeEach(func() {
		input = &TemporalKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "TemporalKubernetes",
			Metadata: &shared.CloudResourceMetadata{
				Name: "temporal-demo",
			},
			Spec: &TemporalKubernetesSpec{
				DisableWebUi: false,
				Ingress: &TemporalKubernetesIngress{
					Frontend: &TemporalKubernetesFrontendIngressEndpoint{
						Enabled:      true,
						GrpcHostname: "temporal-frontend.example.com",
					},
					WebUi: &TemporalKubernetesWebUiIngressEndpoint{
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
				input.Spec.Database = &TemporalKubernetesDatabaseConfig{
					Backend: TemporalKubernetesDatabaseBackend_cassandra,
					ExternalDatabase: &TemporalKubernetesExternalDatabase{
						Host:     "cassandra.example.com",
						Port:     9042,
						Username: "temporal_user",
						Password: "secret",
					},
					DatabaseName:           proto.String("temporal"),
					VisibilityName:         proto.String("temporal_visibility"),
					DisableAutoSchemaSetup: true,
				}
				input.Spec.ExternalElasticsearch = &TemporalKubernetesExternalElasticsearch{
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
				input.Spec.Database = &TemporalKubernetesDatabaseConfig{
					Backend: TemporalKubernetesDatabaseBackend_postgresql,
					ExternalDatabase: &TemporalKubernetesExternalDatabase{
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
				input.Spec.Database = &TemporalKubernetesDatabaseConfig{
					Backend: TemporalKubernetesDatabaseBackend_mysql,
					ExternalDatabase: &TemporalKubernetesExternalDatabase{
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
			input.Spec.Database = &TemporalKubernetesDatabaseConfig{
				Backend: TemporalKubernetesDatabaseBackend_postgresql,
				ExternalDatabase: &TemporalKubernetesExternalDatabase{
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
