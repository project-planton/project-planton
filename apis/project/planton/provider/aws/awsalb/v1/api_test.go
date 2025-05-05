package awsalbv1

import (
	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsAlbSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsAlbSpec Custom Validation Tests")
}

var _ = Describe("AwsAlbSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws_alb", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AwsAlb{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsAlb",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-alb-resource",
					},
					Spec: &AwsAlbSpec{
						Subnets: []*foreignkeyv1.StringValueOrRef{
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-12345678"},
							},
							{
								LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "subnet-12345679"},
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			// Removed: SSL/certificate_arn test cases,
			// since we've temporarily removed that validation rule.
		})
	})

	// Removed: The test block that checks for an error if SSL is enabled but certificate ARN is missing,
	// because that validation no longer exists (we'll re-add these tests later).
})
