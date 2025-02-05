# Create using CLI

Create a YAML file using the examples shown below. After the YAML is created, use the following command to apply:

```shell
platon apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: example-gke-cluster
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
  kubernetesAddons:
    isInstallIngressNginx: true
  ingressDnsDomains:
    - name: example.com
      isTlsEnabled: true
      dnsZoneGcpProjectId: example-gcp-project
```

This basic example creates a GKE cluster in the `us-central1` region with a single node pool. It installs the Ingress
Nginx controller and sets up an ingress DNS domain with TLS enabled.

# Example with Cluster Autoscaling

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
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
  kubernetesAddons:
    isInstallIstio: true
```

In this example, cluster autoscaling is enabled with specified CPU and memory limits. The cluster is set up in the
`europe-west1` region, and Istio is installed as a service mesh.

# Example with Multiple Node Pools

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
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

This example demonstrates a cluster with multiple node pools, including a pool with spot instances. It installs Cert
Manager for certificate management.

# Example with Kubernetes Addons

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: addons-enabled-cluster
spec:
  billingAccountId: 0123AB-4567CD-89EFGH
  gcpCredentialId: gcpcred-addons-credential
  region: us-east1
  zone: us-east1-b
  nodePools:
    - name: default-pool
      machineType: n1-standard-4
      minNodeCount: 1
      maxNodeCount: 3
  kubernetesAddons:
    isInstallIngressNginx: true
    isInstallIstio: true
    isInstallCertManager: true
    isInstallExternalDns: true
    isInstallExternalSecrets: true
    isInstallKafkaOperator: true
    isInstallPostgresOperator: true
    isInstallSolrOperator: true
    isInstallElasticOperator: true
  ingressDnsDomains:
    - name: example.org
      isTlsEnabled: true
      dnsZoneGcpProjectId: dns-project-id
```

This cluster installs a comprehensive set of Kubernetes addons and operators, providing a robust environment for
deploying various applications and services.

# Example with Shared VPC

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: shared-vpc-cluster
spec:
  billingAccountId: 4567CD-89EFGH-0123AB
  gcpCredentialId: gcpcred-sharedvpc-credential
  region: us-west1
  zone: us-west1-a
  isCreateSharedVpc: true
  nodePools:
    - name: default-pool
      machineType: n1-standard-4
      minNodeCount: 1
      maxNodeCount: 3
  kubernetesAddons:
    isInstallIstio: true
  ingressDnsDomains:
    - name: shared.example.com
      isTlsEnabled: true
      dnsZoneGcpProjectId: shared-vpc-project-id
```

In this example, the cluster is configured to use a shared VPC network, allowing for network segmentation and resource
isolation across projects.

# Example with Workload Logs Enabled

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: logging-enabled-cluster
spec:
  billingAccountId: 89EFGH-0123AB-4567CD
  gcpCredentialId: gcpcred-logging-credential
  region: us-central1
  zone: us-central1-f
  isWorkloadLogsEnabled: true
  nodePools:
    - name: logging-pool
      machineType: n1-standard-4
      minNodeCount: 1
      maxNodeCount: 3
```

This cluster has workload logging enabled, which sends logs from Kubernetes pods to Google Cloud Logging. Be aware that
enabling log forwarding may increase cloud billing costs.

# Example with Custom Cluster Autoscaling and Node Pools

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: custom-autoscaling-cluster
spec:
  billingAccountId: 0123AB-4567CD-89EFGH
  gcpCredentialId: gcpcred-custom-autoscaling
  region: australia-southeast1
  zone: australia-southeast1-b
  clusterAutoscalingConfig:
    isEnabled: true
    cpuMinCores: 2
    cpuMaxCores: 32
    memoryMinGb: 4
    memoryMaxGb: 128
  nodePools:
    - name: cpu-intensive-pool
      machineType: n2-highcpu-8
      minNodeCount: 0
      maxNodeCount: 10
    - name: memory-intensive-pool
      machineType: n2-highmem-16
      minNodeCount: 0
      maxNodeCount: 5
  kubernetesAddons:
    isInstallIngressNginx: true
    isInstallCertManager: true
  ingressDnsDomains:
    - name: custom.example.net
      isTlsEnabled: true
      dnsZoneGcpProjectId: custom-dns-project
```

This example showcases a cluster with custom autoscaling settings and specialized node pools for different workloads.

# Example with Ingress DNS Domains without TLS

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: no-tls-cluster
spec:
  billingAccountId: 4567CD-89EFGH-0123AB
  gcpCredentialId: gcpcred-notls-credential
  region: europe-north1
  zone: europe-north1-a
  nodePools:
    - name: default-pool
      machineType: n1-standard-4
      minNodeCount: 1
      maxNodeCount: 3
  ingressDnsDomains:
    - name: notls.example.com
      isTlsEnabled: false
      dnsZoneGcpProjectId: notls-dns-project
```

In this cluster, the ingress DNS domain is configured without TLS enabled. Note that services requiring TLS will not
function with this domain configuration.

# Example with Custom Machine Types and Labels

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: custom-machinetype-cluster
spec:
  billingAccountId: 89EFGH-0123AB-4567CD
  gcpCredentialId: gcpcred-custom-machinetype
  region: southamerica-east1
  zone: southamerica-east1-b
  nodePools:
    - name: custom-pool
      machineType: n2-custom-8-32768
      minNodeCount: 1
      maxNodeCount: 4
  kubernetesAddons:
    isInstallElasticOperator: true
```

This example uses a custom machine type for the node pool, providing 8 vCPUs and 32 GB of memory per node. The Elastic
Operator is installed for managing Elasticsearch clusters.

# Example with Environment Information

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: env-info-cluster
spec:
  billingAccountId: 0123AB-4567CD-89EFGH
  gcpCredentialId: gcpcred-envinfo-credential
  region: us-central1
  zone: us-central1-a
  environmentInfo:
    orgId: example-org
    envId: prod
    envName: production
  nodePools:
    - name: default-pool
      machineType: n1-standard-4
      minNodeCount: 1
      maxNodeCount: 3
```

In this cluster, environment information is provided, which can be useful for labeling and organizing resources
according to organizational structures.

# Example with Stack Job Settings

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GkeCluster
metadata:
  name: stackjob-cluster
spec:
  billingAccountId: 4567CD-89EFGH-0123AB
  gcpCredentialId: gcpcred-stackjob-credential
  region: asia-southeast1
  zone: asia-southeast1-b
  stackJobSettings:
    pulumiBackendCredentialId: pulcred-example-backend
    stackJobRunnerId: sjr-example-runner
  nodePools:
    - name: default-pool
      machineType: e2-medium
      minNodeCount: 1
      maxNodeCount: 2
```

This example includes stack job settings for customizing Pulumi backend credentials and job runner configurations.
