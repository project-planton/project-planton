# Kubernetes CronJob - Terraform Examples

This document provides practical examples for deploying Kubernetes CronJobs using the **KubernetesCronJob** Terraform module. Each example demonstrates different configuration aspects including scheduling, environment variables, secrets, concurrency policies, and retry strategies.

> **Note:** These examples show how to use the Terraform module directly. The module expects variables to match the protobuf specification defined in the KubernetesCronJob API.

---

## 1. Minimal Configuration

A basic example that runs a simple container on a daily schedule at midnight with default concurrency and retry settings.

```hcl
module "daily_backup_cronjob" {
  source = "./path/to/kubernetes-cronjob-module"

  metadata = {
    name = "daily-backup"
    id   = "daily-backup-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    
    namespace = "my-namespace"
    
    schedule = "0 0 * * *"
    
    image = {
      repo = "busybox"
      tag  = "latest"
    }
    
    resources = {
      limits = {
        cpu    = "500m"
        memory = "256Mi"
      }
      requests = {
        cpu    = "100m"
        memory = "128Mi"
      }
    }
  }
}
```

**Key Points:**
- **schedule**: Cron format - runs daily at midnight (00:00 UTC)
- **image**: Uses minimal busybox image for demonstration
- **resources**: Basic CPU/memory requests and limits for resource management

---

## 2. CronJob with Environment Variables

This example shows how to pass non-sensitive configuration values as environment variables.

```hcl
module "weekly_report_cronjob" {
  source = "./path/to/kubernetes-cronjob-module"

  metadata = {
    name = "weekly-report"
    id   = "weekly-report-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    
    namespace = "my-namespace"
    
    schedule = "0 9 * * 1"  # Every Monday at 9:00 AM
    
    image = {
      repo = "my-org/reports-generator"
      tag  = "v1.0.0"
    }
    
    env = {
      variables = {
        REPORT_TYPE    = "weekly"
        S3_BUCKET_NAME = "my-report-bucket"
        OUTPUT_FORMAT  = "pdf"
      }
    }
    
    resources = {
      limits = {
        cpu    = "1"
        memory = "1Gi"
      }
      requests = {
        cpu    = "100m"
        memory = "256Mi"
      }
    }
  }
}
```

**Key Points:**
- **schedule**: `"0 9 * * 1"` runs every Monday at 9:00 AM
- **env.variables**: Map of environment variables for non-sensitive configuration

---

## 3. Using Secret Environment Variables

Sensitive data should be stored in secrets. This example shows how to reference secrets in environment variables.

```hcl
module "db_maintenance_cronjob" {
  source = "./path/to/kubernetes-cronjob-module"

  metadata = {
    name = "db-maintenance"
    id   = "db-maintenance-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    
    namespace = "my-namespace"
    
    schedule = "0 3 * * *"  # Daily at 3:00 AM
    
    image = {
      repo = "my-org/maintenance-tool"
      tag  = "stable"
    }
    
    env = {
      # Sensitive data from secrets
      secrets = {
        DB_PASSWORD = "password-from-secret-manager"
        API_KEY     = "api-key-from-secret-manager"
      }
      
      # Non-sensitive configuration
      variables = {
        DB_HOST = "db-server.prod.svc.cluster.local"
        DB_NAME = "myappdb"
        DB_PORT = "5432"
      }
    }
    
    resources = {
      limits = {
        cpu    = "1"
        memory = "512Mi"
      }
      requests = {
        cpu    = "100m"
        memory = "128Mi"
      }
    }
  }
}
```

**Key Points:**
- **env.secrets**: Sensitive values that will be stored in Kubernetes secrets
- **env.variables**: Non-sensitive configuration values
- The module automatically creates a secret named "main" with the secret values

---

## 4. Concurrency Control & Retry Configuration

Control how CronJobs behave under load and when failures occur.

```hcl
module "heavy_lift_cronjob" {
  source = "./path/to/kubernetes-cronjob-module"

  metadata = {
    name = "heavy-lift-job"
    id   = "heavy-lift-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    
    namespace = "my-namespace"
    
    schedule = "*/30 * * * *"  # Every 30 minutes
    
    # Prevent concurrent executions
    concurrency_policy = "Forbid"
    
    # Retry logic
    backoff_limit = 3
    
    # Job history retention
    successful_jobs_history_limit = 2
    failed_jobs_history_limit      = 2
    
    image = {
      repo = "my-org/heavy-lift"
      tag  = "v2.1"
    }
    
    resources = {
      limits = {
        cpu    = "2"
        memory = "2Gi"
      }
      requests = {
        cpu    = "200m"
        memory = "256Mi"
      }
    }
  }
}
```

