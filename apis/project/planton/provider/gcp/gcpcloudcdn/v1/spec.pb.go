// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcpcloudcdn/v1/spec.proto

package gcpcloudcdnv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// **GcpCloudCdnSpec** defines the configuration for deploying a Google Cloud CDN (Content Delivery Network).
// This message specifies the necessary parameters to create and manage a Cloud CDN within a
// specified GCP project. By providing the project ID, you can set up CDN resources to accelerate
// content delivery by caching content at edge locations globally, improving load times and
// reducing latency for end-users.
type GcpCloudCdnSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ID of the GCP project where the Cloud CDN resources will be created.
	GcpProjectId  string `protobuf:"bytes,1,opt,name=gcp_project_id,json=gcpProjectId,proto3" json:"gcp_project_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GcpCloudCdnSpec) Reset() {
	*x = GcpCloudCdnSpec{}
	mi := &file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GcpCloudCdnSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpCloudCdnSpec) ProtoMessage() {}

func (x *GcpCloudCdnSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpCloudCdnSpec.ProtoReflect.Descriptor instead.
func (*GcpCloudCdnSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *GcpCloudCdnSpec) GetGcpProjectId() string {
	if x != nil {
		return x.GcpProjectId
	}
	return ""
}

var File_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto protoreflect.FileDescriptor

const file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDesc = "" +
	"\n" +
	"6project/planton/provider/gcp/gcpcloudcdn/v1/spec.proto\x12+project.planton.provider.gcp.gcpcloudcdn.v1\x1a\x1bbuf/validate/validate.proto\"?\n" +
	"\x0fGcpCloudCdnSpec\x12,\n" +
	"\x0egcp_project_id\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\fgcpProjectIdB\xfa\x02\n" +
	"/com.project.planton.provider.gcp.gcpcloudcdn.v1B\tSpecProtoP\x01Zigithub.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpcloudcdn/v1;gcpcloudcdnv1\xa2\x02\x05PPPGG\xaa\x02+Project.Planton.Provider.Gcp.Gcpcloudcdn.V1\xca\x02+Project\\Planton\\Provider\\Gcp\\Gcpcloudcdn\\V1\xe2\x027Project\\Planton\\Provider\\Gcp\\Gcpcloudcdn\\V1\\GPBMetadata\xea\x020Project::Planton::Provider::Gcp::Gcpcloudcdn::V1b\x06proto3"

var (
	file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDescData []byte
)

func file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDesc), len(file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDesc)))
	})
	return file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDescData
}

var file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_goTypes = []any{
	(*GcpCloudCdnSpec)(nil), // 0: project.planton.provider.gcp.gcpcloudcdn.v1.GcpCloudCdnSpec
}
var file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_init() }
func file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_init() {
	if File_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDesc), len(file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_msgTypes,
	}.Build()
	File_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto = out.File
	file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_goTypes = nil
	file_project_planton_provider_gcp_gcpcloudcdn_v1_spec_proto_depIdxs = nil
}
