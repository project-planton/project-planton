resource "aws_rds_cluster_instance" "this" {
  count = var.spec.cluster_size

  identifier                   = "${local.resource_id}-${count.index + 1}"
  cluster_identifier           = aws_rds_cluster.this[0].id
  db_subnet_group_name         = aws_rds_cluster.this[0].db_subnet_group_name
  publicly_accessible          = var.spec.is_publicly_accessible
  tags                         = local.final_labels
  engine                       = aws_rds_cluster.this[0].engine
  engine_version               = aws_rds_cluster.this[0].engine_version
  auto_minor_version_upgrade   = true
  apply_immediately            = true
  preferred_maintenance_window = var.spec.maintenance_window
  preferred_backup_window      = var.spec.backup_window
  copy_tags_to_snapshot        = false
  ca_cert_identifier = var.spec.ca_cert_identifier

  # If serverless v2 scaling is defined, we must use "db.serverless"
  instance_class = var.spec.serverlessv2_scaling_configuration != null ? "db.serverless" : var.spec.instance_type

  # Enhanced monitoring
  monitoring_role_arn = var.spec.enhanced_monitoring_role_enabled ? aws_iam_role.enhanced_monitoring[0].arn : null
  monitoring_interval = var.spec.enhanced_monitoring_role_enabled ? var.spec.rds_monitoring_interval : 0

  # Performance Insights
  performance_insights_enabled    = var.spec.is_performance_insights_enabled
  performance_insights_kms_key_id = var.spec.is_performance_insights_enabled ? var.spec.performance_insights_kms_key_id : null

  lifecycle {
    # Match the Pulumi behavior of ignoring engine_version updates
    ignore_changes = [engine_version]
  }

  depends_on = [
    aws_rds_cluster.this
  ]
}
