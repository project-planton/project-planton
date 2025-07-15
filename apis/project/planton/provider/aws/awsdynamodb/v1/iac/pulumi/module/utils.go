package module

import (
    "strings"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// The following helpers convert proto enum values into the exact string literals
// expected by the AWS provider.

func attributeTypeToString(t awsdynamodbv1.AttributeType) string {
    switch t {
    case awsdynamodbv1.AttributeType_STRING:
        return "S"
    case awsdynamodbv1.AttributeType_NUMBER:
        return "N"
    case awsdynamodbv1.AttributeType_BINARY:
        return "B"
    default:
        return "" // Provider will complain which is fine – validation should
                   // have caught this before.
    }
}

func keyTypeToString(t awsdynamodbv1.KeyType) string {
    switch t {
    case awsdynamodbv1.KeyType_HASH:
        return "HASH"
    case awsdynamodbv1.KeyType_RANGE:
        return "RANGE"
    default:
        return ""
    }
}

func projectionTypeToString(t awsdynamodbv1.ProjectionType) string {
    switch t {
    case awsdynamodbv1.ProjectionType_ALL:
        return "ALL"
    case awsdynamodbv1.ProjectionType_KEYS_ONLY:
        return "KEYS_ONLY"
    case awsdynamodbv1.ProjectionType_INCLUDE:
        return "INCLUDE"
    default:
        return ""
    }
}

func streamViewTypeToString(t awsdynamodbv1.StreamViewType) string {
    // Classic provider expects the *constant name* (NEW_IMAGE, OLD_IMAGE …)
    return strings.ToUpper(t.String())
}
