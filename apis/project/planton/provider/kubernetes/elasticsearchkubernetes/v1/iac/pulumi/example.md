
# Example 1: Basic Elasticsearch and Kibana Deployment

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ElasticsearchKubernetes
metadata:
  name: logging-cluster
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  elasticsearch:
    resources:
      requests:
        cpu: 500m
        memory: 1Gi
      limits:
        cpu: 1000m
        memory: 2Gi
  kibana:
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 500m
        memory: 1Gi
```

---

# Example 2: Elasticsearch with Ingress Enabled

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ElasticsearchKubernetes
metadata:
  name: search-service
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  elasticsearch:
    resources:
      requests:
        cpu: 1
        memory: 2Gi
      limits:
        cpu: 2
        memory: 4Gi
  kibana:
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 500m
        memory: 1Gi
  ingress:
    enabled: true
    hostname: search.example.com
```

---

# Example 3: Elasticsearch Deployment with Environment Variables

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ElasticsearchKubernetes
metadata:
  name: logging-app
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
  elasticsearch:
    resources:
      requests:
        cpu: 500m
        memory: 1Gi
      limits:
        cpu: 1000m
        memory: 2Gi
    env:
      variables:
        ELASTIC_PASSWORD: secret-password
        NODE_NAME: "node-1"
  kibana:
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 500m
        memory: 1Gi
```

---

# Example 4: Minimal Elasticsearch Deployment (Empty Spec)

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: ElasticsearchKubernetes
metadata:
  name: minimal-elasticsearch
spec:
  kubernetesClusterCredentialId: my-k8s-credentials
```
