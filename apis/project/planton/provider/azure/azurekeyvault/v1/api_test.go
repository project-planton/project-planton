package azurekeyvaultv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureKeyVault(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureKeyVault Suite")
}

var _ = Describe("AzureKeyVault Custom Validation Tests", func() {
	var input *AzureKeyVault

	BeforeEach(func() {
		input = &AzureKeyVault{
			ApiVersion: "azure.project-planton.org/v1",
			Kind:       "AzureKeyVault",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-kv",
			},
			Spec: &AzureKeyVaultSpec{
				SecretNames: []string{"my-secret-1", "my-secret-2"},
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("azure_key_vault", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
