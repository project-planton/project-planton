# AwsDynamodb

The AWS DynamoDB API resource provides a consistent and streamlined interface for creating and managing Amazon DynamoDB tables within our cloud infrastructure. By abstracting the complexities of DynamoDB table configurations, this resource allows you to define your NoSQL database requirements effortlessly while ensuring consistency and compliance across different environments.

## Why We Created This API Resource

Managing DynamoDB tables directly can be complex due to the various configuration options, key schema design, billing modes, and best practices that need to be considered. To simplify this process and promote a standardized approach, we developed this API resource. It enables you to:

- **Simplify Table Management**: Easily create and configure DynamoDB tables without dealing with low-level AWS configurations.
- **Ensure Consistency**: Maintain uniform DynamoDB table configurations across different environments and projects.
- **Enhance Security**: Control encryption settings and access patterns to ensure data security.
- **Optimize Costs**: Choose appropriate billing modes (provisioned vs pay-per-request) based on your usage patterns.
- **Improve Productivity**: Reduce the time and effort required to manage database resources, allowing you to focus on application development.

## Spec Fields

### Required Fields

- **`table_name`** (string, 3-255 chars): The name of the DynamoDB table. Must be unique within the AWS account and region.
- **`aws_region`** (string): The AWS region where the DynamoDB table will be created.
- **`billing_mode`** (enum): The billing mode for the table:
  - `BILLING_MODE_PROVISIONED`: Pay for provisioned read and write capacity units
  - `BILLING_MODE_PAY_PER_REQUEST`: Pay only for the requests you make
- **`partition_key_name`** (string): The name of the partition key (hash key) for the table.
- **`partition_key_type`** (enum): The data type of the partition key:
  - `ATTRIBUTE_TYPE_STRING`: String data type
  - `ATTRIBUTE_TYPE_NUMBER`: Numeric data type
  - `ATTRIBUTE_TYPE_BINARY`: Binary data type

### Optional Fields

- **`sort_key_name`** (string): The name of the sort key (range key) for the table, if applicable.
- **`sort_key_type`** (enum): The data type of the sort key (same options as partition key type).
- **`read_capacity_units`** (int32, ≥0): The number of read capacity units for provisioned billing mode.
- **`write_capacity_units`** (int32, ≥0): The number of write capacity units for provisioned billing mode.
- **`point_in_time_recovery_enabled`** (bool, default: false): Flag to enable point-in-time recovery for the table.
- **`server_side_encryption_enabled`** (bool, default: true): Flag to enable server-side encryption for the table.

### Validation Rules

- When `billing_mode` is `PROVISIONED`, both `read_capacity_units` and `write_capacity_units` must be greater than 0.
- When `billing_mode` is `PAY_PER_REQUEST`, both capacity units must be 0.
- If `sort_key_name` is provided, `sort_key_type` must also be specified.
- If `sort_key_name` is empty, `sort_key_type` must be `UNSPECIFIED`.

## Stack Outputs

After successful deployment, the following outputs are available:

- **`table_arn`** (string): The ARN of the DynamoDB table (e.g., `arn:aws:dynamodb:us-east-1:123456789012:table/my-table`)
- **`table_id`** (string): The ID of the DynamoDB table (e.g., `12345678-1234-1234-1234-123456789012`)
- **`table_name`** (string): The name of the DynamoDB table
- **`stream_arn`** (string): The stream ARN if point-in-time recovery is enabled
- **`aws_region`** (string): The AWS region where the table is located

## How It Works

This resource supports both Pulumi and Terraform as Infrastructure as Code (IaC) backends:

- **Pulumi**: Uses the AWS provider to create DynamoDB tables with TypeScript/Go/Python
- **Terraform**: Uses the AWS provider to create DynamoDB tables with HCL

The resource automatically handles:
- Table creation with specified key schema
- Billing mode configuration
- Capacity provisioning (for provisioned mode)
- Encryption settings
- Point-in-time recovery configuration
- Output collection for downstream consumption

## References

- [AWS DynamoDB Developer Guide](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/)
- [DynamoDB Table Creation](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.Basics.html)
- [DynamoDB Billing Modes](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html)
- [DynamoDB Encryption](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/EncryptionAtRest.html)
- [DynamoDB Point-in-Time Recovery](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/PointInTimeRecovery.html)

## Benefits

- **Simplified Deployment**: Abstracts the complexities of AWS DynamoDB table configurations into an easy-to-use API.
- **Consistency**: Ensures all DynamoDB tables adhere to organizational standards for security and performance.
- **Scalability**: Allows for efficient management of database resources as your application and data storage needs grow.
- **Security**: Provides control over encryption and access patterns, reducing the risk of unauthorized data access.
- **Cost Optimization**: Enables appropriate billing mode selection based on usage patterns.
- **Flexibility**: Customize table settings to meet specific application requirements without compromising best practices.
