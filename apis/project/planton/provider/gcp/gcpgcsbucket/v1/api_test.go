package gcpgcsbucketv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGcsBucket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGcsBucket Suite")
}

var _ = Describe("GcpGcsBucket Custom Validation Tests", func() {
	var input *GcpGcsBucket

	BeforeEach(func() {
		input = &GcpGcsBucket{
			ApiVersion: "gcp.project-planton.org/v1",
			Kind:       "GcpGcsBucket",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-bucket",
			},
			Spec: &GcpGcsBucketSpec{
				GcpProjectId: "some-project-id",
				GcpRegion:    "us-west2",
				IsPublic:     false,
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gcp_gcs_bucket", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
