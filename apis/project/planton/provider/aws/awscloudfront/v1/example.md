# Create using CLI

Create a YAML file using one of the examples shown below. After the YAML is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: aws.project.planton/v1
kind: AwsCloudFront
metadata:
  name: my-cloudfront-distribution
spec:
  awsCredentialId: my-aws-credential
```

# Example with Custom AWS Credential

```yaml
apiVersion: aws.project.planton/v1
kind: AwsCloudFront
metadata:
  name: custom-cloudfront-distribution
spec:
  awsCredentialId: custom-aws-credential
```

# Example with Additional Metadata

```yaml
apiVersion: aws.project.planton/v1
kind: AwsCloudFront
metadata:
  name: another-cloudfront-instance
  labels:
    environment: production
spec:
  awsCredentialId: prod-aws-credential
```

# Notes

- **awsCredentialId**: This field is required and should reference the ID of the AWS credentials you have set up. Ensure that the credential ID matches one that is configured in your environment.
- Since the `spec` section is minimal for this API resource and the module is not fully implemented, the examples provided are straightforward and focus primarily on the required `awsCredentialId` field.

# Conclusion

These examples demonstrate how to create an `AwsCloudFront` resource using minimal configuration. As the module gets further developed, more comprehensive examples will be provided to showcase additional features and capabilities
