###############################################################################
# 1. Enable required APIs for the GCP project where the GKE cluster is created
###############################################################################
resource "google_project_service" "gke_cluster_project_apis" {
  for_each = toset([
    "compute.googleapis.com",
    "container.googleapis.com",
    "secretmanager.googleapis.com",
    "dns.googleapis.com"
  ])

  project            = var.spec.cluster_project_id
  service            = each.value
  disable_on_destroy = true
}

###################################
# 2. Create the VPC network
###################################
resource "google_compute_network" "gke_network" {
  name                    = "vpc"
  project                 = var.spec.cluster_project_id
  auto_create_subnetworks = false
}

#####################################
# 3. Create the subnetwork
#####################################
resource "google_compute_subnetwork" "gke_subnetwork" {
  name                     = var.metadata.name
  project                  = var.spec.cluster_project_id
  region                   = var.spec.region
  network                  = google_compute_network.gke_network.self_link
  ip_cidr_range = "10.0.0.0/14"     # SubNetworkCidr
  private_ip_google_access = true

  secondary_ip_range {
    range_name    = local.kubernetes_pod_secondary_ip_range_name
    ip_cidr_range = "10.4.0.0/16"           # KubernetesPodSecondaryIpRange
  }
  secondary_ip_range {
    range_name    = local.kubernetes_service_secondary_ip_range_name
    ip_cidr_range = "10.5.0.0/16"           # KubernetesServiceSecondaryIpRange
  }
}

###################################
# 4. Create firewall for webhooks
###################################
resource "google_compute_firewall" "gke_webhook_firewall" {
  name    = "${var.metadata.name}-gke-webhook"
  project = var.spec.cluster_project_id
  network = google_compute_network.gke_network.name

  source_ranges = ["172.16.0.0/28"]

  allow {
    protocol = "tcp"
    ports = ["8443", "15017"]
  }

  target_tags = [local.network_tag]
}

############################
# 5. Create the router
############################
resource "google_compute_router" "gke_router" {
  name    = var.metadata.name
  project = var.spec.cluster_project_id
  region  = var.spec.region
  network = google_compute_network.gke_network.self_link
}

#################################################
# 6. Create an external IP address for the NAT
#################################################
resource "google_compute_address" "gke_router_nat_ip" {
  name         = "${var.metadata.name}-router-nat"
  project      = var.spec.cluster_project_id
  region       = var.spec.region
  address_type = "EXTERNAL"
  labels       = local.final_gcp_labels
}

########################################
# 7. Create the router NAT
########################################
resource "google_compute_router_nat" "gke_nat" {
  name                               = var.metadata.name
  project                            = var.spec.cluster_project_id
  region                             = var.spec.region
  router                             = google_compute_router.gke_router.name
  nat_ip_allocate_option             = "MANUAL_ONLY"
  nat_ips = [google_compute_address.gke_router_nat_ip.self_link]
  source_subnetwork_ip_ranges_to_nat = "ALL_SUBNETWORKS_ALL_IP_RANGES"
}

#######################################################################
# 8. Create the GKE cluster with optional autoscaling configuration
#######################################################################
resource "google_container_cluster" "gke_cluster" {
  name                     = var.metadata.name
  project                  = var.spec.cluster_project_id
  location                 = var.spec.zone
  network                  = google_compute_network.gke_network.self_link
  subnetwork               = google_compute_subnetwork.gke_subnetwork.self_link
  remove_default_node_pool = true
  deletion_protection      = false
  initial_node_count       = 1

  # Workload Identity
  workload_identity_config {
    workload_pool = "${var.spec.cluster_project_id}.svc.id.goog"
  }

  # GKE Release channel
  release_channel {
    channel = "STABLE"
  }

  # Vertical Pod Autoscaling
  vertical_pod_autoscaling {
    enabled = true
  }

  # Standard addons config
  addons_config {
    horizontal_pod_autoscaling {
      disabled = false
    }
    http_load_balancing {
      disabled = true
    }
    network_policy_config {
      disabled = true
    }
  }

  # Private cluster config
  private_cluster_config {
    enable_private_endpoint = false
    enable_private_nodes    = true
    master_ipv4_cidr_block  = "172.16.0.0/28"
  }

  # IP allocation policy
  ip_allocation_policy {
    cluster_secondary_range_name  = local.kubernetes_pod_secondary_ip_range_name
    services_secondary_range_name = local.kubernetes_service_secondary_ip_range_name
  }

  # Master authorized networks
  master_authorized_networks_config {
    cidr_blocks {
      cidr_block   = "0.0.0.0/0"
      display_name = "kubectl-from-anywhere"
    }
  }

  # Autoscaling (enabled only if .spec.cluster_autoscaling_config.is_enabled == true)
  dynamic "cluster_autoscaling" {
    for_each = (
      var.spec.cluster_autoscaling_config != null && var.spec.cluster_autoscaling_config.is_enabled
      ? [var.spec.cluster_autoscaling_config]
      : []
    )

    content {
      enabled             = true
      autoscaling_profile = "OPTIMIZE_UTILIZATION"
      resource_limits {
        resource_type = "cpu"
        minimum       = cluster_autoscaling.value.cpu_min_cores
        maximum       = cluster_autoscaling.value.cpu_max_cores
      }
      resource_limits {
        resource_type = "memory"
        minimum       = cluster_autoscaling.value.memory_min_gb
        maximum       = cluster_autoscaling.value.memory_max_gb
      }
    }
  }

  # Logging config
  logging_config {
    enable_components = local.container_cluster_logging_component_list
  }

  # Make sure APIs are enabled before creating the cluster
  depends_on = [google_project_service.gke_cluster_project_apis]

  lifecycle {
    # Keep ignoring the default node pool removal (already set to remove_default_node_pool = true)
    ignore_changes = [remove_default_node_pool]
  }
}
