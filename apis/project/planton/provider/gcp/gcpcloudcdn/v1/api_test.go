package gcpcloudcdnv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudCdn(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudCdn Suite")
}

var _ = Describe("GcpCloudCdn Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("GCP", func() {
			It("should not return a validation error", func() {
				input := &GcpCloudCdn{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudCdn",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloudcdn",
					},
					Spec: &GcpCloudCdnSpec{
						GcpProjectId: "my-project",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
