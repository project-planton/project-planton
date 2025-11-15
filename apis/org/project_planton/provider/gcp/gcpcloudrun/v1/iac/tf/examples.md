# GCP Cloud Run Terraform Module - Examples

This document provides examples of using the GCP Cloud Run Terraform module with different configurations.

## Example 1: Minimal Configuration

The simplest deployment with only required fields:

```hcl
module "cloud_run_minimal" {
  source = "./path/to/module"

  metadata = {
    name = "simple-api"
    org  = "my-org"
    env  = "dev"
  }

  spec = {
    project_id = "my-gcp-project-123"
    region     = "us-central1"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/my-app"
        tag  = "1.0.0"
      }
      cpu    = 1
      memory = 512
      replicas = {
        min = 0
        max = 10
      }
    }
  }
}

output "service_url" {
  value = module.cloud_run_minimal.url
}
```

## Example 2: With Environment Variables

Deployment with plain environment variables:

```hcl
module "cloud_run_with_env" {
  source = "./path/to/module"

  metadata = {
    name = "todo-api"
    org  = "my-org"
    env  = "production"
    labels = {
      team = "platform"
      cost-center = "engineering"
    }
  }

  spec = {
    project_id = "my-gcp-project-123"
    region     = "us-central1"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/todo-api"
        tag  = "v1.2.3"
      }
      env = {
        variables = {
          DATABASE_NAME = "todos"
          DATABASE_HOST = "10.0.0.5"
          LOG_LEVEL     = "info"
          PORT          = "8080"
        }
      }
      port   = 8080
      cpu    = 2
      memory = 1024
      replicas = {
        min = 1
        max = 50
      }
    }

    max_concurrency = 100
    timeout_seconds = 300
  }
}
```

## Example 3: With Secrets from Secret Manager

Deployment with secrets from GCP Secret Manager:

```hcl
module "cloud_run_with_secrets" {
  source = "./path/to/module"

  metadata = {
    name = "secure-api"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    project_id = "my-gcp-project-123"
    region     = "us-central1"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/secure-api"
        tag  = "v2.0.0"
      }
      env = {
        variables = {
          DATABASE_NAME = "production_db"
          ENVIRONMENT   = "production"
        }
        secrets = {
          DATABASE_PASSWORD = "projects/123456789/secrets/db-password/versions/latest"
          API_KEY           = "projects/123456789/secrets/api-key/versions/1"
          JWT_SECRET        = "projects/123456789/secrets/jwt-secret/versions/latest"
        }
      }
      cpu    = 2
      memory = 2048
      replicas = {
        min = 2
        max = 100
      }
    }

    max_concurrency = 80
    timeout_seconds = 600
  }
}
```

## Example 4: Private Service with Custom Service Account

Internal-only service with a custom service account:

```hcl
module "cloud_run_private" {
  source = "./path/to/module"

  metadata = {
    name = "internal-api"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    project_id      = "my-gcp-project-123"
    region          = "us-central1"
    service_account = "cloud-run-sa@my-gcp-project-123.iam.gserviceaccount.com"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/internal-api"
        tag  = "v1.0.0"
      }
      cpu    = 1
      memory = 512
      replicas = {
        min = 1
        max = 20
      }
    }

    ingress               = "INGRESS_TRAFFIC_INTERNAL_ONLY"
    allow_unauthenticated = false
  }
}
```

## Example 5: With VPC Access (Direct VPC Egress)

Service with access to private VPC resources:

```hcl
module "cloud_run_vpc" {
  source = "./path/to/module"

  metadata = {
    name = "vpc-enabled-api"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    project_id = "my-gcp-project-123"
    region     = "us-central1"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/api"
        tag  = "v1.5.0"
      }
      env = {
        variables = {
          REDIS_HOST = "10.0.1.10"  # Private VPC IP
          DB_HOST    = "10.0.1.20"  # Private VPC IP
        }
      }
      cpu    = 2
      memory = 1024
      replicas = {
        min = 1
        max = 50
      }
    }

    vpc_access = {
      network = "projects/my-gcp-project-123/global/networks/default"
      subnet  = "projects/my-gcp-project-123/regions/us-central1/subnetworks/default"
      egress  = "PRIVATE_RANGES_ONLY"
    }
  }
}
```

