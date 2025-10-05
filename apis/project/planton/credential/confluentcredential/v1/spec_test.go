package confluentcredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestConfluentCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ConfluentCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("ConfluentCredentialSpec Validation Tests", func() {
	var input *ConfluentCredential

	ginkgo.BeforeEach(func() {
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with valid credentials", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("missing required fields", func() {

			ginkgo.It("should return error if api_key is missing", func() {
				input.Spec.ApiKey = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if api_secret is missing", func() {
				input.Spec.ApiSecret = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
