# Exercise the real module's dynamic blocks without credentials or cloud calls.
# Inputs include the catalog defaults normally populated before module rendering.
# Explicitly disable legacy client certificates so the mock can populate the
# computed CA certificate on a configured master_auth block.
mock_provider "google" {
  mock_resource "google_container_cluster" {
    defaults = {
      master_auth = { cluster_ca_certificate = "public-ca" }
    }
  }
}

variables {
  metadata = { name = "policy-test" }
  spec = {
    project_id                = "test-project"
    location                  = "us-central1"
    network                   = "test-vpc"
    subnetwork                = "test-subnet"
    release_channel           = "REGULAR"
    workload_identity_enabled = true
    issue_client_certificate  = false
  }
}

run "omitted_policy" {
  command = plan

  assert {
    condition     = length(google_container_cluster.this.master_authorized_networks_config) == 0
    error_message = "An omitted policy must not emit an authorized-networks restriction."
  }
}

run "explicit_unrestricted" {
  command = plan

  variables {
    spec = {
      project_id                = "test-project"
      location                  = "us-central1"
      network                   = "test-vpc"
      subnetwork                = "test-subnet"
      release_channel           = "REGULAR"
      workload_identity_enabled = true
      issue_client_certificate  = false
      unrestricted_networks     = {}
    }
  }

  assert {
    condition     = length(google_container_cluster.this.master_authorized_networks_config) == 0
    error_message = "Explicit unrestricted access must not emit an authorized-networks restriction."
  }
}

run "empty_restriction" {
  command = plan

  variables {
    spec = {
      project_id                 = "test-project"
      location                   = "us-central1"
      network                    = "test-vpc"
      subnetwork                 = "test-subnet"
      release_channel            = "REGULAR"
      workload_identity_enabled  = true
      issue_client_certificate   = false
      master_authorized_networks = {}
    }
  }

  assert {
    condition     = length(google_container_cluster.this.master_authorized_networks_config) == 1
    error_message = "A configured empty restriction must keep its provider block."
  }
  assert {
    condition     = length(google_container_cluster.this.master_authorized_networks_config[0].cidr_blocks) == 0
    error_message = "An empty restriction must not invent an authorized CIDR."
  }
}

run "explicit_false_switches" {
  command = plan

  variables {
    spec = {
      project_id                = "test-project"
      location                  = "us-central1"
      network                   = "test-vpc"
      subnetwork                = "test-subnet"
      release_channel           = "REGULAR"
      workload_identity_enabled = true
      issue_client_certificate  = false
      master_authorized_networks = {
        gcp_public_cidrs_access_enabled      = false
        private_endpoint_enforcement_enabled = false
      }
    }
  }

  assert {
    condition = (
      google_container_cluster.this.master_authorized_networks_config[0].gcp_public_cidrs_access_enabled == false &&
      google_container_cluster.this.master_authorized_networks_config[0].private_endpoint_enforcement_enabled == false
    )
    error_message = "Explicit false policy switches must remain false."
  }
}

run "populated_restriction" {
  command = plan

  variables {
    spec = {
      project_id                = "test-project"
      location                  = "us-central1"
      network                   = "test-vpc"
      subnetwork                = "test-subnet"
      release_channel           = "REGULAR"
      workload_identity_enabled = true
      issue_client_certificate  = false
      master_authorized_networks = {
        cidr_blocks                          = [{ cidr_block = "203.0.113.0/24", display_name = "office" }]
        private_endpoint_enforcement_enabled = true
      }
    }
  }

  assert {
    condition = (
      one(google_container_cluster.this.master_authorized_networks_config[0].cidr_blocks).cidr_block == "203.0.113.0/24" &&
      one(google_container_cluster.this.master_authorized_networks_config[0].cidr_blocks).display_name == "office" &&
      google_container_cluster.this.master_authorized_networks_config[0].private_endpoint_enforcement_enabled == true
    )
    error_message = "The restriction must preserve its CIDR, display name and private-endpoint switch."
  }
}
