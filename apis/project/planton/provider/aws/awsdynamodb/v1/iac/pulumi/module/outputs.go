package module

// Output keys that are exported by the stack. The constant names use the Planton
// “Op” prefix convention while their *values* match exactly the field names
// defined in the AwsDynamodbStackOutputs protobuf (dot-notation is used for
// nested messages).

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
