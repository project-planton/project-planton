package awsdynamodbv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsDynamodbSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsDynamodbSpec Custom Validation Tests")
}

var _ = Describe("AwsDynamodbSpec Custom Validation Tests", func() {

	var input *AwsDynamodb

	BeforeEach(func() {
		input = &AwsDynamodb{
			ApiVersion: "aws.project-planton.org/v1",
			Kind:       "AwsDynamodb",
			Metadata: &shared.ApiResourceMetadata{
				Name: "test-dynamodb",
			},
			Spec: &AwsDynamodbSpec{
				TableName:   "my-table",
				BillingMode: "PROVISIONED",
				HashKey: &AwsDynamodbTableAttribute{
					Name: "Id",
					Type: "S",
				},
				Attributes: []*AwsDynamodbTableAttribute{
					{
						Name: "Id",
						Type: "S",
					},
				},
			},
		}
	})

	Describe("When valid input is passed", func() {

		Context("Valid Sample Object", func() {
			It("should not return a validation error", func() {
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("billing_mode enumerations", func() {
			It("should accept 'PAY_PER_REQUEST'", func() {
				input.Spec.BillingMode = "PAY_PER_REQUEST"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid billing_mode", func() {
				input.Spec.BillingMode = "INVALID_BILLING"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("stream_view_type enumerations", func() {
			It("should accept 'NEW_IMAGE'", func() {
				input.Spec.StreamViewType = "NEW_IMAGE"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid stream_view_type", func() {
				input.Spec.StreamViewType = "IMAGE_ONLY"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("AwsDynamodbTableAttribute.type enumerations", func() {
			It("should accept 'N'", func() {
				input.Spec.HashKey.Type = "N"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid attribute type", func() {
				input.Spec.HashKey.Type = "STRING"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("AwsDynamodbTableGlobalSecondaryIndex.projection_type enumerations", func() {
			var gsi AwsDynamodbTableGlobalSecondaryIndex

			BeforeEach(func() {
				gsi = AwsDynamodbTableGlobalSecondaryIndex{
					Name:           "example-gsi",
					HashKey:        "Id",
					ProjectionType: "ALL",
				}
				input.Spec.GlobalSecondaryIndexes = []*AwsDynamodbTableGlobalSecondaryIndex{&gsi}
			})

			It("should accept 'KEYS_ONLY'", func() {
				gsi.ProjectionType = "KEYS_ONLY"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid projection_type", func() {
				gsi.ProjectionType = "SOME_INVALID_TYPE"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("AwsDynamodbTableLocalSecondaryIndex.projection_type enumerations", func() {
			var lsi AwsDynamodbTableLocalSecondaryIndex

			BeforeEach(func() {
				lsi = AwsDynamodbTableLocalSecondaryIndex{
					Name:           "example-lsi",
					RangeKey:       "RangeKeyAttr",
					ProjectionType: "ALL",
				}
				input.Spec.LocalSecondaryIndexes = []*AwsDynamodbTableLocalSecondaryIndex{&lsi}
			})

			It("should accept 'INCLUDE'", func() {
				lsi.ProjectionType = "INCLUDE"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid projection_type", func() {
				lsi.ProjectionType = "INVALID_PROJECTION"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("AwsDynamodbTableImport.input_compression_type enumerations", func() {
			var imp AwsDynamodbTableImport

			BeforeEach(func() {
				imp = AwsDynamodbTableImport{
					InputCompressionType: "NONE",
					InputFormat:          "CSV",
				}
				input.Spec.ImportTable = &imp
			})

			It("should accept 'GZIP'", func() {
				imp.InputCompressionType = "GZIP"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid input_compression_type", func() {
				imp.InputCompressionType = "LZ4"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})

		Context("AwsDynamodbTableImport.input_format enumerations", func() {
			var imp AwsDynamodbTableImport

			BeforeEach(func() {
				imp = AwsDynamodbTableImport{
					InputCompressionType: "NONE",
					InputFormat:          "CSV",
				}
				input.Spec.ImportTable = &imp
			})

			It("should accept 'ION'", func() {
				imp.InputFormat = "ION"
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should reject an invalid input_format", func() {
				imp.InputFormat = "XML"
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
			})
		})
	})
})
