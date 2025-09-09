package gcpgkenodepoolv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestGcpGkeNodePoolSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GcpGkeNodePoolSpec Custom Validation Tests")
}

var _ = Describe("GcpGkeNodePoolSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("gcp_gke_node_pool", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &GcpGkeNodePool{
					ApiVersion: "gcp.project-planton.org/v1",
					Kind:       "GcpGkeNodePool",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-node-pool",
					},
					Spec: &GcpGkeNodePoolSpec{
						ClusterProjectId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-project-123"},
						},
						ClusterName: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-gke-cluster"},
						},
						DiskType: "pd-standard", // Valid disk type
						NodePoolSize: &GcpGkeNodePoolSpec_NodeCount{
							NodeCount: 3, // Required node pool size
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
