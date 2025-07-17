package awsdynamodb

import (
    "fmt"

    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    awsdynamodbpb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// applyStreamSpec translates the user-supplied StreamSpecification (coming from the
// protobuf definition) into the corresponding Pulumi arguments that have to be
// set on aws.dynamodb.Table.
//
// The function is intentionally side-effect free with the exception of mutating
// the provided args pointer. In case the specification is nil or streams are
// disabled, the function becomes a no-op so callers do not have to perform any
// pre-checks.
func applyStreamSpec(spec *awsdynamodbpb.StreamSpecification, args *dynamodb.TableArgs) error {
    if spec == nil || !spec.GetStreamEnabled() {
        // Nothing to do – streams are either not configured or explicitly
        // turned off.
        return nil
    }

    viewType, err := convertStreamViewType(spec.GetStreamViewType())
    if err != nil {
        return err
    }

    args.StreamEnabled = pulumi.BoolPtr(true)
    args.StreamViewType = pulumi.StringPtr(viewType)

    return nil
}

// convertStreamViewType converts the protobuf enum into the exact string that
// the AWS provider for Pulumi expects. The mapping is a 1-to-1 pass-through for
// all officially supported view types.
func convertStreamViewType(t awsdynamodbpb.StreamViewType) (string, error) {
    switch t {
    case awsdynamodbpb.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE", nil
    case awsdynamodbpb.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE", nil
    case awsdynamodbpb.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES", nil
    case awsdynamodbpb.StreamViewType_STREAM_KEYS_ONLY:
        return "KEYS_ONLY", nil
    default:
        return "", fmt.Errorf("unsupported StreamViewType %q", t.String())
    }
}

// exportStreamOutputs publishes the dynamically-generated stream ARN and stream
// label as Pulumi stack outputs so that they become visible to users and can be
// referenced by other stacks/resources.
func exportStreamOutputs(ctx *pulumi.Context, table *dynamodb.Table, prefix string) {
    // In case streams are disabled the properties below will resolve to `null`
    // which is perfectly acceptable for stack outputs.
    if table == nil {
        return
    }

    ctx.Export(fmt.Sprintf("%sstreamArn", prefix), table.StreamArn)
    ctx.Export(fmt.Sprintf("%sstreamLabel", prefix), table.StreamLabel)
}
