##############################################
# locals.tf
#
# Includes logic for:
#  - Deriving resource_id, labels, namespace
#  - Determining ingress hostnames
#  - Creating image_pull_secret_data from
#    docker_credential (if provider == "gcp_artifact_registry").
##############################################

locals {
  # Derive a stable resource ID (prefer `metadata.id`, fallback to `metadata.name`)
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "microservice_kubernetes"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? { "organization" = var.metadata.org } : {}

  # Environment label only if var.metadata.env.id is non-empty
  env_label = (
  var.metadata.env != null &&
  try(var.metadata.env.id, "") != ""
  ) ? { "environment" = var.metadata.env.id } : {}

  # Merge base, org, and environment labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Use resource_id as the namespace name
  namespace = local.resource_id

  # The microservice version is used as the Service name (per the code logic).
  kube_service_name = var.spec.version

  # Internal DNS name for the service
  kube_service_fqdn = "${local.kube_service_name}.${local.namespace}.svc.cluster.local"

  # Handy port-forward command
  kube_port_forward_command = "kubectl port-forward -n ${local.namespace} service/${local.kube_service_name} 8080:8080"

  # Safely handle optional ingress values
  ingress_is_enabled = try(var.spec.ingress.is_enabled, false)
  ingress_dns_domain = try(var.spec.ingress.dns_domain, "")

  # External hostname (null if not applicable)
  ingress_external_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}.${local.ingress_dns_domain}" : null

  # Internal hostname (null if not applicable)
  ingress_internal_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}-internal.${local.ingress_dns_domain}" : null

  # For certificate creation
  ingress_cert_cluster_issuer_name = local.ingress_dns_domain != "" ? local.ingress_dns_domain : null
  ingress_cert_secret_name         = local.resource_id

  #############################################################
  # Docker Credential → image_pull_secret_data logic
  #
  # String-based provider check: "gcp_artifact_registry"
  #############################################################

  # True if docker_credential is non-null and
  # provider == "gcp_artifact_registry"
  docker_credential_is_gcp = (
  var.docker_credential != null
  && var.docker_credential.provider == "gcp_artifact_registry"
  )

  # Build the GCP registry hostname if GCP
  docker_repo_hostname = (
    local.docker_credential_is_gcp
    ? "${var.docker_credential.gcp_artifact_registry.gcp_region}-docker.pkg.dev"
    : null
  )

  # Decode the service account key from base64
  docker_sa_key_decoded = (
    local.docker_credential_is_gcp
    ? base64decode(var.docker_credential.gcp_artifact_registry.service_account_key_base64)
    : null
  )

  # Re-encode as _json_key: to produce Docker "auth"
  docker_sa_key_encoded = (
    local.docker_credential_is_gcp
    ? base64encode("_json_key:${local.docker_sa_key_decoded}")
    : null
  )

  # Construct the final .dockerconfigjson data (null if not GCP)
  image_pull_secret_data = (
    local.docker_credential_is_gcp
    ? {
    ".dockerconfigjson" = jsonencode({
      auths = {
        (local.docker_repo_hostname) = {
          username = "_json_key"
          auth     = local.docker_sa_key_encoded
        }
      }
    })
  }
    : null
  )
}
