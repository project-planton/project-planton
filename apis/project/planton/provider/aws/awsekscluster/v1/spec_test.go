package awseksclusterv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
)

func TestAwsEksClusterSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsEksClusterSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AwsEksClusterSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_eks_cluster", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AwsEksCluster{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsEksCluster",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-eks-cluster",
					},
					Spec: &AwsEksClusterSpec{
						SubnetIds: []*foreignkeyv1.StringValueOrRef{
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-abc123"},
							},
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-def456"},
							},
						},
						ClusterRoleArn: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{
								Value: "arn:aws:iam::123456789012:role/EksClusterServiceRole",
							},
						},
						Version: "1.29", // Valid version that matches the regex
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
