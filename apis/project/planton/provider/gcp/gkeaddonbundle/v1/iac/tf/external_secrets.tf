###############################################################################
# External Secrets
#
# 1. Google Service Account (GSA) for External Secrets, plus roles:
#    - roles/secretmanager.secretAccessor
#    - roles/iam.workloadIdentityUser
# 2. Create the external-secrets namespace, labeled with final_kubernetes_labels.
# 3. Create a Kubernetes Service Account (KSA) annotated with the GSA email.
# 4. Deploy the External Secrets Helm chart, disabling default SA creation.
# 5. Create a ClusterSecretStore referencing GCP Secrets Manager for polling.
###############################################################################

##############################################
# 1. Google Service Account
##############################################
resource "google_service_account" "external_secrets_gsa" {
  count        = var.spec.install_external_secrets ? 1 : 0
  project      = var.spec.cluster_project_id
  account_id   = "external-secrets"
  display_name = "external-secrets"
  description  = "GSA for External Secrets to read secrets from Secret Manager"
}

# Bind the GSA to the "secretmanager.secretAccessor" role
resource "google_project_iam_binding" "external_secrets_secret_accessor" {
  count   = var.spec.install_external_secrets ? 1 : 0
  project = var.spec.cluster_project_id
  role    = "roles/secretmanager.secretAccessor"
  members = [
    "serviceAccount:${google_service_account.external_secrets_gsa[count.index].email}"
  ]
}

# IAM binding granting "iam.workloadIdentityUser" to the KSA identity
resource "google_service_account_iam_binding" "external_secrets_workload_identity_binding" {
  count              = var.spec.install_external_secrets ? 1 : 0
  service_account_id = google_service_account.external_secrets_gsa[count.index].name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "serviceAccount:${var.spec.cluster_project_id}.svc.id.goog[external-secrets/external-secrets]"
  ]
}

output "external_secrets_gsa_email" {
  description = "Email of the External Secrets GSA"
  value       = var.spec.install_external_secrets ? google_service_account.external_secrets_gsa[0].email : null
}

##############################################
# 2. external-secrets Namespace
##############################################
resource "kubernetes_namespace_v1" "external_secrets_namespace" {
  count = var.spec.install_external_secrets ? 1 : 0

  metadata {
    name   = "external-secrets"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 3. KSA annotated with GSA email
##############################################
resource "kubernetes_service_account_v1" "external_secrets_ksa" {
  count = var.spec.install_external_secrets ? 1 : 0

  metadata {
    name      = "external-secrets"
    namespace = kubernetes_namespace_v1.external_secrets_namespace[count.index].metadata[0].name

    annotations = {
      "iam.gke.io/gcp-service-account" = google_service_account.external_secrets_gsa[count.index].email
    }
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 4. Helm Release for External Secrets
##############################################
resource "helm_release" "external_secrets" {
  count            = var.spec.install_external_secrets ? 1 : 0
  name             = "external-secrets"
  repository       = "https://charts.external-secrets.io"
  chart            = "external-secrets"
  version          = "v0.9.20"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.external_secrets_namespace[count.index].metadata[0].name
  timeout          = 180
  cleanup_on_fail  = true
  atomic           = false
  wait = true

  # Pass the polling interval, disable service account creation, etc.
  values = [
    yamlencode({
      customResourceManagerDisabled = false
      crds = {
        create = true
      }
      env = {
        POLLER_INTERVAL_MILLISECONDS = 10 * 1000
        LOG_LEVEL                    = "info"
        LOG_MESSAGE_KEY              = "msg"
        METRICS_PORT                 = 3001
      }
      rbac = {
        create = true
      }
      serviceAccount = {
        create = false
        name   = "external-secrets"
      }
      replicaCount = 1
    })
  ]

  depends_on = [
    kubernetes_service_account_v1.external_secrets_ksa,
    google_service_account_iam_binding.external_secrets_workload_identity_binding
  ]
}

# problem: https://github.com/hashicorp/terraform-provider-kubernetes/issues/1367
# workaround 1: https://github.com/hashicorp/terraform-provider-kubernetes/issues/1367#issuecomment-2260043939
# workaround 2: use kubectl_manifest instead of helm_release
# https://github.com/gavinbunney/terraform-provider-kubectl
# https://registry.terraform.io/providers/ddzero2c/kubectl/latest/docs/resources/kubectl_manifest
##############################################
# 5. ClusterSecretStore referencing GCP SM
##############################################
# We define a simple "ClusterSecretStore" that references the GCP project
# from which secrets should be retrieved.
###############################################################################
resource "kubectl_manifest" "external_secrets_cluster_secret_store" {
  count = var.spec.install_external_secrets ? 1 : 0

  # You can inline your YAML in the resource. If you had multiple docs (---),
  # kubectl_manifest supports them out-of-the-box, but here it's just one.
  yaml_body = <<-EOF
apiVersion: external-secrets.io/v1beta1
kind: ClusterSecretStore
metadata:
  name: gcp-secrets-manager
spec:
  provider:
    gcpsm:
      projectId: ${var.spec.cluster_project_id}
  refreshInterval: 10
EOF

  depends_on = [
    helm_release.external_secrets
  ]
}
