# AltinityOperatorKubernetes API Resource Examples

Below are examples demonstrating how to configure and deploy the `AltinityOperatorKubernetes` API resource using various specifications. Follow the instructions to create and apply each YAML configuration using the Planton CLI.

---

## Example 1: Basic Operator Deployment

### Description

This example demonstrates a basic deployment of the Altinity ClickHouse Operator within a Kubernetes cluster. It uses the default resource allocations suitable for most production environments.

### Create and Apply

1. **Create a YAML file** using the example below.
2. **Apply the configuration** using the following command:

    ```shell
    planton pulumi up --manifest <yaml-path>
    ```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: AltinityOperatorKubernetes
metadata:
  name: altinity-operator-prod
spec:
  targetCluster:
    credentialId: my-k8s-cluster-credential
  namespace: altinity-operator  # Optional: defaults to "altinity-operator"
  container:
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 1Gi
```

---

## Example 2: Operator Deployment with Custom Resources

### Description

This example illustrates how to deploy the Altinity ClickHouse Operator with customized resource allocations. This configuration provides more resources to the operator, which may be beneficial in large clusters managing many ClickHouse instances.

### Create and Apply

1. **Create a YAML file** using the example below.
2. **Apply the configuration** using the following command:

    ```shell
    planton pulumi up --manifest <yaml-path>
    ```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: AltinityOperatorKubernetes
metadata:
  name: altinity-operator-large
spec:
  targetCluster:
    credentialId: my-k8s-cluster-credential
  namespace: altinity-operator
  container:
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 2Gi
```

---

## Example 3: Development Environment Deployment

### Description

This example shows a minimal resource configuration suitable for development or testing environments where resource constraints are tighter.

### Create and Apply

1. **Create a YAML file** using the example below.
2. **Apply the configuration** using the following command:

    ```shell
    planton pulumi up --manifest <yaml-path>
    ```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: AltinityOperatorKubernetes
metadata:
  name: altinity-operator-dev
spec:
  targetCluster:
    credentialId: my-local-k8s-cluster
  namespace: altinity-operator-dev  # Custom namespace for dev environment
  container:
    resources:
      requests:
        cpu: 50m
        memory: 128Mi
      limits:
        cpu: 500m
        memory: 512Mi
```

---

## Verifying the Deployment

After deploying the operator, you can verify its status using standard Kubernetes commands:

```shell
# Check operator pod status
kubectl get pods -n altinity-operator

# Check operator logs
kubectl logs -n altinity-operator -l app.kubernetes.io/name=altinity-clickhouse-operator

# Verify CRDs are installed
kubectl get crds | grep clickhouse

# Expected CRDs:
# - clickhouseinstallations.clickhouse.altinity.com
# - clickhouseoperatorconfigurations.clickhouse.altinity.com
```

## Next Steps

Once the operator is deployed, you can:

1. Deploy ClickHouse clusters using the `ClickHouseKubernetes` workload resource
2. Create custom ClickHouseInstallation resources directly
3. Monitor the operator's health and logs
4. Scale your ClickHouse deployments as needed