**Key Points:**
- **concurrency_policy**: `"Forbid"` prevents new runs if previous is still running
  - Other options: `"Allow"` (default), `"Replace"` (kills old run, starts new)
- **backoff_limit**: Number of retries before marking job as failed (default: 6)
- **successful_jobs_history_limit**: Keep last N successful job pods (default: 3)
- **failed_jobs_history_limit**: Keep last N failed job pods (default: 1)

---

## 5. Advanced Configuration with Custom Commands

For complex scheduling and execution requirements, including custom commands and startup deadlines.

```hcl
module "end_of_month_report" {
  source = "./path/to/kubernetes-cronjob-module"

  metadata = {
    name = "end-of-month-report"
    id   = "eom-report-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    
    namespace = "my-namespace"
    
    # Run at 23:30 on the last day of each month
    schedule = "30 23 28-31 * *"
    
    # If schedule is missed for over 2 hours, don't run retroactively
    starting_deadline_seconds = 7200
    
    # Temporarily suspend scheduling (set to false to resume)
    suspend = false
    
    # Replace running job if new schedule arrives
    concurrency_policy = "Replace"
    
    backoff_limit = 1
    
    successful_jobs_history_limit = 5
    failed_jobs_history_limit      = 5
    
    image = {
      repo = "my-org/report-service"
      tag  = "monthly-latest"
    }
    
    env = {
      variables = {
        REPORT_TZ               = "UTC"
        SEND_EMAIL_NOTIFICATIONS = "true"
        REPORT_FORMAT           = "excel"
      }
    }
    
    resources = {
      limits = {
        cpu    = "2"
        memory = "2Gi"
      }
      requests = {
        cpu    = "500m"
        memory = "512Mi"
      }
    }
  }
}
```

**Key Points:**
- **schedule**: `"30 23 28-31 * *"` runs at 23:30 on days 28-31 (catches last day of month)
- **starting_deadline_seconds**: Skip job if it can't start within this time window
- **suspend**: Set to `true` to pause scheduling; `false` to resume
- **concurrency_policy**: `"Replace"` cancels running job to start new one

---

## 6. gRPC Service Invocation Example

Example CronJob that periodically invokes a gRPC service using grpcurl.

```hcl
module "grpc_invoker_cronjob" {
  source = "./path/to/kubernetes-cronjob-module"

  metadata = {
    name = "grpc-invoker"
    id   = "grpc-invoker-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    
    namespace = "my-namespace"
    
    schedule = "0 0 * * *"  # Daily at midnight
    
    concurrency_policy = "Forbid"
    
    restart_policy = "Never"
    
    image = {
      repo = "fullstorydev/grpcurl"
      tag  = "latest"
    }
    
    # Custom command and arguments
    command = ["grpcurl"]
    args = [
      "-plaintext",
      "my-grpc-service:50051",
      "my.package.Service/Method"
    ]
    
    env = {
      variables = {
        GRPC_TIMEOUT = "30s"
      }
    }
    
    resources = {
      limits = {
        cpu    = "200m"
        memory = "256Mi"
      }
      requests = {
        cpu    = "50m"
        memory = "100Mi"
      }
    }
    
    backoff_limit = 6
    
    successful_jobs_history_limit = 3
    failed_jobs_history_limit      = 1
  }
}
```

**Key Points:**
- **command**: Override container's default ENTRYPOINT
- **args**: Override container's default CMD
- **restart_policy**: Typically `"Never"` for CronJobs (default)
  - Other option: `"OnFailure"` to retry within same pod

---

## 7. Docker Registry Authentication

When using private container registries, provide Docker credentials.

```hcl
module "private_registry_cronjob" {
  source = "./path/to/kubernetes-cronjob-module"

  metadata = {
    name = "private-app"
    id   = "private-app-prod"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    
    namespace = "my-namespace"
    
    schedule = "0 */6 * * *"  # Every 6 hours
    
    image = {
      repo = "gcr.io/my-project/my-private-image"
      tag  = "v2.0.0"
    }
    
    resources = {
      limits = {
        cpu    = "1"
        memory = "512Mi"
      }
      requests = {
        cpu    = "100m"
        memory = "128Mi"
      }
    }
  }

  # Docker config for private registry authentication
  docker_config_json = jsonencode({
    auths = {
      "gcr.io" = {
        username = "_json_key"
        password = var.gcp_service_account_key
        email    = "service-account@project.iam.gserviceaccount.com"
        auth     = base64encode("_json_key:${var.gcp_service_account_key}")
      }
    }
  })
}

# Variable for the service account key
variable "gcp_service_account_key" {
  description = "GCP service account key for accessing private container registry"
  type        = string
  sensitive   = true
}
```

