package gcpcloudrunv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudRun(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudRun Suite")
}

var _ = Describe("GcpCloudRun Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("GCP", func() {
			It("should not return a validation error", func() {
				input := &GcpCloudRun{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudRun",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-run",
					},
					Spec: &GcpCloudRunSpec{
						GcpProjectId: "my-project-id",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
