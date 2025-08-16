# AWS DynamoDB Terraform Module Examples

This document provides comprehensive examples of how to use the AWS DynamoDB Terraform module for various use cases.

## Create using CLI

Create a YAML file using one of the examples shown below. After the YAML file is created, use the command below to apply the configuration:

```shell
project-planton tofu apply --manifest <yaml-path> --auto-approve
```

## Minimal manifest (YAML)

This minimal example creates a DynamoDB table with only the required fields using pay-per-request billing.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-minimal-table
spec:
  table_name: my-minimal-table
  aws_region: us-east-1
  billing_mode: BILLING_MODE_PAY_PER_REQUEST
  partition_key_name: id
  partition_key_type: ATTRIBUTE_TYPE_STRING
```

## Provisioned billing with sort key

This example demonstrates a DynamoDB table with provisioned billing mode and a composite key (partition key + sort key).

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-provisioned-table
spec:
  table_name: my-provisioned-table
  aws_region: us-west-2
  billing_mode: BILLING_MODE_PROVISIONED
  partition_key_name: user_id
  partition_key_type: ATTRIBUTE_TYPE_STRING
  sort_key_name: timestamp
  sort_key_type: ATTRIBUTE_TYPE_NUMBER
  read_capacity_units: 5
  write_capacity_units: 5
  point_in_time_recovery_enabled: true
  server_side_encryption_enabled: true
```

## High-performance table with binary keys

This example shows a DynamoDB table optimized for high-performance workloads with binary partition keys and enhanced security features.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-high-performance-table
spec:
  table_name: my-high-performance-table
  aws_region: us-east-1
  billing_mode: BILLING_MODE_PROVISIONED
  partition_key_name: session_id
  partition_key_type: ATTRIBUTE_TYPE_BINARY
  sort_key_name: event_id
  sort_key_type: ATTRIBUTE_TYPE_BINARY
  read_capacity_units: 100
  write_capacity_units: 100
  point_in_time_recovery_enabled: true
  server_side_encryption_enabled: true
```

## Development table with pay-per-request

This example creates a cost-effective development table using pay-per-request billing with minimal configuration.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-dev-table
spec:
  table_name: my-dev-table
  aws_region: us-east-1
  billing_mode: BILLING_MODE_PAY_PER_REQUEST
  partition_key_name: id
  partition_key_type: ATTRIBUTE_TYPE_STRING
  sort_key_name: created_at
  sort_key_type: ATTRIBUTE_TYPE_NUMBER
  point_in_time_recovery_enabled: false
  server_side_encryption_enabled: true
```

## Production table with numeric keys

This example demonstrates a production-ready DynamoDB table with numeric partition keys and comprehensive backup and security features.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-production-table
spec:
  table_name: my-production-table
  aws_region: us-west-2
  billing_mode: BILLING_MODE_PROVISIONED
  partition_key_name: customer_id
  partition_key_type: ATTRIBUTE_TYPE_NUMBER
  sort_key_name: order_date
  sort_key_type: ATTRIBUTE_TYPE_STRING
  read_capacity_units: 50
  write_capacity_units: 25
  point_in_time_recovery_enabled: true
  server_side_encryption_enabled: true
```

## CLI Flows

### Validate manifest

Validate your DynamoDB manifest before deployment:

```bash
project-planton validate --manifest dynamodb-table.yaml
```

### Deploy with Terraform

Deploy using Terraform as the IaC backend:

```bash
# Initialize Terraform
project-planton tofu init --manifest dynamodb-table.yaml

# Preview changes
project-planton tofu plan --manifest dynamodb-table.yaml

# Apply changes
project-planton tofu apply --manifest dynamodb-table.yaml --auto-approve

# Destroy resources
project-planton tofu destroy --manifest dynamodb-table.yaml --auto-approve
```

## Complete Manifest Example

Here's a complete manifest file that can be used with the Project Planton CLI:

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: example-dynamodb-table
  org: my-organization
  env: production
  id: example-dynamodb-table-001
spec:
  table_name: example-dynamodb-table
  aws_region: us-east-1
  billing_mode: BILLING_MODE_PAY_PER_REQUEST
  partition_key_name: id
  partition_key_type: ATTRIBUTE_TYPE_STRING
  sort_key_name: created_at
  sort_key_type: ATTRIBUTE_TYPE_NUMBER
  server_side_encryption_enabled: true
  point_in_time_recovery_enabled: true
```

---

These examples illustrate various configurations of the `AwsDynamodb` API resource, demonstrating how to define DynamoDB tables with different billing modes, key schemas, capacity settings, and security features.

**Note**: Provider credentials are supplied via stack input, not in the spec. Ensure your AWS credentials are properly configured in your Planton Cloud environment before deployment.

**Important considerations**:
- Choose the appropriate billing mode based on your usage patterns
- Design your key schema carefully for optimal performance
- Enable point-in-time recovery for production workloads
- Always enable server-side encryption for security
- Monitor your capacity usage and adjust as needed
