variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(object({
      name = optional(string),
      id = optional(string),
    })),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}

variable "spec" {
  description = "RDS instance specification."
  type = object({

    db_name = optional(string, "")
    manage_master_user_password = optional(bool, false)
    master_user_secret_kms_key_id = optional(string, "")
    username = optional(string, "")
    password = optional(string, "")
    port = optional(number, 0)
    engine = optional(string, "")
    engine_version = optional(string, "")
    major_engine_version = optional(string, "")
    character_set_name = optional(string, "")
    instance_class = optional(string, "")
    allocated_storage = optional(number, 0)
    max_allocated_storage = optional(number, 0)
    storage_encrypted = optional(bool, false)
    kms_key_id = optional(string, "")
    security_group_ids = optional(list(string), [])
    allowed_cidr_blocks = optional(list(string), [])
    associate_security_group_ids = optional(list(string), [])
    subnet_ids = optional(list(string), [])
    availability_zone = optional(string, "")
    db_subnet_group_name = optional(string, "")
    ca_cert_identifier = optional(string, "")
    parameter_group_name = optional(string, "")
    db_parameter_group = optional(string, "")
    parameters = optional(
      list(
        object({
          apply_method = optional(string, "")
          name = optional(string, "")
          value = optional(string, "")
        })
      ),
      []
    )
    option_group_name = optional(string, "")
    options = optional(
      list(
        object({
          db_security_group_memberships = optional(list(string), [])
          option_name = optional(string, "")
          port = optional(number, 0)
          version = optional(string, "")
          vpc_security_group_memberships = optional(list(string), [])
          option_settings = optional(
            list(
              object({
                name = optional(string, "")
                value = optional(string, "")
              })
            ),
            []
          )
        })
      ),
      []
    )

    is_multi_az = optional(bool, false)
    storage_type = optional(string, "")
    iops = optional(number, 0)
    storage_throughput = optional(number, 0)
    is_publicly_accessible = optional(bool, false)
    snapshot_identifier = optional(string, "")
    allow_major_version_upgrade = optional(bool, false)
    auto_minor_version_upgrade = optional(bool, false)
    apply_immediately = optional(bool, false)
    maintenance_window = optional(string, "")
    skip_final_snapshot = optional(bool, false)
    copy_tags_to_snapshot = optional(bool, false)
    backup_retention_period = optional(number, 0)
    backup_window = optional(string, "")
    deletion_protection = optional(bool, false)
    replicate_source_db = optional(string, "")
    timezone = optional(string, "")
    iam_database_authentication_enabled = optional(bool, false)
    enabled_cloudwatch_logs_exports = optional(list(string), [])

    performance_insights = optional(
      object({
        is_enabled = optional(bool, false)
        kms_key_id = optional(string, "")
        retention_period = optional(number, 0)
      })
    )

    monitoring = optional(
      object({
        monitoring_interval = optional(number, 0)
        monitoring_role_arn = optional(string, "")
      })
    )

    restore_to_point_in_time = optional(
      object({
        restore_time = optional(string, "")
        source_db_instance_automated_backups_arn = optional(string, "")
        source_db_instance_identifier = optional(string, "")
        source_dbi_resource_id = optional(string, "")
        use_latest_restorable_time = optional(bool, false)
      })
    )

    vpc_id = optional(string, "")
    license_model = optional(string, "")
  })
}
