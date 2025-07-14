package awsdynamodbv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

// -----------------------------------------------------------------------------
// Test Suite bootstrap
// -----------------------------------------------------------------------------

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(ginkgo.Fail)
    ginkgo.RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

// -----------------------------------------------------------------------------
// Suite set-up
// -----------------------------------------------------------------------------

var validator protovalidate.Validator

var _ = ginkgo.BeforeSuite(func() {
    vPtr, err := protovalidate.New()
    Expect(err).NotTo(HaveOccurred())
    validator = *vPtr
})

// -----------------------------------------------------------------------------
// Helper builders
// -----------------------------------------------------------------------------

func baseSpec() *AwsDynamodbSpec {
    return &AwsDynamodbSpec{
        TableName: "my_table",
        AttributeDefinitions: []*AttributeDefinition{
            {
                AttributeName: "id",
                AttributeType: AttributeType_STRING,
            },
        },
        KeySchema: []*KeySchemaElement{
            {
                AttributeName: "id",
                KeyType:        KeyType_HASH,
            },
        },
        Tags: []*Tag{},
    }
}

// -----------------------------------------------------------------------------
// Validation tests
// -----------------------------------------------------------------------------

var _ = ginkgo.Describe("AwsDynamodbSpec", func() {

    ginkgo.Context("BillingMode = PROVISIONED", func() {
        ginkgo.It("succeeds when provisioned_throughput is positive", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            spec.ProvisionedThroughput = &ProvisionedThroughput{
                ReadCapacityUnits:  5,
                WriteCapacityUnits: 5,
            }

            err := validator.Validate(spec)
            Expect(err).NotTo(HaveOccurred())
        })

        ginkgo.It("fails when provisioned_throughput is missing", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            // no throughput

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("provisioned_throughput"))
        })

        ginkgo.It("fails when provisioned_throughput is zero", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            spec.ProvisionedThroughput = &ProvisionedThroughput{}

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("capacity units"))
        })
    })

    ginkgo.Context("BillingMode = PAY_PER_REQUEST", func() {
        ginkgo.It("succeeds when provisioned_throughput is unset", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST

            err := validator.Validate(spec)
            Expect(err).NotTo(HaveOccurred())
        })

        ginkgo.It("succeeds when provisioned_throughput has zero units", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            spec.ProvisionedThroughput = &ProvisionedThroughput{}

            err := validator.Validate(spec)
            Expect(err).NotTo(HaveOccurred())
        })

        ginkgo.It("fails when provisioned_throughput is positive", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            spec.ProvisionedThroughput = &ProvisionedThroughput{
                ReadCapacityUnits:  1,
                WriteCapacityUnits: 1,
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("provisioned_throughput must be unset"))
        })
    })

    ginkgo.Context("GlobalSecondaryIndex capacity rules", func() {
        ginkgo.It("succeeds for PROVISIONED when each GSI has positive throughput", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            spec.ProvisionedThroughput = &ProvisionedThroughput{ReadCapacityUnits: 5, WriteCapacityUnits: 5}
            spec.GlobalSecondaryIndexes = []*GlobalSecondaryIndex{
                {
                    IndexName: "gsi1",
                    KeySchema: []*KeySchemaElement{{AttributeName: "id", KeyType: KeyType_HASH}},
                    Projection: &Projection{ProjectionType: ProjectionType_ALL},
                    ProvisionedThroughput: &ProvisionedThroughput{ReadCapacityUnits: 1, WriteCapacityUnits: 1},
                },
            }

            err := validator.Validate(spec)
            Expect(err).NotTo(HaveOccurred())
        })

        ginkgo.It("fails for PROVISIONED when a GSI lacks throughput", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PROVISIONED
            spec.ProvisionedThroughput = &ProvisionedThroughput{ReadCapacityUnits: 5, WriteCapacityUnits: 5}
            spec.GlobalSecondaryIndexes = []*GlobalSecondaryIndex{
                {
                    IndexName:  "gsi1",
                    KeySchema:  []*KeySchemaElement{{AttributeName: "id", KeyType: KeyType_HASH}},
                    Projection: &Projection{ProjectionType: ProjectionType_ALL},
                    // Missing ProvisionedThroughput
                },
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("Global secondary index capacity units must follow table billing_mode rules"))
        })

        ginkgo.It("fails for PAY_PER_REQUEST when a GSI has positive throughput", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            spec.GlobalSecondaryIndexes = []*GlobalSecondaryIndex{
                {
                    IndexName: "gsi1",
                    KeySchema: []*KeySchemaElement{{AttributeName: "id", KeyType: KeyType_HASH}},
                    Projection: &Projection{ProjectionType: ProjectionType_ALL},
                    ProvisionedThroughput: &ProvisionedThroughput{ReadCapacityUnits: 1, WriteCapacityUnits: 1},
                },
            }

            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("Global secondary index capacity units must follow table billing_mode rules"))
        })

        ginkgo.It("succeeds for PAY_PER_REQUEST when GSI throughput is unset", func() {
            spec := baseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            spec.GlobalSecondaryIndexes = []*GlobalSecondaryIndex{
                {
                    IndexName:  "gsi1",
                    KeySchema:  []*KeySchemaElement{{AttributeName: "id", KeyType: KeyType_HASH}},
                    Projection: &Projection{ProjectionType: ProjectionType_ALL},
                    // No throughput set
                },
            }

            err := validator.Validate(spec)
            Expect(err).NotTo(HaveOccurred())
        })
    })
})
