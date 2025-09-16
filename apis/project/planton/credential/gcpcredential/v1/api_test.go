package gcpcredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCredential(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpCredential Suite")
}

var _ = ginkgo.Describe("GcpCredentialSpec Custom Validation Tests", func() {
	var input *GcpCredential

	ginkgo.BeforeEach(func() {
		input = &GcpCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "GcpCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-gcp-credential",
			},
			Spec: &GcpCredentialSpec{},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("GCP Credential", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ServiceAccountKeyBase64 = "base64EncodedServiceAccountKey"
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
