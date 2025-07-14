package awsdynamodbv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    "github.com/onsi/ginkgo/v2"
    "github.com/onsi/gomega"
)

var validator protovalidate.Validator

func TestAwsDynamodbSpecValidation(t *testing.T) {
    gomega.RegisterFailHandler(ginkgo.Fail)
    ginkgo.RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = ginkgo.BeforeSuite(func() {
    var err error
    validator, err = protovalidate.New()
    gomega.Expect(err).NotTo(gomega.HaveOccurred())
})

func makeValidSpec() *AwsDynamodbSpec {
    return &AwsDynamodbSpec{
        TableName: "my_table",
        Attributes: []*AttributeDefinition{
            {
                Name: "id",
                Type: AttributeType_STRING,
            },
        },
        PartitionKey: "id",
        BillingMode:  BillingMode_PAY_PER_REQUEST,
    }
}

var _ = ginkgo.Describe("AwsDynamodbSpec buf.validate rules", func() {
    ginkgo.It("accepts a fully valid spec", func() {
        spec := makeValidSpec()
        err := validator.Validate(spec)
        gomega.Expect(err).NotTo(gomega.HaveOccurred())
    })

    ginkgo.It("rejects when table_name is missing", func() {
        spec := makeValidSpec()
        spec.TableName = ""
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())
    })

    ginkgo.It("rejects when attributes are empty", func() {
        spec := makeValidSpec()
        spec.Attributes = []*AttributeDefinition{}
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())
    })

    ginkgo.It("rejects when partition_key is missing", func() {
        spec := makeValidSpec()
        spec.PartitionKey = ""
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())
    })

    ginkgo.It("rejects when billing_mode is unspecified", func() {
        spec := makeValidSpec()
        spec.BillingMode = BillingMode_BILLING_MODE_UNSPECIFIED
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())
    })

    ginkgo.It("enforces provisioned capacities when billing_mode is PROVISIONED", func() {
        spec := makeValidSpec()
        spec.BillingMode = BillingMode_PROVISIONED
        // Leave capacities at zero -> should fail
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())

        // Provide positive capacities -> should pass
        spec.ProvisionedReadCapacity = 5
        spec.ProvisionedWriteCapacity = 10
        err = validator.Validate(spec)
        gomega.Expect(err).NotTo(gomega.HaveOccurred())
    })

    ginkgo.It("requires capacities to be zero when billing_mode is PAY_PER_REQUEST", func() {
        spec := makeValidSpec()
        spec.ProvisionedReadCapacity = 5
        spec.ProvisionedWriteCapacity = 10
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())
    })

    ginkgo.It("validates stream consistency rule", func() {
        // stream_enabled true but stream_view_type unspecified -> fail
        spec := makeValidSpec()
        spec.StreamEnabled = true
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())

        // Set stream_view_type -> pass
        spec.StreamViewType = StreamViewType_NEW_IMAGE
        err = validator.Validate(spec)
        gomega.Expect(err).NotTo(gomega.HaveOccurred())

        // When stream_enabled false but type set -> fail
        spec2 := makeValidSpec()
        spec2.StreamViewType = StreamViewType_NEW_IMAGE
        err = validator.Validate(spec2)
        gomega.Expect(err).To(gomega.HaveOccurred())
    })

    ginkgo.It("requires sort_key when local secondary indexes exist", func() {
        spec := makeValidSpec()
        spec.LocalSecondaryIndexes = []*LocalSecondaryIndex{
            {
                Name:    "lsi1",
                SortKey: "sk",
                Projection: &Projection{
                    Type: ProjectionType_ALL,
                },
            },
        }
        // sort_key missing in table spec -> should fail
        err := validator.Validate(spec)
        gomega.Expect(err).To(gomega.HaveOccurred())

        // Add sort_key -> should pass
        spec.SortKey = "sk"
        err = validator.Validate(spec)
        gomega.Expect(err).NotTo(gomega.HaveOccurred())
    })
})
