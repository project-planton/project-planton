
# Example 1: Basic Helm Release Deployment

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: HelmRelease
metadata:
  name: my-helm-release
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  helm_chart:
    repo: https://charts.bitnami.com/bitnami
    name: nginx
    version: 8.5.0
    values: {}
```

---

# Example 2: Helm Release with Custom Values

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: HelmRelease
metadata:
  name: my-nginx-helm-release
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  helm_chart:
    repo: https://charts.bitnami.com/bitnami
    name: nginx
    version: 8.5.0
    values:
      service:
        type: LoadBalancer
      replicaCount: "3"
```

---

# Example 3: Helm Release with Ingress Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: HelmRelease
metadata:
  name: my-app-helm-release
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  helm_chart:
    repo: https://charts.bitnami.com/bitnami
    name: wordpress
    version: 10.0.0
    values:
      ingress:
        enabled: true
        hostname: wordpress.example.com
```

---

# Example 4: Helm Release with Multiple Values

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: HelmRelease
metadata:
  name: multi-value-helm-release
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  helm_chart:
    repo: https://charts.bitnami.com/bitnami
    name: redis
    version: 14.8.12
    values:
      usePassword: "true"
      sentinel:
        enabled: "true"
      metrics:
        enabled: "true"
```

---

# Example 5: Minimal Helm Release (Empty Spec)

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: HelmRelease
metadata:
  name: minimal-helm-release
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  helm_chart:
    repo: https://charts.bitnami.com/bitnami
    name: nginx
    version: 8.5.0
    values: {}
```
