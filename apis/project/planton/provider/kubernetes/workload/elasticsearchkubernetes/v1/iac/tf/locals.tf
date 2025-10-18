###########################
# locals.tf
###########################

locals {
  # Derive a stable resource ID (fall back to name if ID is missing/empty).
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  # Base labels
  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "elasticsearch_kubernetes"
  }

  # Organization label only if non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if env and env.id are set
  env_label = (
  var.metadata.env != null && var.metadata.env != ""
  ) ? {
    "organization" = var.metadata.env
  } : {}
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # Namespace is the resource_id
  namespace = local.resource_id

  # Service names and endpoints
  elasticsearch_kube_service_name = "${var.metadata.name}-es-http"
  elasticsearch_kube_service_fqdn = "${local.elasticsearch_kube_service_name}.${local.namespace}.svc.cluster.local"
  elasticsearch_kube_port_forward_command = "kubectl port-forward -n ${local.namespace} service/${local.elasticsearch_kube_service_name} 9200:9200"

  kibana_kube_service_name = "${var.metadata.name}-kb-http"
  kibana_kube_service_fqdn = "${local.kibana_kube_service_name}.${local.namespace}.svc.cluster.local"
  kibana_kube_port_forward_command = "kubectl port-forward -n ${local.namespace} service/${local.kibana_kube_service_name} 5601:5601"

  # Elasticsearch ingress
  elasticsearch_ingress_is_enabled = try(var.spec.elasticsearch.ingress.enabled, false)
  elasticsearch_ingress_external_hostname = try(var.spec.elasticsearch.ingress.hostname, null)

  # Kibana ingress
  kibana_is_enabled = try(var.spec.kibana.enabled, false)
  kibana_ingress_is_enabled = local.kibana_is_enabled && try(var.spec.kibana.ingress.enabled, false)
  kibana_ingress_external_hostname = local.kibana_is_enabled ? try(var.spec.kibana.ingress.hostname, null) : null

  # Combine hostnames for certificate
  ingress_hostnames = compact([
    local.elasticsearch_ingress_external_hostname,
    local.kibana_ingress_external_hostname,
  ])

  # Certificate issuer: extract domain from first hostname
  ingress_cert_cluster_issuer_name = length(local.ingress_hostnames) > 0 ? (
    join(".", slice(split(".", local.ingress_hostnames[0]), 1, length(split(".", local.ingress_hostnames[0]))))
  ) : ""

  ingress_cert_secret_name = local.resource_id

  # These match your Pulumi vars. Adjust as needed or make them variables.
  istio_ingress_namespace      = "istio-ingress"
  gateway_ingress_class_name   = "istio"
  gateway_external_lb_hostname = "ingress-external.istio-ingress.svc.cluster.local"

  # The ports for Elasticsearch and Kibana (from your `vars` struct)
  elasticsearch_port = 9200
  kibana_port        = 5601

  elasticsearch_version = "8.15.0"
}
