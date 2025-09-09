package cloudflarednszonev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareDnsZoneSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudflareDnsZoneSpec Custom Validation Tests")
}

var _ = Describe("CloudflareDnsZoneSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("cloudflare_dns_zone", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareDnsZone{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareDnsZone",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-dns-zone",
					},
					Spec: &CloudflareDnsZoneSpec{
						ZoneName:  "example.com",
						AccountId: "test-account-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
