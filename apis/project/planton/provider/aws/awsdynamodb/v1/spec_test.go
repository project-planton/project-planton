package awsdynamodbv1

import (
    "testing"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "github.com/bufbuild/protovalidate-go"
    v1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var validator protovalidate.Validator

var _ = BeforeSuite(func() {
    v, err := protovalidate.New()
    Expect(err).NotTo(HaveOccurred())
    validator = *v
})

// helper that returns a minimally-valid AwsDynamodbSpec. Individual tests may
// mutate the returned instance to exercise specific validation paths.
func baseSpec() *v1.AwsDynamodbSpec {
    return &v1.AwsDynamodbSpec{
        TableName: "my_table",
        KeySchema: []*v1.KeySchemaElement{
            {
                AttributeName: "id",
                KeyType:      v1.KeyType_HASH,
            },
        },
        AttributeDefinitions: []*v1.AttributeDefinition{
            {
                Name: "id",
                Type: v1.AttributeType_S,
            },
        },
    }
}

// helper to create a simple projection of type ALL (required by GSIs / LSIs)
func allProjection() *v1.Projection {
    return &v1.Projection{ProjectionType: v1.ProjectionType_ALL}
}

// helper to create a valid ProvisionedThroughput instance with the provided
// capacity values.
func throughput(rcu, wcu int64) *v1.ProvisionedThroughput {
    return &v1.ProvisionedThroughput{
        ReadCapacityUnits:  rcu,
        WriteCapacityUnits: wcu,
    }
}

var _ = Describe("AwsDynamodbSpec billing_mode_capacities rule", func() {
    Context("billing_mode == PROVISIONED", func() {
        It("passes when provisioned_throughput and every GSI throughput are > 0", func() {
            spec := baseSpec()
            spec.BillingMode = v1.BillingMode_PROVISIONED
            spec.ProvisionedThroughput = throughput(5, 5)
            spec.GlobalSecondaryIndexes = []*v1.GlobalSecondaryIndex{
                {
                    Name:                 "gsi1",
                    KeySchema:            spec.KeySchema,
                    Projection:           allProjection(),
                    ProvisionedThroughput: throughput(3, 3),
                },
            }
            Expect(validator.Validate(spec)).To(Succeed())
        })

        It("fails when table provisioned_throughput is missing", func() {
            spec := baseSpec()
            spec.BillingMode = v1.BillingMode_PROVISIONED
            // spec.ProvisionedThroughput intentionally nil
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when a GSI lacks provisioned_throughput", func() {
            spec := baseSpec()
            spec.BillingMode = v1.BillingMode_PROVISIONED
            spec.ProvisionedThroughput = throughput(5, 5)
            spec.GlobalSecondaryIndexes = []*v1.GlobalSecondaryIndex{
                {
                    Name:       "gsi_no_throughput",
                    KeySchema:  spec.KeySchema,
                    Projection: allProjection(),
                    // ProvisionedThroughput intentionally nil
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })

    Context("billing_mode == PAY_PER_REQUEST", func() {
        It("passes when no throughput is set on table or GSIs", func() {
            spec := baseSpec()
            spec.BillingMode = v1.BillingMode_PAY_PER_REQUEST
            // ProvisionedThroughput nil, GSI without throughput
            spec.GlobalSecondaryIndexes = []*v1.GlobalSecondaryIndex{
                {
                    Name:       "gsi_on_demand",
                    KeySchema:  spec.KeySchema,
                    Projection: allProjection(),
                    // no ProvisionedThroughput
                },
            }
            Expect(validator.Validate(spec)).To(Succeed())
        })

        It("fails when table provisioned_throughput is specified", func() {
            spec := baseSpec()
            spec.BillingMode = v1.BillingMode_PAY_PER_REQUEST
            spec.ProvisionedThroughput = throughput(1, 1)
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })

        It("fails when a GSI has provisioned_throughput specified", func() {
            spec := baseSpec()
            spec.BillingMode = v1.BillingMode_PAY_PER_REQUEST
            spec.GlobalSecondaryIndexes = []*v1.GlobalSecondaryIndex{
                {
                    Name:                 "gsi_with_throughput",
                    KeySchema:            spec.KeySchema,
                    Projection:           allProjection(),
                    ProvisionedThroughput: throughput(1, 1),
                },
            }
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })

    Context("billing_mode == UNSPECIFIED", func() {
        It("always fails the CEL rule", func() {
            spec := baseSpec()
            // billing_mode defaults to UNSPECIFIED (0)
            Expect(validator.Validate(spec)).To(HaveOccurred())
        })
    })
})
