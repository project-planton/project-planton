locals {
  # Safe locals to avoid null dereferences
  safe_metadata = var.metadata
  safe_spec     = var.spec

  # Computed booleans for conditional flows
  has_sort_key = local.safe_spec.sort_key_name != null && local.safe_spec.sort_key_name != ""
  is_provisioned_billing = local.safe_spec.billing_mode == "PROVISIONED"
  is_pay_per_request_billing = local.safe_spec.billing_mode == "PAY_PER_REQUEST"
  
  # Safe defaults for capacity units
  safe_read_capacity_units = local.is_provisioned_billing ? local.safe_spec.read_capacity_units : null
  safe_write_capacity_units = local.is_provisioned_billing ? local.safe_spec.write_capacity_units : null
  
  # Safe defaults for encryption and recovery
  safe_server_side_encryption_enabled = local.safe_spec.server_side_encryption_enabled != null ? local.safe_spec.server_side_encryption_enabled : true
  safe_point_in_time_recovery_enabled = local.safe_spec.point_in_time_recovery_enabled != null ? local.safe_spec.point_in_time_recovery_enabled : false
}
