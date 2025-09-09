package azurednszonev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureDnsZoneSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureDnsZoneSpec Custom Validation Tests")
}

var _ = Describe("AzureDnsZoneSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("azure_dns_zone", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AzureDnsZone{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureDnsZone",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-dns-zone",
					},
					Spec: &AzureDnsZoneSpec{
						ZoneName:      "example.com",
						ResourceGroup: "test-resource-group",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
