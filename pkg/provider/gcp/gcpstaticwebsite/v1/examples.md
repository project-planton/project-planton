# Create using CLI

Create a YAML file using the examples shown below. After the YAML is created, use the command below to apply the configuration.

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: gcp.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
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

# Example with Environment Variables

```yaml
apiVersion: gcp.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
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

# Example with Environment Secrets

The example below assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
  version: main
  container:
    app:
      env:
        secrets:
          # value before dot 'gcpsm-my-org-prod-gcp-secrets' is the ID of the gcp-secret-manager resource on Planton Cloud
          # value after dot 'database-password' is one of the secrets in 'gcpsm-my-org-prod-gcp-secrets'
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
