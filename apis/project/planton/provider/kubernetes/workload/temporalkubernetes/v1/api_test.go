package temporalkubernetesv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
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

		ginkgo.Context("with valid search attribute types", func() {
			ginkgo.It("should not return a validation error for Keyword type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "CustomerId",
						Type: "Keyword",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error for Text type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "Description",
						Type: "Text",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error for Int type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "OrderCount",
						Type: "Int",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error for Double type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "Price",
						Type: "Double",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error for Bool type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "IsActive",
						Type: "Bool",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error for Datetime type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "CreatedAt",
						Type: "Datetime",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error for KeywordList type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "Tags",
						Type: "KeywordList",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with multiple valid search attributes", func() {
			ginkgo.It("should not return a validation error for multiple attributes with different types", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "CustomerId",
						Type: "Keyword",
					},
					{
						Name: "Environment",
						Type: "Keyword",
					},
					{
						Name: "OrderCount",
						Type: "Int",
					},
					{
						Name: "TotalAmount",
						Type: "Double",
					},
					{
						Name: "IsProcessed",
						Type: "Bool",
					},
					{
						Name: "ProcessedAt",
						Type: "Datetime",
					},
					{
						Name: "Labels",
						Type: "KeywordList",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("with invalid search attribute types", func() {
			ginkgo.It("should return a validation error for lowercase keyword type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "CustomerId",
						Type: "keyword",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error for lowercase text type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "Description",
						Type: "text",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error for lowercase int type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "OrderCount",
						Type: "int",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error for invalid type value", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "CustomField",
						Type: "InvalidType",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error for empty type value", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "CustomField",
						Type: "",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error for snake_case type", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "Tags",
						Type: "keyword_list",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("with missing required fields", func() {
			ginkgo.It("should return a validation error when type field is missing", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Name: "CustomField",
						// Type field is not set
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return a validation error when name field is missing", func() {
				input.Spec.SearchAttributes = []*TemporalKubernetesSearchAttribute{
					{
						Type: "Keyword",
						// Name field is not set
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
