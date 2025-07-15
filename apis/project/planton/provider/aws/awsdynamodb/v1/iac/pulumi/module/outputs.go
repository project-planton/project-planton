package module

// This file declares string constants that match the field names defined in the
// AwsDynamodbStackOutputs proto.  Using constants avoids typos when exporting
// values from the Pulumi program.

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
