##############################################
# main.tf
#
# Main orchestration file for KubernetesCronJob
# deployment using Terraform.
#
# This module creates the following resources:
#  1. Kubernetes Namespace (defined here)
#  2. Service Account for the CronJob (cron_job.tf)
#  3. Image Pull Secret (if docker_config_json provided) (cron_job.tf)
#  4. Environment Secrets (if spec.env.secrets provided) (secret.tf)
#  5. CronJob Resource (cron_job.tf)
#
# The module follows best practices:
#  - Uses namespace isolation for each CronJob
#  - Creates dedicated service account for RBAC
#  - Supports private container registries
#  - Implements proper resource management
#  - Provides comprehensive labeling
#
# For examples and usage patterns, see:
#  - examples.md for Terraform examples
#  - ../README.md for component documentation
#  - ../../docs/README.md for research and best practices
##############################################

##############################################
# 1. Create Namespace
#
# Each CronJob gets its own namespace for:
#  - Resource isolation
#  - Security boundary
#  - Easier cleanup and management
##############################################
resource "kubernetes_namespace" "this" {
  metadata {
    name   = local.namespace
    labels = local.final_labels
  }
}

##############################################
# Additional Resources
#
# Other resources are defined in dedicated files:
#  - cron_job.tf: ServiceAccount, ImagePullSecret, and CronJob
#  - secret.tf: Environment secrets (if configured)
#  - outputs.tf: Module outputs
#  - locals.tf: Local variables and computed values
#  - variables.tf: Input variable definitions
##############################################
