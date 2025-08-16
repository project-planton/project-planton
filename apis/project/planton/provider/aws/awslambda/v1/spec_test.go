package awslambdav1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsLambda(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsLambda Suite")
}

var _ = Describe("AwsLambda Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("aws_lambda", func() {
			var input *AwsLambda

			BeforeEach(func() {
				input = &AwsLambda{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsLambda",
					Metadata: &shared.ApiResourceMetadata{
						Name: "my-lambda-resource",
					},
					Spec: &AwsLambdaSpec{
						Function: &AwsLambdaFunction{
							Handler: "testHandler",
						},
					},
				}
			})

			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should fail when api_version is incorrect", func() {
				input.ApiVersion = "invalid-value"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})

			It("should fail when kind is incorrect", func() {
				input.Kind = "SomeOtherKind"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
