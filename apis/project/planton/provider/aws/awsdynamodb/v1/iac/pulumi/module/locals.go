// Code generated for the aws_dynamodb Pulumi component.
// DO NOT EDIT MANUALLY.
//
// This file contains small generic helpers that are useful across the
// component implementation.  Nothing in here depends on the shape of the
// resource inputs or outputs, therefore importing it does **not** create any
// resource in the target cloud provider.
//
// The helpers fall in three categories:
//   1. A Locals struct that can be embedded by the main component to expose
//      computed values that are not strictly part of the public surface area.
//   2. Naming helpers that make it trivial to comply with DynamoDB naming
//      constraints (3-255 chars, letters, numbers, "_", "-", ".") while still
//      guaranteeing uniqueness per-stack.
//   3. Small utility functions (generic pointer helper, map converters, …)
//      that keep the rest of the code tidy.
package awsdynamodb

import (
    "crypto/sha1"
    "encoding/hex"
    "fmt"
    "regexp"
    "strings"

    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// -----------------------------------------------------------------------------
// Constants & regular expressions
// -----------------------------------------------------------------------------

const (
    // DynamoDB table name limits according to AWS documentation.
    minTableNameLength = 3
    maxTableNameLength = 255

    // Pulumi logical name used for the DynamoDB table resource (kept in a
    // constant so every file uses the exact same string and accidental typos
    // do not create duplicated resources).
    pulumiTableLogicalName = "dynamodbTable"
)

var (
    // Allowed characters for DynamoDB table names as per official docs:
    // https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html
    tableNameAllowedChars = regexp.MustCompile(`[A-Za-z0-9_.-]+`)
)

// -----------------------------------------------------------------------------
// Locals struct
// -----------------------------------------------------------------------------

// Locals bundles together handy values derived from the user-supplied inputs or
// from the Pulumi execution context.  It is **not** part of the public API – it
// simply makes the implementation code easier to read by avoiding plumbing of
// the same handful of values through many functions.
//
// Typical usage inside the component constructor:
//
//   l := NewLocals(ctx, inputs.TableName, inputs.Tags)
//   table, err := dynamodb.NewTable(l.ctx, pulumiTableLogicalName, …)
//
// All fields are intentionally exported so that other files in the same
// package can access them without resorting to getter methods.
// -----------------------------------------------------------------------------
type Locals struct {
    // Pulumi context (kept here so helper functions do not need to receive it
    // as an extra parameter).
    ctx *pulumi.Context

    // BaseName is the raw table name requested by the user (prior to any
    // sanitisation or automatic suffix addition).
    BaseName string

    // TableName is the final DynamoDB-compliant name that will be sent to the
    // AWS API.  It is returned as a Pulumi output so other components can rely
    // on it even when the name contains stack-dependent parts.
    TableName pulumi.StringOutput

    // Tags as a Pulumi-native StringMap so they can be spread on any AWS
    // resource without conversion boilerplate.
    Tags pulumi.StringMap
}

// NewLocals instantiates the helper structure and pre-computes all its
// exported fields.
func NewLocals(ctx *pulumi.Context, baseName string, rawTags map[string]string) *Locals {
    l := &Locals{
        ctx:      ctx,
        BaseName: baseName,
        Tags:     toPulumiStringMap(rawTags),
    }

    l.TableName = makeTableName(ctx, baseName)
    return l
}

// -----------------------------------------------------------------------------
// Naming helpers
// -----------------------------------------------------------------------------

// makeTableName returns a DynamoDB-compatible table name and guarantees it is
// unique **per Pulumi stack** by appending ctx.Stack() when the user-supplied
// string does not already end with it.
//
// The function also makes sure the final name never exceeds the 255 chars hard
// limit – if it does, the beginning of the name is truncated and replaced with
// an 8-char SHA-1 checksum to preserve uniqueness while respecting AWS limits.
func makeTableName(ctx *pulumi.Context, requested string) pulumi.StringOutput {
    sanitized := sanitizeTableName(requested)

    // If the user manually added the stack suffix, do not append it twice.
    stack := ctx.Stack()
    if !strings.HasSuffix(sanitized, "-"+stack) {
        sanitized = fmt.Sprintf("%s-%s", sanitized, stack)
    }

    // Truncate + hash when the name exceeds the max length allowed by AWS.
    if len(sanitized) > maxTableNameLength {
        hash := sha1.Sum([]byte(sanitized))
        digest := hex.EncodeToString(hash[:])[:8]

        // We need room for the hyphen that separates the hash from the tail.
        keep := maxTableNameLength - len(digest) - 1
        if keep < 0 {
            // Very unlikely, but guard against negative slice panics.
            keep = 0
        }
        sanitized = fmt.Sprintf("%s-%s", sanitized[len(sanitized)-keep:], digest)
    }

    // Pulumi expects outputs when the value can depend on runtime information –
    // here the only runtime bit is the stack name, accessible synchronously, so
    // we can safely create a constant StringOutput.
    return pulumi.String(sanitized).ToStringOutput()
}

// sanitizeTableName enforces DynamoDB table name rules by:
//   1. Replacing disallowed characters with dashes.
//   2. Trimming consecutive dashes or dots that could appear after replacements.
//   3. Ensuring the name is at least 3 characters long (by padding with "tbl").
func sanitizeTableName(name string) string {
    if name == "" {
        name = "tbl"
    }

    // Replace every disallowed char by "-".
    builder := strings.Builder{}
    for _, r := range name {
        if tableNameAllowedChars.MatchString(string(r)) {
            builder.WriteRune(r)
        } else {
            builder.WriteRune('-')
        }
    }
    sanitized := builder.String()

    // Collapse repeated dashes and dots that could happen after replacements.
    sanitized = strings.ReplaceAll(sanitized, "..", ".")
    sanitized = strings.ReplaceAll(sanitized, "--", "-")

    // DynamoDB minimum length is 3 – pad if necessary.
    for len(sanitized) < minTableNameLength {
        sanitized += "x"
    }
    return sanitized
}

// -----------------------------------------------------------------------------
// Misc. utilities
// -----------------------------------------------------------------------------

// Ptr is a tiny generic helper to easily take the address of composite literals
// without creating an intermediate variable:
//
//   table, _ := dynamodb.NewTable(ctx, name, &dynamodb.TableArgs{
//       BillingMode: pulumi.StringPtr("PAY_PER_REQUEST"),
//       Tags:        locals.Tags,
//   })
func Ptr[T any](v T) *T { return &v }

// toPulumiStringMap converts a plain Go map to the pulumi.StringMap expected by
// every aws.* constructor.
func toPulumiStringMap(m map[string]string) pulumi.StringMap {
    sm := pulumi.StringMap{}
    for k, v := range m {
        sm[k] = pulumi.String(v)
    }
    return sm
}
