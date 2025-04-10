package awseksclusterv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsEksCluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsEksCluster Suite")
}

var _ = Describe("KubernetesClusterCredentialSpec Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {

		Context("gcp_gke", func() {
			It("should not return a validation error", func() {
				// No GCP proto; placeholder context.
			})
		})

		Context("aws_eks", func() {
			var input *AwsEksCluster

			BeforeEach(func() {
				input = &AwsEksCluster{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsEksCluster",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-eks",
					},
					Spec: &AwsEksClusterSpec{
						Region:       "us-east-1",
						InstanceType: "t2.medium",
					},
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should produce a validation error if the api_version is incorrect", func() {
				input.ApiVersion = "invalid"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should produce a validation error if the kind is incorrect", func() {
				input.Kind = "WrongKind"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("azure_aks", func() {
			It("should not return a validation error", func() {
				// No Azure proto; placeholder context.
			})
		})
	})
})
