package module

// List of output keys that will be exported at the end of the stack execution.
// These identifiers are kept in a single place so that other files (unit tests
// for example) can reuse them without hard-coding string literals.
const (
    OpTableArn                   = "table_arn"
    OpTableName                  = "table_name"
    OpTableId                    = "table_id"
    OpStreamStreamArn            = "stream.stream_arn"
    OpStreamStreamLabel          = "stream.stream_label"
    OpKmsKeyArn                  = "kms_key_arn"
    OpGlobalSecondaryIndexNames  = "global_secondary_index_names"
    OpLocalSecondaryIndexNames   = "local_secondary_index_names"
)
