# Terraform Examples for KubernetesOpenFga

This document provides Terraform usage examples for deploying OpenFGA (Fine-Grained Authorization) on Kubernetes.

## Example 1: Basic OpenFGA Deployment with PostgreSQL

```hcl
module "openfga_basic" {
  source = "./path/to/kubernetesopenfga/v1/iac/tf"

  metadata = {
    name = "basic-openfga"
  }

  spec = {
    target_cluster = {
      name = "my-gke-cluster"
    }
    namespace = "basic-openfga"
    create_namespace = true

    container = {
      replicas = 1

      resources = {
        requests = {
          cpu    = "50m"
          memory = "100Mi"
        }
        limits = {
          cpu    = "1000m"
          memory = "1Gi"
        }
      }
    }

    datastore = {
      engine = "postgres"
      uri    = "postgres://user:password@db-host:5432/openfga"
    }

    ingress = {
      enabled  = false
      hostname = ""
    }
  }
}
```

## Example 2: OpenFGA with Ingress Enabled

```hcl
module "openfga_with_ingress" {
  source = "./path/to/kubernetesopenfga/v1/iac/tf"

  metadata = {
    name = "ingress-openfga"
    org  = "my-organization"
    env  = "development"
  }

  spec = {
    target_cluster = {
      name = "my-gke-cluster"
    }
    namespace = "ingress-openfga"
    create_namespace = true

    container = {
      replicas = 2

      resources = {
        requests = {
          cpu    = "50m"
          memory = "256Mi"
        }
        limits = {
          cpu    = "1000m"
          memory = "1Gi"
        }
      }
    }

    datastore = {
      engine = "postgres"
      uri    = "postgres://user:password@db-host:5432/openfga"
    }

    ingress = {
      enabled  = true
      hostname = "openfga.example.com"
    }
  }
}
```

## Example 3: OpenFGA with MySQL Datastore

```hcl
module "openfga_mysql" {
  source = "./path/to/kubernetesopenfga/v1/iac/tf"

  metadata = {
    name = "mysql-openfga"
  }

  spec = {
    target_cluster = {
      name = "my-gke-cluster"
    }
    namespace = "mysql-openfga"
    create_namespace = true

    container = {
      replicas = 1

      resources = {
        requests = {
          cpu    = "50m"
          memory = "256Mi"
        }
        limits = {
          cpu    = "500m"
          memory = "512Mi"
        }
      }
    }

    datastore = {
      engine = "mysql"
      uri    = "mysql://user:password@mysql-db:3306/openfga"
    }

    ingress = {
      enabled  = false
      hostname = ""
    }
  }
}
```

## Example 4: High Availability Production Deployment

```hcl
module "openfga_production" {
  source = "./path/to/kubernetesopenfga/v1/iac/tf"

  metadata = {
    name = "production-openfga"
    id   = "prod-openfga-001"
    org  = "acme-corp"
    env  = "production"
    labels = {
      team    = "platform"
      project = "authorization"
    }
  }

  spec = {
    target_cluster = {
      name = "my-gke-cluster"
    }
    namespace = "production-openfga"
    create_namespace = true

    container = {
      replicas = 5

      resources = {
        requests = {
          cpu    = "500m"
          memory = "512Mi"
        }
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
      }
    }

    datastore = {
      engine = "postgres"
      uri    = "postgres://openfga:securepassword@ha-db-host:5432/openfga?sslmode=require"
    }

    ingress = {
      enabled  = true
      hostname = "openfga.prod.example.com"
    }

    helm_values = {
      "autoscaling.enabled" = "true"
      "autoscaling.minReplicas" = "3"
      "autoscaling.maxReplicas" = "10"
    }
  }
}

# Output the OpenFGA connection details
output "openfga_namespace" {
  value = module.openfga_production.namespace
}

output "openfga_service" {
  value = module.openfga_production.service
}

output "openfga_kube_endpoint" {
  value = module.openfga_production.kube_endpoint
}

output "openfga_external_hostname" {
  value = module.openfga_production.external_hostname
}

output "openfga_port_forward_command" {
  value = module.openfga_production.port_forward_command
}
```

## Example 5: Development Environment with Minimal Resources

```hcl
module "openfga_dev" {
  source = "./path/to/kubernetesopenfga/v1/iac/tf"

  metadata = {
    name = "dev-openfga"
    env  = "development"
  }

  spec = {
    target_cluster = {
      name = "my-gke-cluster"
    }
    namespace = "dev-openfga"
    create_namespace = true

    container = {
      replicas = 1

      resources = {
        requests = {
          cpu    = "50m"
          memory = "128Mi"
        }
        limits = {
          cpu    = "500m"
          memory = "512Mi"
        }
      }
    }

    datastore = {
      engine = "postgres"
      uri    = "postgres://postgres:devpassword@postgres.dev.svc.cluster.local:5432/openfga?sslmode=disable"
    }

    ingress = {
      enabled  = false
      hostname = ""
    }
  }
}
```

