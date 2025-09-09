package awssecuritygroupv1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsSecurityGroupSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsSecurityGroupSpec Custom Validation Tests")
}

var _ = Describe("AwsSecurityGroupSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws_security_group", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AwsSecurityGroup{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsSecurityGroup",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-security-group",
					},
					Spec: &AwsSecurityGroupSpec{
						VpcId: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "vpc-12345678"},
						},
						Description: "Test security group for validation",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
