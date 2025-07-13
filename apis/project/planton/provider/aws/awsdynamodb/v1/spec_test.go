package awsdynamodbv1

import (
    "strings"
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    var (
        spec       *AwsDynamodbSpec
        validator  protovalidate.Validator
    )

    BeforeEach(func() {
        // Instantiate the validator as required
        vPtr, err := protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
        validator = *vPtr

        // Build a minimal valid spec
        spec = &AwsDynamodbSpec{
            TableName: "valid_table",
            AttributeDefinitions: []*AttributeDefinition{
                {
                    Name: "id",
                    Type: AttributeType_STRING,
                },
            },
            KeySchema: &KeySchema{
                PartitionKey: &KeyElement{
                    AttributeName: "id",
                    KeyType:       KeyType_HASH,
                },
            },
            BillingMode: BillingMode_PAY_PER_REQUEST,
            Tags: map[string]string{
                "env": "test",
            },
            TableClass: TableClass_STANDARD,
        }
    })

    It("accepts a fully valid spec", func() {
        Expect(validator.Validate(spec)).To(Succeed())
    })

    Context("table_name", func() {
        It("fails when too short", func() {
            spec.TableName = "ab"
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when invalid characters used", func() {
            spec.TableName = "invalid name!"
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })

    Context("attribute_definitions", func() {
        It("fails when list is empty", func() {
            spec.AttributeDefinitions = []*AttributeDefinition{}
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when attribute name is empty", func() {
            spec.AttributeDefinitions[0].Name = ""
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when attribute type is unspecified", func() {
            spec.AttributeDefinitions[0].Type = AttributeType_ATTRIBUTE_TYPE_UNSPECIFIED
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })

    Context("key_schema", func() {
        It("fails when key_schema is missing", func() {
            spec.KeySchema = nil
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })

    Context("enums that must not be zero", func() {
        It("fails when billing_mode is unspecified", func() {
            spec.BillingMode = BillingMode_BILLING_MODE_UNSPECIFIED
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when table_class is unspecified", func() {
            spec.TableClass = TableClass_TABLE_CLASS_UNSPECIFIED
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })

    Context("tags", func() {
        It("fails when the key is empty", func() {
            spec.Tags = map[string]string{"": "value"}
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when the value is too long", func() {
            spec.Tags = map[string]string{"k": strings.Repeat("a", 257)}
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })
})
