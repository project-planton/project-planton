package awsdynamodbv1

import (
	"testing"

	"github.com/bufbuild/protovalidate-go"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAwsDynamodbSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "AwsDynamodbSpec Validation Suite")
}

var _ = Describe("AwsDynamodbSpec validations", func() {
	var spec *AwsDynamodbSpec

	BeforeEach(func() {
		spec = &AwsDynamodbSpec{
			TableName:                   "test-table",
			AwsRegion:                   "us-east-1",
			BillingMode:                 BillingMode_BILLING_MODE_PROVISIONED,
			PartitionKeyName:            "id",
			PartitionKeyType:            AttributeType_ATTRIBUTE_TYPE_STRING,
			ReadCapacityUnits:           5,
			WriteCapacityUnits:          5,
			PointInTimeRecoveryEnabled:  true,
			ServerSideEncryptionEnabled: true,
		}
	})

	It("accepts a valid DynamoDB spec with provisioned billing", func() {
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts a valid DynamoDB spec with pay-per-request billing", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PAY_PER_REQUEST
		spec.ReadCapacityUnits = 0
		spec.WriteCapacityUnits = 0
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts a valid DynamoDB spec with sort key", func() {
		spec.SortKeyName = "timestamp"
		spec.SortKeyType = AttributeType_ATTRIBUTE_TYPE_NUMBER
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts a valid DynamoDB spec with binary partition key", func() {
		spec.PartitionKeyType = AttributeType_ATTRIBUTE_TYPE_BINARY
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	// Field validation tests
	It("fails when table_name is empty", func() {
		spec.TableName = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when table_name is too short", func() {
		spec.TableName = "ab"
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when table_name is too long", func() {
		spec.TableName = "a" + string(make([]byte, 255))
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when aws_region is empty", func() {
		spec.AwsRegion = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when billing_mode is undefined", func() {
		spec.BillingMode = BillingMode(99)
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when partition_key_name is empty", func() {
		spec.PartitionKeyName = ""
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when partition_key_type is undefined", func() {
		spec.PartitionKeyType = AttributeType(99)
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when read_capacity_units is negative", func() {
		spec.ReadCapacityUnits = -1
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when write_capacity_units is negative", func() {
		spec.WriteCapacityUnits = -1
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// CEL expression tests
	It("fails when billing_mode is PROVISIONED but read_capacity_units is 0 (billing_mode_capacity_validation)", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PROVISIONED
		spec.ReadCapacityUnits = 0
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when billing_mode is PROVISIONED but write_capacity_units is 0 (billing_mode_capacity_validation)", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PROVISIONED
		spec.WriteCapacityUnits = 0
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when billing_mode is PAY_PER_REQUEST but read_capacity_units is greater than 0 (billing_mode_capacity_validation)", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PAY_PER_REQUEST
		spec.ReadCapacityUnits = 5
		spec.WriteCapacityUnits = 0
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when billing_mode is PAY_PER_REQUEST but write_capacity_units is greater than 0 (billing_mode_capacity_validation)", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PAY_PER_REQUEST
		spec.ReadCapacityUnits = 0
		spec.WriteCapacityUnits = 5
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("passes when billing_mode is PROVISIONED and capacity units are greater than 0 (billing_mode_capacity_validation)", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PROVISIONED
		spec.ReadCapacityUnits = 5
		spec.WriteCapacityUnits = 5
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("passes when billing_mode is PAY_PER_REQUEST and capacity units are 0 (billing_mode_capacity_validation)", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PAY_PER_REQUEST
		spec.ReadCapacityUnits = 0
		spec.WriteCapacityUnits = 0
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("fails when partition_key_name is set but partition_key_type is UNSPECIFIED (partition_key_required)", func() {
		spec.PartitionKeyName = "id"
		spec.PartitionKeyType = AttributeType_ATTRIBUTE_TYPE_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when partition_key_type is UNSPECIFIED (partition_key_required)", func() {
		spec.PartitionKeyType = AttributeType_ATTRIBUTE_TYPE_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("passes when partition_key_name is set and partition_key_type is valid (partition_key_required)", func() {
		spec.PartitionKeyName = "id"
		spec.PartitionKeyType = AttributeType_ATTRIBUTE_TYPE_STRING
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("fails when sort_key_name is set but sort_key_type is UNSPECIFIED (sort_key_consistency)", func() {
		spec.SortKeyName = "timestamp"
		spec.SortKeyType = AttributeType_ATTRIBUTE_TYPE_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when sort_key_name is empty but sort_key_type is not UNSPECIFIED (sort_key_consistency)", func() {
		spec.SortKeyName = ""
		spec.SortKeyType = AttributeType_ATTRIBUTE_TYPE_STRING
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("passes when sort_key_name is empty and sort_key_type is UNSPECIFIED (sort_key_consistency)", func() {
		spec.SortKeyName = ""
		spec.SortKeyType = AttributeType_ATTRIBUTE_TYPE_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("passes when sort_key_name is set and sort_key_type is valid (sort_key_consistency)", func() {
		spec.SortKeyName = "timestamp"
		spec.SortKeyType = AttributeType_ATTRIBUTE_TYPE_NUMBER
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	// Edge cases
	It("accepts valid table names with hyphens and underscores", func() {
		spec.TableName = "my-test_table-123"
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts valid partition key names with special characters", func() {
		spec.PartitionKeyName = "user_id_123"
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts valid sort key names with special characters", func() {
		spec.SortKeyName = "created_at_123"
		spec.SortKeyType = AttributeType_ATTRIBUTE_TYPE_NUMBER
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts zero capacity units for pay-per-request billing", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PAY_PER_REQUEST
		spec.ReadCapacityUnits = 0
		spec.WriteCapacityUnits = 0
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts high capacity units for provisioned billing", func() {
		spec.BillingMode = BillingMode_BILLING_MODE_PROVISIONED
		spec.ReadCapacityUnits = 1000
		spec.WriteCapacityUnits = 1000
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})
})
