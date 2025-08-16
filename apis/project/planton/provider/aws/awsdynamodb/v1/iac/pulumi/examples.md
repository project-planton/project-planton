# AWS DynamoDB Pulumi Module Examples

This document provides comprehensive examples of how to use the AWS DynamoDB Pulumi module for various use cases.

## Basic DynamoDB Table with Pay-Per-Request Billing

This example creates a simple DynamoDB table with a string partition key using pay-per-request billing mode.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-simple-table
spec:
  tableName: my-simple-table
  awsRegion: us-east-1
  billingMode: PAY_PER_REQUEST
  partitionKeyName: id
  partitionKeyType: STRING
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table with Provisioned Billing

This example creates a DynamoDB table with provisioned billing mode, specifying read and write capacity units.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-provisioned-table
spec:
  tableName: my-provisioned-table
  awsRegion: us-west-2
  billingMode: PROVISIONED
  partitionKeyName: user_id
  partitionKeyType: STRING
  readCapacityUnits: 5
  writeCapacityUnits: 5
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table with Composite Key

This example creates a DynamoDB table with both a partition key and a sort key for more complex data modeling.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-composite-table
spec:
  tableName: my-composite-table
  awsRegion: eu-west-1
  billingMode: PAY_PER_REQUEST
  partitionKeyName: user_id
  partitionKeyType: STRING
  sortKeyName: timestamp
  sortKeyType: NUMBER
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table with Binary Partition Key

This example creates a DynamoDB table using a binary partition key, useful for storing binary data as keys.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: my-binary-table
spec:
  tableName: my-binary-table
  awsRegion: ap-southeast-1
  billingMode: PROVISIONED
  partitionKeyName: binary_id
  partitionKeyType: BINARY
  readCapacityUnits: 10
  writeCapacityUnits: 10
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table for User Sessions

This example creates a DynamoDB table optimized for storing user session data with TTL support.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: user-sessions-table
spec:
  tableName: user-sessions-table
  awsRegion: us-east-1
  billingMode: PAY_PER_REQUEST
  partitionKeyName: session_id
  partitionKeyType: STRING
  sortKeyName: user_id
  sortKeyType: STRING
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table for E-commerce Orders

This example creates a DynamoDB table for storing e-commerce order data with customer and order information.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: ecommerce-orders-table
spec:
  tableName: ecommerce-orders-table
  awsRegion: us-west-2
  billingMode: PROVISIONED
  partitionKeyName: customer_id
  partitionKeyType: STRING
  sortKeyName: order_id
  sortKeyType: STRING
  readCapacityUnits: 20
  writeCapacityUnits: 15
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table for IoT Device Data

This example creates a DynamoDB table for storing IoT device telemetry data with timestamp-based sorting.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: iot-telemetry-table
spec:
  tableName: iot-telemetry-table
  awsRegion: eu-central-1
  billingMode: PAY_PER_REQUEST
  partitionKeyName: device_id
  partitionKeyType: STRING
  sortKeyName: timestamp
  sortKeyType: NUMBER
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table for Gaming Leaderboards

This example creates a DynamoDB table for storing gaming leaderboard data with score-based sorting.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: gaming-leaderboard-table
spec:
  tableName: gaming-leaderboard-table
  awsRegion: ap-northeast-1
  billingMode: PROVISIONED
  partitionKeyName: game_id
  partitionKeyType: STRING
  sortKeyName: score
  sortKeyType: NUMBER
  readCapacityUnits: 25
  writeCapacityUnits: 20
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table for Content Management

This example creates a DynamoDB table for storing content with category and creation date sorting.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: content-management-table
spec:
  tableName: content-management-table
  awsRegion: sa-east-1
  billingMode: PAY_PER_REQUEST
  partitionKeyName: content_id
  partitionKeyType: STRING
  sortKeyName: category
  sortKeyType: STRING
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## DynamoDB Table for Financial Transactions

This example creates a DynamoDB table for storing financial transaction data with account and transaction date.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: financial-transactions-table
spec:
  tableName: financial-transactions-table
  awsRegion: us-east-1
  billingMode: PROVISIONED
  partitionKeyName: account_id
  partitionKeyType: STRING
  sortKeyName: transaction_date
  sortKeyType: NUMBER
  readCapacityUnits: 30
  writeCapacityUnits: 25
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## Manifest File Example

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
  tableName: example-dynamodb-table
  awsRegion: us-east-1
  billingMode: PAY_PER_REQUEST
  partitionKeyName: id
  partitionKeyType: STRING
  sortKeyName: created_at
  sortKeyType: NUMBER
  serverSideEncryptionEnabled: true
  pointInTimeRecoveryEnabled: true
```

## CLI Commands

### Preview Changes
```bash
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/my-project/dynamodb-stack \
  --module-dir .
```

### Apply Changes
```bash
project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack organization/my-project/dynamodb-stack \
  --module-dir . \
  --yes
```

### Refresh State
```bash
project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack organization/my-project/dynamodb-stack \
  --module-dir .
```

### Destroy Resources
```bash
project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack organization/my-project/dynamodb-stack \
  --module-dir .
```

## Usage Notes

1. **Table Names**: Must be unique within the AWS account and region, between 3-255 characters, and contain only alphanumeric characters, hyphens, and underscores.

2. **Billing Mode Selection**:
   - Use `PAY_PER_REQUEST` for unpredictable workloads or development environments
   - Use `PROVISIONED` for predictable workloads with known capacity requirements

3. **Key Design**:
   - Choose partition keys that distribute data evenly across partitions
   - Use sort keys to enable efficient range queries
   - Consider access patterns when designing your key schema

4. **Security**:
   - Always enable server-side encryption for production tables
   - Enable point-in-time recovery for critical data backup requirements

5. **Capacity Planning**:
   - Monitor CloudWatch metrics to optimize provisioned capacity
   - Use auto-scaling for provisioned tables with variable workloads
