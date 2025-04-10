package awsiamrolev1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsIamRole(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsIamRole Suite")
}

var _ = Describe("AwsIamRole Custom Validation Tests", func() {
	Describe("When valid input is passed", func() {
		Context("aws_iam_role", func() {
			It("should not return a validation error", func() {
				input := &AwsIamRole{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsIamRole",
					Metadata: &shared.ApiResourceMetadata{
						Name: "valid-name",
					},
					Spec: &AwsIamRoleSpec{
						TrustPolicy: &structpb.Struct{},
						ManagedPolicyArns: []string{
							"arn:aws:iam::123456789012:policy/testPolicy",
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
