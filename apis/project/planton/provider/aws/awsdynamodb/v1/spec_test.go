package awsdynamodbv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestAwsDynamodbSpec(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    var validator protovalidate.Validator

    BeforeEach(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    Context("valid spec", func() {
        It("passes validation", func() {
            spec := &AwsDynamodbSpec{
                TableName:    "my_table",
                BillingMode:  "PROVISIONED",
                ReadCapacity: 5,
                WriteCapacity: 5,
                AttributeDefinitions: []*AttributeDefinition{
                    {
                        AttributeName: "id",
                        AttributeType: "S",
                    },
                },
                KeySchema: []*KeySchemaElement{
                    {
                        AttributeName: "id",
                        KeyType:       "HASH",
                    },
                },
                TableClass: TableClass_STANDARD,
            }

            err := validator.Validate(spec)
            Expect(err).NotTo(HaveOccurred())
        })
    })

    Context("invalid table name", func() {
        It("fails when table name is too short", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "ab",
                BillingMode: "PROVISIONED",
                ReadCapacity: 5,
                WriteCapacity: 5,
                AttributeDefinitions: []*AttributeDefinition{
                    {
                        AttributeName: "id",
                        AttributeType: "S",
                    },
                },
                KeySchema: []*KeySchemaElement{
                    {
                        AttributeName: "id",
                        KeyType:       "HASH",
                    },
                },
                TableClass: TableClass_STANDARD,
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when table name contains invalid characters", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "invalid name!",
                BillingMode: "PROVISIONED",
                ReadCapacity: 5,
                WriteCapacity: 5,
                AttributeDefinitions: []*AttributeDefinition{
                    {
                        AttributeName: "id",
                        AttributeType: "S",
                    },
                },
                KeySchema: []*KeySchemaElement{
                    {
                        AttributeName: "id",
                        KeyType:       "HASH",
                    },
                },
                TableClass: TableClass_STANDARD,
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })
    })

    Context("invalid billing mode", func() {
        It("fails when billing mode is not allowed", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "valid_table_name",
                BillingMode: "UNKNOWN",
                ReadCapacity: 5,
                WriteCapacity: 5,
                AttributeDefinitions: []*AttributeDefinition{
                    {
                        AttributeName: "id",
                        AttributeType: "S",
                    },
                },
                KeySchema: []*KeySchemaElement{
                    {
                        AttributeName: "id",
                        KeyType:       "HASH",
                    },
                },
                TableClass: TableClass_STANDARD,
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })
    })

    Context("missing required repeated fields", func() {
        It("fails when attribute_definitions is empty", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "valid_table_name",
                BillingMode: "PROVISIONED",
                ReadCapacity: 5,
                WriteCapacity: 5,
                // AttributeDefinitions intentionally missing
                KeySchema: []*KeySchemaElement{
                    {
                        AttributeName: "id",
                        KeyType:       "HASH",
                    },
                },
                TableClass: TableClass_STANDARD,
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when key_schema is empty", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "valid_table_name",
                BillingMode: "PROVISIONED",
                ReadCapacity: 5,
                WriteCapacity: 5,
                AttributeDefinitions: []*AttributeDefinition{
                    {
                        AttributeName: "id",
                        AttributeType: "S",
                    },
                },
                // KeySchema intentionally missing
                TableClass: TableClass_STANDARD,
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })
    })
})