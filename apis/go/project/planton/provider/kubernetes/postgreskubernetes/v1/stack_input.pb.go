// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/postgreskubernetes/v1/stack_input.proto

package postgreskubernetesv1

import (
	v1 "github.com/project-planton/project-planton/apis/go/project/planton/credential/kubernetesclustercredential/v1"
	pulumi "github.com/project-planton/project-planton/apis/go/project/planton/shared/pulumi"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// postgres-kubernetes stack-input
type PostgresKubernetesStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pulumi input
	Pulumi *pulumi.PulumiStackInfo `protobuf:"bytes,1,opt,name=pulumi,proto3" json:"pulumi,omitempty"`
	// target api-resource
	Target *PostgresKubernetes `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// kubernetes-cluster-credential
	KubernetesCluster *v1.KubernetesClusterCredentialSpec `protobuf:"bytes,3,opt,name=kubernetes_cluster,json=kubernetesCluster,proto3" json:"kubernetes_cluster,omitempty"`
}

func (x *PostgresKubernetesStackInput) Reset() {
	*x = PostgresKubernetesStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostgresKubernetesStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostgresKubernetesStackInput) ProtoMessage() {}

func (x *PostgresKubernetesStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostgresKubernetesStackInput.ProtoReflect.Descriptor instead.
func (*PostgresKubernetesStackInput) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescGZIP(), []int{0}
}

func (x *PostgresKubernetesStackInput) GetPulumi() *pulumi.PulumiStackInfo {
	if x != nil {
		return x.Pulumi
	}
	return nil
}

func (x *PostgresKubernetesStackInput) GetTarget() *PostgresKubernetes {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *PostgresKubernetesStackInput) GetKubernetesCluster() *v1.KubernetesClusterCredentialSpec {
	if x != nil {
		return x.KubernetesCluster
	}
	return nil
}

var File_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto protoreflect.FileDescriptor

var file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDesc = []byte{
	0x0a, 0x4b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63,
	0x6b, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x39, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x43, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x44, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd9, 0x02, 0x0a, 0x1c, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x46, 0x0a, 0x06, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69,
	0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x06, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x12, 0x65, 0x0a, 0x06, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x4d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12,
	0x89, 0x01, 0x0a, 0x12, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x5a, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x53, 0x70, 0x65, 0x63, 0x52, 0x11, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x42, 0xdf, 0x03, 0x0a, 0x3d,
	0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x81, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72,
	0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x4b, 0x50, 0xaa, 0x02, 0x39, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x39, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x50,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x45, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x50, 0x6f, 0x73, 0x74, 0x67,
	0x72, 0x65, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3e, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a,
	0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x50, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescOnce sync.Once
	file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescData = file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDesc
)

func file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescGZIP() []byte {
	file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescData)
	})
	return file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDescData
}

var file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_goTypes = []any{
	(*PostgresKubernetesStackInput)(nil),       // 0: project.planton.provider.kubernetes.postgreskubernetes.v1.PostgresKubernetesStackInput
	(*pulumi.PulumiStackInfo)(nil),             // 1: project.planton.shared.pulumi.PulumiStackInfo
	(*PostgresKubernetes)(nil),                 // 2: project.planton.provider.kubernetes.postgreskubernetes.v1.PostgresKubernetes
	(*v1.KubernetesClusterCredentialSpec)(nil), // 3: project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredentialSpec
}
var file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_depIdxs = []int32{
	1, // 0: project.planton.provider.kubernetes.postgreskubernetes.v1.PostgresKubernetesStackInput.pulumi:type_name -> project.planton.shared.pulumi.PulumiStackInfo
	2, // 1: project.planton.provider.kubernetes.postgreskubernetes.v1.PostgresKubernetesStackInput.target:type_name -> project.planton.provider.kubernetes.postgreskubernetes.v1.PostgresKubernetes
	3, // 2: project.planton.provider.kubernetes.postgreskubernetes.v1.PostgresKubernetesStackInput.kubernetes_cluster:type_name -> project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredentialSpec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_init() }
func file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_init() {
	if File_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto != nil {
		return
	}
	file_project_planton_provider_kubernetes_postgreskubernetes_v1_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*PostgresKubernetesStackInput); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_msgTypes,
	}.Build()
	File_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto = out.File
	file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_rawDesc = nil
	file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_goTypes = nil
	file_project_planton_provider_kubernetes_postgreskubernetes_v1_stack_input_proto_depIdxs = nil
}