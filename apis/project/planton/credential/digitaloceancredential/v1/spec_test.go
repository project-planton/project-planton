package digitaloceancredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	digitalocean "github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestDigitalOceanCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "DigitalOceanCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("DigitalOceanCredentialSpec Validation Tests", func() {
	var input *DigitalOceanCredential

	ginkgo.BeforeEach(func() {
		input = &DigitalOceanCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "DigitalOceanCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-do-cred",
			},
			Spec: &DigitalOceanCredentialSpec{
				ApiToken:      "valid-digitalocean-api-token",
				DefaultRegion: digitalocean.DigitalOceanRegion_nyc3,
			},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("with valid credentials", func() {

			ginkgo.It("should not return a validation error for minimal fields", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})

			ginkgo.It("should not return a validation error with Spaces credentials", func() {
				input.Spec.SpacesAccessId = "spaces-access-id-value"
				input.Spec.SpacesSecretKey = "spaces-secret-key-value"
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
				input.Spec.DefaultRegion = digitalocean.DigitalOceanRegion_digital_ocean_region_unspecified
				err := protovalidate.Validate(input)
				gomega.Expect(err).NotTo(gomega.BeNil())
			})
		})
	})
})
