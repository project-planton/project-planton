package v1

import (
    "context"
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

// -----------------------------------------------------------------------------
//  Test Suite bootstrap
// -----------------------------------------------------------------------------
func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec buf.validate Suite")
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    var (
        validator protovalidate.Validator
        ctx       context.Context
    )

    // Build a fully-populated, buf.validate-compliant AwsDynamodbSpec that we
    // can mutate per test.
    baseSpec := func() *AwsDynamodbSpec {
        return &AwsDynamodbSpec{
            TableName:   "myTable",
            BillingMode: BillingMode_PROVISIONED,
            ProvisionedThroughput: &ProvisionedThroughput{
                ReadCapacityUnits:  5,
                WriteCapacityUnits: 5,
            },
            AttributeDefinitions: []*AttributeDefinition{{
                Name: "pk",
                Type: AttributeType_STRING,
            }},
            KeySchema: []*KeySchemaElement{{
                AttributeName: "pk",
                KeyType:       KeyType_HASH,
            }},
            GlobalSecondaryIndexes: []*GlobalSecondaryIndex{{
                Name: "gsi1",
                KeySchema: []*KeySchemaElement{
                    {AttributeName: "pk", KeyType: KeyType_HASH},
                },
                Projection: &Projection{ProjectionType: Projection_ALL},
                ProvisionedThroughput: &ProvisionedThroughput{
                    ReadCapacityUnits:  5,
                    WriteCapacityUnits: 5,
                },
            }},
        }
    }

    BeforeEach(func() {
        ptr, err := protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
        validator = *ptr // satisfy requirement: variable is non-pointer type
        ctx = context.Background()
    })

    It("accepts a valid PROVISIONED spec", func() {
        spec := baseSpec()
        Expect(validator.Validate(ctx, spec)).To(Succeed())
    })

    It("fails when PROVISIONED but table throughput is missing", func() {
        spec := baseSpec()
        spec.ProvisionedThroughput = nil // violates CEL rule
        err := validator.Validate(ctx, spec)
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(ContainSubstring("capacity fields must align with billing_mode"))
    })

    It("fails when PROVISIONED but GSI throughput is missing", func() {
        spec := baseSpec()
        spec.GlobalSecondaryIndexes[0].ProvisionedThroughput = nil
        err := validator.Validate(ctx, spec)
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(ContainSubstring("capacity fields must align with billing_mode"))
    })

    It("accepts a PAY_PER_REQUEST spec without throughput", func() {
        spec := baseSpec()
        spec.BillingMode = BillingMode_PAY_PER_REQUEST
        spec.ProvisionedThroughput = nil
        spec.GlobalSecondaryIndexes[0].ProvisionedThroughput = nil
        Expect(validator.Validate(ctx, spec)).To(Succeed())
    })

    It("fails when PAY_PER_REQUEST has non-zero throughput", func() {
        spec := baseSpec()
        spec.BillingMode = BillingMode_PAY_PER_REQUEST
        // Keep non-zero throughput ⇒ should error
        err := validator.Validate(ctx, spec)
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(ContainSubstring("capacity fields must align with billing_mode"))
    })

    It("fails when required repeated fields are empty", func() {
        spec := baseSpec()
        spec.AttributeDefinitions = nil // min_items = 1
        err := validator.Validate(ctx, spec)
        Expect(err).To(HaveOccurred())
        Expect(err.Error()).To(ContainSubstring("attribute_definitions"))
    })
})
