# MicroserviceKubernetes - Example Configurations

This document provides a series of examples demonstrating various configurations of the **MicroserviceKubernetes** API
resource. Each example shows a typical use case, with corresponding YAML that can be applied via
`planton apply -f <filename>`.

---

## 1. Minimal Configuration

A simple example deploying a containerized application with default settings for CPU/memory and no ingress exposure.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: minimal-example
spec:
  version: main
  container:
    app:
      image:
        repo: nginx
        tag: latest
      ports:
        - name: http
          containerPort: 80
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
          isIngressPort: false
      resources:
        requests:
          cpu: "100m"
          memory: "128Mi"
        limits:
          cpu: "500m"
          memory: "512Mi"
```

**Key points**:

- Minimal `ports` configuration exposes port 80 inside the cluster.
- `isIngressPort: false` means no external ingress is configured.

---

## 2. Environment Variables

Demonstrates how to pass key-value environment variables to the container. Great for passing non-sensitive config data
like feature flags, hostnames, or numeric parameters.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: env-example
spec:
  version: main
  container:
    app:
      image:
        repo: org/my-app
        tag: "1.0.0"
      env:
        variables:
          LOG_LEVEL: debug
          FEATURE_X_ENABLED: "true"
      resources:
        requests:
          cpu: "200m"
          memory: "128Mi"
        limits:
          cpu: "800m"
          memory: "512Mi"
      ports:
        - name: http
          containerPort: 8080
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
          isIngressPort: false
```

**Key points**:

- `env.variables` sets custom environment variables accessible inside the container.
- Resource requests/limits ensure pods request and cap CPU/memory usage appropriately.

---

## 3. Using Secrets for Sensitive Data

Secrets are injected via a Kubernetes Secret named `"main"`. The reference syntax (`$ref::...`) is just an example
placeholder – your actual solution might differ, e.g., referencing secret keys from a Cloud Secret Manager.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: db-credentials-example
spec:
  version: main
  container:
    app:
      image:
        repo: org/database-connector
        tag: stable
      env:
        variables:
          DB_HOST: "db.prod.svc.cluster.local"
        secrets:
          DB_PASSWORD: $ref::secrets-group::postgres-prod::password
      resources:
        requests:
          cpu: "100m"
          memory: "200Mi"
        limits:
          cpu: "500m"
          memory: "1Gi"
      ports:
        - name: http
          containerPort: 9090
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
          isIngressPort: true
```

**Key points**:

- `DB_PASSWORD` is stored in Kubernetes as a secret.
- This keeps sensitive data out of version control and your container images.
- `isIngressPort: true` on port 9090, potentially enabling external access if `ingress.is_enabled` is set.

---

## 4. Sidecar Containers

Example with a sidecar that collects logs or metrics, using minimal resources. This can be extended to use specialized
logging, caching, or proxy sidecars.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: sidecar-example
spec:
  version: "v2"
  container:
    app:
      image:
        repo: org/main-service
        tag: "2.3.4"
      resources:
        requests:
          cpu: "200m"
          memory: "256Mi"
        limits:
          cpu: "1"
          memory: "1Gi"
      ports:
        - name: main-port
          containerPort: 8080
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
          isIngressPort: false
    sidecars:
      - name: logger
        image: org/log-agent:latest
        ports:
          - name: agent-port
            container_port: 4000
            protocol: "TCP"
        resources:
          limits:
            cpu: "100m"
            memory: "128Mi"
          requests:
            cpu: "50m"
            memory: "64Mi"
        env:
          - name: LOG_LEVEL
            value: "info"
```

**Key points**:

- The main container and a logging sidecar run together in the same pod.
- Each container can have its own resource profile and environment variables.

---

## 5. Enabling Ingress with Istio

By setting `ingress.is_enabled: true` and providing `ingress.dns_domain`, the module generates an Istio Gateway,
VirtualService (HTTPRoute), and TLS certificate resources (if configured). This allows external traffic to reach your
microservice.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: ingress-example
  labels:
    customLabel: "example"
spec:
  version: main
  container:
    app:
      image:
        repo: org/web-api
        tag: v1.1
      ports:
        - name: http
          containerPort: 8080
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
          isIngressPort: true
      resources:
        requests:
          cpu: "100m"
          memory: "128Mi"
        limits:
          cpu: "1"
          memory: "1Gi"
  ingress:
    is_enabled: true
    dns_domain: "example.org"
```

**Key points**:

- The microservice becomes accessible at `<name>.<dns_domain>` (e.g., `ingress-example.example.org`).
- The module configures Istio resources automatically if your cluster is set up to support it.

---

## 6. Scaling with Horizontal Pod Autoscaling (HPA)

Define a minimum number of replicas and enable optional autoscaling to handle increased load.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: hpa-example
spec:
  version: "3.0"
  container:
    app:
      image:
        repo: org/hpa-service
        tag: stable
      ports:
        - name: http
          containerPort: 3000
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
          isIngressPort: false
      resources:
        requests:
          cpu: "250m"
          memory: "256Mi"
        limits:
          cpu: "2"
          memory: "2Gi"
  availability:
    minReplicas: 2
    horizontalPodAutoscaling:
      is_enabled: true
      target_cpu_utilization_percent: 70.0
      target_memory_utilization: "1Gi"
```

**Key points**:

- The deployment will start with 2 replicas.
- When CPU usage rises above ~70%, autoscaling increments the pod count until usage stabilizes or the cluster limit is
  reached.

---

## Conclusion

These examples illustrate the breadth of **MicroserviceKubernetes** features, from basic single-container deployments to
advanced sidecars, secrets management, and ingress configuration. By consolidating Kubernetes manifests behind a concise
API resource definition, you can maintain consistency, reduce error-prone manual config, and accelerate delivery cycles.

> **Getting Started**
> 1. Create a YAML file for your microservice (e.g., `service.yaml`).
> 2. Run:
     >    ```shell
     > planton apply -f service.yaml
     >    ```
> 3. Verify the logs and resources in Kubernetes to ensure your deployment is functioning as expected.

For additional details, see the [MicroserviceKubernetes API documentation](#) (placeholder link), or reach out to our
support team.
