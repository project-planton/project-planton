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

    BeforeSuite(func() {
        v, err := protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
        validator = *v // convert pointer to value, per requirement
    })

    // Helper that returns a minimally valid specification
    buildBaseSpec := func() *AwsDynamodbSpec {
        return &AwsDynamodbSpec{
            TableName:   "valid_table",
            BillingMode: BillingMode_PROVISIONED,
            ProvisionedThroughput: &ProvisionedThroughput{
                ReadCapacityUnits:  5,
                WriteCapacityUnits: 5,
            },
            AttributeDefinitions: []*AttributeDefinition{
                {
                    AttributeName: "id",
                    AttributeType: AttributeType_STRING,
                },
            },
            KeySchema: &KeySchema{
                PartitionKey: "id",
            },
            TableClass: TableClass_STANDARD,
        }
    }

    Context("table name validation", func() {
        It("accepts a valid name", func() {
            spec := buildBaseSpec()
            Expect(validator.Validate(spec)).To(Succeed())
        })

        It("rejects too short name", func() {
            spec := buildBaseSpec()
            spec.TableName = "ab"
            Expect(validator.Validate(spec)).NotTo(Succeed())
        })

        It("rejects invalid characters", func() {
            spec := buildBaseSpec()
            spec.TableName = "invalid name" // space is not allowed by pattern
            Expect(validator.Validate(spec)).NotTo(Succeed())
        })
    })

    Context("billing mode / throughput CEL rule", func() {
        It("accepts PROVISIONED with positive RCU/WCU", func() {
            spec := buildBaseSpec()
            Expect(validator.Validate(spec)).To(Succeed())
        })

        It("rejects PROVISIONED without throughput", func() {
            spec := buildBaseSpec()
            spec.ProvisionedThroughput = nil
            Expect(validator.Validate(spec)).NotTo(Succeed())
        })

        It("accepts PAY_PER_REQUEST without throughput", func() {
            spec := buildBaseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            spec.ProvisionedThroughput = nil
            Expect(validator.Validate(spec)).To(Succeed())
        })

        It("rejects PAY_PER_REQUEST with positive throughput", func() {
            spec := buildBaseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            // Keep positive throughput
            Expect(validator.Validate(spec)).NotTo(Succeed())
        })
    })

    Context("required and repeated field rules", func() {
        It("rejects missing key_schema", func() {
            spec := buildBaseSpec()
            spec.KeySchema = nil
            Expect(validator.Validate(spec)).NotTo(Succeed())
        })

        It("rejects empty attribute_definitions", func() {
            spec := buildBaseSpec()
            spec.AttributeDefinitions = nil
            Expect(validator.Validate(spec)).NotTo(Succeed())
        })
    })
})
