##############################################
# 1. Google Service Account for cert-manager
##############################################
resource "google_service_account" "cert_manager_gsa" {
  count        = var.spec.install_cert_manager ? 1 : 0
  project      = var.spec.cluster_project_id
  account_id   = "cert-manager"
  display_name = "cert-manager"
  description  = "Service Account for Cert-Manager to solve DNS challenges"
}

resource "google_service_account_iam_binding" "cert_manager_workload_identity_binding" {
  count              = var.spec.install_cert_manager ? 1 : 0
  service_account_id = google_service_account.cert_manager_gsa[count.index].name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "serviceAccount:${var.spec.cluster_project_id}.svc.id.goog[cert-manager/cert-manager]"
  ]
}

# Optionally export the GSA email (will be null if install_cert_manager is false)
output "cert_manager_gsa_email" {
  description = "Email of the Cert Manager GSA"
  value       = var.spec.install_cert_manager ? google_service_account.cert_manager_gsa[0].email : null
}

##############################################
# 2. cert-manager Namespace
##############################################
resource "kubernetes_namespace_v1" "cert_manager_namespace" {
  count = var.spec.install_cert_manager ? 1 : 0

  metadata {
    name   = "cert-manager"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 3. KSA annotated with the GSA email
##############################################
resource "kubernetes_service_account_v1" "cert_manager_ksa" {
  count = var.spec.install_cert_manager ? 1 : 0

  metadata {
    name = "cert-manager"
    # Refer to the cert-manager namespace created above
    namespace = kubernetes_namespace_v1.cert_manager_namespace[count.index].metadata[0].name

    annotations = {
      "iam.gke.io/gcp-service-account" = google_service_account.cert_manager_gsa[count.index].email
    }

    labels = local.final_kubernetes_labels
  }
}

##############################################
# 4. Helm Release for cert-manager
##############################################
resource "helm_release" "cert_manager" {
  count           = var.spec.install_cert_manager ? 1 : 0
  name            = "cert-manager"
  repository      = "https://charts.jetstack.io"
  chart           = "cert-manager"
  version         = "v1.15.2"
  create_namespace = false
  # Refer to the cert-manager namespace created above
  namespace       = kubernetes_namespace_v1.cert_manager_namespace[count.index].metadata[0].name
  timeout         = 180
  cleanup_on_fail = true
  atomic          = false
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

  depends_on = [
    kubernetes_service_account_v1.cert_manager_ksa,
    google_service_account_iam_binding.cert_manager_workload_identity_binding
  ]
}
