package gcpcredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("GcpCredentialSpec Validation Tests", func() {
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
		ginkgo.Context("with valid service account key", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.ServiceAccountKeyBase64 = "base64EncodedServiceAccountKey"
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("missing required fields", func() {

			ginkgo.It("should return error if service_account_key_base64 is missing", func() {
				input.Spec.ServiceAccountKeyBase64 = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
