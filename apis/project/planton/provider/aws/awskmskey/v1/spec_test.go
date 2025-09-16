package awskmskeyv1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsKmsKeySpec(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "AwsKmsKeySpec Custom Validation Tests")
}

var _ = ginkgo.Describe("AwsKmsKeySpec Custom Validation Tests", func() {

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("aws_kms_key", func() {

			ginkgo.It("should not return a validation error for minimal valid fields", func() {
				input := &AwsKmsKey{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsKmsKey",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-kms-key",
					},
					Spec: &AwsKmsKeySpec{
						DeletionWindowDays: 30,               // Valid value between 7-30
						AliasName:          "alias/test-key", // Valid alias format
					},
				}
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})
})
