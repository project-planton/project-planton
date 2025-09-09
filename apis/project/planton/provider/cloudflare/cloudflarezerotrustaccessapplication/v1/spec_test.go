package cloudflarezerotrustaccessapplicationv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestCloudflareZeroTrustAccessApplicationSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CloudflareZeroTrustAccessApplicationSpec Custom Validation Tests")
}

var _ = Describe("CloudflareZeroTrustAccessApplicationSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("cloudflare_zero_trust_access_application", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &CloudflareZeroTrustAccessApplication{
					ApiVersion: "cloudflare.project-planton.org/v1",
					Kind:       "CloudflareZeroTrustAccessApplication",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-access-app",
					},
					Spec: &CloudflareZeroTrustAccessApplicationSpec{
						ApplicationName: "Test Access Application",
						ZoneId:          "test-zone-123",
						Hostname:        "app.example.com",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
