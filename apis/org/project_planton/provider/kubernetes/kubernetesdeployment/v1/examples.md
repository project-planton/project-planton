# Create using CLI

Create a yaml using the example shown below. After the yaml is created, use the below command to apply.

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: kubernetes.project-planton.org/v1
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

# Example w/ Environment Variables

```yaml
apiVersion: kubernetes.project-planton.org/v1
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

# Example w/ Environment Secrets  

The below example assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/project-planton/apis/docs/main:ai.planton.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.
```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: MicroserviceKubernetes
metadata:
  name: todo-list-api
spec:
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

# Example w/ ConfigMaps and Volume Mounts

This example shows how to create ConfigMaps and mount them as files in the container.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesDeployment
metadata:
  name: tekton-cloudevents-router
spec:
  namespace:
    value: tekton-pipelines
  version: main
  configMaps:
    router-config: |
      routes:
        - namespace_prefix: planton-dev-
          target: http://servicehub-tekton-webhooks.planton-dev.svc/webhook
        - namespace_prefix: planton-prod-
          target: http://servicehub-tekton-webhooks.planton-prod.svc/webhook
  container:
    app:
      image:
        repo: ghcr.io/plantoncloud/tekton-cloud-event-router
        tag: v0.1.0
      volumeMounts:
        - name: router-config
          mountPath: /etc/router/config.yaml
          configMap:
            name: router-config
            key: router-config
      ports:
        - name: http
          containerPort: 8080
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
      resources:
        requests:
          cpu: 50m
          memory: 64Mi
        limits:
          cpu: 200m
          memory: 256Mi
```

# Example w/ Command and Args Override

This example shows how to override the container's command and arguments.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesDeployment
metadata:
  name: custom-entrypoint-app
spec:
  namespace:
    value: my-namespace
  version: main
  container:
    app:
      image:
        repo: busybox
        tag: latest
      command:
        - /bin/sh
        - -c
      args:
        - echo "Hello, World!" && sleep 3600
      ports:
        - name: http
          containerPort: 8080
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
      resources:
        requests:
          cpu: 50m
          memory: 64Mi
        limits:
          cpu: 100m
          memory: 128Mi
```

# Example w/ Multiple Volume Types

This example demonstrates mounting various volume types including ConfigMaps, Secrets, EmptyDir, and HostPath.

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: KubernetesDeployment
metadata:
  name: multi-volume-app
spec:
  namespace:
    value: my-namespace
  version: main
  configMaps:
    app-config: |
      database:
        host: postgres.default.svc
        port: 5432
      cache:
        host: redis.default.svc
        port: 6379
  container:
    app:
      image:
        repo: my-app
        tag: v1.0.0
      volumeMounts:
        # Mount ConfigMap as a single file
        - name: config-volume
          mountPath: /etc/app/config.yaml
          configMap:
            name: app-config
            key: app-config
        # Mount a Secret
        - name: tls-certs
          mountPath: /etc/tls
          readOnly: true
          secret:
            name: my-tls-secret
        # EmptyDir for temporary storage
        - name: cache
          mountPath: /tmp/cache
          emptyDir:
            medium: Memory
            sizeLimit: 256Mi
      ports:
        - name: http
          containerPort: 8080
          networkProtocol: TCP
          appProtocol: http
          servicePort: 80
      resources:
        requests:
          cpu: 100m
          memory: 128Mi
        limits:
          cpu: 500m
          memory: 512Mi
```
