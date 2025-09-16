package gcpcloudsqlv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudSqlSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "GcpCloudSqlSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("GcpCloudSqlSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_cloud_sql", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &GcpCloudSql{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudSql",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloud-sql",
					},
					Spec: &GcpCloudSqlSpec{
						// No fields currently, spec to be added
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
