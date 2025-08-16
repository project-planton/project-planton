# AWS RDS Cluster Examples

Below are several examples demonstrating how to define an AWS RDS Cluster component in
ProjectPlanton. After creating one of these YAML manifests, apply it with Terraform using the ProjectPlanton CLI:

```shell
project-planton tofu apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic RDS Cluster

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: basic-rds-cluster
spec:
  engine: "aurora-mysql"
  engineVersion: "5.7.mysql_aurora.2.11.2"
  engineMode: "provisioned"
  clusterFamily: "aurora-mysql5.7"
  instanceType: "db.r5.large"
  clusterSize: 2
  databaseName: "mydb"
  masterUser: "admin"
  manageMasterUserPassword: true
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  storageEncrypted: true
  backupWindow: "03:00-04:00"
  retentionPeriod: 7
```

This example creates a basic RDS cluster:
• Aurora MySQL 5.7 engine with provisioned mode.
• 2 instances for high availability.
• Automatic master password management via Secrets Manager.
• VPC integration with private subnets.
• Storage encryption enabled.
• 7-day backup retention.

---

## Serverless RDS Cluster

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: serverless-rds-cluster
spec:
  engine: "aurora-postgresql"
  engineVersion: "13.9"
  engineMode: "serverless"
  clusterFamily: "aurora-postgresql13"
  instanceType: "db.serverless"
  clusterSize: 1
  databaseName: "serverlessdb"
  masterUser: "admin"
  manageMasterUserPassword: true
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  serverlessv2ScalingConfiguration:
    minCapacity: 0.5
    maxCapacity: 16.0
  backupWindow: "02:00-03:00"
  retentionPeriod: 7
```

This example creates a serverless Aurora cluster:
• Aurora PostgreSQL 13 with serverless mode.
• Auto-scaling from 0.5 to 16 ACU.
• Pay-per-use pricing model.
• Automatic scaling based on demand.
• Suitable for variable workloads.

---

## Production RDS Cluster with High Availability

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: production-rds-cluster
spec:
  engine: "aurora-mysql"
  engineVersion: "5.7.mysql_aurora.2.11.2"
  engineMode: "provisioned"
  clusterFamily: "aurora-mysql5.7"
  instanceType: "db.r5.xlarge"
  clusterSize: 3
  databaseName: "productiondb"
  masterUser: "admin"
  manageMasterUserPassword: true
  masterUserSecretKmsKeyId: "arn:aws:kms:us-east-1:123456789012:key/production-key"
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-private-1a"
    - "subnet-private-1b"
    - "subnet-private-1c"
  securityGroupIds:
    - "sg-production-db"
  storageEncrypted: true
  storageKmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/storage-key"
  backupWindow: "01:00-02:00"
  retentionPeriod: 30
  deletionProtection: true
  isPerformanceInsightsEnabled: true
  performanceInsightsKmsKeyId: "arn:aws:kms:us-east-1:123456789012:key/pi-key"
  enhancedMonitoringRoleEnabled: true
  rdsMonitoringInterval: 60
```

This example creates a production-ready cluster:
• High-performance r5.xlarge instances.
• 3 instances for maximum availability.
• KMS encryption for all data.
• 30-day backup retention.
• Deletion protection enabled.
• Performance Insights for monitoring.
• Enhanced monitoring with 60-second intervals.

---

## Multi-AZ RDS Cluster

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: multiaz-rds-cluster
spec:
  engine: "aurora-postgresql"
  engineVersion: "13.9"
  engineMode: "provisioned"
  clusterFamily: "aurora-postgresql13"
  instanceType: "db.r5.large"
  clusterSize: 4
  databaseName: "multiazdb"
  masterUser: "admin"
  manageMasterUserPassword: true
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-az1-private"
    - "subnet-az2-private"
    - "subnet-az3-private"
  securityGroupIds:
    - "sg-multiaz-db"
  storageEncrypted: true
  backupWindow: "00:00-01:00"
  retentionPeriod: 14
  maintenanceWindow: "sun:04:00-sun:05:00"
  allowMajorVersionUpgrade: false
