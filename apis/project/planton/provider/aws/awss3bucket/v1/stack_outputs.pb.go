// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/aws/awss3bucket/v1/stack_outputs.proto

package awss3bucketv1

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

// aws-s3-bucket stack outputs
type AwsS3BucketStackOutputs struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// id of the storage-bucket created on aws
	BucketId      string `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AwsS3BucketStackOutputs) Reset() {
	*x = AwsS3BucketStackOutputs{}
	mi := &file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AwsS3BucketStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsS3BucketStackOutputs) ProtoMessage() {}

func (x *AwsS3BucketStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsS3BucketStackOutputs.ProtoReflect.Descriptor instead.
func (*AwsS3BucketStackOutputs) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDescGZIP(), []int{0}
}

func (x *AwsS3BucketStackOutputs) GetBucketId() string {
	if x != nil {
		return x.BucketId
	}
	return ""
}

var File_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto protoreflect.FileDescriptor

const file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDesc = "" +
	"\n" +
	"?project/planton/provider/aws/awss3bucket/v1/stack_outputs.proto\x12+project.planton.provider.aws.awss3bucket.v1\"6\n" +
	"\x17AwsS3BucketStackOutputs\x12\x1b\n" +
	"\tbucket_id\x18\x01 \x01(\tR\bbucketIdB\x82\x03\n" +
	"/com.project.planton.provider.aws.awss3bucket.v1B\x11StackOutputsProtoP\x01Zigithub.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1;awss3bucketv1\xa2\x02\x05PPPAA\xaa\x02+Project.Planton.Provider.Aws.Awss3bucket.V1\xca\x02+Project\\Planton\\Provider\\Aws\\Awss3bucket\\V1\xe2\x027Project\\Planton\\Provider\\Aws\\Awss3bucket\\V1\\GPBMetadata\xea\x020Project::Planton::Provider::Aws::Awss3bucket::V1b\x06proto3"

var (
	file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDescData []byte
)

func file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDesc)))
	})
	return file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDescData
}

var file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_goTypes = []any{
	(*AwsS3BucketStackOutputs)(nil), // 0: project.planton.provider.aws.awss3bucket.v1.AwsS3BucketStackOutputs
}
var file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_init() }
func file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_init() {
	if File_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDesc), len(file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto = out.File
	file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_goTypes = nil
	file_project_planton_provider_aws_awss3bucket_v1_stack_outputs_proto_depIdxs = nil
}
