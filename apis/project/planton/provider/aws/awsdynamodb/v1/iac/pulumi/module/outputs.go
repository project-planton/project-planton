package module

// Keys exported by this Pulumi stack.  These constants mirror the field names
// declared in AwsDynamodbStackOutputs.  Nested fields use dot-notation to make
// it easy for callers to reference the exact attribute required.
const (
    OpTableArn                    = "table_arn"
    OpTableName                   = "table_name"
    OpTableID                     = "table_id"
    OpStreamStreamArn             = "stream.stream_arn"
    OpStreamStreamLabel           = "stream.stream_label"
    OpKmsKeyArn                   = "kms_key_arn"
    OpGlobalSecondaryIndexNames   = "global_secondary_index_names"
    OpLocalSecondaryIndexNames    = "local_secondary_index_names"
)