**Key Points:**
- **docker_config_json**: Standard Docker config format for registry authentication
- The module creates a Kubernetes secret of type `kubernetes.io/dockerconfigjson`
- Automatically attached as imagePullSecret to the CronJob

---

## Common Cron Schedule Patterns

| Schedule         | Description                                    |
|------------------|------------------------------------------------|
| `"* * * * *"`    | Every minute                                   |
| `"*/5 * * * *"`  | Every 5 minutes                                |
| `"0 * * * *"`    | Every hour (on the hour)                       |
| `"0 0 * * *"`    | Daily at midnight                              |
| `"0 9 * * 1-5"`  | Weekdays at 9:00 AM                            |
| `"0 0 * * 0"`    | Every Sunday at midnight                       |
| `"0 0 1 * *"`    | First day of every month at midnight           |
| `"30 23 * * *"`  | Every day at 11:30 PM                          |
| `"0 */6 * * *"`  | Every 6 hours                                  |
| `"0 0 */2 * *"`  | Every 2 days at midnight                       |

---

## Output Values

The module provides these outputs:

```hcl
output "namespace" {
  description = "The Kubernetes namespace where the CronJob is deployed"
  value       = module.my_cronjob.namespace
}

output "cronjob_name" {
  description = "The name of the CronJob resource"
  value       = module.my_cronjob.cronjob_name
}

output "service_account_name" {
  description = "The service account used by the CronJob"
  value       = module.my_cronjob.service_account_name
}
```

---

## Best Practices

### 1. Resource Limits
Always set resource requests and limits to prevent resource starvation:

```hcl
resources = {
  limits = {
    cpu    = "1"      # Maximum CPU
    memory = "1Gi"    # Maximum memory
  }
  requests = {
    cpu    = "100m"   # Guaranteed CPU
    memory = "256Mi"  # Guaranteed memory
  }
}
```

### 2. Concurrency Control
Choose the right concurrency policy:
- **Forbid**: Prevent overlapping runs (recommended for most cases)
- **Allow**: Multiple instances can run simultaneously
- **Replace**: Kill old run and start new one

### 3. History Limits
Control how many completed jobs are retained:

```hcl
successful_jobs_history_limit = 3  # Keep last 3 successful runs
failed_jobs_history_limit      = 1  # Keep last 1 failed run
```

This helps with debugging while preventing cluster resource exhaustion.

### 4. Startup Deadlines
Prevent "thundering herd" of missed schedules:

```hcl
starting_deadline_seconds = 300  # Skip if can't start within 5 minutes
```

### 5. Restart Policy
For CronJobs, typically use `"Never"`:

```hcl
restart_policy = "Never"  # Don't restart failed containers
```

Use `"OnFailure"` only if you want the pod to retry internally.

---

## Troubleshooting

### Check CronJob Status
```bash
kubectl get cronjobs -n <namespace>
kubectl describe cronjob <cronjob-name> -n <namespace>
```

### View Job History
```bash
kubectl get jobs -n <namespace>
kubectl get pods -n <namespace> --show-all
```

### View Logs
```bash
kubectl logs -n <namespace> <pod-name>
kubectl logs -n <namespace> -l job-name=<job-name>
```

### Test Schedule Manually
```bash
kubectl create job --from=cronjob/<cronjob-name> <test-job-name> -n <namespace>
```

---

## Summary

- **Scheduling**: Use standard cron format in the `schedule` field
- **Resources**: Always define CPU/memory requests and limits
- **Environment**: Use `env.variables` for config, `env.secrets` for sensitive data
- **Concurrency**: Set `concurrency_policy` to control overlapping runs
- **Retry**: Configure `backoff_limit` for failure handling
- **History**: Use history limits to manage completed job retention
- **Commands**: Override container behavior with `command` and `args`
- **Registry Auth**: Provide `docker_config_json` for private images

For more information, consult the [KubernetesCronJob API documentation](../../README.md) or review the [research documentation](../../docs/README.md) for production best practices.

