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
  target_cluster:
    cluster_name: "my-gke-cluster"
  namespace:
    value: cronjobs
  create_namespace: true
  schedule: "0 0 * * *"
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
  target_cluster:
    cluster_name: "my-gke-cluster"
  namespace:
    value: cronjobs
  create_namespace: true
  schedule: "0 0 * * *"
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

The below example assumes that the secrets are managed by Planton Cloud's [GCP Secrets Manager](https://buf.build/project-planton/apis/docs/main:ai.planton.code2cloud.v1.gcp.gcpsecretsmanager) deployment module.
```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: CronJobKubernetes
metadata:
  name: todo-list-api
spec:
  target_cluster:
    cluster_name: "my-gke-cluster"
  namespace:
    value: cronjobs
  create_namespace: true
  schedule: "0 0 * * *"
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
  target_cluster:
    cluster_name: "my-gke-cluster"
  namespace:
    value: cronjobs
  create_namespace: true
  
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

# Namespace Management

The `create_namespace` field controls whether the namespace is automatically created or if an existing namespace should be used.

## Creating a New Namespace

Set `create_namespace: true` to automatically create the namespace:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: CronJobKubernetes
metadata:
  name: isolated-job
spec:
  target_cluster:
    cluster_name: "my-gke-cluster"
  namespace:
    value: my-cronjobs
  create_namespace: true
  schedule: "0 2 * * *"
  image:
    repo: busybox
    tag: latest
  resources:
    requests:
      cpu: "100m"
      memory: "128Mi"
    limits:
      cpu: "500m"
      memory: "256Mi"
```

**Use this when:**
- You want dedicated namespace isolation for the CronJob
- The namespace doesn't exist yet
- You want the CronJob to manage its namespace lifecycle

## Using an Existing Namespace

Set `create_namespace: false` to deploy into a pre-existing namespace:

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: CronJobKubernetes
metadata:
  name: shared-job
spec:
  target_cluster:
    cluster_name: "my-gke-cluster"
  namespace:
    value: shared-batch-jobs
  create_namespace: false
  schedule: "0 4 * * *"
  image:
    repo: busybox
    tag: latest
  resources:
    requests:
      cpu: "100m"
      memory: "128Mi"
    limits:
      cpu: "500m"
      memory: "256Mi"
```

**Use this when:**
- Multiple CronJobs share the same namespace
- The namespace is managed separately (e.g., by a platform team)
- You're following a GitOps pattern where namespaces are pre-created
- **Important:** Ensure the namespace exists before deploying, or the deployment will fail

## Best Practices

- **Isolated workloads**: Use `create_namespace: true` for dedicated, single-purpose CronJobs
- **Shared namespaces**: Use `create_namespace: false` when multiple CronJobs need to coexist in the same namespace
- **GitOps workflows**: Set `create_namespace: false` and manage namespaces separately for better lifecycle control
- **Development environments**: Use `create_namespace: true` for quick iterations and easy cleanup
```
