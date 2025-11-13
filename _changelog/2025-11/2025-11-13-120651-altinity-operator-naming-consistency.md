# AltinityOperator Naming Consistency Refactoring

**Date**: November 13, 2025  
**Type**: Refactoring  
**Components**: API Definitions, Cloud Resource Registry, Documentation, Pulumi CLI Integration

## Summary

Renamed the `AltinityOperatorKubernetes` resource to `AltinityOperator` across all proto definitions, documentation, and implementation code to align with Project Planton's naming conventions for Kubernetes addon operators. This change eliminates redundant "Kubernetes" suffixes and improves consistency with other addon resources like `CertManagerKubernetes`, `ExternalDnsKubernetes`, and similar components that are already scoped to Kubernetes via their provider namespace.

## Problem Statement / Motivation

The Altinity ClickHouse Operator resource was originally named `AltinityOperatorKubernetes`, which included a redundant "Kubernetes" suffix. This naming pattern was inconsistent with Project Planton's design philosophy where:

### Pain Points

- **Redundant Context**: The resource lives under `provider/kubernetes/addon/`, making the "Kubernetes" suffix in the name redundant
- **Verbose API Surface**: Users had to write `kind: AltinityOperatorKubernetes` in manifests, which is unnecessarily long
- **Naming Inconsistency**: Mixed naming patterns across addon operators—some with suffixes, some without
- **Code Verbosity**: Proto message types like `AltinityOperatorKubernetesSpec` and `AltinityOperatorKubernetesStackInput` were excessively long
- **Poor Developer Experience**: The redundancy made code harder to read and type

The provider namespace (`org.project_planton.provider.kubernetes.addon.altinityoperatorkubernetes.v1`) already clearly indicates this is a Kubernetes component, so including "Kubernetes" in every message name adds noise without value.

## Solution / What's New

Performed a comprehensive rename from `AltinityOperatorKubernetes` to `AltinityOperator` across:

1. **Proto API Definitions**: Updated all message types, field references, and validation constraints
2. **Cloud Resource Registry**: Modified the enum entry in `cloud_resource_kind.proto`
3. **Documentation**: Updated all user-facing docs, examples, and implementation guides
4. **Implementation Code**: Modified Go code in Pulumi modules to use renamed types
5. **Test Fixtures**: Updated manifest examples with the new `kind` value

### Naming Convention

The new naming follows Project Planton's established pattern:

```yaml
# Before
apiVersion: kubernetes.project-planton.org/v1
kind: AltinityOperatorKubernetes

# After
apiVersion: kubernetes.project-planton.org/v1
kind: AltinityOperator
```

The provider path remains unchanged (`provider/kubernetes/addon/altinityoperatorkubernetes/`) to avoid breaking existing directory structures and import paths.

## Implementation Details

### Proto File Changes

**File**: `apis/org/project_planton/provider/kubernetes/addon/altinityoperatorkubernetes/v1/api.proto`

```protobuf
// Before
message AltinityOperatorKubernetes {
  string kind = 2 [(buf.validate.field).string.const = 'AltinityOperatorKubernetes'];
  AltinityOperatorKubernetesSpec spec = 4;
  AltinityOperatorKubernetesStatus status = 5;
}

// After
message AltinityOperator {
  string kind = 2 [(buf.validate.field).string.const = 'AltinityOperator'];
  AltinityOperatorSpec spec = 4;
  AltinityOperatorStatus status = 5;
}
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/altinityoperatorkubernetes/v1/spec.proto`

```protobuf
// Before
message AltinityOperatorKubernetesSpec { ... }
message AltinityOperatorKubernetesSpecContainer { ... }

// After
message AltinityOperatorSpec { ... }
message AltinityOperatorSpecContainer { ... }
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/altinityoperatorkubernetes/v1/stack_input.proto`

```protobuf
// Before
message AltinityOperatorKubernetesStackInput {
  AltinityOperatorKubernetes target = 1;
}

// After
message AltinityOperatorStackInput {
  AltinityOperator target = 1;
}
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/altinityoperatorkubernetes/v1/stack_outputs.proto`

```protobuf
// Before
message AltinityOperatorKubernetesStackOutputs { ... }

// After
message AltinityOperatorStackOutputs { ... }
```

### Registry Update

**File**: `apis/org/project_planton/shared/cloudresourcekind/cloud_resource_kind.proto`

```protobuf
// Before
AltinityOperatorKubernetes = 831 [(kind_meta) = {
  provider: kubernetes
  version: v1
  id_prefix: "altopk8s"
  kubernetes_meta: {category: addon}
}];

// After
AltinityOperator = 831 [(kind_meta) = {
  provider: kubernetes
  version: v1
  id_prefix: "altopk8s"
  kubernetes_meta: {category: addon}
}];
```

### Implementation Code Changes

**File**: `apis/org/project_planton/provider/kubernetes/addon/altinityoperatorkubernetes/v1/iac/pulumi/main.go`

```go
// Before
stackInput := &altinityoperatorkubernetesv1.AltinityOperatorKubernetesStackInput{}

// After
stackInput := &altinityoperatorkubernetesv1.AltinityOperatorStackInput{}
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/altinityoperatorkubernetes/v1/iac/pulumi/module/altinity_operator.go`

