##############################################
# ingress.tf
#
# Creates cert-manager Certificate and Gateway
# API resources if ingress is enabled.
##############################################

# Create a certificate using cert-manager
resource "kubernetes_manifest" "ingress_certificate" {
  # Only create if ingress is enabled
  count = local.ingress_is_enabled && local.ingress_dns_domain != "" ? 1 : 0

  manifest = {
    apiVersion = "cert-manager.io/v1"
    kind       = "Certificate"
    metadata = {
      name      = local.resource_id
      namespace = "istio-ingress"
      labels    = local.final_labels
    }
    spec = {
      dnsNames = [
        for hostname in [
          local.ingress_external_hostname,
          local.ingress_internal_hostname
        ]
        : hostname if hostname != null
      ]
      secretName = local.ingress_cert_secret_name
      issuerRef = {
        kind = "ClusterIssuer"
        name = local.ingress_cert_cluster_issuer_name
      }
    }
  }

  depends_on = [
    kubernetes_namespace.this
  ]
}

# Create external Gateway
resource "kubernetes_manifest" "gateway_external" {
  count = local.ingress_is_enabled && local.ingress_dns_domain != "" ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "Gateway"
    metadata = {
      name      = "${local.resource_id}-external"
      namespace = "istio-ingress"
      labels    = local.final_labels
    }
    spec = {
      gatewayClassName = "istio"
      addresses = [
        {
          type  = "Hostname"
          value = "ingress-external.istio-ingress.svc.cluster.local"
        }
      ]
      listeners = [
        {
          name     = "https-external"
          hostname = local.ingress_external_hostname
          port     = 443
          protocol = "HTTPS"
          tls = {
            mode = "Terminate"
            certificateRefs = [
              {
                name = local.ingress_cert_secret_name
              }
            ]
          }
          allowedRoutes = {
            namespaces = {
              from = "All"
            }
          }
        },
        {
          name     = "http-external"
          hostname = local.ingress_external_hostname
          port     = 80
          protocol = "HTTP"
          allowedRoutes = {
            namespaces = {
              from = "All"
            }
          }
        }
      ]
    }
  }

  depends_on = [
    kubernetes_manifest.ingress_certificate
  ]
}

# Create internal Gateway
resource "kubernetes_manifest" "gateway_internal" {
  count = local.ingress_is_enabled && local.ingress_dns_domain != "" ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "Gateway"
    metadata = {
      name      = "${local.resource_id}-internal"
      namespace = "istio-ingress"
      labels    = local.final_labels
    }
    spec = {
      gatewayClassName = "istio"
      addresses = [
        {
          type  = "Hostname"
          value = "ingress-internal.istio-ingress.svc.cluster.local"
        }
      ]
      listeners = [
        {
          name     = "https-internal"
          hostname = local.ingress_internal_hostname
          port     = 443
          protocol = "HTTPS"
          tls = {
            mode = "Terminate"
            certificateRefs = [
              {
                name = local.ingress_cert_secret_name
              }
            ]
          }
          allowedRoutes = {
            namespaces = {
              from = "All"
            }
          }
        },
        {
          name     = "http-internal"
          hostname = local.ingress_internal_hostname
          port     = 80
          protocol = "HTTP"
          allowedRoutes = {
            namespaces = {
              from = "All"
            }
          }
        }
      ]
    }
  }

  depends_on = [
    kubernetes_manifest.ingress_certificate
  ]
}

# Build an easy reference to the service port
# If the user has multiple ports in spec.container.app.ports,
# we pick the one where `is_ingress_port = true`; otherwise fall
# back to port 80. This snippet is purely for demonstration—some
# specialized logic might be needed if multiple ingress ports exist.
locals {
  ingress_port = (
    length([
      for p in try(var.spec.container.app.ports, []) : p.service_port
      if try(p.is_ingress_port, false)
    ]) > 0
    ? [
    for p in var.spec.container.app.ports : p.service_port
    if try(p.is_ingress_port, false)
  ][
  0
  ]
    : 80
  )
}

# -------------
# External Host
# -------------
# 1) HTTP -> HTTPS redirect
resource "kubernetes_manifest" "http_external_redirect" {
  count = local.ingress_is_enabled && local.ingress_external_hostname != null ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "HTTPRoute"
    metadata = {
      name      = "http-external-redirect"
      namespace = kubernetes_namespace.this.metadata[0].name
      labels    = local.final_labels
    }
    spec = {
      hostnames = [local.ingress_external_hostname]
      parentRefs = [
        {
          name        = "${local.resource_id}-external"
          namespace   = "istio-ingress"
          sectionName = "http-external"
        }
      ]
      rules = [
        {
          filters = [
            {
              type = "RequestRedirect"
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
    kubernetes_manifest.gateway_external
  ]
}

# 2) HTTPS route
resource "kubernetes_manifest" "https_external" {
  count = local.ingress_is_enabled && local.ingress_external_hostname != null ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "HTTPRoute"
    metadata = {
      name      = "https-external"
      namespace = kubernetes_namespace.this.metadata[0].name
      labels    = local.final_labels
    }
    spec = {
      hostnames = [local.ingress_external_hostname]
      parentRefs = [
        {
          name        = "${local.resource_id}-external"
          namespace   = "istio-ingress"
          sectionName = "https-external"
        }
      ]
      rules = [
        {
          matches = [
            {
              path = {
                type  = "PathPrefix"
                value = "/"
              }
            }
          ]
          backendRefs = [
            {
              name      = local.kube_service_name
              namespace = kubernetes_namespace.this.metadata[0].name
              port      = local.ingress_port
            }
          ]
        }
      ]
    }
  }

  depends_on = [
    kubernetes_manifest.gateway_external
  ]
}

# -------------
# Internal Host
# -------------
# 3) HTTP -> HTTPS redirect
resource "kubernetes_manifest" "http_internal_redirect" {
  count = local.ingress_is_enabled && local.ingress_internal_hostname != null ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "HTTPRoute"
    metadata = {
      name      = "http-internal-redirect"
      namespace = kubernetes_namespace.this.metadata[0].name
      labels    = local.final_labels
    }
    spec = {
      hostnames = [local.ingress_internal_hostname]
      parentRefs = [
        {
          name        = "${local.resource_id}-internal"
          namespace   = "istio-ingress"
          sectionName = "http-internal"
        }
      ]
      rules = [
        {
          filters = [
            {
              type = "RequestRedirect"
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
    kubernetes_manifest.gateway_internal
  ]
}

# 4) HTTPS route
resource "kubernetes_manifest" "https_internal" {
  count = local.ingress_is_enabled && local.ingress_internal_hostname != null ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "HTTPRoute"
    metadata = {
      name      = "https-internal"
      namespace = kubernetes_namespace.this.metadata[0].name
      labels    = local.final_labels
    }
    spec = {
      hostnames = [local.ingress_internal_hostname]
      parentRefs = [
        {
          name        = "${local.resource_id}-internal"
          namespace   = "istio-ingress"
          sectionName = "https-internal"
        }
      ]
      rules = [
        {
          matches = [
            {
              path = {
                type  = "PathPrefix"
                value = "/"
              }
            }
          ]
          backendRefs = [
            {
              name      = local.kube_service_name
              namespace = kubernetes_namespace.this.metadata[0].name
              port      = local.ingress_port
            }
          ]
        }
      ]
    }
  }

  depends_on = [
    kubernetes_manifest.gateway_internal
  ]
}
