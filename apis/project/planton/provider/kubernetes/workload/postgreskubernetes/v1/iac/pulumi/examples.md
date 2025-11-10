Here are a few examples for the `PostgresKubernetes` API resource, following the same format as provided:

---

# Basic Example

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: postgres-db
spec:
  kubernetes_cluster_credential_id: my-k8s-cluster-credential
  container:
    replicas: 1
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
    diskSize: 10Gi
  ingress:
    enabled: true
    hostname: postgres-db.example.com
```

---

# Example with Environment Variables

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: postgres-db
spec:
  kubernetes_cluster_credential_id: prod-k8s-cluster-credential
  container:
    replicas: 2
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 2Gi
    diskSize: 20Gi
  ingress:
    enabled: true
    hostname: postgres-prod.example.com
  env:
    variables:
      DATABASE_USER: admin
      DATABASE_NAME: prod-db
```

---

# Example with Secrets for Sensitive Information

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: secure-postgres-db
spec:
  kubernetes_cluster_credential_id: secure-k8s-cluster-credential
  container:
    replicas: 3
    resources:
      requests:
        cpu: 300m
        memory: 1Gi
      limits:
        cpu: 4000m
        memory: 4Gi
    diskSize: 50Gi
  ingress:
    enabled: false
  env:
    secrets:
      DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.db-password}
    variables:
      DATABASE_NAME: secure-db
```

---

# Example with Minimum Required Fields

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: minimal-postgres-db
spec:
  kubernetes_cluster_credential_id: basic-k8s-cluster-credential
  container:
    replicas: 1
    resources:
      requests:
        cpu: 50m
        memory: 128Mi
      limits:
        cpu: 500m
        memory: 512Mi
    diskSize: 5Gi
```

---

These examples demonstrate varying configurations for deploying PostgreSQL on Kubernetes using the `PostgresKubernetes` API resource. Each example includes different use cases such as handling environment variables, secrets, basic deployment, and minimum configuration setups for flexibility.