// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/aws/awseksnodegroup/v1/stack_input.proto

package awseksnodegroupv1

import (
	v1 "github.com/project-planton/project-planton/apis/project/planton/credential/awscredential/v1"
	shared "github.com/project-planton/project-planton/apis/project/planton/shared"
	pulumi "github.com/project-planton/project-planton/apis/project/planton/shared/iac/pulumi"
	terraform "github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
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

// aws-eks-node-group stack-input
type AwsEksNodeGroupStackInput struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// iac-provisioner
	Provisioner shared.IacProvisioner `protobuf:"varint,1,opt,name=provisioner,proto3,enum=project.planton.shared.IacProvisioner" json:"provisioner,omitempty"`
	// pulumi input required when the provisioner is pulumi
	Pulumi *pulumi.PulumiStackInfo `protobuf:"bytes,2,opt,name=pulumi,proto3" json:"pulumi,omitempty"`
	// terraform input required when the provisioner is terraform
	Terraform *terraform.TerraformStackInfo `protobuf:"bytes,3,opt,name=terraform,proto3" json:"terraform,omitempty"`
	// target api-resource
	Target *AwsEksNodeGroup `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
	// provider-credential
	ProviderCredential *v1.AwsCredentialSpec `protobuf:"bytes,5,opt,name=provider_credential,json=providerCredential,proto3" json:"provider_credential,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *AwsEksNodeGroupStackInput) Reset() {
	*x = AwsEksNodeGroupStackInput{}
	mi := &file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsEksNodeGroupStackInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsEksNodeGroupStackInput) ProtoMessage() {}

func (x *AwsEksNodeGroupStackInput) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsEksNodeGroupStackInput.ProtoReflect.Descriptor instead.
func (*AwsEksNodeGroupStackInput) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDescGZIP(), []int{0}
}

func (x *AwsEksNodeGroupStackInput) GetProvisioner() shared.IacProvisioner {
	if x != nil {
		return x.Provisioner
	}
	return shared.IacProvisioner(0)
}

func (x *AwsEksNodeGroupStackInput) GetPulumi() *pulumi.PulumiStackInfo {
	if x != nil {
		return x.Pulumi
	}
	return nil
}

func (x *AwsEksNodeGroupStackInput) GetTerraform() *terraform.TerraformStackInfo {
	if x != nil {
		return x.Terraform
	}
	return nil
}

func (x *AwsEksNodeGroupStackInput) GetTarget() *AwsEksNodeGroup {
	if x != nil {
		return x.Target
	}
	return nil
}

func (x *AwsEksNodeGroupStackInput) GetProviderCredential() *v1.AwsCredentialSpec {
	if x != nil {
		return x.ProviderCredential
	}
	return nil
}

var File_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto protoreflect.FileDescriptor

const file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDesc = "" +
	"\n" +
	"Aproject/planton/provider/aws/awseksnodegroup/v1/stack_input.proto\x12/project.planton.provider.aws.awseksnodegroup.v1\x1a6project/planton/credential/awscredential/v1/spec.proto\x1a9project/planton/provider/aws/awseksnodegroup/v1/api.proto\x1a.project/planton/shared/iac/pulumi/pulumi.proto\x1a project/planton/shared/iac.proto\x1a4project/planton/shared/iac/terraform/terraform.proto\"\xd4\x03\n" +
	"\x19AwsEksNodeGroupStackInput\x12H\n" +
	"\vprovisioner\x18\x01 \x01(\x0e2&.project.planton.shared.IacProvisionerR\vprovisioner\x12J\n" +
	"\x06pulumi\x18\x02 \x01(\v22.project.planton.shared.iac.pulumi.PulumiStackInfoR\x06pulumi\x12V\n" +
	"\tterraform\x18\x03 \x01(\v28.project.planton.shared.iac.terraform.TerraformStackInfoR\tterraform\x12X\n" +
	"\x06target\x18\x04 \x01(\v2@.project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupR\x06target\x12o\n" +
	"\x13provider_credential\x18\x05 \x01(\v2>.project.planton.credential.awscredential.v1.AwsCredentialSpecR\x12providerCredentialB\x9c\x03\n" +
	"3com.project.planton.provider.aws.awseksnodegroup.v1B\x0fStackInputProtoP\x01Zqgithub.com/project-planton/project-planton/apis/project/planton/provider/aws/awseksnodegroup/v1;awseksnodegroupv1\xa2\x02\x05PPPAA\xaa\x02/Project.Planton.Provider.Aws.Awseksnodegroup.V1\xca\x02/Project\\Planton\\Provider\\Aws\\Awseksnodegroup\\V1\xe2\x02;Project\\Planton\\Provider\\Aws\\Awseksnodegroup\\V1\\GPBMetadata\xea\x024Project::Planton::Provider::Aws::Awseksnodegroup::V1b\x06proto3"

var (
	file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDescData []byte
)

func file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDesc), len(file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDesc)))
	})
	return file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDescData
}

var file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_goTypes = []any{
	(*AwsEksNodeGroupStackInput)(nil),    // 0: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackInput
	(shared.IacProvisioner)(0),           // 1: project.planton.shared.IacProvisioner
	(*pulumi.PulumiStackInfo)(nil),       // 2: project.planton.shared.iac.pulumi.PulumiStackInfo
	(*terraform.TerraformStackInfo)(nil), // 3: project.planton.shared.iac.terraform.TerraformStackInfo
	(*AwsEksNodeGroup)(nil),              // 4: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroup
	(*v1.AwsCredentialSpec)(nil),         // 5: project.planton.credential.awscredential.v1.AwsCredentialSpec
}
var file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_depIdxs = []int32{
	1, // 0: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackInput.provisioner:type_name -> project.planton.shared.IacProvisioner
	2, // 1: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackInput.pulumi:type_name -> project.planton.shared.iac.pulumi.PulumiStackInfo
	3, // 2: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackInput.terraform:type_name -> project.planton.shared.iac.terraform.TerraformStackInfo
	4, // 3: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackInput.target:type_name -> project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroup
	5, // 4: project.planton.provider.aws.awseksnodegroup.v1.AwsEksNodeGroupStackInput.provider_credential:type_name -> project.planton.credential.awscredential.v1.AwsCredentialSpec
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_init() }
func file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_init() {
	if File_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto != nil {
		return
	}
	file_project_planton_provider_aws_awseksnodegroup_v1_api_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDesc), len(file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto = out.File
	file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_goTypes = nil
	file_project_planton_provider_aws_awseksnodegroup_v1_stack_input_proto_depIdxs = nil
}
