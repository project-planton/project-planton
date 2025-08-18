# Examples for AWSLambda Pulumi Module

## Minimal manifest (S3 zip code)

```yaml
apiVersion: aws.project-planton.org/v1
kind: AWSLambda
metadata:
  name: my-lambda
spec:
  functionName: my-lambda
  roleArn:
    value: arn:aws:iam::<account-id>:role/service-role/lambda-exec-role
  codeSourceType: CODE_SOURCE_TYPE_S3
  runtime: nodejs18.x
  handler: index.handler
  s3:
    bucket: my-artifacts
    key: lambda/hello.zip
```

## Run

```bash
project-planton pulumi preview \
  --manifest ../hack/manifest.yaml \
  --stack organization/<project>/<stack> \
  --module-dir .
```
