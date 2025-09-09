package gcpsubnetworkv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpSubnetworkSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpSubnetworkSpec Custom Validation Tests")
}

var _ = Describe("GcpSubnetworkSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_subnetwork", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpSubnetwork{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpSubnetwork",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-subnetwork",
					},
					Spec: &GcpSubnetworkSpec{
						ProjectId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-project-123"},
						},
						VpcSelfLink: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "projects/test-project-123/global/networks/test-vpc"},
						},
						Region:       "us-central1",
						IpCidrRange: "10.0.0.0/24",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
