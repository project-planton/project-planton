package gcpsecretsmanagerv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestGcpSecretsManagerSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpSecretsManagerSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpSecretsManagerSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_secrets_manager", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with multiple secrets", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "app-secrets",
						Org:  "acme-corp",
						Env:  "production",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "my-gcp-project",
						SecretNames: []string{
							"database-password",
							"api-key",
							"oauth-client-secret",
							"jwt-signing-key",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with empty secret names", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "empty-secrets",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId:   "test-project-456",
						SecretNames: []string{},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with environment metadata", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "prod-secrets",
						Env:  "prod",
						Org:  "engineering",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "prod-project-123",
						SecretNames: []string{
							"stripe-secret-key",
							"sendgrid-api-key",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("gcp_secrets_manager", func() {

			ginkgo.It("should return a validation error when project_id is missing", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "",
						SecretNames: []string{
							"api-key",
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
				gomega.Expect(err.Error()).To(gomega.ContainSubstring("project_id"))
			})

			ginkgo.It("should return a validation error when project_id is empty string", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
				gomega.Expect(err.Error()).To(gomega.ContainSubstring("project_id"))
			})

			ginkgo.It("should return a validation error when spec is nil", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: nil,
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
				gomega.Expect(err.Error()).To(gomega.ContainSubstring("spec"))
			})

			ginkgo.It("should return a validation error when metadata is nil", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata:   nil,
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
				gomega.Expect(err.Error()).To(gomega.ContainSubstring("metadata"))
			})

			ginkgo.It("should return a validation error when api_version is incorrect", func() {
				input := &GcpSecretsManager{
					ApiVersion: "invalid.api.version/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
				gomega.Expect(err.Error()).To(gomega.ContainSubstring("api_version"))
			})

			ginkgo.It("should return a validation error when kind is incorrect", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "InvalidKind",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).ToNot(gomega.BeNil())
				gomega.Expect(err.Error()).To(gomega.ContainSubstring("kind"))
			})
		})
	})
})
