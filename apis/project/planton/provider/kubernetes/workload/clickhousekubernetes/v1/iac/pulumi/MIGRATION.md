# Bitnami Registry Migration Guide

## Issue

Bitnami discontinued free Docker Hub images on September 29, 2025. Existing deployments fail with:
```
ErrImagePull: docker.io/bitnami/clickhouse:24.7.1-debian-12-r0: not found
ErrImagePull: docker.io/bitnami/zookeeper:3.8.4-debian-12-r11: not found
```

## Solution

This module now uses `docker.io/bitnamilegacy` registry by default.

### For Existing Deployments

**Option 1: Destroy and Recreate (Recommended)**

This ensures a clean deployment with the new image registry settings:

```bash
# 1. Destroy the existing stack
project-planton pulumi destroy --manifest clickhuose.yaml \
  --module-dir apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1/iac/pulumi

# 2. Recreate with the new settings
project-planton pulumi up --manifest clickhuose.yaml \
  --module-dir apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1/iac/pulumi
```

⚠️ **Warning**: This will delete your ClickHouse data unless you have external backups!

**Option 2: Manual Override via Helm Values**

If you can't destroy the stack, override the images directly in your manifest:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ClickhouseKubernetes
metadata:
  name: main
  org: planton-cloud
  env: app-prod
spec:
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: "50Gi"
    resources:
      limits:
        cpu: "2000m"
        memory: "4Gi"
      requests:
        cpu: "500m"
        memory: "1Gi"
  helmValues:
    global.imageRegistry: "docker.io/bitnamilegacy"
    image.registry: "docker.io"
    image.repository: "bitnamilegacy/clickhouse"
    zookeeper.image.registry: "docker.io"
    zookeeper.image.repository: "bitnamilegacy/zookeeper"
  ingress:
    enabled: true
    dnsDomain: planton.live
```

Then run:
```bash
project-planton pulumi up --manifest clickhuose.yaml \
  --module-dir apis/project/planton/provider/kubernetes/workload/clickhousekubernetes/v1/iac/pulumi
```

### For New Deployments

New deployments automatically use the `bitnamilegacy` registry. No additional configuration needed.

## Long-term Alternatives

1. **Use Official ClickHouse Images** (Recommended):
   ```yaml
   helmValues:
     image.registry: "docker.io"
     image.repository: "clickhouse/clickhouse-server"
     image.tag: "24.8"
   ```

2. **Altinity ClickHouse Operator**: Consider using the Altinity ClickHouse Operator which uses official images
   - https://github.com/Altinity/clickhouse-operator

3. **Custom Images**: Build your own from Bitnami's open-source code (Apache 2.0)

## References

- Bitnami Registry Changes: https://github.com/bitnami/containers/issues/83267
- ClickHouse Official Images: https://hub.docker.com/r/clickhouse/clickhouse-server
- Bitnami Legacy Registry: https://hub.docker.com/u/bitnamilegacy
