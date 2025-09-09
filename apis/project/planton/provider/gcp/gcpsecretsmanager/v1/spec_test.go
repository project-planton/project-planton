package gcpsecretsmanagerv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpSecretsManagerSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpSecretsManagerSpec Custom Validation Tests")
}

var _ = Describe("GcpSecretsManagerSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_secrets_manager", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpSecretsManager{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSecretsManager",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-secrets-manager",
					},
					Spec: &GcpSecretsManagerSpec{
						ProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
