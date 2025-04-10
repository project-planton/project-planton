package shared

import (
	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/validateutil"
	"testing"
)

const (
	metadataNameFieldPath                 = "name"
	metadataNameMinLengthViolationMessage = "value length must be at least 3 characters"
	metadataNameMaxLengthViolationMessage = "value length must be at most 63 characters"
)

func TestMetadata(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Metadata Suite")
}

var _ = Describe("Metadata", func() {
	Context("when name is not passed", func() {
		It("should return a validation error for 'metadata.name'", func() {
			input := &ApiResourceMetadata{
				Name: "",
			}

			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				expected := &validateutil.ExpectedViolation{
					FieldPath:    metadataNameFieldPath,
					ConstraintId: validateutil.RequiredConstraint,
					Message:      validateutil.RequiredViolationMessage,
				}
				validateutil.Match(validationErr.Violations[0], expected)
			}
		})
	})

	Context("when name length is less than minimum", func() {
		It("should return a validation error for 'metadata.name'", func() {
			input := &ApiResourceMetadata{
				Name: "a",
			}
			err := protovalidate.Validate(input)
			Expect(err).ToNot(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				expected := &validateutil.ExpectedViolation{
					FieldPath:    metadataNameFieldPath,
					ConstraintId: validateutil.StringMinLengthConstraint,
					Message:      metadataNameMinLengthViolationMessage,
				}
				validateutil.Match(validationErr.Violations[0], expected)
			}
		})
	})

	Context("when name length is greater than allowed)", func() {
		It("should return a validation error", func() {
			input := &ApiResourceMetadata{
				Name: "a-really-really-long-name-that-is-greater-than-63-characters-long",
			}
			err := protovalidate.Validate(input)
			Expect(err).ToNot(BeNil())
			var validationErr *protovalidate.ValidationError
			if errors.As(err, &validationErr) {
				Expect(len(validationErr.Violations)).To(Equal(1))
				expected := &validateutil.ExpectedViolation{
					FieldPath:    metadataNameFieldPath,
					ConstraintId: validateutil.StringMaxLengthConstraint,
					Message:      metadataNameMaxLengthViolationMessage,
				}
				validateutil.Match(validationErr.Violations[0], expected)
			}
		})
	})
})
