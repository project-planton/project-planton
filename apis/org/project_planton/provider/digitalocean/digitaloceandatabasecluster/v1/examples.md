# DigitalOcean Database Cluster Examples

Complete, copy-paste ready YAML manifests for common database deployment scenarios.

---

## Example 1: Development PostgreSQL (Single Node)

**Use Case**: Non-production PostgreSQL for local development or testing.

```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: dev-postgres
spec:
  cluster_name: dev-postgres
  engine: postgres
  engine_version: "16"
  region: nyc3
  size_slug: db-s-1vcpu-1gb
  node_count: 1
  enable_public_connectivity: true  # For dev access from local machine
```

**Notes:**
- Single node = no HA (acceptable for dev/test)
- Public connectivity enabled for ease of access
- Smallest/cheapest size slug ($15/month)

---

## Example 2: Production PostgreSQL (High Availability with VPC)

**Use Case**: Mission-critical PostgreSQL with automatic failover.

```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: prod-postgres
spec:
  cluster_name: prod-postgres
  engine: postgres
  engine_version: "16"
  region: nyc3
  size_slug: db-s-4vcpu-8gb
  node_count: 3  # Primary + 2 standbys for maximum HA
  vpc:
    value: "12345678-1234-1234-1234-123456789012"
  storage_gib: 200
  enable_public_connectivity: false
```

**Notes:**
- 3 nodes for maximum resilience
- VPC-private (no public access)
- Custom storage (200GB instead of default)
- Production-sized instance (4 vCPU, 8GB RAM)

**Required Follow-Up:**
1. Configure PgBouncer connection pool (separate resource)
2. Set up firewall rules (tag-based or VPC CIDR)
3. Configure application with private_uri output

---

## Example 3: MySQL for Web Applications

**Use Case**: MySQL cluster for content management systems or e-commerce.

```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: prod-mysql
spec:
  cluster_name: prod-mysql
  engine: mysql
  engine_version: "8"
  region: sfo3
  size_slug: db-s-2vcpu-4gb
  node_count: 2
  vpc:
    value: "vpc-87654321"
  storage_gib: 100
  enable_public_connectivity: false
```

**Notes:**
- MySQL 8 (latest stable)
- 2-node HA configuration
- San Francisco region for US West users
- Moderate storage for typical web app

**MySQL Gotchas:**
- No SUPER privileges (cannot SET GLOBAL variables)
- Restricted mysqldump (use DigitalOcean backups instead)
- See research doc for migration considerations

---

## Example 4: Redis Cache Cluster

**Use Case**: In-memory caching layer for application performance.

```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: prod-redis-cache
spec:
  cluster_name: prod-redis-cache
  engine: redis
  engine_version: "7"
  region: nyc3
  size_slug: db-s-2vcpu-4gb
  node_count: 2
  vpc:
    value: "vpc-cache-12345"
  enable_public_connectivity: false
```

**Notes:**
- Redis 7 (actually Valkey due to SSPL licensing)
- HA configuration (primary + standby)
- VPC-private for security
- No clustering mode (single-instance limit)

**Redis Limitations:**
- No Redis Cluster mode support
- Persistence is available but Redis should be treated as ephemeral
- Use for caching/sessions, not primary data store

---

## Example 5: MongoDB Document Store (High Availability)

**Use Case**: Document-oriented database for flexible schema applications.

```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: prod-mongodb
spec:
  cluster_name: prod-mongodb
  engine: mongodb
  engine_version: "7.0"
  region: fra1
  size_slug: db-s-4vcpu-8gb
  node_count: 3
  vpc:
    value: "vpc-eu-67890"
  storage_gib: 500
  enable_public_connectivity: false
```

**Notes:**
- MongoDB 7.0 (latest)
- 3-node replica set for HA
- Frankfurt region for EU data residency
- Large storage (500GB) for document collections

**MongoDB Limitations:**
- No sharding support (replica sets only)
- Max 3 nodes (cannot scale beyond this)

---

## Example 6: Production PostgreSQL with VPC Reference

**Use Case**: Reference VPC from another Project Planton resource.

```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: app-database
spec:
  cluster_name: app-database
  engine: postgres
  engine_version: "15"
  region: nyc3
  size_slug: db-s-4vcpu-8gb
  node_count: 3
  vpc:
    ref:
      kind: DigitalOceanVpc
      name: production-vpc
      field_path: status.outputs.vpc_id
  storage_gib: 250
  enable_public_connectivity: false
```

