###############################################
# main.tf
###############################################

# ------------------------------------------------------------------------------
# 1) Certificate (only if TLS is enabled)
# ------------------------------------------------------------------------------
resource "kubernetes_manifest" "ingress_certificate" {
  count = local.is_tls_enabled ? 1 : 0

  manifest = {
    apiVersion = "cert-manager.io/v1"
    kind       = "Certificate"
    metadata = {
      name      = local.resource_id
      namespace = var.istio_ingress_namespace
      labels    = local.final_labels
    }
    spec = {
      dnsNames   = [local.endpoint_domain_name]
      secretName = local.ingress_cert_secret_name
      issuerRef = {
        kind = "ClusterIssuer"
        name = local.cert_cluster_issuer_name
      }
    }
  }
}

# ------------------------------------------------------------------------------
# 2) Gateway
# ------------------------------------------------------------------------------
resource "kubernetes_manifest" "gateway" {
  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1"
    kind       = "Gateway"
    metadata = {
      name      = local.resource_id
      namespace = var.istio_ingress_namespace
      labels    = local.final_labels
    }
    spec = {
      gatewayClassName = var.gateway_ingress_class_name
      addresses = [
        {
          type  = "Hostname"
          value = var.gateway_external_load_balancer_service_hostname
        }
      ]
      # The final array is set in locals.tf so that Terraform sees
      # one consistent type for both the "true" and "false" branches.
      listeners = local.gateway_listeners
    }
  }

  depends_on = [
    kubernetes_manifest.ingress_certificate
  ]
}

# ------------------------------------------------------------------------------
# 3) HTTPRoute for redirecting HTTP -> HTTPS (only if TLS is enabled)
# ------------------------------------------------------------------------------
resource "kubernetes_manifest" "http_external_redirect" {
  count = local.is_tls_enabled ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1"
    kind       = "HTTPRoute"
    metadata = {
      name      = "http-external-redirect"
      namespace = var.istio_ingress_namespace
      labels    = local.final_labels
    }
    spec = {
      hostnames = [local.endpoint_domain_name]
      parentRefs = [
        {
          name        = local.resource_id
          namespace   = var.istio_ingress_namespace
          sectionName = "http-external"
        }
      ]
      rules = [
        {
          filters = [
            {
              type            = "RequestRedirect"
              requestRedirect = {
                scheme     = "https"
                statusCode = 301
              }
            }
          ]
        }
      ]
    }
  }

  depends_on = [
    kubernetes_manifest.gateway
  ]
}

# ------------------------------------------------------------------------------
# 4) Primary HTTPRoute (always created)
# ------------------------------------------------------------------------------
resource "kubernetes_manifest" "primary_http_route" {
  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1"
    kind       = "HTTPRoute"
    metadata = {
      name      = local.resource_id
      namespace = var.istio_ingress_namespace
      labels    = local.final_labels
    }
    spec = {
      hostnames = [local.endpoint_domain_name]
      parentRefs = [
        {
          name        = local.resource_id
          namespace   = var.istio_ingress_namespace
          sectionName = local.route_section_name
        }
      ]
      rules = local.http_route_rules
    }
  }

  depends_on = [
    kubernetes_manifest.gateway
  ]
}
