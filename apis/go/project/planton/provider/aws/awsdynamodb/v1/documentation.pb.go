// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/aws/awsdynamodb/v1/documentation.proto

//# Overview
//
//The AWS DynamoDB API resource provides a consistent and streamlined interface for deploying and managing DynamoDB tables within our cloud infrastructure. By abstracting the complexities of AWS DynamoDB configurations, this resource allows you to define your database requirements effortlessly while ensuring consistency and compliance across different environments.
//
//## Why We Created This API Resource
//
//Deploying DynamoDB tables can be intricate due to the numerous configuration options and best practices that need to be considered. To simplify this process and promote a standardized approach, we developed this API resource. It enables you to:
//
//- **Easily Configure Tables**: Define table names, keys, billing modes, and other essential settings without dealing with low-level AWS details.
//- **Maintain Consistency**: Ensure that all DynamoDB tables across various environments adhere to our organization's standards and policies.
//- **Enhance Productivity**: Reduce the time and effort required to deploy and manage DynamoDB tables, allowing you to focus on application development.
//
//## Key Features
//
//### Environment Integration
//
//- **Environment Info**: Integrates with our environment management system to deploy tables within specific environments seamlessly.
//- **Stack Job Settings**: Supports custom stack job settings for infrastructure-as-code deployments.
//
//### AWS Credential Management
//
//- **AWS Credential ID**: Utilizes specified AWS credentials to ensure secure and authorized deployments.
//
//### Customizable Table Specifications
//
//- **Table Name**: Option to provide a custom table name or generate one based on the context.
//- **Billing Mode**: Choose between `PROVISIONED` and `PAY_PER_REQUEST` billing modes.
//- **Key Attributes**: Define the hash (partition) key and optional range (sort) key, including their data types.
//
//### Advanced Configurations
//
//- **Streams**: Enable DynamoDB Streams and specify the `stream_view_type` to control what data is captured.
//- **Server-Side Encryption**: Configure encryption at rest using AWS-managed keys or custom KMS keys.
//- **Point-in-Time Recovery**: Enable point-in-time recovery for data protection against accidental writes or deletes.
//- **Time to Live (TTL)**: Set up TTL configurations to automatically remove expired items from the table.
//- **Auto Scaling**: Define auto-scaling policies for read and write capacities based on target utilization metrics.
//
//### Indexes Support
//
//- **Global Secondary Indexes (GSIs)**: Create GSIs with specific key schemas, projection types, and provisioned capacities.
//- **Local Secondary Indexes (LSIs)**: Define LSIs at table creation for additional query flexibility using alternative sort keys.
//
//### Data Import Capability
//
//- **S3 Data Import**: Import data from Amazon S3 into a new table, supporting various input formats (`CSV`, `DYNAMODB_JSON`, `ION`) and compression types (`GZIP`, `ZSTD`, `NONE`).
//- **Input Format Options**: Customize CSV import options, including delimiters and header definitions.
//
//### Global Tables Configuration
//
//- **Replication**: Configure DynamoDB Global Tables V2 for multi-region replication by specifying replica regions.
//
//## Benefits
//
//- **Simplified Deployment**: Reduce the complexity involved in setting up DynamoDB tables with a user-friendly API.
//- **Consistency**: Ensure all tables comply with organizational standards for security, performance, and scalability.
//- **Scalability**: Leverage auto-scaling features to handle varying workloads efficiently.
//- **Security**: Integrate with AWS KMS for encryption and manage credentials securely.
//- **Flexibility**: Customize tables extensively to meet specific application requirements without compromising on best practices.

package awsdynamodbv1

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

var File_project_planton_provider_aws_awsdynamodb_v1_documentation_proto protoreflect.FileDescriptor

var file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61,
	0x77, 0x73, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e,
	0x61, 0x77, 0x73, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x86,
	0x03, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x2e,
	0x76, 0x31, 0x42, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x64, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x77, 0x73, 0x64, 0x79, 0x6e, 0x61,
	0x6d, 0x6f, 0x64, 0x62, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x41, 0x41, 0xaa, 0x02,
	0x2b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x77, 0x73, 0x2e, 0x41, 0x77,
	0x73, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x2b, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x41, 0x77, 0x73, 0x64,
	0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x37, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x41, 0x77, 0x73, 0x64, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x64, 0x62, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x30, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x64, 0x79, 0x6e, 0x61, 0x6d,
	0x6f, 0x64, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_goTypes = []any{}
var file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_init() }
func file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_init() {
	if File_project_planton_provider_aws_awsdynamodb_v1_documentation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_depIdxs,
	}.Build()
	File_project_planton_provider_aws_awsdynamodb_v1_documentation_proto = out.File
	file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_rawDesc = nil
	file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_goTypes = nil
	file_project_planton_provider_aws_awsdynamodb_v1_documentation_proto_depIdxs = nil
}