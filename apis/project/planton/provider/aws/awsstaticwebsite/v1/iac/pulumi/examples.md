# Create using CLI

Create a YAML file using one of the examples shown below. After the YAML is created, use the command below to apply.

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsStaticWebsite
metadata:
  name: my-static-website
spec:
  awsCredentialId: my-aws-credential-id
```

# Example with Environment Info

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsStaticWebsite
metadata:
  name: my-static-website
spec:
  awsCredentialId: my-aws-credential-id
  environmentInfo:
    envId: production
```

# Example with Stack Job Settings

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsStaticWebsite
metadata:
  name: my-static-website
spec:
  awsCredentialId: my-aws-credential-id
  stackJobSettings:
    jobTimeout: 3600
```

# Example with Website Configuration

*(Note: This example includes speculative fields since the `spec` is currently empty.)*

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsStaticWebsite
metadata:
  name: my-static-website
spec:
  awsCredentialId: my-aws-credential-id
  websiteConfig:
    indexDocument: index.html
    errorDocument: error.html
```

# Notes

Since the `spec` is currently empty and the module is not completely implemented, these examples are provided for illustrative purposes. They demonstrate how you would structure your YAML configuration files to create an `AwsStaticWebsite` resource using the API once the module is fully implemented.

Remember to replace placeholders like `my-aws-credential-id` and `my-static-website` with your actual AWS credential ID and desired resource name.