```go
// Before
func Resources(ctx *pulumi.Context, stackInput *altinityoperatorkubernetesv1.AltinityOperatorKubernetesStackInput) error

// After
func Resources(ctx *pulumi.Context, stackInput *altinityoperatorkubernetesv1.AltinityOperatorStackInput) error
```

### Documentation Updates

Updated all occurrences in:
- `README.md` (main component documentation)
- `examples.md` (YAML manifest examples)
- `iac/pulumi/README.md` (Pulumi module documentation)
- `iac/pulumi/examples.md` (Pulumi usage examples)
- `iac/tf/README.md` (Terraform module documentation)
- `iac/tf/examples.md` (Terraform usage examples)
- `iac/hack/manifest.yaml` (test fixture)

### Build Process

1. **Proto Generation**: Ran `make protos` to regenerate Go stubs from updated proto files
2. **Gazelle Update**: Ran `./bazelw run //:gazelle` to update BUILD.bazel files
3. **Compilation Verification**: Successfully compiled all Go packages
4. **Linter Validation**: Confirmed no linter errors in modified proto files

## Benefits

### Cleaner API Surface

```yaml
# Users write less, understand more
kind: AltinityOperator  # vs. kind: AltinityOperatorKubernetes
```

### Improved Code Readability

Proto message names are now more concise:
- `AltinityOperatorSpec` (was `AltinityOperatorKubernetesSpec`)
- `AltinityOperatorStackInput` (was `AltinityOperatorKubernetesStackInput`)
- `AltinityOperatorStackOutputs` (was `AltinityOperatorKubernetesStackOutputs`)

### Naming Consistency

Aligns with the pattern where the provider namespace provides sufficient context:
- Package: `org.project_planton.provider.kubernetes.addon.altinityoperatorkubernetes.v1`
- Kind: `AltinityOperator` (context is already clear from package)

### Developer Experience

- Shorter type names in code
- Less typing in YAML manifests
- Reduced cognitive load when reading code
- Consistent mental model across all Kubernetes addon operators

## Impact

### User-Facing Changes

**Breaking Change**: Yes, for API consumers

Users must update their YAML manifests:

```yaml
# Update required
apiVersion: kubernetes.project-planton.org/v1
kind: AltinityOperator  # Changed from: AltinityOperatorKubernetes
metadata:
  name: altinity-operator-prod
spec:
  # ... rest unchanged
```

### Non-Breaking Aspects

- **Directory structure**: Unchanged (`provider/kubernetes/addon/altinityoperatorkubernetes/`)
- **Package namespace**: Unchanged (`org.project_planton.provider.kubernetes.addon.altinityoperatorkubernetes.v1`)
- **Import paths**: Unchanged in Go code
- **Functionality**: Zero behavioral changes
- **ID prefix**: Unchanged (`altopk8s`)
- **Enum value**: Unchanged (831)

### Developer Impact

- Proto stubs regenerated automatically
- BUILD files updated via Gazelle
- No manual code migration required for internal implementation
- All existing tests continue to pass

## Files Modified

**Proto Definitions** (4 files):
- `api.proto` - Main API message types
- `spec.proto` - Spec and container message types
- `stack_input.proto` - Stack input message type
- `stack_outputs.proto` - Stack outputs message type

**Registry** (1 file):
- `cloud_resource_kind.proto` - Enum entry

**Documentation** (6 files):
- `README.md` - Component documentation
- `examples.md` - YAML examples
- `iac/pulumi/README.md` - Pulumi module docs
- `iac/pulumi/examples.md` - Pulumi examples
- `iac/tf/README.md` - Terraform module docs
- `iac/tf/examples.md` - Terraform examples

**Implementation** (2 files):
- `iac/pulumi/main.go` - Stack input type reference
- `iac/pulumi/module/altinity_operator.go` - Function signature

**Test Fixtures** (1 file):
- `iac/hack/manifest.yaml` - Test manifest

**Generated Files**:
- `*.pb.go` files (auto-regenerated from proto definitions)
- `BUILD.bazel` files (auto-updated via Gazelle)

**Total**: 14 manually modified files + generated artifacts

## Related Work

This refactoring is part of an ongoing effort to improve naming consistency across Project Planton's Kubernetes addon operators. Similar patterns should be evaluated for:

- Other addon operators that may have redundant suffixes
- Future addon components to follow this cleaner naming convention
- Documentation guidelines to establish naming patterns for new resources

## Migration Notes

For users with existing manifests:

1. **Find and replace** in all manifest files:
   ```bash
   find . -name "*.yaml" -type f -exec sed -i '' 's/kind: AltinityOperatorKubernetes/kind: AltinityOperator/g' {} +
   ```

2. **No CLI changes required** - The `project-planton` CLI will work with the new kind name automatically after updating to the version with this change

3. **No infrastructure impact** - Existing deployed resources are unaffected; this only affects new deployments

---

**Status**: ✅ Production Ready  
**Files Changed**: 14 manual + generated artifacts  
**Build Status**: All tests passing, no linter errors