**Notes:**
- Uses foreign key reference to DigitalOceanVpc resource
- field_path extracts vpc_id from VPC resource status
- Ensures database and VPC are deployed in dependency order

---

## Example 7: Multi-Environment Setup

**Use Case**: Consistent configuration across dev/staging/prod with parameter differences.

**Development:**
```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: dev-app-db
spec:
  cluster_name: dev-app-db
  engine: postgres
  engine_version: "16"
  region: nyc3
  size_slug: db-s-1vcpu-1gb
  node_count: 1
  enable_public_connectivity: true
```

**Staging:**
```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: staging-app-db
spec:
  cluster_name: staging-app-db
  engine: postgres
  engine_version: "16"
  region: nyc3
  size_slug: db-s-2vcpu-4gb
  node_count: 2
  vpc:
    value: "vpc-staging-123"
  enable_public_connectivity: false
```

**Production:**
```yaml
apiVersion: digitalocean.project-planton.org/v1
kind: DigitalOceanDatabaseCluster
metadata:
  name: prod-app-db
spec:
  cluster_name: prod-app-db
  engine: postgres
  engine_version: "16"
  region: nyc3
  size_slug: db-s-4vcpu-8gb
  node_count: 3
  vpc:
    value: "vpc-prod-456"
  storage_gib: 300
  enable_public_connectivity: false
```

**Pattern Highlights:**
- Dev: Single node, public access, smallest size
- Staging: 2 nodes (HA), VPC-private, moderate size
- Production: 3 nodes (max HA), VPC-private, large storage

---

## Configuration Patterns Summary

| Use Case | Engine | Nodes | Size | Public | VPC | Monthly Cost |
|----------|--------|-------|------|--------|-----|--------------|
| Dev/Test | Any | 1 | db-s-1vcpu-1gb | Yes | Optional | ~$15 |
| Small Prod | postgres/mysql | 2 | db-s-2vcpu-4gb | No | Required | ~$120 |
| Large Prod | Any | 3 | db-s-4vcpu-8gb | No | Required | ~$360 |
| Cache Layer | redis | 2 | db-s-2vcpu-4gb | No | Required | ~$120 |
| Document Store | mongodb | 3 | db-s-4vcpu-8gb | No | Required | ~$360 |

---

## Deployment Checklist

Before deploying, ensure:

- ✅ `cluster_name` is unique in your DigitalOcean account
- ✅ `engine_version` is supported (check DigitalOcean docs for current versions)
- ✅ `size_slug` is valid for selected engine
- ✅ `node_count` ≥ 2 for production (HA requirement)
- ✅ `vpc` is specified for production deployments
- ✅ `enable_public_connectivity: false` for production
- ✅ `storage_gib` accounts for expected data growth (cannot shrink)

---

## Next Steps After Deployment

1. **Retrieve Connection Details**:
   ```bash
   project-planton get outputs <cluster-name>
   # Outputs: connection_uri, host, port, username, password
   ```

2. **Configure Firewall** (separate resource):
   ```yaml
   apiVersion: digitalocean.project-planton.org/v1
   kind: DigitalOceanDatabaseFirewall
   metadata:
     name: prod-db-firewall
   spec:
     cluster_id: ${cluster_id_from_outputs}
     rules:
       - type: vpc
         value: "10.10.0.0/16"
   ```

3. **Set Up PgBouncer** (PostgreSQL only):
   ```yaml
   apiVersion: digitalocean.project-planton.org/v1
   kind: DigitalOceanDatabaseConnectionPool
   metadata:
     name: app-pool
   spec:
     cluster_id: ${cluster_id}
     name: app-pool
     mode: transaction
     size: 25
   ```

4. **Test Connectivity**:
   ```bash
   # PostgreSQL
   psql "${connection_uri}"
   
   # MySQL
   mysql -h ${host} -P ${port} -u ${username} -p
   
   # Redis
   redis-cli -h ${host} -p ${port} -a ${password}
   
   # MongoDB
   mongosh "${connection_uri}"
   ```

---

For more details, see:
- [README.md](./README.md) - Component overview and best practices
- [docs/README.md](./docs/README.md) - Comprehensive production guide
- [iac/pulumi/README.md](./iac/pulumi/README.md) - Pulumi module usage
- [iac/tf/README.md](./iac/tf/README.md) - Terraform module usage

