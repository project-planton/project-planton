# Overview

The provided Pulumi module automates the provisioning of a Google Kubernetes Engine (GKE) cluster using Golang and
Pulumi, based on a unified API resource specification. It takes a `GkeCluster` API resource as input and orchestrates
the creation of Google Cloud resources, including projects, folders, VPC networks, subnets, and the GKE cluster itself.
The module supports configurations for shared VPC setups, cluster autoscaling, node pools with specific machine types,
and custom network settings. It also handles IAM roles and service accounts, ensuring secure and appropriate permissions
are set for the cluster operations.
