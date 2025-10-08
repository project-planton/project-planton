# ClickHouse Kubernetes Terraform Module - Examples

This document provides examples demonstrating various ClickHouse cluster configurations using the Altinity ClickHouse Operator.

## Prerequisites

The **Altinity ClickHouse Operator** must be installed on your Kubernetes cluster before deploying these examples.

---

## Example 1: Minimal Standalone Configuration

```hcl
module "clickhouse_basic" {
  source = "../../tf"

  metadata = {
    name = "basic-clickhouse"
  }

  spec = {
    cluster_name = "dev-cluster"
    
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "20Gi"
      resources = {
        requests = {
          cpu    = "500m"
          memory = "1Gi"
        }
        limits = {
          cpu    = "2000m"
          memory = "4Gi"
        }
      }
    }
  }
}
```

---

## Example 2: Production with Version Pinning

```hcl
module "clickhouse_production" {
  source = "../../tf"

  metadata = {
    name = "prod-clickhouse"
    org  = "my-org"
    env  = "production"
  }

  spec = {
    cluster_name = "production-analytics"
    version      = "24.8"
    
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "200Gi"
      resources = {
        requests = {
          cpu    = "2000m"
          memory = "8Gi"
        }
        limits = {
          cpu    = "8000m"
          memory = "32Gi"
        }
      }
    }
  }
}
```

---

## Example 3: Distributed Cluster with Sharding

```hcl
module "clickhouse_distributed" {
  source = "../../tf"

  metadata = {
    name = "distributed-analytics"
  }

  spec = {
    cluster_name = "analytics-cluster"
    
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "100Gi"
      resources = {
        requests = {
          cpu    = "1000m"
          memory = "4Gi"
        }
        limits = {
          cpu    = "4000m"
          memory = "16Gi"
        }
      }
    }
    
    cluster = {
      is_enabled    = true
      shard_count   = 4
      replica_count = 1
    }
  }
}
```

**Key points**:
- 4 shards for horizontal scaling
- ZooKeeper automatically managed by operator
- Distributed query execution

---

## Example 4: High Availability Cluster

```hcl
module "clickhouse_ha" {
  source = "../../tf"

  metadata = {
    name = "ha-clickhouse"
    env  = "production"
  }

  spec = {
    cluster_name = "ha-analytics"
    version      = "24.8"
    
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "150Gi"
      resources = {
        requests = {
          cpu    = "2000m"
          memory = "8Gi"
        }
        limits = {
          cpu    = "6000m"
          memory = "24Gi"
        }
      }
    }
    
    cluster = {
      is_enabled    = true
      shard_count   = 3
      replica_count = 3
    }
  }
}
```

**Key points**:
- 3 shards with 3 replicas each (9 total nodes)
- High availability - survives 2 node failures per shard
- Data replicated across nodes

---

## Example 5: Cluster with External ZooKeeper

```hcl
module "clickhouse_external_zk" {
  source = "../../tf"

  metadata = {
    name = "enterprise-clickhouse"
  }

  spec = {
    cluster_name = "enterprise-cluster"
    
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "200Gi"
      resources = {
        requests = {
          cpu    = "4000m"
          memory = "16Gi"
        }
        limits = {
          cpu    = "12000m"
          memory = "48Gi"
        }
      }
    }
    
    cluster = {
      is_enabled    = true
      shard_count   = 6
      replica_count = 2
    }
    
    zookeeper = {
      use_external = true
      nodes = [
        "zk-0.zookeeper.default.svc.cluster.local:2181",
        "zk-1.zookeeper.default.svc.cluster.local:2181",
        "zk-2.zookeeper.default.svc.cluster.local:2181"
      ]
    }
  }
}
```

**Key points**:
- External ZooKeeper for shared coordination
- 6 shards with 2 replicas (12 total nodes)
- Enterprise-scale resources

---

## Example 6: With Ingress for External Access

```hcl
module "clickhouse_public" {
  source = "../../tf"

  metadata = {
    name = "public-clickhouse"
  }

  spec = {
    cluster_name = "public-cluster"
    
    container = {
      replicas               = 1
      is_persistence_enabled = true
      disk_size              = "100Gi"
      resources = {
        requests = {
          cpu    = "1000m"
          memory = "4Gi"
        }
        limits = {
          cpu    = "4000m"
          memory = "16Gi"
        }
      }
    }
    
    ingress = {
      is_enabled = true
      dns_domain = "example.com"
    }
  }
}
```

**Key points**:
- LoadBalancer service with external DNS
- External hostname: `public-clickhouse.example.com`
- Both HTTP (8123) and native (9000) ports exposed

---

## Deployment Architecture

All deployments use the **Altinity ClickHouse Operator** which:
- Creates ClickHouseInstallation custom resources
- Manages StatefulSets, Services, and ConfigMaps
- Handles ZooKeeper coordination for clusters
- Provides rolling upgrades and self-healing
- Uses official ClickHouse container images

