package azureaksnodepoolv1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/org/project-planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestAzureAksNodePoolSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AzureAksNodePoolSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AzureAksNodePoolSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("azure_aks_node_pool", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AzureAksNodePool{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureAksNodePool",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-node-pool",
					},
					Spec: &AzureAksNodePoolSpec{
						ClusterName: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "test-aks-cluster"},
						},
						VmSize:            "Standard_D4s_v3",
						InitialNodeCount:  2,
						AvailabilityZones: []string{"1", "2"}, // Add at least 2 zones for HA
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
