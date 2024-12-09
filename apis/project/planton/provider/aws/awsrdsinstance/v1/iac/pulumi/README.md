# AWS RDS Instance Pulumi Module

## Introduction

The AWS RDS Instance Pulumi module is designed to automate the provisioning, management, and configuration of Amazon RDS instances using a Kubernetes-like API resource model. This module is part of our Unified APIs, which provide a standardized approach to managing cloud resources across multiple cloud providers. The module takes an `AwsRdsInstance` API resource as input, allowing users to specify various parameters such as database engine, version, storage options, security settings, and networking configurations. Once the YAML configuration is provided, the module provisions an RDS instance in AWS, with the outputs captured in `status.stackOutputs` for easy access and integration.

## Key Features

- **Unified API Structure**: Uses Kubernetes-like resource modeling (`apiVersion`, `kind`, `metadata`, `spec`, and `status`), making it consistent with other cloud resources.
- **Pulumi Integration**: Leverages Pulumi's infrastructure-as-code capabilities to manage and automate AWS RDS instances, ensuring consistent deployments across different environments.
- **Credential Management**: Securely manages AWS credentials to authenticate and provision resources.
- **Flexible Configuration**: Supports detailed configuration of RDS instances, including database engine type, version, instance class, storage, network, and security settings.
- **Support for Replication and Backups**: Includes options for multi-AZ deployments, replication from existing RDS instances, backup windows, and automated snapshots.
- **Output Management**: Exports key outputs such as the RDS instance endpoint, ID, security groups, and parameter group configurations, making it easier to integrate the RDS instance with other services.

## Usage

Refer to the example section for usage instructions.

## Module Details

### API Resource Specification

The `AwsRdsInstance` API resource contains various fields that allow users to define the desired state of their AWS RDS instance. The key components of the `spec` section include:

- **`awsCredentialId`** (required): The identifier for the AWS credentials to authenticate with AWS.
- **Database Configuration**:  
  - `db_name`: The name of the database to create. Not applicable for Oracle or SQL Server engines.
  - `engine`, `engine_version`: Specifies the database engine and version to use (e.g., MySQL, PostgreSQL).
  - `instance_class`: Defines the type of instance (e.g., db.t3.medium).
  - `allocated_storage`, `max_allocated_storage`: Configures the storage capacity for the RDS instance.
  - `multiAz`: Enables multi-AZ deployments for high availability.

- **Security Settings**:  
  - `security_group_ids`, `associate_security_group_ids`: Specifies the security groups to associate with the instance.
  - `kms_key_id`, `storage_encrypted`: Configures encryption options using AWS KMS.
  - `subnet_ids`, `db_subnet_group_name`, `availability_zone`: Defines the networking configuration, including VPC subnets and availability zones.

- **Performance and Monitoring**:  
  - `performance_insights`: Enables AWS Performance Insights with KMS encryption.
  - `monitoring`: Configures enhanced monitoring, including interval and IAM role for CloudWatch.

- **Backup and Maintenance**:  
  - `backup_retention_period`, `backup_window`: Defines backup settings and windows.
  - `maintenance_window`: Specifies the maintenance window for the RDS instance.
  - `snapshot_identifier`: Allows the instance to be restored from an existing snapshot.

- **Replication**:  
  - `replicate_source_db`: Specifies an existing RDS instance as the source for replication, supporting both single-region and cross-region replication.

### Pulumi Module Functionality

The module executes the following tasks:

1. **AWS Provider Initialization**:  
   Sets up the AWS provider within the Pulumi context using the credentials provided in the `AwsRdsInstanceStackInput`. This provider is required to authenticate and manage AWS resources.

2. **Security Group Creation**:  
   Creates a default security group if none is provided, ensuring the RDS instance is secure and accessible based on user-defined ingress rules.

3. **RDS Instance Creation**:  
   Provisions the RDS instance in AWS based on the configuration defined in the `spec`. This includes setting up the database engine, storage, networking, and security groups.

4. **Output Management**:  
   Captures and exports essential information about the RDS instance, including:
   - `rds_instance_endpoint`: The DNS endpoint for database connectivity.
   - `rds_instance_id`, `rds_instance_arn`: Unique identifiers for the RDS instance in AWS.
   - `rds_subnet_group`, `rds_security_group`: The associated subnet group and security group details.
   - `rds_parameter_group`, `rds_options_group`: The parameter and option group configurations.

## Limitations

- **Not Completely Implemented**: The current implementation of this module is not fully complete. Certain features such as the automatic management of master user passwords through AWS Secrets Manager are not yet supported.
- **Placeholder Values for Certain Fields**: Some fields, such as the password and database name, may need to be managed externally or provided directly in the YAML configuration.
- **Limited Error Handling**: Basic error handling is in place, but more advanced error validation and user feedback will be introduced in future releases.

## Future Enhancements

- **Enhanced Security Features**: Full support for AWS Secrets Manager to manage master user passwords.
- **Comprehensive Error Handling**: Improve error reporting and validation for better user feedback.
- **Support for Additional AWS RDS Features**: Extend support for database clustering (e.g., Aurora) and cross-region replication.
- **Advanced Output Handling**: Capture more detailed outputs, such as database performance metrics, to assist with monitoring and optimization.

## Documentation

For detailed API definitions and additional documentation, please refer to the resources available via [buf.build](https://buf.build).

## Contributing

Contributions are welcome! Please open issues or pull requests to help improve this module.

## License

This project is licensed under the [MIT License](LICENSE).