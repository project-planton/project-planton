// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/azure/azurekeyvault/v1/documentation.proto

//# Overview
//
//The **Azure Key Vault API Resource** provides a consistent and standardized interface for deploying and managing secrets using Azure Key Vault within our infrastructure. This resource simplifies the process of storing, retrieving, and managing secrets and cryptographic keys, allowing users to handle sensitive information securely and efficiently.
//
//## Purpose
//
//We developed this API resource to streamline the management of secrets and keys across various applications and services in Azure. By offering a unified interface, it reduces the complexity involved in handling credentials and sensitive data, enabling users to:
//
//- **Create and Manage Secrets**: Effortlessly create and store secrets in Azure Key Vault.
//- **Integrate Seamlessly**: Incorporate secret management into existing workflows and deployments.
//- **Enhance Security**: Centralize secret storage with robust encryption and access control.
//- **Manage Cryptographic Keys**: Handle keys for encryption, decryption, and signing operations.
//- **Focus on Development**: Allow developers to concentrate on application logic without worrying about secret distribution.
//
//## Key Features
//
//- **Consistent Interface**: Aligns with our existing APIs for deploying open-source software, microservices, and cloud infrastructure.
//- **Simplified Configuration**: Abstracts the complexities of Azure Key Vault, enabling quicker setups without deep Azure expertise.
//- **Scalability**: Manage multiple secrets across different environments and applications efficiently.
//- **Security**: Leverages Azure's encryption and Azure Active Directory (AAD) for access control to protect sensitive data.
//- **Integration**: Works seamlessly with other Azure services and can be integrated into CI/CD pipelines.
//
//## Use Cases
//
//- **Credential Management**: Securely store database passwords, API keys, certificates, and other application credentials.
//- **Key Management**: Generate and manage cryptographic keys for encryption and signing.
//- **Automatic Secret Rotation**: Implement policies for regular secret rotation to enhance security and compliance.
//- **Multi-Environment Deployments**: Manage secrets for development, staging, and production environments separately.
//- **Compliance and Auditing**: Meet organizational and regulatory requirements for secret management and auditing.
//
//## Future Enhancements
//
//As this resource is currently in a partial implementation phase, future updates will include:
//
//- **Advanced Secret Features**: Support for certificate management, key versioning, and secret tagging.
//- **Enhanced Access Control**: Fine-grained permissions using Azure RBAC and integration with AAD.
//- **Monitoring and Auditing**: Integration with Azure Monitor and Azure Security Center for tracking secret access and changes.
//- **Automation**: Support for automatic secret rotation and integration with Azure DevOps.
//- **Comprehensive Documentation**: Expanded usage examples, best practices, and troubleshooting guides.

package azurekeyvaultv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_project_planton_provider_azure_azurekeyvault_v1_documentation_proto protoreflect.FileDescriptor

var file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_rawDesc = []byte{
	0x0a, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65,
	0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x9d, 0x03, 0x0a, 0x33, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x12,
	0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61,
	0x7a, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61,
	0x75, 0x6c, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x41, 0x41, 0xaa, 0x02, 0x2f,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x2e, 0x41,
	0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x7a, 0x75, 0x72, 0x65,
	0x5c, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x3b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x5c, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x75, 0x6c, 0x74,
	0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x34, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x41, 0x7a,
	0x75, 0x72, 0x65, 0x3a, 0x3a, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x6b, 0x65, 0x79, 0x76, 0x61, 0x75,
	0x6c, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_goTypes = []any{}
var file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_init() }
func file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_init() {
	if File_project_planton_provider_azure_azurekeyvault_v1_documentation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_depIdxs,
	}.Build()
	File_project_planton_provider_azure_azurekeyvault_v1_documentation_proto = out.File
	file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_rawDesc = nil
	file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_goTypes = nil
	file_project_planton_provider_azure_azurekeyvault_v1_documentation_proto_depIdxs = nil
}