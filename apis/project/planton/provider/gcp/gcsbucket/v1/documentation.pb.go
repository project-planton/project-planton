// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcsbucket/v1/documentation.proto

//# Overview
//
//The GCP GCS Bucket API resource provides a consistent and streamlined interface for creating and managing Google Cloud Storage (GCS) buckets within our cloud infrastructure. By abstracting the complexities of GCS bucket configurations, this resource allows you to define your storage requirements effortlessly while ensuring consistency and compliance across different environments.
//
//## Why We Created This API Resource
//
//Managing GCS buckets directly can be complex due to various configuration options, permission settings, and best practices that need to be considered. To simplify this process and promote a standardized approach, we developed this API resource. It enables you to:
//
//- **Simplify Bucket Management**: Easily create and configure GCS buckets without dealing with low-level GCP configurations.
//- **Ensure Consistency**: Maintain uniform GCS bucket configurations across different environments and projects.
//- **Enhance Security**: Control public access settings to ensure buckets are not unintentionally exposed.
//- **Improve Productivity**: Reduce the time and effort required to manage storage resources, allowing you to focus on application development.
//
//## Key Features
//
//### Environment Integration
//
//- **Environment Info**: Seamlessly integrates with our environment management system to deploy GCS buckets within specific environments.
//- **Stack Job Settings**: Supports custom stack job settings for infrastructure-as-code deployments.
//
//### GCP Credential Management
//
//- **GCP Credential ID**: Utilizes specified GCP credentials to ensure secure and authorized operations within Google Cloud Platform.
//
//### Customizable Bucket Specifications
//
//- **Project ID**: Define the GCP project (`gcp_project_id`) where the storage bucket will be created, ensuring resources are organized within the correct project.
//- **Region Specification**: Specify the GCP region (`gcp_region`) where the GCS bucket will be created. Choosing the appropriate region can optimize data access performance and comply with regional regulations.
//- **Public Access Control**: The `is_public` flag allows you to specify whether the GCS bucket should have external (public) access. By default, this is set to `false` to enhance security.
//
//## Security and Compliance
//
//- **Access Control**: By managing the `is_public` setting, you can ensure that your buckets are only accessible as intended, preventing accidental exposure of sensitive data.
//- **Compliance with Policies**: Standardized creation of buckets helps maintain compliance with organizational and regulatory policies regarding data storage and access.
//
//## Benefits
//
//- **Simplified Deployment**: Abstracts the complexities of GCS bucket configurations into an easy-to-use API.
//- **Consistency**: Ensures all GCS buckets adhere to organizational standards for security and access control.
//- **Scalability**: Allows for efficient management of storage resources as your application and data storage needs grow.
//- **Security**: Provides control over public accessibility, reducing the risk of unauthorized data access.
//- **Flexibility**: Customize bucket settings to meet specific application requirements without compromising best practices.
//- **Cost Efficiency**: Optimize resource allocation by specifying the appropriate GCP region for your storage needs.

package gcsbucketv1

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

var File_project_planton_provider_gcp_gcsbucket_v1_documentation_proto protoreflect.FileDescriptor

var file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_rawDesc = []byte{
	0x0a, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67,
	0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x29, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63,
	0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0xf5, 0x02, 0x0a, 0x2d, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e,
	0x67, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x44, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x65, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70,
	0x2f, 0x67, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x63,
	0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x47,
	0x47, 0xaa, 0x02, 0x29, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x63, 0x70,
	0x2e, 0x47, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x29,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x73,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x35, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a,
	0x47, 0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63, 0x73, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_goTypes = []any{}
var file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_init() }
func file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_init() {
	if File_project_planton_provider_gcp_gcsbucket_v1_documentation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_depIdxs,
	}.Build()
	File_project_planton_provider_gcp_gcsbucket_v1_documentation_proto = out.File
	file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_rawDesc = nil
	file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_goTypes = nil
	file_project_planton_provider_gcp_gcsbucket_v1_documentation_proto_depIdxs = nil
}