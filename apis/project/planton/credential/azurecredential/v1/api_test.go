package azurecredentialv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureCredential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AzureCredential Suite")
}

var _ = Describe("AzureCredentialSpec Custom Validation Tests", func() {
	var input *AzureCredential

	BeforeEach(func() {
		input = &AzureCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "AzureCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-azure-cred",
			},
			Spec: &AzureCredentialSpec{
				ClientId:       "my-client-id",
				ClientSecret:   "my-client-secret",
				TenantId:       "my-tenant-id",
				SubscriptionId: "my-subscription-id",
			},
		}
	})

	Describe("When valid input is passed", func() {
		It("should not return a validation error", func() {
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})
})
