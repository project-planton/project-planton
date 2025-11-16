Here are a few examples for the `KubernetesArgocd` API resource. These examples demonstrate how to configure and deploy Argo CD on a Kubernetes cluster using Project Planton's unified API structure.

---

# Example 1: Basic Argo CD Deployment

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesArgocd
metadata:
  name: my-argocd
spec:
  container:
    resources:
      requests:
        cpu: 50m
        memory: 100Mi
      limits:
        cpu: 1000m
        memory: 1Gi
```

**Description:** A basic Argo CD deployment with default resource allocations.

---

# Example 2: Argo CD with Ingress Enabled

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesArgocd
metadata:
  name: argocd-prod
spec:
  container:
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 2000m
        memory: 2Gi
  ingress:
    is_enabled: true
    dns_domain: example.com
```

**Description:** Production Argo CD deployment with ingress enabled for external access. This will create external hostname at `argo-argocd-prod.example.com` and internal hostname at `argo-argocd-prod-internal.example.com`.

---

# Example 3: Argo CD Deployment with Custom Resources

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesArgocd
metadata:
  name: argocd-large
spec:
  container:
    resources:
      requests:
        cpu: 200m
        memory: 1Gi
      limits:
        cpu: 3000m
        memory: 4Gi
```

**Description:** Argo CD deployment with higher resource allocations for managing large-scale GitOps workflows with many applications.

---

# Example 4: Minimal Argo CD Deployment

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesArgocd
metadata:
  name: minimal-argocd
spec:
  container:
    resources:
      requests:
        cpu: 50m
        memory: 100Mi
      limits:
        cpu: 1000m
        memory: 1Gi
```

**Description:** Minimal Argo CD deployment using default container resources. Suitable for development and testing environments.
