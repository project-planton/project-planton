# MicroserviceKubernetes API - Examples

Here are a few examples of how to create and configure a **MicroserviceKubernetes** API resource using the Planton Cloud CLI. The examples cover a basic setup, configuring environment variables, and handling secrets from GCP Secrets Manager.

## Create using CLI

First, create a YAML file using the examples provided below. After the YAML file is created, you can apply the configuration using the following command:

```shell
planton apply -f <yaml-path>
```

## Basic Example

This basic example demonstrates how to create a simple **MicroserviceKubernetes** resource for deploying an application using the `nginx` container image. The container listens on port 8080, with a service exposing it on port 80.

```yaml
apiVersion: gcp.project-planton.org/v1
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

In this example, an environment variable `DATABASE_NAME` is added to the container configuration. The container is still based on `nginx`, and it listens on port 8080 with the service on port 80.

```yaml
apiVersion: gcp.project-planton.org/v1
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

In this example, secrets are injected into the container using GCP Secrets Manager. The `DATABASE_PASSWORD` is pulled from the GCP Secrets Manager using the secret `gcpsm-my-org-prod-gcp-secrets.database-password`.

```yaml
apiVersion: gcp.project-planton.org/v1
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
          # Reference to the GCP Secrets Manager for the database password
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

These examples highlight how to configure **MicroserviceKubernetes** API resources for basic deployments, environment variables, and secret management via GCP.
