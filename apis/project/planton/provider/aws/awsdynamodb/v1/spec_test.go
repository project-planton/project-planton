package awsdynamodbv1

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "github.com/bufbuild/protovalidate-go"
)

func TestSuite(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    var (
        validator protovalidate.Validator
    )

    BeforeEach(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    It("accepts a fully valid spec", func() {
        spec := &AwsDynamodbSpec{
            TableName:          "my_table",
            BillingMode:        BillingMode_PROVISIONED,
            ReadCapacityUnits:  5,
            WriteCapacityUnits: 5,
            AttributeDefinitions: []*AttributeDefinition{
                {
                    AttributeName: "id",
                    AttributeType: AttributeType_S,
                },
            },
            KeySchema: []*KeySchemaElement{
                {
                    AttributeName: "id",
                    KeyType:       KeyType_HASH,
                },
            },
        }

        err := validator.Validate(spec)
        Expect(err).NotTo(HaveOccurred())
    })

    It("rejects a table name that is too short", func() {
        spec := &AwsDynamodbSpec{
            TableName:          "ab",
            BillingMode:        BillingMode_PROVISIONED,
            ReadCapacityUnits:  5,
            WriteCapacityUnits: 5,
            AttributeDefinitions: []*AttributeDefinition{
                {
                    AttributeName: "id",
                    AttributeType: AttributeType_S,
                },
            },
            KeySchema: []*KeySchemaElement{
                {
                    AttributeName: "id",
                    KeyType:       KeyType_HASH,
                },
            },
        }

        err := validator.Validate(spec)
        Expect(err).To(HaveOccurred())
    })

    It("enforces billing_mode/pay per request capacity rules", func() {
        spec := &AwsDynamodbSpec{
            TableName:          "valid_table",
            BillingMode:        BillingMode_PAY_PER_REQUEST,
            ReadCapacityUnits:  1,
            WriteCapacityUnits: 0,
            AttributeDefinitions: []*AttributeDefinition{
                {
                    AttributeName: "id",
                    AttributeType: AttributeType_S,
                },
            },
            KeySchema: []*KeySchemaElement{
                {
                    AttributeName: "id",
                    KeyType:       KeyType_HASH,
                },
            },
        }

        err := validator.Validate(spec)
        Expect(err).To(HaveOccurred())
    })

    It("requires at least one attribute definition", func() {
        spec := &AwsDynamodbSpec{
            TableName:          "valid_table",
            BillingMode:        BillingMode_PROVISIONED,
            ReadCapacityUnits:  5,
            WriteCapacityUnits: 5,
            KeySchema: []*KeySchemaElement{
                {
                    AttributeName: "id",
                    KeyType:       KeyType_HASH,
                },
            },
        }

        err := validator.Validate(spec)
        Expect(err).To(HaveOccurred())
    })
})