package gcpdnszonev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpDnsZoneSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpDnsZoneSpec Custom Validation Tests")
}

var _ = Describe("GcpDnsZoneSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_dns_zone", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpDnsZone{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpDnsZone",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-dns-zone",
					},
					Spec: &GcpDnsZoneSpec{
						ProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
