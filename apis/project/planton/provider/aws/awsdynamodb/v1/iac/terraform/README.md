# AWS DynamoDB Terraform Module

This module provisions an **Amazon DynamoDB table** together with all commonly-used features such as on-demand / provisioned capacity, global & local secondary indexes, DynamoDB Streams, TTL, point-in-time recovery (PITR), server-side encryption (SSE) and tagging.

The public interface of the module is intentionally kept very close to the `AwsDynamodbSpec` protobuf message shipped with `project-planton` so that the same schema can be reused consistently across tools (Terraform, Kubernetes CRD, etc.).

---

## Usage

### Quick start (on-demand table)
```hcl
module "orders_table" {
  source = "github.com/project-planton/project-planton//modules/aws_dynamodb"

  table_name = "orders"

  attribute_definitions = [
    { attribute_name = "pk", attribute_type = "S" },
    { attribute_name = "sk", attribute_type = "S" }
  ]

  key_schema = [
    { attribute_name = "pk", key_type = "HASH" },
    { attribute_name = "sk", key_type = "RANGE" }
  ]

  billing_mode = "PAY_PER_REQUEST"

  tags = {
    Environment = "dev"
    Service     = "checkout"
  }
}
```

### Provisioned capacity with a global secondary index & PITR
```hcl
module "users_table" {
  source = "github.com/project-planton/project-planton//modules/aws_dynamodb"

  table_name = "users"

  attribute_definitions = [
    { attribute_name = "user_id",   attribute_type = "S" },
    { attribute_name = "email",     attribute_type = "S" },
    { attribute_name = "created_at", attribute_type = "N" }
  ]

  key_schema = [
    { attribute_name = "user_id", key_type = "HASH" }
  ]

  billing_mode           = "PROVISIONED"
  provisioned_throughput = { read_capacity_units = 10, write_capacity_units = 5 }

  global_secondary_indexes = [
    {
      index_name = "email_index"
      key_schema = [
        { attribute_name = "email", key_type = "HASH" }
      ]
      projection = { projection_type = "ALL" }
      # Using table-level capacity because `billing_mode` is PROVISIONED.
    }
  ]

  point_in_time_recovery_enabled = true

  sse_specification = {
    enabled           = true
    sse_type          = "KMS"
    kms_master_key_id = aws_kms_key.dynamodb.arn
  }

  stream_specification = {
    stream_enabled   = true
    stream_view_type = "NEW_AND_OLD_IMAGES"
  }

  tags = {
    Environment = "prod"
    CostCenter  = "42"
  }
}
```

---

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.3 |
| aws | >= 5.0 |

---

## Inputs

| Variable | Type | Default | Required | Description |
|----------|------|---------|:--------:|-------------|
| `table_name` | `string` | n/a | **yes** | Name of the DynamoDB table (3–255 chars). |
| `attribute_definitions` | `list(object({ attribute_name = string, attribute_type = string }))` | n/a | **yes** | List of attributes referenced by the table or any index. `attribute_type` must be one of `"S"`, `"N"`, `"B"`. |
| `key_schema` | `list(object({ attribute_name = string, key_type = string }))` | n/a | **yes** | Primary-key schema (1–2 elements). `key_type` must be `"HASH"` or `"RANGE"`. |
| `billing_mode` | `string` | `"PAY_PER_REQUEST"` | no | Billing mode. Allowed: `"PROVISIONED"`, `"PAY_PER_REQUEST"`. |
| `provisioned_throughput` | `object({ read_capacity_units = number, write_capacity_units = number })` | `null` | cond.* | Required when `billing_mode = "PROVISIONED"`; must be **unset** when `PAY_PER_REQUEST`. |
| `global_secondary_indexes` | `list(object({ index_name = string, key_schema = list(object({ attribute_name = string, key_type = string })), projection = object({ projection_type = string, non_key_attributes = list(string) }), provisioned_throughput = object({ read_capacity_units = number, write_capacity_units = number }) }))` | `[]` | no | GSI definitions. When `billing_mode = "PROVISIONED"`, each GSI must include its own `provisioned_throughput`; omit otherwise. |
| `local_secondary_indexes` | same shape as `global_secondary_indexes` but without `provisioned_throughput` | `[]` | no | LSI definitions. |
| `stream_specification` | `object({ stream_enabled = bool, stream_view_type = string })` | `{ stream_enabled = false }` | no | Configure DynamoDB Streams. `stream_view_type` must be one of `"NEW_IMAGE"`, `"OLD_IMAGE"`, `"NEW_AND_OLD_IMAGES"`, `"KEYS_ONLY"` when `stream_enabled = true`. |
| `ttl_specification` | `object({ ttl_enabled = bool, attribute_name = string })` | `{ ttl_enabled = false }` | no | Time-to-live (TTL) configuration. `attribute_name` is mandatory when `ttl_enabled = true`. |
| `sse_specification` | `object({ enabled = bool, sse_type = string, kms_master_key_id = string })` | `{ enabled = true, sse_type = "AES256" }` | no | Server-side encryption. If `sse_type = "KMS"` you must supply a CMK ARN. |
| `point_in_time_recovery_enabled` | `bool` | `false` | no | Enables point-in-time recovery (continuous backups). |
| `tags` | `map(string)` | `{}` | no | Key/value tags to apply to the table & related resources. |

*cond.* = Conditional requirement depending on other arguments (see above).

---

## Outputs

| Name | Description |
|------|-------------|
| `table_arn` | Fully-qualified ARN of the table. |
| `table_name` | Final table name (with any runtime suffix). |
| `table_id` | Unique identifier assigned by AWS. |
| `stream` | Object with `stream_arn` and `stream_label` when streams are enabled. |
| `kms_key_arn` | ARN of the customer-managed KMS key when SSE uses CMK. |
| `global_secondary_index_names` | List of GSI names provisioned. |
| `local_secondary_index_names` | List of LSI names provisioned. |

---

## Validation rules & best practices

The module mirrors all **CEL validations** present in the protobuf specification, guaranteeing that every plan is rejected **before** hitting AWS when the configuration is internally inconsistent.

* Billing mode vs. capacity – Provisioned throughput must be configured **only** when `billing_mode = PROVISIONED` and must be omitted for on-demand tables. The same applies to every GSI.
* Projections – `non_key_attributes` are allowed **only** when `projection_type = INCLUDE` and must be provided in that case.
* Streams – `stream_view_type` is mandatory when streams are enabled.
* TTL – `attribute_name` is required when TTL is enabled.
* SSE – `sse_type` must be set when encryption is enabled; `kms_master_key_id` is required only when `sse_type = KMS`.

### Security / compliance

* **Encryption at rest** – Enabled by default (`AES256`). Bring your own CMK by setting `sse_type = "KMS"` and passing `kms_master_key_id`.
* **Encryption in transit** – All SDK calls use TLS by default. No additional action needed.
* **Back-ups & disaster recovery** – Enable either PITR (`point_in_time_recovery_enabled = true`) or consider regular export jobs to S3.
* **Data retention** – Use TTL to automatically remove stale items and reduce storage cost.

---

## Development & contribution

Issues and pull-requests are welcome! Please run `terraform fmt -recursive` and `terraform validate` before opening a PR.

---

## License

Apache 2.0 © Project Planton