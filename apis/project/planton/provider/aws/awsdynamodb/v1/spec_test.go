package awsdynamodbv1

import (
    "context"
    "testing"

    protovalidate "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var (
    validator protovalidate.Validator
)

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = BeforeSuite(func() {
    v, err := protovalidate.New()
    Expect(err).NotTo(HaveOccurred())
    validator = *v
})

func validSpec() *AwsDynamodbSpec {
    return &AwsDynamodbSpec{
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
        TableClass:  TableClass_STANDARD,
    }
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    It("accepts a fully valid spec", func() {
        spec := validSpec()
        Expect(validator.Validate(context.Background(), spec)).To(Succeed())
    })

    It("rejects a table name that is too short", func() {
        spec := validSpec()
        spec.TableName = "ab" // min_len is 3
        Expect(validator.Validate(context.Background(), spec)).ToNot(Succeed())
    })

    It("rejects missing key schema", func() {
        spec := validSpec()
        spec.KeySchema = nil // required=true
        Expect(validator.Validate(context.Background(), spec)).ToNot(Succeed())
    })

    It("rejects empty attribute definitions", func() {
        spec := validSpec()
        spec.AttributeDefinitions = []*AttributeDefinition{} // min_items=1
        Expect(validator.Validate(context.Background(), spec)).ToNot(Succeed())
    })

    It("rejects unspecified billing mode (enum zero value)", func() {
        spec := validSpec()
        spec.BillingMode = BillingMode_BILLING_MODE_UNSPECIFIED // not_in: [0]
        Expect(validator.Validate(context.Background(), spec)).ToNot(Succeed())
    })
})
