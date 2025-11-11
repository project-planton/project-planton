###############################################
# locals.tf
###############################################

locals {
  # --------------------------------------------------------------------------
  # Basic resource/label logic
  # --------------------------------------------------------------------------
  resource_id = (
    var.metadata.id != null && var.metadata.id != ""
    ? var.metadata.id
    : var.metadata.name
  )

  base_labels = {
    "resource"      = "true"
    "resource_id"   = local.resource_id
    "resource_kind" = "kubernetes_http_endpoint"
  }

  org_label = (
  var.metadata.org != null && var.metadata.org != ""
  ) ? {
    "organization" = var.metadata.org
  } : {}

  env_label = (
  var.metadata.env != null && var.metadata.env != ""
  ) ? {
    "organization" = var.metadata.env
  } : {}
  final_labels = merge(local.base_labels, local.org_label, local.env_label)

  endpoint_domain_name       = var.metadata.name
  is_tls_enabled             = try(var.spec.is_tls_enabled, false)
  is_grpc_web_compatible     = try(var.spec.is_grpc_web_compatible, false)
  cert_cluster_issuer_name   = try(var.spec.cert_cluster_issuer_name, "")
  ingress_cert_secret_name   = "cert-${local.endpoint_domain_name}"
  routing_rules              = try(var.spec.routing_rules, [])

  # For the primary HTTPRoute parentRef section:
  route_section_name = local.is_tls_enabled ? "https-external" : "http-external"

  # --------------------------------------------------------------------------
  # Gateway Listeners
  # --------------------------------------------------------------------------
  # 1) Base HTTP listener (for protocol=HTTP). Must NOT have any tls block.
  http_listener = {
    name        = "http-external"
    hostname    = local.endpoint_domain_name
    port        = 80
    protocol    = "HTTP"
    allowedRoutes = {
      namespaces = {
        from = "All"
      }
    }
    # Omit tls entirely for HTTP
  }

  # 2) Optional HTTPS listener (for protocol=HTTPS) if TLS is enabled
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

    # TLS is only valid when protocol=HTTPS
    tls = {
      mode = "Terminate"
      # Note: "certificateRefs" is correct for Gateway API v1,
      # but older CRDs might require "certificateRef" instead.
      certificateRefs = [
        {
          name = local.ingress_cert_secret_name
        }
      ]
    }
  }

  # 3) Create a list with the base HTTP listener,
  # and append the HTTPS listener only if TLS is enabled.
  gateway_listeners = concat(
    [local.http_listener],
      local.is_tls_enabled ? [local.https_listener] : []
  )

  # --------------------------------------------------------------------------
  # HTTPRoute rules array
  # --------------------------------------------------------------------------
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
