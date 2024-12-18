// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/solrkubernetes/v1/documentation.proto

//# Overview
//
//The **Solr Kubernetes API resource** provides a structured way to deploy and manage Solr clusters in Kubernetes environments. It includes configurations for Solr, Zookeeper, and ingress management, allowing for a comprehensive setup of a Solr cluster that is optimized for scalability and performance.
//
//## Purpose of the Solr Kubernetes API Resource
//
//Deploying Solr in Kubernetes often involves various components, such as managing Solr configurations, Zookeeper instances, resource allocations, and persistence settings. This API resource simplifies that process by offering a well-defined structure for Solr deployment in Kubernetes, ensuring high availability, efficient resource management, and easy scaling.
//
//## Key Features
//
//### Environment and Stack Integration
//
//### Solr Container Configuration
//
//- **Replicas**: Configure the number of Solr pod replicas, with a recommended default of 1 for initial deployments.
//- **Container Image**: Define the Solr container image, such as `solr:8.7.0`, for deployment.
//- **Resource Allocation**: Solr container resources can be customized. The recommended default values are:
//- **CPU Requests**: `50m`
//- **Memory Requests**: `256Mi`
//- **CPU Limits**: `1`
//- **Memory Limits**: `1Gi`
//- **Disk Size**: Allocate disk storage for persistent data. The default is `1Gi`, ensuring persistent data backup in case of restarts.
//
//### Solr Configuration
//
//- **JVM Memory Settings**: Set JVM memory configurations for Solr. The default is `"-Xms1g -Xmx3g"`.
//- **Custom Solr Options**: Provide additional Solr options, such as `-Dsolr.autoSoftCommit.maxTime=10000`, to tune Solr performance.
//- **Garbage Collection Tuning**: Customize the garbage collection settings for Solr, such as `-XX:SurvivorRatio=4 -XX:TargetSurvivorRatio=90`.
//
//### Zookeeper Container Configuration
//
//- **Replicas**: Configure the number of Zookeeper pod replicas, with a recommended default of 1.
//- **Resource Allocation**: Customize Zookeeper's container resources. The recommended default values are:
//- **CPU Requests**: `50m`
//- **Memory Requests**: `256Mi`
//- **CPU Limits**: `1`
//- **Memory Limits**: `1Gi`
//- **Disk Size**: Allocate disk storage for Zookeeper with a default value of `1Gi`.
//
//### Ingress Configuration
//
//- **Ingress Spec**: Use Kubernetes ingress configurations to expose the Solr service securely, enabling external access as needed.
//
//## Benefits
//
//- **Simplified Deployment**: This API resource abstracts the complexities of deploying and managing Solr in Kubernetes, offering a straightforward approach.
//- **Scalable and Resilient**: Built-in configuration options for replicas, resource management, and persistence ensure a highly available and scalable Solr cluster.
//- **Data Persistence**: Persistent storage options guarantee that Solr data is securely backed up, reducing the risk of data loss during restarts or failures.
//- **Customizable**: Fine-tune resource allocations, JVM settings, and garbage collection configurations to match your performance requirements.
//- **Integrated Zookeeper**: Manage Zookeeper instances alongside Solr with similar configuration options for ease of use.

package solrkubernetesv1

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

var File_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto protoreflect.FileDescriptor

var file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_rawDesc = []byte{
	0x0a, 0x49, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x35, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x73, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0xc2, 0x03, 0x0a, 0x39, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x73,
	0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x76, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x73, 0x6f, 0x6c, 0x72,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x6f,
	0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x05, 0x50, 0x50, 0x50, 0x4b, 0x53, 0xaa, 0x02, 0x35, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x6c,
	0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x35, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x5c, 0x53, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x41, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x53, 0x6f, 0x6c,
	0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3a, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x3a, 0x3a, 0x53, 0x6f, 0x6c, 0x72, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_goTypes = []any{}
var file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_init() }
func file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_init() {
	if File_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_depIdxs,
	}.Build()
	File_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto = out.File
	file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_rawDesc = nil
	file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_goTypes = nil
	file_project_planton_provider_kubernetes_solrkubernetes_v1_documentation_proto_depIdxs = nil
}
