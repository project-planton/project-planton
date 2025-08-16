# Pulumi Module to Deploy AwsDynamodb

## CLI usage (ProjectPlanton pulumi)

```bash
# Preview
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .

# Update (apply)
project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes

# Refresh
project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .

# Destroy
project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```

## Debugging

This module includes a `debug.sh` helper. To enable debugging, edit `Pulumi.yaml` and uncomment the `runtime.options.binary` line so Pulumi runs the program via the script:

```yaml
name: aws-dynamodb-module
runtime:
  name: go
  options:
    binary: ./debug.sh
```

Then make the script executable and run your command (e.g., `preview` or `update`). See `docs/pages/docs/guide/debug-pulumi-modules.mdx` for full instructions.

```bash
chmod +x debug.sh
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```

# AWS DynamoDB Pulumi Module

## Introduction

The AWS DynamoDB Pulumi Module provides a standardized and efficient way to define and deploy AWS DynamoDB tables using a Kubernetes-like API resource model. By leveraging our unified APIs, developers can specify their DynamoDB table configurations in simple YAML files, which the module then uses to create and manage AWS DynamoDB resources through Pulumi. This approach abstracts the complexity of AWS interactions and streamlines the deployment process, enabling consistent infrastructure management across multi-cloud environments.

## Key Features

- **Kubernetes-Like API Resource Model**: Utilizes a familiar structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status`, making it intuitive for developers accustomed to Kubernetes to define AWS DynamoDB resources.
- **Comprehensive DynamoDB Configuration**: Supports all major DynamoDB features including billing modes (provisioned and pay-per-request), key schemas (partition and sort keys), point-in-time recovery, and server-side encryption.
- **Flexible Key Schema**: Supports both simple primary keys (partition key only) and composite keys (partition key + sort key) with configurable data types (String, Number, Binary).
- **Security Features**: Built-in support for server-side encryption and point-in-time recovery for enhanced data protection and backup capabilities.
- **Cost Optimization**: Support for both provisioned and on-demand billing modes to optimize costs based on usage patterns.
- **Multi-Region Support**: Configurable AWS region deployment with proper credential management.

## Architecture

The module follows the standard Project Planton architecture:

1. **API Definition**: DynamoDB resources are defined using protobuf-based APIs with validation
2. **Stack Input**: The module accepts stack input containing the target resource and AWS credentials
3. **Resource Creation**: Pulumi creates the DynamoDB table with the specified configuration
4. **Output Export**: The module exports table ARN, ID, name, region, and stream ARN (if applicable)

## Supported Features

### Billing Modes
- **PROVISIONED**: Pay for provisioned read and write capacity units
- **PAY_PER_REQUEST**: Pay only for the requests you make

### Key Schema
- **Partition Key (Hash Key)**: Required primary key for uniquely identifying items
- **Sort Key (Range Key)**: Optional secondary key for sorting items with the same partition key

### Data Types
- **STRING**: String data type for text-based keys
- **NUMBER**: Numeric data type for numeric keys
- **BINARY**: Binary data type for binary keys

### Security & Backup
- **Server-Side Encryption**: Encrypt data at rest using AWS managed keys
- **Point-in-Time Recovery**: Enable backup and restore capabilities

## Outputs

The module exports the following outputs:

- `table_arn`: The ARN of the DynamoDB table
- `table_id`: The ID of the DynamoDB table
- `table_name`: The name of the DynamoDB table
- `aws_region`: The AWS region where the table is located
- `stream_arn`: The stream ARN (only exported if point-in-time recovery is enabled)

## Error Handling

The module includes comprehensive error handling for:
- Invalid AWS credentials
- DynamoDB table creation failures
- Configuration validation errors
- Resource timeout scenarios

## Performance Considerations

- **Provisioned Capacity**: Set appropriate read and write capacity units for predictable workloads
- **On-Demand Billing**: Use for unpredictable workloads to avoid over-provisioning
- **Key Design**: Design partition keys to distribute data evenly across partitions
- **Indexes**: Consider adding Global Secondary Indexes (GSIs) for complex query patterns

## Best Practices

1. **Table Naming**: Use descriptive, consistent naming conventions
2. **Key Design**: Design partition keys to avoid hot partitions
3. **Capacity Planning**: Monitor usage patterns and adjust capacity accordingly
4. **Security**: Always enable server-side encryption for production tables
5. **Backup**: Enable point-in-time recovery for critical data
6. **Monitoring**: Set up CloudWatch alarms for capacity and error metrics
