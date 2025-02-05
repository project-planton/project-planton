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
  project      = var.spec.cluster_project_id
  account_id   = "external-secrets"
  display_name = "external-secrets"
  description  = "GSA for External Secrets to read secrets from Secret Manager"
}

# Bind the GSA to the "secretmanager.secretAccessor" role
resource "google_project_iam_binding" "external_secrets_secret_accessor" {
  project = var.spec.cluster_project_id
  role    = "roles/secretmanager.secretAccessor"
  members = [
    "serviceAccount:${google_service_account.external_secrets_gsa.email}"
  ]
}

# IAM binding granting "iam.workloadIdentityUser" to the KSA identity
resource "google_service_account_iam_binding" "external_secrets_workload_identity_binding" {
  service_account_id = google_service_account.external_secrets_gsa.name
  role               = "roles/iam.workloadIdentityUser"
  members = [
    "serviceAccount:${var.spec.cluster_project_id}.svc.id.goog[external-secrets/external-secrets]"
  ]
}

output "external_secrets_gsa_email" {
  description = "Email of the External Secrets GSA"
  value       = google_service_account.external_secrets_gsa.email
}

##############################################
# 2. external-secrets Namespace
##############################################
resource "kubernetes_namespace_v1" "external_secrets_namespace" {
  metadata {
    name   = "external-secrets"
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 3. KSA annotated with GSA email
##############################################
resource "kubernetes_service_account_v1" "external_secrets_ksa" {
  metadata {
    name      = "external-secrets"
    namespace = kubernetes_namespace_v1.external_secrets_namespace.metadata[0].name

    annotations = {
      "iam.gke.io/gcp-service-account" = google_service_account.external_secrets_gsa.email
    }
    labels = local.final_kubernetes_labels
  }
}

##############################################
# 4. Helm Release for External Secrets
##############################################
resource "helm_release" "external_secrets" {
  name             = "external-secrets"
  repository       = "https://charts.external-secrets.io"
  chart            = "external-secrets"
  version          = "v0.9.20"
  create_namespace = false
  namespace        = kubernetes_namespace_v1.external_secrets_namespace.metadata[0].name
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
        POLLER_INTERVAL_MILLISECONDS = 10 * 1000  # e.g., 10 seconds
        LOG_LEVEL       = "info"
        LOG_MESSAGE_KEY = "msg"
        METRICS_PORT    = 3001
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

  lifecycle {
    ignore_changes = [
      status,
      description
    ]
  }

  depends_on = [
    kubernetes_service_account_v1.external_secrets_ksa,
    google_service_account_iam_binding.external_secrets_workload_identity_binding
  ]
}

##############################################
# 5. ClusterSecretStore referencing GCP SM
##############################################
# We define a simple "ClusterSecretStore" that references the GCP project
# from which secrets should be retrieved.
###############################################################################
resource "kubernetes_manifest" "external_secrets_cluster_secret_store" {
  manifest = {
    apiVersion = "external-secrets.io/v1beta1"
    kind       = "ClusterSecretStore"
    metadata = {
      name   = "gcp-secrets-manager"
      labels = local.final_kubernetes_labels
    }
    spec = {
      provider = {
        gcpsm = {
          projectId = var.spec.cluster_project_id
        }
      }
      # you can specify a string with the time unit, e.g., "10s"
      refreshInterval = "10s"
    }
  }

  depends_on = [
    helm_release.external_secrets
  ]
}
