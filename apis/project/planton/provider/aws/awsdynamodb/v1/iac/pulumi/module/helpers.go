package awsdynamodb

import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodmpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// -----------------------------------------------------------------------------
// Pointer helpers
// -----------------------------------------------------------------------------
// These helpers keep the resource-building code readable by avoiding the verbose
// &val construction repeatedly all over the place.

// StringPtr returns a *string unless the input is the empty string, in which
// case it returns nil.  This is handy when optional Pulumi inputs should be
// omitted instead of being passed the empty value.
func StringPtr(s string) *string {
    if s == "" {
        return nil
    }
    return &s
}

// Int64Ptr returns a *int64 only when the provided value is non-zero.
func Int64Ptr(v int64) *int64 {
    if v == 0 {
        return nil
    }
    return &v
}

// BoolPtr returns a *bool unless the pointer would be nil (when the caller wants
// to explicitly omit an optional field).
func BoolPtr(b bool) *bool {
    return &b
}

// -----------------------------------------------------------------------------
// Map helpers
// -----------------------------------------------------------------------------
// ToPulumiStringMap converts a plain map[string]string (used in protobufs) to
// Pulumi's strongly-typed StringMap.  This is required because most AWS Pulumi
// resources take pulumi.StringInput values so they can reference other
// resources.
func ToPulumiStringMap(src map[string]string) pulumi.StringMap {
    if len(src) == 0 {
        return nil
    }

    out := make(pulumi.StringMap, len(src))
    for k, v := range src {
        // Pulumi helpers such as pulumi.String() already return the correct
        // input type (pulumi.StringInput) that satisfies pulumi.Input.
        out[k] = pulumi.String(v)
    }
    return out
}

// -----------------------------------------------------------------------------
// Enum translation helpers
// -----------------------------------------------------------------------------
// The protobuf representation uses Go-generated enum values whereas the Pulumi
// AWS provider relies on raw strings that match the AWS API identifiers.  The
// following helpers act as a single source of truth for those conversions and
// avoid stringly-typed code scattered throughout the implementation.

// BillingModeToAWS converts the proto BillingMode enum to the string expected by
// AWS / Pulumi ("PROVISIONED" or "PAY_PER_REQUEST").  When BillingMode is
// unspecified the empty string is returned so the caller can decide whether to
// set the attribute or leave it nil.
func BillingModeToAWS(mode awsdynamodmpb.BillingMode) string {
    switch mode {
    case awsdynamodmpb.BillingMode_PROVISIONED:
        return "PROVISIONED"
    case awsdynamodmpb.BillingMode_PAY_PER_REQUEST:
        return "PAY_PER_REQUEST"
    default:
        return ""
    }
}

// AttributeTypeToAWS converts the proto AttributeType enum to DynamoDB scalar
// type identifiers ("S", "N", "B").
func AttributeTypeToAWS(t awsdynamodmpb.AttributeType) string {
    switch t {
    case awsdynamodmpb.AttributeType_STRING:
        return "S"
    case awsdynamodmpb.AttributeType_NUMBER:
        return "N"
    case awsdynamodmpb.AttributeType_BINARY:
        return "B"
    default:
        return ""
    }
}

// KeyTypeToAWS converts the proto KeyType enum to the strings required by the
// AWS SDK / CloudFormation ("HASH" or "RANGE").
func KeyTypeToAWS(kt awsdynamodmpb.KeyType) string {
    switch kt {
    case awsdynamodmpb.KeyType_HASH:
        return "HASH"
    case awsdynamodmpb.KeyType_RANGE:
        return "RANGE"
    default:
        return ""
    }
}

// ProjectionTypeToAWS converts the proto ProjectionType enum to its AWS string
// counterpart ("ALL", "KEYS_ONLY", or "INCLUDE").
func ProjectionTypeToAWS(pt awsdynamodmpb.ProjectionType) string {
    switch pt {
    case awsdynamodmpb.ProjectionType_ALL:
        return "ALL"
    case awsdynamodmpb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY"
    case awsdynamodmpb.ProjectionType_INCLUDE:
        return "INCLUDE"
    default:
        return ""
    }
}

// StreamViewTypeToAWS converts the proto StreamViewType to the exact strings
// accepted by the AWS API ("NEW_IMAGE", "OLD_IMAGE", "NEW_AND_OLD_IMAGES",
// "KEYS_ONLY").
func StreamViewTypeToAWS(vt awsdynamodmpb.StreamViewType) string {
    switch vt {
    case awsdynamodmpb.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE"
    case awsdynamodmpb.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE"
    case awsdynamodmpb.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES"
    case awsdynamodmpb.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY"
    default:
        return ""
    }
}

// SSETypeToAWS converts the proto SSEType enum to the string expected by the
// AWS provider ("AES256" or "KMS").
func SSETypeToAWS(st awsdynamodmpb.SSEType) string {
    switch st {
    case awsdynamodmpb.SSEType_AES256:
        return "AES256"
    case awsdynamodmpb.SSEType_KMS:
        return "KMS"
    default:
        return ""
    }
}

// -----------------------------------------------------------------------------
// Optional Pulumi helpers
// -----------------------------------------------------------------------------
// Pulumi uses the *PtrInput pattern to represent optional resource arguments.
// The helpers below make it slightly nicer to build those values from the proto
// specification.

// OptionalStringInput converts a plain Go string to a pulumi.StringPtrInput.  If
// the string is empty nil is returned so that the field is cleanly omitted.
func OptionalStringInput(s string) pulumi.StringPtrInput {
    if s == "" {
        return nil
    }
    return pulumi.StringPtr(s)
}

// OptionalBoolInput converts a bool plus a flag that indicates whether the
// value was explicitly set.  This is useful when a proto field is optional but
// defaults to false, which otherwise makes it impossible to distinguish
// "unset" from "set to false" in Go.  The caller passes (value, isSet).
func OptionalBoolInput(v bool, set bool) pulumi.BoolPtrInput {
    if !set {
        return nil
    }
    return pulumi.BoolPtr(v)
}

// OptionalIntPtrInput converts an int64 to a pulumi.IntPtrInput when the value
// is positive.  Zero or negative values are treated as "unset".
func OptionalIntPtrInput(v int64) pulumi.IntPtrInput {
    if v <= 0 {
        return nil
    }
    // Safe conversion: DynamoDB capacity units are always positive and small
    // enough to fit into an int on every architecture Pulumi supports.
    return pulumi.IntPtr(int(v))
}
