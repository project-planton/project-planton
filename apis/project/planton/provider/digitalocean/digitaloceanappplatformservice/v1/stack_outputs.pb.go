// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/digitalocean/digitaloceanappplatformservice/v1/stack_outputs.proto

package digitaloceanappplatformservicev1

import (
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

// DigitalOceanAppPlatformServiceStackOutputs captures the key outputs after provisioning a service on DigitalOcean App Platform.
type DigitalOceanAppPlatformServiceStackOutputs struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// app_id is the unique identifier of the app (DigitalOcean App Platform application ID).
	AppId string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	// default_hostname is the default hostname assigned to the app (usually ending in "ondigitalocean.app").
	DefaultHostname string `protobuf:"bytes,2,opt,name=default_hostname,json=defaultHostname,proto3" json:"default_hostname,omitempty"`
	// live_url is the publicly accessible URL (including protocol) of the deployed service.
	// This may be the same as the default hostname with "https://" prefix, or a custom domain if one was configured.
	LiveUrl       string `protobuf:"bytes,3,opt,name=live_url,json=liveUrl,proto3" json:"live_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DigitalOceanAppPlatformServiceStackOutputs) Reset() {
	*x = DigitalOceanAppPlatformServiceStackOutputs{}
	mi := &file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DigitalOceanAppPlatformServiceStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DigitalOceanAppPlatformServiceStackOutputs) ProtoMessage() {}

func (x *DigitalOceanAppPlatformServiceStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DigitalOceanAppPlatformServiceStackOutputs.ProtoReflect.Descriptor instead.
func (*DigitalOceanAppPlatformServiceStackOutputs) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDescGZIP(), []int{0}
}

func (x *DigitalOceanAppPlatformServiceStackOutputs) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *DigitalOceanAppPlatformServiceStackOutputs) GetDefaultHostname() string {
	if x != nil {
		return x.DefaultHostname
	}
	return ""
}

func (x *DigitalOceanAppPlatformServiceStackOutputs) GetLiveUrl() string {
	if x != nil {
		return x.LiveUrl
	}
	return ""
}

var File_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto protoreflect.FileDescriptor

const file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDesc = "" +
	"\n" +
	"[project/planton/provider/digitalocean/digitaloceanappplatformservice/v1/stack_outputs.proto\x12Gproject.planton.provider.digitalocean.digitaloceanappplatformservice.v1\"\x89\x01\n" +
	"*DigitalOceanAppPlatformServiceStackOutputs\x12\x15\n" +
	"\x06app_id\x18\x01 \x01(\tR\x05appId\x12)\n" +
	"\x10default_hostname\x18\x02 \x01(\tR\x0fdefaultHostname\x12\x19\n" +
	"\blive_url\x18\x03 \x01(\tR\aliveUrlB\xbe\x04\n" +
	"Kcom.project.planton.provider.digitalocean.digitaloceanappplatformservice.v1B\x11StackOutputsProtoP\x01Z\x98\x01github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean/digitaloceanappplatformservice/v1;digitaloceanappplatformservicev1\xa2\x02\x05PPPDD\xaa\x02GProject.Planton.Provider.Digitalocean.Digitaloceanappplatformservice.V1\xca\x02GProject\\Planton\\Provider\\Digitalocean\\Digitaloceanappplatformservice\\V1\xe2\x02SProject\\Planton\\Provider\\Digitalocean\\Digitaloceanappplatformservice\\V1\\GPBMetadata\xea\x02LProject::Planton::Provider::Digitalocean::Digitaloceanappplatformservice::V1b\x06proto3"

var (
	file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDescOnce sync.Once
	file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDescData []byte
)

func file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDescGZIP() []byte {
	file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDesc)))
	})
	return file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDescData
}

var file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_goTypes = []any{
	(*DigitalOceanAppPlatformServiceStackOutputs)(nil), // 0: project.planton.provider.digitalocean.digitaloceanappplatformservice.v1.DigitalOceanAppPlatformServiceStackOutputs
}
var file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() {
	file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_init()
}
func file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_init() {
	if File_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_msgTypes,
	}.Build()
	File_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto = out.File
	file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_goTypes = nil
	file_project_planton_provider_digitalocean_digitaloceanappplatformservice_v1_stack_outputs_proto_depIdxs = nil
}
