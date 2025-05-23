// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/snowflake/snowflakedatabase/v1/stack_outputs.proto

package snowflakedatabasev1

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

// snowflake-database stack outputs
// https://www.pulumi.com/registry/packages/snowflakecloud/api-docs/kafkacluster/#outputs
type SnowflakeDatabaseStackOutputs struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The provider-assigned unique ID for this managed resource.
	// https://www.pulumi.com/registry/packages/snowflakecloud/api-docs/kafkacluster/#id_yaml
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The bootstrap endpoint used by Kafka clients to connect to the Kafka cluster. (e.g., SASL_SSL://pkc-00000.us-central1.gcp.snowflake.cloud:9092).
	// https://www.pulumi.com/registry/packages/snowflakecloud/api-docs/kafkacluster/#bootstrapendpoint_yaml
	BootstrapEndpoint string `protobuf:"bytes,2,opt,name=bootstrap_endpoint,json=bootstrapEndpoint,proto3" json:"bootstrap_endpoint,omitempty"`
	// The Snowflake Resource Name of the Kafka cluster,
	// for example, crn://snowflake.cloud/organization=1111aaaa-11aa-11aa-11aa-111111aaaaaa/environment=env-abc123/cloud-cluster=lkc-abc123.
	// https://www.pulumi.com/registry/packages/snowflakecloud/api-docs/kafkacluster/#rbaccrn_yaml
	Crn string `protobuf:"bytes,3,opt,name=crn,proto3" json:"crn,omitempty"`
	// The REST endpoint of the Kafka cluster (e.g., https://pkc-00000.us-central1.gcp.snowflake.cloud:443).
	// https://www.pulumi.com/registry/packages/snowflakecloud/api-docs/kafkacluster/#restendpoint_yaml
	RestEndpoint  string `protobuf:"bytes,4,opt,name=rest_endpoint,json=restEndpoint,proto3" json:"rest_endpoint,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SnowflakeDatabaseStackOutputs) Reset() {
	*x = SnowflakeDatabaseStackOutputs{}
	mi := &file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SnowflakeDatabaseStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnowflakeDatabaseStackOutputs) ProtoMessage() {}

func (x *SnowflakeDatabaseStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnowflakeDatabaseStackOutputs.ProtoReflect.Descriptor instead.
func (*SnowflakeDatabaseStackOutputs) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDescGZIP(), []int{0}
}

func (x *SnowflakeDatabaseStackOutputs) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SnowflakeDatabaseStackOutputs) GetBootstrapEndpoint() string {
	if x != nil {
		return x.BootstrapEndpoint
	}
	return ""
}

func (x *SnowflakeDatabaseStackOutputs) GetCrn() string {
	if x != nil {
		return x.Crn
	}
	return ""
}

func (x *SnowflakeDatabaseStackOutputs) GetRestEndpoint() string {
	if x != nil {
		return x.RestEndpoint
	}
	return ""
}

var File_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto protoreflect.FileDescriptor

const file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDesc = "" +
	"\n" +
	"Kproject/planton/provider/snowflake/snowflakedatabase/v1/stack_outputs.proto\x127project.planton.provider.snowflake.snowflakedatabase.v1\"\x95\x01\n" +
	"\x1dSnowflakeDatabaseStackOutputs\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\tR\x02id\x12-\n" +
	"\x12bootstrap_endpoint\x18\x02 \x01(\tR\x11bootstrapEndpoint\x12\x10\n" +
	"\x03crn\x18\x03 \x01(\tR\x03crn\x12#\n" +
	"\rrest_endpoint\x18\x04 \x01(\tR\frestEndpointB\xd0\x03\n" +
	";com.project.planton.provider.snowflake.snowflakedatabase.v1B\x11StackOutputsProtoP\x01Z{github.com/project-planton/project-planton/apis/project/planton/provider/snowflake/snowflakedatabase/v1;snowflakedatabasev1\xa2\x02\x05PPPSS\xaa\x027Project.Planton.Provider.Snowflake.Snowflakedatabase.V1\xca\x027Project\\Planton\\Provider\\Snowflake\\Snowflakedatabase\\V1\xe2\x02CProject\\Planton\\Provider\\Snowflake\\Snowflakedatabase\\V1\\GPBMetadata\xea\x02<Project::Planton::Provider::Snowflake::Snowflakedatabase::V1b\x06proto3"

var (
	file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDescOnce sync.Once
	file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDescData []byte
)

func file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDescGZIP() []byte {
	file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDesc)))
	})
	return file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDescData
}

var file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_goTypes = []any{
	(*SnowflakeDatabaseStackOutputs)(nil), // 0: project.planton.provider.snowflake.snowflakedatabase.v1.SnowflakeDatabaseStackOutputs
}
var file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_init() }
func file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_init() {
	if File_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_msgTypes,
	}.Build()
	File_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto = out.File
	file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_goTypes = nil
	file_project_planton_provider_snowflake_snowflakedatabase_v1_stack_outputs_proto_depIdxs = nil
}
