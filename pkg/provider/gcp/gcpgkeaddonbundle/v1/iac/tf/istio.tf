###############################################################################
# Istio Installation
#
# 1. Create the "istio-system" namespace, labeled with our final_kubernetes_labels.
# 2. Deploy the Istio "base" Helm chart in the "istio-system" namespace.
# 3. Deploy the Istio "istiod" Helm chart in the "istio-system" namespace.
# 4. Create the "istio-ingress" namespace, labeled with our final_kubernetes_labels.
# 5. Deploy the Istio "gateway" Helm chart in the "istio-ingress" namespace.
# 6. Add the gRPC-Web EnvoyFilter for the Istio gateway.
# 7. Create internal and external load balancer IP addresses, then corresponding
#    Kubernetes Services in the "istio-ingress" namespace, referencing those IPs.
###############################################################################

###############################################
# 1. istio-system Namespace
###############################################
resource "kubernetes_namespace_v1" "istio_system_namespace" {
  count = var.spec.istio.enabled ? 1 : 0

  metadata {
    name   = "istio-system"
    labels = local.final_kubernetes_labels
  }
}

###############################################
# 2. Istio Base Helm Chart
###############################################
resource "helm_release" "istio_base" {
  count            = var.spec.istio.enabled ? 1 : 0
  name             = "base"
  repository       = "https://istio-release.storage.googleapis.com/charts"
  chart            = "base"
  version          = "1.22.3"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.istio_system_namespace[count.index].metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait             = true
}

###############################################
# 3. Istiod Helm Chart
###############################################
resource "helm_release" "istiod" {
  count            = var.spec.istio.enabled ? 1 : 0
  name             = "istiod"
  repository       = "https://istio-release.storage.googleapis.com/charts"
  chart            = "istiod"
  version          = "1.22.3"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.istio_system_namespace[count.index].metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait = true

  # Mesh config values to match the original logic
  values = [
    yamlencode({
      meshConfig = {
        ingressClass          = "istio"
        ingressControllerMode = "STRICT"
        ingressService        = "ingress-external"
        ingressSelector       = "ingress"
      }
    })
  ]

  depends_on = [
    helm_release.istio_base
  ]
}

###############################################
# 4. istio-ingress Namespace
###############################################
resource "kubernetes_namespace_v1" "istio_gateway_namespace" {
  count = var.spec.istio.enabled ? 1 : 0

  metadata {
    name   = "istio-ingress"
    labels = local.final_kubernetes_labels
  }
}

###############################################
# 5. Istio Gateway Helm Chart
###############################################
resource "helm_release" "istio_gateway" {
  count            = var.spec.istio.enabled ? 1 : 0
  name             = "gateway"
  repository       = "https://istio-release.storage.googleapis.com/charts"
  chart            = "gateway"
  version          = "1.22.3"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.istio_gateway_namespace[count.index].metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait = true

  # Configure service to use ClusterIP with these ports
  values = [
    yamlencode({
      service = {
        type = "ClusterIP"
        ports = [
          {
            name       = "status-port"
            protocol   = "TCP"
            port       = 15021
            targetPort = 15021
          },
          {
            name       = "http2"
            protocol   = "TCP"
            port       = 80
            targetPort = 80
          },
          {
            name       = "https"
            protocol   = "TCP"
            port       = 443
            targetPort = 443
          },
          {
            name       = "debug"
            protocol   = "TCP"
            port       = 5005
            targetPort = 5005
          }
        ]
      }
    })
  ]

  depends_on = [
    helm_release.istiod
  ]
}

###############################################
# 6. EnvoyFilter to support gRPC-Web traffic
###############################################
resource "kubernetes_manifest" "istio_grpc_web_envoy_filter" {
  count = var.spec.istio.enabled ? 1 : 0

  manifest = yamldecode(<<-EOT
apiVersion: networking.istio.io/v1alpha3
kind: EnvoyFilter
metadata:
  name: grpc-web
  namespace: istio-ingress
spec:
  workloadSelector:
    labels:
      app: "gateway"
      istio: "gateway"
  configPatches:
    - applyTo: HTTP_FILTER
      match:
        context: GATEWAY
        listener:
          filterChain:
            filter:
              name: "envoy.filters.network.http_connection_manager"
              subFilter:
                name: "envoy.filters.http.cors"
      patch:
        operation: INSERT_BEFORE
        value:
          name: "envoy.filters.http.grpc_web"
          typed_config:
            "@type": "type.googleapis.com/envoy.extensions.filters.http.grpc_web.v3.GrpcWeb"
EOT
  )

  depends_on = [
    helm_release.istio_gateway
  ]
}

