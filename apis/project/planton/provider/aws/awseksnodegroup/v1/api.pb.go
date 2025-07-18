// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/aws/awseksnodegroup/v1/api.proto

package awseksnodegroupv1

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

// AwsEksNodeGroup represents a containerized application deployed on AWS ECS.
// This resource manages ECS services that can run on either Fargate or EC2.
type AwsEksNodeGroup struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// api-version must be set to "aws.project-planton.org/v1".
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind for this ECS service resource.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata captures identifying information (name, org, version, etc.)
	// and must pass standard validations for resource naming.
	Metadata *shared.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec holds the core configuration data defining how the ECS service is deployed.
	Spec *AwsEksNodeGroupSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status holds runtime or post-deployment information.
	Status        *AwsEksNodeGroupStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsEksNodeGroup) Reset() {
	*x = AwsEksNodeGroup{}
	mi := &file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsEksNodeGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsEksNodeGroup) ProtoMessage() {}

func (x *AwsEksNodeGroup) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsEksNodeGroup.ProtoReflect.Descriptor instead.
func (*AwsEksNodeGroup) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *AwsEksNodeGroup) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AwsEksNodeGroup) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AwsEksNodeGroup) GetMetadata() *shared.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *AwsEksNodeGroup) GetSpec() *AwsEksNodeGroupSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *AwsEksNodeGroup) GetStatus() *AwsEksNodeGroupStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// AwsEksNodeGroupStatus describes the status fields for an ECS service resource.
type AwsEksNodeGroupStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lifecycle indicates if the resource is active or has been marked for removal.
	Lifecycle *shared.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit contains creation and update information for the resource.
	Audit *shared.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack_job_id stores the ID of the Pulumi/Terraform stack job responsible for provisioning.
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack_outputs captures the outputs returned by Pulumi/Terraform after provisioning.
	Outputs       *AwsEksNodeGroupStackOutputs `protobuf:"bytes,1,opt,name=outputs,proto3" json:"outputs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsEksNodeGroupStatus) Reset() {
	*x = AwsEksNodeGroupStatus{}
	mi := &file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsEksNodeGroupStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsEksNodeGroupStatus) ProtoMessage() {}

func (x *AwsEksNodeGroupStatus) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsEksNodeGroupStatus.ProtoReflect.Descriptor instead.
func (*AwsEksNodeGroupStatus) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *AwsEksNodeGroupStatus) GetLifecycle() *shared.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *AwsEksNodeGroupStatus) GetAudit() *shared.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *AwsEksNodeGroupStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *AwsEksNodeGroupStatus) GetOutputs() *AwsEksNodeGroupStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_project_planton_provider_aws_awseksnodegroup_v1_api_proto protoreflect.FileDescriptor

const file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDesc = "" +
	"\n" +
	"9project/planton/provider/aws/awseksnodegroup/v1/api.proto\x12/project.planton.provider.aws.awseksnodegroup.v1\x1a\x1bbuf/validate/validate.proto\x1a:project/planton/provider/aws/awseksnodegroup/v1/spec.proto\x1aCproject/planton/provider/aws/awseksnodegroup/v1/stack_outputs.proto\x1a#project/planton/shared/status.proto\x1a%project/planton/shared/metadata.proto\"\x94\x03\n" +
	"\x0fAwsEksNodeGroup\x12B\n" +
	"\vapi_version\x18\x01 \x01(\tB!\xbaH\x1er\x1c\n" +
	"\x1aaws.project-planton.org/v1R\n" +
	"apiVersion\x12*\n" +
	"\x04kind\x18\x02 \x01(\tB\x16\xbaH\x13r\x11\n" +
	"\x0fAwsEksNodeGroupR\x04kind\x12O\n" +
	"\bmetadata\x18\x03 \x01(\v2+.project.planton.shared.ApiResourceMetadataB\x06\xbaH\x03\xc8\x01\x01R\bmetadata\x12`\n" +
	"\x04spec\x18\x04 \x01(\v2D.project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupSpecB\x06\xbaH\x03\xc8\x01\x01R\x04spec\x12^\n" +
	"\x06status\x18\x05 \x01(\v2F.project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStatusR\x06status\"\xad\x02\n" +
	"\x15AwsEksNodeGroupStatus\x12J\n" +
	"\tlifecycle\x18c \x01(\v2,.project.planton.shared.ApiResourceLifecycleR\tlifecycle\x12>\n" +
	"\x05audit\x18b \x01(\v2(.project.planton.shared.ApiResourceAuditR\x05audit\x12 \n" +
	"\fstack_job_id\x18a \x01(\tR\n" +
	"stackJobId\x12f\n" +
	"\aoutputs\x18\x01 \x01(\v2L.project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackOutputsR\aoutputsB\x95\x03\n" +
	"3com.project.planton.provider.aws.awseksnodegroup.v1B\bApiProtoP\x01Zqgithub.com/project-planton/project-planton/apis/project/planton/provider/aws/awseksnodegroup/v1;awseksnodegroupv1\xa2\x02\x05PPPAA\xaa\x02/Project.Planton.Provider.Aws.Awseksnodegroup.V1\xca\x02/Project\\Planton\\Provider\\Aws\\Awseksnodegroup\\V1\xe2\x02;Project\\Planton\\Provider\\Aws\\Awseksnodegroup\\V1\\GPBMetadata\xea\x024Project::Planton::Provider::Aws::Awseksnodegroup::V1b\x06proto3"

var (
	file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescData []byte
)

func file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDesc), len(file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDesc)))
	})
	return file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDescData
}

var file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_goTypes = []any{
	(*AwsEksNodeGroup)(nil),             // 0: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroup
	(*AwsEksNodeGroupStatus)(nil),       // 1: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStatus
	(*shared.ApiResourceMetadata)(nil),  // 2: project.planton.shared.ApiResourceMetadata
	(*AwsEksNodeGroupSpec)(nil),         // 3: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupSpec
	(*shared.ApiResourceLifecycle)(nil), // 4: project.planton.shared.ApiResourceLifecycle
	(*shared.ApiResourceAudit)(nil),     // 5: project.planton.shared.ApiResourceAudit
	(*AwsEksNodeGroupStackOutputs)(nil), // 6: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackOutputs
}
var file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroup.metadata:type_name -> project.planton.shared.ApiResourceMetadata
	3, // 1: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroup.spec:type_name -> project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupSpec
	1, // 2: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroup.status:type_name -> project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStatus
	4, // 3: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStatus.lifecycle:type_name -> project.planton.shared.ApiResourceLifecycle
	5, // 4: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStatus.audit:type_name -> project.planton.shared.ApiResourceAudit
	6, // 5: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStatus.outputs:type_name -> project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackOutputs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_init() }
func file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_init() {
	if File_project_planton_provider_aws_awseksnodegroup_v1_api_proto != nil {
		return
	}
	file_project_planton_provider_aws_awseksnodegroup_v1_spec_proto_init()
	file_project_planton_provider_aws_awseksnodegroup_v1_stack_outputs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDesc), len(file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awseksnodegroup_v1_api_proto = out.File
	file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_goTypes = nil
	file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_depIdxs = nil
}
