Here are a few examples for the `MicroserviceKubernetes` API resource, showcasing different configurations with environment variables and secrets.

# Create using CLI

Create a YAML file using one of the examples shown below. After the YAML is created, use the following command to apply it:

```shell
planton apply -f <yaml-path>
```

## Basic Example

This example demonstrates a simple microservice configuration using a basic container running `nginx`.

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

This example includes environment variables passed to the container for configuring database information.

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

This example shows how to use secrets managed by Planton Cloud's GCP Secrets Manager module to securely pass sensitive data to the application.

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
          # The key before the dot 'gcpsm-my-org-prod-gcp-secrets' is the ID of the GCP Secret Manager resource
          # The key after the dot 'database-password' is the specific secret stored in the GCP Secret Manager
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
