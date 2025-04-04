##############################################
# 1. Create Service Account for Workload Deployer
##############################################
resource "google_service_account" "workload_deployer_sa" {
  project       = var.spec.cluster_project_id
  account_id    = "workload-deployer"
  display_name  = "workload-deployer"
  description   = "Service account to deploy workloads"
}

##############################################
# 2. Create a Key for the Service Account
##############################################
resource "google_service_account_key" "workload_deployer_sa_key" {
  service_account_id = google_service_account.workload_deployer_sa.name
}

################################################
# 3. Assign Roles to the Workload Deployer SA
################################################

# Role: container.admin
resource "google_project_iam_binding" "workload_deployer_container_admin" {
  project = var.spec.cluster_project_id
  role    = "roles/container.admin"
  members = [
    "serviceAccount:${google_service_account.workload_deployer_sa.email}"
  ]
}

# Role: container.clusterAdmin
resource "google_project_iam_binding" "workload_deployer_cluster_admin" {
  project = var.spec.cluster_project_id
  role    = "roles/container.clusterAdmin"
  members = [
    "serviceAccount:${google_service_account.workload_deployer_sa.email}"
  ]
}
