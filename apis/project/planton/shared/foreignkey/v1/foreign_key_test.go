// filename: foreign_key_test.go
package foreignkeyv1

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestForeignKey(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ForeignKey Suite")
}

var _ = Describe("ForeignKey Oneof Tests", func() {

	Describe("StringValueOrRef usage", func() {
		Context("when setting a literal value only", func() {
			It("should not return a validation error", func() {
				input := &StringValueOrRef{
					LiteralOrRef: &StringValueOrRef_Value{
						Value: "my-string",
					},
				}
				// No custom rule; we expect no errors from protovalidate
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("when setting a ValueFromRef only", func() {
			It("should not return a validation error", func() {
				input := &StringValueOrRef{
					LiteralOrRef: &StringValueOrRef_ValueFrom{
						ValueFrom: &ValueFromRef{
							Kind:      cloudresourcekind.CloudResourceKind_FirstTestCloudApiResource,
							Env:       "dev",
							Slug:      "my-cert",
							FieldPath: "status.outputs.cert_arn",
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("when setting both (proto oneof overwrites the first)", func() {
			It("should end up with the last field set and not produce a validation error", func() {
				// The 'value' gets overwritten by the oneof assignment to 'value_from'
				input := &StringValueOrRef{
					LiteralOrRef: &StringValueOrRef_Value{
						Value: "my-string",
					},
				}
				input.LiteralOrRef = &StringValueOrRef_ValueFrom{
					ValueFrom: &ValueFromRef{
						Kind: cloudresourcekind.CloudResourceKind_FirstTestCloudApiResource,
						Env:  "dev",
						Slug: "overwrites-literal",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())

				// Because this is a oneof, the 'Value' is no longer set
				Expect(input.GetValue()).To(Equal(""))
				Expect(input.GetValueFrom().GetSlug()).To(Equal("overwrites-literal"))
			})
		})
	})

	Describe("Int32ValueOrRef usage", func() {
		Context("when setting an int32 literal only", func() {
			It("should not return a validation error", func() {
				input := &Int32ValueOrRef{
					LiteralOrRef: &Int32ValueOrRef_Value{
						Value: 123,
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("when setting a ValueFromRef only", func() {
			It("should not return a validation error", func() {
				input := &Int32ValueOrRef{
					LiteralOrRef: &Int32ValueOrRef_ValueFrom{
						ValueFrom: &ValueFromRef{
							Kind: cloudresourcekind.CloudResourceKind_FirstTestCloudApiResource,
							Env:  "dev",
							Slug: "ref-int32",
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})

		Context("when setting both fields in code sequentially", func() {
			It("should overwrite the first field with the second (no error)", func() {
				input := &Int32ValueOrRef{
					LiteralOrRef: &Int32ValueOrRef_Value{
						Value: 456,
					},
				}

				// Overwrite with ValueFromRef
				input.LiteralOrRef = &Int32ValueOrRef_ValueFrom{
					ValueFrom: &ValueFromRef{
						Kind: cloudresourcekind.CloudResourceKind_FirstTestCloudApiResource,
						Env:  "dev",
						Slug: "ref-overwrites-int",
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())

				// Because it's a oneof, the integer literal is no longer set
				Expect(input.GetValue()).To(BeEquivalentTo(0))
				Expect(input.GetValueFrom().GetSlug()).To(Equal("ref-overwrites-int"))
			})
		})
	})
})
