# GCP Components ValueFrom Migration Analysis

## Executive Summary

This document provides a comprehensive analysis of all GCP deployment components in the Project Planton codebase, identifying fields that currently use plain `string` types for `project-id`, `network/vpc-id`, and `subnetwork-id` references that should be migrated to use `StringValueOrRef` for cross-resource references.

**Total Components Analyzed:** 17
**Components Requiring Changes:** 9
**Components Already Compliant:** 8

---

## Background: StringValueOrRef Pattern

The `StringValueOrRef` type enables flexible resource references:

```protobuf
message StringValueOrRef {
  oneof literal_or_ref {
    string value = 1;              // Direct string value
    ValueFromRef value_from = 2;   // Reference to another resource
  }
}

message ValueFromRef {
  CloudResourceKind kind = 1;      // e.g., GcpProject, GcpVpc
  string env = 2;                  // Environment name
  string name = 3;                 // Resource name
  string field_path = 4;           // Path to field, e.g., "status.outputs.project_id"
}
```

This pattern allows users to:
- **Hard-code values** when appropriate: `{value: "my-project-id"}`
- **Reference other resources** for dynamic dependencies: `{value_from: {kind: GcpProject, name: "main-project", field_path: "status.outputs.project_id"}}`

---

## Components Already Using StringValueOrRef (✅ Compliant)

These components have already been migrated and follow best practices:

### 1. GcpVpc
**File:** `apis/org/project_planton/provider/gcp/gcpvpc/v1/spec.proto`

**Compliant Fields:**
```protobuf
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

---

### 2. GcpSubnetwork
**File:** `apis/org/project_planton/provider/gcp/gcpsubnetwork/v1/spec.proto`

**Compliant Fields:**
```protobuf
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];

org.project_planton.shared.foreignkey.v1.StringValueOrRef vpc_self_link = 2 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpVpc,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.network_self_link"
];
```

---

### 3. GcpGkeCluster
**File:** `apis/org/project_planton/provider/gcp/gcpgkecluster/v1/spec.proto`

**Compliant Fields:**
```protobuf
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];

org.project_planton.shared.foreignkey.v1.StringValueOrRef network_self_link = 2 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpVpc,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.self_link"
];

org.project_planton.shared.foreignkey.v1.StringValueOrRef subnetwork_self_link = 4 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpSubnetwork,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.self_link"
];
```

---

### 4. GcpGkeNodePool
**File:** `apis/org/project_planton/provider/gcp/gcpgkenodepool/v1/spec.proto`

**Compliant Fields:**
```protobuf
org.project_planton.shared.foreignkey.v1.StringValueOrRef cluster_project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpGkeCluster,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "spec.project_id"
];

org.project_planton.shared.foreignkey.v1.StringValueOrRef cluster_location = 3 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpGkeCluster,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "spec.location"
];
```

---

### 5. GcpRouterNat
**File:** `apis/org/project_planton/provider/gcp/gcprouternat/v1/spec.proto`

**Compliant Fields:**
```protobuf
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];

org.project_planton.shared.foreignkey.v1.StringValueOrRef vpc_self_link = 2 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpVpc,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.network_self_link"
];

repeated org.project_planton.shared.foreignkey.v1.StringValueOrRef subnetwork_self_links = 4;
```

---

### 6. GcpCertManagerCert
**File:** `apis/org/project_planton/provider/gcp/gcpcertmanagercert/v1/spec.proto`

**Compliant Fields:**
```protobuf
org.project_planton.shared.foreignkey.v1.StringValueOrRef cloud_dns_zone_id = 4 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpDnsZone,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.zone_name"
];
```

**Note:** `gcp_project_id` (line 15) is still a plain string - **NEEDS CHANGE**.

---

### 7. GcpGkeWorkloadIdentityBinding
**File:** `apis/org/project_planton/provider/gcp/gcpgkeworkloadidentitybinding/v1/spec.proto`

**Compliant Fields:**
```protobuf
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];

org.project_planton.shared.foreignkey.v1.StringValueOrRef service_account_email = 2 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpServiceAccount,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.email"
];
```

---

### 8. GcpProject
**File:** `apis/org/project_planton/provider/gcp/gcpproject/v1/spec.proto`

**Current Definition:**
```protobuf
string project_id = 1 [(buf.validate.field) = {
  required: true
  string: {
    min_len: 6
    max_len: 30
    pattern: "^[a-z][a-z0-9-]*[a-z0-9]$"
  }
}];
```

**Status:** ✅ CORRECT - This is the project ID being **created**, not a reference to another project. Should remain as `string`.

---

## Components Requiring Migration (❌ Needs Changes)

### 1. GcpCloudRun ⚠️ HIGH PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpcloudrun/v1/spec.proto`

