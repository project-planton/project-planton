package azurecredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAzureCredential(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AzureCredential Suite")
}

var (
	_ = ginkgo.Describe("AzureCredentialSpec Custom Validation Tests", func() {
		var input *AzureCredential

		ginkgo.BeforeEach(func() {
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

		ginkgo.Describe("When valid input is passed", func() {
			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
)
