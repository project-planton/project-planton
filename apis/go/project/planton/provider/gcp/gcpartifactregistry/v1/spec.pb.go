// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcpartifactregistry/v1/spec.proto

package gcpartifactregistryv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// **GcpArtifactRegistrySpec** defines the configuration for deploying a Google Cloud Artifact Registry.
// This message specifies the necessary parameters to create and manage an Artifact Registry within a
// specified GCP project and region. It allows you to set the project ID and region, and configure
// access settings such as enabling unauthenticated external access, which is particularly useful for
// open-source projects that require public availability of their artifacts.
type GcpArtifactRegistrySpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// **Required.** The ID of the GCP project where the Artifact Registry resources will be created.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// **Required.** The GCP region where the Artifact Registry will be created (e.g., "us-west2").
	// Selecting a region close to your Kubernetes clusters can reduce service startup time
	// by enabling faster downloads of container images.
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// A flag indicating whether to allow unauthenticated access to artifacts published in the repositories.
	// Enable this for publishing artifacts for open-source projects that require public access.
	IsExternal bool `protobuf:"varint,3,opt,name=is_external,json=isExternal,proto3" json:"is_external,omitempty"`
}

func (x *GcpArtifactRegistrySpec) Reset() {
	*x = GcpArtifactRegistrySpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpArtifactRegistrySpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpArtifactRegistrySpec) ProtoMessage() {}

func (x *GcpArtifactRegistrySpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpArtifactRegistrySpec.ProtoReflect.Descriptor instead.
func (*GcpArtifactRegistrySpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *GcpArtifactRegistrySpec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GcpArtifactRegistrySpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GcpArtifactRegistrySpec) GetIsExternal() bool {
	if x != nil {
		return x.IsExternal
	}
	return false
}

var File_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto protoreflect.FileDescriptor

var file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x3e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x33, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x17, 0x47, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x45, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0xb5, 0x03, 0x0a, 0x37, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x7c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70,
	0x2f, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x05,
	0x50, 0x50, 0x50, 0x47, 0x47, 0xaa, 0x02, 0x33, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x33, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x3f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70,
	0x5c, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x38, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescData = file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDesc
)

func file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescData)
	})
	return file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDescData
}

var file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_goTypes = []any{
	(*GcpArtifactRegistrySpec)(nil), // 0: project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistrySpec
}
var file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_init() }
func file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_init() {
	if File_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GcpArtifactRegistrySpec); i {
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
			RawDescriptor: file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_msgTypes,
	}.Build()
	File_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto = out.File
	file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_rawDesc = nil
	file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_goTypes = nil
	file_project_planton_provider_gcp_gcpartifactregistry_v1_spec_proto_depIdxs = nil
}