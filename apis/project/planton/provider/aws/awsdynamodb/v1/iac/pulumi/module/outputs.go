package module

import (
    "github.com/pulumi/pulumi-aws/sdk/v5/go/aws/dynamodb"
    "github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Export keys – keep these in sync with AwsDynamodbStackOutputs proto fields.
const (
    TableArn                  = "table_arn"
    TableName                 = "table_name"
    TableID                   = "table_id"
    StreamStreamArn           = "stream.stream_arn"   // nested: Stream.stream_arn
    StreamStreamLabel         = "stream.stream_label" // nested: Stream.stream_label
    KmsKeyArn                 = "kms_key_arn"
    GlobalSecondaryIndexNames = "global_secondary_index_names"
    LocalSecondaryIndexNames  = "local_secondary_index_names"
)

// buildOutputs converts selected properties of the created DynamoDB table into a
// map keyed by the constants above. The returned map can be iterated to export
// values via ctx.Export.
func buildOutputs(table *dynamodb.Table) map[string]pulumi.Output {
    outputs := map[string]pulumi.Output{
        TableArn:  table.Arn,
        TableName: table.Name,
        TableID:   table.ID(),
    }

    // Optional / nullable attributes ------------------------------------------------
    if table.StreamArn != nil {
        outputs[StreamStreamArn] = table.StreamArn
    }
    if table.StreamLabel != nil {
        outputs[StreamStreamLabel] = table.StreamLabel
    }
    if table.KmsKeyArn != nil {
        outputs[KmsKeyArn] = table.KmsKeyArn
    }

    // Global Secondary Index names ---------------------------------------------------
    outputs[GlobalSecondaryIndexNames] = table.GlobalSecondaryIndexes.ApplyT(
        func(gs []dynamodb.TableGlobalSecondaryIndex) []string {
            names := make([]string, len(gs))
            for i, g := range gs {
                names[i] = g.Name
            }
            return names
        },
    ).(pulumi.StringArrayOutput)

    // Local Secondary Index names ----------------------------------------------------
    outputs[LocalSecondaryIndexNames] = table.LocalSecondaryIndexes.ApplyT(
        func(ls []dynamodb.TableLocalSecondaryIndex) []string {
            names := make([]string, len(ls))
            for i, l := range ls {
                names[i] = l.Name
            }
            return names
        },
    ).(pulumi.StringArrayOutput)

    return outputs
}
