# Basic Example

This example demonstrates a basic setup of a Snowflake database with minimal configuration.

```yaml
apiVersion: snowflake.project.planton/v1
kind: SnowflakeDatabase
metadata:
  name: analytics-db
spec:
  snowflake_credential_id: snowflake-cred-123
  catalog: default_catalog
  comment: "Analytics database for reporting"
  data_retention_time_in_days: 30
  default_ddl_collation: "en_US"
  drop_public_schema_on_creation: false
  enable_console_output: true
  external_volume: "external_vol_1"
  is_transient: false
  log_level: "INFO"
  max_data_extension_time_in_days: 10
  name: analytics_db
  quoted_identifiers_ignore_case: true
  replace_invalid_characters: false
  storage_serialization_policy: "COMPATIBLE"
  suspend_task_after_num_failures: 3
  task_auto_retry_attempts: 2
  trace_level: "OFF"
```

# Example with Advanced Configuration

This example includes advanced configurations such as environment isolation and detailed security settings.

```yaml
apiVersion: snowflake.project.planton/v1
kind: SnowflakeDatabase
metadata:
  name: finance-db
spec:
  snowflake_credential_id: snowflake-cred-finance
  catalog: finance_catalog
  comment: "Finance database for transactional data"
  data_retention_time_in_days: 90
  default_ddl_collation: "en_US"
  drop_public_schema_on_creation: true
  enable_console_output: false
  external_volume: "external_vol_finance"
  is_transient: true
  log_level: "DEBUG"
  max_data_extension_time_in_days: 15
  name: finance_db
  quoted_identifiers_ignore_case: false
  replace_invalid_characters: true
  storage_serialization_policy: "OPTIMIZED"
  suspend_task_after_num_failures: 5
  task_auto_retry_attempts: 3
  trace_level: "ALWAYS ON"
```
