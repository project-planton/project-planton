package digitaloceanvpcv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestDigitalOceanVpcSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DigitalOceanVpcSpec Custom Validation Tests")
}

var _ = Describe("DigitalOceanVpcSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("digitalocean_vpc", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &DigitalOceanVpc{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanVpc",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-vpc",
					},
					Spec: &DigitalOceanVpcSpec{
						Region:      digitalocean.DigitalOceanRegion_nyc3,
						IpRangeCidr: "10.10.0.0/16",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
