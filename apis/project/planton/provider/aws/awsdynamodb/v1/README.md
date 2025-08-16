# AwsDynamodb

AWS DynamoDB table resource for Project Planton. This module provisions a DynamoDB table with common options like billing mode, key schema, GSIs/LSIs, TTL, streams, PITR, SSE, and table class.

## Spec fields (summary)
- **billing_mode**: Billing mode for the table. BILLING_MODE_PROVISIONED or BILLING_MODE_PAY_PER_REQUEST.
- **provisioned_throughput**: RCUs and WCUs used when billing_mode is PROVISIONED.
- **attribute_definitions**: List of attributes and their types used by the table and indexes.
- **key_schema**: Primary key schema (one HASH, optional RANGE).
- **global_secondary_indexes**: GSIs with key schema, projection, and optional throughput.
- **local_secondary_indexes**: LSIs with key schema and projection.
- **ttl**: Time-to-live configuration.
- **stream_enabled/stream_view_type**: Enable and configure DynamoDB Streams.
- **point_in_time_recovery_enabled**: PITR toggle.
- **server_side_encryption**: SSE with optional KMS key.
- **table_class**: STANDARD or STANDARD_INFREQUENT_ACCESS.
- **deletion_protection_enabled**: Prevent accidental deletion.
- **contributor_insights_enabled**: Enable Contributor Insights.

## Stack outputs
- **table_name**: Table name.
- **table_arn**: Table ARN.
- **table_id**: Provider-assigned table ID.
- **stream_arn**: Stream ARN when streams are enabled.
- **stream_label**: Stream label when streams are enabled.

## How it works
This resource plugs into Project Planton’s IaC backends:
- Pulumi module (Go) and Terraform module define the provider-specific provisioning recipe.
- The CLI passes AwsDynamodbStackInput with provisioner, stack info, target resource, and AWS credentials.

## References
- DynamoDB: Table concepts — https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html
- DynamoDB: Billing — https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html
- DynamoDB: Streams — https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Streams.html
- DynamoDB: TTL — https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/TTL.html
- DynamoDB: Global secondary indexes — https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.html
- DynamoDB: Local secondary indexes — https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LSI.html


