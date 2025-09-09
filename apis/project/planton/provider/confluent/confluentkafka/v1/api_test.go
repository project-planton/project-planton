package confluentkafkav1

import (
	"testing"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestConfluentKafka(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ConfluentKafka Suite")
}

var _ = Describe("ConfluentKafka Custom Validation Tests", func() {
	var input *ConfluentKafka

	BeforeEach(func() {
		input = &ConfluentKafka{
			ApiVersion: "confluent.project-planton.org/v1",
			Kind:       "ConfluentKafka",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-resource",
			},
			Spec: &ConfluentKafkaSpec{
				Cloud:        "AWS",
				Availability: "SINGLE_ZONE",
				Environment:  "sample-env",
			},
		}
	})

	Describe("When valid input is passed", func() {
		Context("AWS", func() {
			It("should not return a validation error", func() {
				input.Spec.Cloud = "AWS"
				input.Spec.Availability = "SINGLE_ZONE"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("GCP", func() {
			It("should not return a validation error", func() {
				input.Spec.Cloud = "GCP"
				input.Spec.Availability = "MULTI_ZONE"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("AZURE", func() {
			It("should not return a validation error", func() {
				input.Spec.Cloud = "AZURE"
				input.Spec.Availability = "HIGH"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})

	Describe("Cloud Field Validation", func() {
		It("should fail validation if the cloud field is invalid", func() {
			input.Spec.Cloud = "IBM"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should allow the cloud field to be omitted", func() {
			input.Spec.Cloud = ""
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})

	Describe("Availability Field Validation", func() {
		It("should fail validation if the availability field is invalid", func() {
			input.Spec.Availability = "UNSUPPORTED_ZONE"
			err := protovalidate.Validate(input)
			Expect(err).NotTo(BeNil())
		})

		It("should allow the availability field to be omitted", func() {
			input.Spec.Availability = ""
			err := protovalidate.Validate(input)
			Expect(err).To(BeNil())
		})
	})
})
