## Minimal manifest (YAML)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AWSLambda
metadata:
  name: hello-lambda
spec:
  functionName: hello-lambda
  roleArn:
    value: arn:aws:iam::123456789012:role/service-role/hello-role
  codeSourceType: CODE_SOURCE_TYPE_S3
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
spec:
  functionName: payments-worker
  description: Processes payment events from the queue
  roleArn:
    value: arn:aws:iam::123456789012:role/service-role/payments-role
  codeSourceType: CODE_SOURCE_TYPE_IMAGE
  imageUri: 123456789012.dkr.ecr.us-east-1.amazonaws.com/payments:1.2.3
  memoryMb: 1024
  timeoutSeconds: 120
  reservedConcurrency: 10
  environment:
    LOG_LEVEL: info
    QUEUE_URL: https://sqs.us-east-1.amazonaws.com/123456789012/payments
  architecture: ARM64
  subnets:
    - value: subnet-aaaa1111
    - value: subnet-bbbb2222
  securityGroups:
    - value: sg-0123456789abcdef0
  layerArns:
    - value: arn:aws:lambda:us-east-1:123456789012:layer:powertools:3
  kmsKeyArn:
    value: arn:aws:kms:us-east-1:123456789012:key/abcde-12345-ffff-9999-0000
```

## CLI flows (Pulumi)

```bash
# Preview
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .

# Apply
project-planton pulumi update \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir . \
  --yes

# Refresh
project-planton pulumi refresh \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .

# Destroy
project-planton pulumi destroy \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```


