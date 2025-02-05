###############################################################################
# External DNS
#
# 1. Create a Google Service Account (GSA) for external-dns, with Workload Identity binding.
# 2. Create an external-dns namespace, labeled with final_kubernetes_labels.
# 3. Create a Kubernetes Service Account (KSA), annotated with the GSA email.
# 4. For each TLS domain in var.spec.ingress_dns_domains, create a Helm release
#    that installs external-dns configured with domainFilters = [that domain].
###############################################################################

##############################################
# 1. Google Service Account for external-dns
##############################################
resource "google_service_account" "external_dns_gsa" {
  project      = var.spec.cluster_project_id
  account_id   = "external-dns"
  display_name = "external-dns"
  description  = "GSA for external-dns to manage DNS recordsets in Cloud DNS zones"
}

# IAM binding granting "iam.workloadIdentityUser" to the KSA identity
resource "google_service_account_iam_binding" "external_dns_workload_identity_binding" {
  service_account_id = google_service_account.external_dns_gsa.name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "serviceAccount:${var.spec.cluster_project_id}.svc.id.goog[external-dns/external-dns]"
  ]
}

output "external_dns_gsa_email" {
  description = "Email of the external-dns GSA"
  value       = google_service_account.external_dns_gsa.email
}

##############################################
# 2. external-dns Namespace
##############################################
resource "kubernetes_namespace_v1" "external_dns_namespace" {
  metadata {
    name   = "external-dns"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 3. KSA annotated with GSA email
##############################################
resource "kubernetes_service_account_v1" "external_dns_ksa" {
  metadata {
    name      = "external-dns"
    namespace = kubernetes_namespace_v1.external_dns_namespace.metadata[0].name

    annotations = {
      "iam.gke.io/gcp-service-account" = google_service_account.external_dns_gsa.email
    }
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 4. Helm Release(s) for each DNS domain
##############################################
resource "helm_release" "external_dns" {
  # We'll create one release per domain that appears in var.spec.ingress_dns_domains.
  for_each = {
    for domain in var.spec.ingress_dns_domains :
    domain.name => domain
  }

  name             = "external-dns-${replace(each.value.name, ".", "-")}"
  repository       = "https://kubernetes-sigs.github.io/external-dns/"
  chart            = "external-dns"
  version          = "1.14.4"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.external_dns_namespace.metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait             = true

  # This release is domain-specific. We set .Values.txtOwnerId to local.resource_id
  # and .Values.domainFilters to [the domain.name].
  # We also disable service account creation in the chart, so we can apply the KSA with annotation.
  values = [
    yamlencode({
      txtOwnerId    = local.resource_id
      serviceAccount = {
        create = false
        name   = "external-dns"
      }
      domainFilters = [each.value.name]
      # also specify gateway + service source if you need them
      sources = [
        "service",
        "gateway-httproute"
      ]
      provider = "google"
      extraArgs = [
        "--google-zone-visibility=public",
        "--google-project=${each.value.dns_zone_gcp_project_id}",
      ]
    })
  ]

  lifecycle {
    ignore_changes = [
      status,
      description
    ]
  }

  depends_on = [
    kubernetes_service_account_v1.external_dns_ksa,
    google_service_account_iam_binding.external_dns_workload_identity_binding
  ]
}
