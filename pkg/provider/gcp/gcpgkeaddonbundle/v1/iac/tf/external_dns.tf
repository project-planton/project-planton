###############################################################################
# External DNS
#
# 1. Create a Google Service Account (GSA) for external-dns, with Workload Identity binding.
# 2. Create an external-dns namespace, labeled with final_kubernetes_labels.
# 3. Create a Kubernetes Service Account (KSA), annotated with the GSA email.
###############################################################################

##############################################
# 1. Google Service Account for external-dns
##############################################
resource "google_service_account" "external_dns_gsa" {
  count        = var.spec.install_external_dns ? 1 : 0
  project      = var.spec.cluster_project_id
  account_id   = "external-dns"
  display_name = "external-dns"
  description  = "GSA for external-dns to manage DNS recordsets in Cloud DNS zones"
}

# IAM binding granting "iam.workloadIdentityUser" to the KSA identity
resource "google_service_account_iam_binding" "external_dns_workload_identity_binding" {
  count              = var.spec.install_external_dns ? 1 : 0
  service_account_id = google_service_account.external_dns_gsa[count.index].name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "serviceAccount:${var.spec.cluster_project_id}.svc.id.goog[external-dns/external-dns]"
  ]
}

output "external_dns_gsa_email" {
  description = "Email of the external-dns GSA"
  # Return null if not installed
  value       = var.spec.install_external_dns ? google_service_account.external_dns_gsa[0].email : null
}

##############################################
# 2. external-dns Namespace
##############################################
resource "kubernetes_namespace_v1" "external_dns_namespace" {
  count = var.spec.install_external_dns ? 1 : 0

  metadata {
    name   = "external-dns"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 3. KSA annotated with GSA email
##############################################
resource "kubernetes_service_account_v1" "external_dns_ksa" {
  count = var.spec.install_external_dns ? 1 : 0

  metadata {
    name      = "external-dns"
    namespace = kubernetes_namespace_v1.external_dns_namespace[count.index].metadata[0].name

    annotations = {
      "iam.gke.io/gcp-service-account" = google_service_account.external_dns_gsa[count.index].email
    }
    labels = local.final_kubernetes_labels
  }
}
