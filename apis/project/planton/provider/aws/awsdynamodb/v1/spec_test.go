//go:build unit

package awsdynamodbv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
)

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    var validator protovalidate.Validator

    BeforeSuite(func() {
        // Instantiate the validator as required.
        validatorPtr, err := protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
        validator = *validatorPtr
    })

    // Helper producing a fully-valid base spec that callers can tweak.
    newBaseSpec := func() *AwsDynamodbSpec {
        return &AwsDynamodbSpec{
            TableName:        "MyTable_123",
            BillingMode:       BillingMode_PROVISIONED,
            ReadCapacityUnits: 5,
            WriteCapacityUnits: 10,
            AttributeDefinitions: []*AttributeDefinition{
                {AttributeName: "PK", AttributeType: AttributeType_STRING},
            },
            KeySchema: []*KeySchemaElement{
                {AttributeName: "PK", KeyType: KeyType_HASH},
            },
            GlobalSecondaryIndexes: []*GlobalSecondaryIndex{
                {
                    IndexName: "GSI1",
                    KeySchema: []*KeySchemaElement{
                        {AttributeName: "PK", KeyType: KeyType_HASH},
                    },
                    Projection: &Projection{ProjectionType: ProjectionType_ALL},
                    ReadCapacityUnits: 1,
                    WriteCapacityUnits: 1,
                },
            },
        }
    }

    Context("PROVISIONED billing mode", func() {
        It("allows positive capacity values", func() {
            spec := newBaseSpec()
            Expect(validator.Validate(spec)).To(Succeed())
        })

        It("rejects zero read capacity", func() {
            spec := newBaseSpec()
            spec.ReadCapacityUnits = 0
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("read/write capacity"))
        })

        It("rejects GSI capacity units of zero", func() {
            spec := newBaseSpec()
            spec.GlobalSecondaryIndexes[0].ReadCapacityUnits = 0
            spec.GlobalSecondaryIndexes[0].WriteCapacityUnits = 0
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("GSI capacity units must be > 0"))
        })
    })

    Context("PAY_PER_REQUEST billing mode", func() {
        It("requires table-level capacity to be zero", func() {
            spec := newBaseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            spec.ReadCapacityUnits = 1
            spec.WriteCapacityUnits = 1
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("read/write capacity must be 0"))
        })

        It("requires GSI capacities to be zero", func() {
            spec := newBaseSpec()
            spec.BillingMode = BillingMode_PAY_PER_REQUEST
            spec.ReadCapacityUnits = 0
            spec.WriteCapacityUnits = 0
            spec.GlobalSecondaryIndexes[0].ReadCapacityUnits = 2
            spec.GlobalSecondaryIndexes[0].WriteCapacityUnits = 2
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("GSI capacity units must be 0"))
        })
    })

    Context("field-level validation", func() {
        It("rejects table names that are too short", func() {
            spec := newBaseSpec()
            spec.TableName = "ab" // min_len is 3
            err := validator.Validate(spec)
            Expect(err).To(HaveOccurred())
            Expect(err.Error()).To(ContainSubstring("value length must be at least 3"))
        })
    })
})
