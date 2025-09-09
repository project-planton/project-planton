package gcpgcsbucketv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGcsBucketSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGcsBucketSpec Custom Validation Tests")
}

var _ = Describe("GcpGcsBucketSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_gcs_bucket", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpGcsBucket{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpGcsBucket",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-gcs-bucket",
					},
					Spec: &GcpGcsBucketSpec{
						GcpProjectId: "test-project-123",
						GcpRegion:    "us-central1",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
