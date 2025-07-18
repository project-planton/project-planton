# AWS – DynamoDB table Terraform module

## Overview
This module creates and manages an **Amazon DynamoDB** table with the full feature-set that is exposed in the AWS API today – encryption, streams, TTL, point-in-time recovery, GSIs/LSIs, tagging, different billing modes and much more.  
The whole configuration surface is represented by **one single input variable – `spec` – whose schema is generated from the protobuf message `AwsDynamodbSpec`** (see `proto/` directory). This keeps the module interface predictable and future-proof: when AWS ships new capabilities we only need to update the protobuf and the internal translation layer – your Terraform code stays unchanged.

## Supported features
* Pay-per-request **or** provisioned capacity (including per-GSI overrides)
* Global and local secondary indexes (GSIs / LSIs)
* Streams – any of the four stream view types
* Time-to-live (TTL) for automatic item expiration
* Server-side encryption – `AES256` or customer-managed KMS keys
* Point-in-time recovery (continuous backups)
* Arbitrary key/value **tags**

---

## Quick-start example
```hcl
module "orders_table" {
  source = "github.com/project-planton/project-planton//modules/aws_dynamodb"
  # or "git::https://github.com/project-planton/project-planton.git//modules/aws_dynamodb?ref=v0.1.0"

  # The only required input
  spec = {
    table_name = "orders"

    attribute_definitions = [
      { attribute_name = "pk", attribute_type = "STRING" },
      { attribute_name = "sk", attribute_type = "STRING" }
    ]

    key_schema = [
      { attribute_name = "pk", key_type = "HASH" },
      { attribute_name = "sk", key_type = "RANGE" }
    ]

    billing_mode = "PAY_PER_REQUEST"  # On-demand pricing

    point_in_time_recovery_enabled = true

    tags = {
      Environment = "production"
      Service     = "orders"
    }
  }
}
```

### Provisioned capacity + GSI example
```hcl
module "shopping_cart" {
  source = "github.com/project-planton/project-planton//modules/aws_dynamodb"

  spec = {
    table_name = "shopping_cart"

    attribute_definitions = [
      { attribute_name = "user_id", attribute_type = "STRING" },
      { attribute_name = "cart_id", attribute_type = "STRING" },
      { attribute_name = "status",  attribute_type = "STRING" }
    ]

    key_schema = [
      { attribute_name = "user_id", key_type = "HASH" },
      { attribute_name = "cart_id", key_type = "RANGE" }
    ]

    billing_mode = "PROVISIONED"
    provisioned_throughput = {
      read_capacity_units  = 5
      write_capacity_units = 5
    }

    global_secondary_indexes = [
      {
        index_name = "status-gsi"
        key_schema = [
          { attribute_name = "status",  key_type = "HASH" },
          { attribute_name = "user_id", key_type = "RANGE" }
        ]
        projection = {
          projection_type = "ALL"
        }
        # GSI-level capacity (overrides table-level numbers)
        provisioned_throughput = {
          read_capacity_units  = 10
          write_capacity_units = 2
        }
      }
    ]
  }
}
```

---

## Input variables
| Name | Type | Description | Default | Required |
|------|------|-------------|---------|----------|
| `spec` | `object` (<br>see [`AwsDynamodbSpec`](#spec-schema) below<br>) | Full specification of the table, indexes and all optional settings. Validation is performed by CEL rules that are compiled from the protobuf definition. | n/a | **yes** |

There are **no other inputs** – region, credentials, etc. are supplied through the standard Terraform AWS provider configuration.

### <a id="spec-schema"></a>`spec` schema (simplified)
The table below is a human-friendly rendering of the proto message. Refer to the `.proto` file for the authoritative definition and validation rules.

| Field | Type | Default | Notes |
|-------|------|---------|-------|
| `table_name` | `string` | – | 3–255 chars. Will be post-fixed with a random suffix when `create_before_destroy` is performed. |
| `attribute_definitions` | list(object) | – | All attributes referenced by the main key schema **and** by every index. |
| `key_schema` | list(object) | – | 1–2 elements: **partition (HASH)** key and optional **sort (RANGE)** key. |
| `billing_mode` | `string` | – | `PROVISIONED` or `PAY_PER_REQUEST`. |
| `provisioned_throughput` | object | – | Required when `billing_mode = PROVISIONED`. RCUs / WCUs must be `> 0`. |
| `global_secondary_indexes` | list(object) | `[]` | Each GSI may have its own `provisioned_throughput`. |
| `local_secondary_indexes` | list(object) | `[]` | LSI shares the same HASH key with the base table. |
| `stream_specification` | object | – | Enable **DynamoDB Streams** and choose one `stream_view_type`. |
| `ttl_specification` | object | – | Enable TTL and provide the attribute holding the expiry epoch time. |
| `sse_specification` | object | – | Enable server-side encryption (`AES256` or `KMS`). `kms_master_key_id` is required only for `KMS`. |
| `point_in_time_recovery_enabled` | `bool` | `false` | Enables 35-day point-in-time recovery. |
| `tags` | map(string) | `{}` | Arbitrary key/values applied to the table and propagated to supporting resources (stream, KMS key, etc.). |

> Validation logic such as *“`provisioned_throughput` must be set when `billing_mode` is PROVISIONED”* is enforced automatically by [buf-build/validate](https://github.com/bufbuild/validate) CEL constraints included in the proto file. Your plan will fail fast if any rule is violated.

---

## Outputs
| Name | Description |
|------|-------------|
| `table_arn` | Fully-qualified ARN of the table. |
| `table_name` | Final name of the table (it might include a random suffix in replacement scenarios). |
| `table_id` | AWS-assigned unique identifier of the table. |
| `stream` | Object containing `stream_arn` and `stream_label`. Present only when streams are enabled. |
| `kms_key_arn` | ARN of the customer-managed KMS key when encryption type is `KMS`. Empty when `AES256` or encryption is disabled. |
| `global_secondary_index_names` | List of names of all created GSIs. |
| `local_secondary_index_names` | List of names of all created LSIs. |

---

## Requirements
* Terraform **1.3+** (for full `optional`/`nullable` object support)
* AWS provider **5.x**

---

## Contributing
Bug reports and pull requests are welcome on GitHub at <https://github.com/project-planton/project-planton>.  
When adding new DynamoDB functionality update **both** the proto definition and the implementation so validation and documentation stay in sync.

---

© Project Planton – released under the Apache 2.0 license.
