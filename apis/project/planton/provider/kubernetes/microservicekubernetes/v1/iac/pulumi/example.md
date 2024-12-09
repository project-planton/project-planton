# MicroserviceKubernetes API - Example Configurations

## Example w/ Basic Configuration

### Create Using CLI

Create a YAML file using the example shown below. After the YAML is created, use the following command to apply it:

```shell
planton apply -f <yaml-path>
```

### Basic Example

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: basic-microservice
spec:
  environmentInfo:
    envId: default-env
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
          cpu: 1000m
          memory: 1Gi
```

---

## Example w/ Environment Variables

In this example, we configure environment variables for the microservice application container. These environment variables can be used to configure your application logic (e.g., database names, configuration settings, etc.).

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

---

## Example w/ Environment Secrets

This example demonstrates how to manage sensitive information (e.g., passwords) using environment secrets. These secrets are pulled from a resource like Planton Cloud's GCP Secrets Manager.

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

## Example w/ Horizontal Pod Autoscaling (HPA)

This example includes horizontal pod autoscaling, which automatically scales the number of pod replicas based on CPU utilization. The minimum number of replicas is set to 2, and autoscaling is triggered when CPU utilization exceeds 60%.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: auto-scaling-microservice
spec:
  environmentInfo:
    envId: auto-scaling-env
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
          cpu: 200m
          memory: 200Mi
        limits:
          cpu: 2000m
          memory: 2Gi
  availability:
    minReplicas: 2
    horizontalPodAutoscaling:
      isEnabled: true
      targetCpuUtilizationPercent: 60.0
```

---

## Example w/ Istio Ingress Enabled

The following example enables Istio-based ingress for the microservice, allowing the service to be accessed externally via Istio gateway configurations.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MicroserviceKubernetes
metadata:
  name: istio-microservice
spec:
  environmentInfo:
    envId: istio-enabled-env
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
          cpu: 1000m
          memory: 1Gi
  ingress:
    isEnabled: true
```
