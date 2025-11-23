# Example 1: Basic GitLab Deployment

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: GitlabKubernetes
metadata:
  name: gitlab-instance
spec:
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: gitlab-instance
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
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: gitlab-production
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
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: gitlab-custom
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
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: gitlab-env-vars
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

# Example 5: Minimal GitLab Deployment

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: GitlabKubernetes
metadata:
  name: minimal-gitlab
spec:
  target_cluster:
    cluster_name: my-gke-cluster
  namespace:
    value: minimal-gitlab
  container:
    resources:
      requests:
        cpu: 50m
        memory: 100Mi
      limits:
        cpu: 1
        memory: 1Gi
```
