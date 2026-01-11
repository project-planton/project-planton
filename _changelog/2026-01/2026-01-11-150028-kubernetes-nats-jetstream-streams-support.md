# KubernetesNats JetStream Streams Support with NACK Controller

**Date**: January 11, 2026
**Type**: Feature
**Components**: Kubernetes Provider, API Definitions, Pulumi CLI Integration, IAC Stack Runner

## Summary

Extended the KubernetesNats deployment component to support NACK (NATS Controllers for Kubernetes) and declarative JetStream stream creation. Users can now opt-in to deploy the NACK controller alongside NATS and define streams directly in their YAML manifests, which are reconciled as Kubernetes custom resources.

## Problem Statement / Motivation

NATS JetStream provides powerful streaming capabilities, but managing streams required either:
1. Manual stream creation via NATS CLI after deployment
2. Custom scripts or external tooling
3. Separate Kubernetes manifests for NACK CRs

This created operational friction and prevented infrastructure-as-code for the complete NATS ecosystem.

### Pain Points

- No declarative way to define JetStream streams in KubernetesNats manifests
- Stream configuration scattered across multiple tools and workflows
- NACK controller deployment was a separate manual process
- No type safety or validation for stream/consumer configurations
- Race conditions when deploying CRDs and CRs in the same module

## Solution / What's New

Integrated NACK controller deployment and JetStream stream management directly into the KubernetesNats Pulumi module with a clean API.

### Key Features

1. **Opt-in NACK Controller**: Deploy NACK alongside NATS with a simple `nack_controller.enabled: true`
2. **Declarative Streams**: Define streams with subjects, retention, storage, and limits in YAML
3. **Nested Consumers**: Attach consumers to streams for complete JetStream configuration
4. **Strongly-typed Resources**: Generated Go types from NACK CRDs for type-safe Pulumi code
5. **Configurable Versions**: Both Helm chart version and app version exposed for flexibility

### Deployment Flow

```mermaid
flowchart TB
    subgraph Pulumi["Pulumi Module Execution"]
        A[NATS Helm Chart] --> B[NACK CRDs]
        B --> C[NACK Controller Helm]
        C --> D[Stream CRs]
        D --> E[Consumer CRs]
    end
    
    subgraph K8s["Kubernetes Cluster"]
        F[NATS StatefulSet]
        G[NACK Deployment]
        H[Stream Resources]
        I[Consumer Resources]
    end
    
    A --> F
    C --> G
    D --> H
    E --> I
    G -.->|reconciles| H
    G -.->|reconciles| I
```

## Implementation Details

### API Design (spec.proto)

Extended `KubernetesNatsSpec` with:

```protobuf
// NACK controller configuration
KubernetesNatsNackController nack_controller = 10;

// JetStream streams
repeated KubernetesNatsStream streams = 11;

// NATS Helm chart version
optional string nats_helm_chart_version = 12;
```

**NACK Controller Message**:
```protobuf
message KubernetesNatsNackController {
  bool enabled = 1;
  bool enable_control_loop = 2;
  optional string helm_chart_version = 3;  // Default: "0.31.1"
  optional string app_version = 4;          // Default: "0.21.1"
}
```

**Key Design Decision**: Chart version (`0.31.1`) differs from app version (`0.21.1`). The app version maps to GitHub tags and CRD schemas, while chart version is for Helm. Both are exposed to users.

### Enum Wrapper Pattern

Protobuf enum scoping conflicts required a wrapper message pattern:

```protobuf
message StreamStorageEnum {
  enum Value {
    unspecified = 0;
    file = 1;
    memory = 2;
  }
}
```

This preserves NATS-native values (`file`, `memory`) in YAML output while avoiding C++ scoping collisions.

### Pulumi Module Architecture

```mermaid
flowchart LR
    subgraph Module["Pulumi Module"]
        main["main.go"] --> helm["helm_chart.go"]
        main --> crds["nack_crds.go"]
        main --> ctrl["nack_controller.go"]
        main --> streams["streams.go"]
    end
    
    subgraph Types["Generated Types"]
        nack["nack/v1beta2"]
        nack --> Stream
        nack --> Consumer
    end
    
    streams --> nack
```

**New Files**:
- `nack_crds.go` - Deploys CRDs from versioned GitHub URL
- `nack_controller.go` - Deploys NACK Helm chart with NATS connection config
- `streams.go` - Creates Stream/Consumer CRs using strongly-typed resources

