package gcpstaticwebsitev1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpStaticWebsiteSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpStaticWebsiteSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpStaticWebsiteSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_static_website", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpStaticWebsite{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpStaticWebsite",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-static-website",
					},
					Spec: &GcpStaticWebsiteSpec{
						GcpProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
