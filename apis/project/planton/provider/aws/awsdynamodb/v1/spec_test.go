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
			BillingMode: AwsDynamodbSpec_BILLING_MODE_PAY_PER_REQUEST,
			AttributeDefinitions: []*AwsDynamodbSpec_AttributeDefinition{
				{Name: "pk", Type: AwsDynamodbSpec_ATTRIBUTE_TYPE_S},
			},
			KeySchema: []*AwsDynamodbSpec_KeySchemaElement{
				{AttributeName: "pk", KeyType: AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_HASH},
			},
		}
	})

	It("accepts a minimal valid PAY_PER_REQUEST table", func() {
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	It("accepts a valid PROVISIONED table with throughput", func() {
		spec.BillingMode = AwsDynamodbSpec_BILLING_MODE_PROVISIONED
		spec.ProvisionedThroughput = &AwsDynamodbSpec_ProvisionedThroughput{ReadCapacityUnits: 5, WriteCapacityUnits: 5}
		err := protovalidate.Validate(spec)
		Expect(err).To(BeNil())
	})

	// Field-level validations
	It("fails when attribute_definitions is empty", func() {
		spec.AttributeDefinitions = []*AwsDynamodbSpec_AttributeDefinition{}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when key_schema is empty", func() {
		spec.AttributeDefinitions = []*AwsDynamodbSpec_AttributeDefinition{{Name: "pk", Type: AwsDynamodbSpec_ATTRIBUTE_TYPE_S}}
		spec.KeySchema = []*AwsDynamodbSpec_KeySchemaElement{}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when attribute name is empty", func() {
		spec.AttributeDefinitions = []*AwsDynamodbSpec_AttributeDefinition{{Name: "", Type: AwsDynamodbSpec_ATTRIBUTE_TYPE_S}}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when enum values are undefined", func() {
		spec.AttributeDefinitions = []*AwsDynamodbSpec_AttributeDefinition{{Name: "pk", Type: AwsDynamodbSpec_AttributeType(99)}}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// CEL: billing mode must be non-zero
	It("fails when billing_mode is UNSPECIFIED", func() {
		spec.BillingMode = AwsDynamodbSpec_BILLING_MODE_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// CEL: PROVISIONED requires throughput > 0
	It("fails when PROVISIONED without throughput", func() {
		spec.BillingMode = AwsDynamodbSpec_BILLING_MODE_PROVISIONED
		spec.ProvisionedThroughput = nil
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when PAY_PER_REQUEST has throughput set > 0", func() {
		spec.BillingMode = AwsDynamodbSpec_BILLING_MODE_PAY_PER_REQUEST
		spec.ProvisionedThroughput = &AwsDynamodbSpec_ProvisionedThroughput{ReadCapacityUnits: 1, WriteCapacityUnits: 1}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// CEL: key schema shape
	It("fails when key_schema has two HASH keys", func() {
		spec.KeySchema = []*AwsDynamodbSpec_KeySchemaElement{
			{AttributeName: "pk", KeyType: AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_HASH},
			{AttributeName: "sk", KeyType: AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_HASH},
		}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// Projection INCLUDE requires non_key_attributes
	It("fails when projection type is INCLUDE without attributes", func() {
		gsi := &AwsDynamodbSpec_GlobalSecondaryIndex{
			Name:       "gsi1",
			KeySchema:  []*AwsDynamodbSpec_KeySchemaElement{{AttributeName: "pk", KeyType: AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_HASH}},
			Projection: &AwsDynamodbSpec_Projection{Type: AwsDynamodbSpec_PROJECTION_TYPE_INCLUDE, NonKeyAttributes: []string{}},
		}
		spec.GlobalSecondaryIndexes = []*AwsDynamodbSpec_GlobalSecondaryIndex{gsi}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// Streams rule
	It("fails when stream_enabled is true but view type is UNSPECIFIED", func() {
		spec.StreamEnabled = true
		spec.StreamViewType = AwsDynamodbSpec_STREAM_VIEW_TYPE_UNSPECIFIED
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when stream_enabled is false but view type is set", func() {
		spec.StreamEnabled = false
		spec.StreamViewType = AwsDynamodbSpec_STREAM_VIEW_TYPE_KEYS_ONLY
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// TTL rule
	It("fails when TTL enabled without attribute_name", func() {
		spec.Ttl = &AwsDynamodbSpec_TimeToLive{Enabled: true, AttributeName: ""}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	It("fails when TTL disabled but attribute_name is set", func() {
		spec.Ttl = &AwsDynamodbSpec_TimeToLive{Enabled: false, AttributeName: "expiresAt"}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})

	// GSI throughput rule matches billing mode
	It("fails when PROVISIONED table has GSI without throughput", func() {
		spec.BillingMode = AwsDynamodbSpec_BILLING_MODE_PROVISIONED
		spec.ProvisionedThroughput = &AwsDynamodbSpec_ProvisionedThroughput{ReadCapacityUnits: 5, WriteCapacityUnits: 5}
		gsi := &AwsDynamodbSpec_GlobalSecondaryIndex{
			Name:                  "gsi1",
			KeySchema:             []*AwsDynamodbSpec_KeySchemaElement{{AttributeName: "pk", KeyType: AwsDynamodbSpec_KeySchemaElement_KEY_TYPE_HASH}},
			Projection:            &AwsDynamodbSpec_Projection{Type: AwsDynamodbSpec_PROJECTION_TYPE_KEYS_ONLY},
			ProvisionedThroughput: nil,
		}
		spec.GlobalSecondaryIndexes = []*AwsDynamodbSpec_GlobalSecondaryIndex{gsi}
		err := protovalidate.Validate(spec)
		Expect(err).NotTo(BeNil())
	})
})
