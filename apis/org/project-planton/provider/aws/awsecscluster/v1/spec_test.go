package awsecsclusterv1

import (
	"testing"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared"

	"buf.build/go/protovalidate"
)

func TestAwsEcsCluster(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsEcsCluster Suite")
}

var _ = ginkgo.Describe("KubernetesProviderConfig Custom Validation Tests", func() {
	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_ecs_cluster", func() {
			var input *AwsEcsCluster

			ginkgo.BeforeEach(func() {
				input = &AwsEcsCluster{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsEcsCluster",
					Metadata: &shared.CloudResourceMetadata{
						Name: "valid-name",
					},
					Spec: &AwsEcsClusterSpec{
						EnableContainerInsights: true,
						CapacityProviders:       []string{"FARGATE", "FARGATE_SPOT"},
						EnableExecuteCommand:    false,
					},
				}
			})

			ginkgo.It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
