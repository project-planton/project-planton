# Example 1: Basic MongoDB Atlas Cluster Setup

This example demonstrates a basic setup for a MongoDB Atlas cluster with minimal configuration. It provisions a replica set cluster with default values for node distribution and MongoDB version.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbAtlas
metadata:
  name: basic-mongo-cluster
spec:
  mongodb_atlas_credential_id: my-mongodb-atlas-cred
  cluster_config:
    project_id: my-mongodb-project
    cluster_type: REPLICASET
    electable_nodes: 3
    priority: 7
    read_only_nodes: 0
    cloud_backup: true
    auto_scaling_disk_gb_enabled: true
    mongo_db_major_version: 6.0
    provider_name: AWS
    provider_instance_size_name: M10
```

# Example 2: Sharded Cluster with Read-Only Nodes

This example sets up a sharded MongoDB Atlas cluster with both electable and read-only nodes. Auto-scaling and cloud backup are enabled for the cluster.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbAtlas
metadata:
  name: sharded-mongo-cluster
spec:
  mongodb_atlas_credential_id: atlas-cred-123
  cluster_config:
    project_id: my-sharded-cluster-project
    cluster_type: SHARDED
    electable_nodes: 5
    priority: 7
    read_only_nodes: 2
    cloud_backup: true
    auto_scaling_disk_gb_enabled: true
    mongo_db_major_version: 6.0
    provider_name: GCP
    provider_instance_size_name: M30
```

# Example 3: Geo-sharded Cluster on Azure

This example provisions a geo-sharded MongoDB Atlas cluster on Microsoft Azure, using a specific MongoDB version and enabling cloud backups. This setup is ideal for distributing data across different geographical regions for high availability.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbAtlas
metadata:
  name: geo-sharded-cluster
spec:
  mongodb_atlas_credential_id: geo-cred-567
  cluster_config:
    project_id: geo-project-id
    cluster_type: GEOSHARDED
    electable_nodes: 7
    priority: 7
    read_only_nodes: 3
    cloud_backup: true
    auto_scaling_disk_gb_enabled: true
    mongo_db_major_version: 5.0
    provider_name: AZURE
    provider_instance_size_name: M50
```

# Example 4: Minimal Replica Set Cluster

This is an example of a minimal configuration for deploying a MongoDB Atlas replica set cluster. It skips the optional read-only nodes and disables cloud backups.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbAtlas
metadata:
  name: minimal-replicaset-cluster
spec:
  mongodb_atlas_credential_id: minimal-atlas-cred
  cluster_config:
    project_id: minimal-project-id
    cluster_type: REPLICASET
    electable_nodes: 3
    priority: 7
    read_only_nodes: 0
    cloud_backup: false
    auto_scaling_disk_gb_enabled: false
    mongo_db_major_version: 7.0
    provider_name: AWS
    provider_instance_size_name: M10
```

# Example 5: Sharded Cluster with Disk Auto-Scaling Disabled

This example demonstrates a MongoDB Atlas sharded cluster configuration with auto-scaling disabled, providing a more controlled approach to storage management.

```yaml
apiVersion: code2cloud.planton.cloud/v1
kind: MongodbAtlas
metadata:
  name: sharded-cluster-no-autoscaling
spec:
  mongodb_atlas_credential_id: atlas-shard-cred
  cluster_config:
    project_id: shard-project-id
    cluster_type: SHARDED
    electable_nodes: 5
    priority: 7
    read_only_nodes: 1
    cloud_backup: true
    auto_scaling_disk_gb_enabled: false
    mongo_db_major_version: 6.0
    provider_name: AWS
    provider_instance_size_name: M20
```

---
