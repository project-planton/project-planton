# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates an AWS Lambda function with minimal configuration.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsLambda
metadata:
  name: my-basic-lambda
spec:
  awsCredentialId: my-aws-credential-id
  function:
    handler: index.handler
    runtime: nodejs14.x
    s3Bucket: my-lambda-functions
    s3Key: my-basic-lambda.zip
```

# Example with Environment Variables

This example creates an AWS Lambda function and sets environment variables that can be accessed within the function code.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsLambda
metadata:
  name: my-env-lambda
spec:
  awsCredentialId: my-aws-credential-id
  function:
    handler: app.lambdaHandler
    runtime: python3.8
    s3Bucket: my-lambda-functions
    s3Key: my-env-lambda.zip
    variables:
      DATABASE_NAME: mydatabase
      STAGE: production
```

# Example with Environment Secrets

The below example assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/plantoncloud/planton-cloud-apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsLambda
metadata:
  name: my-secret-lambda
spec:
  awsCredentialId: my-aws-credential-id
  function:
    handler: index.handler
    runtime: nodejs14.x
    s3Bucket: my-lambda-functions
    s3Key: my-secret-lambda.zip
    variables:
      DATABASE_NAME: mydatabase
      DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.database-password}
```

In this example:

- **DATABASE_PASSWORD** is referenced from the GCP Secrets Manager. The value before the dot (`gcpsm-my-org-prod-gcp-secrets`) is the ID of the GCP Secrets Manager resource on Planton Cloud, and the value after the dot (`database-password`) is the name of the secret within that resource.

# Example with VPC Configuration

This example deploys a Lambda function within a VPC, specifying subnet IDs and security group IDs.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsLambda
metadata:
  name: my-vpc-lambda
spec:
  awsCredentialId: my-aws-credential-id
  function:
    handler: index.handler
    runtime: nodejs14.x
    s3Bucket: my-lambda-functions
    s3Key: my-vpc-lambda.zip
    vpcConfig:
      subnetIds:
        - subnet-0123456789abcdef0
        - subnet-0fedcba9876543210
      securityGroupIds:
        - sg-0123456789abcdef0
```

# Example with IAM Role and Policies

This example specifies an IAM role with custom policies for the Lambda function.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsLambda
metadata:
  name: my-custom-iam-lambda
spec:
  awsCredentialId: my-aws-credential-id
  function:
    handler: index.handler
    runtime: nodejs14.x
    s3Bucket: my-lambda-functions
    s3Key: my-custom-iam-lambda.zip
  iamRole:
    customIamPolicyArns:
      - arn:aws:iam::aws:policy/AmazonS3ReadOnlyAccess
    inlineIamPolicy: |
      {
        "Version": "2012-10-17",
        "Statement": [
          {
            "Effect": "Allow",
            "Action": [
              "dynamodb:PutItem",
              "dynamodb:GetItem"
            ],
            "Resource": "arn:aws:dynamodb:*:*:table/MyTable"
          }
        ]
      }
```

# Example with All Available Fields

This comprehensive example includes multiple configurations to demonstrate the full capabilities of the `AwsLambda` resource.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsLambda
metadata:
  name: my-full-config-lambda
spec:
  awsCredentialId: my-aws-credential-id
  function:
    handler: index.handler
    runtime: nodejs14.x
    s3Bucket: my-lambda-functions
    s3Key: my-full-config-lambda.zip
    s3ObjectVersion: "1"
    sourceCodeHash: abcdef1234567890
    memorySize: 256
    timeout: 15
    publish: true
    reservedConcurrentExecutions: 10
    architectures:
      - x86_64
      - arm64
    layers:
      - arn:aws:lambda:us-east-1:123456789012:layer:my-layer:1
    variables:
      KEY1: value1
      KEY2: value2
    deadLetterConfigTargetArn: arn:aws:sqs:us-east-1:123456789012:my-dlq
    kmsKeyArn: arn:aws:kms:us-east-1:123456789012:key/your-kms-key-id
    vpcConfig:
      subnetIds:
        - subnet-0123456789abcdef0
      securityGroupIds:
        - sg-0123456789abcdef0
    tracingConfigMode: Active
    fileSystemConfig:
      arn: arn:aws:elasticfilesystem:us-east-1:123456789012:access-point/fsap-0123456789abcdef0
      localMountPath: /mnt/efs
    ephemeralStorageSize: 1024
  iamRole:
    permissionsBoundary: arn:aws:iam::123456789012:policy/my-permissions-boundary
    cloudwatchLambdaInsightsEnabled: true
    ssmParameterNames:
      - /my/parameter/name
    customIamPolicyArns:
      - arn:aws:iam::aws:policy/AmazonS3FullAccess
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
  cloudwatchLogGroup:
    retentionInDays: 30
    kmsKeyArn: arn:aws:kms:us-east-1:123456789012:key/your-log-group-kms-key
  invokeFunctionPermissions:
    - principal: sns.amazonaws.com
      sourceArn: arn:aws:sns:us-east-1:123456789012:my-sns-topic
    - principal: events.amazonaws.com
      sourceArn: arn:aws:events:us-east-1:123456789012:rule/my-scheduled-rule
```

---

These examples illustrate various configurations of the `AwsLambda` API resource, demonstrating how to define Lambda functions with different features such as environment variables, environment secrets, VPC settings, IAM roles, layers, file system configurations, CloudWatch Logs, and invoke permissions.

Please ensure that you replace placeholder values like `my-aws-credential-id`, `my-lambda-functions`, bucket names, ARNs, and resource IDs with your actual configuration details.
