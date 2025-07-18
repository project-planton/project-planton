package module

import (
    "strings"

    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Ptr returns a pointer to the supplied value. It can be used with any
// basic or user-defined type and helps when Pulumi (or AWS SDK) expects
// pointer inputs.
func Ptr[T any](v T) *T {
    return &v
}

// String returns a *string for the provided value.
func String(v string) *string { return &v }

// Bool returns a *bool for the provided value.
func Bool(v bool) *bool { return &v }

// Int returns a *int for the provided value.
func Int(v int) *int { return &v }

// Int64 returns a *int64 for the provided value.
func Int64(v int64) *int64 { return &v }

// StringSliceToPulumi converts a slice of Go strings to a Pulumi
// StringArray that can be assigned to Pulumi resource inputs.
func StringSliceToPulumi(ss []string) pulumi.StringArray {
    arr := make(pulumi.StringArray, len(ss))
    for i, s := range ss {
        arr[i] = pulumi.String(s)
    }
    return arr
}

// MergeLabels combines one or more maps into a new map. When the same key
// exists in multiple maps, the value from the *last* map wins.
func MergeLabels(base map[string]string, others ...map[string]string) map[string]string {
    merged := make(map[string]string)

    // copy base first
    for k, v := range base {
        merged[k] = v
    }

    // overlay additional maps, letting later ones override earlier ones
    for _, m := range others {
        for k, v := range m {
            merged[k] = v
        }
    }
    return merged
}

// IsBlank returns true when the string is empty or contains only
// whitespace characters.
func IsBlank(s string) bool {
    return strings.TrimSpace(s) == ""
}

// CopyStringMap makes a shallow copy of a map[string]string.
func CopyStringMap(src map[string]string) map[string]string {
    if src == nil {
        return nil
    }
    dst := make(map[string]string, len(src))
    for k, v := range src {
        dst[k] = v
    }
    return dst
}
