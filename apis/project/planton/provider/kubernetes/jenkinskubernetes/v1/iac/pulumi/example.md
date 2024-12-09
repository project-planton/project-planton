# Example 1: Basic Jenkins Kubernetes Setup

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: JenkinsKubernetes
metadata:
  name: jenkins-instance-basic
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  container:
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
  helm_values:
    persistence:
      enabled: true
      size: 10Gi
  ingress:
    enabled: false
```

# Example 2: Jenkins Kubernetes with Custom Helm Values

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: JenkinsKubernetes
metadata:
  name: jenkins-instance-custom
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  container:
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 2
        memory: 2Gi
  helm_values:
    controller:
      adminUser: custom-admin
      adminPassword: custom-password
    agent:
      enabled: false
    persistence:
      size: 50Gi
  ingress:
    enabled: true
    ingressClassName: "nginx"
    hosts:
      - host: jenkins.mydomain.com
        paths:
          - /
```

# Example 3: Jenkins Kubernetes with Ingress Disabled and Port Forwarding

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: JenkinsKubernetes
metadata:
  name: jenkins-no-ingress
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  container:
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
  helm_values:
    controller:
      adminUser: jenkins-admin
      adminPassword: secure-password
  ingress:
    enabled: false
```

# Example 4: Jenkins Kubernetes with Large Resource Allocation

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: JenkinsKubernetes
metadata:
  name: jenkins-high-resources
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  container:
    resources:
      requests:
        cpu: 500m
        memory: 1Gi
      limits:
        cpu: 4
        memory: 8Gi
  helm_values:
    persistence:
      enabled: true
      size: 100Gi
    controller:
      adminUser: jenkins-superadmin
      adminPassword: supersecurepassword
  ingress:
    enabled: true
    ingressClassName: "nginx"
    hosts:
      - host: jenkins.large-resources.com
        paths:
          - /
```

# Example 5: Jenkins Kubernetes with Minimal Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: JenkinsKubernetes
metadata:
  name: jenkins-minimal
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  container:
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
  ingress:
    enabled: false
```
