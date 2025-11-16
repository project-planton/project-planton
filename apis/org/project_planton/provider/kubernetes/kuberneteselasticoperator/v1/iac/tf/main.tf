##############################################
# main.tf
#
# Main orchestration file for deploying the
# Elastic Cloud on Kubernetes (ECK) operator.
#
# This module installs the ECK operator using Helm,
# which provides automated lifecycle management for
# Elasticsearch, Kibana, APM Server, Enterprise Search,
# Beats, Elastic Agent, and Logstash on Kubernetes.
#
# Resources Created:
#  1. Kubernetes Namespace (elastic-system)
#  2. Helm Release (eck-operator chart)
#
# The ECK operator extends the Kubernetes API with
# Custom Resource Definitions (CRDs) for Elastic Stack
# components and provides automated operations including:
#  - Certificate management and rotation
#  - Rolling upgrades
#  - Scaling operations
#  - Cross-cluster replication
#
# For more information see:
#  - examples.md for usage examples
#  - ../README.md for component documentation
#  - ../../docs/README.md for deployment patterns
##############################################

##############################################
# 1. Create Namespace
#
# ECK operator runs in the elastic-system namespace
# by default. This provides isolation and makes it
# easy to manage the operator lifecycle.
##############################################
resource "kubernetes_namespace" "elastic_system" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}

##############################################
# 2. Deploy ECK Operator via Helm
#
# Installs the Elastic Cloud on Kubernetes operator
# from the official Elastic Helm repository.
#
# The operator will:
#  - Install CRDs for Elastic Stack components
#  - Run as a controller that watches for custom resources
#  - Automatically manage Elastic Stack deployments
#  - Handle certificate management and rotation
#  - Orchestrate rolling upgrades
##############################################
resource "helm_release" "eck_operator" {
  name       = local.helm_chart_name
  repository = local.helm_chart_repo
  chart      = local.helm_chart_name
  version    = local.helm_chart_version
  namespace  = kubernetes_namespace.elastic_system.metadata[0].name

  # Disable namespace creation since we create it above
  create_namespace = false

  # Helm release options
  atomic          = false # Allow partial deployment for debugging
  cleanup_on_fail = true  # Clean up resources if deployment fails
  wait            = true  # Wait for resources to be ready
  wait_for_jobs   = true  # Wait for any jobs to complete
  timeout         = 180   # 3 minutes timeout

  # Configure ECK to inherit Planton labels
  values = [yamlencode({
    # Propagate Planton labels to all ECK-managed resources
    configKubernetes = {
      inherited_labels = local.inherited_labels
    }

    # Resource limits and requests for the ECK operator pod
    resources = {
      limits = {
        cpu    = try(var.spec.container.resources.limits.cpu, "1000m")
        memory = try(var.spec.container.resources.limits.memory, "1Gi")
      }
      requests = {
        cpu    = try(var.spec.container.resources.requests.cpu, "50m")
        memory = try(var.spec.container.resources.requests.memory, "100Mi")
      }
    }
  })]

  # Ignore changes to these fields to prevent drift
  lifecycle {
    ignore_changes = [
      metadata,
    ]
  }

  depends_on = [
    kubernetes_namespace.elastic_system
  ]
}

