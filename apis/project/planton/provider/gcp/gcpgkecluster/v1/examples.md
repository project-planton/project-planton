# Create using CLI

Create a YAML file using the examples shown below. After the YAML is created, use the following command to apply:

```shell
platon apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpGkeCluster
metadata:
  name: example-gcp-gke-cluster
spec:
  billingAccountId: 0123AB-4567CD-89EFGH
  gcpCredentialId: gcpcred-example-credential
  region: us-central1
  zone: us-central1-a
  nodePools:
    - name: default-pool
      machineType: n1-standard-4
      minNodeCount: 1
      maxNodeCount: 3
```

This basic example creates a GKE cluster in the `us-central1` region with a single node pool. It installs the Ingress
Nginx controller and sets up an ingress DNS domain with TLS enabled.

# Example with Cluster Autoscaling

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpGkeCluster
metadata:
  name: autoscaling-cluster
spec:
  billingAccountId: 4567CD-89EFGH-0123AB
  gcpCredentialId: gcpcred-autoscaling-credential
  region: europe-west1
  zone: europe-west1-b
  clusterAutoscalingConfig:
    isEnabled: true
    cpuMinCores: 4
    cpuMaxCores: 16
    memoryMinGb: 8
    memoryMaxGb: 64
  nodePools:
    - name: autoscaling-pool
      machineType: e2-standard-4
      minNodeCount: 1
      maxNodeCount: 5
```

In this example, cluster autoscaling is enabled with specified CPU and memory limits. The cluster is set up in the
`europe-west1` region, and Istio is installed as a service mesh.

# Example with Multiple Node Pools

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpGkeCluster
metadata:
  name: multi-nodepool-cluster
spec:
  billingAccountId: 89EFGH-0123AB-4567CD
  gcpCredentialId: gcpcred-multinode-credential
  region: asia-east1
  zone: asia-east1-a
  nodePools:
    - name: general-pool
      machineType: n1-standard-4
      minNodeCount: 1
      maxNodeCount: 3
    - name: high-memory-pool
      machineType: n1-highmem-8
      minNodeCount: 0
      maxNodeCount: 2
    - name: spot-pool
      machineType: n1-standard-2
      minNodeCount: 0
      maxNodeCount: 5
      isSpotEnabled: true
  kubernetesAddons:
    isInstallCertManager: true
```
