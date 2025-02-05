###############################################################################
# Generate a random UUID to approximate ULID behavior for final snapshot name
###############################################################################
resource "random_uuid" "this" {}

locals {
  # If skip_final_snapshot is false, we generate a unique snapshot name
  rds_final_snapshot_name = "${local.resource_id}-${random_uuid.this.result}"
}

###############################################################################
# Create the AWS RDS Instance
###############################################################################
resource "aws_db_instance" "this" {
  # A stable identifier for the RDS instance
  identifier = local.resource_id

  ###########################################################################
  # Replication logic (skip engine/allocated if replicate_source_db is used)
  ###########################################################################
  replicate_source_db = (
    (var.spec.replicate_source_db != "")
    ? var.spec.replicate_source_db
    : null
  )

  # Skip engine details if replicating from an existing DB
  engine = (
    var.spec.replicate_source_db == ""
    ? var.spec.engine
    : null
  )
  engine_version = (
    var.spec.replicate_source_db == ""
    ? var.spec.engine_version
    : null
  )

  # Skip allocated_storage if replicating from an existing DB
  allocated_storage = (
    var.spec.replicate_source_db == ""
    ? var.spec.allocated_storage
    : null
  )

  ###########################################################################
  # Basic DB configuration
  ###########################################################################
  db_name = (
    var.spec.db_name != ""
    ? var.spec.db_name
    : null
  )
  port = (
    var.spec.port != 0
    ? var.spec.port
    : null
  )
  character_set_name = (
    var.spec.character_set_name != ""
    ? var.spec.character_set_name
    : null
  )
  instance_class = var.spec.instance_class

  max_allocated_storage = (
    var.spec.max_allocated_storage != 0
    ? var.spec.max_allocated_storage
    : null
  )
  storage_encrypted = var.spec.storage_encrypted
  kms_key_id        = (
    var.spec.kms_key_id != ""
    ? var.spec.kms_key_id
    : null
  )
  multi_az           = var.spec.is_multi_az
  ca_cert_identifier = (
    var.spec.ca_cert_identifier != ""
    ? var.spec.ca_cert_identifier
    : null
  )
  license_model = (
    var.spec.license_model != ""
    ? var.spec.license_model
    : null
  )
  storage_type = (
    var.spec.storage_type != ""
    ? var.spec.storage_type
    : null
  )
  iops = (
    var.spec.iops != 0
    ? var.spec.iops
    : null
  )
  publicly_accessible = var.spec.is_publicly_accessible

  ###########################################################################
  # Snapshot logic (create from snapshot if provided)
  ###########################################################################
  snapshot_identifier = (
    var.spec.snapshot_identifier != ""
    ? var.spec.snapshot_identifier
    : null
  )

  ###########################################################################
  # Version upgrades and Maintenance
  ###########################################################################
  allow_major_version_upgrade = var.spec.allow_major_version_upgrade
  auto_minor_version_upgrade  = var.spec.auto_minor_version_upgrade
  apply_immediately           = var.spec.apply_immediately
  maintenance_window          = (
    var.spec.maintenance_window != ""
    ? var.spec.maintenance_window
    : null
  )
  copy_tags_to_snapshot = var.spec.copy_tags_to_snapshot

  ###########################################################################
  # Backup configuration
  ###########################################################################
  backup_retention_period = (
    var.spec.backup_retention_period != 0
    ? var.spec.backup_retention_period
    : null
  )
  backup_window = (
    var.spec.backup_window != ""
    ? var.spec.backup_window
    : null
  )

  ###########################################################################
  # Final snapshot
  ###########################################################################
  deletion_protection       = var.spec.deletion_protection
  skip_final_snapshot       = var.spec.skip_final_snapshot
  final_snapshot_identifier = (
    var.spec.skip_final_snapshot
    ? null
    : local.rds_final_snapshot_name
  )

  ###########################################################################
  # Timezone & IAM DB Auth
  ###########################################################################
  timezone = (
    var.spec.timezone != ""
    ? var.spec.timezone
    : null
  )
  iam_database_authentication_enabled = var.spec.iam_database_authentication_enabled

  ###########################################################################
  # Export DB logs to CloudWatch if specified
  ###########################################################################
  enabled_cloudwatch_logs_exports = (
    length(var.spec.enabled_cloudwatch_logs_exports) > 0
    ? var.spec.enabled_cloudwatch_logs_exports
    : null
  )

  ###########################################################################
  # Master user credentials & secret management
  ###########################################################################
  username = (
    var.spec.replicate_source_db == ""
    ? var.spec.username
    : null
  )

  manage_master_user_password = (
    (
    var.spec.manage_master_user_password
    && var.spec.replicate_source_db == ""
    )
    ? true
    : null
  )

  master_user_secret_kms_key_id = (
    (
    var.spec.manage_master_user_password
    && var.spec.master_user_secret_kms_key_id != ""
    && var.spec.replicate_source_db == ""
    )
    ? var.spec.master_user_secret_kms_key_id
    : null
  )

  password = (
    (
    !var.spec.manage_master_user_password
    && var.spec.replicate_source_db == ""
    )
    ? var.spec.password
    : null
  )

  ###########################################################################
  # DB Subnet Group (either user-provided or newly created)
  ###########################################################################
  db_subnet_group_name = (
    local.final_db_subnet_group_name != ""
    ? local.final_db_subnet_group_name
    : null
  )

  ###########################################################################
  # Parameter Group (either user-provided or newly created)
  ###########################################################################
  parameter_group_name = (
    local.final_parameter_group_name != ""
    ? local.final_parameter_group_name
    : null
  )

  ###########################################################################
  # Option Group (either user-provided or newly created)
  ###########################################################################
  option_group_name = (
    local.final_option_group_name != ""
    ? local.final_option_group_name
    : null
  )

  ###########################################################################
  # Availability Zone if not Multi-AZ
  ###########################################################################
  availability_zone = (
    (!var.spec.is_multi_az && var.spec.availability_zone != "")
    ? var.spec.availability_zone
    : null
  )

  ###########################################################################
  # GP3 storage throughput
  ###########################################################################
  storage_throughput = (
    (
    var.spec.storage_type == "gp3"
    && var.spec.storage_throughput != 0
    )
    ? var.spec.storage_throughput
    : null
  )

  ###########################################################################
  # Performance Insights
  ###########################################################################
  # If performance_insights is null, try() returns fallback values so no error occurs
  performance_insights_enabled = (
    (
    var.spec.performance_insights != null
    && try(var.spec.performance_insights.is_enabled, false)
    )
    ? true
    : false
  )

  performance_insights_kms_key_id = (
    (
    var.spec.performance_insights != null
    && try(var.spec.performance_insights.is_enabled, false)
    && try(var.spec.performance_insights.kms_key_id, "") != ""
    )
    ? var.spec.performance_insights.kms_key_id
    : null
  )

  performance_insights_retention_period = (
    (
    var.spec.performance_insights != null
    && try(var.spec.performance_insights.is_enabled, false)
    && try(var.spec.performance_insights.retention_period, 0) != 0
    )
    ? var.spec.performance_insights.retention_period
    : null
  )

  ###########################################################################
  # Enhanced Monitoring
  ###########################################################################
  monitoring_interval = (
    (
    var.spec.monitoring != null
    && try(var.spec.monitoring.monitoring_interval, 0) != 0
    )
    ? var.spec.monitoring.monitoring_interval
    : null
  )

  monitoring_role_arn = (
    (
    var.spec.monitoring != null
    && try(var.spec.monitoring.monitoring_role_arn, "") != ""
    )
    ? var.spec.monitoring.monitoring_role_arn
    : null
  )

  ###########################################################################
  # Point-in-time restore block (only if no snapshot_identifier is given)
  ###########################################################################
  dynamic "restore_to_point_in_time" {
    for_each = (
      (
      var.spec.snapshot_identifier == ""
      && var.spec.restore_to_point_in_time != null
      )
      ? [1]
      : []
    )

    content {
      # If user requested use_latest_restorable_time = true,
      # override restore_time to null (they cannot coexist)
      restore_time = (
        var.spec.restore_to_point_in_time.use_latest_restorable_time
        ? null
        : (
        var.spec.restore_to_point_in_time.restore_time != ""
        ? var.spec.restore_to_point_in_time.restore_time
        : null
      )
      )

      # If user supplied a restore_time, automatically set use_latest_restorable_time=false
      use_latest_restorable_time = (
        var.spec.restore_to_point_in_time.restore_time != ""
        ? false
        : try(var.spec.restore_to_point_in_time.use_latest_restorable_time, false)
      )

      source_db_instance_automated_backups_arn = try(
        var.spec.restore_to_point_in_time.source_db_instance_automated_backups_arn,
        null
      )
      source_db_instance_identifier = try(
        var.spec.restore_to_point_in_time.source_db_instance_identifier,
        null
      )
      source_dbi_resource_id = try(
        var.spec.restore_to_point_in_time.source_dbi_resource_id,
        null
      )
    }
  }

  ###########################################################################
  # Attach security group references
  ###########################################################################
  vpc_security_group_ids = concat(
    var.spec.associate_security_group_ids,
    [aws_security_group.default.id]
  )

  # Final tags
  tags = local.final_labels
}
