// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcpartifactregistry/v1/stack_input.proto

package gcpartifactregistryv1

import (
	v1 "github.com/project-planton/project-planton/apis/project/planton/credential/gcpcredential/v1"
	pulumi "github.com/project-planton/project-planton/apis/project/planton/shared/pulumi"
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

// gcp-artifact-registry stack-input
type GcpArtifactRegistryStackInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pulumi input
	Pulumi *pulumi.PulumiStackInfo `protobuf:"bytes,1,opt,name=pulumi,proto3" json:"pulumi,omitempty"`
	// target api-resource
	Target *GcpArtifactRegistry `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	// gcp-credential
	GcpCredential *v1.GcpCredentialSpec `protobuf:"bytes,3,opt,name=gcp_credential,json=gcpCredential,proto3" json:"gcp_credential,omitempty"`
}

func (x *GcpArtifactRegistryStackInput) Reset() {
	*x = GcpArtifactRegistryStackInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpArtifactRegistryStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpArtifactRegistryStackInput) ProtoMessage() {}

func (x *GcpArtifactRegistryStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpArtifactRegistryStackInput.ProtoReflect.Descriptor instead.
func (*GcpArtifactRegistryStackInput) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescGZIP(), []int{0}
}

func (x *GcpArtifactRegistryStackInput) GetPulumi() *pulumi.PulumiStackInfo {
	if x != nil {
		return x.Pulumi
	}
	return nil
}

func (x *GcpArtifactRegistryStackInput) GetTarget() *GcpArtifactRegistry {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *GcpArtifactRegistryStackInput) GetGcpCredential() *v1.GcpCredentialSpec {
	if x != nil {
		return x.GcpCredential
	}
	return nil
}

var File_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto protoreflect.FileDescriptor

var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDesc = []byte{
	0x0a, 0x45, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x3d, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x67, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x70, 0x75, 0x6c, 0x75,
	0x6d, 0x69, 0x2f, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xb0, 0x02, 0x0a, 0x1d, 0x47, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x46, 0x0a, 0x06, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x75, 0x6c, 0x75, 0x6d,
	0x69, 0x2e, 0x50, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x70, 0x75, 0x6c, 0x75, 0x6d, 0x69, 0x12, 0x60, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x48, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x65, 0x0a, 0x0e, 0x67,
	0x63, 0x70, 0x5f, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x3e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c,
	0x2e, 0x67, 0x63, 0x70, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x63, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x0d, 0x67, 0x63, 0x70, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x42, 0xb8, 0x03, 0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0f,
	0x53, 0x74, 0x61, 0x63, 0x6b, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x79, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70,
	0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f,
	0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50,
	0x50, 0x50, 0x47, 0x47, 0xaa, 0x02, 0x33, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x47, 0x63, 0x70, 0x2e, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x33, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x3f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c,
	0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x38, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a,
	0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescOnce sync.Once
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescData = file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDesc
)

func file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescGZIP() []byte {
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescData)
	})
	return file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDescData
}

var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_goTypes = []any{
	(*GcpArtifactRegistryStackInput)(nil), // 0: project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistryStackInput
	(*pulumi.PulumiStackInfo)(nil),        // 1: project.planton.shared.pulumi.PulumiStackInfo
	(*GcpArtifactRegistry)(nil),           // 2: project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistry
	(*v1.GcpCredentialSpec)(nil),          // 3: project.planton.credential.gcpcredential.v1.GcpCredentialSpec
}
var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_depIdxs = []int32{
	1, // 0: project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistryStackInput.pulumi:type_name -> project.planton.shared.pulumi.PulumiStackInfo
	2, // 1: project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistryStackInput.target:type_name -> project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistry
	3, // 2: project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistryStackInput.gcp_credential:type_name -> project.planton.credential.gcpcredential.v1.GcpCredentialSpec
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_init() }
func file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_init() {
	if File_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto != nil {
		return
	}
	file_project_planton_provider_gcp_gcpartifactregistry_v1_api_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GcpArtifactRegistryStackInput); i {
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
			RawDescriptor: file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_msgTypes,
	}.Build()
	File_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto = out.File
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_rawDesc = nil
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_goTypes = nil
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_input_proto_depIdxs = nil
}
