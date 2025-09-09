package digitaloceandnszonev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestDigitalOceanDnsZoneSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DigitalOceanDnsZoneSpec Custom Validation Tests")
}

var _ = Describe("DigitalOceanDnsZoneSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("digitalocean_dns_zone", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &DigitalOceanDnsZone{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanDnsZone",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-dns-zone",
					},
					Spec: &DigitalOceanDnsZoneSpec{
						DomainName: "example.com",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
