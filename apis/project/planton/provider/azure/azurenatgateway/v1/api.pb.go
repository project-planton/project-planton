// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/azure/azurenatgateway/v1/api.proto

package azurenatgatewayv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	shared "github.com/project-planton/project-planton/apis/project/planton/shared"
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

// azure-nat-gateway
type AzureNatGateway struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	Metadata *shared.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *AzureNatGatewaySpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status        *AzureNatGatewayStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AzureNatGateway) Reset() {
	*x = AzureNatGateway{}
	mi := &file_project_planton_provider_azure_azurenatgateway_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureNatGateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureNatGateway) ProtoMessage() {}

func (x *AzureNatGateway) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_azure_azurenatgateway_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureNatGateway.ProtoReflect.Descriptor instead.
func (*AzureNatGateway) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *AzureNatGateway) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AzureNatGateway) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AzureNatGateway) GetMetadata() *shared.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AzureNatGateway) GetSpec() *AzureNatGatewaySpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *AzureNatGateway) GetStatus() *AzureNatGatewayStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// azure-nat-gateway status
type AzureNatGatewayStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lifecycle
	Lifecycle *shared.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *shared.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack-job id
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	//
	//	azure-nat-gateway stack-outputs
	Outputs       *AzureNatGatewayStackOutputs `protobuf:"bytes,1,opt,name=outputs,proto3" json:"outputs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AzureNatGatewayStatus) Reset() {
	*x = AzureNatGatewayStatus{}
	mi := &file_project_planton_provider_azure_azurenatgateway_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureNatGatewayStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureNatGatewayStatus) ProtoMessage() {}

func (x *AzureNatGatewayStatus) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_azure_azurenatgateway_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureNatGatewayStatus.ProtoReflect.Descriptor instead.
func (*AzureNatGatewayStatus) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *AzureNatGatewayStatus) GetLifecycle() *shared.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *AzureNatGatewayStatus) GetAudit() *shared.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *AzureNatGatewayStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *AzureNatGatewayStatus) GetOutputs() *AzureNatGatewayStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_project_planton_provider_azure_azurenatgateway_v1_api_proto protoreflect.FileDescriptor

const file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDesc = "" +
	"\n" +
	";project/planton/provider/azure/azurenatgateway/v1/api.proto\x121project.planton.provider.azure.azurenatgateway.v1\x1a\x1bbuf/validate/validate.proto\x1a<project/planton/provider/azure/azurenatgateway/v1/spec.proto\x1aEproject/planton/provider/azure/azurenatgateway/v1/stack_outputs.proto\x1a#project/planton/shared/status.proto\x1a%project/planton/shared/metadata.proto\"\x9a\x03\n" +
	"\x0fAzureNatGateway\x12D\n" +
	"\vapi_version\x18\x01 \x01(\tB#\xbaH r\x1e\n" +
	"\x1cazure.project-planton.org/v1R\n" +
	"apiVersion\x12*\n" +
	"\x04kind\x18\x02 \x01(\tB\x16\xbaH\x13r\x11\n" +
	"\x0fAzureNatGatewayR\x04kind\x12O\n" +
	"\bmetadata\x18\x03 \x01(\v2+.project.planton.shared.ApiResourceMetadataB\x06\xbaH\x03\xc8\x01\x01R\bmetadata\x12b\n" +
	"\x04spec\x18\x04 \x01(\v2F.project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewaySpecB\x06\xbaH\x03\xc8\x01\x01R\x04spec\x12`\n" +
	"\x06status\x18\x05 \x01(\v2H.project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStatusR\x06status\"\xaf\x02\n" +
	"\x15AzureNatGatewayStatus\x12J\n" +
	"\tlifecycle\x18c \x01(\v2,.project.planton.shared.ApiResourceLifecycleR\tlifecycle\x12>\n" +
	"\x05audit\x18b \x01(\v2(.project.planton.shared.ApiResourceAuditR\x05audit\x12 \n" +
	"\fstack_job_id\x18a \x01(\tR\n" +
	"stackJobId\x12h\n" +
	"\aoutputs\x18\x01 \x01(\v2N.project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStackOutputsR\aoutputsB\xa1\x03\n" +
	"5com.project.planton.provider.azure.azurenatgateway.v1B\bApiProtoP\x01Zsgithub.com/project-planton/project-planton/apis/project/planton/provider/azure/azurenatgateway/v1;azurenatgatewayv1\xa2\x02\x05PPPAA\xaa\x021Project.Planton.Provider.Azure.Azurenatgateway.V1\xca\x021Project\\Planton\\Provider\\Azure\\Azurenatgateway\\V1\xe2\x02=Project\\Planton\\Provider\\Azure\\Azurenatgateway\\V1\\GPBMetadata\xea\x026Project::Planton::Provider::Azure::Azurenatgateway::V1b\x06proto3"

var (
	file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescOnce sync.Once
	file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescData []byte
)

func file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescGZIP() []byte {
	file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDesc), len(file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDesc)))
	})
	return file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDescData
}

var file_project_planton_provider_azure_azurenatgateway_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_azure_azurenatgateway_v1_api_proto_goTypes = []any{
	(*AzureNatGateway)(nil),             // 0: project.planton.provider.azure.azurenatgateway.v1.AzureNatGateway
	(*AzureNatGatewayStatus)(nil),       // 1: project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStatus
	(*shared.ApiResourceMetadata)(nil),  // 2: project.planton.shared.ApiResourceMetadata
	(*AzureNatGatewaySpec)(nil),         // 3: project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewaySpec
	(*shared.ApiResourceLifecycle)(nil), // 4: project.planton.shared.ApiResourceLifecycle
	(*shared.ApiResourceAudit)(nil),     // 5: project.planton.shared.ApiResourceAudit
	(*AzureNatGatewayStackOutputs)(nil), // 6: project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStackOutputs
}
var file_project_planton_provider_azure_azurenatgateway_v1_api_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.azure.azurenatgateway.v1.AzureNatGateway.metadata:type_name -> project.planton.shared.ApiResourceMetadata
	3, // 1: project.planton.provider.azure.azurenatgateway.v1.AzureNatGateway.spec:type_name -> project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewaySpec
	1, // 2: project.planton.provider.azure.azurenatgateway.v1.AzureNatGateway.status:type_name -> project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStatus
	4, // 3: project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStatus.lifecycle:type_name -> project.planton.shared.ApiResourceLifecycle
	5, // 4: project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStatus.audit:type_name -> project.planton.shared.ApiResourceAudit
	6, // 5: project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStatus.outputs:type_name -> project.planton.provider.azure.azurenatgateway.v1.AzureNatGatewayStackOutputs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_project_planton_provider_azure_azurenatgateway_v1_api_proto_init() }
func file_project_planton_provider_azure_azurenatgateway_v1_api_proto_init() {
	if File_project_planton_provider_azure_azurenatgateway_v1_api_proto != nil {
		return
	}
	file_project_planton_provider_azure_azurenatgateway_v1_spec_proto_init()
	file_project_planton_provider_azure_azurenatgateway_v1_stack_outputs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDesc), len(file_project_planton_provider_azure_azurenatgateway_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_azure_azurenatgateway_v1_api_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_azure_azurenatgateway_v1_api_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_azure_azurenatgateway_v1_api_proto_msgTypes,
	}.Build()
	File_project_planton_provider_azure_azurenatgateway_v1_api_proto = out.File
	file_project_planton_provider_azure_azurenatgateway_v1_api_proto_goTypes = nil
	file_project_planton_provider_azure_azurenatgateway_v1_api_proto_depIdxs = nil
}
