// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/rediskubernetes/v1/spec.proto

package rediskubernetesv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	kubernetes "github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	_ "github.com/project-planton/project-planton/apis/project/planton/shared/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// **RedisKubernetesSpec** defines the configuration for deploying Redis on a Kubernetes cluster.
// This message specifies the parameters needed to create and manage a Redis deployment within a Kubernetes environment.
// It includes container specifications and ingress settings to control resource allocation and external access.
type RedisKubernetesSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The container specifications for the Redis deployment.
	Container *RedisKubernetesContainer `protobuf:"bytes,1,opt,name=container,proto3" json:"container,omitempty"`
	// The ingress configuration for the Redis deployment.
	Ingress *kubernetes.IngressSpec `protobuf:"bytes,2,opt,name=ingress,proto3" json:"ingress,omitempty"`
}

func (x *RedisKubernetesSpec) Reset() {
	*x = RedisKubernetesSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisKubernetesSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisKubernetesSpec) ProtoMessage() {}

func (x *RedisKubernetesSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisKubernetesSpec.ProtoReflect.Descriptor instead.
func (*RedisKubernetesSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *RedisKubernetesSpec) GetContainer() *RedisKubernetesContainer {
	if x != nil {
		return x.Container
	}
	return nil
}

func (x *RedisKubernetesSpec) GetIngress() *kubernetes.IngressSpec {
	if x != nil {
		return x.Ingress
	}
	return nil
}

type RedisKubernetesContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of Redis pods to deploy.
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// The CPU and memory resources allocated to the Redis container.
	Resources *kubernetes.ContainerResources `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	// *
	// A flag to enable or disable data persistence for Redis.
	// When enabled, in-memory data is persisted to a storage volume, allowing data to survive pod restarts.
	// The backup data from the persistent volume is restored into Redis memory between pod restarts.
	// Defaults to `false`.
	IsPersistenceEnabled bool   `protobuf:"varint,3,opt,name=is_persistence_enabled,json=isPersistenceEnabled,proto3" json:"is_persistence_enabled,omitempty"`
	DiskSize             string `protobuf:"bytes,4,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
}

func (x *RedisKubernetesContainer) Reset() {
	*x = RedisKubernetesContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisKubernetesContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisKubernetesContainer) ProtoMessage() {}

func (x *RedisKubernetesContainer) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisKubernetesContainer.ProtoReflect.Descriptor instead.
func (*RedisKubernetesContainer) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescGZIP(), []int{1}
}

func (x *RedisKubernetesContainer) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *RedisKubernetesContainer) GetResources() *kubernetes.ContainerResources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *RedisKubernetesContainer) GetIsPersistenceEnabled() bool {
	if x != nil {
		return x.IsPersistenceEnabled
	}
	return false
}

func (x *RedisKubernetesContainer) GetDiskSize() string {
	if x != nil {
		return x.DiskSize
	}
	return ""
}

var file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*RedisKubernetesContainer)(nil),
		Field:         537001,
		Name:          "project.planton.provider.kubernetes.rediskubernetes.v1.default_container",
		Tag:           "bytes,537001,opt,name=default_container",
		Filename:      "project/planton/provider/kubernetes/rediskubernetes/v1/spec.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesContainer default_container = 537001;
	E_DefaultContainer = &file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_extTypes[0]
)

var File_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto protoreflect.FileDescriptor

