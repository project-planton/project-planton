resource "google_container_node_pool" "gke_node_pool" {
  # Loop over each node pool specified in var.spec.node_pools
  for_each = {
    for np in var.spec.node_pools :
    np.name => np
  }

  name       = each.key
  project    = var.spec.cluster_project_id
  location   = var.spec.zone
  cluster    = google_container_cluster.gke_cluster.name
  node_count = each.value.min_node_count

  autoscaling {
    min_node_count = each.value.min_node_count
    max_node_count = each.value.max_node_count
  }

  management {
    auto_repair  = true
    auto_upgrade = true
  }

  node_config {
    labels = local.final_gcp_labels

    machine_type = each.value.machine_type

    metadata = {
      "disable-legacy-endpoints" = "true"
    }

    oauth_scopes = [
      "https://www.googleapis.com/auth/monitoring",
      "https://www.googleapis.com/auth/monitoring.write",
      "https://www.googleapis.com/auth/devstorage.read_only",
      "https://www.googleapis.com/auth/logging.write"
    ]

    preemptible = each.value.is_spot_enabled

    tags = [
      local.network_tag
    ]

    workload_metadata_config {
      mode = "GKE_METADATA"
    }
  }

  upgrade_settings {
    max_surge       = 2
    max_unavailable = 1
  }

  # Ignore changes to node_count so Terraform doesn't forcibly re-scale
  # the node pool if someone adjusts node counts via the GCP console or UI
  lifecycle {
    ignore_changes = [node_count]
    create_before_destroy = false
  }

  depends_on = [
    google_container_cluster.gke_cluster
  ]
}
