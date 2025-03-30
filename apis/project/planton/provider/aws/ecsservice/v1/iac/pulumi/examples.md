# Create using CLI

Create a YAML file using the examples shown below. After the YAML file is created, use the following command to apply:

```shell
planton apply -f <yaml-path>
```

# Basic Example

This basic example creates an AWS VPC with default settings.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsVpc
metadata:
  name: my-basic-vpc
spec:
  awsCredentialId: my-aws-credential-id
  vpcCidr: 10.0.0.0/16
  availabilityZones:
    - us-west-2a
    - us-west-2b
  subnetsPerAvailabilityZone: 1
  subnetSize: 256
  isNatGatewayEnabled: false
  isDnsHostnamesEnabled: true
  isDnsSupportEnabled: true
```

# Example with Environment Variables

This example uses environment variables to parameterize the VPC configuration.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsVpc
metadata:
  name: my-env-vpc
spec:
  awsCredentialId: ${AWS_CREDENTIAL_ID}
  vpcCidr: ${VPC_CIDR}
  availabilityZones:
    - ${AVAILABILITY_ZONE_1}
    - ${AVAILABILITY_ZONE_2}
  subnetsPerAvailabilityZone: ${SUBNETS_PER_AZ}
  subnetSize: ${SUBNET_SIZE}
  isNatGatewayEnabled: ${ENABLE_NAT_GATEWAY}
  isDnsHostnamesEnabled: ${ENABLE_DNS_HOSTNAMES}
  isDnsSupportEnabled: ${ENABLE_DNS_SUPPORT}
```

In this example, replace the placeholders like `${AWS_CREDENTIAL_ID}` with your actual environment variable names or values.

# Example with Environment Secrets

The below example assumes that the secrets are managed by Planton Cloud's [AWS Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.aws.awssecretsmanager) deployment module.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsVpc
metadata:
  name: my-secret-vpc
spec:
  awsCredentialId: my-aws-credential-id
  vpcCidr: 10.1.0.0/16
  availabilityZones:
    - us-east-1a
    - us-east-1b
  subnetsPerAvailabilityZone: 2
  subnetSize: 512
  isNatGatewayEnabled: true
  isDnsHostnamesEnabled: true
  isDnsSupportEnabled: true
  someSecretConfig: ${awssm-my-org-prod-aws-secrets.secret-key}
```

In this example:

- **someSecretConfig** is a placeholder for any configuration value you want to retrieve from AWS Secrets Manager.
- The value before the dot (`awssm-my-org-prod-aws-secrets`) is the ID of the AWS Secrets Manager resource on Planton Cloud.
- The value after the dot (`secret-key`) is the name of the secret within that resource.

# Example with All Available Fields

This comprehensive example demonstrates the full capabilities of the `AwsVpc` resource.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: AwsVpc
metadata:
  name: my-full-config-vpc
spec:
  awsCredentialId: my-aws-credential-id
  vpcCidr: 10.2.0.0/16
  availabilityZones:
    - us-east-1a
    - us-east-1b
    - us-east-1c
  subnetsPerAvailabilityZone: 3
  subnetSize: 256
  isNatGatewayEnabled: true
  isDnsHostnamesEnabled: true
  isDnsSupportEnabled: true
```

---

These examples illustrate various configurations of the `AwsVpc` API resource, demonstrating how to define VPCs with different features such as environment variables, environment secrets, and comprehensive settings.

Please ensure that you replace placeholder values like `my-aws-credential-id`, environment variable names, and secret references with your actual configuration details.
