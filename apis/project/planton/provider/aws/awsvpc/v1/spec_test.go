package awsvpcv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsVpcSpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsVpcSpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AwsVpcSpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_vpc", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AwsVpc{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsVpc",
					Metadata: &shared.CloudResourceMetadata{
						Name: "test-vpc",
					},
					Spec: &AwsVpcSpec{
						VpcCidr:                    "10.0.0.0/16",
						SubnetsPerAvailabilityZone: 1,
						SubnetSize:                 256,
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
