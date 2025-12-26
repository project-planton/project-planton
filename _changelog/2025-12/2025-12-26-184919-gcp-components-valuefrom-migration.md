# GCP Components ValueFrom Migration: Cross-Resource Reference Support

**Date**: December 26, 2025
**Type**: Enhancement
**Components**: GcpGcsBucket, GcpDnsZone, GcpArtifactRegistryRepo, API Definitions, Pulumi IAC, Terraform IAC

## Summary

Migrated three GCP deployment components (`GcpGcsBucket`, `GcpDnsZone`, `GcpArtifactRegistryRepo`) from plain `string` types to `StringValueOrRef` for project ID fields, enabling cross-resource references. This allows users to dynamically reference outputs from other resources (like `GcpProject`) instead of hardcoding values, improving infrastructure composition and dependency management.

## Problem Statement / Motivation

The existing GCP components used plain `string` types for `project_id` fields, which required users to hardcode GCP project IDs directly in their manifests. This created several issues:

### Pain Points

- **Hardcoded Dependencies**: Users had to manually copy project IDs between resources
- **No Dynamic References**: Couldn't reference outputs from other Project Planton resources
- **Manual Coordination**: Changes to project IDs required updating multiple manifests
- **Limited Composability**: Resources couldn't be chained together declaratively
- **Inconsistent Patterns**: Some GCP components already supported `StringValueOrRef`, others didn't

## Solution / What's New

Implemented the `StringValueOrRef` pattern across three GCP components, enabling both literal values and cross-resource references:

### Before (Hardcoded)

```yaml
spec:
  gcpProjectId: my-gcp-project-123  # Plain string
```

### After (With ValueFrom Support)

```yaml
# Option 1: Literal value (backward compatible)
spec:
  gcpProjectId:
    value: my-gcp-project-123

# Option 2: Cross-resource reference
spec:
  gcpProjectId:
    valueFrom:
      kind: GcpProject
      name: main-project
      fieldPath: status.outputs.project_id
```

### Components Migrated

| Component | Field Changed | Default Reference Kind |
|-----------|--------------|----------------------|
| GcpGcsBucket | `gcp_project_id` | GcpProject |
| GcpDnsZone | `project_id` | GcpProject |
| GcpArtifactRegistryRepo | `project_id` | GcpProject |

## Implementation Details

### Proto Schema Changes

Added `StringValueOrRef` type with field options for default kind resolution:

```protobuf
import "org/project_planton/shared/foreignkey/v1/foreign_key.proto";

message GcpGcsBucketSpec {
  org.project_planton.shared.foreignkey.v1.StringValueOrRef gcp_project_id = 1 [
    (buf.validate.field).required = true,
    (org.project_planton.shared.foreignkey.v1.default_kind) = GcpProject,
    (org.project_planton.shared.foreignkey.v1.default_kind_field_path) = "status.outputs.project_id"
  ];
}
```

### Pulumi IAC Updates

Updated Go code to use `.GetValue()` method for extracting the string value:

```go
// Before
Project: pulumi.String(locals.GcpGcsBucket.Spec.GcpProjectId),

// After
Project: pulumi.String(locals.GcpGcsBucket.Spec.GcpProjectId.GetValue()),
```

### Terraform IAC Updates

Updated `variables.tf` to accept object type with `value` field:

```hcl
variable "spec" {
  type = object({
    gcp_project_id = object({
      value = string
    })
    # ... other fields
  })
}
```

Updated `main.tf` to reference nested value:

```hcl
resource "google_storage_bucket" "main" {
  project = var.spec.gcp_project_id.value
  # ...
}
```

### Test Updates

Added tests for both literal values and `valueFrom` references:

```go
ginkgo.It("should not return a validation error with valueFrom reference", func() {
    input := &GcpGcsBucket{
        Spec: &GcpGcsBucketSpec{
            GcpProjectId: &foreignkeyv1.StringValueOrRef{
                LiteralOrRef: &foreignkeyv1.StringValueOrRef_ValueFrom{
                    ValueFrom: &foreignkeyv1.ValueFromRef{
                        Name:      "main-project",
                        FieldPath: "status.outputs.project_id",
                    },
                },
            },
            // ...
        },
    }
    err := protovalidate.Validate(input)
    gomega.Expect(err).To(gomega.BeNil())
})
```

## Files Changed

### GcpGcsBucket (11 files)
- `spec.proto` - Added import and changed field type
- `spec.pb.go` - Regenerated protobuf code
- `spec_test.go` - Added valueFrom test cases
- `BUILD.bazel` - Added foreignkey dependency
- `iac/pulumi/module/main.go` - Updated to use `.GetValue()`
- `iac/tf/variables.tf` - Updated variable structure
- `iac/tf/main.tf` - Complete overhaul with all spec features
- `iac/tf/locals.tf` - Added user label merging
- `iac/tf/hack/manifest.yaml` - Updated test manifest
- `examples.md` - Added valueFrom examples
- `iac/pulumi/examples.md` - Added valueFrom examples

### GcpDnsZone (11 files)
- Similar changes to spec.proto, IAC code, and examples

### GcpArtifactRegistryRepo (13 files)
- Similar changes to spec.proto, IAC code, and examples

### Frontend TypeScript (3 files)
- Regenerated TypeScript protobuf stubs

## Benefits

### For Users
- **Dynamic Infrastructure**: Reference project outputs instead of hardcoding
- **Reduced Errors**: No more copying/pasting project IDs incorrectly
- **Better Composition**: Chain resources together declaratively
- **Environment Flexibility**: Use `env` field to reference resources in specific environments

### For Developers
- **Consistent Pattern**: Aligns with already-migrated components (GcpVpc, GcpGkeCluster, etc.)
- **Type Safety**: Field options provide hints for reference resolution
- **Extensible**: Same pattern can be applied to other fields (VPC, subnet, etc.)

## Impact

### Breaking Changes

⚠️ **YAML Manifest Format Change**

Users must update their manifests to use the new nested format:

```yaml
# Old format (no longer valid)
spec:
  gcpProjectId: my-project

# New format (required)
spec:
  gcpProjectId:
    value: my-project
```

### Backward Compatibility

The protobuf wire format is backward compatible - existing code using the proto messages will continue to work. Only the YAML/JSON manifest format changes.

### Migration Path

1. Update manifests to use `{ value: "..." }` format
2. Optionally convert to `valueFrom` references for dynamic dependencies

## Related Work

- **Analysis Document**: `apis/gcp-value-from-analysis.md` - Comprehensive analysis of all GCP components
- **Already Migrated**: GcpVpc, GcpSubnetwork, GcpGkeCluster, GcpRouterNat, GcpGkeWorkloadIdentityBinding
- **Future Work**: GcpCloudRun, GcpCloudSql, GcpCloudFunction, GcpSecretsManager, GcpServiceAccount, GcpCloudCdn

## Verification

All changes verified with:
- ✅ Proto regeneration (`make protos`)
- ✅ Go build validation
- ✅ Component-specific tests (4 tests pass)
- ✅ TypeScript stub generation

---

**Status**: ✅ Production Ready
**Timeline**: Single session implementation
**Code Metrics**: 38 files changed, 897 insertions, 204 deletions