###############################################
# 7. Internal Load Balancer IP address
###############################################
resource "google_compute_address" "istio_ingress_internal_lb_ip" {
  count        = var.spec.istio.enabled ? 1 : 0
  name         = "gke-${local.resource_id}-ingress-internal"
  project      = var.spec.cluster_project_id
  region       = var.spec.istio.cluster_region
  address_type = "INTERNAL"
  labels = local.final_gcp_labels

  # Must reference the same subnetwork if you want an internal LB
  subnetwork = var.spec.istio.sub_network_self_link
}

###############################################
# 8. Internal LB Service in "istio-ingress" NS
###############################################
resource "kubernetes_service_v1" "istio_ingress_internal_lb" {
  count = var.spec.istio.enabled ? 1 : 0

  metadata {
    name = "ingress-internal"
    namespace = kubernetes_namespace_v1.istio_gateway_namespace[count.index].metadata[0].name

    # Merge final_kubernetes_labels with the app/istio labels
    labels = merge(
      local.final_kubernetes_labels,
      {
        "app"   = "gateway"
        "istio" = "gateway"
      }
    )

    annotations = {
      "cloud.google.com/load-balancer-type" = "internal"
    }
  }

  spec {
    type             = "LoadBalancer"
    load_balancer_ip = google_compute_address.istio_ingress_internal_lb_ip[count.index].address

    selector = {
      "app"   = "gateway"
      "istio" = "gateway"
    }

    port {
      name        = "status-port"
      protocol    = "TCP"
      port        = 15021
      target_port = 15021
    }
    port {
      name        = "http2"
      protocol    = "TCP"
      port        = 80
      target_port = 80
    }
    port {
      name        = "https"
      protocol    = "TCP"
      port        = 443
      target_port = 443
    }
  }
}

###############################################
# 9. External Load Balancer IP address
###############################################
resource "google_compute_address" "istio_ingress_external_lb_ip" {
  count        = var.spec.istio.enabled ? 1 : 0
  name         = "gke-${local.resource_id}-ingress-external"
  project      = var.spec.cluster_project_id
  region       = var.spec.istio.cluster_region
  address_type = "EXTERNAL"
  labels       = local.final_gcp_labels
}

###############################################
# 10. External LB Service in "istio-ingress" NS
###############################################
resource "kubernetes_service_v1" "istio_ingress_external_lb" {
  count = var.spec.istio.enabled ? 1 : 0

  metadata {
    name = "ingress-external"
    namespace = kubernetes_namespace_v1.istio_gateway_namespace[count.index].metadata[0].name

    # Merge final_kubernetes_labels with the app/istio labels
    labels = merge(
      local.final_kubernetes_labels,
      {
        "app"   = "gateway"
        "istio" = "gateway"
      }
    )

    annotations = {
      "cloud.google.com/load-balancer-type" = "external"
    }
  }

  spec {
    type             = "LoadBalancer"
    load_balancer_ip = google_compute_address.istio_ingress_external_lb_ip[count.index].address

    selector = {
      "app"   = "gateway"
      "istio" = "gateway"
    }

    port {
      name        = "status-port"
      protocol    = "TCP"
      port        = 15021
      target_port = 15021
    }
    port {
      name        = "http2"
      protocol    = "TCP"
      port        = 80
      target_port = 80
    }
    port {
      name        = "https"
      protocol    = "TCP"
      port        = 443
      target_port = 443
    }
  }
}

###############################################
# Optional: Export LB IP addresses
###############################################
output "ingress_internal_ip" {
  description = "Internal IP for the istio ingress internal load balancer"
  value       = var.spec.istio.enabled ? google_compute_address.istio_ingress_internal_lb_ip[0].address : null
}

output "ingress_external_ip" {
  description = "External IP for the istio ingress external load balancer"
  value       = var.spec.istio.enabled ? google_compute_address.istio_ingress_external_lb_ip[0].address : null
}
