# Examples for AWSLambda Pulumi Module

## Minimal manifest (YAML)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AWSLambda
metadata:
  name: hello-lambda
  org: my-org
spec:
  function_name: hello-lambda
  role_arn: arn:aws:iam::123456789012:role/service-role/hello-role
  code_source_type: CODE_SOURCE_TYPE_S3
  runtime: nodejs18.x
  handler: index.handler
  s3:
    bucket: my-artifacts
    key: lambda/hello.zip
```

## Container image-based Lambda with VPC and layers

```yaml
apiVersion: aws.project-planton.org/v1
kind: AWSLambda
metadata:
  name: payments-worker
  org: my-org
  tags:
    app: payments
    env: prod
spec:
  function_name: payments-worker
  description: Processes payment events from the queue
  role_arn: arn:aws:iam::123456789012:role/service-role/payments-role
  code_source_type: CODE_SOURCE_TYPE_IMAGE
  image_uri: 123456789012.dkr.ecr.us-east-1.amazonaws.com/payments:1.2.3
  memory_mb: 1024
  timeout_seconds: 120
  reserved_concurrency: 10
  environment:
    LOG_LEVEL: info
    QUEUE_URL: https://sqs.us-east-1.amazonaws.com/123456789012/payments
  architecture: ARM64
  subnets:
    - value: subnet-aaaa1111
    - value: subnet-bbbb2222
  security_groups:
    - value: sg-0123456789abcdef0
  layer_arns:
    - arn:aws:lambda:us-east-1:123456789012:layer:powertools:3
  kms_key_arn: arn:aws:kms:us-east-1:123456789012:key/abcde-12345-ffff-9999-0000
```

## CLI flows

Validate manifest:
```bash
project-planton validate --manifest ./lambda.yaml
```

Pulumi deploy:
```bash
project-planton pulumi update --manifest ./lambda.yaml --stack my-org/project/dev --module-dir apis/project/planton/provider/aws/awslambda/v1/iac/pulumi
```

