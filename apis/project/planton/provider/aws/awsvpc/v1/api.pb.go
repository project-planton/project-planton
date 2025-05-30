// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/aws/awsvpc/v1/api.proto

package awsvpcv1

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

// aws-vpc
type AwsVpc struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	Metadata *shared.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *AwsVpcSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status        *AwsVpcStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsVpc) Reset() {
	*x = AwsVpc{}
	mi := &file_project_planton_provider_aws_awsvpc_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsVpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsVpc) ProtoMessage() {}

func (x *AwsVpc) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awsvpc_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsVpc.ProtoReflect.Descriptor instead.
func (*AwsVpc) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *AwsVpc) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AwsVpc) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AwsVpc) GetMetadata() *shared.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AwsVpc) GetSpec() *AwsVpcSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *AwsVpc) GetStatus() *AwsVpcStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// aws-vpc status
type AwsVpcStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lifecycle
	Lifecycle *shared.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *shared.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack-job id
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	Outputs       *AwsVpcStackOutputs `protobuf:"bytes,1,opt,name=outputs,proto3" json:"outputs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsVpcStatus) Reset() {
	*x = AwsVpcStatus{}
	mi := &file_project_planton_provider_aws_awsvpc_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsVpcStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsVpcStatus) ProtoMessage() {}

func (x *AwsVpcStatus) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awsvpc_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsVpcStatus.ProtoReflect.Descriptor instead.
func (*AwsVpcStatus) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *AwsVpcStatus) GetLifecycle() *shared.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *AwsVpcStatus) GetAudit() *shared.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *AwsVpcStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *AwsVpcStatus) GetOutputs() *AwsVpcStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_project_planton_provider_aws_awsvpc_v1_api_proto protoreflect.FileDescriptor

const file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDesc = "" +
	"\n" +
	"0project/planton/provider/aws/awsvpc/v1/api.proto\x12&project.planton.provider.aws.awsvpc.v1\x1a\x1bbuf/validate/validate.proto\x1a1project/planton/provider/aws/awsvpc/v1/spec.proto\x1a:project/planton/provider/aws/awsvpc/v1/stack_outputs.proto\x1a#project/planton/shared/status.proto\x1a%project/planton/shared/metadata.proto\"\xde\x02\n" +
	"\x06AwsVpc\x12B\n" +
	"\vapi_version\x18\x01 \x01(\tB!\xbaH\x1er\x1c\n" +
	"\x1aaws.project-planton.org/v1R\n" +
	"apiVersion\x12!\n" +
	"\x04kind\x18\x02 \x01(\tB\r\xbaH\n" +
	"r\b\n" +
	"\x06AwsVpcR\x04kind\x12O\n" +
	"\bmetadata\x18\x03 \x01(\v2+.project.planton.shared.ApiResourceMetadataB\x06\xbaH\x03\xc8\x01\x01R\bmetadata\x12N\n" +
	"\x04spec\x18\x04 \x01(\v22.project.planton.provider.aws.awsvpc.v1.AwsVpcSpecB\x06\xbaH\x03\xc8\x01\x01R\x04spec\x12L\n" +
	"\x06status\x18\x05 \x01(\v24.project.planton.provider.aws.awsvpc.v1.AwsVpcStatusR\x06status\"\x92\x02\n" +
	"\fAwsVpcStatus\x12J\n" +
	"\tlifecycle\x18c \x01(\v2,.project.planton.shared.ApiResourceLifecycleR\tlifecycle\x12>\n" +
	"\x05audit\x18b \x01(\v2(.project.planton.shared.ApiResourceAuditR\x05audit\x12 \n" +
	"\fstack_job_id\x18a \x01(\tR\n" +
	"stackJobId\x12T\n" +
	"\aoutputs\x18\x01 \x01(\v2:.project.planton.provider.aws.awsvpc.v1.AwsVpcStackOutputsR\aoutputsB\xd6\x02\n" +
	"*com.project.planton.provider.aws.awsvpc.v1B\bApiProtoP\x01Z_github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsvpc/v1;awsvpcv1\xa2\x02\x05PPPAA\xaa\x02&Project.Planton.Provider.Aws.Awsvpc.V1\xca\x02&Project\\Planton\\Provider\\Aws\\Awsvpc\\V1\xe2\x022Project\\Planton\\Provider\\Aws\\Awsvpc\\V1\\GPBMetadata\xea\x02+Project::Planton::Provider::Aws::Awsvpc::V1b\x06proto3"

var (
	file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescData []byte
)

func file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDesc), len(file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDesc)))
	})
	return file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDescData
}

var file_project_planton_provider_aws_awsvpc_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_aws_awsvpc_v1_api_proto_goTypes = []any{
	(*AwsVpc)(nil),                      // 0: project.planton.provider.aws.awsvpc.v1.AwsVpc
	(*AwsVpcStatus)(nil),                // 1: project.planton.provider.aws.awsvpc.v1.AwsVpcStatus
	(*shared.ApiResourceMetadata)(nil),  // 2: project.planton.shared.ApiResourceMetadata
	(*AwsVpcSpec)(nil),                  // 3: project.planton.provider.aws.awsvpc.v1.AwsVpcSpec
	(*shared.ApiResourceLifecycle)(nil), // 4: project.planton.shared.ApiResourceLifecycle
	(*shared.ApiResourceAudit)(nil),     // 5: project.planton.shared.ApiResourceAudit
	(*AwsVpcStackOutputs)(nil),          // 6: project.planton.provider.aws.awsvpc.v1.AwsVpcStackOutputs
}
var file_project_planton_provider_aws_awsvpc_v1_api_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.aws.awsvpc.v1.AwsVpc.metadata:type_name -> project.planton.shared.ApiResourceMetadata
	3, // 1: project.planton.provider.aws.awsvpc.v1.AwsVpc.spec:type_name -> project.planton.provider.aws.awsvpc.v1.AwsVpcSpec
	1, // 2: project.planton.provider.aws.awsvpc.v1.AwsVpc.status:type_name -> project.planton.provider.aws.awsvpc.v1.AwsVpcStatus
	4, // 3: project.planton.provider.aws.awsvpc.v1.AwsVpcStatus.lifecycle:type_name -> project.planton.shared.ApiResourceLifecycle
	5, // 4: project.planton.provider.aws.awsvpc.v1.AwsVpcStatus.audit:type_name -> project.planton.shared.ApiResourceAudit
	6, // 5: project.planton.provider.aws.awsvpc.v1.AwsVpcStatus.outputs:type_name -> project.planton.provider.aws.awsvpc.v1.AwsVpcStackOutputs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awsvpc_v1_api_proto_init() }
func file_project_planton_provider_aws_awsvpc_v1_api_proto_init() {
	if File_project_planton_provider_aws_awsvpc_v1_api_proto != nil {
		return
	}
	file_project_planton_provider_aws_awsvpc_v1_spec_proto_init()
	file_project_planton_provider_aws_awsvpc_v1_stack_outputs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDesc), len(file_project_planton_provider_aws_awsvpc_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awsvpc_v1_api_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awsvpc_v1_api_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_aws_awsvpc_v1_api_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awsvpc_v1_api_proto = out.File
	file_project_planton_provider_aws_awsvpc_v1_api_proto_goTypes = nil
	file_project_planton_provider_aws_awsvpc_v1_api_proto_depIdxs = nil
}
