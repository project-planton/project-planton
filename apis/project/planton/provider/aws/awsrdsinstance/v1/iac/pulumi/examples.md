# Create using CLI

Create a YAML file using one of the examples shown below. After the YAML is created, use the command below to apply.

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsInstance
metadata:
  name: my-rds-instance
spec:
  awsCredentialId: my-aws-credential-id
  db_name: mydb
  engine: mysql
  engine_version: 8.0
  instance_class: db.t3.micro
  allocated_storage: 20
  username: admin
  password: adminpassword
  port: 3306
  subnet_ids:
    - subnet-12345678
    - subnet-87654321
  security_group_ids:
    - sg-12345678
  is_publicly_accessible: true
```

# Example with Multi-AZ and Backup Settings

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsInstance
metadata:
  name: my-rds-instance
spec:
  awsCredentialId: my-aws-credential-id
  db_name: mydb
  engine: postgres
  engine_version: 13.3
  instance_class: db.t3.medium
  allocated_storage: 50
  max_allocated_storage: 100
  username: admin
  password: supersecurepassword
  port: 5432
  isMultiAz: true
  backup_retention_period: 7
  backup_window: "02:00-03:00"
  maintenance_window: "Sun:23:00-Mon:00:30"
  subnet_ids:
    - subnet-12345678
    - subnet-87654321
  security_group_ids:
    - sg-12345678
  is_publicly_accessible: false
```

# Example with Replication

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsInstance
metadata:
  name: replica-rds-instance
spec:
  awsCredentialId: my-aws-credential-id
  replicate_source_db: arn:aws:rds:us-west-2:123456789012:db:source-db-instance
  instance_class: db.t3.large
  port: 3306
  isMultiAz: true
  subnet_ids:
    - subnet-12345678
    - subnet-87654321
  security_group_ids:
    - sg-12345678
  is_publicly_accessible: false
```

# Example with Performance Insights and Monitoring

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsRdsInstance
metadata:
  name: monitored-rds-instance
spec:
  awsCredentialId: my-aws-credential-id
  db_name: mydb
  engine: mysql
  engine_version: 8.0
  instance_class: db.r5.large
  allocated_storage: 100
  username: admin
  password: securepassword123
  port: 3306
  subnet_ids:
    - subnet-12345678
    - subnet-87654321
  security_group_ids:
    - sg-12345678
  performance_insights:
    is_enabled: true
    kms_key_id: arn:aws:kms:us-west-2:123456789012:key/my-key
    retention_period: 731
  monitoring:
    monitoring_interval: 30
    monitoring_role_arn: arn:aws:iam::123456789012:role/monitoring-role
  is_publicly_accessible: true
```