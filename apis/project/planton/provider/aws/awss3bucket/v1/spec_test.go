package awss3bucketv1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsS3BucketSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsS3BucketSpec Custom Validation Tests")
}

var _ = Describe("AwsS3BucketSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("aws_s3_bucket", func() {

			It("should not return a validation error for minimal valid fields", func() {
				input := &AwsS3Bucket{
					ApiVersion: "aws.project-planton.org/v1",
					Kind:       "AwsS3Bucket",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-s3-bucket",
					},
					Spec: &AwsS3BucketSpec{
						AwsRegion: "us-east-1",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})