```

This example creates a multi-AZ cluster:
• 4 instances across multiple availability zones.
• Aurora PostgreSQL for advanced features.
• Cross-AZ high availability.
• 14-day backup retention.
• Maintenance window on Sundays.
• Major version upgrades disabled for stability.

---

## Development RDS Cluster

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: development-rds-cluster
spec:
  engine: "aurora-mysql"
  engineVersion: "5.7.mysql_aurora.2.11.2"
  engineMode: "provisioned"
  clusterFamily: "aurora-mysql5.7"
  instanceType: "db.t3.medium"
  clusterSize: 1
  databaseName: "devdb"
  masterUser: "admin"
  manageMasterUserPassword: true
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-dev-private"
  securityGroupIds:
    - "sg-dev-db"
  storageEncrypted: true
  backupWindow: "04:00-05:00"
  retentionPeriod: 3
  skipFinalSnapshot: true
```

This example creates a development cluster:
• Single t3.medium instance for cost efficiency.
• 3-day backup retention for cost control.
• Skip final snapshot for faster cleanup.
• Suitable for development and testing.

---

## Publicly Accessible RDS Cluster

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: public-rds-cluster
spec:
  engine: "aurora-mysql"
  engineVersion: "5.7.mysql_aurora.2.11.2"
  engineMode: "provisioned"
  clusterFamily: "aurora-mysql5.7"
  instanceType: "db.r5.large"
  clusterSize: 2
  databaseName: "publicdb"
  masterUser: "admin"
  manageMasterUserPassword: true
  isPubliclyAccessible: true
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-public-1a"
    - "subnet-public-1b"
  securityGroupIds:
    - "sg-public-db"
  storageEncrypted: true
  backupWindow: "03:00-04:00"
  retentionPeriod: 7
```

This example creates a publicly accessible cluster:
• Public subnets for external access.
• Security groups for access control.
• Suitable for applications requiring external database access.
• 2 instances for high availability.

---

## RDS Cluster with Custom Parameters

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: custom-params-rds-cluster
spec:
  engine: "aurora-mysql"
  engineVersion: "5.7.mysql_aurora.2.11.2"
  engineMode: "provisioned"
  clusterFamily: "aurora-mysql5.7"
  instanceType: "db.r5.large"
  clusterSize: 2
  databaseName: "customdb"
  masterUser: "admin"
  manageMasterUserPassword: true
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  clusterParameterGroupName: "custom-aurora-mysql5.7"
  storageEncrypted: true
  backupWindow: "02:00-03:00"
  retentionPeriod: 7
  databasePort: 3306
  iamDatabaseAuthenticationEnabled: true
```

This example includes custom configurations:
• Custom parameter group for engine tuning.
• IAM database authentication enabled.
• Custom database port configuration.
• Suitable for applications requiring custom database settings.

---

## RDS Cluster with Auto Scaling

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsRdsCluster
metadata:
  name: autoscaling-rds-cluster
spec:
  engine: "aurora-mysql"
  engineVersion: "5.7.mysql_aurora.2.11.2"
  engineMode: "provisioned"
  clusterFamily: "aurora-mysql5.7"
  instanceType: "db.r5.large"
  clusterSize: 2
  databaseName: "autoscaledb"
  masterUser: "admin"
  manageMasterUserPassword: true
  vpcId: "vpc-0123456789abcdef0"
  subnetIds:
    - "subnet-0123456789abcdef0"
    - "subnet-0fedcba9876543210"
  securityGroupIds:
    - "sg-0123456789abcdef0"
  storageEncrypted: true
  backupWindow: "01:00-02:00"
  retentionPeriod: 7
  autoScaling:
    isEnabled: true
    policyType: "TargetTrackingScaling"
    targetMetrics: "CPUUtilization"
    targetValue: 70.0
    minCapacity: 2
    maxCapacity: 5
```

This example includes auto scaling:
• Target tracking based on CPU utilization.
• Auto-scaling from 2 to 5 instances.
• 70% CPU threshold for scaling.
• Suitable for variable workloads.

---

## After Deploying

Once you've applied your manifest with ProjectPlanton tofu, you can confirm that the RDS cluster is active via the AWS console or by
using the AWS CLI:

```shell
aws rds describe-db-clusters --db-cluster-identifier <your-cluster-id>
```

For detailed cluster information:

```shell
aws rds describe-db-cluster-parameters --db-cluster-parameter-group-name <your-parameter-group>
```

To list cluster instances:

```shell
aws rds describe-db-cluster-instances --db-cluster-identifier <your-cluster-id>
```

This will show the RDS cluster details including endpoints, instance status, and configuration information.
