# AWS RDS Instance Examples

Below are several examples demonstrating how to define an AWS RDS Instance component in
ProjectPlanton. After creating one of these YAML manifests, apply it with Terraform using the ProjectPlanton CLI:

```shell
project-planton tofu apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic RDS Instance

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: basic-rds-instance
spec:
  dbName: "mydb"
  engine: "mysql"
  engineVersion: "8.0.35"
  instanceClass: "db.t3.micro"
  allocatedStorage: 20
  username: "admin"
  manageMasterUserPassword: true
  port: 3306
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  isPubliclyAccessible: true
  storageEncrypted: true
```

This example creates a basic RDS instance:
• MySQL 8.0 engine with t3.micro instance class.
• 20GB allocated storage with encryption.
• Automatic master password management via Secrets Manager.
• Public accessibility for external connections.
• VPC integration with private subnets.

---

## Production RDS Instance with Multi-AZ

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: production-rds-instance
spec:
  dbName: "productiondb"
  engine: "postgres"
  engineVersion: "15.4"
  instanceClass: "db.r5.large"
  allocatedStorage: 100
  maxAllocatedStorage: 200
  username: "admin"
  manageMasterUserPassword: true
  masterUserSecretKmsKeyId: "arn:aws:kms:us-east-1:123456789012:key/production-key"
  port: 5432
  isMultiAz: true
  backupRetentionPeriod: 30
  backupWindow: "02:00-03:00"
  maintenanceWindow: "sun:04:00-sun:05:00"
  subnetIds:
    - "subnet-private-1a"
    - "subnet-private-1b"
  securityGroupIds:
    - "sg-production-db"
  isPubliclyAccessible: false
  storageEncrypted: true
  storageKmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/storage-key"
  deletionProtection: true
  isPerformanceInsightsEnabled: true
  performanceInsightsKmsKeyId: "arn:aws:kms:us-east-1:123456789012:key/pi-key"
  monitoring:
    monitoringInterval: 60
    monitoringRoleArn: "arn:aws:iam::123456789012:role/rds-monitoring-role"
```

This example creates a production-ready instance:
• PostgreSQL 15.4 with r5.large for performance.
• Multi-AZ deployment for high availability.
• 30-day backup retention with encryption.
• Performance Insights for monitoring.
• Enhanced monitoring with 60-second intervals.
• Deletion protection enabled.

---

## Development RDS Instance

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: development-rds-instance
spec:
  dbName: "devdb"
  engine: "mysql"
  engineVersion: "8.0.35"
  instanceClass: "db.t3.medium"
  allocatedStorage: 50
  maxAllocatedStorage: 100
  username: "admin"
  manageMasterUserPassword: true
  port: 3306
  backupRetentionPeriod: 7
  backupWindow: "04:00-05:00"
  maintenanceWindow: "sun:05:00-sun:06:00"
  subnetIds:
    - "subnet-dev-private"
  securityGroupIds:
    - "sg-dev-db"
  isPubliclyAccessible: false
  storageEncrypted: true
  skipFinalSnapshot: true
```

This example creates a development instance:
• MySQL 8.0 with t3.medium for cost efficiency.
• 7-day backup retention for cost control.
• Skip final snapshot for faster cleanup.
• Private network access only.
• Suitable for development and testing.

---

## RDS Instance with Custom Parameters

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: custom-params-rds-instance
spec:
  dbName: "customdb"
  engine: "mysql"
  engineVersion: "8.0.35"
  instanceClass: "db.r5.large"
  allocatedStorage: 100
  username: "admin"
  manageMasterUserPassword: true
  port: 3306
  characterSetName: "utf8mb4"
  parameterGroupName: "custom-mysql8-params"
  optionGroupName: "custom-mysql8-options"
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  storageEncrypted: true
  backupRetentionPeriod: 14
