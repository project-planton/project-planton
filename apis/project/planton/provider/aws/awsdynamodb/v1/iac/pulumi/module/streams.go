package awsdynamodb

import (
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

    pb "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

// buildStreamSpecification translates the user-supplied protobuf definition into
// a Pulumi dynamodb.TableStreamSpecificationPtrInput. When streams aren’t
// enabled, the function returns nil so the resulting Table resource will be
// created without a StreamSpecification block.
func buildStreamSpecification(spec *pb.AwsDynamodbSpec) dynamodb.TableStreamSpecificationPtrInput {
    if spec == nil {
        return nil
    }

    ss := spec.GetStreamSpecification()
    if ss == nil || !ss.GetStreamEnabled() {
        // Streams are not requested – don’t include the block.
        return nil
    }

    viewType, ok := mapStreamViewType(ss.GetStreamViewType())
    if !ok {
        // Unsupported/unknown view type – let the provider surface an error by
        // returning nil; this keeps the helper side-effect-free.
        return nil
    }

    return &dynamodb.TableStreamSpecificationArgs{
        StreamEnabled:  pulumi.BoolPtr(true),
        StreamViewType: pulumi.StringPtr(viewType),
    }
}

// mapStreamViewType converts the protobuf enum to the string expected by the
// AWS provider. The boolean indicates whether a mapping was found.
func mapStreamViewType(t pb.StreamViewType) (string, bool) {
    switch t {
    case pb.StreamViewType_NEW_IMAGE:
        return "NEW_IMAGE", true
    case pb.StreamViewType_OLD_IMAGE:
        return "OLD_IMAGE", true
    case pb.StreamViewType_NEW_AND_OLD_IMAGES:
        return "NEW_AND_OLD_IMAGES", true
    case pb.StreamViewType_STREAM_KEYS_ONLY:
        // The Terraform/AWS name is "KEYS_ONLY".
        return "KEYS_ONLY", true
    default:
        return "", false
    }
}

// exportStreamOutputs publishes the Stream identifiers (ARN & label) as a
// nested object called "stream" on the stack. The shape mirrors the
// AwsDynamodbStackOutputs.Stream message so that callers can unmarshal it
// directly.
func exportStreamOutputs(ctx *pulumi.Context, table *dynamodb.Table) {
    if ctx == nil || table == nil {
        return
    }

    stream := pulumi.All(table.StreamArn, table.StreamLabel).ApplyT(func(vs []interface{}) interface{} {
        arn, label := "", ""
        if vs[0] != nil {
            arn = vs[0].(string)
        }
        if vs[1] != nil {
            label = vs[1].(string)
        }

        // No ARN => streams disabled; omit from outputs.
        if arn == "" {
            return nil
        }

        return map[string]interface{}{
            "stream_arn":   arn,
            "stream_label": label,
        }
    })

    ctx.Export("stream", stream)
}
