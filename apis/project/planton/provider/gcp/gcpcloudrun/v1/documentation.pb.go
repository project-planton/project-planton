// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcpcloudrun/v1/documentation.proto

//# Overview
//
//The **GCP Cloud Run API Resource** provides a consistent and standardized interface for deploying and managing applications on Google Cloud Run within our infrastructure. This resource simplifies the process of running containerized applications in a fully managed serverless environment on Google Cloud Platform (GCP), allowing users to build and deploy scalable web applications and APIs without managing servers.
//
//## Purpose
//
//We developed this API resource to streamline the deployment and management of containerized applications using GCP Cloud Run. By offering a unified interface, it reduces the complexity involved in setting up and configuring serverless containers, enabling users to:
//
//- **Easily Deploy Cloud Run Services**: Quickly create and deploy services in specified GCP projects.
//- **Simplify Configuration**: Abstract the complexities of setting up GCP Cloud Run, including environment settings and permissions.
//- **Integrate Seamlessly**: Utilize existing GCP credentials and integrate with other GCP services.
//- **Focus on Code**: Allow developers to concentrate on writing code rather than managing infrastructure.
//
//## Key Features
//
//- **Consistent Interface**: Aligns with our existing APIs for deploying open-source software, microservices, and cloud infrastructure.
//- **Simplified Deployment**: Automates the provisioning of Cloud Run services, including setting up necessary permissions and environment variables.
//- **Scalability**: Leverages GCP's serverless infrastructure to automatically scale applications based on demand.
//- **Flexible Configuration**: Supports specifying GCP projects and credentials for seamless integration.
//- **Integration**: Works seamlessly with other GCP services like Cloud Storage, Cloud SQL, and Firestore.
//
//## Use Cases
//
//- **Web Application Deployment**: Deploy containerized web applications without worrying about server management.
//- **API Hosting**: Host scalable APIs and microservices with automatic scaling and high availability.
//- **Event-Driven Applications**: Build applications that respond to events from Pub/Sub, Cloud Storage, or other services.
//- **Background Processing**: Run background tasks and asynchronous processing in a serverless environment.
//- **Continuous Deployment**: Integrate with CI/CD pipelines for automated deployments and updates.
//
//## Future Enhancements
//
//As this resource is currently in a partial implementation phase, future updates will include:
//
//- **Advanced Configuration Options**: Support for custom domain mappings, traffic splitting, and concurrency settings.
//- **Enhanced Security Features**: Integration with VPC connectors, IAM roles, and secret management.
//- **Monitoring and Logging**: Improved support for logging, tracing, and monitoring using Google Cloud Logging and Monitoring.
//- **Automation and CI/CD Integration**: Streamlined deployment processes with integration into continuous deployment pipelines.
//- **Comprehensive Documentation**: Expanded usage examples, best practices, and troubleshooting guides.

package gcpcloudrunv1

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

var File_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto protoreflect.FileDescriptor

var file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_rawDesc = []byte{
	0x0a, 0x3f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67,
	0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f,
	0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x2b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e,
	0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x83,
	0x03, 0x0a, 0x2f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75, 0x6e, 0x2e,
	0x76, 0x31, 0x42, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x69, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72,
	0x75, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75,
	0x6e, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x47, 0x47, 0xaa, 0x02, 0x2b, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x63, 0x70, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x72, 0x75, 0x6e, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x2b, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x72, 0x75, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x37, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72,
	0x75, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x30, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a,
	0x47, 0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75, 0x6e,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_goTypes = []any{}
var file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_init() }
func file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_init() {
	if File_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_depIdxs,
	}.Build()
	File_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto = out.File
	file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_rawDesc = nil
	file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_goTypes = nil
	file_project_planton_provider_gcp_gcpcloudrun_v1_documentation_proto_depIdxs = nil
}