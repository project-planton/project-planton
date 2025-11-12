package azurekeyvaultv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared"
)

func TestAzureKeyVaultSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AzureKeyVaultSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AzureKeyVaultSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("azure_key_vault", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AzureKeyVault{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureKeyVault",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-key-vault",
					},
					Spec: &AzureKeyVaultSpec{
						// No required fields
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
