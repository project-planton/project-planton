package azurecontainerregistryv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"
)

func TestAzureContainerRegistrySpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AzureContainerRegistrySpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AzureContainerRegistrySpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("azure_container_registry", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AzureContainerRegistry{
					ApiVersion: "azure.project-planton.org/v1",
					Kind:       "AzureContainerRegistry",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-container-registry",
					},
					Spec: &AzureContainerRegistrySpec{
						RegistryName: "testregistry123",
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
