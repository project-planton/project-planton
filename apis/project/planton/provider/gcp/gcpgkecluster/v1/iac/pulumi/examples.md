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

In this example, cluster autoscaling is enabled with specified CPU and memory limits. The cluster is set up in the
`europe-west1` region.

# Example with Multiple Node Pools

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpGkeCluster
metadata:
  name: multi-nodepool-cluster
spec:
  clusterProjectId: gcp-project-id
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
```
