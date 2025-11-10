# GCP Cloud SQL API - Examples

Here are examples of how to create and configure a **GcpCloudSql** API resource using the Project Planton CLI. The examples cover MySQL and PostgreSQL instances with various configuration options.

## Create using CLI

First, create a YAML file using the examples provided below. After the YAML file is created, you can apply the configuration using the following command:

```shell
project-planton pulumi up --manifest <yaml-path> --stack <org>/<stack-name>/<environment>
```

Or using Terraform:

```shell
project-planton tofu apply --manifest <yaml-path> --auto-approve
```

## Basic MySQL Example

This example demonstrates how to create a basic MySQL 8.0 Cloud SQL instance with minimal configuration.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: mysql-db
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: MYSQL
  database_version: MYSQL_8_0
  tier: db-n1-standard-1
  storage_gb: 10
  root_password: MySecurePassword123!
```

## Basic PostgreSQL Example

This example demonstrates how to create a basic PostgreSQL 15 Cloud SQL instance.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: postgres-db
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: POSTGRESQL
  database_version: POSTGRES_15
  tier: db-n1-standard-2
  storage_gb: 20
  root_password: MySecurePassword123!
```

## MySQL with Private IP and VPC

This example creates a MySQL instance with private IP connectivity through a VPC network.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: mysql-private
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: MYSQL
  database_version: MYSQL_8_0
  tier: db-n1-standard-1
  storage_gb: 10
  network:
    vpc_id: projects/my-gcp-project/global/networks/my-vpc
    private_ip_enabled: true
  root_password: MySecurePassword123!
```

## PostgreSQL with Public IP and Authorized Networks

This example creates a PostgreSQL instance with public IP access restricted to specific CIDR ranges.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: postgres-public
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: POSTGRESQL
  database_version: POSTGRES_15
  tier: db-n1-standard-1
  storage_gb: 10
  network:
    authorized_networks:
      - 203.0.113.0/24
      - 198.51.100.0/24
  root_password: MySecurePassword123!
```

## MySQL with High Availability

This example creates a highly available MySQL instance with automatic failover to a secondary zone.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: mysql-ha
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: MYSQL
  database_version: MYSQL_8_0
  tier: db-n1-standard-2
  storage_gb: 50
  high_availability:
    enabled: true
    zone: us-central1-b
  root_password: MySecurePassword123!
```

## PostgreSQL with Automated Backups

This example creates a PostgreSQL instance with automated daily backups configured.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: postgres-backup
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: POSTGRESQL
  database_version: POSTGRES_15
  tier: db-n1-standard-1
  storage_gb: 20
  backup:
    enabled: true
    start_time: "03:00"
    retention_days: 7
  root_password: MySecurePassword123!
```

## Production-Grade PostgreSQL Instance

This comprehensive example creates a production-ready PostgreSQL instance with high availability, private networking, automated backups, and custom database flags.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: postgres-production
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: POSTGRESQL
  database_version: POSTGRES_15
  tier: db-n1-highmem-4
  storage_gb: 100
  network:
    vpc_id: projects/my-gcp-project/global/networks/production-vpc
    private_ip_enabled: true
  high_availability:
    enabled: true
    zone: us-central1-c
  backup:
    enabled: true
    start_time: "02:00"
    retention_days: 30
  database_flags:
    max_connections: "200"
    shared_buffers: "262144"
    effective_cache_size: "786432"
  root_password: ProductionSecurePassword123!
```

## MySQL with Database Flags

This example creates a MySQL instance with custom database configuration flags.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: mysql-custom
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: MYSQL
  database_version: MYSQL_8_0
  tier: db-n1-standard-2
  storage_gb: 30
  database_flags:
    max_connections: "150"
    innodb_buffer_pool_size: "268435456"
    slow_query_log: "on"
  root_password: MySecurePassword123!
```

## Large Storage PostgreSQL Instance

This example creates a PostgreSQL instance optimized for large datasets.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: postgres-large-storage
spec:
  projectId: my-gcp-project
  region: us-central1
  database_engine: POSTGRESQL
  database_version: POSTGRES_15
  tier: db-n1-highmem-8
  storage_gb: 1000
  network:
    vpc_id: projects/my-gcp-project/global/networks/data-vpc
    private_ip_enabled: true
  backup:
    enabled: true
    start_time: "01:00"
    retention_days: 14
  root_password: MySecurePassword123!
```

## Development MySQL Instance

This example creates a small, cost-effective MySQL instance suitable for development environments.

```yaml
apiVersion: gcp.project-planton.org/v1
kind: GcpCloudSql
metadata:
  name: mysql-dev
spec:
  projectId: my-dev-project
  region: us-central1
  database_engine: MYSQL
  database_version: MYSQL_8_0
  tier: db-f1-micro
  storage_gb: 10
  root_password: DevPassword123!
```

## Notes

- **Root Password**: Always use strong passwords with at least 8 characters including letters, numbers, and special characters.
- **Private IP**: Requires VPC peering to be configured between your VPC and the Cloud SQL service.
- **High Availability**: Increases cost but provides better uptime and automatic failover.
- **Backup Retention**: Longer retention periods increase storage costs but provide more recovery options.
- **Database Flags**: Refer to MySQL or PostgreSQL documentation for available flags and their values.
- **Instance Tiers**: Choose appropriate tier based on CPU and memory requirements of your workload.
- **Storage**: Start with minimum required and enable automatic storage increase in production.

