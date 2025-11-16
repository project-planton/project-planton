# Terraform Examples for ElasticsearchKubernetes

This document provides examples of how to use the Terraform module for deploying Elasticsearch on Kubernetes.

---

## Example 1: Basic Elasticsearch and Kibana Deployment

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ElasticsearchKubernetes
metadata:
  name: logging-cluster
spec:
  kubernetesProviderConfigId: my-k8s-credentials
  elasticsearch:
    resources:
      requests:
        cpu: 500m
        memory: 1Gi
      limits:
        cpu: 1000m
        memory: 2Gi
  kibana:
    resources:
      requests:
        cpu: 200m
        memory: 512Mi
      limits:
        cpu: 500m
        memory: 1Gi
```

### Terraform Usage

```hcl
module "logging_cluster" {
  source = "../"

  metadata = {
    name = "logging-cluster"
  }

  spec = {
    kubernetesProviderConfigId = "my-k8s-credentials"
    elasticsearch = {
      resources = {
        requests = {
          cpu    = "500m"
          memory = "1Gi"
        }
        limits = {
          cpu    = "1000m"
          memory = "2Gi"
        }
      }
    }
    kibana = {
      resources = {
        requests = {
          cpu    = "200m"
          memory = "512Mi"
        }
        limits = {
          cpu    = "500m"
          memory = "1Gi"
        }
      }
    }
  }
}
```

---

## Example 2: Elasticsearch with Ingress Enabled

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ElasticsearchKubernetes
metadata:
  name: search-service
spec:
  kubernetesProviderConfigId: my-k8s-credentials
  elasticsearch:
    container:
      resources:
        requests:
          cpu: 1
          memory: 2Gi
        limits:
          cpu: 2
          memory: 4Gi
    ingress:
      enabled: true
      hostname: search.example.com
  kibana:
    enabled: true
    container:
      resources:
        requests:
          cpu: 200m
          memory: 512Mi
        limits:
          cpu: 500m
          memory: 1Gi
    ingress:
      enabled: true
      hostname: search-kibana.example.com
```

### Terraform Usage

```hcl
module "search_service" {
  source = "../"

  metadata = {
    name = "search-service"
  }

  spec = {
    kubernetesProviderConfigId = "my-k8s-credentials"
    elasticsearch = {
      container = {
        resources = {
          requests = {
            cpu    = "1"
            memory = "2Gi"
          }
          limits = {
            cpu    = "2"
            memory = "4Gi"
          }
        }
      }
      ingress = {
        enabled  = true
        hostname = "search.example.com"
      }
    }
    kibana = {
      enabled = true
      container = {
        resources = {
          requests = {
            cpu    = "200m"
            memory = "512Mi"
          }
          limits = {
            cpu    = "500m"
            memory = "1Gi"
          }
        }
      }
      ingress = {
        enabled  = true
        hostname = "search-kibana.example.com"
      }
    }
  }
}
```

---

## Example 3: Elasticsearch with Persistent Storage

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ElasticsearchKubernetes
metadata:
  name: persistent-cluster
spec:
  kubernetesProviderConfigId: my-k8s-credentials
  elasticsearch:
    container:
      persistenceEnabled: true
      diskSize: 20Gi
      replicas: 3
      resources:
        requests:
          cpu: 1
          memory: 2Gi
        limits:
          cpu: 2
          memory: 4Gi
    ingress:
      enabled: true
      hostname: elasticsearch.example.com
  kibana:
    enabled: true
    container:
      replicas: 1
      resources:
        requests:
          cpu: 200m
          memory: 512Mi
        limits:
          cpu: 500m
          memory: 1Gi
    ingress:
      enabled: true
      hostname: kibana.example.com
```

### Terraform Usage

```hcl
module "persistent_cluster" {
  source = "../"

  metadata = {
    name = "persistent-cluster"
  }

  spec = {
    kubernetesProviderConfigId = "my-k8s-credentials"
    elasticsearch = {
      container = {
        persistenceEnabled = true
        diskSize          = "20Gi"
        replicas          = 3
        resources = {
          requests = {
            cpu    = "1"
            memory = "2Gi"
          }
          limits = {
            cpu    = "2"
            memory = "4Gi"
          }
        }
      }
      ingress = {
        enabled  = true
        hostname = "elasticsearch.example.com"
      }
    }
    kibana = {
      enabled = true
      container = {
        replicas = 1
        resources = {
          requests = {
            cpu    = "200m"
            memory = "512Mi"
          }
          limits = {
            cpu    = "500m"
            memory = "1Gi"
          }
        }
      }
      ingress = {
        enabled  = true
        hostname = "kibana.example.com"
      }
    }
  }
}
```

---

## Example 4: Minimal Elasticsearch Deployment

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: ElasticsearchKubernetes
metadata:
  name: minimal-elasticsearch
spec:
  kubernetesProviderConfigId: my-k8s-credentials
```

### Terraform Usage

```hcl
module "minimal_elasticsearch" {
  source = "../"

  metadata = {
    name = "minimal-elasticsearch"
  }

  spec = {
    kubernetesProviderConfigId = "my-k8s-credentials"
  }
}
```

This minimal example uses all the default values configured in the Terraform module.

---

## Running Terraform

To use any of these examples:

1. Save the Terraform code to a `.tf` file (e.g., `main.tf`)
2. Initialize Terraform:

```bash
terraform init
```

3. Review the planned changes:

```bash
terraform plan
```

4. Apply the configuration:

```bash
terraform apply
```

5. To destroy the resources:

```bash
terraform destroy
```

---

## Output Variables

After applying, you can access the following outputs:

```hcl
output "elasticsearch_endpoint" {
  value = module.search_service.elasticsearch_endpoint
}

output "kibana_endpoint" {
  value = module.search_service.kibana_endpoint
}

output "port_forward_command_elasticsearch" {
  value = module.search_service.port_forward_command_elasticsearch
}

output "port_forward_command_kibana" {
  value = module.search_service.port_forward_command_kibana
}
```

---

## Notes

- All examples assume you have a valid `kubernetesProviderConfigId` configured
- Adjust resource limits and requests based on your workload requirements
- For production deployments, always enable persistence and use multiple replicas
- Ingress hostnames should be valid DNS names that resolve to your Kubernetes cluster's ingress controller

