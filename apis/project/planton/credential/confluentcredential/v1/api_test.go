package confluentcredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestConfluentCredential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConfluentCredential Suite")
}

var _ = Describe("KubernetesClusterCredentialSpec Custom Validation Tests", func() {
	var input *ConfluentCredential

	BeforeEach(func() {
		input = &ConfluentCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "ConfluentCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-confluent-cred",
			},
			Spec: &ConfluentCredentialSpec{
				ApiKey:    "some-api-key",
				ApiSecret: "some-api-secret",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("confluent_credential", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("When invalid domain-specific constraints are passed", func() {
		Context("api_version mismatch", func() {
			It("should return a validation error if api_version is incorrect", func() {
				input.ApiVersion = "invalid.version"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("kind mismatch", func() {
			It("should return a validation error if kind is incorrect", func() {
				input.Kind = "NotConfluentCredential"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
