# Overview

The provided Pulumi module automates the provisioning of a Google Kubernetes Engine (GKE) cluster using Golang and
Pulumi, based on a unified API resource specification. It takes a `GkeCluster` API resource as input and orchestrates
the creation of Google Cloud resources, including projects, folders, VPC networks, subnets, and the GKE cluster itself.
The module supports configurations for shared VPC setups, cluster autoscaling, node pools with specific machine types,
and custom network settings. It also handles IAM roles and service accounts, ensuring secure and appropriate permissions
are set for the cluster operations.

Beyond the basic cluster setup, the module automates the installation of essential Kubernetes addons such as Istio
service mesh, Ingress Nginx controller, Cert Manager, External DNS, and various operators like Postgres, Kafka, Solr,
and Elastic. It manages workload identity bindings and integrates with Google Cloud services like Cloud DNS and Secrets
Manager. This comprehensive automation enables developers to deploy complex infrastructure and applications by simply
providing a YAML configuration, streamlining the deployment process in a multi-cloud environment and adhering to the
standardized API structures defined by the organization.
