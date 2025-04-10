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
	Context("with an empty input", func() {
		It("should not return a validation error", func() {
			input := &AwsCertManagerCert{}
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})

	Context("when version message is not passed", func() {
		It("should return a validation error indicating 'version.message' is missing", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
				Metadata:   &shared.ApiResourceMetadata{},
			}
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[metadata.version.message"))
			Expect(err.Error()).To(ContainSubstring("Version message is required and cannot be empty"))
		})
	})

	Context("when metadata is not passed", func() {
		It("should return a validation error for missing metadata", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
			}
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[metadata]"))
			Expect(err.Error()).To(ContainSubstring("value is required"))
		})
	})

	Context("when version message is passed", func() {
		It("should not return a validation error", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
				Metadata: &shared.ApiResourceMetadata{
					Version: &shared.ApiResourceMetadataVersion{
						Message: "test aws-cert-manager-cert",
					},
				},
			}
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})

	Context("when name is empty", func() {
		It("should return a validation error for 'metadata.name'", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
				Metadata: &shared.ApiResourceMetadata{
					Version: &shared.ApiResourceMetadataVersion{
						Message: "test aws-cert-manager-cert",
					},
				},
			}
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[metadata.name]"))
			Expect(err.Error()).To(ContainSubstring("Name must be between 3 and 63 characters long"))
		})
	})

	Context("when name length is greater than 63", func() {
		It("should return a validation error for 'metadata.name'", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
				Metadata: &shared.ApiResourceMetadata{
					Name: "this is test name to check length validation. adding additional text to make it greater than 63",
					Version: &shared.ApiResourceMetadataVersion{
						Message: "test aws-cert-manager-cert",
					},
				},
			}
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			Expect(err.Error()).To(ContainSubstring("[metadata.name]"))
			Expect(err.Error()).To(ContainSubstring("Name must be between 3 and 63 characters long"))
		})
	})

	Context("when name length is valid (less than or equal to 63)", func() {
		It("should not return a validation error", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
				Metadata: &shared.ApiResourceMetadata{
					Name: "test",
					Version: &shared.ApiResourceMetadataVersion{
						Message: "test aws-cert-manager-cert",
					},
				},
			}
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})

	Context("when validation method is invalid", func() {
		It("should return a validation error for 'spec.validation_method'", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
				Metadata: &shared.ApiResourceMetadata{
					Name: "testName",
					Version: &shared.ApiResourceMetadataVersion{
						Message: "test aws-cert-manager-cert",
					},
				},
				Spec: &AwsCertManagerCertSpec{
					PrimaryDomainName:   "example.com",
					Route53HostedZoneId: "Z123456ABCXYZ",
					ValidationMethod:    "FAKE",
				},
			}
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			var validationErr *protovalidate.ValidationError
			if err != nil {
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
			}
		})
	})

	Context("when validation method is set to DNS", func() {
		It("should not return a validation error", func() {
			input := &AwsCertManagerCert{
				ApiVersion: "aws.project-planton.org/v1",
				Metadata: &shared.ApiResourceMetadata{
					Name: "example.com",
					Version: &shared.ApiResourceMetadataVersion{
						Message: "aws-cert-manager-cert test",
					},
				},
				Spec: &AwsCertManagerCertSpec{
					PrimaryDomainName:   "example.com",
					Route53HostedZoneId: "Z123456ABCXYZ",
					ValidationMethod:    "DNS",
				},
			}
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})
})
