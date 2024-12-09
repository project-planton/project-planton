# GCP Static Website Pulumi Module

## Overview

The GCP Static Website Pulumi module is designed to simplify the process of deploying static websites on Google Cloud Platform (GCP) by using Google Cloud Storage as the hosting infrastructure. This module integrates with the unified API framework developed by Planton Cloud, which models every API resource in a Kubernetes-like structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status` fields. The `GcpStaticWebsite` resource defines the necessary specifications for creating the infrastructure required to serve static websites using GCP Storage.

By utilizing this Pulumi module, developers can automate the provisioning of a GCP Storage bucket and configure it to serve static content as a public website. The module interacts with the GCP credentials, project ID, and other specifications provided in the resource, and ensures that the deployment process is streamlined. The outputs from the deployment, such as the bucket ID, are captured in the resource's status, allowing users to track and manage their infrastructure directly from the `GcpStaticWebsite` resource.

## Key Features

- **Automated Deployment**: The module automates the entire process of setting up the infrastructure needed to host static websites on GCP. It provisions a GCP Storage bucket configured for static website hosting based on the input YAML specification.
  
- **API-Driven Integration**: This module integrates seamlessly with the Planton Cloud API, following the standardized resource structure of `apiVersion`, `kind`, `metadata`, `spec`, and `status`. This enables easy integration into CI/CD pipelines or automated deployment workflows.

- **Credentials and Security**: The module requires GCP credentials to authenticate and configure the necessary services. It supports custom credential specifications, ensuring secure and flexible deployments.

- **Custom Configuration**: Users can specify the GCP project and other configuration options like the credential ID and project ID, making the module adaptable to various project needs.

- **Status Tracking**: The module automatically captures key output values, such as the bucket ID, in the status field, which helps in managing and referencing the created infrastructure within the same resource context.

## Usage

The GCP Static Website Pulumi module can be triggered via the CLI using the `planton pulumi up` command. Developers are required to provide an input YAML file that defines the `GcpStaticWebsite` resource with its required fields. If no module repository is specified, the default module associated with this resource will be used for the deployment.

Refer to the example section for usage instructions.

## Notes

**Important:** The current `GcpStaticWebsite` API resource spec is fully implemented and supports the provisioning of GCP storage buckets for static website hosting. Any missing fields or custom configurations can be added in future iterations of the resource definition.

