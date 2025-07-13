package awsdynamodbv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

var validator protovalidate.Validator

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = BeforeSuite(func() {
    var err error
    validator, err = protovalidate.New()
    Expect(err).NotTo(HaveOccurred())
})

// validSpec returns a minimal, fully valid AwsDynamodbSpec instance.
func validSpec() *AwsDynamodbSpec {
    return &AwsDynamodbSpec{
        TableName: "my-table",
        HashKey:   "pk",
        Attributes: []*AttributeDefinition{
            {
                Name: "pk",
                Type: AttributeType_STRING,
            },
        },
        BillingMode:   BillingMode_PROVISIONED,
        ReadCapacity:  1,
        WriteCapacity: 1,
        TableClass:    TableClass_STANDARD,
    }
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    It("accepts a fully valid spec", func() {
        spec := validSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    It("rejects a table_name that is too short", func() {
        spec := validSpec()
        spec.TableName = "ab" // min length is 3
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects a table_name with invalid characters", func() {
        spec := validSpec()
        spec.TableName = "invalid name" // space not allowed by pattern
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when no attributes are provided", func() {
        spec := validSpec()
        spec.Attributes = nil // repeated min_items = 1
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when an attribute has unspecified type", func() {
        spec := validSpec()
        spec.Attributes = []*AttributeDefinition{{Name: "pk", Type: AttributeType_ATTRIBUTE_TYPE_UNSPECIFIED}}
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects an unspecified billing_mode", func() {
        spec := validSpec()
        spec.BillingMode = BillingMode_BILLING_MODE_UNSPECIFIED // not_in: [0]
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects read_capacity below minimum", func() {
        spec := validSpec()
        spec.ReadCapacity = 0 // gte: 1
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects write_capacity below minimum", func() {
        spec := validSpec()
        spec.WriteCapacity = 0 // gte: 1
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects tags with empty key", func() {
        spec := validSpec()
        spec.Tags = map[string]string{"": "value"} // key must be min_len 1
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects tags with empty value", func() {
        spec := validSpec()
        spec.Tags = map[string]string{"key": ""} // value must be min_len 1
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })
})
