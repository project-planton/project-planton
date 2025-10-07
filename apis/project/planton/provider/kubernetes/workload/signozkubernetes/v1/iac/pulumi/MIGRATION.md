# Bitnami Registry Migration Guide

## Issue

Bitnami discontinued free Docker Hub images on September 29, 2025. Existing SigNoz deployments may fail with:
```
ErrImagePull: docker.io/bitnami/clickhouse:XX.X.X-debian-12-rX: not found
ErrImagePull: docker.io/bitnami/zookeeper:X.X.X-debian-12-rXX: not found
```

## Solution

This module now uses `docker.io/bitnamilegacy` registry by default for ClickHouse and ZooKeeper dependencies.

### For Existing Deployments

**Option 1: Destroy and Recreate (Recommended)**

This ensures a clean deployment with the new image registry settings:

```bash
# 1. Destroy the existing stack
project-planton pulumi destroy --manifest signoz.yaml \
  --module-dir apis/project/planton/provider/kubernetes/workload/signozkubernetes/v1/iac/pulumi

# 2. Recreate with the new settings
project-planton pulumi up --manifest signoz.yaml \
  --module-dir apis/project/planton/provider/kubernetes/workload/signozkubernetes/v1/iac/pulumi
```

⚠️ **Warning**: This will delete your SigNoz telemetry data unless you have external backups!

**Option 2: Manual Override via Helm Values**

If you can't destroy the stack, override the images directly in your manifest:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: SignozKubernetes
metadata:
  name: main
  org: planton-cloud
  env: app-prod
spec:
  signozContainer:
    replicas: 1
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 1000m
        memory: 2Gi
  otelCollectorContainer:
    replicas: 2
    resources:
      requests:
        cpu: 500m
        memory: 1Gi
      limits:
        cpu: 2000m
        memory: 4Gi
  database:
    isExternal: false
    managedDatabase:
      container:
        replicas: 1
        isPersistenceEnabled: true
        diskSize: 20Gi
        resources:
          requests:
            cpu: 500m
            memory: 1Gi
          limits:
            cpu: 2000m
            memory: 4Gi
      cluster:
        isEnabled: false
      zookeeper:
        isEnabled: false
  helmValues:
    global.imageRegistry: "docker.io/bitnamilegacy"
    clickhouse.image.repository: "clickhouse"
    zookeeper.image.repository: "zookeeper"
```

Then run:
```bash
project-planton pulumi up --manifest signoz.yaml \
  --module-dir apis/project/planton/provider/kubernetes/workload/signozkubernetes/v1/iac/pulumi
```

### For New Deployments

New deployments automatically use the `bitnamilegacy` registry. No additional configuration needed.

## Long-term Alternatives

### For ClickHouse Dependency

1. **Use Official ClickHouse Images** (Recommended):
   ```yaml
   helmValues:
     clickhouse.image.registry: "docker.io"
     clickhouse.image.repository: "clickhouse/clickhouse-server"
     clickhouse.image.tag: "24.8"
   ```

2. **Altinity ClickHouse Operator**: Consider using the Altinity ClickHouse Operator which uses official images
   - https://github.com/Altinity/clickhouse-operator

3. **External ClickHouse**: Use the external database configuration to connect to a separately managed ClickHouse instance

### For ZooKeeper Dependency

1. **Use Official ZooKeeper Images**:
   ```yaml
   helmValues:
     zookeeper.image.registry: "docker.io"
     zookeeper.image.repository: "zookeeper"
     zookeeper.image.tag: "3.8"
   ```

2. **External ZooKeeper**: Deploy ZooKeeper separately using the official operator or helm chart

### SigNoz-Specific Considerations

Since SigNoz Helm chart manages ClickHouse as a dependency, the image registry override applies to the embedded ClickHouse and ZooKeeper deployments. If you're using external ClickHouse (`spec.database.isExternal: true`), this issue doesn't affect you.

## References

- Bitnami Registry Changes: https://github.com/bitnami/containers/issues/83267
- SigNoz Helm Chart: https://github.com/SigNoz/charts
- ClickHouse Official Images: https://hub.docker.com/r/clickhouse/clickhouse-server
- ZooKeeper Official Images: https://hub.docker.com/_/zookeeper
- Bitnami Legacy Registry: https://hub.docker.com/u/bitnamilegacy

