# Only create resources if ingress is enabled and dns_domain is not empty
resource "kubernetes_manifest" "ingress_certificate" {
  count = local.ingress_is_enabled && local.ingress_dns_domain != "" ? 1 : 0

  manifest = yamldecode(
    yamlencode({
      apiVersion = "cert-manager.io/v1"
      kind       = "Certificate"
      metadata = {
        name      = local.resource_id
        namespace = local.istio_ingress_namespace
        labels    = local.final_labels
      }
      spec = {
        dnsNames   = local.ingress_hostnames
        secretName = local.ingress_cert_secret_name
        issuerRef = {
          kind = "ClusterIssuer"
          name = local.ingress_cert_cluster_issuer_name
        }
      }
    })
  )
}

resource "kubernetes_manifest" "jenkins_external_gateway" {
  count = local.ingress_is_enabled && local.ingress_dns_domain != "" ? 1 : 0

  manifest = yamldecode(
    yamlencode({
      apiVersion = "gateway.networking.k8s.io/v1beta1"
      kind       = "Gateway"
      metadata = {
        name      = "${local.resource_id}-external"
        namespace = local.istio_ingress_namespace
        labels    = local.final_labels
      }
      spec = {
        gatewayClassName = local.gateway_ingress_class_name
        addresses = [
          {
            type  = "Hostname"
            value = local.gateway_external_lb_hostname
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
    })
  )

  depends_on = [
    kubernetes_manifest.ingress_certificate
  ]
}

resource "kubernetes_manifest" "jenkins_http_external_redirect" {
  count = local.ingress_is_enabled && local.ingress_dns_domain != "" ? 1 : 0

  manifest = yamldecode(
    yamlencode({
      apiVersion = "gateway.networking.k8s.io/v1beta1"
      kind       = "HTTPRoute"
      metadata = {
        name      = "http-external-redirect"
        namespace = local.namespace
        labels    = local.final_labels
      }
      spec = {
        hostnames = [local.ingress_external_hostname]
        parentRefs = [
          {
            name        = "${local.resource_id}-external"
            namespace   = local.istio_ingress_namespace
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
    })
  )

  depends_on = [
    kubernetes_manifest.jenkins_external_gateway
  ]
}

resource "kubernetes_manifest" "jenkins_https_external" {
  count = local.ingress_is_enabled && local.ingress_dns_domain != "" ? 1 : 0

  manifest = yamldecode(
    yamlencode({
      apiVersion = "gateway.networking.k8s.io/v1beta1"
      kind       = "HTTPRoute"
      metadata = {
        name      = "https-external"
        namespace = local.namespace
        labels    = local.final_labels
      }
      spec = {
        hostnames = [local.ingress_external_hostname]
        parentRefs = [
          {
            name        = "${local.resource_id}-external"
            namespace   = local.istio_ingress_namespace
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
                name      = local.jenkins_kube_service_name
                namespace = local.namespace
                port      = 80
              }
            ]
          }
        ]
      }
    })
  )

  depends_on = [
    kubernetes_manifest.jenkins_external_gateway
  ]
}
