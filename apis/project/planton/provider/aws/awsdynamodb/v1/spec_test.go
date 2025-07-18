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

var _ = Describe("AwsDynamodbSpec validation", func() {
    It("accepts a valid PROVISIONED spec", func() {
        spec := validProvisionedSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    It("rejects PROVISIONED spec without table provisioned throughput", func() {
        spec := validProvisionedSpec()
        spec.ProvisionedThroughput = nil // remove table capacity
        Expect(validator.Validate(spec)).To(Not(Succeed()))
    })

    It("rejects PROVISIONED spec with a GSI missing provisioned throughput", func() {
        spec := validProvisionedSpec()
        spec.GlobalSecondaryIndexes[0].ProvisionedThroughput = nil // remove GSI capacity
        Expect(validator.Validate(spec)).To(Not(Succeed()))
    })

    It("accepts a valid PAY_PER_REQUEST spec", func() {
        spec := validPayPerRequestSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    It("rejects PAY_PER_REQUEST spec that sets provisioned throughput", func() {
        spec := validPayPerRequestSpec()
        spec.ProvisionedThroughput = &ProvisionedThroughput{ReadCapacityUnits: 5, WriteCapacityUnits: 5}
        Expect(validator.Validate(spec)).To(Not(Succeed()))
    })
})

// Helper constructors -------------------------------------------------------

func validProvisionedSpec() *AwsDynamodbSpec {
    return &AwsDynamodbSpec{
        TableName: "mytable",
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
        BillingMode: BillingMode_PROVISIONED,
        ProvisionedThroughput: &ProvisionedThroughput{
            ReadCapacityUnits:  5,
            WriteCapacityUnits: 5,
        },
        GlobalSecondaryIndexes: []*GlobalSecondaryIndex{
            {
                IndexName: "gsi1",
                KeySchema: []*KeySchemaElement{
                    {
                        AttributeName: "id",
                        KeyType:        KeyType_HASH,
                    },
                },
                Projection: &Projection{
                    ProjectionType: ProjectionType_ALL,
                },
                ProvisionedThroughput: &ProvisionedThroughput{
                    ReadCapacityUnits:  5,
                    WriteCapacityUnits: 5,
                },
            },
        },
    }
}

func validPayPerRequestSpec() *AwsDynamodbSpec {
    spec := validProvisionedSpec()
    spec.BillingMode = BillingMode_PAY_PER_REQUEST
    spec.ProvisionedThroughput = nil
    for _, g := range spec.GlobalSecondaryIndexes {
        g.ProvisionedThroughput = nil
    }
    return spec
}
