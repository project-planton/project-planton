package gcpcredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCredential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCredential Suite")
}

var _ = Describe("GcpCredentialSpec Custom Validation Tests", func() {
	var input *GcpCredential

	BeforeEach(func() {
		input = &GcpCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "GcpCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-gcp-credential",
			},
			Spec: &GcpCredentialSpec{},
		}
	})

	Describe("When valid input is passed", func() {
		Context("GCP Credential", func() {
			It("should not return a validation error", func() {
				input.Spec.ServiceAccountKeyBase64 = "base64EncodedServiceAccountKey"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
