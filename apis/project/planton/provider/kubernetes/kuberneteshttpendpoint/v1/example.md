# Example 1: Basic Kubernetes HTTP Endpoint

```yaml
apiVersion: kubernetes.project.planton/v1
kind: KubernetesHttpEndpoint
metadata:
  name: basic-http-endpoint
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  is_tls_enabled: false
  routing_rules:
    - url_path_prefix: /api
      backend_service:
        name: my-backend-service
        namespace: default
        port: 8080
```

# Example 2: HTTPS-Enabled Kubernetes HTTP Endpoint with TLS

```yaml
apiVersion: kubernetes.project.planton/v1
kind: KubernetesHttpEndpoint
metadata:
  name: secure-http-endpoint
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  is_tls_enabled: true
  cert_cluster_issuer_name: my-cluster-issuer
  routing_rules:
    - url_path_prefix: /secure-api
      backend_service:
        name: secure-service
        namespace: default
        port: 443
```

# Example 3: Kubernetes HTTP Endpoint with Multiple Routes

```yaml
apiVersion: kubernetes.project.planton/v1
kind: KubernetesHttpEndpoint
metadata:
  name: multi-route-http-endpoint
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  is_tls_enabled: false
  routing_rules:
    - url_path_prefix: /public
      backend_service:
        name: public-service
        namespace: default
        port: 80
    - url_path_prefix: /private
      backend_service:
        name: private-service
        namespace: default
        port: 8080
```

# Example 4: gRPC-Web Compatible Kubernetes HTTP Endpoint

```yaml
apiVersion: kubernetes.project.planton/v1
kind: KubernetesHttpEndpoint
metadata:
  name: grpc-web-endpoint
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  is_tls_enabled: true
  cert_cluster_issuer_name: my-cluster-issuer
  is_grpc_web_compatible: true
  routing_rules:
    - url_path_prefix: /grpc
      backend_service:
        name: grpc-service
        namespace: default
        port: 50051
```

# Example 5: Minimal Kubernetes HTTP Endpoint Configuration

```yaml
apiVersion: kubernetes.project.planton/v1
kind: KubernetesHttpEndpoint
metadata:
  name: minimal-http-endpoint
spec:
  kubernetes_cluster_credential_id: my-cluster-creds
  is_tls_enabled: false
```
