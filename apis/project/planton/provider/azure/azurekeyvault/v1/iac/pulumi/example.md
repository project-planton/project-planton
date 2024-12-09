# MicroserviceKubernetes API Resource Examples

This document provides multiple examples for the `MicroserviceKubernetes` API resource, showcasing different configurations such as basic setups, usage with environment variables, and secrets management using Planton Cloud's GCP Secrets Manager.

## Create using CLI

To create and apply the `MicroserviceKubernetes` resource, follow the steps below:

1. Create a YAML file using one of the examples provided.
2. Apply the resource using the following command:

```shell
planton apply -f <yaml-path>
```

## Basic Example

This basic example demonstrates how to define a `MicroserviceKubernetes` resource that deploys a simple containerized application using the `nginx` image. 

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

This example extends the basic setup by introducing environment variables. The `DATABASE_NAME` environment variable is defined within the container configuration.

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

This example shows how to integrate Planton Cloud's GCP Secrets Manager for managing sensitive data like database credentials. The `DATABASE_PASSWORD` secret is retrieved from the GCP Secrets Manager, while the `DATABASE_NAME` is provided as an environment variable.

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

These examples show the flexibility of the `MicroserviceKubernetes` API resource, allowing you to customize deployments using environment variables and secrets management, while adhering to a standard API resource structure.
