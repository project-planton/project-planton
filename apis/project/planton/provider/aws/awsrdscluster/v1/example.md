# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsCloudFront
metadata:
  name: my-cloudfront-distribution
spec:
  awsCredentialId: my-aws-credential-id
```

# Example with Environment Information

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsCloudFront
metadata:
  name: my-env-cloudfront-distribution
spec:
  environmentInfo:
    envId: production-environment
  awsCredentialId: prod-aws-credential-id
```

# Example with Stack Job Settings

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsCloudFront
metadata:
  name: advanced-cloudfront-distribution
spec:
  stackJobSettings:
    pulumiBackendCredentialId: my-pulumi-backend-credential
    stackJobRunnerId: my-stack-job-runner
  awsCredentialId: advanced-aws-credential-id
```

# Example with All Available Fields

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsCloudFront
metadata:
  name: full-config-cloudfront-distribution
spec:
  environmentInfo:
    envId: dev-environment
  stackJobSettings:
    pulumiBackendCredentialId: dev-pulumi-backend-credential
    stackJobRunnerId: dev-stack-job-runner
  awsCredentialId: dev-aws-credential-id
```

Please note that since the `spec` is currently empty in the API resource definition, these examples are illustrative and may not reflect actual configuration options. Future updates to the module will include additional fields in the `spec` to allow for more detailed configuration of AWS CloudFront distributions.
