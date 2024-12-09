Here are a few example configurations for the `OpenfgaKubernetes` API resource based on the similar format provided. Since the specification in your case has specific fields like `container`, `datastore`, and `ingress`, I have created examples that demonstrate different configurations for deploying an OpenFGA service on Kubernetes.

---

# Example with Ingress Enabled

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: OpenfgaKubernetes
metadata:
  name: openfga-service
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
apiVersion: code2cloud.planton.cloud/v1
kind: OpenfgaKubernetes
metadata:
  name: openfga-service
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
apiVersion: code2cloud.planton.cloud/v1
kind: OpenfgaKubernetes
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
apiVersion: code2cloud.planton.cloud/v1
kind: OpenfgaKubernetes
metadata:
  name: openfga-high-availability
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
    host: openfga-ha.example.com
    path: /openfga
  datastore:
    engine: postgres
    uri: postgres://user:securepassword@ha-db-host:5432/openfga
```
