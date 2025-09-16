package confluentcredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestConfluentCredential(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ConfluentCredential Suite")
}

var _ = ginkgo.Describe("KubernetesClusterCredentialSpec Custom Validation Tests", func() {
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
		ginkgo.Context("confluent_credential", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid domain-specific constraints are passed", func() {
		ginkgo.Context("api_version mismatch", func() {
			ginkgo.It("should return a validation error if api_version is incorrect", func() {
				input.ApiVersion = "invalid.version"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})

		ginkgo.Context("kind mismatch", func() {
			ginkgo.It("should return a validation error if kind is incorrect", func() {
				input.Kind = "NotConfluentCredential"
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
