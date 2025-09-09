package gcpprojectv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpProjectSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpProjectSpec Custom Validation Tests")
}

var _ = Describe("GcpProjectSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_project", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpProject{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpProject",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-gcp-project",
					},
					Spec: &GcpProjectSpec{
						BillingAccountId: "0123AB-4567CD-89EFGH", // Valid billing account format
						OwnerMember:      "user@example.com",     // Valid email address
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
