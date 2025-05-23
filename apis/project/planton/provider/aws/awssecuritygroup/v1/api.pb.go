// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/aws/awssecuritygroup/v1/api.proto

package awssecuritygroupv1

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

// aws-security-group
type AwsSecurityGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	Metadata *shared.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *AwsSecurityGroupSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status        *AwsSecurityGroupStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsSecurityGroup) Reset() {
	*x = AwsSecurityGroup{}
	mi := &file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsSecurityGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsSecurityGroup) ProtoMessage() {}

func (x *AwsSecurityGroup) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsSecurityGroup.ProtoReflect.Descriptor instead.
func (*AwsSecurityGroup) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *AwsSecurityGroup) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AwsSecurityGroup) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AwsSecurityGroup) GetMetadata() *shared.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AwsSecurityGroup) GetSpec() *AwsSecurityGroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *AwsSecurityGroup) GetStatus() *AwsSecurityGroupStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// aws-security-group status
type AwsSecurityGroupStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lifecycle
	Lifecycle *shared.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *shared.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack-job id
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	Outputs       *AwsSecurityGroupStackOutputs `protobuf:"bytes,1,opt,name=outputs,proto3" json:"outputs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsSecurityGroupStatus) Reset() {
	*x = AwsSecurityGroupStatus{}
	mi := &file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsSecurityGroupStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsSecurityGroupStatus) ProtoMessage() {}

func (x *AwsSecurityGroupStatus) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsSecurityGroupStatus.ProtoReflect.Descriptor instead.
func (*AwsSecurityGroupStatus) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *AwsSecurityGroupStatus) GetLifecycle() *shared.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *AwsSecurityGroupStatus) GetAudit() *shared.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *AwsSecurityGroupStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *AwsSecurityGroupStatus) GetOutputs() *AwsSecurityGroupStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_project_planton_provider_aws_awssecuritygroup_v1_api_proto protoreflect.FileDescriptor

const file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDesc = "" +
	"\n" +
	":project/planton/provider/aws/awssecuritygroup/v1/api.proto\x120project.planton.provider.aws.awssecuritygroup.v1\x1a\x1bbuf/validate/validate.proto\x1a;project/planton/provider/aws/awssecuritygroup/v1/spec.proto\x1aDproject/planton/provider/aws/awssecuritygroup/v1/stack_outputs.proto\x1a#project/planton/shared/status.proto\x1a%project/planton/shared/metadata.proto\"\x9a\x03\n" +
	"\x10AwsSecurityGroup\x12B\n" +
	"\vapi_version\x18\x01 \x01(\tB!\xbaH\x1er\x1c\n" +
	"\x1aaws.project-planton.org/v1R\n" +
	"apiVersion\x12+\n" +
	"\x04kind\x18\x02 \x01(\tB\x17\xbaH\x14r\x12\n" +
	"\x10AwsSecurityGroupR\x04kind\x12O\n" +
	"\bmetadata\x18\x03 \x01(\v2+.project.planton.shared.ApiResourceMetadataB\x06\xbaH\x03\xc8\x01\x01R\bmetadata\x12b\n" +
	"\x04spec\x18\x04 \x01(\v2F.project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupSpecB\x06\xbaH\x03\xc8\x01\x01R\x04spec\x12`\n" +
	"\x06status\x18\x05 \x01(\v2H.project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStatusR\x06status\"\xb0\x02\n" +
	"\x16AwsSecurityGroupStatus\x12J\n" +
	"\tlifecycle\x18c \x01(\v2,.project.planton.shared.ApiResourceLifecycleR\tlifecycle\x12>\n" +
	"\x05audit\x18b \x01(\v2(.project.planton.shared.ApiResourceAuditR\x05audit\x12 \n" +
	"\fstack_job_id\x18a \x01(\tR\n" +
	"stackJobId\x12h\n" +
	"\aoutputs\x18\x01 \x01(\v2N.project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStackOutputsR\aoutputsB\x9c\x03\n" +
	"4com.project.planton.provider.aws.awssecuritygroup.v1B\bApiProtoP\x01Zsgithub.com/project-planton/project-planton/apis/project/planton/provider/aws/awssecuritygroup/v1;awssecuritygroupv1\xa2\x02\x05PPPAA\xaa\x020Project.Planton.Provider.Aws.Awssecuritygroup.V1\xca\x020Project\\Planton\\Provider\\Aws\\Awssecuritygroup\\V1\xe2\x02<Project\\Planton\\Provider\\Aws\\Awssecuritygroup\\V1\\GPBMetadata\xea\x025Project::Planton::Provider::Aws::Awssecuritygroup::V1b\x06proto3"

var (
	file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescData []byte
)

func file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDesc), len(file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDesc)))
	})
	return file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDescData
}

var file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_goTypes = []any{
	(*AwsSecurityGroup)(nil),             // 0: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroup
	(*AwsSecurityGroupStatus)(nil),       // 1: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStatus
	(*shared.ApiResourceMetadata)(nil),   // 2: project.planton.shared.ApiResourceMetadata
	(*AwsSecurityGroupSpec)(nil),         // 3: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupSpec
	(*shared.ApiResourceLifecycle)(nil),  // 4: project.planton.shared.ApiResourceLifecycle
	(*shared.ApiResourceAudit)(nil),      // 5: project.planton.shared.ApiResourceAudit
	(*AwsSecurityGroupStackOutputs)(nil), // 6: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStackOutputs
}
var file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroup.metadata:type_name -> project.planton.shared.ApiResourceMetadata
	3, // 1: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroup.spec:type_name -> project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupSpec
	1, // 2: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroup.status:type_name -> project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStatus
	4, // 3: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStatus.lifecycle:type_name -> project.planton.shared.ApiResourceLifecycle
	5, // 4: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStatus.audit:type_name -> project.planton.shared.ApiResourceAudit
	6, // 5: project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStatus.outputs:type_name -> project.planton.provider.aws.awssecuritygroup.v1.AwsSecurityGroupStackOutputs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_init() }
func file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_init() {
	if File_project_planton_provider_aws_awssecuritygroup_v1_api_proto != nil {
		return
	}
	file_project_planton_provider_aws_awssecuritygroup_v1_spec_proto_init()
	file_project_planton_provider_aws_awssecuritygroup_v1_stack_outputs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDesc), len(file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awssecuritygroup_v1_api_proto = out.File
	file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_goTypes = nil
	file_project_planton_provider_aws_awssecuritygroup_v1_api_proto_depIdxs = nil
}
