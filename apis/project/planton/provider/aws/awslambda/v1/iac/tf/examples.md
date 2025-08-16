# AWS Lambda Examples

Below are several examples demonstrating how to define an AWS Lambda component in
ProjectPlanton. After creating one of these YAML manifests, apply it with Terraform using the ProjectPlanton CLI:

```shell
project-planton tofu apply --manifest <yaml-path> --stack <stack-name>
```

---

## Basic Lambda Function

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: basic-lambda
spec:
  function:
    handler: "index.handler"
    runtime: "nodejs18.x"
    s3Bucket: "my-lambda-functions"
    s3Key: "basic-lambda.zip"
    description: "Basic Lambda function with minimal configuration"
```

This example creates a basic Lambda function:
• Node.js 18.x runtime environment.
• S3-based deployment package.
• Default memory (128MB) and timeout (3 seconds).
• Automatic IAM role creation with basic permissions.
• CloudWatch logging enabled by default.

---

## Lambda Function with Environment Variables

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: env-lambda
spec:
  function:
    handler: "app.lambdaHandler"
    runtime: "python3.11"
    s3Bucket: "my-lambda-functions"
    s3Key: "env-lambda.zip"
    description: "Lambda function with environment variables"
    variables:
      DATABASE_NAME: "mydatabase"
      STAGE: "production"
      LOG_LEVEL: "INFO"
    memorySize: 256
    timeout: 30
```

This example includes environment configuration:
• Python 3.11 runtime for data processing.
• Environment variables for configuration.
• Increased memory and timeout for performance.
• Suitable for applications requiring configuration.

---

## Lambda Function with Container Image

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: container-lambda
spec:
  function:
    imageUri: "123456789012.dkr.ecr.us-east-1.amazonaws.com/my-lambda:latest"
    packageType: "Image"
    description: "Lambda function using container image"
    memorySize: 512
    timeout: 60
    architectures:
      - "x86_64"
      - "arm64"
```

This example uses container deployment:
• Container image from ECR repository.
• Multi-architecture support (x86_64 and ARM64).
• Larger memory allocation for container workloads.
• Extended timeout for container startup.
• Suitable for complex applications with dependencies.

---

## Lambda Function with VPC Integration

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: vpc-lambda
spec:
  function:
    handler: "index.handler"
    runtime: "nodejs18.x"
    s3Bucket: "my-lambda-functions"
    s3Key: "vpc-lambda.zip"
    description: "Lambda function with VPC integration"
    vpcConfig:
      subnetIds:
        - "subnet-0123456789abcdef0"
        - "subnet-0fedcba9876543210"
      securityGroupIds:
        - "sg-0123456789abcdef0"
    memorySize: 512
    timeout: 30
```

This example includes VPC integration:
• Private subnets for network isolation.
• Security groups for access control.
• VPC access for database and internal service connectivity.
• Increased memory for VPC networking overhead.
• Suitable for applications requiring private network access.

---

## Lambda Function with Custom IAM Role

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: custom-iam-lambda
spec:
  function:
    handler: "index.handler"
    runtime: "nodejs18.x"
    s3Bucket: "my-lambda-functions"
    s3Key: "custom-iam-lambda.zip"
    description: "Lambda function with custom IAM permissions"
  iamRole:
    cloudwatchLambdaInsightsEnabled: true
    ssmParameterNames:
      - "/my/parameter/name"
    customIamPolicyArns:
      - "arn:aws:iam::aws:policy/AmazonS3FullAccess"
    inlineIamPolicy: |
      {
        "Version": "2012-10-17",
        "Statement": [
          {
            "Effect": "Allow",
            "Action": "dynamodb:*",
            "Resource": "*"
          }
        ]
      }
```

This example includes custom IAM configuration:
• CloudWatch Lambda Insights for enhanced monitoring.
• SSM Parameter Store access for configuration.
• S3 full access via managed policy.
• Custom inline policy for DynamoDB access.
• Suitable for applications requiring specific AWS service access.

---

## Lambda Function with CloudWatch Logging

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: logging-lambda
spec:
  function:
    handler: "index.handler"
    runtime: "nodejs18.x"
    s3Bucket: "my-lambda-functions"
    s3Key: "logging-lambda.zip"
    description: "Lambda function with enhanced logging"
  cloudwatchLogGroup:
    retentionInDays: 30
    kmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/log-encryption-key"
```