```

This example includes custom configurations:
• Custom parameter group for engine tuning.
• Custom option group for additional features.
• UTF8MB4 character set for full Unicode support.
• 14-day backup retention.
• Suitable for applications requiring custom database settings.

---

## RDS Instance with Replication

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: replica-rds-instance
spec:
  replicateSourceDb: "arn:aws:rds:us-east-1:123456789012:db:source-db-instance"
  instanceClass: "db.r5.large"
  port: 3306
  isMultiAz: true
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  isPubliclyAccessible: false
  storageEncrypted: true
  backupRetentionPeriod: 7
  copyTagsToSnapshot: true
```

This example creates a read replica:
• Replicates from an existing RDS instance.
• Multi-AZ deployment for high availability.
• Inherits engine and version from source.
• Copy tags to snapshots for organization.
• Suitable for read scaling and disaster recovery.

---

## RDS Instance with Point-in-Time Recovery

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: pitr-rds-instance
spec:
  dbName: "pitrdb"
  engine: "postgres"
  engineVersion: "15.4"
  instanceClass: "db.r5.large"
  allocatedStorage: 100
  username: "admin"
  manageMasterUserPassword: true
  port: 5432
  isMultiAz: true
  backupRetentionPeriod: 35
  backupWindow: "01:00-02:00"
  maintenanceWindow: "sun:03:00-sun:04:00"
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  storageEncrypted: true
  restoreToPointInTime:
    sourceDbInstanceIdentifier: "source-db-instance"
    useLatestRestorableTime: true
```

This example includes point-in-time recovery:
• PostgreSQL 15.4 with 35-day backup retention.
• Point-in-time recovery from source instance.
• Multi-AZ deployment for high availability.
• Suitable for disaster recovery scenarios.

---

## RDS Instance with Enhanced Monitoring

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: monitored-rds-instance
spec:
  dbName: "monitoreddb"
  engine: "mysql"
  engineVersion: "8.0.35"
  instanceClass: "db.r5.large"
  allocatedStorage: 100
  username: "admin"
  manageMasterUserPassword: true
  port: 3306
  isMultiAz: true
  backupRetentionPeriod: 14
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  storageEncrypted: true
  isPerformanceInsightsEnabled: true
  performanceInsightsKmsKeyId: "arn:aws:kms:us-east-1:123456789012:key/pi-key"
  monitoring:
    monitoringInterval: 30
    monitoringRoleArn: "arn:aws:iam::123456789012:role/rds-monitoring-role"
```

This example includes comprehensive monitoring:
• Performance Insights for query analysis.
• Enhanced monitoring with 30-second intervals.
• KMS encryption for performance data.
• Multi-AZ deployment for high availability.
• Suitable for production workloads requiring detailed monitoring.

---

## RDS Instance with Storage Optimization

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsInstance
metadata:
  name: storage-optimized-rds-instance
spec:
  dbName: "storageoptdb"
  engine: "mysql"
  engineVersion: "8.0.35"
  instanceClass: "db.r5.large"
  allocatedStorage: 500
  maxAllocatedStorage: 1000
  storageType: "io1"
  iops: 3000
  username: "admin"
  manageMasterUserPassword: true
  port: 3306
  isMultiAz: true
  backupRetentionPeriod: 14
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  storageEncrypted: true
```

This example is optimized for storage performance:
• Provisioned IOPS (io1) storage type.
• 3000 IOPS for high-performance storage.
• Auto-scaling storage from 500GB to 1TB.
• Multi-AZ deployment for high availability.
• Suitable for I/O-intensive workloads.

---

## After Deploying

Once you've applied your manifest with ProjectPlanton tofu, you can confirm that the RDS instance is active via the AWS console or by
using the AWS CLI:

```shell
aws rds describe-db-instances --db-instance-identifier <your-instance-id>
```

For detailed instance information:

```shell
aws rds describe-db-instances --db-instance-identifier <your-instance-id> --query 'DBInstances[0].{Endpoint:Endpoint,Port:Endpoint.Port,Status:DBInstanceStatus}'
```

To test database connectivity:

```shell
mysql -h <your-endpoint> -P <port> -u <username> -p <db-name>
```

This will show the RDS instance details including endpoint, status, and configuration information.