var file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x41, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x36, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfe, 0x01,
	0x0a, 0x13, 0x52, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x9c, 0x01, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x72, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x2c, 0xca, 0x9a, 0x86,
	0x02, 0x27, 0x08, 0x01, 0x12, 0x1c, 0x0a, 0x0c, 0x0a, 0x05, 0x31, 0x30, 0x30, 0x30, 0x6d, 0x12,
	0x03, 0x31, 0x47, 0x69, 0x12, 0x0c, 0x0a, 0x03, 0x35, 0x30, 0x6d, 0x12, 0x05, 0x31, 0x30, 0x30,
	0x4d, 0x69, 0x18, 0x01, 0x22, 0x03, 0x31, 0x47, 0x69, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x22, 0xbe,
	0x04, 0x0a, 0x18, 0x52, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x73, 0x12, 0x53, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x73, 0x68, 0x61,
	0x72, 0x65, 0x64, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x52, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x16,
	0x69, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x69, 0x73,
	0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x3a,
	0xdd, 0x02, 0xba, 0x48, 0xd9, 0x02, 0x1a, 0xd6, 0x02, 0x0a, 0x21, 0x73, 0x70, 0x65, 0x63, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73,
	0x69, 0x7a, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x49, 0x44, 0x69,
	0x73, 0x6b, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x20, 0x69, 0x73, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x64, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x20, 0x74, 0x68, 0x65, 0x20, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x20, 0x69, 0x66,
	0x20, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x20, 0x69, 0x73, 0x20,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x1a, 0xe5, 0x01, 0x28, 0x28, 0x21, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x69, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65,
	0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x20, 0x26, 0x26, 0x20, 0x28, 0x73, 0x69, 0x7a,
	0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x29, 0x20, 0x3d, 0x3d, 0x20, 0x30, 0x20, 0x7c, 0x7c, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x64,
	0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x29, 0x29,
	0x20, 0x7c, 0x7c, 0x20, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x69, 0x73, 0x5f, 0x70, 0x65, 0x72,
	0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x20, 0x26, 0x26, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x64, 0x69,
	0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x29, 0x20, 0x3e, 0x20, 0x30, 0x20, 0x26, 0x26, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x2e, 0x6d,
	0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x5c, 0x5c, 0x64, 0x2b, 0x28, 0x5c, 0x5c,
	0x2e, 0x5c, 0x5c, 0x64, 0x2b, 0x29, 0x3f, 0x5c, 0x5c, 0x73, 0x3f, 0x28, 0x4b, 0x69, 0x7c, 0x4d,
	0x69, 0x7c, 0x47, 0x69, 0x7c, 0x54, 0x69, 0x7c, 0x50, 0x69, 0x7c, 0x45, 0x69, 0x7c, 0x4b, 0x7c,
	0x4d, 0x7c, 0x47, 0x7c, 0x54, 0x7c, 0x50, 0x7c, 0x45, 0x29, 0x24, 0x27, 0x29, 0x29, 0x29, 0x3a,
	0x9e, 0x01, 0x0a, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xa9, 0xe3, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x72, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x4b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x10,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x42, 0xc0, 0x03, 0x0a, 0x3a, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x72, 0x65, 0x64,
	0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x78, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x4b, 0x52, 0xaa, 0x02,
	0x36, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x36, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x42, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x52, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x3b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a,
	0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x3a, 0x3a, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a,
	0x52, 0x65, 0x64, 0x69, 0x73, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescData = file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDesc
)

func file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescData)
	})
	return file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDescData
}

var file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_goTypes = []any{
	(*RedisKubernetesSpec)(nil),           // 0: project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesSpec
	(*RedisKubernetesContainer)(nil),      // 1: project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesContainer
	(*kubernetes.IngressSpec)(nil),        // 2: project.planton.shared.kubernetes.IngressSpec
	(*kubernetes.ContainerResources)(nil), // 3: project.planton.shared.kubernetes.ContainerResources
	(*descriptorpb.FieldOptions)(nil),     // 4: google.protobuf.FieldOptions
}
var file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_depIdxs = []int32{
	1, // 0: project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesSpec.container:type_name -> project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesContainer
	2, // 1: project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesSpec.ingress:type_name -> project.planton.shared.kubernetes.IngressSpec
	3, // 2: project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesContainer.resources:type_name -> project.planton.shared.kubernetes.ContainerResources
	4, // 3: project.planton.provider.kubernetes.rediskubernetes.v1.default_container:extendee -> google.protobuf.FieldOptions
	1, // 4: project.planton.provider.kubernetes.rediskubernetes.v1.default_container:type_name -> project.planton.provider.kubernetes.rediskubernetes.v1.RedisKubernetesContainer
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	4, // [4:5] is the sub-list for extension type_name
	3, // [3:4] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_init() }
func file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_init() {
	if File_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*RedisKubernetesSpec); i {
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
		file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RedisKubernetesContainer); i {
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
			RawDescriptor: file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_msgTypes,
		ExtensionInfos:    file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_extTypes,
	}.Build()
	File_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto = out.File
	file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_rawDesc = nil
	file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_goTypes = nil
	file_project_planton_provider_kubernetes_rediskubernetes_v1_spec_proto_depIdxs = nil
}
