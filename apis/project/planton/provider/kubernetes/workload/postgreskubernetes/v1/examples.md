# Multiple Examples for `PostgresKubernetes` API-Resource

## Example: Basic PostgreSQL Database

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
project-planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: my-app-db
  org: my-org
  env: dev
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: "250m"
        memory: "256Mi"
      limits:
        cpu: "1000m"
        memory: "1Gi"
    diskSize: "10Gi"
```

---

## Example: PostgreSQL with Ingress

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
project-planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: my-public-db
  org: my-org
  env: prod
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: "500m"
        memory: "512Mi"
      limits:
        cpu: "2000m"
        memory: "2Gi"
    diskSize: "50Gi"
  
  ingress:
    enabled: true
    hostname: postgres-prod.example.com
```

---

## Example: PostgreSQL with Custom Backup Configuration

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
project-planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: my-app-db
  org: my-org
  env: prod
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: "500m"
        memory: "512Mi"
      limits:
        cpu: "2000m"
        memory: "2Gi"
    diskSize: "100Gi"
  
  # Custom backup configuration (overrides operator-level defaults)
  backup_config:
    # Custom S3 prefix for this database's backups
    # $(PGVERSION) will be replaced by the PostgreSQL version
    s3_prefix: "backups/critical/my-app-db/$(PGVERSION)"
    
    # Custom backup schedule (every 6 hours instead of operator default)
    backup_schedule: "0 */6 * * *"
    
    # Explicitly enable backups for this database
    enable_backup: true
```

---

## Example: Disaster Recovery - Restore from Backup (Stage 1: Standby)

### Create using CLI

This example demonstrates cross-cluster disaster recovery by restoring a database from R2/S3 backups.

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
project-planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: restored-db
  org: my-org
  env: prod
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: "500m"
        memory: "512Mi"
      limits:
        cpu: "2000m"
        memory: "2Gi"
    diskSize: "100Gi"
  
  # Disaster recovery configuration
  backup_config:
    restore:
      # Stage 1: Bootstrap from backup (read-only standby)
      enabled: true
      
      # S3/R2 bucket containing the backup
      bucketName: "my-db-backups-prod"
      
      # Path to backup directory (without s3:// prefix or bucket name)
      # This path should contain basebackups_005/ and wal_005/ directories
      s3_path: "backups/source-db-name/14"
      
      # R2/S3 credentials for restore access
      # Allows independent disaster recovery without operator dependencies
      r2_config:
        cloudflare_account_id: "your-account-id"
        access_key_id: "your-r2-access-key"
        secret_access_key: "your-r2-secret-key"
```

### Verification

After deployment, verify the database is in standby mode:

```shell
# Get the pod name
POD=$(kubectl get pods -n <namespace> -l application=spilo -o jsonpath='{.items[0].metadata.name}')

# Check Patroni status (should show "Standby Leader" or "Replica")
kubectl exec -n <namespace> $POD -- patronictl list

# Verify read-only mode (should return 't' for true)
kubectl exec -n <namespace> $POD -- psql -U postgres -c "SELECT pg_is_in_recovery();"

# Test read-only enforcement (should FAIL)
kubectl exec -n <namespace> $POD -- psql -U postgres -c "CREATE TABLE test (id int);"
# Expected: ERROR: cannot execute CREATE TABLE in a read-only transaction
```

---

## Example: Disaster Recovery - Promote to Primary (Stage 2)

### Update Configuration

After validating the restored data in Stage 1, promote the database to primary by updating the manifest.

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: restored-db
  org: my-org
  env: prod
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: "500m"
        memory: "512Mi"
      limits:
        cpu: "2000m"
        memory: "2Gi"
    diskSize: "100Gi"
  
  # Disaster recovery configuration
  backup_config:
    restore:
      # Stage 2: Promote to primary (read-write)
      enabled: false  # Changed from true
      
      # Other fields can be kept for documentation or removed entirely
      bucketName: "my-db-backups-prod"
      s3_path: "backups/source-db-name/14"
```

### Deploy Promotion

```shell
project-planton apply -f <yaml-path>
```

### Verification

After promotion, verify the database is now a read-write primary:

```shell
# Get the pod name
POD=$(kubectl get pods -n <namespace> -l application=spilo -o jsonpath='{.items[0].metadata.name}')

# Check Patroni status (should show "Leader", Timeline advanced to 2)
kubectl exec -n <namespace> $POD -- patronictl list

# Verify read-write mode (should return 'f' for false)
kubectl exec -n <namespace> $POD -- psql -U postgres -c "SELECT pg_is_in_recovery();"

