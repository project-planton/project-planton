# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates an AWS Fargate service with minimal configuration.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsFargate
metadata:
  name: my-fargate-service
spec:
  awsCredentialId: my-aws-credential-id
```

# Example with Environment Information

This example includes environment information, specifying the environment ID.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsFargate
metadata:
  name: my-env-fargate-service
spec:
  environmentInfo:
    envId: production-environment
  awsCredentialId: prod-aws-credential-id
```

# Example with Stack Job Settings

This example includes stack job settings for more advanced configurations.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsFargate
metadata:
  name: my-advanced-fargate-service
spec:
  stackJobSettings:
    pulumiBackendCredentialId: my-pulumi-backend-credential
    stackJobRunnerId: my-stack-job-runner
  awsCredentialId: advanced-aws-credential-id
```

---

Please note that since the `spec` is currently empty in the API resource definition, these examples are illustrative and may not reflect actual configuration options. Future updates to the module will include additional fields in the `spec` to allow for more detailed configuration of AWS Fargate services.

