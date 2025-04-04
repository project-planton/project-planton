# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates an AWS AWS EKS Cluster with default settings.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEksCluster
metadata:
  name: my-basic-aws-eks-cluster
spec:
  awsCredentialId: my-aws-credential-id
  region: us-west-2
  workersManagementMode: MANAGED
```

# Example with Existing VPC

This example creates an AWS EKS Cluster in an existing VPC.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEksCluster
metadata:
  name: my-aws-eks-cluster-with-vpc
spec:
  awsCredentialId: my-aws-credential-id
  region: us-east-1
  vpcId: vpc-0123456789abcdef0
  workersManagementMode: MANAGED
```

# Example with Environment Variables

This example uses environment variables to parameterize the AWS EKS Cluster configuration.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEksCluster
metadata:
  name: my-env-aws-eks-cluster
spec:
  awsCredentialId: ${AWS_CREDENTIAL_ID}
  region: ${AWS_REGION}
  workersManagementMode: ${WORKERS_MANAGEMENT_MODE}
```

In this example, replace the placeholders like `${AWS_CREDENTIAL_ID}` with your actual environment variable names or values.

# Example with Environment Secrets

The below example assumes that the secrets are managed by Planton Cloud's [AWS Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.aws.awssecretsmanager) deployment module.

```yaml
apiVersion: aws.project-planton.org/v1
kind: AwsEksCluster
metadata:
  name: my-secret-aws-eks-cluster
spec:
  awsCredentialId: my-aws-credential-id
  region: us-west-2
  workersManagementMode: MANAGED
  someSecretConfig: ${awssm-my-org-prod-aws-secrets.secret-key}
```

In this example:

- **someSecretConfig** is a placeholder for any configuration value you want to retrieve from AWS Secrets Manager.
- The value before the dot (`awssm-my-org-prod-aws-secrets`) is the ID of the AWS Secrets Manager resource on Planton Cloud.
- The value after the dot (`secret-key`) is the name of the secret within that resource.

---

These examples illustrate how to define an AWS EKS Cluster using the `AwsEksCluster` API resource. Since the `spec` is minimal, we have provided a few examples to demonstrate how to specify different configurations.

Please ensure that you replace placeholder values like `my-aws-credential-id`, `vpc-0123456789abcdef0`, environment variable names, and secret references with your actual configuration details.
