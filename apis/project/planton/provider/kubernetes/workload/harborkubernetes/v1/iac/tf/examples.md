# Harbor Kubernetes Terraform - Examples

## Basic Example

```hcl
module "harbor_basic" {
  source = "../"

  harbor_kubernetes = {
    metadata = {
      name = "harbor-dev"
    }
    spec = {
      database = {
        is_external = false
      }
      cache = {
        is_external = false
      }
      storage = {
        type = "filesystem"
        filesystem = {
          disk_size = "100Gi"
        }
      }
    }
  }
}
```

## Production Example with External Dependencies

```hcl
module "harbor_production" {
  source = "../"

  harbor_kubernetes = {
    metadata = {
      name = "harbor-prod"
    }
    spec = {
      database = {
        is_external = true
        external_database = {
          host     = "postgres.example.com"
          port     = 5432
          username = "harbor"
          password = var.db_password
          use_ssl  = true
        }
      }
      cache = {
        is_external = true
        external_cache = {
          host     = "redis.example.com"
          port     = 6379
          password = var.redis_password
        }
      }
      storage = {
        type = "s3"
        s3 = {
          bucket     = "my-harbor-artifacts"
          region     = "us-west-2"
          access_key = var.aws_access_key
          secret_key = var.aws_secret_key
        }
      }
      ingress = {
        core = {
          enabled  = true
          hostname = "harbor.example.com"
        }
      }
    }
  }
}
```

## Note

These examples show the variable structure. The actual Terraform module would need to be
extended with Helm release resources to fully deploy Harbor. Refer to the Pulumi module
for a complete implementation reference.

