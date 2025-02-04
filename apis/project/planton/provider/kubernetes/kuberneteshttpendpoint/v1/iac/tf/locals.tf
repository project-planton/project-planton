###############################################
# locals.tf
###############################################

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
    "resource_kind" = "kubernetes_http_endpoint"
  }

  # Organization label only if var.metadata.org is non-empty
  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  # Environment label only if var.metadata.env.id is non-empty
  env_label = (
  var.metadata.env != null &&
  try(var.metadata.env.id, "") != ""
  ) ? {
    "environment" = var.metadata.env.id
  } : {}

  # Merge base, org, and environment labels
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  # The endpoint domain name (used for Gateway, HTTPRoute, etc.)
  endpoint_domain_name = var.metadata.name

  # TLS and gRPC-Web flags
  is_tls_enabled         = try(var.spec.is_tls_enabled, false)
  is_grpc_web_compatible = try(var.spec.is_grpc_web_compatible, false)

  # Name of the ClusterIssuer to use for TLS, if needed
  cert_cluster_issuer_name = try(var.spec.cert_cluster_issuer_name, "")

  # Certificate secret name (only relevant if TLS is enabled)
  ingress_cert_secret_name = "cert-${local.endpoint_domain_name}"

  # The array of routing rules from the proto-based input
  routing_rules = try(var.spec.routing_rules, [])

  # For the primary HTTPRoute, if TLS is enabled, we attach to "https-external";
  # otherwise "http-external".
  route_section_name = local.is_tls_enabled ? "https-external" : "http-external"

  # ----------------------------------------------------------------------------
  # GATEWAY LISTENERS
  # ----------------------------------------------------------------------------
  # 1) Base HTTP listener (always present)
  base_http_listener = {
    name        = "http-external"
    hostname    = local.endpoint_domain_name
    port        = 80
    protocol    = "HTTP"
    allowedRoutes = {
      namespaces = {
        from = "All"
      }
    }
    # For consistent object shape, define `tls` as an empty map here
    tls = {}
  }

  # 2) Optional HTTPS listener (only if TLS is enabled)
  #    Same object shape: name, hostname, port, protocol, allowedRoutes, tls
  https_listener = {
    name        = "https-external"
    hostname    = local.endpoint_domain_name
    port        = 443
    protocol    = "HTTPS"
    allowedRoutes = {
      namespaces = {
        from = "All"
      }
    }
    tls = {
      mode = "Terminate"
      certificateRefs = [
        {
          name = local.ingress_cert_secret_name
        }
      ]
    }
  }

  # 3) Conditionally include the HTTPS listener
  maybe_https_listener = local.is_tls_enabled ? [local.https_listener] : []

  # 4) Final array of listeners for the Gateway
  #    We always have the base HTTP listener, plus 0 or 1 HTTPS listener
  gateway_listeners = concat(
    [local.base_http_listener],
    local.maybe_https_listener
  )

  # ----------------------------------------------------------------------------
  # HTTP ROUTE RULES
  # ----------------------------------------------------------------------------
  # Convert each item in routing_rules into a Gateway-API style "rules" entry
  http_route_rules = [
    for rule in local.routing_rules : {
      matches = [
        {
          path = {
            type  = "PathPrefix"
            value = rule.url_path_prefix
          }
        }
      ]
      backendRefs = [
        {
          name      = rule.backend_service.name
          namespace = rule.backend_service.namespace
          port      = rule.backend_service.port
        }
      ]
    }
  ]
}
