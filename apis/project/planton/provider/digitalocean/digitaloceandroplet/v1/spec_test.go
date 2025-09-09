package digitaloceandropletv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestDigitalOceanDropletSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DigitalOceanDropletSpec Custom Validation Tests")
}

var _ = Describe("DigitalOceanDropletSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("digitalocean_droplet", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &DigitalOceanDroplet{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanDroplet",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-droplet",
					},
					Spec: &DigitalOceanDropletSpec{
						DropletName: "test-droplet",
						Region:      digitalocean.DigitalOceanRegion_nyc3,
						Size:        "s-2vcpu-4gb",
						Image:       "ubuntu-22-04-x64",
						Vpc: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-vpc-id"},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
