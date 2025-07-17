package awsdynamodb

// Code generated manually to bridge Pulumi-facing string enums with the
// protobuf enums declared in
// github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1.
//
// All types below are thin string aliases that provide a developer friendly
// experience in Pulumi programs while still allowing the provider
// implementation to convert loss-lessly to/from the canonical protobuf
// representations.

import (
    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// -----------------------------------------------------------------------------
// AttributeType
// -----------------------------------------------------------------------------

type AttributeType string

const (
    // AttributeTypeUnspecified represents an absent / zero value. It is encoded
    // as an empty string because Pulumi uses zero values to model the absence
    // of user input.
    AttributeTypeUnspecified AttributeType = ""

    AttributeTypeString AttributeType = "S" // String attribute.
    AttributeTypeNumber AttributeType = "N" // Number attribute.
    AttributeTypeBinary AttributeType = "B" // Binary attribute.
)

// ToProto converts the Pulumi-side AttributeType into the generated protobuf
// enum.
func (t AttributeType) ToProto() awsdynamodbpb.AttributeType {
    switch t {
    case AttributeTypeString:
        return awsdynamodbpb.AttributeType_STRING
    case AttributeTypeNumber:
        return awsdynamodbpb.AttributeType_NUMBER
    case AttributeTypeBinary:
        return awsdynamodbpb.AttributeType_BINARY
    default:
        return awsdynamodbpb.AttributeType_ATTRIBUTE_TYPE_UNSPECIFIED
    }
}

// AttributeTypeFromProto creates a Pulumi-side AttributeType from the protobuf
// value.
func AttributeTypeFromProto(t awsdynamodbpb.AttributeType) AttributeType {
    switch t {
    case awsdynamodbpb.AttributeType_STRING:
        return AttributeTypeString
    case awsdynamodbpb.AttributeType_NUMBER:
        return AttributeTypeNumber
    case awsdynamodbpb.AttributeType_BINARY:
        return AttributeTypeBinary
    default:
        return AttributeTypeUnspecified
    }
}

// -----------------------------------------------------------------------------
// KeyType
// -----------------------------------------------------------------------------

type KeyType string

const (
    KeyTypeUnspecified KeyType = ""

    KeyTypeHash  KeyType = "HASH"  // Partition key.
    KeyTypeRange KeyType = "RANGE" // Sort/Range key.
)

func (t KeyType) ToProto() awsdynamodbpb.KeyType {
    switch t {
    case KeyTypeHash:
        return awsdynamodbpb.KeyType_HASH
    case KeyTypeRange:
        return awsdynamodbpb.KeyType_RANGE
    default:
        return awsdynamodbpb.KeyType_KEY_TYPE_UNSPECIFIED
    }
}

func KeyTypeFromProto(t awsdynamodbpb.KeyType) KeyType {
    switch t {
    case awsdynamodbpb.KeyType_HASH:
        return KeyTypeHash
    case awsdynamodbpb.KeyType_RANGE:
        return KeyTypeRange
    default:
        return KeyTypeUnspecified
    }
}

// -----------------------------------------------------------------------------
// BillingMode
// -----------------------------------------------------------------------------

type BillingMode string

const (
    BillingModeUnspecified   BillingMode = ""
    BillingModeProvisioned   BillingMode = "PROVISIONED"
    BillingModePayPerRequest BillingMode = "PAY_PER_REQUEST"
)

func (m BillingMode) ToProto() awsdynamodbpb.BillingMode {
    switch m {
    case BillingModeProvisioned:
        return awsdynamodbpb.BillingMode_PROVISIONED
    case BillingModePayPerRequest:
        return awsdynamodbpb.BillingMode_PAY_PER_REQUEST
    default:
        return awsdynamodbpb.BillingMode_BILLING_MODE_UNSPECIFIED
    }
}

func BillingModeFromProto(m awsdynamodbpb.BillingMode) BillingMode {
    switch m {
    case awsdynamodbpb.BillingMode_PROVISIONED:
        return BillingModeProvisioned
    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        return BillingModePayPerRequest
    default:
        return BillingModeUnspecified
    }
}

// -----------------------------------------------------------------------------
// ProjectionType
// -----------------------------------------------------------------------------

type ProjectionType string

const (
    ProjectionTypeUnspecified ProjectionType = ""
    ProjectionTypeAll         ProjectionType = "ALL"
    ProjectionTypeKeysOnly    ProjectionType = "KEYS_ONLY"
    ProjectionTypeInclude     ProjectionType = "INCLUDE"
)

func (p ProjectionType) ToProto() awsdynamodbpb.ProjectionType {
    switch p {
    case ProjectionTypeAll:
        return awsdynamodbpb.ProjectionType_ALL
    case ProjectionTypeKeysOnly:
        return awsdynamodbpb.ProjectionType_KEYS_ONLY
    case ProjectionTypeInclude:
        return awsdynamodbpb.ProjectionType_INCLUDE
    default:
        return awsdynamodbpb.ProjectionType_PROJECTION_TYPE_UNSPECIFIED
    }
}

func ProjectionTypeFromProto(p awsdynamodbpb.ProjectionType) ProjectionType {
    switch p {
    case awsdynamodbpb.ProjectionType_ALL:
        return ProjectionTypeAll
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        return ProjectionTypeKeysOnly
    case awsdynamodbpb.ProjectionType_INCLUDE:
        return ProjectionTypeInclude
    default:
        return ProjectionTypeUnspecified
    }
}

// -----------------------------------------------------------------------------
// StreamViewType
// -----------------------------------------------------------------------------

type StreamViewType string

const (
    StreamViewTypeUnspecified     StreamViewType = ""
    StreamViewTypeNewImage        StreamViewType = "NEW_IMAGE"
    StreamViewTypeOldImage        StreamViewType = "OLD_IMAGE"
    StreamViewTypeNewAndOldImages StreamViewType = "NEW_AND_OLD_IMAGES"
    StreamViewTypeKeysOnly        StreamViewType = "KEYS_ONLY"
)

func (s StreamViewType) ToProto() awsdynamodbpb.StreamViewType {
    switch s {
    case StreamViewTypeNewImage:
        return awsdynamodbpb.StreamViewType_NEW_IMAGE
    case StreamViewTypeOldImage:
        return awsdynamodbpb.StreamViewType_OLD_IMAGE
    case StreamViewTypeNewAndOldImages:
        return awsdynamodbpb.StreamViewType_NEW_AND_OLD_IMAGES
    case StreamViewTypeKeysOnly:
        return awsdynamodbpb.StreamViewType_STREAM_KEYS_ONLY
    default:
        return awsdynamodbpb.StreamViewType_STREAM_VIEW_TYPE_UNSPECIFIED
    }
}

func StreamViewTypeFromProto(s awsdynamodbpb.StreamViewType) StreamViewType {
    switch s {
    case awsdynamodbpb.StreamViewType_NEW_IMAGE:
        return StreamViewTypeNewImage
    case awsdynamodbpb.StreamViewType_OLD_IMAGE:
        return StreamViewTypeOldImage
    case awsdynamodbpb.StreamViewType_NEW_AND_OLD_IMAGES:
        return StreamViewTypeNewAndOldImages
    case awsdynamodbpb.StreamViewType_STREAM_KEYS_ONLY:
        return StreamViewTypeKeysOnly
    default:
        return StreamViewTypeUnspecified
    }
}

// -----------------------------------------------------------------------------
// SSEType
// -----------------------------------------------------------------------------

type SSEType string

const (
    SSETypeUnspecified SSEType = ""
    SSETypeAES256      SSEType = "AES256"
    SSETypeKMS         SSEType = "KMS"
)

func (s SSEType) ToProto() awsdynamodbpb.SSEType {
    switch s {
    case SSETypeAES256:
        return awsdynamodbpb.SSEType_AES256
    case SSETypeKMS:
        return awsdynamodbpb.SSEType_KMS
    default:
        return awsdynamodbpb.SSEType_SSE_TYPE_UNSPECIFIED
    }
}

func SSETypeFromProto(s awsdynamodbpb.SSEType) SSEType {
    switch s {
    case awsdynamodbpb.SSEType_AES256:
        return SSETypeAES256
    case awsdynamodbpb.SSEType_KMS:
        return SSETypeKMS
    default:
        return SSETypeUnspecified
    }
}
