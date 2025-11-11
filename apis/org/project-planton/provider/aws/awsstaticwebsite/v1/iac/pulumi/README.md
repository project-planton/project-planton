**Note:** This module is not completely implemented as the API resource specification is currently empty.

# AWS Static Website Pulumi Module

## Introduction

This Pulumi module is designed to automate the deployment of static websites on AWS using our Unified APIs that mimic Kubernetes' resource modeling. By leveraging this module, developers can define their static website configurations in a standardized YAML file, simplifying the process of hosting static content on AWS services such as S3 and CloudFront.

## Key Features

- **Unified API Structure**: Consistent resource modeling with `apiVersion`, `kind`, `metadata`, `spec`, and `status`, similar to Kubernetes.
- **Multi-Cloud Support**: Designed to work seamlessly across multiple cloud providers, starting with AWS.
- **Pulumi Integration**: Utilizes Pulumi's infrastructure-as-code capabilities to automate the provisioning and management of AWS resources required for hosting static websites.
- **Credential Management**: Securely handles AWS credentials for authenticating with AWS services.
- **Simplified Deployment**: Allows developers to deploy static websites using a single YAML configuration file.
- **Output Handling**: Captures and exports essential output parameters, such as the S3 bucket ID, in `status.outputs`.
- **Standardized Documentation**: Comprehensive documentation available via buf.build for easy reference.

## Usage

Refer to the example section for usage instructions.

## Module Details

### API Resource Specification

The module expects an `api-resource.yaml` file defining the desired state of the AWS static website. The key components of this file include:

- **`awsProviderConfigId`** (required): The identifier for the AWS credentials used to authenticate with AWS services.
- **`environmentInfo`**: Contains environment-specific information (currently not utilized).
- **`stackJobSettings`**: Settings related to the stack job execution (currently not utilized).

### Pulumi Module Functionality

The core functionality of this module revolves around setting up the AWS provider within the Pulumi context using the provided AWS credentials. This setup is essential for any resource creation and management within AWS, such as creating S3 buckets and configuring CloudFront distributions for static website hosting.

#### Steps Performed:

1. **AWS Provider Initialization**:  
   Initializes the AWS provider in Pulumi using credentials supplied in the `AwsStaticWebsiteStackInput`. The credentials required are:

   - `AccessKeyId`
   - `SecretAccessKey`
   - `Region`

2. **Resource Provisioning**:  
   *(Not yet implemented)* The module will provision AWS resources such as S3 buckets for storing static content and CloudFront distributions for content delivery based on the specifications provided in the `api-resource.yaml` file.

3. **Output Handling**:  
   *(Not yet implemented)* Captures and exports essential outputs, like the S3 bucket ID, storing them in `status.outputs` for easy reference by other resources or applications.

## Limitations

- **Incomplete Implementation**: The module currently does not implement resource creation due to the empty API resource specification.
- **Unused Spec Fields**: Fields like `environmentInfo` and `stackJobSettings` are included in the spec but are not utilized in the current implementation.
- **No Error Handling**: Advanced error handling and validation mechanisms are yet to be implemented.

## Future Enhancements

- **Implement Resource Creation**: Extend the module to create S3 buckets, set up static website hosting configurations, and optionally configure CloudFront distributions.
- **Utilize Spec Fields**: Make use of `environmentInfo` and `stackJobSettings` to allow for more granular control over the deployment environment and stack job configurations.
- **Enhance Output Management**: Capture and expose essential output parameters such as bucket URLs, website endpoints, and CloudFront distribution details.
- **Error Handling and Validation**: Introduce comprehensive error handling and input validation to improve reliability and user experience.

## Documentation

For detailed API definitions and additional documentation, please refer to our resources available via [buf.build](https://buf.build).

## Contributing

Contributions are welcome! Please open issues or pull requests to help improve this module.

## License

This project is licensed under the [MIT License](LICENSE).