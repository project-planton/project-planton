# GcpCloudSql Pulumi Module

## Overview

The **GcpCloudSql** Pulumi module is part of the Planton Cloud ecosystem, which provides a unified API interface for managing multi-cloud infrastructure. This module enables developers to create and manage Google Cloud SQL instances by using a Kubernetes-style API resource model. The module automates the provisioning of Cloud SQL databases and other related resources in Google Cloud, based on the specifications defined in a YAML configuration file. It simplifies complex infrastructure management, offering an abstraction layer that allows developers to focus on application logic rather than infrastructure.

The key benefit of this module is its seamless integration with the Planton CLI and the standardized API resource format used across the entire platform. It allows you to provide configurations in a declarative manner, automatically handles resource creation on Google Cloud, and captures all outputs, such as instance details, in the `status.stackOutputs`. This reduces manual intervention and the complexity involved in managing cloud resources. By using this module, you can easily deploy, configure, and manage Cloud SQL instances in a consistent and predictable way across your environments.

## Features

- **Kubernetes-style API Resource Modeling**: Follows a Kubernetes-like structure, making it familiar for developers used to working with Kubernetes resources (with `apiVersion`, `kind`, `metadata`, `spec`, and `status` fields).
- **Multi-Cloud Infrastructure Management**: This module is part of the broader Planton Cloud API ecosystem, which enables multi-cloud management with a unified API interface.
- **Automated Google Cloud SQL Provisioning**: Takes an API resource as input and provisions Cloud SQL databases in Google Cloud automatically, based on the defined specification.
- **Output Capture**: Captures important resource outputs, such as instance names, connection strings, and access details, in the `status.stackOutputs` for later use or reference.
- **Planton CLI Integration**: Fully integrated with the Planton CLI, allowing for easy deployment and management of infrastructure resources via the `planton pulumi up` command.
- **Google Cloud Provider Integration**: This module directly interacts with Google Cloud using GCP credentials, ensuring secure and compliant management of infrastructure.
- **Declarative Infrastructure**: Define your infrastructure requirements declaratively in a YAML format, enabling reproducible and version-controlled infrastructure.

## GcpCloudSql API Resource

The **GcpCloudSql** resource is modeled after Kubernetes API resources, making it straightforward for developers familiar with Kubernetes to adopt. The resource includes:

- **Metadata**: Metadata section for resource identification (e.g., `name`, `namespace`).
- **Spec**: Contains details about the configuration and parameters required for provisioning Google Cloud SQL, including project ID, credentials, and other necessary configurations.
- **Status**: Outputs the created resources and their details, such as connection information and instance identifiers.

## Module Structure

This Pulumi module is written in Golang and interacts directly with Google Cloud to create the required SQL infrastructure. It handles various tasks such as:

- Setting up the Google Cloud provider using the provided GCP credentials.
- Provisioning Cloud SQL instances.
- Capturing and exporting resource outputs to the status for downstream use.

The module is designed to abstract the underlying complexity of managing Google Cloud infrastructure, providing a simple interface for declaratively defining and managing Cloud SQL resources.

## Usage

To use the module, define your Cloud SQL resource in a YAML file, and run the `planton pulumi up` command to deploy the resource. This will create the necessary Cloud SQL infrastructure in Google Cloud based on the defined specification.

**Refer to the example section for usage instructions.**
