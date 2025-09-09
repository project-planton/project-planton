package gcpcloudfunctionv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudFunction(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudFunction Suite")
}

var _ = Describe("GcpCloudFunction Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("GCP", func() {
			It("should not return a validation error", func() {
				input := &GcpCloudFunction{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudFunction",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-function",
					},
					Spec: &GcpCloudFunctionSpec{
						GcpProjectId: "my-gcp-project",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
