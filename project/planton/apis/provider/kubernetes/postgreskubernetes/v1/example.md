# Multiple Examples for `MicroserviceKubernetes` API-Resource

## Example with Environment Variables

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
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
          LOG_LEVEL: debug
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

---

## Example with Environment Secrets

*Note: This example assumes that secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/plantoncloud/planton-cloud-apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.*

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
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
          # Format: ${<secret-manager-id>.<secret-key>}
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

---

## Example with Multiple Containers

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: multi-container-app
spec:
  environmentInfo:
    envId: my-org-staging
  version: develop
  container:
    app:
      image:
        repo: myorg/multi-container-app
        tag: v1.2.3
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: main-api
          networkProtocol: TCP
          servicePort: 80
        - appProtocol: grpc
          containerPort: 9090
          isIngressPort: false
          name: grpc-api
          networkProtocol: TCP
          servicePort: 9090
      resources:
        requests:
          cpu: 250m
          memory: 256Mi
        limits:
          cpu: 1000m
          memory: 1Gi
    sidecar:
      image:
        repo: myorg/log-collector
        tag: stable
      ports:
        - appProtocol: tcp
          containerPort: 514
          isIngressPort: false
          name: log-collector
          networkProtocol: TCP
          servicePort: 514
      resources:
        requests:
          cpu: 50m
          memory: 64Mi
        limits:
          cpu: 200m
          memory: 128Mi
```

---

## Example with Custom Ingress Settings

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: custom-ingress-app
spec:
  environmentInfo:
    envId: my-org-development
  version: feature-branch
  container:
    app:
      env:
        variables:
          API_KEY: your-api-key
      image:
        repo: myorg/custom-ingress-app
        tag: beta
      ports:
        - appProtocol: https
          containerPort: 8443
          isIngressPort: true
          name: secure-api
          networkProtocol: TCP
          servicePort: 443
      resources:
        requests:
          cpu: 150m
          memory: 200Mi
        limits:
          cpu: 1500m
          memory: 1.5Gi
      ingress:
        isEnabled: true
        annotations:
          kubernetes.io/ingress.class: "nginx"
          cert-manager.io/cluster-issuer: "letsencrypt-prod"
        hosts:
          - host: api.dev.myorg.com
            paths:
              - path: /
                pathType: Prefix
```

---

## Example with Different Datastore Configuration

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: datastore-config-app
spec:
  environmentInfo:
    envId: my-org-testing
  version: release-1.0
  container:
    app:
      env:
        variables:
          DATABASE_NAME: testdb
      image:
        repo: myorg/datastore-config-app
        tag: stable
      ports:
        - appProtocol: http
          containerPort: 8000
          isIngressPort: true
          name: api-server
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 200m
          memory: 256Mi
        limits:
          cpu: 2500m
          memory: 2.5Gi
      datastore:
        engine: postgres
        uri: postgres://user:password@postgres-service:5432/testdb
```

---

## Example with Minimal Configuration

*Note: This module is not completely implemented.*

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: minimal-app
spec: {}
```

---

## Example with Advanced Resource Allocation

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: advanced-resources-app
spec:
  environmentInfo:
    envId: enterprise-prod
  version: v2.0.1
  container:
    app:
      env:
        variables:
          SERVICE_MODE: high-performance
          MAX_CONNECTIONS: "5000"
      image:
        repo: myorg/advanced-resources-app
        tag: v2.0.1
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: high-perf-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 500m
          memory: 512Mi
        limits:
          cpu: 4000m
          memory: 4Gi
```

---

## Example with Multiple Ingress Hosts

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: multi-ingress-host-app
spec:
  environmentInfo:
    envId: my-org-multi
  version: multi-ingress
  container:
    app:
      env:
        variables:
          FEATURE_FLAG: enabled
      image:
        repo: myorg/multi-ingress-app
        tag: release
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: api-http
          networkProtocol: TCP
          servicePort: 80
        - appProtocol: https
          containerPort: 8443
          isIngressPort: true
          name: api-https
          networkProtocol: TCP
          servicePort: 443
      resources:
        requests:
          cpu: 300m
          memory: 300Mi
        limits:
          cpu: 1500m
          memory: 1.5Gi
      ingress:
        isEnabled: true
        annotations:
          kubernetes.io/ingress.class: "nginx"
          cert-manager.io/cluster-issuer: "letsencrypt-staging"
        hosts:
          - host: api1.myorg.com
            paths:
              - path: /api1
                pathType: Prefix
          - host: api2.myorg.com
            paths:
              - path: /api2
                pathType: Prefix
```

---

## Example with Scaling Configuration

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: scalable-app
spec:
  environmentInfo:
    envId: my-org-scalable
  version: scalable-v1
  container:
    app:
      env:
        variables:
          SCALABLE_FEATURE: true
      image:
        repo: myorg/scalable-app
        tag: v1.0.0
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: scalable-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 200m
          memory: 256Mi
        limits:
          cpu: 2000m
          memory: 2Gi
      scaling:
        minReplicas: 2
        maxReplicas: 10
        targetCPUUtilizationPercentage: 75
```

---

## Example with Logging and Monitoring Integration

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: monitoring-app
spec:
  environmentInfo:
    envId: my-org-monitoring
  version: monitoring-v2
  container:
    app:
      env:
        variables:
          LOG_LEVEL: info
          MONITORING_ENABLED: true
      image:
        repo: myorg/monitoring-app
        tag: v2.1.0
      ports:
        - appProtocol: http
          containerPort: 9090
          isIngressPort: true
          name: monitoring-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 150m
          memory: 200Mi
        limits:
          cpu: 1500m
          memory: 1.5Gi
      logging:
        enabled: true
        logLevel: info
      monitoring:
        enabled: true
        endpoint: /metrics
```

---

## Example with Service Mesh Integration

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: service-mesh-app
spec:
  environmentInfo:
    envId: my-org-mesh
  version: mesh-v1
  container:
    app:
      env:
        variables:
          MESH_ENABLED: true
      image:
        repo: myorg/service-mesh-app
        tag: v1.0.0
      ports:
        - appProtocol: http
          containerPort: 8080
          isIngressPort: true
          name: mesh-api
          networkProtocol: TCP
          servicePort: 80
      resources:
        requests:
          cpu: 250m
          memory: 256Mi
        limits:
          cpu: 2500m
          memory: 2.5Gi
      serviceMesh:
        enabled: true
        sidecar:
          image: istio/proxyv2
          resources:
            requests:
              cpu: 100m
              memory: 128Mi
            limits:
              cpu: 500m
              memory: 512Mi
```

---

## Example with Feature Flags

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: feature-flag-app
spec:
  environmentInfo:
    envId: my-org-feature
  version: feature-flags
  container:
    app:
      env:
        variables:
          FEATURE_X_ENABLED: "true"
          FEATURE_Y_ENABLED: "false"
      image:
        repo: myorg/feature-flag-app
        tag: v1.1.0
      ports:
        - appProtocol: http
          containerPort: 8081
          isIngressPort: true
          name: feature-api
          networkProtocol: TCP
          servicePort: 81
      resources:
        requests:
          cpu: 120m
          memory: 150Mi
        limits:
          cpu: 1800m
          memory: 1.8Gi
      featureFlags:
        featureX:
          enabled: true
        featureY:
          enabled: false
```

---

## Example with Empty Spec

*Note: This module is not completely implemented.*

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: empty-spec-app
spec: {}
```
