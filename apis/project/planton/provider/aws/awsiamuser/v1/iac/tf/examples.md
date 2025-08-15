# AWS IAM User Examples

Below are several examples demonstrating how to define an AWS IAM User component in
ProjectPlanton. After creating one of these YAML manifests, apply it with Terraform using the ProjectPlanton CLI:

```shell
project-planton tofu apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic IAM User

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: basic-iam-user
spec:
  userName: "basic-user"
  managedPolicyArns:
    - "arn:aws:iam::aws:policy/ReadOnlyAccess"
```

This example creates a basic IAM user:
• User name following AWS naming conventions.
• Read-only access via managed policy.
• Access key automatically created (default behavior).
• Suitable for read-only operations and monitoring.

---

## IAM User for CI/CD Pipeline

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: cicd-iam-user
spec:
  userName: "cicd-user"
  managedPolicyArns:
    - "arn:aws:iam::aws:policy/AmazonS3FullAccess"
    - "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryPowerUser"
  inlinePolicies:
    deploymentAccess:
      Version: "2012-10-17"
      Statement:
        - Effect: "Allow"
          Action:
            - "ecs:DescribeServices"
            - "ecs:UpdateService"
            - "ecs:DescribeTasks"
            - "ecs:RunTask"
          Resource: "*"
```

This example creates a CI/CD user:
• Full S3 access for artifact storage.
• ECR access for container image management.
• Custom inline policy for ECS deployment.
• Access key for programmatic access.
• Suitable for automated deployment pipelines.

---

## IAM User for Application Access

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: app-iam-user
spec:
  userName: "application-user"
  managedPolicyArns:
    - "arn:aws:iam::aws:policy/AmazonDynamoDBFullAccess"
    - "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"
  inlinePolicies:
    applicationSpecific:
      Version: "2012-10-17"
      Statement:
        - Effect: "Allow"
          Action:
            - "secretsmanager:GetSecretValue"
            - "kms:Decrypt"
          Resource:
            - "arn:aws:secretsmanager:*:*:secret:app-*"
            - "arn:aws:kms:*:*:key/*"
```

This example creates an application user:
• DynamoDB full access for data operations.
• S3 read-only access for data retrieval.
• Custom policy for secrets and KMS access.
• Access key for application authentication.
• Suitable for application service accounts.

---

## IAM User with No Access Keys

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: console-only-user
spec:
  userName: "console-user"
  managedPolicyArns:
    - "arn:aws:iam::aws:policy/ReadOnlyAccess"
  disableAccessKeys: true
```

This example creates a console-only user:
• Read-only access via managed policy.
• Access keys disabled for security.
• Console access only.
• Suitable for human users who only need console access.

---

## IAM User for Development

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: dev-iam-user
spec:
  userName: "developer-user"
  managedPolicyArns:
    - "arn:aws:iam::aws:policy/PowerUserAccess"
  inlinePolicies:
    developmentAccess:
      Version: "2012-10-17"
      Statement:
        - Effect: "Allow"
          Action:
            - "ec2:*"
            - "rds:*"
            - "lambda:*"
            - "cloudformation:*"
          Resource: "*"
        - Effect: "Deny"
          Action:
            - "iam:*"
            - "organizations:*"
          Resource: "*"
```

This example creates a developer user:
• Power user access for development work.
• Custom policy for development services.
• Explicit denials for sensitive operations.
• Access key for development tools.
• Suitable for development team members.

---

## IAM User for Monitoring

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: monitoring-iam-user
spec:
  userName: "monitoring-user"
  managedPolicyArns:
    - "arn:aws:iam::aws:policy/CloudWatchReadOnlyAccess"
    - "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"
  inlinePolicies:
    monitoringAccess:
      Version: "2012-10-17"
      Statement:
        - Effect: "Allow"
          Action:
            - "logs:DescribeLogGroups"
            - "logs:DescribeLogStreams"
            - "logs:GetLogEvents"
            - "xray:GetTraceSummaries"
            - "xray:BatchGetTraces"
          Resource: "*"
```

This example creates a monitoring user:
• CloudWatch read-only access for metrics.
• S3 read-only access for log analysis.
• Custom policy for logs and X-Ray access.
• Access key for monitoring tools.
• Suitable for monitoring and observability teams.

---

## IAM User for Backup Operations

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: backup-iam-user
spec:
  userName: "backup-user"
  managedPolicyArns:
    - "arn:aws:iam::aws:policy/AmazonS3FullAccess"
    - "arn:aws:iam::aws:policy/AmazonRDSFullAccess"
  inlinePolicies:
    backupOperations:
      Version: "2012-10-17"
      Statement:
        - Effect: "Allow"
          Action:
            - "rds:CreateDBSnapshot"
            - "rds:CopyDBSnapshot"
            - "rds:RestoreDBInstanceFromDBSnapshot"
            - "s3:PutObject"
            - "s3:GetObject"
            - "s3:DeleteObject"
          Resource:
            - "arn:aws:rds:*:*:db:*"
            - "arn:aws:rds:*:*:snapshot:*"
            - "arn:aws:s3:::backup-bucket/*"
```

This example creates a backup user:
• Full S3 access for backup storage.
• Full RDS access for database operations.
• Custom policy for backup-specific operations.
• Access key for automated backup scripts.
• Suitable for backup and disaster recovery operations.

---

## IAM User with Minimal Configuration

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsIamUser
metadata:
  name: minimal-iam-user
spec:
  userName: "minimal-user"
```

This example creates a minimal IAM user:
• Only required user name specified.
• No policies attached (no permissions).
• Access key automatically created.
• Suitable as a starting point for custom permissions.

---

## After Deploying

Once you've applied your manifest with ProjectPlanton tofu, you can confirm that the IAM user is active via the AWS console or by
using the AWS CLI:

```shell
aws iam get-user --user-name <your-user-name>
```

For detailed user information including attached policies:

```shell
aws iam list-attached-user-policies --user-name <your-user-name>
```

To list inline policies:

```shell
aws iam list-user-policies --user-name <your-user-name>
```

To list access keys (if created):

```shell
aws iam list-access-keys --user-name <your-user-name>
```

This will show the IAM user details including attached managed policies, inline policies, and access keys.

