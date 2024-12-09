# MongoDB Kubernetes API - Example Configurations

## Example w/ Basic Configuration

### Create Using CLI

Create a YAML file using the example shown below. After the YAML is created, use the following command to apply it:

```shell
planton apply -f <yaml-path>
```

### Basic Example

This basic example demonstrates a minimal configuration for deploying a MongoDB Kubernetes instance using the default settings, including 1 replica and no persistence.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbKubernetes
metadata:
  name: basic-mongodb
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 1Gi
```

---

## Example w/ Persistence Enabled

In this example, MongoDB persistence is enabled, and a persistent volume is created for each MongoDB pod to ensure data durability. The `disk_size` field defines the storage size allocated to the MongoDB pods.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbKubernetes
metadata:
  name: persistent-mongodb
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 3
    isPersistenceEnabled: true
    diskSize: 10Gi
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 1Gi
```

---

## Example w/ Custom Helm Values

This example demonstrates how to customize the MongoDB deployment using Helm chart values. In this case, we use `helm_values` to set specific resource limits and other options available in the Helm chart for MongoDB.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbKubernetes
metadata:
  name: custom-mongodb
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 2
    isPersistenceEnabled: true
    diskSize: 20Gi
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 2Gi
  helmValues:
    mongodbUsername: myuser
    mongodbDatabase: mydatabase
    mongodbRootPassword: secretpassword
```

---

## Example w/ Ingress Enabled

In this example, ingress is enabled to allow external access to the MongoDB service. This is particularly useful when MongoDB needs to be accessed by clients outside the Kubernetes cluster.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbKubernetes
metadata:
  name: ingress-mongodb
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: true
    diskSize: 5Gi
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 1Gi
  ingress:
    isEnabled: true
```

---

## Example w/ Random Password and Kubernetes Secrets

This example demonstrates how to automatically generate a random password for MongoDB using Kubernetes secrets. The password is securely stored in the Kubernetes secret and used for MongoDB authentication.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbKubernetes
metadata:
  name: secret-mongodb
spec:
  kubernetesClusterCredentialId: my-cluster-credential-id
  container:
    replicas: 1
    isPersistenceEnabled: false
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1000m
        memory: 1Gi
  helmValues:
    mongodbRootPassword: ${kubernetes-secret-root-password}
```
