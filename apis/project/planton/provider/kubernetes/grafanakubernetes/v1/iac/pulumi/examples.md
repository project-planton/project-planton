
# Example 1: Basic Grafana Deployment

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GrafanaKubernetes
metadata:
  name: grafana-instance
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  container:
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 500m
        memory: 512Mi
```

---

# Example 2: Grafana with Ingress Enabled

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GrafanaKubernetes
metadata:
  name: grafana-prod
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  container:
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 1
        memory: 1Gi
  ingress:
    enabled: true
    hostname: grafana.example.com
```

---

# Example 3: Grafana Deployment with Environment Variables

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GrafanaKubernetes
metadata:
  name: grafana-with-env
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  container:
    env:
      variables:
        GF_SECURITY_ADMIN_PASSWORD: securepassword
        GF_SECURITY_ADMIN_USER: admin
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
```

---

# Example 4: Grafana with Environment Secrets

The below example assumes that secrets are managed by Planton Cloud’s [GCP Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GrafanaKubernetes
metadata:
  name: grafana-secure
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  container:
    env:
      secrets:
        GF_SECURITY_ADMIN_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.admin-password}
      variables:
        GF_SECURITY_ADMIN_USER: admin
    resources:
      requests:
        cpu: 100m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
```

---

# Example 5: Minimal Grafana Deployment (Empty Spec)

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: GrafanaKubernetes
metadata:
  name: minimal-grafana
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
```
