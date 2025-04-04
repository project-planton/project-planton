// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/aws/awsecrrepo/v1/spec.proto

package awsecrrepov1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/project-planton/project-planton/apis/project/planton/shared/options"
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

// ImageTagMutability defines whether tags in the ECR repository can be overwritten
// after being pushed. Most production scenarios recommend IMMUTABLE to prevent
// accidental or malicious overwrites.
type ImageTagMutability int32

const (
	// image_tag_mutability_unspecified is not a valid setting but serves as a placeholder.
	ImageTagMutability_IMAGE_TAG_MUTABILITY_UNSPECIFIED ImageTagMutability = 0
	// 'MUTABLE' allows overwriting tags once pushed.
	ImageTagMutability_MUTABLE ImageTagMutability = 1
	// 'IMMUTABLE' prevents overwriting tags once pushed.
	ImageTagMutability_IMMUTABLE ImageTagMutability = 2
)

// Enum value maps for ImageTagMutability.
var (
	ImageTagMutability_name = map[int32]string{
		0: "IMAGE_TAG_MUTABILITY_UNSPECIFIED",
		1: "MUTABLE",
		2: "IMMUTABLE",
	}
	ImageTagMutability_value = map[string]int32{
		"IMAGE_TAG_MUTABILITY_UNSPECIFIED": 0,
		"MUTABLE":                          1,
		"IMMUTABLE":                        2,
	}
)

func (x ImageTagMutability) Enum() *ImageTagMutability {
	p := new(ImageTagMutability)
	*p = x
	return p
}

func (x ImageTagMutability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImageTagMutability) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_enumTypes[0].Descriptor()
}

func (ImageTagMutability) Type() protoreflect.EnumType {
	return &file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_enumTypes[0]
}

func (x ImageTagMutability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImageTagMutability.Descriptor instead.
func (ImageTagMutability) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescGZIP(), []int{0}
}

// EncryptionType determines how images are encrypted at rest in ECR.
// By default, AWS uses AES-256 (service-managed keys). Choose KMS to use
// a customer-managed key (CMK).
type EncryptionType int32

const (
	// encryption_type_unspecified is not a valid setting but serves as a placeholder.
	EncryptionType_ENCRYPTION_TYPE_UNSPECIFIED EncryptionType = 0
	// 'AES256' uses AWS-managed encryption (AES-256 SSE) for ECR.
	EncryptionType_AES256 EncryptionType = 1
	// 'KMS' uses an AWS KMS key for encryption.
	EncryptionType_KMS EncryptionType = 2
)

// Enum value maps for EncryptionType.
var (
	EncryptionType_name = map[int32]string{
		0: "ENCRYPTION_TYPE_UNSPECIFIED",
		1: "AES256",
		2: "KMS",
	}
	EncryptionType_value = map[string]int32{
		"ENCRYPTION_TYPE_UNSPECIFIED": 0,
		"AES256":                      1,
		"KMS":                         2,
	}
)

func (x EncryptionType) Enum() *EncryptionType {
	p := new(EncryptionType)
	*p = x
	return p
}

func (x EncryptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncryptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_enumTypes[1].Descriptor()
}

func (EncryptionType) Type() protoreflect.EnumType {
	return &file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_enumTypes[1]
}

func (x EncryptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncryptionType.Descriptor instead.
func (EncryptionType) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescGZIP(), []int{1}
}

// AwsEcrRepoSpec defines the configuration for creating an AWS ECR repository
// to store and manage Docker images. Most fields are optional, with recommended
// defaults aligned to best practices (immutable tags, scanning enabled, encryption).
type AwsEcrRepoSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repositoryName is the name of the ECR repository. Must be unique within
	// the AWS account and region. Commonly includes the microservice or project name.
	// Example: "github.com/teamblue/my-microservice"
	RepositoryName string `protobuf:"bytes,1,opt,name=repository_name,json=repositoryName,proto3" json:"repository_name,omitempty"`
	// imageTagMutability indicates whether image tags can be overwritten (MUTABLE)
	// or not (IMMUTABLE). Recommended default is immutable for production safety.
	ImageTagMutability ImageTagMutability `protobuf:"varint,2,opt,name=image_tag_mutability,json=imageTagMutability,proto3,enum=project.planton.provider.aws.awsecrrepo.v1.ImageTagMutability" json:"image_tag_mutability,omitempty"`
	// encryptionType determines how ECR encrypts images at rest. Default is AES256,
	// using AWS-managed encryption. Use KMS to specify your own KMS key for compliance.
	EncryptionType EncryptionType `protobuf:"varint,3,opt,name=encryption_type,json=encryptionType,proto3,enum=project.planton.provider.aws.awsecrrepo.v1.EncryptionType" json:"encryption_type,omitempty"`
	// kmsKeyId is the ARN or ID of a KMS key used when encryption_type = KMS.
	// If omitted, AWS uses the default service-managed key for ECR.
	// Ignored if encryption_type = AES256.
	KmsKeyId string `protobuf:"bytes,4,opt,name=kms_key_id,json=kmsKeyId,proto3" json:"kms_key_id,omitempty"`
	// forceDelete, if true, allows deleting the repository even when it contains
	// images (all images get removed on delete). By default, it is false, preventing
	// accidental data loss.
	ForceDelete bool `protobuf:"varint,5,opt,name=force_delete,json=forceDelete,proto3" json:"force_delete,omitempty"`
}

func (x *AwsEcrRepoSpec) Reset() {
	*x = AwsEcrRepoSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsEcrRepoSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsEcrRepoSpec) ProtoMessage() {}

