package gcpsecretsmanagerv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpSecretsManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpSecretsManager Suite")
}

var _ = Describe("GcpSecretsManager Custom Validation Tests", func() {
	var input *GcpSecretsManager

	BeforeEach(func() {
		input = &GcpSecretsManager{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpSecretsManager",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-secrets",
			},
			Spec: &GcpSecretsManagerSpec{
				ProjectId:   "some-project-id",
				SecretNames: []string{"api-key", "db-password"},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gcp_secrets_manager", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
