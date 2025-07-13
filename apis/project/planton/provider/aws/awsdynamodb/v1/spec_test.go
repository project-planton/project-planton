package v1

import (
    "context"
    "testing"

    "github.com/bufbuild/protovalidate-go"

    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    awsdynamodb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = Describe("AwsDynamodbSpec validation", func() {
    var (
        validator protovalidate.Validator
        err       error
    )

    BeforeEach(func() {
        validator, err = protovalidate.New()
        Expect(err).NotTo(HaveOccurred())
    })

    // helper to create a valid spec
    newValidSpec := func() *awsdynamodb.AwsDynamodbSpec {
        return &awsdynamodb.AwsDynamodbSpec{
            TableName: "myTable",
            HashKey:   "id",
            RangeKey:  "sort",
            Attributes: []*awsdynamodb.AttributeDefinition{
                {Name: "id", Type: awsdynamodb.AttributeType_STRING},
                {Name: "sort", Type: awsdynamodb.AttributeType_STRING},
            },
            BillingMode:   awsdynamodb.BillingMode_PROVISIONED,
            ReadCapacity:  5,
            WriteCapacity: 5,
            StreamEnabled: false,
            TtlEnabled:    false,
            TtlAttributeName: "expires_at",
            TableClass:       awsdynamodb.TableClass_STANDARD,
        }
    }

    It("accepts a fully valid spec", func() {
        spec := newValidSpec()
        Expect(validator.Validate(context.Background(), spec)).To(Succeed())
    })

    It("rejects a table name that is too short", func() {
        spec := newValidSpec()
        spec.TableName = "ab" // min_len = 3
        Expect(validator.Validate(context.Background(), spec)).To(HaveOccurred())
    })

    It("rejects when no attributes are provided", func() {
        spec := newValidSpec()
        spec.Attributes = nil // min_items = 1
        Expect(validator.Validate(context.Background(), spec)).To(HaveOccurred())
    })

    It("rejects an unspecified billing mode", func() {
        spec := newValidSpec()
        spec.BillingMode = awsdynamodb.BillingMode_BILLING_MODE_UNSPECIFIED // not_in = [0]
        Expect(validator.Validate(context.Background(), spec)).To(HaveOccurred())
    })

    It("rejects read capacity below the minimum", func() {
        spec := newValidSpec()
        spec.ReadCapacity = 0 // gte = 1
        Expect(validator.Validate(context.Background(), spec)).To(HaveOccurred())
    })
})