// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcpartifactregistry/v1/stack_outputs.proto

package gcpartifactregistryv1

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

// gcp-artifact-registry stack outputs
type GcpArtifactRegistryStackOutputs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// email of the reader service account.
	ReaderServiceAccountEmail string `protobuf:"bytes,1,opt,name=reader_service_account_email,json=readerServiceAccountEmail,proto3" json:"reader_service_account_email,omitempty"`
	// base64 encoded key of the reader service account.
	ReaderServiceAccountKeyBase64 string `protobuf:"bytes,2,opt,name=reader_service_account_key_base64,json=readerServiceAccountKeyBase64,proto3" json:"reader_service_account_key_base64,omitempty"`
	// email of the writer service account.
	WriterServiceAccountEmail string `protobuf:"bytes,3,opt,name=writer_service_account_email,json=writerServiceAccountEmail,proto3" json:"writer_service_account_email,omitempty"`
	// base64 encoded key of the writer service account.
	WriterServiceAccountKeyBase64 string `protobuf:"bytes,4,opt,name=writer_service_account_key_base64,json=writerServiceAccountKeyBase64,proto3" json:"writer_service_account_key_base64,omitempty"`
	// name of the docker repo.
	DockerRepoName string `protobuf:"bytes,5,opt,name=docker_repo_name,json=dockerRepoName,proto3" json:"docker_repo_name,omitempty"`
	// hostname of the docker repo.
	DockerRepoHostname string `protobuf:"bytes,6,opt,name=docker_repo_hostname,json=dockerRepoHostname,proto3" json:"docker_repo_hostname,omitempty"`
	// url for the docker repository.
	DockerRepoUrl string `protobuf:"bytes,7,opt,name=docker_repo_url,json=dockerRepoUrl,proto3" json:"docker_repo_url,omitempty"`
	// name of the maven repo.
	MavenRepoName string `protobuf:"bytes,8,opt,name=maven_repo_name,json=mavenRepoName,proto3" json:"maven_repo_name,omitempty"`
	// url for the maven repository.
	MavenRepoUrl string `protobuf:"bytes,9,opt,name=maven_repo_url,json=mavenRepoUrl,proto3" json:"maven_repo_url,omitempty"`
	// name of the npm repo.
	NpmRepoName string `protobuf:"bytes,10,opt,name=npm_repo_name,json=npmRepoName,proto3" json:"npm_repo_name,omitempty"`
	// name of the python repo.
	PythonRepoName string `protobuf:"bytes,11,opt,name=python_repo_name,json=pythonRepoName,proto3" json:"python_repo_name,omitempty"`
}

func (x *GcpArtifactRegistryStackOutputs) Reset() {
	*x = GcpArtifactRegistryStackOutputs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpArtifactRegistryStackOutputs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpArtifactRegistryStackOutputs) ProtoMessage() {}

func (x *GcpArtifactRegistryStackOutputs) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpArtifactRegistryStackOutputs.ProtoReflect.Descriptor instead.
func (*GcpArtifactRegistryStackOutputs) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescGZIP(), []int{0}
}

func (x *GcpArtifactRegistryStackOutputs) GetReaderServiceAccountEmail() string {
	if x != nil {
		return x.ReaderServiceAccountEmail
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetReaderServiceAccountKeyBase64() string {
	if x != nil {
		return x.ReaderServiceAccountKeyBase64
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetWriterServiceAccountEmail() string {
	if x != nil {
		return x.WriterServiceAccountEmail
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetWriterServiceAccountKeyBase64() string {
	if x != nil {
		return x.WriterServiceAccountKeyBase64
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetDockerRepoName() string {
	if x != nil {
		return x.DockerRepoName
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetDockerRepoHostname() string {
	if x != nil {
		return x.DockerRepoHostname
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetDockerRepoUrl() string {
	if x != nil {
		return x.DockerRepoUrl
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetMavenRepoName() string {
	if x != nil {
		return x.MavenRepoName
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetMavenRepoUrl() string {
	if x != nil {
		return x.MavenRepoUrl
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetNpmRepoName() string {
	if x != nil {
		return x.NpmRepoName
	}
	return ""
}

func (x *GcpArtifactRegistryStackOutputs) GetPythonRepoName() string {
	if x != nil {
		return x.PythonRepoName
	}
	return ""
}

var File_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto protoreflect.FileDescriptor

var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDesc = []byte{
	0x0a, 0x47, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x5f, 0x6f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x33, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x22, 0xd7,
	0x04, 0x0a, 0x1f, 0x47, 0x63, 0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x12, 0x3f, 0x0a, 0x1c, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x19, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x48, 0x0a, 0x21, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65,
	0x79, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d,
	0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x12, 0x3f, 0x0a,
	0x1c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x19, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x48,
	0x0a, 0x21, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x5f, 0x62, 0x61, 0x73,
	0x65, 0x36, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1d, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b,
	0x65, 0x79, 0x42, 0x61, 0x73, 0x65, 0x36, 0x34, 0x12, 0x28, 0x0a, 0x10, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70,
	0x6f, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x12, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x48, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64,
	0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x26, 0x0a, 0x0f,
	0x6d, 0x61, 0x76, 0x65, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6d, 0x61, 0x76, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6f,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x6d, 0x61, 0x76, 0x65, 0x6e, 0x5f, 0x72, 0x65,
	0x70, 0x6f, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x61,
	0x76, 0x65, 0x6e, 0x52, 0x65, 0x70, 0x6f, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x70,
	0x6d, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6e, 0x70, 0x6d, 0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70, 0x79, 0x74, 0x68, 0x6f, 0x6e,
	0x52, 0x65, 0x70, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0xba, 0x03, 0x0a, 0x37, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x53, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x79, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x67,
	0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x47, 0x47, 0xaa, 0x02, 0x33, 0x50,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x63, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x33, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63,
	0x70, 0x5c, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x3f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x38, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescOnce sync.Once
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescData = file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDesc
)

func file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescGZIP() []byte {
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescData)
	})
	return file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDescData
}

var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_goTypes = []any{
	(*GcpArtifactRegistryStackOutputs)(nil), // 0: project.planton.provider.gcp.gcpartifactregistry.v1.GcpArtifactRegistryStackOutputs
}
var file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_init() }
func file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_init() {
	if File_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GcpArtifactRegistryStackOutputs); i {
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
			RawDescriptor: file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_msgTypes,
	}.Build()
	File_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto = out.File
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_rawDesc = nil
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_goTypes = nil
	file_project_planton_provider_gcp_gcpartifactregistry_v1_stack_outputs_proto_depIdxs = nil
}