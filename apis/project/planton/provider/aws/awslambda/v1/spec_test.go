package awslambdav1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAWSLambdaSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AWSLambdaSpec Validation Suite")
}

var _ = Describe("AWSLambdaSpec validations", func() {
	var spec *AWSLambdaSpec

	BeforeEach(func() {
		spec = &AWSLambdaSpec{
			FunctionName:        "my-func",
			RoleArn:             "arn:aws:iam::123456789012:role/service-role/my-role",
			Runtime:             "nodejs18.x",
			Handler:             "index.handler",
			MemoryMb:            512,
			TimeoutSeconds:      30,
			ReservedConcurrency: 1,
			Environment:         map[string]string{"FOO": "bar"},
			Subnets:             nil, // optional
			SecurityGroups:      nil, // optional
			Architecture:        Architecture_X86_64,
			LayerArns:           []string{"arn:aws:lambda:us-east-1:123456789012:layer:my-layer:1"},
			KmsKeyArn:           "",
			CodeSourceType:      CodeSourceType_CODE_SOURCE_TYPE_S3,
			S3: &S3Code{
				Bucket: "my-bucket",
				Key:    "artifacts/app.zip",
			},
			ImageUri: "",
		}
	})

	It("accepts a valid S3-based spec", func() {
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("fails when function_name is empty", func() {
		spec.FunctionName = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when role_arn is an invalid format", func() {
		spec.RoleArn = "not-an-arn"
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when code_source_type is unspecified", func() {
		spec.CodeSourceType = CodeSourceType_CODE_SOURCE_TYPE_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails CEL when S3 type but missing runtime/handler or image_uri set", func() {
		spec.CodeSourceType = CodeSourceType_CODE_SOURCE_TYPE_S3
		spec.S3 = &S3Code{Bucket: "my-bucket", Key: "app.zip"}
		spec.Runtime = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())

		spec.Runtime = "nodejs18.x"
		spec.Handler = "index.handler"
		spec.ImageUri = "123456789012.dkr.ecr.us-east-1.amazonaws.com/repo:tag"
		err = protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("accepts a valid image-based spec and fails if s3 is set", func() {
		spec.CodeSourceType = CodeSourceType_CODE_SOURCE_TYPE_IMAGE
		spec.S3 = nil
		spec.ImageUri = "123456789012.dkr.ecr.us-east-1.amazonaws.com/repo:tag"
		// runtime/handler are ignored for image type in CEL
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())

		spec.S3 = &S3Code{Bucket: "b", Key: "k"}
		err = protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("enforces memory and timeout bounds when set", func() {
		spec.MemoryMb = 64
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())

		spec.MemoryMb = 128
		spec.TimeoutSeconds = 901
		err = protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("enforces reserved_concurrency semantics", func() {
		spec.ReservedConcurrency = -2
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())

		spec.ReservedConcurrency = 0
		err = protovalidate.Validate(spec)
		Expect(err).To(BeNil())

		spec.ReservedConcurrency = 5
		err = protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})
})
