// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/credential/kubernetesclustercredential/v1/api.proto

package kubernetesclustercredentialv1

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

// kubernetes-cluster-credential
type KubernetesClusterCredential struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	Metadata *shared.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *KubernetesClusterCredentialSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status        *shared.ApiResourceLifecycleAndAuditStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KubernetesClusterCredential) Reset() {
	*x = KubernetesClusterCredential{}
	mi := &file_project_planton_credential_kubernetesclustercredential_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubernetesClusterCredential) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubernetesClusterCredential) ProtoMessage() {}

func (x *KubernetesClusterCredential) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_credential_kubernetesclustercredential_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubernetesClusterCredential.ProtoReflect.Descriptor instead.
func (*KubernetesClusterCredential) Descriptor() ([]byte, []int) {
	return file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *KubernetesClusterCredential) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *KubernetesClusterCredential) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *KubernetesClusterCredential) GetMetadata() *shared.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *KubernetesClusterCredential) GetSpec() *KubernetesClusterCredentialSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *KubernetesClusterCredential) GetStatus() *shared.ApiResourceLifecycleAndAuditStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

var File_project_planton_credential_kubernetesclustercredential_v1_api_proto protoreflect.FileDescriptor

const file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDesc = "" +
	"\n" +
	"Cproject/planton/credential/kubernetesclustercredential/v1/api.proto\x129project.planton.credential.kubernetesclustercredential.v1\x1a\x1bbuf/validate/validate.proto\x1a#project/planton/shared/status.proto\x1a%project/planton/shared/metadata.proto\x1aDproject/planton/credential/kubernetesclustercredential/v1/spec.proto\"\xb5\x03\n" +
	"\x1bKubernetesClusterCredential\x12I\n" +
	"\vapi_version\x18\x01 \x01(\tB(\xbaH%r#\n" +
	"!credential.project-planton.org/v1R\n" +
	"apiVersion\x126\n" +
	"\x04kind\x18\x02 \x01(\tB\"\xbaH\x1fr\x1d\n" +
	"\x1bKubernetesClusterCredentialR\x04kind\x12O\n" +
	"\bmetadata\x18\x03 \x01(\v2+.project.planton.shared.ApiResourceMetadataB\x06\xbaH\x03\xc8\x01\x01R\bmetadata\x12n\n" +
	"\x04spec\x18\x04 \x01(\v2Z.project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredentialSpecR\x04spec\x12R\n" +
	"\x06status\x18\x05 \x01(\v2:.project.planton.shared.ApiResourceLifecycleAndAuditStatusR\x06statusB\xdc\x03\n" +
	"=com.project.planton.credential.kubernetesclustercredential.v1B\bApiProtoP\x01Z\x87\x01github.com/project-planton/project-planton/apis/project/planton/credential/kubernetesclustercredential/v1;kubernetesclustercredentialv1\xa2\x02\x04PPCK\xaa\x029Project.Planton.Credential.Kubernetesclustercredential.V1\xca\x029Project\\Planton\\Credential\\Kubernetesclustercredential\\V1\xe2\x02EProject\\Planton\\Credential\\Kubernetesclustercredential\\V1\\GPBMetadata\xea\x02=Project::Planton::Credential::Kubernetesclustercredential::V1b\x06proto3"

var (
	file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDescOnce sync.Once
	file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDescData []byte
)

func file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDescGZIP() []byte {
	file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDescOnce.Do(func() {
		file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDesc), len(file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDesc)))
	})
	return file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDescData
}

var file_project_planton_credential_kubernetesclustercredential_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_credential_kubernetesclustercredential_v1_api_proto_goTypes = []any{
	(*KubernetesClusterCredential)(nil),               // 0: project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredential
	(*shared.ApiResourceMetadata)(nil),                // 1: project.planton.shared.ApiResourceMetadata
	(*KubernetesClusterCredentialSpec)(nil),           // 2: project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredentialSpec
	(*shared.ApiResourceLifecycleAndAuditStatus)(nil), // 3: project.planton.shared.ApiResourceLifecycleAndAuditStatus
}
var file_project_planton_credential_kubernetesclustercredential_v1_api_proto_depIdxs = []int32{
	1, // 0: project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredential.metadata:type_name -> project.planton.shared.ApiResourceMetadata
	2, // 1: project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredential.spec:type_name -> project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredentialSpec
	3, // 2: project.planton.credential.kubernetesclustercredential.v1.KubernetesClusterCredential.status:type_name -> project.planton.shared.ApiResourceLifecycleAndAuditStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_project_planton_credential_kubernetesclustercredential_v1_api_proto_init() }
func file_project_planton_credential_kubernetesclustercredential_v1_api_proto_init() {
	if File_project_planton_credential_kubernetesclustercredential_v1_api_proto != nil {
		return
	}
	file_project_planton_credential_kubernetesclustercredential_v1_spec_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDesc), len(file_project_planton_credential_kubernetesclustercredential_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_credential_kubernetesclustercredential_v1_api_proto_goTypes,
		DependencyIndexes: file_project_planton_credential_kubernetesclustercredential_v1_api_proto_depIdxs,
		MessageInfos:      file_project_planton_credential_kubernetesclustercredential_v1_api_proto_msgTypes,
	}.Build()
	File_project_planton_credential_kubernetesclustercredential_v1_api_proto = out.File
	file_project_planton_credential_kubernetesclustercredential_v1_api_proto_goTypes = nil
	file_project_planton_credential_kubernetesclustercredential_v1_api_proto_depIdxs = nil
}
