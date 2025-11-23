# Terraform Examples for KubernetesNats

This document provides Terraform usage examples for deploying NATS on Kubernetes using the official NATS Helm chart.

## Example 1: Basic NATS Cluster with Default Settings

Deploy a simple NATS cluster with default replicas, resources, and JetStream enabled.

```hcl
module "nats_basic" {
  source = "./path/to/kubernetesnats/v1/iac/tf"

  metadata = {
    name = "nats-basic"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "nats-basic"

    server_container = {
      replicas = 3
      resources = {
        limits = {
          cpu    = "1000m"
          memory = "2Gi"
        }
        requests = {
          cpu    = "100m"
          memory = "256Mi"
        }
      }
      disk_size = "10Gi"
    }

    disable_jet_stream = false
    tls_enabled        = false
    disable_nats_box   = false
  }
}

# Output the connection details
output "nats_namespace" {
  value = module.nats_basic.namespace
}

output "nats_internal_url" {
  value = module.nats_basic.internal_client_url
}
```

## Example 2: NATS Cluster with Bearer Token Authentication

Set up a secure NATS cluster using bearer token authentication.

```hcl
module "nats_secure" {
  source = "./path/to/kubernetesnats/v1/iac/tf"

  metadata = {
    name = "nats-secure"
    org  = "my-organization"
    env  = "production"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "nats-secure"

    server_container = {
      replicas = 5
      resources = {
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
        requests = {
          cpu    = "500m"
          memory = "1Gi"
        }
      }
      disk_size = "20Gi"
    }

    auth = {
      enabled = true
      scheme  = "bearer_token"
    }

    tls_enabled      = true
    disable_jet_stream = false
  }
}

# Output connection and auth details
output "nats_namespace" {
  value = module.nats_secure.namespace
}

output "nats_internal_url" {
  value = module.nats_secure.internal_client_url
}

output "auth_secret_name" {
  value = "auth-nats"
}

output "auth_secret_key" {
  value = "token"
}
```

## Example 3: NATS Cluster with Basic Authentication

Deploy NATS with basic username/password authentication.

```hcl
module "nats_basic_auth" {
  source = "./path/to/kubernetesnats/v1/iac/tf"

  metadata = {
    name = "nats-basic-auth"
  }

  spec = {
    target_cluster = {
      cluster_name = "my-gke-cluster"
    }
    namespace = "nats-basic-auth"

    server_container = {
      replicas = 3
      resources = {
        limits = {
          cpu    = "1000m"
          memory = "2Gi"
        }
        requests = {
          cpu    = "100m"
          memory = "256Mi"
        }
      }
      disk_size = "10Gi"
    }

    auth = {
      enabled = true
      scheme  = "basic_auth"
    }

    tls_enabled = false
  }
}

# Output authentication details
output "nats_namespace" {
  value = module.nats_basic_auth.namespace
}

output "username" {
  value = "nats" # Default admin username
}

output "auth_secret_name" {
  value = "auth-nats"
}

output "password_key" {
  value = "password"
}
```

## Example 4: NATS Cluster with External Access via Ingress

Deploy NATS configured with ingress to allow external clients to access messaging services.

```hcl
module "nats_external" {
  source = "./path/to/kubernetesnats/v1/iac/tf"

  metadata = {
    name = "nats-external"
  }

  spec = {
    server_container = {
      replicas = 3
      resources = {
        limits = {
          cpu    = "1000m"
          memory = "2Gi"
        }
        requests = {
          cpu    = "100m"
          memory = "256Mi"
        }
      }
      disk_size = "10Gi"
    }

    ingress = {
      enabled  = true
      hostname = "nats.example.com"
    }

    auth = {
      enabled = true
      scheme  = "basic_auth"
    }

    tls_enabled = true
  }
}

# Output connection details
output "nats_namespace" {
  value = module.nats_external.namespace
}

output "external_hostname" {
  value = module.nats_external.external_hostname
}

output "internal_client_url" {
  value = module.nats_external.internal_client_url
}

output "external_client_url" {
  value = "nats://nats.example.com:4222"
}
```

## Example 5: Lightweight NATS Cluster Without JetStream

Set up a lightweight NATS messaging cluster without JetStream persistence.

```hcl
module "nats_minimal" {
  source = "./path/to/kubernetesnats/v1/iac/tf"

  metadata = {
    name = "nats-minimal"
    env  = "development"
  }

  spec = {
    server_container = {
      replicas = 1
      resources = {
        limits = {
          cpu    = "500m"
          memory = "512Mi"
        }
        requests = {
          cpu    = "100m"
          memory = "128Mi"
        }
      }
      disk_size = "1Gi"
    }

    disable_jet_stream = true
    tls_enabled        = false
    disable_nats_box   = true
  }
}

# Output minimal connection details
output "nats_namespace" {
  value = module.nats_minimal.namespace
}

output "nats_internal_url" {
  value = module.nats_minimal.internal_client_url
}
```

## Example 6: High Availability NATS with Basic Auth and No-Auth User

Deploy a highly available NATS cluster with both authenticated and unauthenticated access.