### Strongly-Typed Resources

Generated Go types from NACK CRDs using `crd2pulumi`:

```go
// Instead of generic apiextensions.CustomResource
stream, err := nackv1beta2.NewStream(ctx, resourceName,
    &nackv1beta2.StreamArgs{
        Metadata: &kubernetesmeta.ObjectMetaArgs{...},
        Spec: &nackv1beta2.StreamSpecArgs{
            Name:      pulumi.StringPtr(stream.Name),
            Subjects:  pulumi.ToStringArray(stream.Subjects),
            Storage:   pulumi.StringPtr("file"),
            Replicas:  pulumi.IntPtr(1),
        },
    },
    pulumi.DependsOn([]pulumi.Resource{nackController}),
)
```

### Version Management

```mermaid
flowchart TB
    subgraph Versions["Version Sources"]
        proto["spec.proto defaults"]
        vars["vars.go defaults"]
    end
    
    subgraph URLs["Resource URLs"]
        helm_url["Helm: nats-io.github.io/k8s"]
        crd_url["CRDs: raw.githubusercontent.com/nats-io/nack/vX.Y.Z/deploy/crds.yml"]
    end
    
    proto -->|"chart: 0.31.1"| helm_url
    proto -->|"app: 0.21.1"| crd_url
```

**Critical Fix**: Initial implementation used chart version for CRD URL, but CRDs are tagged by app version. This caused silent failures since `v0.31.1` doesn't exist as a GitHub tag.

## YAML Example

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesNats
metadata:
  name: nats-gcp-dev
spec:
  namespace:
    value: planton-gcp-dev
  nack_controller:
    enabled: true
    enable_control_loop: true
  streams:
    - name: api-resources
      subjects:
        - "api.resources.>"
      replicas: 1
      max_age: 5m
      storage: file
      retention: limits
    - name: webhooks
      subjects:
        - "webhooks.>"
      replicas: 1
      storage: file
```

## Benefits

### For Users

- **Single manifest**: Define NATS + NACK + Streams in one YAML file
- **Validated configuration**: Proto validation catches errors before deployment
- **Version control**: Infrastructure-as-code for complete NATS ecosystem
- **Clean YAML**: Enum values like `file`, `memory`, `limits` map naturally

### For Developers

- **Type-safe Pulumi code**: Compile-time checks via generated Go types
- **Clear dependency chain**: Explicit ordering prevents race conditions
- **Configurable versions**: Users can pin chart/app versions as needed

### Deployment Reliability

```mermaid
flowchart TB
    subgraph Before["Before: Race Condition Risk"]
        A1[Deploy CRDs] --> B1[Deploy CRs]
        B1 -.->|"May fail: CRD not ready"| X1[Error]
    end
    
    subgraph After["After: Explicit Dependencies"]
        A2[NATS Helm] --> B2[NACK CRDs]
        B2 --> C2[NACK Controller]
        C2 --> D2[Stream CRs]
        D2 --> E2[Consumer CRs]
    end
```

## Impact

### Components Affected

| Component | Changes |
|-----------|---------|
| `spec.proto` | +70 lines (NACK, Stream, Consumer messages, enums) |
| `vars.go` | +10 lines (NACK constants, CRD URL template) |
| `locals.go` | +15 lines (version management, URL derivation) |
| `nack_crds.go` | New file (~45 lines) |
| `nack_controller.go` | New file (~60 lines) |
| `streams.go` | New file (~280 lines) |
| `outputs.go` | +5 lines (new output constants) |
| `main.go` | +20 lines (4-step deployment orchestration) |
| `kubernetestypes/Makefile` | +5 lines (gen-nack target) |

### Generated Code

- `nack/kubernetes/jetstream/v1beta2/` - ~19K lines of generated Go types
- Proto-generated Go: Updated with new messages and enums

## Related Work

- **ChatGPT Consultation**: Recommendations on CRD/CR deployment ordering to avoid race conditions
- **NACK Repository**: `https://github.com/nats-io/nack` - Official NATS Kubernetes controller
- **Helm Charts**: `https://nats-io.github.io/k8s/` - NATS and NACK Helm charts

## Testing

Validated deployment on GKE cluster:

```
Resources:
    + 14 created (6 CRDs + 8 Streams)
    15 unchanged (existing NATS resources)

Duration: 29s
```

All 8 streams successfully created and managed by NACK controller.

---

**Status**: ✅ Production Ready
**Timeline**: ~3 hours implementation + debugging
