package awscertmanagercertv1

import (
	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/shared/validateutil"
	"testing"
)

func TestAwsCertManagerCert(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsCertManagerCert Suite")
}

var _ = Describe("AwsCertManagerCert", func() {

	var input *AwsCertManagerCert

	BeforeEach(func() {
		input = &AwsCertManagerCert{
			ApiVersion: "aws.project-planton.org/v1",
			Metadata: &shared.ApiResourceMetadata{
				Name: "a-test-name",
				Version: &shared.ApiResourceMetadataVersion{
					Message: "a version message",
				},
			},
			Spec: &AwsCertManagerCertSpec{
				PrimaryDomainName: "example.com",
				AlternateDomainNames: []string{
					"www.example.com",
					"test.example.com",
				},
				Route53HostedZoneId: "a-route53-hosted-zone-id",
				ValidationMethod:    "DNS",
			},
		}
	})

	Context("when valid input is passed", func() {
		It("should not return a validation error", func() {
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})

	Context("when name is not passed", func() {
		It("should return a validation error for 'metadata.name'", func() {
			input.Metadata.Name = ""

			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				expected := &validateutil.ExpectedViolation{
					FieldPath:    validateutil.MetadataFieldPath,
					ConstraintId: validateutil.MetadataNameConstraintId,
					Message:      "Name must be between 3 and 63 characters long",
				}
				validateutil.Match(validationErr.Violations[0], expected)
			}
		})
	})

	Context("when version message is not passed", func() {
		It("should return a validation error indicating 'version.message' is missing", func() {
			input.Metadata.Version.Message = ""
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				validateutil.Match(validationErr.Violations[0], validateutil.VersionMessageViolation)
			}
		})
	})

	Context("when metadata is not passed", func() {
		It("should return a validation error for missing metadata", func() {
			input.Metadata = nil
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				expected := &validateutil.ExpectedViolation{
					FieldPath:    validateutil.MetadataFieldPath,
					ConstraintId: validateutil.RequiredConstraint,
					Message:      validateutil.RequiredViolationMessage,
				}
				validateutil.Match(validationErr.Violations[0], expected)
			}
		})
	})

	Context("when name is empty", func() {
		It("should return a validation error for 'metadata.name'", func() {
			input.Metadata.Name = ""
			err := protovalidate.Validate(input)
			Expect(err).ToNot(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				expected := &validateutil.ExpectedViolation{
					FieldPath:    validateutil.MetadataFieldPath,
					ConstraintId: validateutil.MetadataNameConstraintId,
					Message:      "Name must be between 3 and 63 characters long",
				}
				validateutil.Match(validationErr.Violations[0], expected)
			}
		})
	})

	Context("when name length is greater than allowed)", func() {
		It("should return a validation error", func() {
			input.Metadata.Name = "a-really-really-long-name-that-is-greater-than-63-characters-long"
			err := protovalidate.Validate(input)
			Expect(err).ToNot(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				expected := &validateutil.ExpectedViolation{
					FieldPath:    validateutil.MetadataFieldPath,
					ConstraintId: validateutil.MetadataNameConstraintId,
					Message:      "Name must be between 3 and 63 characters long",
				}
				validateutil.Match(validationErr.Violations[0], expected)
			}
		})
	})

	Context("when validation method is invalid", func() {
		It("should return a validation error for 'spec.validation_method'", func() {
			input.Spec.ValidationMethod = "FAKE"

			err := protovalidate.Validate(input)
			Expect(err).ToNot(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				for _, violation := range validationErr.Violations {
					expected := &validateutil.ExpectedViolation{
						FieldPath:    "spec.validation_method",
						ConstraintId: validateutil.StringInConstraint,
						Message:      "value must be in list [\"DNS\", \"EMAIL\"]",
					}
					validateutil.Match(violation, expected)
				}
			}
		})
	})
})
