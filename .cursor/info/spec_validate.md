# Spec Validate Authoring Guide

Purpose: add validation rules to an existing `spec.proto` without changing its schema.

## Scope
- Do not rename/add/remove fields or messages.
- Only add validations and brief comments.

## Import
- Add once near the top:
  - `import "buf/validate/validate.proto";`

## Field-level validations (80/20)
- Strings:
  - Names/IDs: `(buf.validate.field).string.min_len = 1`
  - ARNs/URIs: add `min_len` and, if stable, a simple `pattern` (avoid brittle regexes)
  - Domains/emails: prefer minimal `pattern` or `min_len`
- Enums:
  - `(buf.validate.field).enum.defined_only = true`
  - If field must be set and enum has `*_UNSPECIFIED = 0`, add CEL to forbid 0
- Numbers:
  - Use `gt/gte/lt/lte` as appropriate
- Booleans:
  - Typically no direct validation; enforce via CEL when tied to other fields
- Repeated:
  - `(buf.validate.field).repeated.min_items = 1` when at least one is required
  - Consider `(buf.validate.field).repeated.unique = true` for sets like aliases
- Bytes/Maps:
  - Apply min/max sizes if applicable

## Message-level CEL validations
- Require B when A is set:
  - Example: if `aliases` non-empty, `certificate_arn` must be non-empty
  - `this.aliases.size() == 0 || this.certificate_arn != ""`
- Mutually exclusive fields:
  - `this.x == "" || this.y == ""`
- Enum-dependent constraints (e.g., DynamoDB billing mode):
  - If `billing_mode == PROVISIONED`, then `read_capacity_units > 0 && write_capacity_units > 0`, else both 0/unset
- Ordered ranges:
  - `(this.min_ttl == 0 || this.default_ttl >= this.min_ttl) && (this.max_ttl == 0 || this.default_ttl <= this.max_ttl)`

## Example (adapt to your schema)
```proto
syntax = "proto3";
package project.planton.provider.aws.awscloudfront.v1;

import "buf/validate/validate.proto";

message AwsCloudFrontSpec {
  repeated string aliases = 1 [(buf.validate.field).repeated = {min_items: 1, unique: true}];

  string certificate_arn = 2 [(buf.validate.field).string.min_len = 1];

  enum PriceClass {
    PRICE_CLASS_UNSPECIFIED = 0;
    PRICE_CLASS_100 = 1;
    PRICE_CLASS_200 = 2;
    PRICE_CLASS_ALL = 3;
  }
  PriceClass price_class = 3 [(buf.validate.field).enum.defined_only = true];

  option (buf.validate.message).cel = {
    id: "aliases_require_cert",
    message: "certificate_arn must be set when aliases are provided",
    expression: "this.aliases.size() == 0 || this.certificate_arn != \"\""
  };
}
```

## Notes
- Prefer pragmatic rules with low false positives.
- Ensure compatibility with protovalidate-go (uses `buf/validate/validate.proto`).
