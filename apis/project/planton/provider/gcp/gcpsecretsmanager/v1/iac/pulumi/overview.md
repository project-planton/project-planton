# GCP Secrets Manager Pulumi Module

## Overview

The GCP Secrets Manager Pulumi module enables developers to declaratively manage secrets in Google Cloud's Secret Manager using Kubernetes-like API resource definitions. This module is designed to integrate seamlessly with Planton Cloud's unified API resources, where developers define infrastructure through YAML specifications. The module takes an API resource as input and provisions Google Cloud secrets based on the defined specifications. Key outputs, such as secret IDs, are captured in `status.stackOutputs`, making it easy to reference and use these secrets across different cloud services or components.

This module allows you to automate the creation of secrets in Google Cloud, initializing each secret with a placeholder version that can be updated with real values post-deployment. The developer can run the module using the `planton pulumi up --stack-input <api-resource.yaml>` command, which automatically sets up the GCP provider using the credentials provided in the API resource, then proceeds to create secrets in the defined Google Cloud project.