#### Changes Required:

**Change 1: project_id (Line 16)**
```protobuf
// BEFORE
string project_id = 1 [
  (buf.validate.field).required = true,
  (buf.validate.field).string = {pattern: "^[a-z][a-z0-9-]{4,28}[a-z0-9]$"}
];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Change 2: VpcAccess.network (Line 237)**
```protobuf
// BEFORE
string network = 1 [
  (buf.validate.field).ignore = IGNORE_IF_ZERO_VALUE,
  (buf.validate.field).string = {min_len: 1}
];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef network = 1 [
  (buf.validate.field).ignore = IGNORE_IF_ZERO_VALUE,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpVpc,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.network_name"
];
```

**Change 3: VpcAccess.subnet (Line 243)**
```protobuf
// BEFORE
string subnet = 2 [
  (buf.validate.field).ignore = IGNORE_IF_ZERO_VALUE,
  (buf.validate.field).string = {min_len: 1}
];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef subnet = 2 [
  (buf.validate.field).ignore = IGNORE_IF_ZERO_VALUE,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpSubnetwork,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.subnetwork_name"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpcloudrun/v1/iac/` to handle reference resolution
- Update examples in `gcpcloudrun/v1/examples.md`

---

### 2. GcpCloudSql ⚠️ HIGH PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpcloudsql/v1/spec.proto`

#### Changes Required:

**Change 1: project_id (Line 15)**
```protobuf
// BEFORE
string project_id = 1 [
  (buf.validate.field).required = true,
  (buf.validate.field).string = {pattern: "^[a-z][a-z0-9-]{4,28}[a-z0-9]$"}
];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Change 2: GcpCloudSqlNetwork.vpc_id (Line 122)**
```protobuf
// BEFORE
string vpc_id = 1;

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef vpc_id = 1 [
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpVpc,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.network_id"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpcloudsql/v1/iac/` to handle reference resolution
- Update examples in `gcpcloudsql/v1/examples.md`
- Update validation in line 148 to handle StringValueOrRef

---

### 3. GcpCloudFunction ⚠️ MEDIUM PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpcloudfunction/v1/spec.proto`

#### Changes Required:

**Change 1: project_id (Line 14)**
```protobuf
// BEFORE
string project_id = 1 [
  (buf.validate.field).required = true,
  (buf.validate.field).string = {pattern: "^[a-z][a-z0-9-]{4,28}[a-z0-9]$"}
];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Note on vpc_connector (Line 165):** The `vpc_connector` field contains a full resource path format `projects/{project}/locations/{region}/connectors/{connector-name}`, not just a VPC reference. This is a composite string that likely should **NOT** be changed to StringValueOrRef as it's a complete resource identifier, not a simple reference.

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpcloudfunction/v1/iac/` to handle reference resolution
- Update examples in `gcpcloudfunction/v1/examples.md`

---

### 4. GcpServiceAccount ⚠️ MEDIUM PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpserviceaccount/v1/spec.proto`

#### Changes Required:

**Change 1: project_id (Line 23)**
```protobuf
// BEFORE
string project_id = 2;

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 2 [
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpserviceaccount/v1/iac/` to handle reference resolution
- Update examples in `gcpserviceaccount/v1/examples.md`

---

### 5. GcpSecretsManager ⚠️ MEDIUM PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpsecretsmanager/v1/spec.proto`

#### Changes Required:

**Change 1: project_id (Line 13)**
```protobuf
// BEFORE
string project_id = 1 [(buf.validate.field).required = true];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpsecretsmanager/v1/iac/` to handle reference resolution
- Update examples in `gcpsecretsmanager/v1/examples.md`

---

### 6. GcpDnsZone ⚠️ MEDIUM PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpdnszone/v1/spec.proto`

#### Changes Required:

**Change 1: project_id (Line 14)**
```protobuf
// BEFORE
string project_id = 1 [(buf.validate.field).required = true];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpdnszone/v1/iac/` to handle reference resolution
- Update examples in `gcpdnszone/v1/examples.md`

---

### 7. GcpArtifactRegistryRepo ⚠️ MEDIUM PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpartifactregistryrepo/v1/spec.proto`

#### Changes Required:

**Change 1: project_id (Line 31)**
```protobuf
// BEFORE
string project_id = 2 [(buf.validate.field).required = true];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef project_id = 2 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpartifactregistryrepo/v1/iac/` to handle reference resolution
- Update examples in `gcpartifactregistryrepo/v1/examples.md`

---

### 8. GcpCloudCdn ⚠️ LOW PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpcloudcdn/v1/spec.proto`

#### Changes Required:

**Change 1: gcp_project_id (Line 17)**
```protobuf
// BEFORE
string gcp_project_id = 1 [
  (buf.validate.field).required = true,
  (buf.validate.field).string.pattern = "^[a-z][a-z0-9-]{4,28}[a-z0-9]$"
];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef gcp_project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpcloudcdn/v1/iac/` to handle reference resolution
- Update examples in `gcpcloudcdn/v1/examples.md`

---

### 9. GcpGcsBucket ⚠️ LOW PRIORITY
**File:** `apis/org/project_planton/provider/gcp/gcpgcsbucket/v1/spec.proto`

#### Changes Required:

**Change 1: gcp_project_id (Line 16)**
```protobuf
// BEFORE
string gcp_project_id = 1 [
  (buf.validate.field).required = true,
  (buf.validate.field).string.pattern = "^[a-z][a-z0-9-]{4,28}[a-z0-9]$"
];

// AFTER
org.project_planton.shared.foreignkey.v1.StringValueOrRef gcp_project_id = 1 [
  (buf.validate.field).required = true,
  (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
  (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
];
```

**Additional Changes:**
- Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- Update IAC code in `gcpgcsbucket/v1/iac/` to handle reference resolution
- Update examples in `gcpgcsbucket/v1/examples.md`

---

## Migration Checklist for Each Component

For each component that needs changes, the following steps must be performed:

### 1. Proto File Changes
- [ ] Add import: `import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";`
- [ ] Change field type from `string` to `org.project_planton.shared.foreignkey.v1.StringValueOrRef`
- [ ] Add field options:
  - `(org.project_planton.shared.foreignkey.v1.default_kind)` - specify the default resource kind
  - `(org.project_planton.shared.foreignkey.v1.default_kind_field_path)` - specify the field path
- [ ] Remove string pattern validation (handled by reference resolution)
- [ ] Keep `(buf.validate.field).required = true` if field was required

### 2. Generated Code
- [ ] Run `make buf-generate` to regenerate Go/TypeScript code
- [ ] Verify generated code compiles without errors
- [ ] Update any BUILD.bazel files if needed (usually auto-generated by Gazelle)

### 3. IAC Code Updates (Pulumi/Terraform)
- [ ] Update `iac/pulumi/main.go` or `iac/tf/main.tf` to resolve references
- [ ] Implement reference resolution logic using shared utilities
- [ ] Handle both literal values and references
- [ ] Add error handling for unresolved references

### 4. Examples and Documentation
- [ ] Update `examples.md` with both literal and reference examples
- [ ] Add migration guide in component README
- [ ] Document breaking changes

### 5. Testing
- [ ] Update unit tests for spec validation
- [ ] Add integration tests for reference resolution
- [ ] Test both literal values and cross-resource references
- [ ] Verify backward compatibility (if maintaining old format)

---

## Implementation Order (Recommended)

### Phase 1: High Priority (Foundation Services)
1. **GcpCloudRun** - Widely used serverless compute
2. **GcpCloudSql** - Critical database service with VPC dependencies

### Phase 2: Medium Priority (Supporting Services)
3. **GcpCloudFunction** - Serverless functions
4. **GcpServiceAccount** - IAM foundation
5. **GcpSecretsManager** - Security foundation
6. **GcpDnsZone** - Network foundation
7. **GcpArtifactRegistryRepo** - Build artifact storage

### Phase 3: Low Priority (Advanced Services)
8. **GcpCloudCdn** - Content delivery
9. **GcpGcsBucket** - Object storage

---

## Breaking Changes Notice

⚠️ **IMPORTANT:** This migration introduces breaking changes to the API specification.

### Impact Assessment:

1. **Existing YAML Manifests:**
   - Manifests with hardcoded string values will continue to work if backward compatibility is implemented
   - Example: `project_id: "my-project"` → `project_id: {value: "my-project"}`

2. **Generated Code:**
   - Go structs will change from `string` to `*StringValueOrRef`
   - TypeScript interfaces will change similarly
   - Client code will need updates

3. **IAC Code:**
   - Pulumi/Terraform code must handle both value types
   - Reference resolution must be implemented

### Mitigation Strategies:

1. **Backward Compatibility Layer:**
   - Implement auto-conversion from plain strings to `{value: "string"}`
   - Deprecation warnings for old format
   - Grace period before removing compatibility

2. **Phased Rollout:**
   - Implement per component over multiple releases
   - Provide migration tools/scripts
   - Comprehensive documentation

3. **Testing:**
   - Maintain test suites for both formats during transition
   - Integration tests for reference resolution
   - Example manifests for both patterns

---

## Reference Resolution Implementation Notes

### Current State
Based on the codebase analysis, the reference resolution pattern exists but **is not yet fully implemented** in the IAC layer.

From `apis/org/project_planton/provider/civo/civofirewall/v1/iac/pulumi/overview.md`:
> **Current limitation**: Reference resolution is not yet implemented. Only literal `value` is used. References will fail silently (empty string).
> **Future work**: Implement reference resolution in a shared library that all modules can use.

### Required Implementation

A shared reference resolution library must be created at:
- **Go:** `pkg/valuefrom/resolver.go` (already exists at `internal/valuefrom/value_from.go` - may need enhancement)
- **TypeScript:** `app/frontend/src/lib/valuefrom/resolver.ts`

The resolver should:
1. Accept a `StringValueOrRef` input
2. If `value` is set, return it directly
3. If `value_from` is set:
   - Look up the referenced resource by `kind`, `env`, `name`
   - Extract the field at `field_path`
   - Return the resolved value
4. Handle errors gracefully (missing resources, invalid paths)

---

## Summary of Components Needing Changes

| Component | File | Fields to Change | Priority |
|-----------|------|------------------|----------|
| GcpCloudRun | `gcpcloudrun/v1/spec.proto` | `project_id`, `network`, `subnet` | HIGH |
| GcpCloudSql | `gcpcloudsql/v1/spec.proto` | `project_id`, `vpc_id` | HIGH |
| GcpCloudFunction | `gcpcloudfunction/v1/spec.proto` | `project_id` | MEDIUM |
| GcpServiceAccount | `gcpserviceaccount/v1/spec.proto` | `project_id` | MEDIUM |
| GcpSecretsManager | `gcpsecretsmanager/v1/spec.proto` | `project_id` | MEDIUM |
| GcpDnsZone | `gcpdnszone/v1/spec.proto` | `project_id` | MEDIUM |
| GcpArtifactRegistryRepo | `gcpartifactregistryrepo/v1/spec.proto` | `project_id` | MEDIUM |
| GcpCloudCdn | `gcpcloudcdn/v1/spec.proto` | `gcp_project_id` | LOW |
| GcpGcsBucket | `gcpgcsbucket/v1/spec.proto` | `gcp_project_id` | LOW |

---

## Example Migration: GcpCloudRun

### Before (Current)
```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudRun
metadata:
  name: my-api
spec:
  project_id: "my-gcp-project-123"
  region: "us-central1"
  vpc_access:
    network: "my-vpc-network"
    subnet: "my-subnet-name"
  container:
    # ... container config
```

### After (With ValueFrom)
```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudRun
metadata:
  name: my-api
spec:
  project_id:
    value_from:
      kind: GcpProject
      name: main-project
      field_path: "status.outputs.project_id"
  region: "us-central1"
  vpc_access:
    network:
      value_from:
        kind: GcpVpc
        name: main-vpc
        field_path: "status.outputs.network_name"
    subnet:
      value_from:
        kind: GcpSubnetwork
        name: private-subnet
        field_path: "status.outputs.subnetwork_name"
  container:
    # ... container config
```

### After (With Literal Values - Backward Compat)
```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudRun
metadata:
  name: my-api
spec:
  project_id:
    value: "my-gcp-project-123"
  region: "us-central1"
  vpc_access:
    network:
      value: "my-vpc-network"
    subnet:
      value: "my-subnet-name"
  container:
    # ... container config
```

---

## Additional Considerations

### 1. Field Path Standards

Establish consistent field path patterns:
- Project ID: `"status.outputs.project_id"`
- VPC Network Self Link: `"status.outputs.network_self_link"`
- VPC Network Name: `"status.outputs.network_name"`
- VPC Network ID: `"status.outputs.network_id"`
- Subnetwork Self Link: `"status.outputs.subnetwork_self_link"`
- Subnetwork Name: `"status.outputs.subnetwork_name"`

### 2. Validation Enhancement

Add CEL validation to ensure either `value` or `value_from` is set:
```protobuf
option (buf.validate.message).cel = {
  id: "string_value_or_ref.oneof_required"
  message: "Either value or value_from must be set"
  expression: "has(this.value) || has(this.value_from)"
};
```

### 3. Output Standardization

Ensure all GCP components have standardized outputs in `stack_outputs.proto`:
- Always include `project_id` if component creates GCP resources
- Always include self-links for referenceable resources
- Always include resource names for user-friendly references

---

## Conclusion

This migration will significantly improve the flexibility and composability of GCP deployment components by enabling cross-resource references. While it introduces breaking changes, careful phased implementation with backward compatibility can minimize disruption.

**Next Steps:**
1. Review and approve this migration plan
2. Implement reference resolution shared library
3. Start with Phase 1 (GcpCloudRun, GcpCloudSql)
4. Test thoroughly with both literal and reference patterns
5. Proceed to subsequent phases

---

**Document Version:** 1.0
**Last Updated:** 2025-12-26
**Author:** AI Assistant (Claude Sonnet 4.5)

