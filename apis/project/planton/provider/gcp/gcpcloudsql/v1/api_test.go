package gcpcloudsqlv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpCloudSql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpCloudSql Suite")
}

var _ = Describe("GcpCloudSql Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("GCP", func() {
			It("should not return a validation error", func() {
				input := &GcpCloudSql{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpCloudSql",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-cloudsql",
					},
					Spec: &GcpCloudSqlSpec{},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
