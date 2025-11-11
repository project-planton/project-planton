# Create using CLI

Create a yaml using the example shown below. After the yaml is created, use the below command to apply.

```shell
planton apply -f <yaml-path>
```

# Basic Example

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: CronJobKubernetes
metadata:
  name: todo-list-api
spec:
  image:
    repo: nginx
    tag: latest
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
kind: CronJobKubernetes
metadata:
  name: todo-list-api
spec:
  env:
    variables:
      DATABASE_NAME: todo
  image:
    repo: nginx
    tag: latest
  resources:
    requests:
      cpu: 100m
      memory: 100Mi
    limits:
      cpu: 2000m
      memory: 2Gi
```

# Example w/ Environment Secrets  

The below example assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/project-planton/apis/docs/main:cloud.planton.apis.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.
```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: CronJobKubernetes
metadata:
  name: todo-list-api
spec:
  env:
    secrets:
      # value before dot 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
      # value after dot 'database-password' is one of the secrets list in 'gcpsm-my-org-prod-gcp-secrets' is the id of the gcp-secret-manager resource on planton-cloud
      DATABASE_PASSWORD: ${gcpsm-my-org-prod-gcp-secrets.database-password}
    variables:
      DATABASE_NAME: todo
  image:
    repo: nginx
  resources:
    requests:
      cpu: 100m
      memory: 100Mi
    limits:
      cpu: 2000m
      memory: 2Gi
```

# Advanced Example w/ grpc api call

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: CronJobKubernetes
metadata:
  name: grpc-invoker
spec:
  # The container image to use. In this case, a public image with grpcurl.
  image:
    repo: fullstorydev/grpcurl
    tag: latest

  # The schedule for the cron job. Runs at midnight (00:00) every day.
  schedule: "0 0 * * *"

  # How concurrency is handled. "Forbid" is the default if omitted.
  concurrencyPolicy: "Forbid"

  # Environment variables (optional). Adjust or remove as needed.
  env:
    variables:
      FOO: "BAR"
    secrets: {}

  # The commands/arguments for grpcurl. Calls "my.package.Service/Method" on "my-grpc-service:50051".
  command:
    - "grpcurl"
  args:
    - "-plaintext"
    - "my-grpc-service:50051"
    - "my.package.Service/Method"

  # Basic resource limits and requests. Adjust these to match your usage.
  resources:
    requests:
      cpu: "50m"
      memory: "100Mi"
    limits:
      cpu: "200m"
      memory: "256Mi"

  # Typically "Never" is recommended for CronJobs to avoid restarts. 
  restartPolicy: "Never"

  # Keep the last 3 successful runs and the last 1 failure in history.
  successfulJobsHistoryLimit: 3
  failedJobsHistoryLimit: 1

  # Number of retries if the job fails before considering it permanently failed.
  backoffLimit: 6

  # If you need to temporarily pause further runs, set this to true.
  # suspend: false

  # If you miss a schedule, how long (in seconds) to still try to start. 0 = no deadline enforced.
  # startingDeadlineSeconds: 0
```
