# PrometheusKubernetes API-Resource Examples

## Example 1: Basic Prometheus Kubernetes Deployment

This example demonstrates a simple Prometheus deployment in a Kubernetes cluster with one Prometheus pod and no persistence enabled. The resources are set to reasonable defaults for basic monitoring purposes.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PrometheusKubernetes
metadata:
  name: prometheus-instance-basic
spec:
  kubernetesClusterCredentialId: cluster-credential-12345
  container:
    replicas: 1
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
```

## Example 2: Prometheus with Persistence Enabled

In this example, Prometheus is deployed with persistence enabled to ensure that monitoring data is retained across pod restarts. The disk size is configured to 5Gi for storing the Prometheus data.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PrometheusKubernetes
metadata:
  name: prometheus-instance-with-persistence
spec:
  kubernetesClusterCredentialId: cluster-credential-67890
  container:
    replicas: 2
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 2
        memory: 2Gi
    isPersistenceEnabled: true
    diskSize: "5Gi"
```

## Example 3: Prometheus with Ingress Configuration

This example deploys Prometheus with an ingress resource, allowing external access to Prometheus via a public URL. The ingress class and host are specified for routing traffic to the Prometheus service.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PrometheusKubernetes
metadata:
  name: prometheus-instance-with-ingress
spec:
  kubernetesClusterCredentialId: cluster-credential-24680
  container:
    replicas: 1
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 2
        memory: 2Gi
  ingress:
    isEnabled: true
    ingressClass: "nginx"
    host: "prometheus.example.com"
```

## Example 4: Prometheus with Custom Resource Limits and No Persistence

This configuration specifies custom CPU and memory limits for the Prometheus container, ensuring the monitoring system uses fewer cluster resources. Persistence is disabled for simplicity.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: PrometheusKubernetes
metadata:
  name: prometheus-instance-custom-limits
spec:
  kubernetesClusterCredentialId: cluster-credential-112233
  container:
    replicas: 3
    resources:
      requests:
        cpu: 200m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 1Gi
    isPersistenceEnabled: false
```
