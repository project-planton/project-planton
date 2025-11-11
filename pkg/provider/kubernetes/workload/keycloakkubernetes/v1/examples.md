# Example 1: Basic Keycloak Kubernetes Setup

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KeycloakKubernetes
metadata:
  name: keycloak-instance-basic
spec:
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

# Example 2: Keycloak Kubernetes with Ingress Enabled

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KeycloakKubernetes
metadata:
  name: keycloak-instance-ingress
spec:
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
    ingressClassName: "nginx"
    hosts:
      - host: keycloak.mydomain.com
        paths:
          - /
```

# Example 3: Keycloak Kubernetes with Minimal Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KeycloakKubernetes
metadata:
  name: keycloak-minimal
spec:
  container:
    resources:
      requests:
        cpu: 50m
        memory: 256Mi
      limits:
        cpu: 1
        memory: 1Gi
```

# Example 4: Keycloak Kubernetes with High Resource Allocation

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KeycloakKubernetes
metadata:
  name: keycloak-high-resources
spec:
  container:
    resources:
      requests:
        cpu: 500m
        memory: 2Gi
      limits:
        cpu: 4
        memory: 8Gi
  ingress:
    enabled: true
    ingressClassName: "nginx"
    hosts:
      - host: keycloak-large.mydomain.com
        paths:
          - /
```

# Example 5: Keycloak Kubernetes with Port Forwarding for Local Access

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KeycloakKubernetes
metadata:
  name: keycloak-port-forward
spec:
  container:
    resources:
      requests:
        cpu: 100m
        memory: 512Mi
      limits:
        cpu: 2
        memory: 2Gi
  ingress:
    enabled: false
```
