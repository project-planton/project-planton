// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/aws/awsiamrole/v1/spec.proto

package awsiamrolev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/project-planton/project-planton/apis/project/planton/shared/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AwsIamRoleSpec defines the minimal fields needed to create an AWS IAM Role.
// It includes the trust policy JSON, managed policies, inline policies, and more.
type AwsIamRoleSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// description is an optional description of the IAM role.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	// path is the IAM path for the role. Defaults to "/" if omitted.
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	// trust_policy_json is the JSON string describing the trust relationship for the role.
	// Example: a trust policy allowing ECS tasks to assume this role.
	TrustPolicy *structpb.Struct `protobuf:"bytes,3,opt,name=trust_policy,json=trustPolicy,proto3" json:"trust_policy,omitempty"`
	// managed_policy_arns is a list of ARNs for AWS-managed or customer-managed IAM policies
	// you want to attach to this role.
	ManagedPolicyArns []string `protobuf:"bytes,4,rep,name=managed_policy_arns,json=managedPolicyArns,proto3" json:"managed_policy_arns,omitempty"`
	// inline_policy_jsons is a map of inline policy names to a JSON policy doc.
	// Key is policy name. Value is the raw JSON for that policy.
	InlinePolicies map[string]*structpb.Struct `protobuf:"bytes,5,rep,name=inline_policies,json=inlinePolicies,proto3" json:"inline_policies,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *AwsIamRoleSpec) Reset() {
	*x = AwsIamRoleSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_aws_awsiamrole_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsIamRoleSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsIamRoleSpec) ProtoMessage() {}

func (x *AwsIamRoleSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awsiamrole_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsIamRoleSpec.ProtoReflect.Descriptor instead.
func (*AwsIamRoleSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *AwsIamRoleSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AwsIamRoleSpec) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *AwsIamRoleSpec) GetTrustPolicy() *structpb.Struct {
	if x != nil {
		return x.TrustPolicy
	}
	return nil
}

func (x *AwsIamRoleSpec) GetManagedPolicyArns() []string {
	if x != nil {
		return x.ManagedPolicyArns
	}
	return nil
}

func (x *AwsIamRoleSpec) GetInlinePolicies() map[string]*structpb.Struct {
	if x != nil {
		return x.InlinePolicies
	}
	return nil
}

var File_project_planton_provider_aws_awsiamrole_v1_spec_proto protoreflect.FileDescriptor

var file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61,
	0x77, 0x73, 0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x03, 0x0a,
	0x0e, 0x41, 0x77, 0x73, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x19, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x05, 0x92, 0xa6, 0x1d, 0x01, 0x2f, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x42, 0x0a, 0x0c,
	0x74, 0x72, 0x75, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x06, 0xba, 0x48, 0x03,
	0xc8, 0x01, 0x01, 0x52, 0x0b, 0x74, 0x72, 0x75, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x12, 0x38, 0x0a, 0x13, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x5f, 0x70, 0x6f, 0x6c, 0x69,
	0x63, 0x79, 0x5f, 0x61, 0x72, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08, 0xba,
	0x48, 0x05, 0x92, 0x01, 0x02, 0x18, 0x01, 0x52, 0x11, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x41, 0x72, 0x6e, 0x73, 0x12, 0x77, 0x0a, 0x0f, 0x69, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x4e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x77, 0x73, 0x49, 0x61, 0x6d, 0x52, 0x6f, 0x6c, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e,
	0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0e, 0x69, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x69, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x13, 0x49, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42,
	0xf3, 0x02, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65, 0x2e,
	0x76, 0x31, 0x42, 0x09, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77,
	0x73, 0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x77, 0x73, 0x69,
	0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x41, 0x41,
	0xaa, 0x02, 0x2a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x77, 0x73, 0x2e,
	0x41, 0x77, 0x73, 0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x2a,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x41, 0x77, 0x73,
	0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x36, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x41, 0x77, 0x73, 0x69, 0x61, 0x6d,
	0x72, 0x6f, 0x6c, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50,
	0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x69, 0x61, 0x6d, 0x72, 0x6f, 0x6c,
	0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescData = file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDesc
)

func file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescData)
	})
	return file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDescData
}

var file_project_planton_provider_aws_awsiamrole_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_aws_awsiamrole_v1_spec_proto_goTypes = []any{
	(*AwsIamRoleSpec)(nil),  // 0: project.planton.provider.aws.awsiamrole.v1.AwsIamRoleSpec
	nil,                     // 1: project.planton.provider.aws.awsiamrole.v1.AwsIamRoleSpec.InlinePoliciesEntry
	(*structpb.Struct)(nil), // 2: google.protobuf.Struct
}
var file_project_planton_provider_aws_awsiamrole_v1_spec_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.aws.awsiamrole.v1.AwsIamRoleSpec.trust_policy:type_name -> google.protobuf.Struct
	1, // 1: project.planton.provider.aws.awsiamrole.v1.AwsIamRoleSpec.inline_policies:type_name -> project.planton.provider.aws.awsiamrole.v1.AwsIamRoleSpec.InlinePoliciesEntry
	2, // 2: project.planton.provider.aws.awsiamrole.v1.AwsIamRoleSpec.InlinePoliciesEntry.value:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awsiamrole_v1_spec_proto_init() }
func file_project_planton_provider_aws_awsiamrole_v1_spec_proto_init() {
	if File_project_planton_provider_aws_awsiamrole_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_aws_awsiamrole_v1_spec_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AwsIamRoleSpec); i {
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
			RawDescriptor: file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awsiamrole_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awsiamrole_v1_spec_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_aws_awsiamrole_v1_spec_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awsiamrole_v1_spec_proto = out.File
	file_project_planton_provider_aws_awsiamrole_v1_spec_proto_rawDesc = nil
	file_project_planton_provider_aws_awsiamrole_v1_spec_proto_goTypes = nil
	file_project_planton_provider_aws_awsiamrole_v1_spec_proto_depIdxs = nil
}