# Test writes (should SUCCEED)
kubectl exec -n <namespace> $POD -- psql -U postgres -c "CREATE TABLE test (id int); INSERT INTO test VALUES (1);"
```

---

## Example: Disaster Recovery with Fallback to Operator Bucket

### Create using CLI

This example uses operator-level bucket configuration as fallback, requiring only the S3 path to be specified.

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
project-planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: restored-db-simple
  org: my-org
  env: prod
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: "500m"
        memory: "512Mi"
      limits:
        cpu: "2000m"
        memory: "2Gi"
    diskSize: "100Gi"
  
  backup_config:
    restore:
      enabled: true
      
      # bucket_name not specified - will use operator-level bucket
      # r2_config not specified - will use operator-level credentials (if supported by operator)
      
      # Only S3 path is required
      s3_path: "backups/source-db-name/14"
```

**Note**: This approach requires the PostgreSQL operator to be configured with R2/S3 credentials. If operator-level credentials are not available, provide per-database `r2_config` for complete independence.

---

## Example: High Availability PostgreSQL with Backups

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
project-planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: ha-db
  org: my-org
  env: prod
spec:
  container:
    # Multiple replicas for high availability
    replicas: 3
    resources:
      requests:
        cpu: "1000m"
        memory: "2Gi"
      limits:
        cpu: "4000m"
        memory: "8Gi"
    diskSize: "500Gi"
  
  # Custom backup configuration for critical database
  backup_config:
    # Aggressive backup schedule (every 3 hours)
    backup_schedule: "0 */3 * * *"
    
    # Custom S3 prefix for critical data
    s3_prefix: "backups/critical/ha-db/$(PGVERSION)"
    
    # Explicitly enable backups
    enable_backup: true
  
  ingress:
    enabled: true
    hostname: postgres-ha-prod.example.com
```

---

## Example: Development Database with Backups Disabled

### Create using CLI

Create a YAML file using the example shown below. After the YAML is created, use the command below to apply it.

```shell
project-planton apply -f <yaml-path>
```

### YAML Configuration

```yaml
apiVersion: kubernetes.project-planton.org/v1
kind: PostgresKubernetes
metadata:
  name: dev-db
  org: my-org
  env: dev
spec:
  container:
    replicas: 1
    resources:
      requests:
        cpu: "100m"
        memory: "128Mi"
      limits:
        cpu: "500m"
        memory: "512Mi"
    diskSize: "5Gi"
  
  # Disable backups for ephemeral development database
  backup_config:
    enable_backup: false
```

---

## Backup and Restore Best Practices

### Backup Strategy

1. **Operator-Level Configuration**: Configure R2/S3 credentials and default backup schedule at the operator level for centralized management
2. **Per-Database Overrides**: Override backup schedule and S3 prefix for critical databases requiring more frequent backups
3. **Backup Schedule**: Use cron format (e.g., `"0 2 * * *"` for daily at 2 AM UTC)
4. **S3 Prefix Naming**: Use descriptive prefixes with `$(PGVERSION)` variable for automatic version separation

### Disaster Recovery Workflow

**When to Use**:
- Source cluster destroyed or inaccessible
- Cross-cluster failover required
- Testing backup integrity
- Creating database copies for analytics/testing

**Two-Stage Process**:

1. **Stage 1 - Bootstrap as Standby** (`restore.enabled: true`)
   - Database restores from R2/S3 backup
   - Runs in read-only mode
   - Allows data validation before committing to failover
   - Zero risk of accidental writes during validation

2. **Stage 2 - Promote to Primary** (`restore.enabled: false`)
   - Controlled, deliberate promotion decision
   - Database becomes read-write primary
   - Clear audit trail of when failover occurred
   - Can be automated or manual based on confidence

**Expected Restore Times**:
- Small DB (<10GB): 5-10 minutes
- Medium DB (50GB): 20-30 minutes
- Large DB (100GB+): 30-60+ minutes
- Promotion time: <10 seconds (seamless)

### Credentials Management

**Option 1: Per-Database Credentials** (Recommended for DR)
```yaml
backup_config:
  restore:
    enabled: true
    bucketName: "my-backups"
    s3_path: "backups/db-name/14"
    r2_config:
      cloudflare_account_id: "xxx"
      access_key_id: "yyy"
      secret_access_key: "zzz"
```
- Complete independence from operator configuration
- Enables true cross-cluster disaster recovery
- No dependencies on operator-level secrets

**Option 2: Operator-Level Fallback**
```yaml
backup_config:
  restore:
    enabled: true
    # bucket_name omitted - uses operator config
    s3_path: "backups/db-name/14"
    # r2_config omitted - uses operator config
```
- Simpler configuration
- Requires operator to have R2/S3 credentials configured
- May not work for true cross-cluster scenarios

### Operator Compatibility

This API is designed to be technology-agnostic and works with multiple PostgreSQL operators:

| Operator | Restore Implementation | Status |
|----------|------------------------|--------|
| **Zalando** | `spec:standby` with `STANDBY_*` env vars | ✅ Implemented |
| **Percona** | `spec:dataSource` | 🔄 Future |
| **CloudNativePG** | `spec:bootstrap.recovery` | 🔄 Future |

The same API manifest works across operators - only the Pulumi module implementation differs.
