// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/aws/route53zone/v1/documentation.proto

//# Overview
//
//The AWS Route 53 Zone API resource provides a consistent and streamlined interface for creating and managing DNS zones and records within Amazon Route 53, AWS's scalable DNS web service. By abstracting the complexities of DNS management, this resource allows you to define your DNS configurations effortlessly while ensuring consistency and compliance across different environments.
//
//## Why We Created This API Resource
//
//Managing DNS zones and records can be complex due to the intricacies of DNS configurations, record types, and best practices. To simplify this process and promote a standardized approach, we developed this API resource. It enables you to:
//
//- **Simplify DNS Management**: Easily create and manage DNS zones and records without dealing with low-level AWS Route 53 configurations.
//- **Ensure Consistency**: Maintain uniform DNS configurations across different environments and applications.
//- **Enhance Productivity**: Reduce the time and effort required to manage DNS settings, allowing you to focus on application development and deployment.
//
//## Key Features
//
//### Environment Integration
//
//- **Environment Info**: Integrates seamlessly with our environment management system to deploy DNS configurations within specific environments.
//- **Stack Job Settings**: Supports custom stack job settings for infrastructure-as-code deployments.
//
//### AWS Credential Management
//
//- **AWS Credential ID**: Utilizes specified AWS credentials to ensure secure and authorized operations within AWS Route 53.
//
//### Simplified DNS Zone and Record Management
//
//- **DNS Zone Creation**: Automatically creates Route 53 DNS zones based on specified domain names.
//- **DNS Record Management**: Define DNS records within the zone, specifying record types, names, values, and TTLs.
//- **Record Types Supported**: Supports various DNS record types as defined in the `DnsRecordType` enum, such as `A`, `AAAA`, `CNAME`, `MX`, `TXT`, etc.
//- **Record Names**: Specify the fully qualified domain name (FQDN) for each record. The name should end with a dot (e.g., `example.com.`).
//- **Record Values**: Provide the values for each DNS record. For `CNAME` records, the value should also end with a dot.
//- **TTL Configuration**: Set the Time-To-Live (TTL) for each DNS record in seconds, controlling how long the record is cached by DNS resolvers.
//
//### Validation and Compliance
//
//- **Input Validation**: Implements validation rules to ensure that DNS names and record values conform to DNS standards.
//- **DNS Name Validation**: Ensures that the domain names provided are valid DNS domain names using regular expressions.
//- **Required Fields**: Enforces the presence of essential fields like `record_type` and `name`.
//
//## Benefits
//
//- **Simplified Deployment**: Abstracts the complexities of AWS Route 53 configurations into an easy-to-use API.
//- **Consistency**: Ensures all DNS zones and records adhere to organizational standards and best practices.
//- **Scalability**: Allows for efficient management of DNS settings as your application and infrastructure grow.
//- **Security**: Manages DNS configurations securely using specified AWS credentials, reducing the risk of unauthorized changes.
//- **Flexibility**: Customize DNS records extensively to meet specific application requirements without compromising on best practices.
//- **Compliance**: Helps maintain compliance with DNS standards and organizational policies through input validation and enforced field requirements.

package route53zonev1

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

var File_project_planton_provider_aws_route53zone_v1_documentation_proto protoreflect.FileDescriptor

var file_project_planton_provider_aws_route53zone_v1_documentation_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e,
	0x72, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x86,
	0x03, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x61, 0x77, 0x73, 0x2e, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x6c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x35,
	0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33,
	0x7a, 0x6f, 0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x41, 0x52, 0xaa, 0x02,
	0x2b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x77, 0x73, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x2b, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x35, 0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x37, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35,
	0x33, 0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x30, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x35, 0x33, 0x7a,
	0x6f, 0x6e, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_project_planton_provider_aws_route53zone_v1_documentation_proto_goTypes = []any{}
var file_project_planton_provider_aws_route53zone_v1_documentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_route53zone_v1_documentation_proto_init() }
func file_project_planton_provider_aws_route53zone_v1_documentation_proto_init() {
	if File_project_planton_provider_aws_route53zone_v1_documentation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_aws_route53zone_v1_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_route53zone_v1_documentation_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_route53zone_v1_documentation_proto_depIdxs,
	}.Build()
	File_project_planton_provider_aws_route53zone_v1_documentation_proto = out.File
	file_project_planton_provider_aws_route53zone_v1_documentation_proto_rawDesc = nil
	file_project_planton_provider_aws_route53zone_v1_documentation_proto_goTypes = nil
	file_project_planton_provider_aws_route53zone_v1_documentation_proto_depIdxs = nil
}