This example includes enhanced logging:
• 30-day log retention for compliance.
• KMS encryption for log data security.
• Structured logging for better observability.
• Suitable for production applications with logging requirements.

---

## Lambda Function with Invoke Permissions

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: event-driven-lambda
spec:
  function:
    handler: "index.handler"
    runtime: "nodejs18.x"
    s3Bucket: "my-lambda-functions"
    s3Key: "event-driven-lambda.zip"
    description: "Lambda function with event source permissions"
  invokeFunctionPermissions:
    - principal: "sns.amazonaws.com"
      sourceArn: "arn:aws:sns:us-east-1:123456789012:my-sns-topic"
    - principal: "events.amazonaws.com"
      sourceArn: "arn:aws:events:us-east-1:123456789012:rule/my-scheduled-rule"
```

This example includes event source permissions:
• SNS topic invocation for event processing.
• EventBridge rule invocation for scheduled execution.
• Event-driven architecture support.
• Suitable for event-driven applications and scheduled tasks.

---

## Production Lambda Function

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: production-lambda
spec:
  function:
    handler: "index.handler"
    runtime: "nodejs18.x"
    s3Bucket: "my-lambda-functions"
    s3Key: "production-lambda.zip"
    description: "Production Lambda function with comprehensive configuration"
    memorySize: 1024
    timeout: 60
    publish: true
    reservedConcurrentExecutions: 10
    architectures:
      - "x86_64"
    layers:
      - "arn:aws:lambda:us-east-1:123456789012:layer:my-layer:1"
    variables:
      ENVIRONMENT: "production"
      LOG_LEVEL: "WARN"
    deadLetterConfigTargetArn: "arn:aws:sqs:us-east-1:123456789012:my-dlq"
    kmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/encryption-key"
    vpcConfig:
      subnetIds:
        - "subnet-private-1a"
        - "subnet-private-1b"
      securityGroupIds:
        - "sg-lambda-production"
    tracingConfigMode: "Active"
    ephemeralStorageSize: 1024
  iamRole:
    cloudwatchLambdaInsightsEnabled: true
    ssmParameterNames:
      - "/production/database/url"
      - "/production/api/key"
    customIamPolicyArns:
      - "arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess"
  cloudwatchLogGroup:
    retentionInDays: 90
    kmsKeyArn: "arn:aws:kms:us-east-1:123456789012:key/log-encryption-key"
```

This example is production-ready:
• High memory allocation for performance.
• Reserved concurrency for predictable scaling.
• Lambda layers for shared code.
• Dead letter queue for error handling.
• KMS encryption for environment variables.
• VPC integration for private network access.
• X-Ray tracing for distributed tracing.
• Enhanced ephemeral storage.
• Comprehensive IAM permissions.
• Extended log retention for compliance.

---

## Development Lambda Function

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsLambda
metadata:
  name: development-lambda
spec:
  function:
    handler: "index.handler"
    runtime: "nodejs18.x"
    s3Bucket: "my-lambda-functions"
    s3Key: "development-lambda.zip"
    description: "Development Lambda function"
    memorySize: 256
    timeout: 30
    variables:
      ENVIRONMENT: "development"
      LOG_LEVEL: "DEBUG"
      DEBUG: "true"
  cloudwatchLogGroup:
    retentionInDays: 7
```

This example is optimized for development:
• Moderate memory allocation for cost efficiency.
• Debug logging enabled.
• Development environment variables.
• Short log retention for cost control.
• Suitable for development and testing environments.

---

## After Deploying

Once you've applied your manifest with ProjectPlanton tofu, you can confirm that the Lambda function is active via the AWS console or by
using the AWS CLI:

```shell
aws lambda get-function --function-name <your-function-name>
```

For detailed function configuration:

```shell
aws lambda get-function-configuration --function-name <your-function-name>
```

To test the function:

```shell
aws lambda invoke --function-name <your-function-name> --payload '{"key": "value"}' response.json
```

This will show the Lambda function details including configuration, environment variables, and execution results.
