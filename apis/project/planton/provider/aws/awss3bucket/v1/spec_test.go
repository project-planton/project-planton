package awss3bucketv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsS3Bucket(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsS3Bucket Suite")
}

var _ = Describe("AwsS3Bucket Custom Validation Tests", func() {
	var input *AwsS3Bucket

	BeforeEach(func() {
		input = &AwsS3Bucket{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsS3Bucket",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-s3-bucket",
			},
			Spec: &AwsS3BucketSpec{
				IsPublic:  false,
				AwsRegion: "us-east-1",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("aws_s3", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Context("api_version field validations", func() {
		It("should reject an incorrect apiVersion", func() {
			input.ApiVersion = "invalid-version"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})

	Context("kind field validations", func() {
		It("should reject an incorrect kind", func() {
			input.Kind = "NotAwsS3Bucket"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})
	})
})
