package temporalkubernetesv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
)

func TestTemporalKubernetes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TemporalKubernetes Suite")
}

var _ = Describe("TemporalKubernetes Custom Validation Tests", func() {
	var input *TemporalKubernetes

	BeforeEach(func() {
		input = &TemporalKubernetes{
			ApiVersion: "kubernetes.project-planton.org/v1",
			Kind:       "TemporalKubernetes",
			Metadata: &shared.ApiResourceMetadata{
				Name: "temporal-demo",
			},
			Spec: &TemporalKubernetesSpec{
				DisableWebUi: false,
				Ingress: &kubernetes.IngressSpec{
					DnsDomain: "temporal.example.com",
				},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("with cassandra external database and external elasticsearch", func() {
			BeforeEach(func() {
				input.Spec.Database = &TemporalKubernetesDatabaseConfig{
					Backend: TemporalKubernetesDatabaseBackend_cassandra,
					ExternalDatabase: &TemporalKubernetesExternalDatabase{
						Host:     "cassandra.example.com",
						Port:     9042,
						User:     "temporal_user",
						Password: "secret",
					},
					DatabaseName:           "temporal",
					VisibilityName:         "temporal_visibility",
					DisableAutoSchemaSetup: true,
				}
				input.Spec.ExternalElasticsearch = &TemporalKubernetesExternalElasticsearch{
					Host:     "es.example.com",
					Port:     9200,
					User:     "es_user",
					Password: "es_password",
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("with postgresql external database", func() {
			BeforeEach(func() {
				input.Spec.Database = &TemporalKubernetesDatabaseConfig{
					Backend: TemporalKubernetesDatabaseBackend_postgresql,
					ExternalDatabase: &TemporalKubernetesExternalDatabase{
						Host:     "postgres.example.com",
						Port:     5432,
						User:     "pg_user",
						Password: "pg_password",
					},
					DatabaseName:           "temporal_pg",
					VisibilityName:         "temporal_visibility_pg",
					DisableAutoSchemaSetup: false,
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("with mysql external database", func() {
			BeforeEach(func() {
				input.Spec.Database = &TemporalKubernetesDatabaseConfig{
					Backend: TemporalKubernetesDatabaseBackend_mysql,
					ExternalDatabase: &TemporalKubernetesExternalDatabase{
						Host:     "mysql.example.com",
						Port:     3306,
						User:     "mysql_user",
						Password: "mysql_password",
					},
					DatabaseName:   "temporal_mysql",
					VisibilityName: "temporal_visibility_mysql",
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
