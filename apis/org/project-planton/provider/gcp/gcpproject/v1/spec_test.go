package gcpprojectv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestGcpProjectSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpProjectSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpProjectSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_project", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpProject{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpProject",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-gcp-project",
					},
					Spec: &GcpProjectSpec{
						BillingAccountId: "0123AB-4567CD-89EFGH", // Valid billing account format
						OwnerMember:      "user@example.com",     // Valid email address
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
