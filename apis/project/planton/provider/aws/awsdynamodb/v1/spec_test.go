package awsdynamodbv1_test

import (
    "testing"

    pv "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"
    "google.golang.org/protobuf/proto"

    pb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

func TestAwsDynamodbSpecValidation(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsDynamodbSpec validation suite")
}

var _ = Describe("AwsDynamodbSpec protovalidate rules", func() {
    var (
        validator *pv.Validator
        err       error
    )

    BeforeSuite(func() {
        validator, err = pv.New()
        Expect(err).NotTo(HaveOccurred())
    })

    newValidSpec := func() *pb.AwsDynamodbSpec {
        return &pb.AwsDynamodbSpec{
            TableName: "my_table",
            AttributeDefinitions: []*pb.AttributeDefinition{
                {Name: "pk", Type: pb.ScalarAttributeType_S},
            },
            HashKey:        "pk",
            RangeKey:       "sk",
            BillingMode:    pb.BillingMode_PAY_PER_REQUEST,
            ReadCapacity:   5,
            WriteCapacity:  5,
            Tags:           map[string]string{"env": "prod"},
            TableClass:     pb.TableClass_STANDARD,
            GlobalSecondaryIndexes: []*pb.GlobalSecondaryIndex{
                {
                    Name:       "gsi1",
                    HashKey:    "pk",
                    Projection: &pb.Projection{Type: pb.Projection_ALL},
                },
            },
        }
    }

    It("accepts a fully valid spec", func() {
        spec := newValidSpec()
        Expect(validator.Validate(spec)).To(Succeed())
    })

    It("rejects a table name that is too short", func() {
        spec := proto.Clone(newValidSpec()).(*pb.AwsDynamodbSpec)
        spec.TableName = "ab" // < min_len 3
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when attribute_definitions is empty", func() {
        spec := proto.Clone(newValidSpec()).(*pb.AwsDynamodbSpec)
        spec.AttributeDefinitions = nil
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when billing_mode is unspecified (0)", func() {
        spec := proto.Clone(newValidSpec()).(*pb.AwsDynamodbSpec)
        spec.BillingMode = pb.BillingMode_BILLING_MODE_UNSPECIFIED
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when read_capacity is 0", func() {
        spec := proto.Clone(newValidSpec()).(*pb.AwsDynamodbSpec)
        spec.ReadCapacity = 0
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })

    It("rejects when tags contain empty key", func() {
        spec := proto.Clone(newValidSpec()).(*pb.AwsDynamodbSpec)
        spec.Tags = map[string]string{"": "value"}
        Expect(validator.Validate(spec)).To(HaveOccurred())
    })
})
