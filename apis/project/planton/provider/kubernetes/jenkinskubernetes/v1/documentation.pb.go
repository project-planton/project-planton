// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/jenkinskubernetes/v1/documentation.proto

//# Overview
//
//The **Jenkins Kubernetes API Resource** provides a standardized and efficient way to deploy Jenkins onto Kubernetes clusters. This API resource simplifies the deployment process by encapsulating all necessary configurations, allowing for consistent and repeatable Jenkins deployments across different environments.
//
//## Purpose
//
//Deploying Jenkins on Kubernetes involves complex configurations, including container resources, environment settings, and customization via Helm charts. The Jenkins Kubernetes API Resource aims to:
//
//- **Standardize Deployments**: Offer a consistent interface for deploying Jenkins, reducing complexity and potential errors.
//- **Simplify Configuration Management**: Centralize all deployment settings, making it easier to manage, update, and replicate configurations.
//- **Enhance Customization**: Allow granular customization through Helm values to meet specific organizational needs.
//
//## Key Features
//
//### Environment Configuration
//
//- **Environment Info**: Tailor Jenkins deployments to specific environments (development, staging, production) using environment-specific information.
//- **Stack Job Settings**: Integrate with infrastructure-as-code (IaC) tools through stack job settings for automated and repeatable deployments.
//
//### Credential Management
//
//  - **Kubernetes Cluster Credential ID**: Specify credentials required to access and configure the target Kubernetes cluster securely.
//
//### Container Specification
//
//- **Jenkins Container Resources**: Define CPU and memory resources for the Jenkins container to optimize performance and resource utilization. Recommended defaults are:
//- **CPU Requests**: 50m
//- **Memory Requests**: 256Mi
//- **CPU Limits**: 1
//- **Memory Limits**: 1Gi
//
//### Helm Chart Customization
//
//- **Helm Values**: Provide a map of key-value pairs for additional customization options via the Jenkins Helm chart. This allows for:
//- Customizing resource limits
//- Setting environment variables
//- Specifying version tags
//- For detailed options, refer to the [Jenkins Helm Chart values.yaml](https://github.com/jenkinsci/helm-charts/blob/main/charts/jenkins/values.yaml)
//
//### Networking and Ingress
//
//- **Ingress Configuration**: Set up Kubernetes Ingress resources to manage external access to Jenkins, including hostname and path routing.
//
//## Benefits
//
//    - **Consistency Across Deployments**: By using a standardized API resource, Jenkins deployments become more predictable and maintainable.
//    - **Reduced Complexity**: Simplifies the deployment process by abstracting complex Kubernetes and Helm configurations.
//    - **Scalability**: Allows for resource adjustments to meet performance requirements.
//    - **Customization**: Enables detailed customization through Helm values to fit specific use cases.
//
//## Use Cases
//
//- **Automated CI/CD Pipelines**: Deploy Jenkins as part of a continuous integration and deployment pipeline, automating the setup of the CI server.
//- **Multi-Environment Deployments**: Consistently deploy Jenkins across different environments with environment-specific configurations.
//- **Resource Optimization**: Adjust resource allocations for Jenkins to optimize performance and cost based on usage patterns.
//- **Custom Jenkins Configurations**: Utilize Helm values to customize Jenkins installations, including plugins, security settings, and more.

package jenkinskubernetesv1

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

var File_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto protoreflect.FileDescriptor

var file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_rawDesc = []byte{
	0x0a, 0x4c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x38,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0xd7, 0x03, 0x0a, 0x3c, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x44, 0x6f, 0x63, 0x75, 0x6d,
	0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x7c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2f, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6a, 0x65, 0x6e, 0x6b, 0x69, 0x6e,
	0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x05,
	0x50, 0x50, 0x50, 0x4b, 0x4a, 0xaa, 0x02, 0x38, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4a, 0x65, 0x6e, 0x6b,
	0x69, 0x6e, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x38, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x44, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5c, 0x4a, 0x65, 0x6e, 0x6b, 0x69, 0x6e, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x3d, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a,
	0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x4a, 0x65, 0x6e,
	0x6b, 0x69, 0x6e, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_goTypes = []any{}
var file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_init() }
func file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_init() {
	if File_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_depIdxs,
	}.Build()
	File_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto = out.File
	file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_rawDesc = nil
	file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_goTypes = nil
	file_project_planton_provider_kubernetes_jenkinskubernetes_v1_documentation_proto_depIdxs = nil
}