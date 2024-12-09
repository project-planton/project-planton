# MongoDB Atlas Pulumi Module for Planton Cloud

## Key Features

- **Unified API Resource Modeling**: The module utilizes a Kubernetes-style API resource structure, which ensures a standardized format across different cloud infrastructures and services.
  
- **Multi-Cloud Support**: MongoDB Atlas clusters can be deployed on AWS, GCP, and Azure, with support for various instance types and cluster configurations.

- **Cluster Configuration Flexibility**: The module supports multiple cluster types (Replica Set, Sharded, Geo-sharded) and allows configuration of electable and read-only nodes, cloud backup options, and MongoDB versions. Developers can fine-tune cluster settings to match specific requirements.

- **Automation via CLI**: The Planton CLI enables developers to deploy infrastructure with a single command (`planton pulumi up --stack-input <api-resource.yaml>`), either specifying a custom Pulumi module via a Git repository or using the default Pulumi module. This makes it easy to integrate into CI/CD pipelines and automate infrastructure deployment.

- **Cloud Provider-Specific Settings**: The module allows specifying cloud provider-specific configurations, such as instance size and type, provider name, and disk auto-scaling options, ensuring that deployments are optimized for the chosen cloud environment.

- **Pulumi Integration**: Leverages Pulumi to manage cloud infrastructure as code. The Pulumi outputs, such as resource IDs, endpoints, and other metadata, are captured in the `status.stackOutputs`, making it easier to track and use in subsequent processes.

## Module Inputs

### MongoDB Atlas Resource Specification

The MongoDB Atlas resource specification allows the user to define the desired state of a MongoDB Atlas cluster. The following are key configurable fields:

- **`mongodb_atlas_credential_id`**: A reference to the MongoDB Atlas credential used to authenticate the provisioning of the cluster.
  
- **`cluster_config`**: The configuration details for the MongoDB Atlas cluster, including:
  - `project_id`: The unique ID of the MongoDB project.
  - `cluster_type`: The type of cluster (Replica Set, Sharded, Geo-sharded).
  - `electable_nodes`: Number of electable nodes in the region.
  - `priority`: The election priority of the region.
  - `read_only_nodes`: Number of read-only nodes.
  - `cloud_backup`: Boolean value to enable or disable cloud backups.
  - `auto_scaling_disk_gb_enabled`: Boolean value to enable or disable auto-scaling of disk storage.
  - `mongo_db_major_version`: Version of MongoDB to deploy.
  - `provider_name`: The cloud service provider (AWS, GCP, Azure).
  - `provider_instance_size_name`: The instance size for the MongoDB Atlas cluster.

### MongoDB Atlas Stack Inputs

The `MongodbAtlasStackInput` is the key input for the module. It includes:
- **`pulumi`**: Pulumi configuration specific to the stack, ensuring that the correct Pulumi resources are used.
- **`target`**: The target MongoDB Atlas resource to be deployed.
- **`mongodb_atlas_credential`**: Credential information required for authenticating with MongoDB Atlas.

## Module Outputs

The module returns a set of outputs that provide key information about the MongoDB Atlas cluster once it has been deployed. These outputs are critical for further integration and management tasks.

- **`id`**: The unique ID assigned to the MongoDB Atlas cluster by the provider.
- **`bootstrap_endpoint`**: The bootstrap endpoint used by clients to connect to the MongoDB cluster.
- **`crn`**: The Cloud Resource Name (CRN) of the MongoDB Atlas cluster, detailing the organization, environment, and cluster details.
- **`rest_endpoint`**: The REST endpoint of the MongoDB Atlas cluster for management operations.

## Usage

Refer to the example section for usage instructions.

## Notes

This module is continuously evolving, and future versions may add more features or support for additional configurations. Please ensure that the API resource specifications are correctly populated for the desired MongoDB Atlas configuration.

## Future Work

Future updates to this module will include:
- Extended support for additional cloud providers.
- More granular control over cluster scaling and security settings.
- Enhanced documentation and examples to illustrate common usage patterns.
