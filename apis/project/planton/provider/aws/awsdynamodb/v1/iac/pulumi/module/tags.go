package awsdynamodb

import (
    "sort"

    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// MergeTags returns a new map that contains the union of defaultTags and
// userTags, with user-supplied tags overriding default values when keys
// collide.  Tags whose key or value is an empty string are silently ignored
// because AWS rejects such entries.
func MergeTags(userTags, defaultTags map[string]string) map[string]string {
    merged := make(map[string]string)

    // Helper that inserts a set of tags, skipping empty keys/values.
    add := func(tags map[string]string) {
        for k, v := range tags {
            if k == "" || v == "" {
                continue
            }
            merged[k] = v
        }
    }

    // Defaults first so that user tags win on conflicts.
    add(defaultTags)
    add(userTags)

    return merged
}

// GetDefaultTags produces a minimal set of recommended default tags that
// identify the Pulumi project and stack.  Callers can pass the result to
// MergeTags to combine them with user-specified tags.
func GetDefaultTags(ctx *pulumi.Context) map[string]string {
    if ctx == nil {
        return map[string]string{}
    }
    return map[string]string{
        "Pulumi:Project": ctx.Project(),
        "Pulumi:Stack":   ctx.Stack(),
    }
}

// BuildAwsTags converts an ordinary Go map into a pulumi.StringMap suitable
// for use with the AWS provider.  The function sorts the keys to ensure a
// deterministic result, which makes unit tests and previews more stable even
// though Pulumi itself does not require a particular order.
func BuildAwsTags(tags map[string]string) pulumi.StringMap {
    if len(tags) == 0 {
        return pulumi.StringMap{}
    }

    keys := make([]string, 0, len(tags))
    for k := range tags {
        keys = append(keys, k)
    }
    sort.Strings(keys)

    result := make(pulumi.StringMap, len(tags))
    for _, k := range keys {
        result[k] = pulumi.String(tags[k])
    }
    return result
}
