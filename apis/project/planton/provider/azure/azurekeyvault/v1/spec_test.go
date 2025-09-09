package azurekeyvaultv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureKeyVaultSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureKeyVaultSpec Custom Validation Tests")
}

var _ = Describe("AzureKeyVaultSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("azure_key_vault", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AzureKeyVault{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureKeyVault",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-key-vault",
					},
					Spec: &AzureKeyVaultSpec{
						// No required fields
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
