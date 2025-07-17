# AWS DynamoDB Table Module (`aws_dynamodb`)

Terraform module that creates and manages an **Amazon DynamoDB table** with optional
secondary indexes, streams, TTL, encryption, point-in-time recovery and tagging.
The interface follows the `AwsDynamodbSpec` protocol buffer schema defined in
`project.planton.provider.aws.awsdynamodb.v1` and the module reports the values
listed in the `AwsDynamodbStackOutputs` schema.

---

## Example Usage

```hcl
module "orders_table" {
  source  = "github.com/project-planton/project-planton//modules/aws_dynamodb"
  version = "x.y.z"

  # ─────────────────────────────────────────────────────────────────────────────
  # Required arguments
  # ─────────────────────────────────────────────────────────────────────────────
  table_name = "orders"

  attribute_definitions = [
    # Partition key
    {
      attribute_name = "order_id"
      attribute_type = "STRING"  # S
    },

    # Sort key (optional)
    {
      attribute_name = "customer_id"
      attribute_type = "STRING"
    },

    # Attribute referenced by the GSI below
    {
      attribute_name = "status"
      attribute_type = "STRING"
    }
  ]

  key_schema = [
    {
      attribute_name = "order_id"
      key_type       = "HASH"   # Partition key
    },
    {
      attribute_name = "customer_id"
      key_type       = "RANGE"  # Sort key
    }
  ]

  billing_mode = "PROVISIONED"
  provisioned_throughput = {
    read_capacity_units  = 5
    write_capacity_units = 5
  }

  # ─────────────────────────────────────────────────────────────────────────────
  # Optional arguments
  # ─────────────────────────────────────────────────────────────────────────────
  global_secondary_indexes = [{
    index_name = "status-index"
    key_schema = [{
      attribute_name = "status"
      key_type       = "HASH"
    }]
    projection = {
      projection_type     = "ALL"
      non_key_attributes  = []  # ignored for ALL|KEYS_ONLY
    }
    provisioned_throughput = {
      read_capacity_units  = 2
      write_capacity_units = 2
    }
  }]

  stream_specification = {
    stream_enabled   = true
    stream_view_type = "NEW_AND_OLD_IMAGES"
  }

  ttl_specification = {
    ttl_enabled   = true
    attribute_name = "expires_at"  # UNIX epoch time in seconds
  }

  sse_specification = {
    enabled            = true
    sse_type           = "KMS"
    kms_master_key_id  = "arn:aws:kms:us-east-1:111122223333:key/abcd-1234-efgh-5678"
  }

  point_in_time_recovery_enabled = true

  tags = {
    environment = "prod"
    team        = "platform"
  }
}
```

---

## Inputs

| Name | Required | Type | Description |
|------|----------|------|-------------|
| `table_name` | **Yes** | `string` | Name of the DynamoDB table (3–255 chars). |
| `attribute_definitions` | **Yes** | `list(object)` | Full list of attributes referenced by the table and any index (see `AttributeDefinition`). |
| `key_schema` | **Yes** | `list(object)` | Table primary key definition (`KeySchemaElement`). Must contain 1 (HASH) or 2 (HASH + RANGE) elements. |
| `billing_mode` | **Yes** | `string` | `PROVISIONED` or `PAY_PER_REQUEST`. When `PROVISIONED`, `provisioned_throughput` is mandatory. |
| `provisioned_throughput` | Conditional | `object` | RCU/WCU settings for the table when `billing_mode = PROVISIONED`. |
| `global_secondary_indexes` | No | `list(object)` | List of Global Secondary Index definitions (`GlobalSecondaryIndex`). |
| `local_secondary_indexes` | No | `list(object)` | List of Local Secondary Index definitions (`LocalSecondaryIndex`). |
| `stream_specification` | No | `object` | Enable DynamoDB Streams and configure the captured image type. |
| `ttl_specification` | No | `object` | TTL settings for automatic item expiry. |
| `sse_specification` | No | `object` | Server-side encryption configuration; required for customer-managed CMKs (`sse_type = KMS`). |
| `point_in_time_recovery_enabled` | No | `bool` | Enable point-in-time recovery (PITR) continuous backups. Defaults to `false`. |
| `tags` | No | `map(string)` | Key/value tags applied to the table and related resources. |

> **Note** – Each complex input object exactly mimics the field names and value
> constraints specified in the protobuf messages so that clients can generate
> Terraform JSON configurations automatically.

---

## Outputs

| Name | Description |
|------|-------------|
| `table_arn` | Fully-qualified ARN of the table. |
| `table_name` | Name of the table (with any runtime suffixes added by the provider). |
| `table_id` | Unique identifier assigned by AWS. |
| `stream` | Object with `stream_arn` and `stream_label`; only exported when streams are enabled. |
| `kms_key_arn` | ARN of the customer-managed CMK used for encryption (when applicable). |
| `global_secondary_index_names` | List of provisioned GSI names. |
| `local_secondary_index_names` | List of provisioned LSI names. |

---

## Requirements

| Name | Version |
|------|---------|
| Terraform | ≥ 1.2 |
| AWS Provider | ≥ 5.0 |

The module is developed and tested with Terraform 1.5 and AWS provider 5.x but
should work with any newer compatible versions.

---

## Contributing

Issues and pull requests are welcome! Please run `terraform fmt` and `terraform
validate` before submitting changes. All new features must maintain backward
compatibility and include documentation updates.

---

## License

Apache 2.0 © Project Planton.
