package gcpgcsbucketv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestGcpGcsBucketSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpGcsBucketSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpGcsBucketSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_gcs_bucket", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpGcsBucket{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpGcsBucket",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-gcs-bucket",
					},
					Spec: &GcpGcsBucketSpec{
						GcpProjectId: "test-project-123",
						GcpRegion:    "us-central1",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
