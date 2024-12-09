# Create using CLI

Create a YAML file using one of the examples shown below. After the YAML file is created, use the command below to apply the configuration:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates a private S3 bucket with default settings.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: S3Bucket
metadata:
  name: my-private-bucket
spec:
  awsCredentialId: my-aws-credential-id
  awsRegion: us-west-2
  isPublic: false
```

# Example with Environment Variables

This example demonstrates how to use environment variables to parameterize the S3 bucket configuration.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: S3Bucket
metadata:
  name: my-env-bucket
spec:
  awsCredentialId: ${AWS_CREDENTIAL_ID}
  awsRegion: ${AWS_REGION}
  isPublic: ${IS_PUBLIC}
```

In this example, replace the placeholders like `${AWS_CREDENTIAL_ID}` with your actual environment variable names or values.

# Example with Environment Secrets

The below example assumes that the secrets are managed by Planton Cloud's [AWS Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.aws.awssecretsmanager) deployment module.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: S3Bucket
metadata:
  name: my-secret-bucket
spec:
  awsCredentialId: my-aws-credential-id
  awsRegion: us-east-1
  isPublic: false
  someSecretConfig: ${awssm-my-org-prod-aws-secrets.secret-key}
```

In this example:

- **someSecretConfig** is a placeholder for any configuration value you want to retrieve from AWS Secrets Manager.
- The value before the dot (`awssm-my-org-prod-aws-secrets`) is the ID of the AWS Secrets Manager resource on Planton Cloud.
- The value after the dot (`secret-key`) is the name of the secret within that resource.

# Example with Public Access

This example creates a public S3 bucket.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: S3Bucket
metadata:
  name: my-public-bucket
spec:
  awsCredentialId: my-aws-credential-id
  awsRegion: us-west-2
  isPublic: true
```

# Example with All Available Fields

This comprehensive example demonstrates the full capabilities of the `S3Bucket` resource.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: S3Bucket
metadata:
  name: my-full-config-bucket
spec:
  awsCredentialId: my-aws-credential-id
  awsRegion: us-east-1
  isPublic: false
```

---

These examples illustrate various configurations of the `S3Bucket` API resource, demonstrating how to define S3 buckets with different features such as public access settings, environment variables, and environment secrets.

Please ensure that you replace placeholder values like `my-aws-credential-id`, `my-private-bucket`, environment variable names, and secret references with your actual configuration details.
