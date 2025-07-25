// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/shared/iac.proto

package shared

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

type IacProvisioner int32

const (
	IacProvisioner_iac_provisioner_unspecified IacProvisioner = 0
	IacProvisioner_terraform                   IacProvisioner = 1
	IacProvisioner_pulumi                      IacProvisioner = 2
)

// Enum value maps for IacProvisioner.
var (
	IacProvisioner_name = map[int32]string{
		0: "iac_provisioner_unspecified",
		1: "terraform",
		2: "pulumi",
	}
	IacProvisioner_value = map[string]int32{
		"iac_provisioner_unspecified": 0,
		"terraform":                   1,
		"pulumi":                      2,
	}
)

func (x IacProvisioner) Enum() *IacProvisioner {
	p := new(IacProvisioner)
	*p = x
	return p
}

func (x IacProvisioner) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IacProvisioner) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_shared_iac_proto_enumTypes[0].Descriptor()
}

func (IacProvisioner) Type() protoreflect.EnumType {
	return &file_project_planton_shared_iac_proto_enumTypes[0]
}

func (x IacProvisioner) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IacProvisioner.Descriptor instead.
func (IacProvisioner) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_shared_iac_proto_rawDescGZIP(), []int{0}
}

var File_project_planton_shared_iac_proto protoreflect.FileDescriptor

const file_project_planton_shared_iac_proto_rawDesc = "" +
	"\n" +
	" project/planton/shared/iac.proto\x12\x16project.planton.shared*L\n" +
	"\x0eIacProvisioner\x12\x1f\n" +
	"\x1biac_provisioner_unspecified\x10\x00\x12\r\n" +
	"\tterraform\x10\x01\x12\n" +
	"\n" +
	"\x06pulumi\x10\x02B\xe8\x01\n" +
	"\x1acom.project.planton.sharedB\bIacProtoP\x01ZFgithub.com/project-planton/project-planton/apis/project/planton/shared\xa2\x02\x03PPS\xaa\x02\x16Project.Planton.Shared\xca\x02\x16Project\\Planton\\Shared\xe2\x02\"Project\\Planton\\Shared\\GPBMetadata\xea\x02\x18Project::Planton::Sharedb\x06proto3"

var (
	file_project_planton_shared_iac_proto_rawDescOnce sync.Once
	file_project_planton_shared_iac_proto_rawDescData []byte
)

func file_project_planton_shared_iac_proto_rawDescGZIP() []byte {
	file_project_planton_shared_iac_proto_rawDescOnce.Do(func() {
		file_project_planton_shared_iac_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_shared_iac_proto_rawDesc), len(file_project_planton_shared_iac_proto_rawDesc)))
	})
	return file_project_planton_shared_iac_proto_rawDescData
}

var file_project_planton_shared_iac_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_project_planton_shared_iac_proto_goTypes = []any{
	(IacProvisioner)(0), // 0: project.planton.shared.IacProvisioner
}
var file_project_planton_shared_iac_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_shared_iac_proto_init() }
func file_project_planton_shared_iac_proto_init() {
	if File_project_planton_shared_iac_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_shared_iac_proto_rawDesc), len(file_project_planton_shared_iac_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_shared_iac_proto_goTypes,
		DependencyIndexes: file_project_planton_shared_iac_proto_depIdxs,
		EnumInfos:         file_project_planton_shared_iac_proto_enumTypes,
	}.Build()
	File_project_planton_shared_iac_proto = out.File
	file_project_planton_shared_iac_proto_goTypes = nil
	file_project_planton_shared_iac_proto_depIdxs = nil
}
