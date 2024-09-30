# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates a DynamoDB table with a simple primary key.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-simple-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-simple-table
    billingMode: PAY_PER_REQUEST
    hashKey:
      name: id
      type: S
    attributes:
      - name: id
        type: S
```

# Example with Provisioned Capacity and Autoscaling

This example creates a DynamoDB table with provisioned read and write capacity, along with autoscaling configurations.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-autoscaling-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-autoscaling-table
    billingMode: PROVISIONED
    readCapacity: 5
    writeCapacity: 5
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

# Example with Global Secondary Index

This example creates a DynamoDB table with a global secondary index (GSI) to enable queries on a non-primary key attribute.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-gsi-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-gsi-table
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
        hashKey:
          name: customerId
          type: S
        projectionType: ALL
```

# Example with Stream Enabled and Server-Side Encryption

This example creates a DynamoDB table with stream enabled and server-side encryption using a specified AWS KMS key.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-secure-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-secure-table
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

# Example with Time to Live (TTL) Configuration

This example creates a DynamoDB table with TTL enabled on a specified attribute.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-ttl-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-ttl-table
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

# Example with Point-in-Time Recovery

This example creates a DynamoDB table with point-in-time recovery enabled for data protection.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-pitr-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-pitr-table
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

# Example with Local Secondary Index

This example creates a DynamoDB table with a local secondary index (LSI).

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-lsi-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-lsi-table
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
        rangeKey:
          name: status
          type: S
        projectionType: ALL
```

# Example with Data Import from S3

This example creates a DynamoDB table and imports data from an S3 bucket.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-import-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-import-table
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

# Example with Multiple Replica Regions (Global Table)

This example creates a DynamoDB Global Table with replicas in multiple regions.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsDynamodb
metadata:
  name: my-global-table
spec:
  awsCredentialId: my-aws-credential-id
  table:
    tableName: my-global-table
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

---

These examples illustrate various configurations of the `AwsDynamodb` API resource, demonstrating how to define DynamoDB tables with different features such as autoscaling, indexes, encryption, TTL, point-in-time recovery, data import, and global tables across multiple regions.

Please ensure that you replace placeholder values like `my-aws-credential-id`, `your-kms-key-id`, bucket names, and region codes with your actual configuration details.
