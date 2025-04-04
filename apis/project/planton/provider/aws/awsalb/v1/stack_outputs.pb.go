// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/aws/awsalb/v1/stack_outputs.proto

package awsalbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AwsAlbStackOutputs describes the outputs returned by Pulumi/Terraform after creating an ALB.
type AwsAlbStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// load_balancer_arn is the ARN of the created Application Load Balancer.
	LoadBalancerArn string `protobuf:"bytes,1,opt,name=load_balancer_arn,json=loadBalancerArn,proto3" json:"load_balancer_arn,omitempty"`
	// load_balancer_name is the final name assigned to the ALB (may differ from metadata.name).
	LoadBalancerName string `protobuf:"bytes,2,opt,name=load_balancer_name,json=loadBalancerName,proto3" json:"load_balancer_name,omitempty"`
	// load_balancer_dns_name is the DNS name automatically assigned to the ALB.
	LoadBalancerDnsName string `protobuf:"bytes,3,opt,name=load_balancer_dns_name,json=loadBalancerDnsName,proto3" json:"load_balancer_dns_name,omitempty"`
	// load_balancer_hosted_zone_id is the Route53 hosted zone ID for the ALB's DNS entry.
	LoadBalancerHostedZoneId string `protobuf:"bytes,4,opt,name=load_balancer_hosted_zone_id,json=loadBalancerHostedZoneId,proto3" json:"load_balancer_hosted_zone_id,omitempty"`
}

func (x *AwsAlbStackOutputs) Reset() {
	*x = AwsAlbStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsAlbStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsAlbStackOutputs) ProtoMessage() {}

func (x *AwsAlbStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsAlbStackOutputs.ProtoReflect.Descriptor instead.
func (*AwsAlbStackOutputs) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescGZIP(), []int{0}
}

func (x *AwsAlbStackOutputs) GetLoadBalancerArn() string {
	if x != nil {
		return x.LoadBalancerArn
	}
	return ""
}

func (x *AwsAlbStackOutputs) GetLoadBalancerName() string {
	if x != nil {
		return x.LoadBalancerName
	}
	return ""
}

func (x *AwsAlbStackOutputs) GetLoadBalancerDnsName() string {
	if x != nil {
		return x.LoadBalancerDnsName
	}
	return ""
}

func (x *AwsAlbStackOutputs) GetLoadBalancerHostedZoneId() string {
	if x != nil {
		return x.LoadBalancerHostedZoneId
	}
	return ""
}

var File_project_planton_provider_aws_awsalb_v1_stack_outputs_proto protoreflect.FileDescriptor

var file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDesc = []byte{
	0x0a, 0x3a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61,
	0x77, 0x73, 0x61, 0x6c, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x61, 0x6c,
	0x62, 0x2e, 0x76, 0x31, 0x22, 0xe3, 0x01, 0x0a, 0x12, 0x41, 0x77, 0x73, 0x41, 0x6c, 0x62, 0x53,
	0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x12, 0x2a, 0x0a, 0x11, 0x6c,
	0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x61, 0x72, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x41, 0x72, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x6f, 0x61, 0x64, 0x5f,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x33, 0x0a, 0x16, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x64, 0x6e, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x72, 0x44, 0x6e, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x1c, 0x6c, 0x6f,
	0x61, 0x64, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x5f, 0x68, 0x6f, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x18, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x48, 0x6f,
	0x73, 0x74, 0x65, 0x64, 0x5a, 0x6f, 0x6e, 0x65, 0x49, 0x64, 0x42, 0xdf, 0x02, 0x0a, 0x2a, 0x63,
	0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e,
	0x61, 0x77, 0x73, 0x61, 0x6c, 0x62, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x74, 0x61, 0x63, 0x6b,
	0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x5f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73,
	0x61, 0x6c, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x77, 0x73, 0x61, 0x6c, 0x62, 0x76, 0x31, 0xa2,
	0x02, 0x05, 0x50, 0x50, 0x50, 0x41, 0x41, 0xaa, 0x02, 0x26, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x41, 0x77, 0x73, 0x2e, 0x41, 0x77, 0x73, 0x61, 0x6c, 0x62, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x26, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c,
	0x41, 0x77, 0x73, 0x61, 0x6c, 0x62, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x32, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77, 0x73, 0x5c, 0x41, 0x77, 0x73, 0x61, 0x6c, 0x62, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x2b, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x41, 0x77, 0x73,
	0x3a, 0x3a, 0x41, 0x77, 0x73, 0x61, 0x6c, 0x62, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescData = file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDesc
)

func file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescData)
	})
	return file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDescData
}

var file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_goTypes = []any{
	(*AwsAlbStackOutputs)(nil), // 0: project.planton.provider.aws.awsalb.v1.AwsAlbStackOutputs
}
var file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_init() }
func file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_init() {
	if File_project_planton_provider_aws_awsalb_v1_stack_outputs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AwsAlbStackOutputs); i {
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
			RawDescriptor: file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awsalb_v1_stack_outputs_proto = out.File
	file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_rawDesc = nil
	file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_goTypes = nil
	file_project_planton_provider_aws_awsalb_v1_stack_outputs_proto_depIdxs = nil
}
