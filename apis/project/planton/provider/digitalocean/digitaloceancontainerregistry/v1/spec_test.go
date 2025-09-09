package digitaloceancontainerregistryv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean"
)

func TestDigitalOceanContainerRegistrySpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DigitalOceanContainerRegistrySpec Custom Validation Tests")
}

var _ = Describe("DigitalOceanContainerRegistrySpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("digitalocean_container_registry", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &DigitalOceanContainerRegistry{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanContainerRegistry",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-container-registry",
					},
					Spec: &DigitalOceanContainerRegistrySpec{
						Name:              "test-registry",
						SubscriptionTier: DigitalOceanContainerRegistryTier_STARTER,
						Region:           digitalocean.DigitalOceanRegion_nyc3,
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
