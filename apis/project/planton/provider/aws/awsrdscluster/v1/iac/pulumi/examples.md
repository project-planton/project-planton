Here are a few examples for the `AwsRdsCluster` API resource, showing different configurations for an RDS cluster on AWS. These examples demonstrate basic and more advanced usage of the resource, including autoscaling and security configurations.

---

### Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsCluster
metadata:
  name: my-basic-rds-cluster
spec:
  awsCredentialId: my-aws-cred
  engine: aurora
  engineVersion: "5.6.10a"
  engineMode: provisioned
  instanceType: db.r5.large
  clusterSize: 3
  databaseName: mydb
  masterUser: admin
  masterPassword: supersecretpassword
  vpcId: vpc-abc123
  subnetIds:
    - subnet-12345
    - subnet-67890
  securityGroupIds:
    - sg-00112233
  storageEncrypted: true
  backupWindow: 23:00-01:00
  retentionPeriod: 7
```

---

### Example with Autoscaling

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsCluster
metadata:
  name: auto-scaling-rds-cluster
spec:
  awsCredentialId: my-aws-cred
  engine: aurora-postgresql
  engineVersion: "11.6"
  engineMode: serverless
  instanceType: db.r5.large
  clusterSize: 3
  databaseName: mydb
  masterUser: admin
  masterPassword: supersecretpassword
  vpcId: vpc-abc123
  subnetIds:
    - subnet-12345
    - subnet-67890
  securityGroupIds:
    - sg-00112233
  autoScaling:
    isEnabled: true
    policyType: TargetTrackingScaling
    targetMetrics: CPUUtilization
    targetValue: 70
    minCapacity: 1
    maxCapacity: 5
  storageEncrypted: true
  retentionPeriod: 7
```

---

### Example with Public Access and KMS Encryption

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsCluster
metadata:
  name: public-rds-cluster
spec:
  awsCredentialId: my-aws-cred
  engine: aurora-mysql
  engineVersion: "5.7.12"
  engineMode: provisioned
  instanceType: db.t3.medium
  clusterSize: 2
  databaseName: mypublicdb
  masterUser: admin
  masterPassword: supersecretpassword
  isPubliclyAccessible: true
  vpcId: vpc-abc123
  subnetIds:
    - subnet-abc123
    - subnet-def456
  securityGroupIds:
    - sg-44556677
  storageEncrypted: true
  storageKmsKeyArn: arn:aws:kms:us-west-2:123456789012:key/abcd1234
  backupWindow: 00:00-02:00
  retentionPeriod: 14
```

---

### Example with Performance Insights

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsCluster
metadata:
  name: performance-insights-rds
spec:
  awsCredentialId: my-aws-cred
  engine: aurora
  engineVersion: "5.6.10a"
  engineMode: provisioned
  instanceType: db.r5.large
  clusterSize: 3
  databaseName: perfinsightsdb
  masterUser: admin
  masterPassword: supersecretpassword
  vpcId: vpc-abc123
  subnetIds:
    - subnet-12345
    - subnet-67890
  securityGroupIds:
    - sg-00112233
  isPerformanceInsightsEnabled: true
  performanceInsightsKmsKeyId: arn:aws:kms:us-west-2:123456789012:key/abcd1234
  storageEncrypted: true
  backupWindow: 01:00-03:00
  retentionPeriod: 7
```
