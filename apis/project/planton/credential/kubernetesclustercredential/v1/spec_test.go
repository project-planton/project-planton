package kubernetesclustercredentialv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestKubernetesClusterCredentialSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "KubernetesClusterCredentialSpec Validation Tests")
}

var _ = ginkgo.Describe("KubernetesClusterCredentialSpec Validation Tests", func() {
	var input *KubernetesClusterCredential

	ginkgo.BeforeEach(func() {
		input = &KubernetesClusterCredential{
			ApiVersion: "credential.project-planton.org/v1",
			Kind:       "KubernetesClusterCredential",
			Metadata: &shared.ApiResourceMetadata{
				Name: "my-kube-cred",
			},
			Spec: &KubernetesClusterCredentialSpec{},
		}
	})

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("gcp_gke", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.Provider = KubernetesProvider_gcp_gke
				input.Spec.GcpGke = &KubernetesClusterCredentialGcpGke{
					ClusterEndpoint:         "https://example-gke-endpoint.com",
					ClusterCaData:           "base64CACertificateData",
					ServiceAccountKeyBase64: "base64ServiceAccountKeyData",
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("aws_eks", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.Provider = KubernetesProvider_aws_eks
				input.Spec.AwsEks = &KubernetesClusterCredentialAwsEks{}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("azure_aks", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.Provider = KubernetesProvider_azure_aks
				input.Spec.AzureAks = &KubernetesClusterCredentialAzureAks{}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
