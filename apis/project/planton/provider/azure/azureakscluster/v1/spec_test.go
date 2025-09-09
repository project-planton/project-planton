package azureaksclusterv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureAksClusterSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureAksClusterSpec Custom Validation Tests")
}

var _ = Describe("AzureAksClusterSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("azure_aks_cluster", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AzureAksCluster{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureAksCluster",
					Metadata: &shared.ApiResourceMetadata{
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
				Expect(err).To(BeNil())
			})
		})
	})
})
