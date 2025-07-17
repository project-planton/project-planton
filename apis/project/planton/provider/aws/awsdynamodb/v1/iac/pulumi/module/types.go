package module

// This file provides lightweight Go representations of the main proto messages
// found in project/planton/provider/aws/awsdynamodb/v1. They are **not** meant
// to be 1-to-1 replacements for the generated pb types – the canonical
// contract remains in protobuf. Instead, they offer a more ergonomic surface
// tailored to this Pulumi implementation (e.g. string-backed enums matching
// the AWS provider input expectations).
//
// Because these helpers live entirely inside the Pulumi stack code, we keep
// them minimal and self-contained: only the fields actually required while
// composing AWS resources are included.

// -------------------------
// Enumerations (string-backed)
// -------------------------

// AttributeType represents the scalar DynamoDB attribute types expected by the
// aws.dynamodb.Table resource ("S", "N", "B").
// See https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html
// for their meaning.
//
// NOTE: Using string constants allows us to pass the values directly into the
// Pulumi AWS provider without additional conversion.
type AttributeType string

const (
    AttributeTypeString AttributeType = "S"
    AttributeTypeNumber AttributeType = "N"
    AttributeTypeBinary AttributeType = "B"
)

// KeyType identifies whether a key element is a HASH (partition) or RANGE
// (sort) key.
type KeyType string

const (
    KeyTypeHash  KeyType = "HASH"
    KeyTypeRange KeyType = "RANGE"
)

// BillingMode determines how read/write capacity is billed.
type BillingMode string

const (
    BillingModeProvisioned   BillingMode = "PROVISIONED"
    BillingModePayPerRequest BillingMode = "PAY_PER_REQUEST"
)

// ProjectionType decides which attributes are projected into an index.
type ProjectionType string

const (
    ProjectionTypeAll       ProjectionType = "ALL"
    ProjectionTypeKeysOnly  ProjectionType = "KEYS_ONLY"
    ProjectionTypeInclude   ProjectionType = "INCLUDE"
)

// StreamViewType controls the information written to a DynamoDB Stream.
type StreamViewType string

const (
    StreamViewTypeNewImage        StreamViewType = "NEW_IMAGE"
    StreamViewTypeOldImage        StreamViewType = "OLD_IMAGE"
    StreamViewTypeNewAndOldImages StreamViewType = "NEW_AND_OLD_IMAGES"
    StreamViewTypeKeysOnly        StreamViewType = "KEYS_ONLY"
)

// SSEType selects the server-side encryption mechanism.
type SSEType string

const (
    SSETypeAES256 SSEType = "AES256"
    SSETypeKMS    SSEType = "KMS"
)

// -------------------------
// Helper Structs mirroring proto definitions
// -------------------------

// AttributeDefinition describes a single attribute that appears in a key
// schema or index.
type AttributeDefinition struct {
    Name string        `json:"attribute_name"`
    Type AttributeType `json:"attribute_type"`
}

// KeySchemaElement identifies the role of an attribute in a primary or
// secondary index key schema.
type KeySchemaElement struct {
    AttributeName string  `json:"attribute_name"`
    KeyType       KeyType `json:"key_type"`
}

// ProvisionedThroughput represents fixed RCU/WCU capacity numbers.
type ProvisionedThroughput struct {
    ReadCapacityUnits  int64 `json:"read_capacity_units"`
    WriteCapacityUnits int64 `json:"write_capacity_units"`
}

// Projection controls the attribute projection behaviour for an index.
type Projection struct {
    ProjectionType   ProjectionType `json:"projection_type"`
    NonKeyAttributes []string       `json:"non_key_attributes,omitempty"`
}

// GlobalSecondaryIndex config for a GSI.
type GlobalSecondaryIndex struct {
    Name                string                 `json:"index_name"`
    KeySchema           []KeySchemaElement     `json:"key_schema"`
    Projection          Projection             `json:"projection"`
    ProvisionedThroughput *ProvisionedThroughput `json:"provisioned_throughput,omitempty"`
}

// LocalSecondaryIndex config for an LSI.
type LocalSecondaryIndex struct {
    Name       string             `json:"index_name"`
    KeySchema  []KeySchemaElement `json:"key_schema"`
    Projection Projection         `json:"projection"`
}

// StreamSpecification enables DynamoDB Streams on a table.
type StreamSpecification struct {
    Enabled      bool            `json:"stream_enabled"`
    StreamViewType StreamViewType `json:"stream_view_type,omitempty"`
}

// TimeToLiveSpecification configures item TTL expiration.
type TimeToLiveSpecification struct {
    Enabled        bool   `json:"ttl_enabled"`
    AttributeName  string `json:"attribute_name,omitempty"`
}

// SSESpecification contains server-side encryption settings.
type SSESpecification struct {
    Enabled        bool    `json:"enabled"`
    SSEType        SSEType `json:"sse_type,omitempty"`
    KmsMasterKeyID string  `json:"kms_master_key_id,omitempty"`
}

// AwsDynamodbSpec aggregates all configuration knobs for a DynamoDB table.
// It purposefully mirrors only what the Pulumi program consumes.
type AwsDynamodbSpec struct {
    TableName             string                  `json:"table_name"`
    AttributeDefinitions  []AttributeDefinition   `json:"attribute_definitions"`
    KeySchema             []KeySchemaElement      `json:"key_schema"`
    BillingMode           BillingMode             `json:"billing_mode"`
    ProvisionedThroughput *ProvisionedThroughput  `json:"provisioned_throughput,omitempty"`
    GlobalSecondaryIndexes []GlobalSecondaryIndex `json:"global_secondary_indexes,omitempty"`
    LocalSecondaryIndexes  []LocalSecondaryIndex  `json:"local_secondary_indexes,omitempty"`
    StreamSpecification   *StreamSpecification    `json:"stream_specification,omitempty"`
    TTLSpecification      *TimeToLiveSpecification `json:"ttl_specification,omitempty"`
    SseSpecification      *SSESpecification        `json:"sse_specification,omitempty"`
    PointInTimeRecovery   bool                    `json:"point_in_time_recovery_enabled"`
    Tags                  map[string]string       `json:"tags,omitempty"`
}

// -------------------------
// Output helper structs
// -------------------------

// Stream encapsulates the identifiers returned when DynamoDB Streams are
// enabled on a table.
type Stream struct {
    Arn   string `json:"stream_arn"`
    Label string `json:"stream_label"`
}

// AwsDynamodbStackOutputs mirrors the proto outputs message so that our Pulumi
// program can marshal data before calling ctx.Export.
type AwsDynamodbStackOutputs struct {
    TableArn                   string   `json:"table_arn"`
    TableName                  string   `json:"table_name"`
    TableID                    string   `json:"table_id"`
    Stream                     *Stream  `json:"stream,omitempty"`
    KmsKeyArn                  string   `json:"kms_key_arn,omitempty"`
    GlobalSecondaryIndexNames  []string `json:"global_secondary_index_names,omitempty"`
    LocalSecondaryIndexNames   []string `json:"local_secondary_index_names,omitempty"`
}
