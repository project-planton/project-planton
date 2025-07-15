package module

// Output keys exported by the stack so they can later be consumed by other
// stacks / resources. The names must exactly match the StackOutputs proto so
// glue code can rely on them.
const (
    OpTableArn                    = "table_arn"
    OpTableName                   = "table_name"
    OpTableId                     = "table_id"
    OpStreamStreamArn             = "stream.stream_arn"
    OpStreamStreamLabel           = "stream.stream_label"
    OpKmsKeyArn                   = "kms_key_arn"
    OpGlobalSecondaryIndexNames   = "global_secondary_index_names"
    OpLocalSecondaryIndexNames    = "local_secondary_index_names"
)
