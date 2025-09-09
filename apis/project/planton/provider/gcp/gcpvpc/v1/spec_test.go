package gcpvpcv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpVpcSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpVpcSpec Custom Validation Tests")
}

var _ = Describe("GcpVpcSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_vpc", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpVpc{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpVpc",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-gcp-vpc",
					},
					Spec: &GcpVpcSpec{
						ProjectId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-project-123"},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
