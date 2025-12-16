# Create a certificate using cert-manager (requires cert-manager CRDs).
# Only create it if ingress is enabled.
resource "kubernetes_manifest" "ingress_certificate" {
  count = local.ingress_is_enabled ? 1 : 0

  manifest = {
    apiVersion = "cert-manager.io/v1"
    kind       = "Certificate"
    metadata = {
      name      = local.resource_id
      namespace = "istio-ingress"
      labels    = local.final_labels
    }
    spec = {
      dnsNames   = [local.ingress_external_hostname]
      secretName = local.ingress_cert_secret_name
      issuerRef = {
        kind = "ClusterIssuer"
        name = local.ingress_cert_cluster_issuer_name
      }
    }
  }

  # Implicit dependency through local.namespace_name reference
}

# Create a Gateway for external ingress (requires Gateway API CRDs).
# Only create it if ingress is enabled.
resource "kubernetes_manifest" "gateway" {
  count = local.ingress_is_enabled ? 1 : 0

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

# Create an HTTPRoute that redirects HTTP to HTTPS for the external hostname.
# Only create it if ingress is enabled.
resource "kubernetes_manifest" "http_route_external_redirect" {
  count = local.ingress_is_enabled ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "HTTPRoute"
    metadata = {
      name      = "http-external-redirect"
      namespace = local.namespace_name
      labels    = local.final_labels
    }
    spec = {
      hostnames = [
        local.ingress_external_hostname
      ]
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
    kubernetes_manifest.gateway
  ]
}

# Create an HTTPS route for the external hostname to route traffic to the OpenFGA service.
# Only create it if ingress is enabled.
resource "kubernetes_manifest" "http_route_https_external" {
  count = local.ingress_is_enabled ? 1 : 0

  manifest = {
    apiVersion = "gateway.networking.k8s.io/v1beta1"
    kind       = "HTTPRoute"
    metadata = {
      name      = "https-external"
      namespace = local.namespace_name
      labels    = local.final_labels
    }
    spec = {
      hostnames = [
        local.ingress_external_hostname
      ]
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
              namespace = local.namespace_name
              port      = 8080
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
