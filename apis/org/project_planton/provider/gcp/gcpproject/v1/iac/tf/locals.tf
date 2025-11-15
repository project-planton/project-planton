# locals.tf - Local value transformations for GCP Project

locals {
  # Generate a random 3-character suffix for globally unique project_id
  # Project ID format: <metadata.name>-<3-char-suffix>
  # Ensures no collisions even if someone else used the same name before
  project_id_suffix = random_string.project_suffix.result

  # Make metadata.name safe for GCP project_id constraints:
  # - lowercase letters, digits, hyphens only
  # - must start with letter
  # - 6-30 chars total
  # - cannot end with hyphen
  safe_name = lower(
    replace(
      replace(var.metadata.name, "_", "-"), # Replace underscores with hyphens
      "/[^a-z0-9-]/", "-"                   # Replace any invalid chars with hyphen
    )
  )

  # Ensure safe_name starts with a letter
  safe_name_prefix = can(regex("^[a-z]", local.safe_name)) ? local.safe_name : "p${local.safe_name}"

  # Trim to leave space for suffix ("-xyz")
  # Max length is 30, minus 4 for "-xyz" = 26 chars for name
  safe_name_trimmed = substr(local.safe_name_prefix, 0, min(length(local.safe_name_prefix), 26))

  # Remove trailing hyphens if any
  safe_name_clean = trimright(local.safe_name_trimmed, "-")

  # Final project_id with suffix
  project_id = "${local.safe_name_clean}-${local.project_id_suffix}"

  # GCP labels: merge user-provided labels with standard Planton labels
  # User labels come first, then standard labels (which override if there's a conflict)
  gcp_labels = merge(
    var.spec.labels != null ? var.spec.labels : {},
    {
      "planton-cloud-resource"      = "true"
      "planton-cloud-resource-name" = var.metadata.name
      "planton-cloud-resource-kind" = "gcpproject"
      "planton-cloud-resource-id"   = var.metadata.id != null ? var.metadata.id : var.metadata.name
      "planton-cloud-org"           = var.metadata.org != null ? var.metadata.org : "default"
      "planton-cloud-env"           = var.metadata.env != null ? var.metadata.env : "default"
    }
  )

  # Determine parent resource configuration
  parent_type      = var.spec.parent_type
  parent_org_id    = local.parent_type == "organization" ? var.spec.parent_id : null
  parent_folder_id = local.parent_type == "folder" ? var.spec.parent_id : null

  # Billing account configuration
  billing_account_id = var.spec.billing_account_id

  # Network configuration
  # disable_default_network defaults to true if not specified
  auto_create_network = var.spec.disable_default_network != null ? !var.spec.disable_default_network : false

  # APIs to enable
  enabled_apis = var.spec.enabled_apis != null ? var.spec.enabled_apis : []

  # IAM owner member (optional)
  owner_member = var.spec.owner_member != null && var.spec.owner_member != "" ? var.spec.owner_member : null
}

