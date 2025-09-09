package gcpcloudfunctionv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudFunctionSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudFunctionSpec Custom Validation Tests")
}

var _ = Describe("GcpCloudFunctionSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_cloud_function", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpCloudFunction{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudFunction",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloud-function",
					},
					Spec: &GcpCloudFunctionSpec{
						GcpProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
