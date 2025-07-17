# AWS DynamoDB Terraform Module

A reusable Terraform module that provisions an **Amazon DynamoDB table** together with every frequently-used option – keys, indexes, billing modes, encryption, TTL, Streams, tags, etc.

The public interface (inputs/outputs) is generated from the [AWS DynamoDB spec proto](./awsdynamodb.proto) shipped in this repo, so the module always stays in sync with the service API.

---

## Usage examples

### 1 – Quick start (on-demand billing)
```hcl
module "orders_table" {
  source = "github.com/project-planton/terraform-aws-dynamodb"

  table_name            = "orders"
  billing_mode          = "PAY_PER_REQUEST"
  attribute_definitions = [
    {
      attribute_name = "pk"
      attribute_type = "S"
    },
    {
      attribute_name = "sk"
      attribute_type = "S"
    }
  ]
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

  tags = {
    environment = var.environment
    app         = "checkout"
  }
}
```

### 2 – Provisioned capacity with a GSI
```hcl
module "users_table" {
  source = "github.com/project-planton/terraform-aws-dynamodb"

  table_name                   = "users"
  billing_mode                 = "PROVISIONED"
  provisioned_throughput = {
    read_capacity_units  = 10
    write_capacity_units = 5
  }

  attribute_definitions = [
    { attribute_name = "user_id",   attribute_type = "S" },
    { attribute_name = "email",     attribute_type = "S" },
    { attribute_name = "created_at", attribute_type = "N" }
  ]

  key_schema = [
    { attribute_name = "user_id", key_type = "HASH" }
  ]

  global_secondary_indexes = [
    {
      index_name = "email-gsi"

      key_schema = [
        { attribute_name = "email", key_type = "HASH" }
      ]

      projection = {
        projection_type    = "INCLUDE"
        non_key_attributes = ["created_at"]
      }

      provisioned_throughput = {
        read_capacity_units  = 5
        write_capacity_units = 2
      }
    }
  ]

  stream_specification = {
    stream_enabled   = true
    stream_view_type = "NEW_AND_OLD_IMAGES"
  }

  point_in_time_recovery_enabled = true
  sse_specification = {
    enabled  = true
    sse_type = "KMS"
    kms_master_key_id = aws_kms_key.ddb.arn
  }
}
```

> 💡 **Tip** – every argument corresponds 1-to-1 to the proto field, so you can always look them up in `awsdynamodb.proto`.

---

## Inputs
| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:------:|
| `table_name` | Name of the DynamoDB table. | `string` | n/a | **yes** |
| `attribute_definitions` | List of attribute definitions referenced by the table or any index. | `list(object({ attribute_name = string, attribute_type = string }))` | n/a | **yes** |
| `key_schema` | Primary key schema (partition key and optional sort key). | `list(object({ attribute_name = string, key_type = string }))` | n/a | **yes** |
| `billing_mode` | How the table is billed – `PROVISIONED` or `PAY_PER_REQUEST`. | `string` | `PAY_PER_REQUEST` | no |
| `provisioned_throughput` | Provisioned capacity when `billing_mode = "PROVISIONED"`. | `object({ read_capacity_units = number, write_capacity_units = number })` | `null` | conditional |
| `global_secondary_indexes` | Definitions of global secondary indexes (GSIs). | `list(any)` | `[]` | no |
| `local_secondary_indexes` | Definitions of local secondary indexes (LSIs). | `list(any)` | `[]` | no |
| `stream_specification` | DynamoDB Streams configuration. | `object({ stream_enabled = bool, stream_view_type = string })` | `null` | no |
| `ttl_specification` | TTL (Time-to-Live) settings. | `object({ ttl_enabled = bool, attribute_name = string })` | `null` | no |
| `sse_specification` | Server-side encryption settings. | `object({ enabled = bool, sse_type = string, kms_master_key_id = string })` | `null` | no |
| `point_in_time_recovery_enabled` | Enables point-in-time recovery (PITR). | `bool` | `true` | no |
| `tags` | Key/value tags applied to the table. | `map(string)` | `{}` | no |

*Refer to `awsdynamodb.proto` for the complete, strongly-typed schema including nested index objects, projections, etc.*

---

## Outputs
| Name | Description |
|------|-------------|
| `table_arn` | Fully-qualified ARN of the table. |
| `table_name` | Name of the DynamoDB table (with any runtime suffixes). |
| `table_id` | AWS-assigned unique identifier of the table. |
| `stream` | Object with the latest stream ARN & label (only when Streams are enabled). |
| `kms_key_arn` | ARN of the CMK when encryption uses KMS. |
| `global_secondary_index_names` | List of provisioned GSI names. |
| `local_secondary_index_names` | List of provisioned LSI names. |

---

## Terraform requirements
| Name | Version |
|------|---------|
| Terraform | >= 1.3 |
| AWS Provider | >= 5.0 |

---

## Module architecture
* **Single resource** – the module wraps a single `aws_dynamodb_table` resource plus optional helper resources (KMS key, IAM policies) if encryption or Streams are enabled.
* **Proto-driven** – `awsdynamodb.proto` captures 100 % of the public API. CI regenerates the Terraform variables and docs from the proto on every commit, eliminating drift.
* **Validation first** – All cross-field rules (e.g. *RCU/WCU must be set when billing is PROVISIONED*) are expressed in the proto and re-used in Terraform via generated `validation` blocks, preventing mis-configuration before the plan reaches AWS.
* **No surprises** – The module does not create IAM roles, alarms or autoscaling policies on your behalf; instead it exposes granular outputs so you can plug them into dedicated modules.

---

## Developing & contributing
1. Make your changes in `awsdynamodb.proto`.
2. Run `make generate` to re-create variable definitions & docs.
3. Open a PR – CI will run `buf`, `tfsec`, `terraform validate` and `terraform test`.

---

© Project Planton – Licensed under the Apache 2.0 License.
