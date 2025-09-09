package azurecontainerregistryv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureContainerRegistrySpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureContainerRegistrySpec Custom Validation Tests")
}

var _ = Describe("AzureContainerRegistrySpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("azure_container_registry", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AzureContainerRegistry{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureContainerRegistry",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-container-registry",
					},
					Spec: &AzureContainerRegistrySpec{
						RegistryName: "testregistry123",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
