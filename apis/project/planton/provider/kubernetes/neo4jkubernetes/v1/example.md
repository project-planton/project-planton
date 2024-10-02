# Neo4j Kubernetes API - Example Configurations

## Example w/ Basic Configuration

### Create Using CLI

Create a YAML file using the example shown below. After the YAML is created, use the following command to apply it:

```shell
planton apply -f <yaml-path>
```

### Basic Example

This is a basic configuration for deploying a Neo4j instance on Kubernetes. It specifies the necessary resources like CPU and memory, along with the Kubernetes cluster credential.

```yaml
apiVersion: kubernetes.project.planton/v1
kind: Neo4jKubernetes
metadata:
  name: basic-neo4j
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 1000m
        memory: 2Gi
```

---

## Example w/ Ingress Enabled

In this example, ingress is enabled to allow external access to the Neo4j instance. This setup is particularly useful if the Neo4j database needs to be accessed from outside the Kubernetes cluster.

```yaml
apiVersion: kubernetes.project.planton/v1
kind: Neo4jKubernetes
metadata:
  name: ingress-neo4j
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    resources:
      requests:
        cpu: 200m
        memory: 1Gi
      limits:
        cpu: 1500m
        memory: 3Gi
  ingress:
    isEnabled: true
```

---

## Example w/ Custom Resource Limits

This example demonstrates how to configure custom resource limits for the Neo4j instance, tailoring the deployment to specific performance needs.

```yaml
apiVersion: kubernetes.project.planton/v1
kind: Neo4jKubernetes
metadata:
  name: custom-resources-neo4j
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    resources:
      requests:
        cpu: 250m
        memory: 1Gi
      limits:
        cpu: 2000m
        memory: 4Gi
```

---

## Example w/ Minimal Configuration

This minimal example specifies only the mandatory fields for deploying a Neo4j instance on Kubernetes. Although functional, additional configurations like ingress and persistence can be added as needed.

```yaml
apiVersion: kubernetes.project.planton/v1
kind: Neo4jKubernetes
metadata:
  name: minimal-neo4j
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 1Gi
```
