package module

import (
    "github.com/pkg/errors"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// ToPulumiStringMap converts a regular Go map[string]string into a Pulumi
// StringMap that can be supplied to any `Tags`/`labels` style argument.
func ToPulumiStringMap(input map[string]string) pulumi.StringMap {
    if len(input) == 0 {
        return nil
    }
    out := make(pulumi.StringMap, len(input))
    for k, v := range input {
        // Use pulumi.String() to obtain a pulumi.StringInput literal.
        out[k] = pulumi.String(v)
    }
    return out
}

// ToPulumiStringArray converts a []string into a Pulumi StringArray.
func ToPulumiStringArray(input []string) pulumi.StringArray {
    if len(input) == 0 {
        return nil
    }
    out := make(pulumi.StringArray, len(input))
    for i, v := range input {
        out[i] = pulumi.String(v)
    }
    return out
}

// ---------------------------------------------------------------------------
// Enum → provider-string helpers
// ---------------------------------------------------------------------------

// AttributeTypeToString converts proto AttributeType to the one-letter AWS
// representation expected by the Terraform AWS provider (and therefore Pulumi).
func AttributeTypeToString(t awsdynamodbpb.AttributeType) (string, error) {
    switch t {
    case awsdynamodbpb.AttributeType_STRING:
        return "S", nil
    case awsdynamodbpb.AttributeType_NUMBER:
        return "N", nil
    case awsdynamodbpb.AttributeType_BINARY:
        return "B", nil
    default:
        return "", errors.Errorf("unsupported AttributeType %q", t)
    }
}

// KeyTypeToString converts proto KeyType to the upper-case string required by
// the provider ("HASH" or "RANGE").
func KeyTypeToString(t awsdynamodbpb.KeyType) (string, error) {
    switch t {
    case awsdynamodbpb.KeyType_HASH:
        return "HASH", nil
    case awsdynamodbpb.KeyType_RANGE:
        return "RANGE", nil
    default:
        return "", errors.Errorf("unsupported KeyType %q", t)
    }
}

// BillingModeToString converts proto BillingMode to the provider value.
func BillingModeToString(m awsdynamodbpb.BillingMode) (string, error) {
    switch m {
    case awsdynamodbpb.BillingMode_PROVISIONED:
        return "PROVISIONED", nil
    case awsdynamodbpb.BillingMode_PAY_PER_REQUEST:
        return "PAY_PER_REQUEST", nil
    default:
        return "", errors.Errorf("unsupported BillingMode %q", m)
    }
}

// ProjectionTypeToString converts proto ProjectionType to the provider value.
func ProjectionTypeToString(p awsdynamodbpb.ProjectionType) (string, error) {
    switch p {
    case awsdynamodbpb.ProjectionType_ALL:
        return "ALL", nil
    case awsdynamodbpb.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY", nil
    case awsdynamodbpb.ProjectionType_INCLUDE:
        return "INCLUDE", nil
    default:
        return "", errors.Errorf("unsupported ProjectionType %q", p)
    }
}

// StreamViewTypeToString converts proto StreamViewType to the provider value.
func StreamViewTypeToString(s awsdynamodbpb.StreamViewType) (string, error) {
    switch s {
    case awsdynamodbpb.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE", nil
    case awsdynamodbpb.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE", nil
    case awsdynamodbpb.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES", nil
    case awsdynamodbpb.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY", nil
    default:
        return "", errors.Errorf("unsupported StreamViewType %q", s)
    }
}

// SSETypeToString converts proto SSEType to the provider value. Note that the
// AWS provider only needs this when server-side encryption is enabled.
func SSETypeToString(s awsdynamodbpb.SSEType) (string, error) {
    switch s {
    case awsdynamodbpb.SSEType_AES256:
        return "AES256", nil
    case awsdynamodbpb.SSEType_KMS:
        return "KMS", nil
    default:
        return "", errors.Errorf("unsupported SSEType %q", s)
    }
}
