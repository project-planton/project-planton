package module

// This file declares the string constants that act as keys when exporting
// values from the Pulumi program. Each constant maps 1-to-1 to a field in
// the AwsDynamodbStackOutputs proto definition. For nested message fields,
// the key path is represented using dot-notation (e.g. "stream.stream_arn").

const (
    // TableArn is the fully-qualified Amazon Resource Name of the DynamoDB table.
    TableArn = "table_arn"

    // TableName is the (potentially suffixed) name of the provisioned table.
    TableName = "table_name"

    // TableID is the AWS-assigned unique identifier of the table.
    TableID = "table_id"

    // StreamArn refers to the ARN of the most recent DynamoDB Stream, when streams are enabled.
    StreamArn = "stream.stream_arn"

    // StreamLabel is the timestamp-based label that uniquely identifies the stream.
    StreamLabel = "stream.stream_label"

    // KmsKeyArn is the ARN of the customer-managed KMS key used for server-side encryption (when applicable).
    KmsKeyArn = "kms_key_arn"

    // GlobalSecondaryIndexNames lists the names of provisioned global secondary indexes (GSIs).
    GlobalSecondaryIndexNames = "global_secondary_index_names"

    // LocalSecondaryIndexNames lists the names of provisioned local secondary indexes (LSIs).
    LocalSecondaryIndexNames = "local_secondary_index_names"
)
