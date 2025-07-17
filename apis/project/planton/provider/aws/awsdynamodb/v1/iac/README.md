# AWS DynamoDB (aws_dynamodb)

Terraform module that provisions an Amazon DynamoDB table together with all the frequently-used options such as billing modes, secondary indexes, TTL, Streams, point-in-time recovery and server-side encryption.

---

## Usage

```hcl
module "orders_ddb" {
  source = "<registry-or-git-url>"

  table_name = "orders"

  # Attribute definitions referenced by either the table or any index
  attribute_definitions = [
    {
      attribute_name = "pk"
      attribute_type = "STRING" # (S)
    },
    {
      attribute_name = "sk"
      attribute_type = "STRING" # (S)
    },
    {
      attribute_name = "gs1pk"
      attribute_type = "STRING"
    }
  ]

  # Primary (partition/sort) key schema
  key_schema = [
    {
      attribute_name = "pk"
      key_type       = "HASH"
    },
    {
      attribute_name = "sk"
      key_type       = "RANGE"
    }
  ]

  # On-demand (PAY_PER_REQUEST) billing
  billing_mode = "PAY_PER_REQUEST"

  # Global secondary index (on-demand inherits billing mode)
  global_secondary_indexes = [
    {
      index_name = "gs1"

      key_schema = [
        {
          attribute_name = "gs1pk"
          key_type       = "HASH"
        },
        {
          attribute_name = "sk"
          key_type       = "RANGE"
        }
      ]

      projection = {
        projection_type    = "INCLUDE"
        non_key_attributes = ["status", "total"]
      }
    }
  ]

  # TTL
  ttl_specification = {
    ttl_enabled    = true
    attribute_name = "ttl"
  }

  # Enable Streams (NEW_AND_OLD_IMAGES)
  stream_specification = {
    stream_enabled    = true
    stream_view_type  = "NEW_AND_OLD_IMAGES"
  }

  point_in_time_recovery_enabled = true

  tags = {
    Environment = "prod"
    Service     = "orders"
  }
}
```

See the [examples](./examples/) directory for more elaborate scenarios, including **PROVISIONED** billing with per-index capacity.

---

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.3 |
| aws | >= 5.0 |

---

## Providers

| Name | Version |
|------|---------|
| aws | >= 5.0 |

---

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| `table_name` | Name of the DynamoDB table. Must be 3-255 characters long. | `string` | n/a | **yes** |
| `attribute_definitions` | List of attribute definitions used by the table or any index.<br>`attribute_name` (string)<br>`attribute_type` (STRING \| NUMBER \| BINARY) | `list(object({ attribute_name = string, attribute_type = string }))` | n/a | **yes** |
| `key_schema` | Primary key schema. One or two elements.<br>`attribute_name` (string)<br>`key_type` (HASH \| RANGE) | `list(object({ attribute_name = string, key_type = string }))` | n/a | **yes** |
| `billing_mode` | Billing mode: `PROVISIONED` or `PAY_PER_REQUEST`. | `string` | `"PAY_PER_REQUEST"` | no |
| `provisioned_throughput` | Provisioned capacity for the table when `billing_mode` = `PROVISIONED`.<br>`read_capacity_units` & `write_capacity_units` > 0 | `object({ read_capacity_units = number, write_capacity_units = number })` | `null` | conditional |
| `global_secondary_indexes` | List of GSI objects.<br>Each object:<br>`index_name`, `key_schema`, `projection`, `provisioned_throughput` (same shape as table). | `list(any)` | `[]` | no |
| `local_secondary_indexes` | List of LSI objects (similar to GSI but must share HASH key with the table). | `list(any)` | `[]` | no |
| `stream_specification` | Streams configuration.<br>`stream_enabled` (bool)<br>`stream_view_type` (NEW_IMAGE \| OLD_IMAGE \| NEW_AND_OLD_IMAGES \| STREAM_KEYS_ONLY) | `object({ stream_enabled = bool, stream_view_type = string })` | `{ stream_enabled = false }` | no |
| `ttl_specification` | TTL configuration.<br>`ttl_enabled` (bool)<br>`attribute_name` (string) | `object({ ttl_enabled = bool, attribute_name = string })` | `{ ttl_enabled = false }` | no |
| `sse_specification` | Server-side encryption settings.<br>`enabled` (bool)<br>`sse_type` (AES256 \| KMS)<br>`kms_master_key_id` (string, only for KMS) | `object({ enabled = bool, sse_type = string, kms_master_key_id = string })` | `{ enabled = false }` | no |
| `point_in_time_recovery_enabled` | Enable point-in-time recovery (continuous backup). | `bool` | `false` | no |
| `tags` | Map of tags applied to the table. | `map(string)` | `{}` | no |

---

## Outputs

| Name | Description |
|------|-------------|
| `table_arn` | Fully-qualified ARN of the table. |
| `table_name` | Runtime name of the table (may include suffixes). |
| `table_id` | AWS-assigned unique identifier. |
| `stream` | Object with `stream_arn` and `stream_label` when Streams are enabled. |
| `kms_key_arn` | ARN of the customer-managed CMK used for encryption (when SSE = KMS). |
| `global_secondary_index_names` | Names of the provisioned GSIs. |
| `local_secondary_index_names` | Names of the provisioned LSIs. |

---

## Development

Run the test suite:

```bash
make test
```

Format and validate:

```bash
terraform fmt -recursive
terraform validate
```

---

## License

Apache License 2.0 © 2024 Project Planton
