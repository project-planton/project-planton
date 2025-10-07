# ClickHouse Kubernetes Terraform Module - Examples

## Example 1: Minimal Configuration

```hcl
module "clickhouse_basic" {
  source = "../../tf"

  metadata = {
    name = "basic-clickhouse"
  }

  spec = {
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "8Gi"
      resources = {
        requests = {
          cpu    = "100m"
          memory = "256Mi"
        }
        limits = {
          cpu    = "1000m"
          memory = "2Gi"
        }
      }
    }
  }
}
```

## Example 2: Production with Ingress

```hcl
module "clickhouse_production" {
  source = "../../tf"

  metadata = {
    name = "production-clickhouse"
    org  = "acme-corp"
    env  = "production"
  }

  spec = {
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "100Gi"
      resources = {
        requests = {
          cpu    = "500m"
          memory = "2Gi"
        }
        limits = {
          cpu    = "4000m"
          memory = "8Gi"
        }
      }
    }
    ingress = {
      is_enabled = true
      dns_domain = "example.com"
    }
  }
}

# Outputs
output "clickhouse_endpoint" {
  value = module.clickhouse_production.kube_endpoint
}

output "clickhouse_external_url" {
  value = module.clickhouse_production.external_hostname
}
```

## Example 3: Clustered Deployment

```hcl
module "clickhouse_cluster" {
  source = "../../tf"

  metadata = {
    name = "clustered-clickhouse"
    org  = "data-team"
    env  = "production"
  }

  spec = {
    container = {
      replicas               = 3
      is_persistence_enabled = true
      disk_size              = "50Gi"
      resources = {
        requests = {
          cpu    = "500m"
          memory = "2Gi"
        }
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
      }
    }
    cluster = {
      is_enabled    = true
      shard_count   = 3
      replica_count = 2
    }
    ingress = {
      is_enabled = true
      dns_domain = "analytics.example.com"
    }
  }
}
```

## Example 4: Custom Helm Values

```hcl
module "clickhouse_custom" {
  source = "../../tf"

  metadata = {
    name = "custom-clickhouse"
  }

  spec = {
    container = {
      replicas               = 2
      is_persistence_enabled = true
      disk_size              = "30Gi"
      resources = {
        requests = {
          cpu    = "200m"
          memory = "512Mi"
        }
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
      }
    }
    helm_values = {
      "auth.username" = "clickhouse_user"
      "defaultConfigurationOverrides" = <<-EOT
        <clickhouse>
          <max_connections>100</max_connections>
          <max_concurrent_queries>50</max_concurrent_queries>
        </clickhouse>
      EOT
    }
  }
}
```

## Example 5: Development Environment

```hcl
module "clickhouse_dev" {
  source = "../../tf"

  metadata = {
    name = "dev-clickhouse"
    env  = "development"
  }

  spec = {
    container = {
      replicas               = 1
      is_persistence_enabled = false
      disk_size              = "1Gi"
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
  }
}
```

## Example 6: Using Official ClickHouse Images

```hcl
module "clickhouse_official" {
  source = "../../tf"

  metadata = {
    name = "official-clickhouse"
  }

  spec = {
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "20Gi"
      resources = {
        requests = {
          cpu    = "200m"
          memory = "512Mi"
        }
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
      }
    }
    helm_values = {
      "global.imageRegistry" = ""
      "image.registry"       = "docker.io"
      "image.repository"     = "clickhouse/clickhouse-server"
      "image.tag"            = "24.8"
    }
  }
}
```
