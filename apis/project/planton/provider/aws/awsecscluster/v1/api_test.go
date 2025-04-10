package awsecsclusterv1

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAwsEcsCluster(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsEcsCluster Suite")
}

var _ = Describe("KubernetesClusterCredentialSpec Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("aws_ecs_cluster", func() {
			var input *AwsEcsCluster

			BeforeEach(func() {
				input = &AwsEcsCluster{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsEcsCluster",
					Metadata: &shared.ApiResourceMetadata{
						Name: "valid-name",
					},
					Spec: &AwsEcsClusterSpec{
						EnableContainerInsights: true,
						CapacityProviders:       []string{"FARGATE", "FARGATE_SPOT"},
						EnableExecuteCommand:    false,
					},
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
