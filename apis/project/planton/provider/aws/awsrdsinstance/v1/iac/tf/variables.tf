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
    labels = optional(object({
      key = string, value = string
    })),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}


variable "spec" {
  description = "spec"
  type = object({

    # The name of the database to create when the DB instance is created.
    # If this parameter is not specified, no database is created in the DB instance.
    # Note that this does not apply for Oracle or SQL Server engines.
    # See the [AWS documentation](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/rds/create-db-instance.html)
    # for more details on what applies for those engines. \
    # If you are providing an Oracle db name, it needs to be in all upper case.
    # Cannot be specified for a replica.
    db_name = string

    # Set to true to allow RDS to manage the master user password in Secrets Manager. Cannot be set if `password` is provided.
    manage_master_user_password = bool

    # The Amazon Web Services KMS key identifier is the key ARN, key ID, alias ARN, or alias name for the KMS key.
    # To use a KMS key in a different Amazon Web Services account, specify the key ARN or alias ARN.
    # If not specified, the default KMS key for your Amazon Web Services account is used.
    master_user_secret_kms_key_id = string

    # (Required unless a `snapshotIdentifier` or `replicateSourceDb` is provided)
    # Username for the master DB user. Cannot be specified for a replica.
    username = string

    # (Required unless `manageMasterUserPassword` is set to true or unless a `snapshotIdentifier` or `replicateSourceDb`
    # is provided or `manageMasterUserPassword` is set.) Password for the master DB user. Note that this may show up in
    # logs, and it will be stored in the state file. Cannot be set if `manageMasterUserPassword` is set to `true`.
    password = string

    # The port on which the DB accepts connections.
    port = number

    # The database engine to use. For supported values, see the Engine parameter in [API action CreateDBInstance]
    # (https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).
    # Note that for Amazon Aurora instances the engine must match the DB cluster's engine'.
    # For information on the difference between the available Aurora MySQL engines see
    # [Comparison between Aurora MySQL 1 and Aurora MySQL 2](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/AuroraMySQL.Updates.20180206.html)
    # in the Amazon RDS User Guide.
    engine = string

    # The engine version to use. If `autoMinorVersionUpgrade` is enabled, you can provide a prefix of the version such
    # as `8.0` (for `8.0.36`). The actual engine version used is returned in the attribute `engineVersionActual`,
    # see Attribute Reference below. For supported values, see the EngineVersion parameter in
    # [API action CreateDBInstance](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBInstance.html).
    # Note that for Amazon Aurora instances the engine version must match the DB cluster's engine version'.
    engine_version = string

    # Database MAJOR engine version, depends on engine type
    # https://docs.aws.amazon.com/cli/latest/reference/rds/create-option-group.html
    major_engine_version = string

    # The character set name to use for DB encoding in Oracle and Microsoft SQL instances (collation).
    # This can't be changed.
    # See [Oracle Character Sets Supported in Amazon RDS](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.OracleCharacterSets.html) or
    # [Server-Level Collation for Microsoft SQL Server](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Appendix.SQLServer.CommonDBATasks.Collation.html) for more information.
    # Cannot be set  with `replicateSourceDb`, `restoreToPointInTime`, `s3Import`, or `snapshotIdentifier`.
    character_set_name = string

    # The instance type of the RDS instance.
    instance_class = string

    # The allocated storage in gibibytes. If `maxAllocatedStorage` is configured, this argument represents the initial
    # storage allocation and differences from the configuration will be ignored automatically when Storage Autoscaling
    # occurs. If `replicateSourceDb` is set, the value is ignored during the creation of the instance.
    allocated_storage = number

    # When configured, the upper limit to which Amazon RDS can automatically scale the storage of the DB instance.
    # Configuring this will automatically ignore differences to `allocatedStorage`. Must be greater than or equal to
    # `allocatedStorage` or `0` to disable Storage Autoscaling.
    max_allocated_storage = number

    # Specifies whether the DB instance is
    # encrypted. Note that if you are creating a cross-region read replica this field
    # is ignored and you should instead declare `kmsKeyId` with a valid ARN. The
    # default is `false` if not specified.
    storage_encrypted = bool

    # The ARN for the KMS encryption key. If creating an
    # encrypted replica, set this to the destination KMS ARN.
    kms_key_id = string

    # The IDs of the security groups from which to allow `ingress` traffic to the DB instance
    security_group_ids = list(string)

    # The whitelisted CIDRs which to allow `ingress` traffic to the DB instance
    allowed_cidr_blocks = list(string)

    # The IDs of the existing security groups to associate with the DB instance
    associate_security_group_ids = list(string)

    # List of subnet IDs for the DB. DB instance will be created in the VPC associated with the DB subnet group provisioned using the subnet IDs.
    # Specify one of `subnet_ids`, `db_subnet_group_name` or `availability_zone`
    subnet_ids = list(string)

    # The AZ for the RDS instance. Specify one of `subnet_ids`, `db_subnet_group_name` or `availability_zone`.
    # If `availability_zone` is provided, the instance will be placed into the default VPC or EC2 Classic
    availability_zone = string

    # Name of DB subnet group. DB instance will be created in the VPC associated with the DB subnet group.
    # Specify one of `subnet_ids`, `db_subnet_group_name` or `availability_zone`
    db_subnet_group_name = string

    # The identifier of the CA certificate for the DB instance.
    ca_cert_identifier = string

    # Name of the DB parameter group to associate.
    parameter_group_name = string

    # The DB parameter group family name. The value depends on DB engine used.
    # See [DBParameterGroupFamily](https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_CreateDBParameterGroup.html#API_CreateDBParameterGroup_RequestParameters)
    # for instructions on how to retrieve applicable value.
    db_parameter_group = string

    # A list of DB parameters to apply. Note that parameters may differ from a DB family to another
    parameters = list(object({

      # "immediate" (default), or "pending-reboot". Some
      # engines can't apply some parameters without a reboot, and you will need to
      # specify "pending-reboot" here.
      apply_method = string

      # The name of the DB parameter.
      name = string

      # The value of the DB parameter.
      value = string
    }))

    # Name of the DB option group to associate
    option_group_name = string

    # A list of DB options to apply with an option group. Depends on DB engine
    options = list(object({

      # List of DB Security Groups for which the option is enabled.
      db_security_group_memberships = list(string)

      # Name of the option (e.g., MEMCACHED).
      option_name = string

      # Port number when connecting to the option (e.g., 11211). Leaving out or removing `port` from your
      # configuration does not remove or clear a port from the option in AWS. AWS may assign a default port.
      # Not including `port` in your configuration means that the AWS provider will ignore a previously set value,
      # a value set by AWS, and any port changes.
      port = number

      # Version of the option (e.g., 13.1.0.0). Leaving out or removing `version` from your configuration does not
      # remove or clear a version from the option in AWS. AWS may assign a default version. Not including `version`
      # in your configuration means that the AWS provider will ignore a previously set value, a value set by AWS,
      # and any version changes.
      version = string

      # List of VPC Security Groups for which the option is enabled.
      vpc_security_group_memberships = list(string)

      # The option settings to apply.
      option_settings = list(object({

        # Name of the setting.
        name = string

        # Value of the setting.
        value = string
      }))
    }))

    # Specifies if the RDS instance is multi-AZ
    is_multi_az = bool

    # One of "standard" (magnetic), "gp2" (general
    # purpose SSD), "gp3" (general purpose SSD that needs `iops` independently)
    # or "io1" (provisioned IOPS SSD). The default is "io1" if `iops` is specified,
    # "gp2" if not.
    storage_type = string

    # The amount of provisioned IOPS. Setting this implies a storageType of "io1".
    # Can only be set when `storageType` is `"io1"` or `"gp3"`.
    # Cannot be specified for gp3 storage if the `allocatedStorage` value is below a per-`engine` threshold.
    # See the [RDS User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_Storage.html#gp3-storage) for details.
    iops = number

    # The storage throughput value for the DB instance. Can only be set when `storage_type` is `gp3`.
    # Cannot be specified if the `allocated_storage` value is below a per-engine threshold.
    storage_throughput = number

    # Bool to control if instance is publicly accessible. Default is `false`.
    is_publicly_accessible = bool

    # Snapshot identifier e.g: `rds:production-2019-06-26-06-05` for automated or `manual-backup-2023-11-16` for manual.
    # If specified, the module create the instance from the snapshot.
    snapshot_identifier = string

    # Allow major version upgrade
    allow_major_version_upgrade = bool

    # Allow automated minor version upgrade (e.g. from Postgres 9.5.3 to Postgres 9.5.4)
    auto_minor_version_upgrade = bool

    # Specifies whether any database modifications are applied immediately, or during the next maintenance window
    apply_immediately = bool

    # The window to perform maintenance in. Syntax: 'ddd:hh24:mi-ddd:hh24:mi' UTC
    maintenance_window = string

    # If true (default), no snapshot will be made before deleting DB
    skip_final_snapshot = bool

    # Copy tags from DB to a snapshot
    copy_tags_to_snapshot = bool

    # Backup retention period in days. Must be > 0 to enable backups
    backup_retention_period = number

    # When AWS can perform DB snapshots, can't overlap with maintenance window
    backup_window = string

    # Set to true to enable deletion protection on the RDS instance
    deletion_protection = bool

    # Specifies that this resource is a Replicate database, and to use this value as the source database.
    # This correlates to the `identifier` of another Amazon RDS Database to replicate (if replicating within a single region)
    # or ARN of the Amazon RDS Database to replicate (if replicating cross-region).
    # Note that if you are creating a cross-region replica of an encrypted database you will also need to
    # specify a `kms_key_id`. See [DB Instance Replication](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Overview.Replication.html)
    # and [Working with PostgreSQL and MySQL Read Replicas](https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ReadRepl.html)
    # for more information on using Replication.
    replicate_source_db = string

    # Time zone of the DB instance. timezone is currently only supported by Microsoft SQL Server. The timezone can only
    # be set on creation. See [MSSQL User Guide](http://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_SQLServer.html#SQLServer.Concepts.General.TimeZone)
    # for more information.
    timezone = string

    # Specifies whether or mappings of AWS Identity and Access Management (IAM) accounts to database accounts is enabled
    iam_database_authentication_enabled = bool

    # List of log types to enable for exporting to CloudWatch logs. If omitted, no logs will be exported.
    # Valid values (depending on engine): alert, audit, error, general, listener, slowquery, trace, postgresql (PostgreSQL),
    # upgrade (PostgreSQL).
    enabled_cloudwatch_logs_exports = list(string)

    # performance insights settings
    performance_insights = object({

      # Specifies whether Performance Insights are enabled.
      is_enabled = bool

      # The ARN for the KMS key to encrypt Performance Insights data. Once KMS key is set, it can never be changed.
      kms_key_id = string

      # The amount of time in days to retain Performance Insights data. Either 7 (7 days) or 731 (2 years).
      retention_period = number
    })

    # enhanced monitoring settings
    monitoring = object({

      # The interval, in seconds, between points when Enhanced Monitoring metrics are collected for the DB instance.
      # To disable collecting Enhanced Monitoring metrics, specify 0. Valid Values are 0, 1, 5, 10, 15, 30, 60.
      monitoring_interval = number

      # The ARN for the IAM role that permits RDS to send enhanced monitoring metrics to CloudWatch Logs
      monitoring_role_arn = string
    })

    # An object specifying the restore point in time for the DB instance to restore from. Only used when
    # `snapshot_identifier` is not provided.
    restore_to_point_in_time = object({

      # The date and time to restore from. Value must be a time in Universal Coordinated Time (UTC) format and must be
      # before the latest restorable time for the DB instance. Cannot be specified with `useLatestRestorableTime`.
      restore_time = string

      # The ARN of the automated backup from which to restore.
      # Required if `sourceDbInstanceIdentifier` or `sourceDbiResourceId` is not specified.
      source_db_instance_automated_backups_arn = string

      # The identifier of the source DB instance from which to restore. Must match the identifier of an existing DB instance.
      # Required if `sourceDbInstanceAutomatedBackupsArn` or `sourceDbiResourceId` is not specified.
      source_db_instance_identifier = string

      # The resource ID of the source DB instance from which to restore.
      # Required if `sourceDbInstanceIdentifier` or `sourceDbInstanceAutomatedBackupsArn` is not specified.
      source_dbi_resource_id = string

      # A boolean value that indicates whether the DB instance is restored from the latest backup time.
      # Defaults to `false`. Cannot be specified with `restoreTime`.
      use_latest_restorable_time = bool
    })

    # VPC ID the DB instance will be created in
    vpc_id = string

    # License model for this DB. Optional, but required for some DB Engines.
    # Valid values: license-included | bring-your-own-license | general-public-license
    license_model = string
  })
}