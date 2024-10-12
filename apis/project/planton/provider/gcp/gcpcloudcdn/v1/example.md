# Examples

## Basic Example

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

## Example with Environment Variables

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      env:
        variables:
          DATABASE_NAME: todo
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: rest-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

## Example with Environment Secrets

The below example assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  environmentInfo:
    envId: my-org-prod
  version: main
  container:
    app:
      env:
        secrets:
          # value before dot 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
          # value after dot 'database-password' is one of the secrets list in 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
          DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.database-password}
        variables:
          DATABASE_NAME: todo
      image:
        repo: nginx
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: rest-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 100Mi
        limits:
          cpu: 2000m
          memory: 2Gi
```

## Example with Multiple Containers

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: multi-container-app
spec:
  environmentInfo:
    envId: dev-env
  version: v1.0.0
  container:
    app:
      image:
        repo: myapp/frontend
        tag: v1.0.0
      ports:
        - appProtocol: http
          containerPort: 80
          isIngressPort: true
          servicePort: 8080
      resources:
        requests:
          cpu: 200m
          memory: 256Mi
        limits:
          cpu: 1000m
          memory: 1Gi
    sidecar:
      image:
        repo: myapp/logging
        tag: v1.0.0
      ports:
        - appProtocol: tcp
          containerPort: 5000
          isIngressPort: false
          servicePort: 5000
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 500m
          memory: 512Mi
```

## Example with Different Resource Limits

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: high-memory-service
spec:
  environmentInfo:
    envId: staging-env
  version: beta
  container:
    app:
      image:
        repo: highmemapp/backend
        tag: latest
      ports:
        - appProtocol: http
          containerPort: 9090
          isIngressPort: true
          servicePort: 9090
      resources:
        requests:
          cpu: 500m
          memory: 512Mi
        limits:
          cpu: 4000m
          memory: 8Gi
```

## Example with Annotations and Labels

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: annotated-service
  labels:
    app: annotated-app
    tier: backend
  annotations:
    description: "This service handles user authentication."
spec:
  environmentInfo:
    envId: production-env
  version: release
  container:
    app:
      image:
        repo: auth-service/image
        tag: release
      ports:
        - appProtocol: https
          containerPort: 8443
          isIngressPort: true
          servicePort: 443
      resources:
        requests:
          cpu: 250m
          memory: 256Mi
        limits:
          cpu: 1500m
          memory: 2Gi
```

## Example with Health Checks

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: healthcheck-service
spec:
  environmentInfo:
    envId: test-env
  version: test
  container:
    app:
      image:
        repo: healthapp/service
        tag: test
      ports:
        - appProtocol: http
          containerPort: 8000
          isIngressPort: true
          servicePort: 8000
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 500m
          memory: 1Gi
      livenessProbe:
        httpGet:
          path: /healthz
          port: 8000
        initialDelaySeconds: 30
        periodSeconds: 10
      readinessProbe:
        httpGet:
          path: /ready
          port: 8000
        initialDelaySeconds: 10
        periodSeconds: 5
```

## Example with Empty Spec

*Note: This module is not completely implemented. Certain features may be missing or not fully functional. Future updates will address these limitations.*

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: incomplete-service
spec: {}
```

```yaml
apiVersion: gcp.project.planton/v1
kind: MicroserviceKubernetes
metadata:
  name: another-incomplete-service
spec: {}
```

---

*Thank you for choosing Planton Cloud's Microservice Kubernetes Pulumi Module. We are dedicated to supporting your infrastructure management needs and look forward to helping you achieve seamless and efficient microservice deployments.*