```hcl
module "nats_ha" {
  source = "./path/to/kubernetesnats/v1/iac/tf"

  metadata = {
    name = "nats-ha-metrics"
    org  = "acme-corp"
    env  = "production"
    labels = {
      team    = "platform"
      project = "messaging"
    }
  }

  spec = {
    server_container = {
      replicas = 7
      resources = {
        limits = {
          cpu    = "4000m"
          memory = "8Gi"
        }
        requests = {
          cpu    = "1000m"
          memory = "2Gi"
        }
      }
      disk_size = "50Gi"
    }

    auth = {
      enabled = true
      scheme  = "basic_auth"
      no_auth_user = {
        enabled = true
        publish_subjects = [
          "telemetry.>",
          "metrics.>",
        ]
      }
    }

    tls_enabled      = true
    disable_jet_stream = false

    ingress = {
      enabled  = true
      hostname = "nats-ha.example.com"
    }
  }
}

# Output comprehensive connection details
output "nats_namespace" {
  value = module.nats_ha.namespace
}

output "external_hostname" {
  value = module.nats_ha.external_hostname
}

output "internal_client_url" {
  value = module.nats_ha.internal_client_url
}

output "metrics_endpoint" {
  value = "http://nats-prom.${module.nats_ha.namespace}.svc.cluster.local:7777/metrics"
}
```

## Example 7: Complete Production Setup with All Features

A production-ready configuration with all security features enabled.

```hcl
module "nats_production" {
  source = "./path/to/kubernetesnats/v1/iac/tf"

  metadata = {
    name = "nats-prod"
    id   = "nats-prod-001"
    org  = "acme-corp"
    env  = "production"
    labels = {
      tier        = "platform"
      component   = "messaging"
      managed_by  = "terraform"
      cost_center = "engineering"
    }
  }

  spec = {
    server_container = {
      replicas = 5
      resources = {
        limits = {
          cpu    = "3000m"
          memory = "6Gi"
        }
        requests = {
          cpu    = "750m"
          memory = "1.5Gi"
        }
      }
      disk_size = "100Gi"
    }

    auth = {
      enabled = true
      scheme  = "bearer_token"
    }

    tls_enabled      = true
    disable_jet_stream = false
    disable_nats_box   = false

    ingress = {
      enabled  = true
      hostname = "nats.prod.example.com"
    }
  }
}

# Comprehensive outputs
output "nats_production_namespace" {
  description = "The namespace where NATS is deployed"
  value       = module.nats_production.namespace
}

output "nats_production_internal_url" {
  description = "Internal cluster URL for NATS clients"
  value       = module.nats_production.internal_client_url
}

output "nats_production_external_hostname" {
  description = "External hostname for NATS access"
  value       = module.nats_production.external_hostname
}
```

## Retrieving Authentication Credentials

### For Bearer Token Authentication:

```bash
# Get the bearer token from the secret
kubectl get secret auth-nats -n <namespace> -o jsonpath='{.data.token}' | base64 -d
```

### For Basic Authentication:

```bash
# Get username
kubectl get secret auth-nats -n <namespace> -o jsonpath='{.data.user}' | base64 -d

# Get password
kubectl get secret auth-nats -n <namespace> -o jsonpath='{.data.password}' | base64 -d
```

### Using Terraform Data Sources:

```hcl
data "kubernetes_secret_v1" "nats_auth" {
  metadata {
    name      = "auth-nats"
    namespace = module.nats_basic_auth.namespace
  }

  depends_on = [module.nats_basic_auth]
}

# Output the credentials (sensitive)
output "nats_username" {
  value     = data.kubernetes_secret_v1.nats_auth.data["user"]
  sensitive = true
}

output "nats_password" {
  value     = data.kubernetes_secret_v1.nats_auth.data["password"]
  sensitive = true
}
```

## Connecting to NATS

### From within the Kubernetes cluster:

```bash
# Use the internal client URL
nats://nats-basic-nats.nats-basic.svc.cluster.local:4222
```

### Using kubectl port-forward:

```bash
# Port-forward to NATS service
kubectl port-forward -n <namespace> service/<name>-nats 4222:4222

# Then connect to localhost:4222
nats://localhost:4222
```

### Using external hostname (if ingress is enabled):

```bash
# Connect using the external hostname
nats://<external-hostname>:4222
```

## Notes

- The deployment uses the official NATS Helm chart (version 1.3.6) for production-grade message streaming
- JetStream provides persistent streaming capabilities with file-based storage backed by PersistentVolumeClaims
- For production deployments, use an odd number of replicas (3, 5, 7) for optimal quorum in clustered mode
- TLS encryption is automatically configured with self-signed certificates when `tls_enabled` is true
- The `nats-box` utility pod is deployed by default for debugging and testing (can be disabled)
- Metrics are exposed at `http://nats-prom.<namespace>.svc.cluster.local:7777/metrics` for Prometheus scraping
- When using basic auth with no-auth user, the no-auth user can only publish to specified subjects
- The module automatically handles secret generation for authentication credentials

