###############################################################################
# Cert Manager (Fully Inline ClusterIssuer)
###############################################################################

##############################################
# 1. Google Service Account for cert-manager
##############################################
resource "google_service_account" "cert_manager_gsa" {
  project      = var.spec.cluster_project_id
  account_id   = "cert-manager"
  display_name = "cert-manager"
  description  = "Service Account for Cert-Manager to solve DNS challenges"
}

resource "google_service_account_iam_binding" "cert_manager_workload_identity_binding" {
  service_account_id = google_service_account.cert_manager_gsa.name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "serviceAccount:${var.spec.cluster_project_id}.svc.id.goog[cert-manager/cert-manager]"
  ]
}

output "cert_manager_gsa_email" {
  description = "Email of the Cert Manager GSA"
  value       = google_service_account.cert_manager_gsa.email
}

##############################################
# 2. cert-manager Namespace
##############################################
resource "kubernetes_namespace_v1" "cert_manager_namespace" {
  metadata {
    name   = "cert-manager"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 3. KSA annotated with the GSA email
##############################################
resource "kubernetes_service_account_v1" "cert_manager_ksa" {
  metadata {
    name      = "cert-manager"
    namespace = kubernetes_namespace_v1.cert_manager_namespace.metadata[0].name

    annotations = {
      "iam.gke.io/gcp-service-account" = google_service_account.cert_manager_gsa.email
    }

    labels = local.final_kubernetes_labels
  }
}

##############################################
# 4. Helm Release for cert-manager
##############################################
resource "helm_release" "cert_manager" {
  name             = "cert-manager"
  repository       = "https://charts.jetstack.io"
  chart            = "cert-manager"
  version          = "v1.15.2"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.cert_manager_namespace.metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait = true

  # Disabling the chart's default ServiceAccount creation
  values = [
    yamlencode({
      installCRDs = true
      extraArgs = [
        "--dns01-recursive-nameservers-only=true",
        "--dns01-recursive-nameservers=8.8.8.8:53,1.1.1.1:53"
      ]
      serviceAccount = {
        create = false
        name   = "cert-manager"
      }
    })
  ]

  lifecycle {
    ignore_changes = [
      status,
      description
    ]
  }

  depends_on = [
    kubernetes_service_account_v1.cert_manager_ksa,
    google_service_account_iam_binding.cert_manager_workload_identity_binding
  ]
}

##############################################
# 5. Fully Inline ClusterIssuer for each TLS-enabled domain
##############################################
resource "kubernetes_manifest" "cert_manager_cluster_issuer" {
  # Create a resource for each domain that has is_tls_enabled = true
  for_each = {
    for domain in var.spec.ingress_dns_domains :
    domain.name => domain
    if domain.is_tls_enabled
  }

  # Use a direct HCL object for the YAML manifest (Terraform converts to JSON or YAML automatically)
  manifest = {
    apiVersion = "cert-manager.io/v1"
    kind       = "ClusterIssuer"
    metadata = {
      name   = each.key
      labels = local.final_kubernetes_labels
    }
    spec = {
      acme = {
        server         = "https://acme-v02.api.letsencrypt.org/directory"
        preferredChain = ""
        privateKeySecretRef = {
          name = "letsencrypt-production"
        }
        solvers = [
          {
            dns01 = {
              cloudDNS = {
                project = each.value.dns_zone_gcp_project_id
              }
            }
          },
          {
            http01 = {
              ingress = {
                class = "istio"
              }
            }
          }
        ]
      }
    }
  }

  depends_on = [
    helm_release.cert_manager
  ]
}
