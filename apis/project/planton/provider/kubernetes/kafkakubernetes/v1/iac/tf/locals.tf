locals {
  # Derive a stable resource ID
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "kafka_kubernetes"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if var.metadata.env.id is non-empty
  env_label = (
  var.metadata.env != null && try(var.metadata.env.id, "") != ""
  ) ? {
    "environment" = var.metadata.env.id
  } : {}

  # Merge base, org, and environment labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Use resource_id as the namespace name
  namespace = local.resource_id

  # Kafka broker container replicas (for convenience)
  broker_replicas = try(var.spec.broker_container.replicas, 1)

  # Zookeeper container replicas (though the Pulumi code currently uses broker_container.replicas for ZK, too)
  zookeeper_replicas = try(var.spec.zookeeper_container.replicas, 1)

  # Basic flags for ingress
  ingress_is_enabled = try(var.spec.ingress.is_enabled, false)
  ingress_dns_domain = try(var.spec.ingress.dns_domain, "")

  # Name and FQDN for the Kafka bootstrap service inside the cluster
  bootstrap_kube_service_name = "${local.resource_id}-kafka-bootstrap"
  bootstrap_kube_service_fqdn = "${local.bootstrap_kube_service_name}.${local.namespace}.svc"

  # External and internal bootstrap hostnames (null if ingress is disabled)
  ingress_external_bootstrap_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}-bootstrap.${local.ingress_dns_domain}" : null

  ingress_internal_bootstrap_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? "${local.resource_id}-bootstrap-internal.${local.ingress_dns_domain}" : null

  # Generate per-broker hostnames for external and internal listeners (empty lists if ingress is disabled)
  ingress_external_broker_hostnames = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? [
    for i in range(local.broker_replicas) :
    "${local.resource_id}-broker-b${i}.${local.ingress_dns_domain}"
  ] : []

  ingress_internal_broker_hostnames = (
  local.ingress_is_enabled && local.ingress_dns_domain != ""
  ) ? [
    for i in range(local.broker_replicas) :
    "${local.resource_id}-broker-b${i}-internal.${local.ingress_dns_domain}"
  ] : []

  # Aggregate list of hostnames used by Kafka's load-balanced ingress
  ingress_hostnames = concat(
    compact([local.ingress_external_bootstrap_hostname, local.ingress_internal_bootstrap_hostname]),
    local.ingress_external_broker_hostnames,
    local.ingress_internal_broker_hostnames
  )

  # Flag to indicate whether Schema Registry is enabled
  is_schema_registry_enabled = try(var.spec.schema_registry_container.is_enabled, false)

  # Schema Registry external and internal hostnames (null if ingress or registry is disabled)
  schema_registry_external_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != "" && local.is_schema_registry_enabled
  ) ? "${local.resource_id}-schema-registry.${local.ingress_dns_domain}" : null

  schema_registry_internal_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != "" && local.is_schema_registry_enabled
  ) ? "${local.resource_id}-schema-registry-internal.${local.ingress_dns_domain}" : null

  schema_registry_hostnames = concat(
    compact([local.schema_registry_external_hostname, local.schema_registry_internal_hostname])
  )

  # Flag to indicate whether we deploy a Kafka UI (Kowl)
  is_deploy_kafka_ui = try(var.spec.is_deploy_kafka_ui, true)

  # Kowl external hostname (null if ingress is disabled or if we're not deploying it)
  kowl_external_hostname = (
  local.ingress_is_enabled && local.ingress_dns_domain != "" && local.is_deploy_kafka_ui
  ) ? "${local.resource_id}-kowl.${local.ingress_dns_domain}" : null
}
