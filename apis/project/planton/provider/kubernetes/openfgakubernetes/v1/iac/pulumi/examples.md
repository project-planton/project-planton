Here are a few example configurations for the `OpenFgaKubernetes` API resource based on the similar format provided. Since the specification in your case has specific fields like `container`, `datastore`, and `ingress`, I have created examples that demonstrate different configurations for deploying an OpenFGA service on Kubernetes.

---

# Example with Ingress Enabled

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: open-fga-service
spec:
  kubernetes_cluster_credential_id: my-k8s-cluster-credential
  container:
    replicas: 3
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
  ingress:
    isEnabled: true
    host: openfga.mycluster.example.com
    path: /openfga
  datastore:
    engine: postgres
    uri: postgres://user:password@db-host:5432/openfga
```

---

# Example with Ingress Disabled and MySQL Datastore

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: open-fga-service
spec:
  kubernetes_cluster_credential_id: another-k8s-cluster-credential
  container:
    replicas: 2
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 512Mi
  ingress:
    isEnabled: false
  datastore:
    engine: mysql
    uri: mysql://user:password@mysql-db:3306/openfga
```

---

# Example with Minimum Required Fields

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: basic-openfga
spec:
  kubernetes_cluster_credential_id: my-cluster-credential
  container:
    replicas: 1
    resources:
      requests:
        cpu: 50m
        memory: 128Mi
      limits:
        cpu: 200m
        memory: 512Mi
  datastore:
    engine: postgres
    uri: postgres://user:password@db-host:5432/openfga
```

---

# Example with High Availability Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: OpenFgaKubernetes
metadata:
  name: open-fga-high-availability
spec:
  kubernetes_cluster_credential_id: high-availability-k8s-credential
  container:
    replicas: 5
    resources:
      requests:
        cpu: 500m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 4Gi
  ingress:
    isEnabled: true
    host: open-fga-ha.example.com
    path: /openfga
  datastore:
    engine: postgres
    uri: postgres://user:securepassword@ha-db-host:5432/openfga
```
