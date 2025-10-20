package azureaksclusterv1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureAksClusterSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AzureAksClusterSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AzureAksClusterSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("azure_aks_cluster", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AzureAksCluster{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureAksCluster",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-aks-cluster",
					},
					Spec: &AzureAksClusterSpec{
						Region: "eastus",
						VnetSubnetId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "/subscriptions/sub-123/resourceGroups/rg/providers/Microsoft.Network/virtualNetworks/vnet/subnets/subnet"},
						},
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
