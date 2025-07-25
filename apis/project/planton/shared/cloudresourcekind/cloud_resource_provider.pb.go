// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/shared/cloudresourcekind/cloud_resource_provider.proto

package cloudresourcekind

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

type ProjectPlantonCloudResourceProvider int32

const (
	ProjectPlantonCloudResourceProvider_project_planton_cloud_resource_provider_unspecified ProjectPlantonCloudResourceProvider = 0
	ProjectPlantonCloudResourceProvider_test                                                ProjectPlantonCloudResourceProvider = 1
	ProjectPlantonCloudResourceProvider_atlas                                               ProjectPlantonCloudResourceProvider = 2
	ProjectPlantonCloudResourceProvider_aws                                                 ProjectPlantonCloudResourceProvider = 3
	ProjectPlantonCloudResourceProvider_azure                                               ProjectPlantonCloudResourceProvider = 4
	ProjectPlantonCloudResourceProvider_confluent                                           ProjectPlantonCloudResourceProvider = 5
	ProjectPlantonCloudResourceProvider_digital_ocean                                       ProjectPlantonCloudResourceProvider = 6
	ProjectPlantonCloudResourceProvider_gcp                                                 ProjectPlantonCloudResourceProvider = 7
	ProjectPlantonCloudResourceProvider_kubernetes                                          ProjectPlantonCloudResourceProvider = 8
	ProjectPlantonCloudResourceProvider_snowflake                                           ProjectPlantonCloudResourceProvider = 9
)

// Enum value maps for ProjectPlantonCloudResourceProvider.
var (
	ProjectPlantonCloudResourceProvider_name = map[int32]string{
		0: "project_planton_cloud_resource_provider_unspecified",
		1: "test",
		2: "atlas",
		3: "aws",
		4: "azure",
		5: "confluent",
		6: "digital_ocean",
		7: "gcp",
		8: "kubernetes",
		9: "snowflake",
	}
	ProjectPlantonCloudResourceProvider_value = map[string]int32{
		"project_planton_cloud_resource_provider_unspecified": 0,
		"test":          1,
		"atlas":         2,
		"aws":           3,
		"azure":         4,
		"confluent":     5,
		"digital_ocean": 6,
		"gcp":           7,
		"kubernetes":    8,
		"snowflake":     9,
	}
)

func (x ProjectPlantonCloudResourceProvider) Enum() *ProjectPlantonCloudResourceProvider {
	p := new(ProjectPlantonCloudResourceProvider)
	*p = x
	return p
}

func (x ProjectPlantonCloudResourceProvider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectPlantonCloudResourceProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_enumTypes[0].Descriptor()
}

func (ProjectPlantonCloudResourceProvider) Type() protoreflect.EnumType {
	return &file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_enumTypes[0]
}

func (x ProjectPlantonCloudResourceProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectPlantonCloudResourceProvider.Descriptor instead.
func (ProjectPlantonCloudResourceProvider) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescGZIP(), []int{0}
}

type ProjectPlantonKubernetesResourceType int32

const (
	ProjectPlantonKubernetesResourceType_project_planton_kubernetes_resource_type_unspecified ProjectPlantonKubernetesResourceType = 0
	ProjectPlantonKubernetesResourceType_addon                                                ProjectPlantonKubernetesResourceType = 1
	ProjectPlantonKubernetesResourceType_workload                                             ProjectPlantonKubernetesResourceType = 2
)

// Enum value maps for ProjectPlantonKubernetesResourceType.
var (
	ProjectPlantonKubernetesResourceType_name = map[int32]string{
		0: "project_planton_kubernetes_resource_type_unspecified",
		1: "addon",
		2: "workload",
	}
	ProjectPlantonKubernetesResourceType_value = map[string]int32{
		"project_planton_kubernetes_resource_type_unspecified": 0,
		"addon":    1,
		"workload": 2,
	}
)

func (x ProjectPlantonKubernetesResourceType) Enum() *ProjectPlantonKubernetesResourceType {
	p := new(ProjectPlantonKubernetesResourceType)
	*p = x
	return p
}

func (x ProjectPlantonKubernetesResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ProjectPlantonKubernetesResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_enumTypes[1].Descriptor()
}

func (ProjectPlantonKubernetesResourceType) Type() protoreflect.EnumType {
	return &file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_enumTypes[1]
}

func (x ProjectPlantonKubernetesResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ProjectPlantonKubernetesResourceType.Descriptor instead.
func (ProjectPlantonKubernetesResourceType) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescGZIP(), []int{1}
}

var File_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto protoreflect.FileDescriptor

const file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDesc = "" +
	"\n" +
	"Fproject/planton/shared/cloudresourcekind/cloud_resource_provider.proto\x12(project.planton.shared.cloudresourcekind*\xd1\x01\n" +
	"#ProjectPlantonCloudResourceProvider\x127\n" +
	"3project_planton_cloud_resource_provider_unspecified\x10\x00\x12\b\n" +
	"\x04test\x10\x01\x12\t\n" +
	"\x05atlas\x10\x02\x12\a\n" +
	"\x03aws\x10\x03\x12\t\n" +
	"\x05azure\x10\x04\x12\r\n" +
	"\tconfluent\x10\x05\x12\x11\n" +
	"\rdigital_ocean\x10\x06\x12\a\n" +
	"\x03gcp\x10\a\x12\x0e\n" +
	"\n" +
	"kubernetes\x10\b\x12\r\n" +
	"\tsnowflake\x10\t*y\n" +
	"$ProjectPlantonKubernetesResourceType\x128\n" +
	"4project_planton_kubernetes_resource_type_unspecified\x10\x00\x12\t\n" +
	"\x05addon\x10\x01\x12\f\n" +
	"\bworkload\x10\x02B\xe8\x02\n" +
	",com.project.planton.shared.cloudresourcekindB\x1aCloudResourceProviderProtoP\x01ZXgithub.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind\xa2\x02\x04PPSC\xaa\x02(Project.Planton.Shared.Cloudresourcekind\xca\x02(Project\\Planton\\Shared\\Cloudresourcekind\xe2\x024Project\\Planton\\Shared\\Cloudresourcekind\\GPBMetadata\xea\x02+Project::Planton::Shared::Cloudresourcekindb\x06proto3"

var (
	file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescOnce sync.Once
	file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescData []byte
)

func file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescGZIP() []byte {
	file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescOnce.Do(func() {
		file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDesc), len(file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDesc)))
	})
	return file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDescData
}

var file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_goTypes = []any{
	(ProjectPlantonCloudResourceProvider)(0),  // 0: project.planton.shared.cloudresourcekind.ProjectPlantonCloudResourceProvider
	(ProjectPlantonKubernetesResourceType)(0), // 1: project.planton.shared.cloudresourcekind.ProjectPlantonKubernetesResourceType
}
var file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_init() }
func file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_init() {
	if File_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDesc), len(file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_goTypes,
		DependencyIndexes: file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_depIdxs,
		EnumInfos:         file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_enumTypes,
	}.Build()
	File_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto = out.File
	file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_goTypes = nil
	file_project_planton_shared_cloudresourcekind_cloud_resource_provider_proto_depIdxs = nil
}
