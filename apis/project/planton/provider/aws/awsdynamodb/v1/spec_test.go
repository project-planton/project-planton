package v1

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "github.com/bufbuild/protovalidate-go"
)

var validator protovalidate.Validator

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = BeforeSuite(func() {
    v, err := protovalidate.New()
    Expect(err).NotTo(HaveOccurred())
    validator = *v
})

var _ = Describe("AwsDynamodbSpec buf.validate rules", func() {
    Context("with a valid PROVISIONED spec", func() {
        It("passes validation", func() {
            spec := &AwsDynamodbSpec{
                TableName: "my-table-001",
                Attributes: []*Attribute{
                    {Name: "UserId", Type: "S"},
                    {Name: "Timestamp", Type: "N"},
                },
                KeySchema: &KeySchema{
                    HashKey:  "UserId",
                    RangeKey: "Timestamp",
                },
                BillingMode: "PROVISIONED",
                ProvisionedThroughput: &ProvisionedThroughput{
                    ReadCapacityUnits:  5,
                    WriteCapacityUnits: 5,
                },
                GlobalSecondaryIndexes: []*GlobalSecondaryIndex{
                    {
                        Name:     "GSI1",
                        HashKey:  "Timestamp",
                        Projection: &Projection{
                            Type: "ALL",
                        },
                        ProvisionedThroughput: &ProvisionedThroughput{
                            ReadCapacityUnits:  5,
                            WriteCapacityUnits: 5,
                        },
                    },
                },
                LocalSecondaryIndexes: []*LocalSecondaryIndex{
                    {
                        Name:     "LSI1",
                        RangeKey: "Timestamp",
                        Projection: &Projection{
                            Type: "KEYS_ONLY",
                        },
                    },
                },
                TableClass: "STANDARD",
            }

            Expect(validator.Validate(spec)).To(Succeed())
        })
    })

    Context("with a valid PAY_PER_REQUEST spec", func() {
        It("passes validation", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "my-ondemand-table",
                BillingMode: "PAY_PER_REQUEST",
                Attributes: []*Attribute{
                    {Name: "PK", Type: "S"},
                },
                KeySchema: &KeySchema{
                    HashKey: "PK",
                },
                GlobalSecondaryIndexes: []*GlobalSecondaryIndex{
                    {
                        Name:    "GSI-OnDemand",
                        HashKey: "PK",
                        Projection: &Projection{
                            Type: "ALL",
                        },
                    },
                },
            }

            Expect(validator.Validate(spec)).To(Succeed())
        })
    })

    Context("with PROVISIONED billing mode but missing throughput", func() {
        It("fails validation", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "bad-provisioned-table",
                BillingMode: "PROVISIONED",
                Attributes: []*Attribute{
                    {Name: "PK", Type: "S"},
                },
                KeySchema: &KeySchema{HashKey: "PK"},
            }

            Expect(validator.Validate(spec)).To(Not(Succeed()))
        })
    })

    Context("with PAY_PER_REQUEST billing mode but throughput supplied", func() {
        It("fails validation", func() {
            spec := &AwsDynamodbSpec{
                TableName:   "bad-ondemand-table",
                BillingMode: "PAY_PER_REQUEST",
                Attributes: []*Attribute{
                    {Name: "PK", Type: "S"},
                },
                KeySchema: &KeySchema{HashKey: "PK"},
                ProvisionedThroughput: &ProvisionedThroughput{
                    ReadCapacityUnits:  1,
                    WriteCapacityUnits: 1,
                },
            }

            Expect(validator.Validate(spec)).To(Not(Succeed()))
        })
    })
})
