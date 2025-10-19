# Example 1: Basic GitLab Deployment

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: GitlabKubernetes
metadata:
  name: gitlab-instance
spec:
  kubernetesProviderConfigId: my-k8s-credentials
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

# Example 2: GitLab with Ingress and Custom Hostname

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: GitlabKubernetes
metadata:
  name: gitlab-production
spec:
  kubernetesProviderConfigId: my-k8s-credentials
  container:
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 2
        memory: 2Gi
  ingress:
    enabled: true
    hostname: gitlab.example.com
```

---

# Example 3: GitLab Deployment with Custom Resources

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: GitlabKubernetes
metadata:
  name: gitlab-custom
spec:
  kubernetesProviderConfigId: my-k8s-credentials
  container:
    resources:
      requests:
        cpu: 500m
        memory: 1Gi
      limits:
        cpu: 3
        memory: 4Gi
```

---

# Example 4: GitLab with Environment Variables

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: GitlabKubernetes
metadata:
  name: gitlab-env-vars
spec:
  kubernetesProviderConfigId: my-k8s-credentials
  container:
    env:
      variables:
        GITLAB_OMNIBUS_CONFIG: |
          external_url 'http://gitlab.example.com'
          gitlab_rails['gitlab_shell_ssh_port'] = 2222
    resources:
      requests:
        cpu: 250m
        memory: 512Mi
      limits:
        cpu: 1
        memory: 1Gi
```

---

# Example 5: Minimal GitLab Deployment (Empty Spec)

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: GitlabKubernetes
metadata:
  name: minimal-gitlab
spec:
  kubernetesProviderConfigId: my-k8s-credentials
```
