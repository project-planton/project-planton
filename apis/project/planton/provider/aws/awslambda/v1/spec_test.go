package awslambdav1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsLambdaSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsLambdaSpec Custom Validation Tests")
}

var _ = Describe("AwsLambdaSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws_lambda", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AwsLambda{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsLambda",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-lambda-function",
					},
					Spec: &AwsLambdaSpec{
						FunctionName: "test-function",
						RoleArn: &foreignkeyv1.StringValueOrRef{
							LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "arn:aws:iam::123456789012:role/lambda-execution-role"},
						},
						CodeSourceType: CodeSourceType_CODE_SOURCE_TYPE_S3,
						S3: &S3Code{
							Bucket: "my-lambda-bucket",
							Key:    "functions/test-function.zip",
						},
						Runtime: "python3.11",
						Handler: "main.handler",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