## Example 6: Production-Ready with Custom DNS

Full production configuration with custom domain:

```hcl
module "cloud_run_production" {
  source = "./path/to/module"

  metadata = {
    name = "production-api"
    org  = "my-org"
    env  = "production"
    labels = {
      team        = "platform"
      service     = "api"
      criticality = "high"
    }
  }

  spec = {
    project_id      = "my-gcp-project-123"
    region          = "us-central1"
    service_name    = "api"
    service_account = "api-sa@my-gcp-project-123.iam.gserviceaccount.com"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/production-api"
        tag  = "v3.2.1"
      }
      env = {
        variables = {
          ENVIRONMENT = "production"
          LOG_LEVEL   = "warn"
        }
        secrets = {
          DATABASE_URL = "projects/123456789/secrets/prod-db-url/versions/latest"
          API_KEY      = "projects/123456789/secrets/prod-api-key/versions/latest"
        }
      }
      port   = 8080
      cpu    = 4
      memory = 4096
      replicas = {
        min = 3
        max = 100
      }
    }

    max_concurrency       = 100
    timeout_seconds       = 300
    ingress               = "INGRESS_TRAFFIC_ALL"
    allow_unauthenticated = true
    execution_environment = "EXECUTION_ENVIRONMENT_GEN2"

    vpc_access = {
      network = "projects/my-gcp-project-123/global/networks/prod-vpc"
      subnet  = "projects/my-gcp-project-123/regions/us-central1/subnetworks/prod-subnet"
      egress  = "PRIVATE_RANGES_ONLY"
    }

    dns = {
      enabled      = true
      hostnames    = ["api.example.com", "api.example.org"]
      managed_zone = "example-com-zone"
    }
  }
}

output "production_api_url" {
  value = module.cloud_run_production.url
}

output "production_api_revision" {
  value = module.cloud_run_production.revision
}
```

## Example 7: GEN1 Execution Environment

Service using the first-generation execution environment:

```hcl
module "cloud_run_gen1" {
  source = "./path/to/module"

  metadata = {
    name = "legacy-api"
    org  = "my-org"
    env  = "staging"
  }

  spec = {
    project_id = "my-gcp-project-123"
    region     = "us-central1"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/legacy-api"
        tag  = "v0.9.0"
      }
      cpu    = 1
      memory = 512
      replicas = {
        min = 0
        max = 10
      }
    }

    execution_environment = "EXECUTION_ENVIRONMENT_GEN1"
  }
}
```

## Example 8: High-Concurrency Configuration

Service optimized for high concurrent requests:

```hcl
module "cloud_run_high_concurrency" {
  source = "./path/to/module"

  metadata = {
    name = "high-throughput-api"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    project_id = "my-gcp-project-123"
    region     = "us-central1"

    container = {
      image = {
        repo = "us-docker.pkg.dev/my-project/my-repo/high-throughput-api"
        tag  = "v2.0.0"
      }
      cpu    = 4
      memory = 8192
      replicas = {
        min = 5
        max = 200
      }
    }

    max_concurrency = 250  # Higher than default
    timeout_seconds = 60   # Short timeout for fast requests
  }
}
```

## Provider Configuration

Don't forget to configure the GCP provider in your root Terraform configuration:

```hcl
terraform {
  required_version = ">= 1.0"

  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = "my-gcp-project-123"
  region  = "us-central1"
}
```

## Notes

1. **Image Repository**: Use Artifact Registry (`pkg.dev`) instead of the legacy Container Registry (`gcr.io`)
2. **CPU Values**: Only 1, 2, or 4 vCPU are allowed
3. **Memory Range**: Must be between 128 MiB and 32768 MiB (32 GiB)
4. **Max Concurrency**: Must be between 1 and 1000
5. **Timeout**: Must be between 1 and 3600 seconds
6. **Scale to Zero**: Set `replicas.min = 0` to enable scale-to-zero for cost savings
7. **Public Access**: Set `allow_unauthenticated = false` for private services
8. **Secrets**: Use GCP Secret Manager format: `projects/{project}/secrets/{secret}/versions/{version}`

