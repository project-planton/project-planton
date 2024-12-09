# Overview

The provided Pulumi module is designed to create AWS resources by setting up an AWS provider using credentials specified in the `AwsCloudFrontStackInput`. This module adheres to a standardized API structure that mimics Kubernetes resource modeling, utilizing fields like `apiVersion`, `kind`, `metadata`, `spec`, and `status`. By supplying AWS credentials through the `aws_credential_id` in the specification, the module initializes an AWS provider that can be used to manage AWS resources.

In the code, the `Resources` function creates a new AWS provider using the access key, secret key, and region from the input credentials. This setup allows developers to deploy AWS CloudFront distributions and other resources seamlessly. The module emphasizes simplicity and consistency, enabling the deployment of complex infrastructure through straightforward YAML configurations.
