# CertManager Naming Consistency Refactoring

**Date**: November 13, 2025  
**Type**: Refactoring  
**Components**: API Definitions, Cloud Resource Registry, Documentation, Pulumi CLI Integration, Test Files

## Summary

Renamed the `CertManagerKubernetes` resource to `CertManager` across all proto definitions, documentation, implementation code, and test files to align with Project Planton's naming conventions for Kubernetes addon operators. This change eliminates redundant "Kubernetes" suffixes and improves consistency with other addon resources that are already scoped to Kubernetes via their provider namespace.

## Problem Statement / Motivation

The Cert-Manager addon resource was originally named `CertManagerKubernetes`, which included a redundant "Kubernetes" suffix. This naming pattern was inconsistent with Project Planton's design philosophy where:

### Pain Points

- **Redundant Context**: The resource lives under `provider/kubernetes/addon/`, making the "Kubernetes" suffix in the name redundant
- **Verbose API Surface**: Users had to write `kind: CertManagerKubernetes` in manifests, which is unnecessarily long
- **Naming Inconsistency**: Mixed naming patterns across addon operators—some with suffixes, some without
- **Code Verbosity**: Proto message types like `CertManagerKubernetesSpec` and `CertManagerKubernetesStackInput` were excessively long
- **Poor Developer Experience**: The redundancy made code harder to read and type

The provider namespace (`org.project_planton.provider.kubernetes.addon.certmanager.v1`) already clearly indicates this is a Kubernetes component, so including "Kubernetes" in every message name adds noise without value.

## Solution / What's New

Performed a comprehensive rename from `CertManagerKubernetes` to `CertManager` across:

1. **Proto API Definitions**: Updated all message types, field references, and validation constraints
2. **Cloud Resource Registry**: Modified the enum entry in `cloud_resource_kind.proto`
3. **Documentation**: Updated README, examples, and YAML manifests
4. **Implementation Code**: Modified Go code in Pulumi modules to use renamed types
5. **Test Files**: Updated test YAML files and example configurations

### Naming Convention

The new naming follows Project Planton's established pattern:

```yaml
# Before
apiVersion: kubernetes.project-planton.org/v1
kind: CertManagerKubernetes

# After
apiVersion: kubernetes.project-planton.org/v1
kind: CertManager
```

The provider path remains unchanged (`provider/kubernetes/addon/certmanager/`) to avoid breaking existing directory structures and import paths.

## Implementation Details

### Proto File Changes

**File**: `apis/org/project_planton/provider/kubernetes/addon/certmanager/v1/api.proto`

```protobuf
// Before
message CertManagerKubernetes {
  string kind = 2 [(buf.validate.field).string.const = 'CertManagerKubernetes'];
  CertManagerKubernetesSpec spec = 4;
  CertManagerKubernetesStatus status = 5;
}

// After
message CertManager {
  string kind = 2 [(buf.validate.field).string.const = 'CertManager'];
  CertManagerSpec spec = 4;
  CertManagerStatus status = 5;
}
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/certmanager/v1/spec.proto`

```protobuf
// Before
message CertManagerKubernetesSpec { ... }

// After
message CertManagerSpec { ... }
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/certmanager/v1/stack_input.proto`

```protobuf
// Before
message CertManagerKubernetesStackInput {
  CertManagerKubernetes target = 1;
}

// After
message CertManagerStackInput {
  CertManager target = 1;
}
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/certmanager/v1/stack_outputs.proto`

```protobuf
// Before
message CertManagerKubernetesStackOutputs { ... }

// After
message CertManagerStackOutputs { ... }
```

### Registry Update

**File**: `apis/org/project_planton/shared/cloudresourcekind/cloud_resource_kind.proto`

```protobuf
// Before
CertManagerKubernetes = 821 [(kind_meta) = {
  provider: kubernetes
  version: v1
  id_prefix: "cmk8s"
  kubernetes_meta: {category: addon}
}];

// After
CertManager = 821 [(kind_meta) = {
  provider: kubernetes
  version: v1
  id_prefix: "cmk8s"
  kubernetes_meta: {category: addon}
}];
```

### Implementation Code Changes

**File**: `apis/org/project_planton/provider/kubernetes/addon/certmanager/v1/iac/pulumi/main.go`

```go
// Before
stackInput := &certmanagerv1.CertManagerKubernetesStackInput{}

// After
stackInput := &certmanagerv1.CertManagerStackInput{}
```

**File**: `apis/org/project_planton/provider/kubernetes/addon/certmanager/v1/iac/pulumi/module/main.go`

