// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcprouternat/v1/stack_outputs.proto

package gcprouternatv1

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

// Outputs produced after provisioning a GCP Cloud Router and NAT.
type GcpRouterNatStackOutputs struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Name of the Cloud NAT gateway (as created in GCP).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Self-link URL of the Cloud Router that was created (or used) for this NAT.
	RouterSelfLink string `protobuf:"bytes,2,opt,name=router_self_link,json=routerSelfLink,proto3" json:"router_self_link,omitempty"`
	// List of external IP addresses utilized by this NAT (e.g. auto-allocated or static IPs provided).
	NatIpAddresses []string `protobuf:"bytes,3,rep,name=nat_ip_addresses,json=natIpAddresses,proto3" json:"nat_ip_addresses,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GcpRouterNatStackOutputs) Reset() {
	*x = GcpRouterNatStackOutputs{}
	mi := &file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GcpRouterNatStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpRouterNatStackOutputs) ProtoMessage() {}

func (x *GcpRouterNatStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpRouterNatStackOutputs.ProtoReflect.Descriptor instead.
func (*GcpRouterNatStackOutputs) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDescGZIP(), []int{0}
}

func (x *GcpRouterNatStackOutputs) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GcpRouterNatStackOutputs) GetRouterSelfLink() string {
	if x != nil {
		return x.RouterSelfLink
	}
	return ""
}

func (x *GcpRouterNatStackOutputs) GetNatIpAddresses() []string {
	if x != nil {
		return x.NatIpAddresses
	}
	return nil
}

var File_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto protoreflect.FileDescriptor

const file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDesc = "" +
	"\n" +
	"@project/planton/provider/gcp/gcprouternat/v1/stack_outputs.proto\x12,project.planton.provider.gcp.gcprouternat.v1\"\x82\x01\n" +
	"\x18GcpRouterNatStackOutputs\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12(\n" +
	"\x10router_self_link\x18\x02 \x01(\tR\x0erouterSelfLink\x12(\n" +
	"\x10nat_ip_addresses\x18\x03 \x03(\tR\x0enatIpAddressesB\x89\x03\n" +
	"0com.project.planton.provider.gcp.gcprouternat.v1B\x11StackOutputsProtoP\x01Zkgithub.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcprouternat/v1;gcprouternatv1\xa2\x02\x05PPPGG\xaa\x02,Project.Planton.Provider.Gcp.Gcprouternat.V1\xca\x02,Project\\Planton\\Provider\\Gcp\\Gcprouternat\\V1\xe2\x028Project\\Planton\\Provider\\Gcp\\Gcprouternat\\V1\\GPBMetadata\xea\x021Project::Planton::Provider::Gcp::Gcprouternat::V1b\x06proto3"

var (
	file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDescOnce sync.Once
	file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDescData []byte
)

func file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDescGZIP() []byte {
	file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDesc)))
	})
	return file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDescData
}

var file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_goTypes = []any{
	(*GcpRouterNatStackOutputs)(nil), // 0: project.planton.provider.gcp.gcprouternat.v1.GcpRouterNatStackOutputs
}
var file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_init() }
func file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_init() {
	if File_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_msgTypes,
	}.Build()
	File_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto = out.File
	file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_goTypes = nil
	file_project_planton_provider_gcp_gcprouternat_v1_stack_outputs_proto_depIdxs = nil
}
