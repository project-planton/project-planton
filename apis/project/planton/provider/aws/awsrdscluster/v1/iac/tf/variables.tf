variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(string),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "The specification for the AWS RDS Cluster."
  type = object({

    # The engine to use; e.g. aurora-mysql, aurora-postgresql, mysql, postgres.
    engine = optional(string, "")

    # Engine version (e.g. 5.7.mysql_aurora.2.03.2)
    engine_version = optional(string, "")

    # Engine mode: global, parallelquery, provisioned, serverless (default "provisioned")
    engine_mode = optional(string, "provisioned")

    # Cluster parameter group family
    cluster_family = optional(string, "")

    # Instance class (e.g. db.r5.large). For Serverless v2, use "db.serverless"
    instance_type = optional(string, "")

    # Number of instances in the cluster
    cluster_size = optional(number, 0)

    # Manage master user password in Secrets Manager
    manage_master_user_password = optional(bool, false)

    # The KMS key ID for master user secret, if manage_master_user_password = true
    master_user_secret_kms_key_id = optional(string, "")

    # Master username
    master_user = optional(string, "master")

    # Master password (ignored if manage_master_user_password = true)
    master_password = optional(string, "")

    # Database name
    database_name = optional(string, "")

    # Publicly accessible?
    is_publicly_accessible = optional(bool, false)

    # Database port
    database_port = optional(number, 0)

    # The VPC ID where this cluster should be placed
    vpc_id = optional(string, "")

    # Subnet IDs (if using a custom DB subnet group)
    subnet_ids = optional(list(string), [])

    # Existing DB subnet group name (alternative to subnet_ids)
    db_subnet_group_name = optional(string, "")

    # Security groups from which to allow ingress
    security_group_ids = optional(list(string), [])

    # Whether to allow traffic among resources within the same security group
    intra_security_group_traffic_enabled = optional(bool, false)

    # List of CIDR blocks allowed to connect
    allowed_cidr_blocks = optional(list(string), [])

    # Additional SGs to associate with the DB
    associate_security_group_ids = optional(list(string), [])

    # IAM DB Authentication
    iam_database_authentication_enabled = optional(bool, false)

    # Whether the cluster is encrypted at rest
    storage_encrypted = optional(bool, false)

    # KMS key ARN for encryption at rest
    storage_kms_key_arn = optional(string, "")

    # Enable Performance Insights
    is_performance_insights_enabled = optional(bool, false)

    # KMS key for Performance Insights
    performance_insights_kms_key_id = optional(string, "")

    # Weekly maintenance window
    maintenance_window = optional(string, "")

    # CloudWatch logs exports
    enabled_cloudwatch_logs_exports = optional(list(string), [])

    # Whether to create the Enhanced Monitoring IAM role
    enhanced_monitoring_role_enabled = optional(bool, false)

    # Attributes used in naming the Enhanced Monitoring IAM role
    enhanced_monitoring_attributes = optional(list(string), [])

    # Monitoring interval if Enhanced Monitoring is enabled
    rds_monitoring_interval = optional(number, 0)

    # Skip final snapshot on deletion
    skip_final_snapshot = optional(bool, false)

    # Deletion protection
    deletion_protection = optional(bool, false)

    # Use an existing snapshot to create the cluster
    snapshot_identifier = optional(string, "")

    # Allow major version upgrade
    allow_major_version_upgrade = optional(bool, false)

    # CA certificate identifier
    ca_cert_identifier = optional(string, "")

    # Retention period for backups (days)
    retention_period = optional(number, 0)

    # Daily backup window
    backup_window = optional(string, "")

    # Auto-scaling configuration
    auto_scaling = optional(object({
      is_enabled = optional(bool, false)
      policy_type = optional(string, "")
      target_metrics = optional(string, "")
      target_value = optional(number, 0)
      scale_in_cooldown = optional(number, 0)
      scale_out_cooldown = optional(number, 0)
      min_capacity = optional(number, 0)
      max_capacity = optional(number, 0)
    }), {
      is_enabled         = false,
      policy_type        = "",
      target_metrics     = "",
      target_value       = 0,
      scale_in_cooldown  = 0,
      scale_out_cooldown = 0,
      min_capacity       = 0,
      max_capacity       = 0
    })

    # Serverless v1 scaling configuration
    scaling_configuration = optional(object({
      auto_pause = optional(bool, false)
      max_capacity = optional(number, 0)
      min_capacity = optional(number, 0)
      seconds_until_auto_pause = optional(number, 0)
      timeout_action = optional(string, "")
    }), null)

    # Serverless v2 scaling configuration
    serverlessv2_scaling_configuration = optional(object({
      max_capacity = optional(number, 0)
      min_capacity = optional(number, 0)
    }), null)

    # DB cluster parameter group name
    cluster_parameter_group_name = optional(string, "")

    # List of cluster parameters
    cluster_parameters = optional(list(object({
      apply_method = optional(string, "")
      name = optional(string, "")
      value = optional(string, "")
    })), [])
  })
}
