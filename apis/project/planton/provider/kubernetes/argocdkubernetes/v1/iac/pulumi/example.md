Here are a few examples for the `ArgocdKubernetes` API resource, modeled in a similar way to the `MicroserviceKubernetes` examples you provided. These examples demonstrate how to configure and deploy ArgoCD on a Kubernetes cluster using Planton Cloud’s unified API structure.

---

# Example 1: Basic ArgoCD Deployment

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ArgocdKubernetes
metadata:
  name: argocd-instance
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  container:
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
```

---

# Example 2: ArgoCD with Ingress Enabled

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ArgocdKubernetes
metadata:
  name: argocd-prod
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  container:
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 2
        memory: 2Gi
  ingress:
    enabled: true
    hostname: argocd.example.com
```

---

# Example 3: ArgoCD Deployment with Custom Resources

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ArgocdKubernetes
metadata:
  name: argocd-custom
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  container:
    resources:
      requests:
        cpu: 200m
        memory: 1Gi
      limits:
        cpu: 3
        memory: 4Gi
```

---

# Example 4: Minimal ArgoCD Deployment (Empty Spec)

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ArgocdKubernetes
metadata:
  name: minimal-argocd
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
```

