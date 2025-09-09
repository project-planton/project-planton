package gcpstaticwebsitev1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpStaticWebsite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpStaticWebsite Suite")
}

var _ = Describe("GcpStaticWebsite Custom Validation Tests", func() {
	var input *GcpStaticWebsite

	BeforeEach(func() {
		input = &GcpStaticWebsite{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpStaticWebsite",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-static-website",
			},
			Spec: &GcpStaticWebsiteSpec{
				GcpProjectId: "my-gcp-project-id",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gcp_static_website", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
