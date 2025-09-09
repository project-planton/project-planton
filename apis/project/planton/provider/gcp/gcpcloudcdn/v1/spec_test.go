package gcpcloudcdnv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudCdnSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudCdnSpec Custom Validation Tests")
}

var _ = Describe("GcpCloudCdnSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_cloud_cdn", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpCloudCdn{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudCdn",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloud-cdn",
					},
					Spec: &GcpCloudCdnSpec{
						GcpProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
