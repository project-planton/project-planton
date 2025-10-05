package civocredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	civo "github.com/project-planton/project-planton/apis/project/planton/provider/civo"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCivoCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "CivoCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("CivoCredentialSpec Validation Tests", func() {
	var input *CivoCredential

	ginkgo.BeforeEach(func() {
		input = &CivoCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "CivoCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-civo-cred",
			},
			Spec: &CivoCredentialSpec{
				ApiToken:      "valid-api-token-for-civo",
				DefaultRegion: civo.CivoRegion_lon1,
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with valid credentials", func() {

			ginkgo.It("should not return a validation error for minimal fields", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with object store credentials", func() {
				input.Spec.ObjectStoreAccessId = "object-store-access-id"
				input.Spec.ObjectStoreSecretKey = "object-store-secret-key"
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("When invalid input is passed", func() {
		ginkgo.Context("missing required fields", func() {

			ginkgo.It("should return error if api_token is missing", func() {
				input.Spec.ApiToken = ""
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})

			ginkgo.It("should return error if default_region is missing", func() {
				input.Spec.DefaultRegion = civo.CivoRegion_civo_region_unspecified
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
