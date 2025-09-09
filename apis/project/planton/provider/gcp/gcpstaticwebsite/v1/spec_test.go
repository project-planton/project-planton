package gcpstaticwebsitev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpStaticWebsiteSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpStaticWebsiteSpec Custom Validation Tests")
}

var _ = Describe("GcpStaticWebsiteSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_static_website", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpStaticWebsite{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpStaticWebsite",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-static-website",
					},
					Spec: &GcpStaticWebsiteSpec{
						GcpProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
