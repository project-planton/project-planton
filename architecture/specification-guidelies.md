# Specification Design Guidelines

## The Philosophy: Deployment-Agnostic Specifications

When writing `spec.proto` files for Planton Cloud resources, we follow a fundamental principle: **specifications should describe what users want to deploy, not how it gets deployed**.

Think of it like ordering food at a restaurant. You tell the waiter "I'd like a medium-rare steak with mashed potatoes," not "Please cook beef at 135°F for 4 minutes per side, boil potatoes, mash them with butter..." The kitchen (deployment layer) handles the implementation details.

## The 80/20 Rule (Pareto's Principle)

Our specifications capture the **20% of configurations that 80% of users need**. We deliberately avoid exposing every knob and dial of the underlying deployment tool.

### Why This Matters

Consider MongoDB deployment options:
- Bitnami Helm chart has 200+ configuration values
- Percona operator CRD has 150+ fields
- Raw Kubernetes manifests have unlimited possibilities

Our `MongodbKubernetesSpec` has just 3 fields:
- `container` - replicas, resources, persistence, disk size
- `ingress` - external access configuration
- `helm_values` - escape hatch for advanced users

This covers the vast majority of use cases while keeping the API intuitive and maintainable.

## Deployment-Agnostic Design

Specifications should **never** be coupled to a specific deployment tool. The same spec should work whether we deploy using:
- Helm charts (Bitnami, Percona)
- Kubernetes operators (Percona, Zalando)
- Raw Kubernetes YAML manifests
- Cloud-managed services (in the future)

### Real-World Example

The `MongodbKubernetes` specification was originally implemented with Bitnami Helm charts. When we migrated to Percona operator CRDs, **we didn't change a single line in spec.proto**. The mapping worked seamlessly:

```
spec.container.replicas        → Percona replica set size
spec.container.resources       → Percona pod resources
spec.container.disk_size       → Percona storage requests
spec.ingress                   → LoadBalancer service (unchanged)
```

This is deployment-agnostic design in action.

## Intuitive User Experience

Specifications should feel natural for engineers deploying resources on Kubernetes. Use terminology and concepts that match how users think about the technology.

### Good Examples

**MongoDB Container Configuration:**
```protobuf
message MongodbKubernetesContainer {
  int32 replicas = 1;
  ContainerResources resources = 2;
  bool is_persistence_enabled = 3;
  string disk_size = 4;
}
```

This mirrors how engineers think: "I want 3 replicas with 2GB memory and 10GB disk storage."

**ClickHouse Coordination (Modern):**
```protobuf
message ClickHouseKubernetesCoordinationConfig {
  CoordinationType type = 1;  // keeper, external_keeper, external_zookeeper
  ClickHouseKubernetesKeeperConfig keeper_config = 2;
  ClickHouseKubernetesExternalConfig external_config = 3;
}
```

Users specify the coordination type they want, not the implementation details.

### Bad Examples (What to Avoid)

❌ Exposing Helm chart structure:
```protobuf
message BadSpec {
  map<string, HelmValue> helm_values = 1;  // Forces users to know Helm
}
```

❌ Tight coupling to specific tools:
```protobuf
message BadSpec {
  BitnamisMongoDBChartConfig bitnami = 1;  // Locks us into Bitnami
}
```

❌ Over-configuration:
```protobuf
message BadSpec {
  string sysctl_vm_swappiness = 1;
  string kernel_transparent_hugepage = 2;
  int32 mongodb_net_max_incoming_connections = 3;
  // ... 200 more fields
}
```

## Future Flexibility

Well-designed specifications enable future enhancements without breaking changes.

### Versioning Strategy

When specifications need evolution, use optional fields and sane defaults:

```protobuf
message MongodbKubernetesSpec {
  MongodbKubernetesContainer container = 1;
  
  // Added later without breaking existing users
  MongodbKubernetesBackupConfig backup = 4;
  MongodbKubernetesMonitoringConfig monitoring = 5;
}
```

Existing deployments continue to work. New features are opt-in.

### Deprecation Pattern

When replacing fields, keep old ones and document the migration path:

```protobuf
message ClickHouseKubernetesSpec {
  // DEPRECATED: Use coordination field instead
  // Will be removed in v2.0
  ClickHouseKubernetesZookeeperConfig zookeeper = 3 [deprecated = true];
  
  // New field with better design
  ClickHouseKubernetesCoordinationConfig coordination = 6;
}
```

Implementation code handles both fields with priority logic (new field takes precedence).

## Practical Guidelines

### 1. Start with User Intent

Ask: "What does a user want to accomplish?" Not: "What does the deployment tool expose?"

**User intent:** "Deploy a highly available MongoDB cluster with 10GB storage"  
**Good spec:** `replicas: 3, disk_size: "10Gi"`  
**Bad spec:** `helm_values: {"architecture": "replicaset", "persistence.size": "10Gi", ...}`

### 2. Use Sensible Defaults

Every field should have a default value that works for 80% of cases:

```protobuf
MongodbKubernetesContainer container = 1 [
  (default_container) = {
    replicas: 1,
    resources: {
      limits { cpu: "1000m", memory: "1Gi" }
      requests { cpu: "50m", memory: "100Mi" }
    },
    is_persistence_enabled: true,
    disk_size: "1Gi"
  }
];
```

Users can deploy MongoDB with just:
```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MongodbKubernetes
metadata:
  name: my-db
spec: {}  # All defaults work!
```

### 3. Provide Escape Hatches

For the 20% of users with special needs, provide escape hatches:

```protobuf
// Advanced users can override specific Helm values
map<string, string> helm_values = 3;
```

This doesn't compromise the core design but gives power users flexibility.

### 4. Document Examples, Not Implementation

README files should show what users want to achieve, not how it's implemented:

**Good:** "Deploy a production MongoDB cluster with backup enabled"  
**Bad:** "Configure Percona operator CRD with specific replset parameters"

## Validation

When you write a spec, ask yourself:

1. ✅ **Deployment-agnostic?** Could this spec work with a different deployment tool?
2. ✅ **80/20 compliant?** Does it cover common use cases without overwhelming users?
3. ✅ **Intuitive?** Would a Kubernetes engineer understand it without reading documentation?
4. ✅ **Future-proof?** Can we add features without breaking changes?
5. ✅ **Sensible defaults?** Can users deploy with minimal configuration?

If you answer "no" to any question, revisit the design.

## Summary

Great specifications are like great APIs: they hide complexity, anticipate user needs, and remain stable over time. By following these guidelines, we create specifications that serve users today and adapt to technologies tomorrow.

Remember: **Specify outcomes, not implementations. Focus on the what, not the how.**

