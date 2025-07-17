package module

import (
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// -----------------------------------------------------------------------------
// Output keys (must be used when calling ctx.Export)
// -----------------------------------------------------------------------------
const (
    // Top-level fields
    TableArnKey                  = "TableArn"
    TableNameKey                 = "TableName"
    TableIDKey                   = "TableId"
    KmsKeyArnKey                 = "KmsKeyArn"
    GlobalSecondaryIndexNamesKey = "GlobalSecondaryIndexNames"
    LocalSecondaryIndexNamesKey  = "LocalSecondaryIndexNames"

    // Nested "stream" message fields – use dot notation
    StreamArnKey   = "Stream.StreamArn"
    StreamLabelKey = "Stream.StreamLabel"
)

// BuildOutputsArgs bundles all values that should be exported from the stack.
// The individual fields mirror AwsDynamodbStackOutputs. Optional values use
// the Ptr-variant output types so callers can omit them by passing nil.
//
// IMPORTANT: Always provide non-nil outputs for required fields (TableArn,
// TableName, TableID). Optional fields may be nil if they are not applicable
// for the current table configuration.
//
//   - All string-scalars       -> pulumi.StringOutput / pulumi.StringPtrOutput
//   - Repeated string fields   -> pulumi.StringArrayOutput
//
// This indirection keeps main.go focused on resource creation while the
// mapping logic lives in a single location.
type BuildOutputsArgs struct {
    TableArn  pulumi.StringOutput
    TableName pulumi.StringOutput
    TableID   pulumi.StringOutput

    // Optional nested Stream outputs
    StreamArn   pulumi.StringPtrOutput
    StreamLabel pulumi.StringPtrOutput

    // Optional KMS CMK ARN (present when SSE uses a CMK)
    KmsKeyArn pulumi.StringPtrOutput

    // Index names
    GlobalSecondaryIndexNames pulumi.StringArrayOutput
    LocalSecondaryIndexNames  pulumi.StringArrayOutput
}

// BuildOutputs exports the stack outputs defined by AwsDynamodbStackOutputs
// using the canonical keys declared in this file. It gracefully skips optional
// values that are not provided (i.e. the corresponding output is nil).
func BuildOutputs(ctx *pulumi.Context, args *BuildOutputsArgs) error {
    if ctx == nil || args == nil {
        return nil
    }

    // Required fields — always exported
    ctx.Export(TableArnKey, args.TableArn)
    ctx.Export(TableNameKey, args.TableName)
    ctx.Export(TableIDKey, args.TableID)

    // Optional nested Stream outputs
    if args.StreamArn != nil {
        ctx.Export(StreamArnKey, args.StreamArn)
    }
    if args.StreamLabel != nil {
        ctx.Export(StreamLabelKey, args.StreamLabel)
    }

    // Optional KMS CMK output
    if args.KmsKeyArn != nil {
        ctx.Export(KmsKeyArnKey, args.KmsKeyArn)
    }

    // Index names (can be nil when no GSIs/LSIs are defined)
    if args.GlobalSecondaryIndexNames != nil {
        ctx.Export(GlobalSecondaryIndexNamesKey, args.GlobalSecondaryIndexNames)
    }
    if args.LocalSecondaryIndexNames != nil {
        ctx.Export(LocalSecondaryIndexNamesKey, args.LocalSecondaryIndexNames)
    }

    return nil
}
