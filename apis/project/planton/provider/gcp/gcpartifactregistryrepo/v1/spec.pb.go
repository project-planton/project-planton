// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcpartifactregistryrepo/v1/spec.proto

package gcpartifactregistryrepov1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

// enumeration for supported formats - https://cloud.google.com/artifact-registry/docs/supported-formats
// note: naming the values using uppercase letters to match the naming convention of the supported formats
type GcpArtifactRegistryRepoFormat int32

const (
	GcpArtifactRegistryRepoFormat_gcp_artifact_registry_repo_format_unspecified GcpArtifactRegistryRepoFormat = 0
	GcpArtifactRegistryRepoFormat_DOCKER                                        GcpArtifactRegistryRepoFormat = 1
	GcpArtifactRegistryRepoFormat_GENERIC                                       GcpArtifactRegistryRepoFormat = 2
	GcpArtifactRegistryRepoFormat_GO                                            GcpArtifactRegistryRepoFormat = 3
	GcpArtifactRegistryRepoFormat_KUBEFLOW                                      GcpArtifactRegistryRepoFormat = 4
	GcpArtifactRegistryRepoFormat_MAVEN                                         GcpArtifactRegistryRepoFormat = 5
	GcpArtifactRegistryRepoFormat_NPM                                           GcpArtifactRegistryRepoFormat = 6
	GcpArtifactRegistryRepoFormat_PYTHON                                        GcpArtifactRegistryRepoFormat = 7
	GcpArtifactRegistryRepoFormat_YUM                                           GcpArtifactRegistryRepoFormat = 8
)

// Enum value maps for GcpArtifactRegistryRepoFormat.
var (
	GcpArtifactRegistryRepoFormat_name = map[int32]string{
		0: "gcp_artifact_registry_repo_format_unspecified",
		1: "DOCKER",
		2: "GENERIC",
		3: "GO",
		4: "KUBEFLOW",
		5: "MAVEN",
		6: "NPM",
		7: "PYTHON",
		8: "YUM",
	}
	GcpArtifactRegistryRepoFormat_value = map[string]int32{
		"gcp_artifact_registry_repo_format_unspecified": 0,
		"DOCKER":   1,
		"GENERIC":  2,
		"GO":       3,
		"KUBEFLOW": 4,
		"MAVEN":    5,
		"NPM":      6,
		"PYTHON":   7,
		"YUM":      8,
	}
)

func (x GcpArtifactRegistryRepoFormat) Enum() *GcpArtifactRegistryRepoFormat {
	p := new(GcpArtifactRegistryRepoFormat)
	*p = x
	return p
}

func (x GcpArtifactRegistryRepoFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GcpArtifactRegistryRepoFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_enumTypes[0].Descriptor()
}

func (GcpArtifactRegistryRepoFormat) Type() protoreflect.EnumType {
	return &file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_enumTypes[0]
}

func (x GcpArtifactRegistryRepoFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GcpArtifactRegistryRepoFormat.Descriptor instead.
func (GcpArtifactRegistryRepoFormat) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescGZIP(), []int{0}
}

// **GcpArtifactRegistrySpec** defines the configuration for deploying a Google Cloud Artifact Registry.
// This message specifies the necessary parameters to create and manage an Artifact Registry within a
// specified GCP project and region. It allows you to set the project ID and region, and configure
// access settings such as enabling unauthenticated external access, which is particularly useful for
// open-source projects that require public availability of their artifacts.
type GcpArtifactRegistryRepoSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The format of the repository in the Artifact Registry.
	RepoFormat GcpArtifactRegistryRepoFormat `protobuf:"varint,1,opt,name=repo_format,json=repoFormat,proto3,enum=project.planton.provider.gcp.gcpartifactregistryrepo.v1.GcpArtifactRegistryRepoFormat" json:"repo_format,omitempty"`
	// The ID of the GCP project where the Artifact Registry resources will be created.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// The GCP region where the Artifact Registry will be created (e.g., "us-west2").
	// Selecting a region close to your Kubernetes clusters can reduce service startup time
	// by enabling faster downloads of container images.
	Region string `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	// A flag indicating whether to allow unauthenticated access to artifacts published in the repositories.
	// Enable this for publishing artifacts for open-source projects that require public access.
	EnablePublicAccess bool `protobuf:"varint,4,opt,name=enable_public_access,json=enablePublicAccess,proto3" json:"enable_public_access,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *GcpArtifactRegistryRepoSpec) Reset() {
	*x = GcpArtifactRegistryRepoSpec{}
	mi := &file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GcpArtifactRegistryRepoSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpArtifactRegistryRepoSpec) ProtoMessage() {}

func (x *GcpArtifactRegistryRepoSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpArtifactRegistryRepoSpec.ProtoReflect.Descriptor instead.
func (*GcpArtifactRegistryRepoSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *GcpArtifactRegistryRepoSpec) GetRepoFormat() GcpArtifactRegistryRepoFormat {
	if x != nil {
		return x.RepoFormat
	}
	return GcpArtifactRegistryRepoFormat_gcp_artifact_registry_repo_format_unspecified
}

func (x *GcpArtifactRegistryRepoSpec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GcpArtifactRegistryRepoSpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *GcpArtifactRegistryRepoSpec) GetEnablePublicAccess() bool {
	if x != nil {
		return x.EnablePublicAccess
	}
	return false
}

var File_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto protoreflect.FileDescriptor

const file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDesc = "" +
	"\n" +
	"Bproject/planton/provider/gcp/gcpartifactregistryrepo/v1/spec.proto\x127project.planton.provider.gcp.gcpartifactregistryrepo.v1\x1a\x1bbuf/validate/validate.proto\"\x97\x02\n" +
	"\x1bGcpArtifactRegistryRepoSpec\x12\x7f\n" +
	"\vrepo_format\x18\x01 \x01(\x0e2V.project.planton.provider.gcp.gcpartifactregistryrepo.v1.GcpArtifactRegistryRepoFormatB\x06\xbaH\x03\xc8\x01\x01R\n" +
	"repoFormat\x12%\n" +
	"\n" +
	"project_id\x18\x02 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\tprojectId\x12\x1e\n" +
	"\x06region\x18\x03 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\x06region\x120\n" +
	"\x14enable_public_access\x18\x04 \x01(\bR\x12enablePublicAccess*\xaa\x01\n" +
	"\x1dGcpArtifactRegistryRepoFormat\x121\n" +
	"-gcp_artifact_registry_repo_format_unspecified\x10\x00\x12\n" +
	"\n" +
	"\x06DOCKER\x10\x01\x12\v\n" +
	"\aGENERIC\x10\x02\x12\x06\n" +
	"\x02GO\x10\x03\x12\f\n" +
	"\bKUBEFLOW\x10\x04\x12\t\n" +
	"\x05MAVEN\x10\x05\x12\a\n" +
	"\x03NPM\x10\x06\x12\n" +
	"\n" +
	"\x06PYTHON\x10\a\x12\a\n" +
	"\x03YUM\x10\bB\xcf\x03\n" +
	";com.project.planton.provider.gcp.gcpartifactregistryrepo.v1B\tSpecProtoP\x01Z\x81\x01github.com/project-planton/project-planton/apis/project/planton/provider/gcp/gcpartifactregistryrepo/v1;gcpartifactregistryrepov1\xa2\x02\x05PPPGG\xaa\x027Project.Planton.Provider.Gcp.Gcpartifactregistryrepo.V1\xca\x027Project\\Planton\\Provider\\Gcp\\Gcpartifactregistryrepo\\V1\xe2\x02CProject\\Planton\\Provider\\Gcp\\Gcpartifactregistryrepo\\V1\\GPBMetadata\xea\x02<Project::Planton::Provider::Gcp::Gcpartifactregistryrepo::V1b\x06proto3"

var (
	file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescData []byte
)

func file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDesc), len(file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDesc)))
	})
	return file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDescData
}

var file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_goTypes = []any{
	(GcpArtifactRegistryRepoFormat)(0),  // 0: project.planton.provider.gcp.gcpartifactregistryrepo.v1.GcpArtifactRegistryRepoFormat
	(*GcpArtifactRegistryRepoSpec)(nil), // 1: project.planton.provider.gcp.gcpartifactregistryrepo.v1.GcpArtifactRegistryRepoSpec
}
var file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_depIdxs = []int32{
	0, // 0: project.planton.provider.gcp.gcpartifactregistryrepo.v1.GcpArtifactRegistryRepoSpec.repo_format:type_name -> project.planton.provider.gcp.gcpartifactregistryrepo.v1.GcpArtifactRegistryRepoFormat
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_init() }
func file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_init() {
	if File_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDesc), len(file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_depIdxs,
		EnumInfos:         file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_enumTypes,
		MessageInfos:      file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_msgTypes,
	}.Build()
	File_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto = out.File
	file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_goTypes = nil
	file_project_planton_provider_gcp_gcpartifactregistryrepo_v1_spec_proto_depIdxs = nil
}
