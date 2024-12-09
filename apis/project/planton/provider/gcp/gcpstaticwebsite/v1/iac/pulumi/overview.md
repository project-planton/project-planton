# GCP Static Website Pulumi Module

## Overview

The GCP Static Website Pulumi module is designed to simplify the process of deploying static websites on Google Cloud Platform (GCP) by using Google Cloud Storage as the hosting infrastructure. This module integrates with the unified API framework developed by Planton Cloud, which models every API resource in a Kubernetes-like structure with `apiVersion`, `kind`, `metadata`, `spec`, and `status` fields. The `GcpStaticWebsite` resource defines the necessary specifications for creating the infrastructure required to serve static websites using GCP Storage.

By utilizing this Pulumi module, developers can automate the provisioning of a GCP Storage bucket and configure it to serve static content as a public website. The module interacts with the GCP credentials, project ID, and other specifications provided in the resource, and ensures that the deployment process is streamlined. The outputs from the deployment, such as the bucket ID, are captured in the resource's status, allowing users to track and manage their infrastructure directly from the `GcpStaticWebsite` resource.
