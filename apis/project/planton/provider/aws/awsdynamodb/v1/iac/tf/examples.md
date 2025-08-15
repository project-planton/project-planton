# AWS DynamoDB Examples

Below are several examples demonstrating how to define an AWS DynamoDB component in
ProjectPlanton. After creating one of these YAML manifests, apply it with Terraform using the ProjectPlanton CLI:

```shell
project-planton tofu apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic DynamoDB Table

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: basic-table
spec:
  tableName: basic-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: id
    type: S
  attributes:
    - name: id
      type: S
```

This example creates a basic DynamoDB table:
• Uses pay-per-request billing for cost optimization.
• Simple primary key with string type.
• No additional indexes or advanced features.

---

## DynamoDB with Provisioned Capacity and Autoscaling

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: autoscaling-table
spec:
  tableName: autoscaling-table
  billingMode: PROVISIONED
  hashKey:
    name: userId
    type: S
  attributes:
    - name: userId
      type: S
  autoScale:
    isEnabled: true
    readCapacity:
      minCapacity: 5
      maxCapacity: 100
      targetUtilization: 70
    writeCapacity:
      minCapacity: 5
      maxCapacity: 100
      targetUtilization: 70
```

This example includes autoscaling configuration:
• Uses provisioned billing mode with autoscaling.
• Scales read/write capacity based on utilization.
• Sets minimum and maximum capacity limits.

---

## DynamoDB with Global Secondary Index

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: gsi-table
spec:
  tableName: gsi-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: orderId
    type: S
  attributes:
    - name: orderId
      type: S
    - name: customerId
      type: S
  globalSecondaryIndexes:
    - name: CustomerIndex
      hashKey: customerId
      projectionType: ALL
```

This example demonstrates GSI usage:
• Creates a global secondary index on customerId.
• Enables efficient queries by customer.
• Uses ALL projection for complete attribute access.

---

## DynamoDB with Streams and Encryption

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: secure-table
spec:
  tableName: secure-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: sessionId
    type: S
  attributes:
    - name: sessionId
      type: S
  enableStreams: true
  streamViewType: NEW_AND_OLD_IMAGES
  serverSideEncryption:
    isEnabled: true
    kmsKeyArn: arn:aws:kms:us-west-2:123456789012:key/your-kms-key-id
```

This example includes security features:
• Enables DynamoDB Streams for change tracking.
• Uses custom KMS key for encryption.
• Captures both old and new images in streams.

---

## DynamoDB with Time to Live (TTL)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: ttl-table
spec:
  tableName: ttl-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: tokenId
    type: S
  attributes:
    - name: tokenId
      type: S
    - name: expirationTime
      type: N
  ttl:
    isEnabled: true
    attributeName: expirationTime
```

This example configures TTL:
• Automatically deletes expired items.
• Uses numeric timestamp for expiration.
• Reduces storage costs for temporary data.

---

## DynamoDB with Point-in-Time Recovery

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: pitr-table
spec:
  tableName: pitr-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: recordId
    type: S
  attributes:
    - name: recordId
      type: S
  pointInTimeRecovery:
    isEnabled: true
```

This example enables data protection:
• Allows restoration to any point in last 35 days.
• Provides continuous backup capability.
• Essential for production data protection.

---

## DynamoDB with Local Secondary Index

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: lsi-table
spec:
  tableName: lsi-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: userId
    type: S
  rangeKey:
    name: timestamp
    type: N
  attributes:
    - name: userId
      type: S
    - name: timestamp
      type: N
    - name: status
      type: S
  localSecondaryIndexes:
    - name: StatusIndex
      rangeKey: status
      projectionType: ALL
```

This example uses LSI:
• Requires composite primary key (hash + range).
• Creates alternative sort key for queries.
• Must be defined at table creation.

---

## DynamoDB Global Table

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: global-table
spec:
  tableName: global-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: itemId
    type: S
  attributes:
    - name: itemId
      type: S
  replicaRegionNames:
    - us-east-1
    - eu-west-1
    - ap-southeast-1
```

This example creates a global table:
• Replicates data across multiple regions.
• Enables low-latency global access.
• Automatically enables streams for replication.

---

## DynamoDB with Data Import

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsDynamodb
metadata:
  name: import-table
spec:
  tableName: import-table
  billingMode: PAY_PER_REQUEST
  hashKey:
    name: productId
    type: S
  attributes:
    - name: productId
      type: S
  importTable:
    inputCompressionType: GZIP
    inputFormat: DYNAMODB_JSON
    s3BucketSource:
      bucket: my-data-bucket
      keyPrefix: data/imports/
```

This example includes data import:
• Imports data from S3 during table creation.
• Supports various compression and format types.
• Useful for bulk data loading.

---

## After Deploying

Once you've applied your manifest with ProjectPlanton tofu, you can confirm that the DynamoDB table is active via the AWS console or by
using the AWS CLI:

```shell
aws dynamodb describe-table --table-name <your-table-name>
```

You should see your new DynamoDB table with its configuration details, including ARN, status, and any configured indexes or streams.