## Example 6: Staging with Custom Helm Values

```hcl
module "openfga_staging" {
  source = "./path/to/kubernetesopenfga/v1/iac/tf"

  metadata = {
    name = "staging-openfga"
    env  = "staging"
  }

  spec = {
    target_cluster = {
      name = "my-gke-cluster"
    }
    namespace = "staging-openfga"
    create_namespace = true

    container = {
      replicas = 3

      resources = {
        requests = {
          cpu    = "200m"
          memory = "512Mi"
        }
        limits = {
          cpu    = "2000m"
          memory = "2Gi"
        }
      }
    }

    datastore = {
      engine = "postgres"
      uri    = "postgres://openfga:stagingpass@postgres.staging.svc.cluster.local:5432/openfga"
    }

    ingress = {
      enabled  = true
      hostname = "openfga.staging.example.com"
    }

    helm_values = {
      "playground.enabled" = "true"
      "log.level" = "info"
      "metrics.enabled" = "true"
    }
  }
}

```

## Example 7: Using Pre-existing Namespace

```hcl
module "openfga_shared_namespace" {
  source = "./path/to/kubernetesopenfga/v1/iac/tf"

  metadata = {
    name = "openfga-shared"
  }

  spec = {
    target_cluster = {
      name = "my-gke-cluster"
    }
    namespace = "shared-services"
    create_namespace = false  # Namespace is managed externally

    container = {
      replicas = 2

      resources = {
        requests = {
          cpu    = "50m"
          memory = "256Mi"
        }
        limits = {
          cpu    = "1000m"
          memory = "1Gi"
        }
      }
    }

    datastore = {
      engine = "postgres"
      uri    = "postgres://user:password@db-host:5432/openfga"
    }

    ingress = {
      enabled  = false
      hostname = ""
    }
  }
}
```

## Output Values

The module provides the following outputs:

- `namespace` - The namespace in which OpenFGA is deployed
- `service` - The base name of the OpenFGA service
- `kube_endpoint` - Internal DNS name of the OpenFGA service within the cluster
- `port_forward_command` - Command to port-forward traffic to the OpenFGA service
- `external_hostname` - The external hostname for OpenFGA if ingress is enabled
- `internal_hostname` - The internal hostname for OpenFGA if ingress is enabled

## Connecting to OpenFGA

After deploying OpenFGA, you can connect to it using the following methods:

### From within the Kubernetes cluster:

```bash
# Use the internal endpoint
http://<kube_endpoint>:8080
```

### Using kubectl port-forward:

```bash
# Use the port-forward command from outputs
kubectl port-forward -n <namespace> service/<service-name> 8080:8080

# Then connect to localhost
http://localhost:8080
```

### Using external hostname (if ingress is enabled):

```bash
# Connect using the external hostname
https://<external_hostname>
```

## OpenFGA API Endpoints

Once deployed, the following OpenFGA endpoints are available:

- **Health Check**: `GET /healthz`
- **gRPC API**: Port 8081 (for authorization checks)
- **HTTP API**: Port 8080 (for model management)
- **Playground** (if enabled): `/playground`

## Datastore Configuration

### PostgreSQL Connection String Format

```
postgres://username:password@hostname:port/database?options
```

**Common options:**
- `sslmode=require` - Enforce SSL connection (recommended for production)
- `sslmode=disable` - Disable SSL (development only)

### MySQL Connection String Format

```
mysql://username:password@hostname:port/database
```

## Best Practices

1. **Production Deployments:**
   - Use at least 3 replicas for high availability
   - Enable ingress with TLS
   - Use SSL for database connections
   - Set appropriate resource limits based on load testing
   - Enable autoscaling via Helm values

2. **Security:**
   - Use secrets management for database credentials (do not hardcode passwords)
   - Enable SSL/TLS for both ingress and database connections
   - Use network policies to restrict traffic
   - Consider using a secrets operator for sensitive data

3. **Performance:**
   - Monitor resource usage and adjust CPU/memory accordingly
   - Use connection pooling for database connections
   - Scale replicas based on authorization request volume

4. **Development:**
   - Single replica is sufficient
   - Ingress can be disabled (use port-forwarding)
   - Minimal resources are acceptable

## Notes

- OpenFGA requires a persistent datastore (PostgreSQL or MySQL)
- The datastore must be created before deploying OpenFGA
- For production, ensure the database is highly available
- OpenFGA is stateless; all state is stored in the configured datastore
- The Helm chart configures OpenFGA with security best practices by default

