package awsdynamodbv1

import (
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "testing"

    "github.com/bufbuild/protovalidate-go"
)

// -----------------------------------------------------------------------------
//  Ginkgo test bootstrap
// -----------------------------------------------------------------------------

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

// -----------------------------------------------------------------------------
//  Suite
// -----------------------------------------------------------------------------

var _ = Describe("AwsDynamodbSpec validation", func() {
    var (
        validator protovalidate.Validator
    )

    // -------------------------------------------------------------------------
    //  Helpers
    // -------------------------------------------------------------------------

    minimalValidSpec := func() *AwsDynamodbSpec {
        return &AwsDynamodbSpec{
            TableName: "users",
            AttributeDefinitions: []*AttributeDefinition{
                {
                    AttributeName: "pk",
                    AttributeType: AttributeType_STRING,
                },
            },
            KeySchema: []*KeySchemaElement{
                {
                    AttributeName: "pk",
                    KeyType:       KeyType_HASH,
                },
            },
            BillingMode: BillingMode_PAY_PER_REQUEST,
        }
    }

    // -------------------------------------------------------------------------
    //  Setup
    // -------------------------------------------------------------------------

    BeforeEach(func() {
        var err error
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    // -------------------------------------------------------------------------
    //  Positive cases
    // -------------------------------------------------------------------------

    Context("valid specifications", func() {
        It("accepts a minimal PAY_PER_REQUEST spec", func() {
            spec := minimalValidSpec()
            Expect(validator.Validate(spec)).To(Succeed())
        })

        It("accepts a valid PROVISIONED spec with throughput", func() {
            spec := minimalValidSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            spec.ProvisionedThroughput = &ProvisionedThroughput{
                ReadCapacityUnits:  5,
                WriteCapacityUnits: 10,
            }
            Expect(validator.Validate(spec)).To(Succeed())
        })
    })

    // -------------------------------------------------------------------------
    //  Negative cases – field level rules
    // -------------------------------------------------------------------------

    Context("field validation failures", func() {
        It("fails when table_name is missing", func() {
            spec := minimalValidSpec()
            spec.TableName = ""
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when table_name is too short", func() {
            spec := minimalValidSpec()
            spec.TableName = "ab"
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when attribute_definitions are empty", func() {
            spec := minimalValidSpec()
            spec.AttributeDefinitions = nil
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when key_schema has more than two elements", func() {
            spec := minimalValidSpec()
            spec.KeySchema = []*KeySchemaElement{
                {AttributeName: "pk", KeyType: KeyType_HASH},
                {AttributeName: "sk1", KeyType: KeyType_RANGE},
                {AttributeName: "sk2", KeyType: KeyType_RANGE},
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when billing_mode is unspecified", func() {
            spec := minimalValidSpec()
            spec.BillingMode = BillingMode_BILLING_MODE_UNSPECIFIED
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })

    // -------------------------------------------------------------------------
    //  Negative cases – cross-field CEL rules
    // -------------------------------------------------------------------------

    Context("cross-field validation failures", func() {
        It("fails when billing_mode=PROVISIONED but throughput is missing", func() {
            spec := minimalValidSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when billing_mode=PROVISIONED but capacity units are zero", func() {
            spec := minimalValidSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            spec.ProvisionedThroughput = &ProvisionedThroughput{
                ReadCapacityUnits:  0,
                WriteCapacityUnits: 0,
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when billing_mode=PAY_PER_REQUEST and throughput has positive units", func() {
            spec := minimalValidSpec()
            spec.ProvisionedThroughput = &ProvisionedThroughput{
                ReadCapacityUnits:  1,
                WriteCapacityUnits: 1,
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })
})
