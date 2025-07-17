# AWS DynamoDB Terraform Module

A reusable Terraform module that provisions and manages an **Amazon DynamoDB table** with optional secondary indexes, on-demand or provisioned capacity, server-side encryption, streams, TTL and other advanced features.  
The public interface of the module is generated from `AwsDynamodbSpec` and `AwsDynamodbStackOutputs` protobuf contracts (see `./proto`).

---

## Usage

```hcl
module "orders_table" {
  source = "git::https://github.com/project-planton/terraform-aws-dynamodb.git?ref=v1.0.0"

  table_name             = "orders"
  billing_mode           = "PAY_PER_REQUEST"   # or "PROVISIONED"

  attribute_definitions = [
    { attribute_name = "pk", attribute_type = "STRING" },
    { attribute_name = "sk", attribute_type = "STRING" },
    { attribute_name = "gsi1pk", attribute_type = "STRING" },
    { attribute_name = "gsi1sk", attribute_type = "STRING" },
  ]

  key_schema = [
    { attribute_name = "pk", key_type = "HASH" },
    { attribute_name = "sk", key_type = "RANGE" },
  ]

  global_secondary_indexes = [
    {
      index_name = "gsi1"
      key_schema = [
        { attribute_name = "gsi1pk", key_type = "HASH" },
        { attribute_name = "gsi1sk", key_type = "RANGE" },
      ]
      projection = {
        projection_type     = "INCLUDE"
        non_key_attributes  = ["status", "total"]
      }
    }
  ]

  stream_specification = {
    stream_enabled   = true
    stream_view_type = "NEW_AND_OLD_IMAGES"
  }

  ttl_specification = {
    ttl_enabled    = true
    attribute_name = "expires_at"
  }

  tags = {
    environment = "prod"
    service     = "billing"
  }
}
```

Refer to the [examples](./examples/) directory for more complete configurations:

* `on_demand` – minimal PAY_PER_REQUEST table
* `provisioned_with_gsi` – table + GSIs on provisioned capacity
* `encrypted_kms` – customer-managed CMK encryption

---

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| `table_name` | Name of the DynamoDB table (3 – 255 chars). | `string` | n/a | yes |
| `attribute_definitions` | All attributes referenced by the table or any index. | `list(object({ attribute_name = string, attribute_type = string }))` | n/a | yes |
| `key_schema` | Primary key schema (HASH and optional RANGE). | `list(object({ attribute_name = string, key_type = string }))` | n/a | yes |
| `billing_mode` | Billing mode: `PROVISIONED` or `PAY_PER_REQUEST`. | `string` | `PAY_PER_REQUEST` | no |
| `provisioned_throughput` | Table-level RCUs/WCUs (required when `billing_mode = "PROVISIONED"`). | `object({ read_capacity_units = number, write_capacity_units = number })` | `null` | conditional |
| `global_secondary_indexes` | List of global secondary indexes (GSIs). | `list(object({ index_name = string, key_schema = list(object({ attribute_name = string, key_type = string })), projection = object({ projection_type = string, non_key_attributes = list(string) }), provisioned_throughput = object({ read_capacity_units = number, write_capacity_units = number }) }))` | `[]` | no |
| `local_secondary_indexes` | List of local secondary indexes (LSIs). | `list(object({ index_name = string, key_schema = list(object({ attribute_name = string, key_type = string })), projection = object({ projection_type = string, non_key_attributes = list(string) }) }))` | `[]` | no |
| `stream_specification` | DynamoDB Streams configuration. | `object({ stream_enabled = bool, stream_view_type = string })` | `{ stream_enabled = false, stream_view_type = null }` | no |
| `ttl_specification` | TTL (Time-to-Live) configuration. | `object({ ttl_enabled = bool, attribute_name = string })` | `{ ttl_enabled = false, attribute_name = null }` | no |
| `sse_specification` | Server-side encryption settings. | `object({ enabled = bool, sse_type = string, kms_master_key_id = string })` | `{ enabled = false, sse_type = null, kms_master_key_id = null }` | no |
| `point_in_time_recovery_enabled` | Enable point-in-time recovery (continuous backups). | `bool` | `false` | no |
| `tags` | Map of resource tags. | `map(string)` | `{}` | no |

> **Note** – Validation rules from the proto spec (e.g. capacity requirements, SSE constraints, TTL name when enabled, etc.) are enforced by this module via `terraform`-side `validation` blocks as well as by the AWS API at deploy time.

---

## Outputs

| Name | Description |
|------|-------------|
| `table_arn` | Fully-qualified ARN of the table. |
| `table_name` | Name of the table (may include runtime suffixes). |
| `table_id` | AWS-assigned unique identifier of the table. |
| `stream` | Object with `stream_arn` and `stream_label` when Streams are enabled. |
| `kms_key_arn` | ARN of the CMK used for SSE (KMS only). |
| `global_secondary_index_names` | List of GSI names created. |
| `local_secondary_index_names` | List of LSI names created. |

---

## Requirements

| Name | Version |
|------|---------|
| `terraform` | >= 1.3 |
| `aws` provider | >= 5.0 |

The module is tested against AWS provider `>= 5.0` and Terraform `>= 1.3`.  
Support for older versions is *not* guaranteed.

---

## Contributing

Issues, feature requests and pull requests are welcome! Please run `make test` before opening a PR to ensure generated code, docs and examples are up-to-date with the protobuf contracts.

---

## License

Apache 2.0 © Project Planton