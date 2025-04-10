package awscloudfrontv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestKubernetesClusterCredentialSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "KubernetesClusterCredentialSpec Custom Validation Tests")
}

var _ = Describe("KubernetesClusterCredentialSpec Custom Validation Tests", func() {

	var input *AwsCloudFront

	BeforeEach(func() {
		input = &AwsCloudFront{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsCloudFront",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-cloud-front",
			},
			Spec: &AwsCloudFrontSpec{},
		}
	})

	Describe("When valid input is passed", func() {

		Context("gcp_gke", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("aws_eks", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("azure_aks", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
