###############################################################################
# Cert Manager
#
# 1. Create a Google Service Account (GSA) for cert-manager, with Workload Identity binding.
# 2. Create a namespace for cert-manager, labeled with our final_kubernetes_labels.
# 3. Create a Kubernetes Service Account (KSA) in that namespace, annotated with the GSA email.
# 4. Deploy the cert-manager Helm chart with CRDs enabled, linking to the KSA.
# 5. For each ingress DNS domain that has is_tls_enabled = true, create a ClusterIssuer.
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

# IAM binding granting "iam.workloadIdentityUser" to the KSA identity
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
  wait             = true

  # We disable the creation of a service account in the Helm chart, since we
  # create it ourselves with the GSA annotation for Workload Identity.
  values = [
    yamlencode({
      installCRDs = true
      extraArgs   = [
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
# 5. ClusterIssuer for each TLS-enabled domain
##############################################
# For each domain where is_tls_enabled = true, create a cert-manager ClusterIssuer resource
resource "kubernetes_manifest" "cert_manager_cluster_issuer" {
  for_each = {
    for domain in var.spec.ingress_dns_domains :
    domain.name => domain
    if domain.is_tls_enabled
  }

  manifest = yamldecode(
    templatefile(
      "${path.module}/templates/cluster_issuer.yaml.tpl",
      {
        issuer_name                = each.key
        dns_zone_gcp_project_id    = each.value.dns_zone_gcp_project_id
      }
    )
  )

  depends_on = [
    helm_release.cert_manager
  ]
}

# Optional example of the cluster_issuer.yaml.tpl template:
#
#  apiVersion: cert-manager.io/v1
#  kind: ClusterIssuer
#  metadata:
#    name: ${issuer_name}
#    labels:
#      <<LABELS>>
#  spec:
#    acme:
#      server: https://acme-v02.api.letsencrypt.org/directory
#      preferredChain: ""
#      privateKeySecretRef:
#        name: "letsencrypt-production"
#      solvers:
#      - dns01:
#          cloudDNS:
#            project: ${dns_zone_gcp_project_id}
#      - http01:
#          ingress:
#            class: istio
#
#  NOTE: In your actual usage, replace <<LABELS>> with local.final_kubernetes_labels
#        or merge them similarly. If you're using a single-blob template, you'd
#        embed them just like in other resources. Or you can construct them
#        via string interpolation in the templatefile() call if desired.
#
# The above approach allows dynamic creation of ClusterIssuers for each domain that has TLS enabled.
