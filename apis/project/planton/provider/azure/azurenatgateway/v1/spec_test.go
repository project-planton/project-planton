package azurenatgatewayv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureNatGatewaySpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureNatGatewaySpec Custom Validation Tests")
}

var _ = Describe("AzureNatGatewaySpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("azure_nat_gateway", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AzureNatGateway{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureNatGateway",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-nat-gateway",
					},
					Spec: &AzureNatGatewaySpec{
						SubnetId: &foreignkeyv1.StringValueOrRef{
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
