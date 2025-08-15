package awsvpcv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsVpc(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsVpc Suite")
}

var _ = Describe("AwsVpc Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("aws_vpc", func() {
			var input *AwsVpc

			BeforeEach(func() {
				input = &AwsVpc{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsVpc",
					Metadata: &shared.ApiResourceMetadata{
						Name: "example-vpc",
					},
					Spec: &AwsVpcSpec{
						VpcCidr:                    "10.0.0.0/16",
						AvailabilityZones:          []string{"us-west-2a", "us-west-2b"},
						SubnetsPerAvailabilityZone: 1,
						SubnetSize:                 1,
						IsNatGatewayEnabled:        true,
						IsDnsHostnamesEnabled:      true,
						IsDnsSupportEnabled:        true,
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
