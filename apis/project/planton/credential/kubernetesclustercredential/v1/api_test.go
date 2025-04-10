package kubernetesclustercredentialv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestKubernetesClusterCredential(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KubernetesClusterCredential Suite")
}

var _ = Describe("KubernetesClusterCredentialSpec Custom Validation Tests", func() {
	var input *KubernetesClusterCredential

	BeforeEach(func() {
		input = &KubernetesClusterCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "KubernetesClusterCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-kube-cred",
			},
			Spec: &KubernetesClusterCredentialSpec{},
		}
	})

	Describe("When valid input is passed", func() {
		Context("gcp_gke", func() {
			It("should not return a validation error", func() {
				input.Spec.Provider = KubernetesProvider_gcp_gke
				input.Spec.GcpGke = &KubernetesClusterCredentialGcpGke{
					ClusterEndpoint:         "https://example-gke-endpoint.com",
					ClusterCaData:           "base64CACertificateData",
					ServiceAccountKeyBase64: "base64ServiceAccountKeyData",
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("aws_eks", func() {
			It("should not return a validation error", func() {
				input.Spec.Provider = KubernetesProvider_aws_eks
				input.Spec.AwsEks = &KubernetesClusterCredentialAwsEks{}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("azure_aks", func() {
			It("should not return a validation error", func() {
				input.Spec.Provider = KubernetesProvider_azure_aks
				input.Spec.AzureAks = &KubernetesClusterCredentialAzureAks{}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
