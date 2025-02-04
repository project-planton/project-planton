resource "aws_rds_cluster" "this" {
  count = 1

  # Basic settings
  cluster_identifier                  = local.resource_id
  engine                              = var.spec.engine
  engine_version                      = var.spec.engine_version
  engine_mode                         = var.spec.engine_mode
  allow_major_version_upgrade         = var.spec.allow_major_version_upgrade
  iam_database_authentication_enabled = var.spec.iam_database_authentication_enabled
  port                                = var.spec.database_port > 0 ? var.spec.database_port : null
  network_type                        = "IPV4"
  apply_immediately                   = true
  copy_tags_to_snapshot               = false
  deletion_protection                 = var.spec.deletion_protection
  preferred_maintenance_window        = var.spec.maintenance_window
  preferred_backup_window             = var.spec.backup_window
  backup_retention_period             = var.spec.retention_period > 0 ? var.spec.retention_period : 5
  skip_final_snapshot                 = var.spec.skip_final_snapshot
  final_snapshot_identifier           = var.spec.skip_final_snapshot ? null : "${local.resource_id}-final-snapshot"
  snapshot_identifier                 = var.spec.snapshot_identifier != "" ? var.spec.snapshot_identifier : null
  tags = local.final_labels

  # Master user / password handling
  manage_master_user_password = var.spec.manage_master_user_password
  master_username             = var.spec.master_user
  master_password             = var.spec.manage_master_user_password ? null : var.spec.master_password
  master_user_secret_kms_key_id = (var.spec.manage_master_user_password && var.spec.master_user_secret_kms_key_id != ""
    ? var.spec.master_user_secret_kms_key_id
    : null)

  # Encryption at rest
  storage_encrypted = (
    var.spec.engine_mode != "serverless"
    ? var.spec.storage_encrypted
    : false
  )
  kms_key_id = (
    var.spec.engine_mode != "serverless" && var.spec.storage_encrypted
    ? var.spec.storage_kms_key_arn
    : null
  )

  # VPC Security Groups
  vpc_security_group_ids = concat(
    [aws_security_group.default.id],
    var.spec.associate_security_group_ids
  )

  # Subnet group
  db_subnet_group_name = (
    length(var.spec.subnet_ids) > 0 && (
    var.spec.db_subnet_group_name == null || var.spec.db_subnet_group_name == ""
    )
    ? aws_db_subnet_group.default[0].name
    : var.spec.db_subnet_group_name
  )

  # Cluster parameter group
  db_cluster_parameter_group_name = (
    var.spec.cluster_parameter_group_name == null || var.spec.cluster_parameter_group_name == ""
    ? aws_rds_cluster_parameter_group.this[0].name
    : var.spec.cluster_parameter_group_name
  )

  # Serverless v1 scaling
  dynamic "scaling_configuration" {
    for_each = (
      var.spec.engine_mode == "serverless" && var.spec.scaling_configuration != null
      ? [var.spec.scaling_configuration]
      : []
    )
    content {
      auto_pause = lookup(scaling_configuration.value, "auto_pause", false)
      max_capacity = (lookup(scaling_configuration.value, "max_capacity", 0) > 0
        ? scaling_configuration.value.max_capacity
        : 16)
      min_capacity = (lookup(scaling_configuration.value, "min_capacity", 0) > 0
        ? scaling_configuration.value.min_capacity
        : 2)
      seconds_until_auto_pause = (lookup(scaling_configuration.value, "seconds_until_auto_pause", 0) > 0
        ? scaling_configuration.value.seconds_until_auto_pause
        : 300)
      timeout_action = (lookup(scaling_configuration.value, "timeout_action", "") != ""
        ? scaling_configuration.value.timeout_action
        : "RollbackCapacityChange")
    }
  }

  # Serverless v2 scaling
  dynamic "serverlessv2_scaling_configuration" {
    for_each = (
      var.spec.engine_mode == "provisioned" && var.spec.serverlessv2_scaling_configuration != null
      ? [var.spec.serverlessv2_scaling_configuration]
      : []
    )
    content {
      max_capacity = serverlessv2_scaling_configuration.value.max_capacity
      min_capacity = serverlessv2_scaling_configuration.value.min_capacity
    }
  }

  depends_on = [
    aws_security_group.default
  ]
}