```go
// Before
func Resources(ctx *pulumi.Context, stackInput *certmanagerv1.CertManagerKubernetesStackInput) error

// After
func Resources(ctx *pulumi.Context, stackInput *certmanagerv1.CertManagerStackInput) error
```

### Test and Example File Updates

**File**: `hack/test-cloudflare.yaml`

```yaml
# Before
kind: CertManagerKubernetes

# After
kind: CertManager
```

**File**: `hack/example-clusterissuer-cloudflare.yaml`

```yaml
# Before
# ClusterIssuers Created Automatically by CertManagerKubernetes Addon
# The CertManagerKubernetes addon automatically creates ONE ClusterIssuer PER DOMAIN
# based on the dns_providers configuration in the CertManagerKubernetesSpec.

# After
# ClusterIssuers Created Automatically by CertManager Addon
# The CertManager addon automatically creates ONE ClusterIssuer PER DOMAIN
# based on the dns_providers configuration in the CertManagerSpec.
```

### Documentation Updates

Updated all occurrences in:
- `README.md` - Component documentation with usage examples
- `hack/` test files - Example YAML manifests

## Benefits

### Cleaner API Surface

```yaml
# Users write less, understand more
kind: CertManager  # vs. kind: CertManagerKubernetes
```

### Improved Code Readability

Proto message names are now more concise:
- `CertManagerSpec` (was `CertManagerKubernetesSpec`)
- `CertManagerStackInput` (was `CertManagerKubernetesStackInput`)
- `CertManagerStackOutputs` (was `CertManagerKubernetesStackOutputs`)

### Naming Consistency

Aligns with the pattern where the provider namespace provides sufficient context:
- Package: `org.project_planton.provider.kubernetes.addon.certmanager.v1`
- Kind: `CertManager` (context is already clear from package)

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
kind: CertManager  # Changed from: CertManagerKubernetes
metadata:
  name: cert-manager-prod
spec:
  # ... rest unchanged
```

### Non-Breaking Aspects

- **Directory structure**: Unchanged (`provider/kubernetes/addon/certmanager/`)
- **Package namespace**: Unchanged (`org.project_planton.provider.kubernetes.addon.certmanager.v1`)
- **Import paths**: Unchanged in Go code
- **Functionality**: Zero behavioral changes
- **ID prefix**: Unchanged (`cmk8s`)
- **Enum value**: Unchanged (821)

### Developer Impact

- Proto stubs regenerated automatically via `make protos`
- BUILD files updated via Gazelle
- No manual code migration required for internal implementation
- All existing tests and functionality preserved

## Files Modified

**Proto Definitions** (4 files):
- `api.proto` - Main API message types
- `spec.proto` - Spec message type
- `stack_input.proto` - Stack input message type
- `stack_outputs.proto` - Stack outputs message type

**Registry** (1 file):
- `cloud_resource_kind.proto` - Enum entry

**Documentation** (1 file):
- `README.md` - Component documentation with examples

**Test Files** (2 files):
- `hack/test-cloudflare.yaml` - Test manifest
- `hack/example-clusterissuer-cloudflare.yaml` - Example documentation

**Implementation** (2 files):
- `iac/pulumi/main.go` - Stack input type reference
- `iac/pulumi/module/main.go` - Function signature

**Generated Files**:
- `*.pb.go` files (auto-regenerated from proto definitions)
- `BUILD.bazel` files (auto-updated via Gazelle)

**Total**: 10 manually modified files + generated artifacts

## Related Work

This refactoring follows the pattern established by:
- **AltinityOperator** (2025-11-13): Already follows the cleaner naming pattern without Kubernetes suffix
- **ElasticOperator** (enum value 822): Already follows the cleaner naming pattern without Kubernetes suffix
- **ExternalDns** (2025-11-13): Renamed from ExternalDnsKubernetes to ExternalDns

This is part of an ongoing effort to improve naming consistency across Project Planton's Kubernetes addon operators.

## Migration Notes

For users with existing manifests:

1. **Find and replace** in all manifest files:
   ```bash
   find . -name "*.yaml" -type f -exec sed -i 's/kind: CertManagerKubernetes/kind: CertManager/g' {} +
   ```

2. **No CLI changes required** - The `project-planton` CLI will work with the new kind name automatically after updating to the version with this change

3. **No infrastructure impact** - Existing deployed resources are unaffected; this only affects new deployments

---

**Status**: ✅ Production Ready  
**Files Changed**: 10 manual + generated artifacts  
**Build Status**: All builds passing, no linter errors

