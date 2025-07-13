package awsdynamodbv1

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "github.com/bufbuild/protovalidate-go"
)

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = Describe("AwsDynamodbSpec", func() {
    var validator protovalidate.Validator

    BeforeEach(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    Context("valid messages", func() {
        It("accepts a fully valid spec", func() {
            spec := &AwsDynamodbSpec{
                TableName: "my_table",
                AttributeDefinitions: []*AttributeDefinition{
                    {
                        Name: "pk",
                        Type: AttributeType_STRING,
                    },
                },
                KeySchema: &KeySchema{
                    PartitionKey: &KeyElement{
                        AttributeName: "pk",
                        KeyType:       KeyType_HASH,
                    },
                },
                BillingMode: BillingMode_PAY_PER_REQUEST,
                Tags: map[string]string{
                    "env": "prod",
                },
                TableClass: TableClass_STANDARD,
            }

            Expect(validator.Validate(spec)).To(Succeed())
        })
    })

    Context("invalid messages", func() {
        It("rejects a table name that is too short", func() {
            spec := &AwsDynamodbSpec{
                TableName: "ab", // min_len = 3
                AttributeDefinitions: []*AttributeDefinition{{
                    Name: "pk",
                    Type: AttributeType_STRING,
                }},
                KeySchema: &KeySchema{
                    PartitionKey: &KeyElement{
                        AttributeName: "pk",
                        KeyType:       KeyType_HASH,
                    },
                },
                BillingMode: BillingMode_PAY_PER_REQUEST,
            }

            Expect(validator.Validate(spec)).ToNot(Succeed())
        })

        It("rejects when no attribute definitions are provided", func() {
            spec := &AwsDynamodbSpec{
                TableName:            "valid_name",
                AttributeDefinitions: []*AttributeDefinition{}, // violates min_items = 1
                KeySchema: &KeySchema{
                    PartitionKey: &KeyElement{
                        AttributeName: "pk",
                        KeyType:       KeyType_HASH,
                    },
                },
                BillingMode: BillingMode_PAY_PER_REQUEST,
            }

            Expect(validator.Validate(spec)).ToNot(Succeed())
        })

        It("rejects when key schema is missing (required=true)", func() {
            spec := &AwsDynamodbSpec{
                TableName: "valid_name",
                AttributeDefinitions: []*AttributeDefinition{{
                    Name: "pk",
                    Type: AttributeType_STRING,
                }},
                BillingMode: BillingMode_PAY_PER_REQUEST,
            }

            Expect(validator.Validate(spec)).ToNot(Succeed())
        })

        It("rejects an unspecified billing mode (enum not_in=[0])", func() {
            spec := &AwsDynamodbSpec{
                TableName: "valid_name",
                AttributeDefinitions: []*AttributeDefinition{{
                    Name: "pk",
                    Type: AttributeType_STRING,
                }},
                KeySchema: &KeySchema{
                    PartitionKey: &KeyElement{
                        AttributeName: "pk",
                        KeyType:       KeyType_HASH,
                    },
                },
                BillingMode: BillingMode_BILLING_MODE_UNSPECIFIED, // 0 is forbidden
            }

            Expect(validator.Validate(spec)).ToNot(Succeed())
        })
    })
})
