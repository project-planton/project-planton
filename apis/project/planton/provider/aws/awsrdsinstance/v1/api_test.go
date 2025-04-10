package awsrdsinstancev1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsRdsInstance(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsRdsInstance Suite")
}

var _ = Describe("AwsRdsInstance Custom Validation Tests", func() {
	var input *AwsRdsInstance

	BeforeEach(func() {
		input = &AwsRdsInstance{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsRdsInstance",
			Metadata: &shared.ApiResourceMetadata{
				Name: "valid-name",
			},
			Spec: &AwsRdsInstanceSpec{
				Engine:           "postgres",
				EngineVersion:    "14.6",
				InstanceClass:    "db.t3.medium",
				DbParameterGroup: "default.postgres14",
				StorageType:      "gp2",
				LicenseModel:     "bring-your-own-license",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("aws_rds", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			Context("api_version field", func() {
				It("should fail if api_version is invalid", func() {
					input.ApiVersion = "invalid.version"
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})
			})

			Context("kind field", func() {
				It("should fail if kind is invalid", func() {
					input.Kind = "InvalidKind"
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})
			})

			Context("storage_type field", func() {
				It("should fail if storage_type is not in the allowed set", func() {
					input.Spec.StorageType = "unrecognizedType"
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})
			})

			Context("license_model field", func() {
				It("should fail if license_model is not in the allowed set", func() {
					input.Spec.LicenseModel = "unknown-license"
					err := protovalidate.Validate(input)
					Expect(err).NotTo(BeNil())
				})
			})
		})
	})
})
