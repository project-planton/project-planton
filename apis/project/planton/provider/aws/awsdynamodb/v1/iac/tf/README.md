# Terraform Module to Deploy AwsDynamodb

This Terraform module deploys AWS DynamoDB tables using the Project Planton CLI (tofu) with a unified API resource model.

## CLI Usage (ProjectPlanton tofu)

The module uses the default local backend for state management. Use the following commands to manage your DynamoDB table:

```bash
# Initialize Terraform and download providers
project-planton tofu init --manifest hack/manifest.yaml

# Preview changes
project-planton tofu plan --manifest hack/manifest.yaml

# Apply changes
project-planton tofu apply --manifest hack/manifest.yaml --auto-approve

# Destroy resources
project-planton tofu destroy --manifest hack/manifest.yaml --auto-approve
```

## Manifest Configuration

The module uses the manifest file located at `hack/manifest.yaml` (created by rule 008). This manifest contains the DynamoDB table specification in the unified API format.

**Note**: Provider credentials are provided via stack input through the CLI, not in the manifest `spec`. Ensure your AWS credentials are properly configured in your environment before deployment.

## Module Features

### Billing Modes
- **PAY_PER_REQUEST**: Pay only for the requests you make (on-demand)
- **PROVISIONED**: Pay for provisioned read and write capacity units

### Key Schema Support
- **Partition Key (Hash Key)**: Required primary key for uniquely identifying items
- **Sort Key (Range Key)**: Optional secondary key for sorting items with the same partition key

### Data Types
- **STRING**: String data type for text-based keys
- **NUMBER**: Numeric data type for numeric keys  
- **BINARY**: Binary data type for binary keys

### Security & Backup Features
- **Server-Side Encryption**: Encrypt data at rest using AWS managed keys
- **Point-in-Time Recovery**: Enable backup and restore capabilities

## Outputs

The module provides the following outputs:

- `table_arn`: The ARN of the DynamoDB table
- `table_id`: The ID of the DynamoDB table
- `table_name`: The name of the DynamoDB table
- `aws_region`: The AWS region where the table is located
- `stream_arn`: The stream ARN (only exported if point-in-time recovery is enabled)

## Examples

See `examples.md` for comprehensive examples of different DynamoDB table configurations, including:
- Minimal tables with pay-per-request billing
- Provisioned tables with composite keys
- High-performance tables with binary keys
- Development and production configurations

## Best Practices

1. **Table Naming**: Use descriptive, consistent naming conventions
2. **Key Design**: Design partition keys to avoid hot partitions
3. **Capacity Planning**: Monitor usage patterns and adjust capacity accordingly
4. **Security**: Always enable server-side encryption for production tables
5. **Backup**: Enable point-in-time recovery for critical data
6. **Monitoring**: Set up CloudWatch alarms for capacity and error metrics

## State Management

This module uses the default local backend for state management. For production environments, consider using a remote backend (S3/DynamoDB) for persistent state storage and team collaboration.
