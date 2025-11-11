package gcpcloudfunctionv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestGcpCloudFunctionSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpCloudFunctionSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpCloudFunctionSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_cloud_function", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpCloudFunction{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudFunction",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-cloud-function",
					},
					Spec: &GcpCloudFunctionSpec{
						GcpProjectId: "test-project-123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
