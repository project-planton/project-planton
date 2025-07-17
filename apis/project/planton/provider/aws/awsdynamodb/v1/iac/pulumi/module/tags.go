package module

import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ToPulumiStringMap converts a plain Go string map into a Pulumi StringMap. It
// returns nil if the provided map is nil to make it convenient to propagate
// nils directly into resource options.
func ToPulumiStringMap(tags map[string]string) pulumi.StringMap {
    if tags == nil {
        return nil
    }

    m := make(pulumi.StringMap, len(tags))
    for k, v := range tags {
        m[k] = pulumi.String(v)
    }
    return m
}

// MergeTags combines two maps containing AWS resource tags. Values from
// overrides take precedence over defaultTags. The returned value is already
// converted into a pulumi.StringMap so it can be assigned directly to the
// Tags argument of any AWS resource.
//
// Passing nil maps is supported and results in the other map being returned as
// is (after conversion). If both are nil, the function returns nil.
func MergeTags(defaultTags, overrides map[string]string) pulumi.StringMap {
    // Fast-path: both are nil.
    if defaultTags == nil && overrides == nil {
        return nil
    }

    merged := make(map[string]string, len(defaultTags)+len(overrides))

    // Copy defaults first so they can be overridden.
    for k, v := range defaultTags {
        merged[k] = v
    }
    // Apply overrides (if any).
    for k, v := range overrides {
        merged[k] = v
    }

    return ToPulumiStringMap(merged)
}
