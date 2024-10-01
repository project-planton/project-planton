# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates an AWS Elastic Container Service (ECS) resource with default settings.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ElasticContainerService
metadata:
  name: my-basic-ecs
spec:
  awsCredentialId: my-aws-credential-id
```

# Example with Environment Variables

Even though the `spec` is currently empty, here's an example that includes an environment variable reference.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ElasticContainerService
metadata:
  name: my-env-ecs
spec:
  awsCredentialId: ${AWS_CREDENTIAL_ID}
```

---

Since the `spec` for `ElasticContainerService` is currently empty, these examples are minimal and primarily demonstrate how to reference the resource in your configuration. Future updates to the module will allow for more detailed configurations.