func (x *AwsEcrRepoSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsEcrRepoSpec.ProtoReflect.Descriptor instead.
func (*AwsEcrRepoSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *AwsEcrRepoSpec) GetRepositoryName() string {
	if x != nil {
		return x.RepositoryName
	}
	return ""
}

func (x *AwsEcrRepoSpec) GetImageTagMutability() ImageTagMutability {
	if x != nil {
		return x.ImageTagMutability
	}
	return ImageTagMutability_IMAGE_TAG_MUTABILITY_UNSPECIFIED
}

func (x *AwsEcrRepoSpec) GetEncryptionType() EncryptionType {
	if x != nil {
		return x.EncryptionType
	}
	return EncryptionType_ENCRYPTION_TYPE_UNSPECIFIED
}

func (x *AwsEcrRepoSpec) GetKmsKeyId() string {
	if x != nil {
		return x.KmsKeyId
	}
	return ""
}

func (x *AwsEcrRepoSpec) GetForceDelete() bool {
	if x != nil {
		return x.ForceDelete
	}
	return false
}

var File_project_planton_provider_aws_awsecrrepo_v1_spec_proto protoreflect.FileDescriptor

var file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61,
	0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86,
	0x03, 0x0a, 0x0e, 0x41, 0x77, 0x73, 0x45, 0x63, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x36, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8,
	0x01, 0x01, 0x72, 0x05, 0x10, 0x02, 0x18, 0x80, 0x02, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x7f, 0x0a, 0x14, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x6d, 0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x3e, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70,
	0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67, 0x4d, 0x75, 0x74,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x42, 0x0d, 0x8a, 0xa6, 0x1d, 0x09, 0x49, 0x4d, 0x4d,
	0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x52, 0x12, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x54, 0x61, 0x67,
	0x4d, 0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x6f, 0x0a, 0x0f, 0x65, 0x6e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61,
	0x77, 0x73, 0x2e, 0x61, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x0a, 0x8a, 0xa6, 0x1d, 0x06, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x52, 0x0e, 0x65, 0x6e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x0a, 0x6b,
	0x6d, 0x73, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6b, 0x6d, 0x73, 0x4b, 0x65, 0x79, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x0c, 0x66, 0x6f, 0x72,
	0x63, 0x65, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x09, 0x8a, 0xa6, 0x1d, 0x05, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x52, 0x0b, 0x66, 0x6f, 0x72, 0x63,
	0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2a, 0x56, 0x0a, 0x12, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x54, 0x61, 0x67, 0x4d, 0x75, 0x74, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x24, 0x0a,
	0x20, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x5f, 0x54, 0x41, 0x47, 0x5f, 0x4d, 0x55, 0x54, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x49, 0x4d, 0x4d, 0x55, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x02, 0x2a,
	0x46, 0x0a, 0x0e, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x4e, 0x43, 0x52, 0x59, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x41, 0x45, 0x53, 0x32, 0x35, 0x36, 0x10, 0x01, 0x12, 0x07,
	0x0a, 0x03, 0x4b, 0x4d, 0x53, 0x10, 0x02, 0x42, 0xf3, 0x02, 0x0a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x61, 0x77, 0x73, 0x2e, 0x61, 0x77, 0x73,
	0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x70, 0x65, 0x63,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x61, 0x77, 0x73, 0x2f, 0x61, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x76, 0x31,
	0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x41, 0x41, 0xaa, 0x02, 0x2a, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x41, 0x77, 0x73, 0x2e, 0x41, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65,
	0x70, 0x6f, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x2a, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c,
	0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x5c, 0x41, 0x77, 0x73, 0x5c, 0x41, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x36, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x41, 0x77,
	0x73, 0x5c, 0x41, 0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x2f, 0x50, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x41, 0x77, 0x73, 0x3a, 0x3a, 0x41,
	0x77, 0x73, 0x65, 0x63, 0x72, 0x72, 0x65, 0x70, 0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescData = file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDesc
)

func file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescData)
	})
	return file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDescData
}

var file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_goTypes = []any{
	(ImageTagMutability)(0), // 0: project.planton.provider.aws.awsecrrepo.v1.ImageTagMutability
	(EncryptionType)(0),     // 1: project.planton.provider.aws.awsecrrepo.v1.EncryptionType
	(*AwsEcrRepoSpec)(nil),  // 2: project.planton.provider.aws.awsecrrepo.v1.AwsEcrRepoSpec
}
var file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_depIdxs = []int32{
	0, // 0: project.planton.provider.aws.awsecrrepo.v1.AwsEcrRepoSpec.image_tag_mutability:type_name -> project.planton.provider.aws.awsecrrepo.v1.ImageTagMutability
	1, // 1: project.planton.provider.aws.awsecrrepo.v1.AwsEcrRepoSpec.encryption_type:type_name -> project.planton.provider.aws.awsecrrepo.v1.EncryptionType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_init() }
func file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_init() {
	if File_project_planton_provider_aws_awsecrrepo_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AwsEcrRepoSpec); i {
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
			RawDescriptor: file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_depIdxs,
		EnumInfos:         file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_enumTypes,
		MessageInfos:      file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_msgTypes,
	}.Build()
	File_project_planton_provider_aws_awsecrrepo_v1_spec_proto = out.File
	file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_rawDesc = nil
	file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_goTypes = nil
	file_project_planton_provider_aws_awsecrrepo_v1_spec_proto_depIdxs = nil
}
