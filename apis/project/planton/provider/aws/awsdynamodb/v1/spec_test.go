package awsdynamodb_test

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    awsdynamodb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// Ginkgo entry point
func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec validation suite")
}

var _ = Describe("AwsDynamodbSpec protovalidate rules", func() {
    var validator *protovalidate.Validator

    BeforeSuite(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    Context("valid spec", func() {
        It("accepts a minimal valid spec", func() {
            spec := &awsdynamodb.AwsDynamodbSpec{
                TableName: "my_table",
                AttributeDefinitions: []*awsdynamodb.AwsDynamodbSpec_AttributeDefinition{
                    {
                        Name: "pk",
                        Type: "S",
                    },
                },
                KeySchema: []*awsdynamodb.AwsDynamodbSpec_KeySchemaElement{
                    {
                        AttributeName: "pk",
                        KeyType:       awsdynamodb.AwsDynamodbSpec_KeySchemaElement_HASH,
                    },
                },
                BillingMode:        awsdynamodb.AwsDynamodbSpec_PROVISIONED,
                ReadCapacityUnits:  5,
                WriteCapacityUnits: 5,
                TableClass:         awsdynamodb.AwsDynamodbSpec_STANDARD,
            }

            Expect(validator.Validate(spec)).To(Succeed())
        })
    })

    Context("invalid specs", func() {
        It("fails when table_name is empty", func() {
            spec := &awsdynamodb.AwsDynamodbSpec{}
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when billing_mode is unspecified (enum zero)", func() {
            spec := &awsdynamodb.AwsDynamodbSpec{
                TableName: "tbl",
                AttributeDefinitions: []*awsdynamodb.AwsDynamodbSpec_AttributeDefinition{
                    {Name: "pk", Type: "S"},
                },
                KeySchema: []*awsdynamodb.AwsDynamodbSpec_KeySchemaElement{
                    {AttributeName: "pk", KeyType: awsdynamodb.AwsDynamodbSpec_KeySchemaElement_HASH},
                },
                // BillingMode left as zero value (unspecified)
                ReadCapacityUnits:  1,
                WriteCapacityUnits: 1,
                TableClass:         awsdynamodb.AwsDynamodbSpec_STANDARD,
            }
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })

        It("fails when attribute type is not one of S/N/B", func() {
            spec := &awsdynamodb.AwsDynamodbSpec{
                TableName: "good_name",
                AttributeDefinitions: []*awsdynamodb.AwsDynamodbSpec_AttributeDefinition{{
                    Name: "pk",
                    Type: "X", // invalid
                }},
                KeySchema: []*awsdynamodb.AwsDynamodbSpec_KeySchemaElement{{
                    AttributeName: "pk",
                    KeyType:       awsdynamodb.AwsDynamodbSpec_KeySchemaElement_HASH,
                }},
                BillingMode:        awsdynamodb.AwsDynamodbSpec_PROVISIONED,
                ReadCapacityUnits:  1,
                WriteCapacityUnits: 1,
                TableClass:         awsdynamodb.AwsDynamodbSpec_STANDARD,
            }
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
        })
    })
})
