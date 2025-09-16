package confluentkafkav1

import (
	"testing"

	"buf.build/go/protovalidate"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestConfluentKafka(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "ConfluentKafka Suite")
}

var _ = ginkgo.Describe("ConfluentKafka Custom Validation Tests", func() {
	var input *ConfluentKafka

	ginkgo.BeforeEach(func() {
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

	ginkgo.Describe("When valid input is passed", func() {
		ginkgo.Context("AWS", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.Cloud = "AWS"
				input.Spec.Availability = "SINGLE_ZONE"
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("GCP", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.Cloud = "GCP"
				input.Spec.Availability = "MULTI_ZONE"
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})

		ginkgo.Context("AZURE", func() {
			ginkgo.It("should not return a validation error", func() {
				input.Spec.Cloud = "AZURE"
				input.Spec.Availability = "HIGH"
				err := protovalidate.Validate(input)
				gomega.Expect(err).To(gomega.BeNil())
			})
		})
	})

	ginkgo.Describe("Cloud Field Validation", func() {
		ginkgo.It("should fail validation if the cloud field is invalid", func() {
			input.Spec.Cloud = "IBM"
			err := protovalidate.Validate(input)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("should allow the cloud field to be omitted", func() {
			input.Spec.Cloud = ""
			err := protovalidate.Validate(input)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})

	ginkgo.Describe("Availability Field Validation", func() {
		ginkgo.It("should fail validation if the availability field is invalid", func() {
			input.Spec.Availability = "UNSUPPORTED_ZONE"
			err := protovalidate.Validate(input)
			gomega.Expect(err).NotTo(gomega.BeNil())
		})

		ginkgo.It("should allow the availability field to be omitted", func() {
			input.Spec.Availability = ""
			err := protovalidate.Validate(input)
			gomega.Expect(err).To(gomega.BeNil())
		})
	})
})
