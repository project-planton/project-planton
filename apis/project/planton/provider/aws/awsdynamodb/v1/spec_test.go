package awsdynamodbv1

import (
    "context"
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    awsdbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
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
    ctx := context.Background()

    It("accepts a valid PROVISIONED spec", func() {
        spec := validProvisionedSpec()
        Expect(validator.Validate(ctx, spec)).To(Succeed())
    })

    It("rejects PROVISIONED spec without table provisioned throughput", func() {
        spec := validProvisionedSpec()
        spec.ProvisionedThroughput = nil // remove table capacity
        Expect(validator.Validate(ctx, spec)).To(Not(Succeed()))
    })

    It("rejects PROVISIONED spec with a GSI missing provisioned throughput", func() {
        spec := validProvisionedSpec()
        spec.GlobalSecondaryIndexes[0].ProvisionedThroughput = nil // remove GSI capacity
        Expect(validator.Validate(ctx, spec)).To(Not(Succeed()))
    })

    It("accepts a valid PAY_PER_REQUEST spec", func() {
        spec := validPayPerRequestSpec()
        Expect(validator.Validate(ctx, spec)).To(Succeed())
    })

    It("rejects PAY_PER_REQUEST spec that sets provisioned throughput", func() {
        spec := validPayPerRequestSpec()
        spec.ProvisionedThroughput = &awsdbpb.ProvisionedThroughput{ReadCapacityUnits: 5, WriteCapacityUnits: 5}
        Expect(validator.Validate(ctx, spec)).To(Not(Succeed()))
    })
})

// Helper constructors -------------------------------------------------------

func validProvisionedSpec() *awsdbpb.AwsDynamodbSpec {
    return &awsdbpb.AwsDynamodbSpec{
        TableName: "mytable",
        AttributeDefinitions: []*awsdbpb.AttributeDefinition{
            {
                AttributeName: "id",
                AttributeType: awsdbpb.AttributeType_STRING,
            },
        },
        KeySchema: []*awsdbpb.KeySchemaElement{
            {
                AttributeName: "id",
                KeyType:        awsdbpb.KeyType_HASH,
            },
        },
        BillingMode: awsdbpb.BillingMode_PROVISIONED,
        ProvisionedThroughput: &awsdbpb.ProvisionedThroughput{
            ReadCapacityUnits:  5,
            WriteCapacityUnits: 5,
        },
        GlobalSecondaryIndexes: []*awsdbpb.GlobalSecondaryIndex{
            {
                IndexName: "gsi1",
                KeySchema: []*awsdbpb.KeySchemaElement{
                    {
                        AttributeName: "id",
                        KeyType:        awsdbpb.KeyType_HASH,
                    },
                },
                Projection: &awsdbpb.Projection{
                    ProjectionType: awsdbpb.ProjectionType_ALL,
                },
                ProvisionedThroughput: &awsdbpb.ProvisionedThroughput{
                    ReadCapacityUnits:  5,
                    WriteCapacityUnits: 5,
                },
            },
        },
    }
}

func validPayPerRequestSpec() *awsdbpb.AwsDynamodbSpec {
    spec := validProvisionedSpec()
    spec.BillingMode = awsdbpb.BillingMode_PAY_PER_REQUEST
    spec.ProvisionedThroughput = nil
    for _, g := range spec.GlobalSecondaryIndexes {
        g.ProvisionedThroughput = nil
    }
    return spec
}
