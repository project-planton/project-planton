// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/workload/signozkubernetes/v1/api.proto

package signozkubernetesv1

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

// signoz-kubernetes
type SignozKubernetes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	Metadata *shared.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *SignozKubernetesSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status        *SignozKubernetesStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignozKubernetes) Reset() {
	*x = SignozKubernetes{}
	mi := &file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignozKubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignozKubernetes) ProtoMessage() {}

func (x *SignozKubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignozKubernetes.ProtoReflect.Descriptor instead.
func (*SignozKubernetes) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *SignozKubernetes) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *SignozKubernetes) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *SignozKubernetes) GetMetadata() *shared.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SignozKubernetes) GetSpec() *SignozKubernetesSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *SignozKubernetes) GetStatus() *SignozKubernetesStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// signoz-kubernetes status.
type SignozKubernetesStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lifecycle
	Lifecycle *shared.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *shared.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack-job id
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	Outputs       *SignozKubernetesStackOutputs `protobuf:"bytes,1,opt,name=outputs,proto3" json:"outputs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignozKubernetesStatus) Reset() {
	*x = SignozKubernetesStatus{}
	mi := &file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignozKubernetesStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignozKubernetesStatus) ProtoMessage() {}

func (x *SignozKubernetesStatus) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignozKubernetesStatus.ProtoReflect.Descriptor instead.
func (*SignozKubernetesStatus) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *SignozKubernetesStatus) GetLifecycle() *shared.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *SignozKubernetesStatus) GetAudit() *shared.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *SignozKubernetesStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *SignozKubernetesStatus) GetOutputs() *SignozKubernetesStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto protoreflect.FileDescriptor

const file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDesc = "" +
	"\n" +
	"Jproject/planton/provider/kubernetes/workload/signozkubernetes/v1/api.proto\x12@project.planton.provider.kubernetes.workload.signozkubernetes.v1\x1a\x1bbuf/validate/validate.proto\x1aKproject/planton/provider/kubernetes/workload/signozkubernetes/v1/spec.proto\x1aTproject/planton/provider/kubernetes/workload/signozkubernetes/v1/stack_outputs.proto\x1a#project/planton/shared/status.proto\x1a%project/planton/shared/metadata.proto\"\xc1\x03\n" +
	"\x10SignozKubernetes\x12I\n" +
	"\vapi_version\x18\x01 \x01(\tB(\xbaH%r#\n" +
	"!kubernetes.project-planton.org/v1R\n" +
	"apiVersion\x12+\n" +
	"\x04kind\x18\x02 \x01(\tB\x17\xbaH\x14r\x12\n" +
	"\x10SignozKubernetesR\x04kind\x12O\n" +
	"\bmetadata\x18\x03 \x01(\v2+.project.planton.shared.ApiResourceMetadataB\x06\xbaH\x03\xc8\x01\x01R\bmetadata\x12r\n" +
	"\x04spec\x18\x04 \x01(\v2V.project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesSpecB\x06\xbaH\x03\xc8\x01\x01R\x04spec\x12p\n" +
	"\x06status\x18\x05 \x01(\v2X.project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStatusR\x06status\"\xc0\x02\n" +
	"\x16SignozKubernetesStatus\x12J\n" +
	"\tlifecycle\x18c \x01(\v2,.project.planton.shared.ApiResourceLifecycleR\tlifecycle\x12>\n" +
	"\x05audit\x18b \x01(\v2(.project.planton.shared.ApiResourceAuditR\x05audit\x12 \n" +
	"\fstack_job_id\x18a \x01(\tR\n" +
	"stackJobId\x12x\n" +
	"\aoutputs\x18\x01 \x01(\v2^.project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStackOutputsR\aoutputsB\xff\x03\n" +
	"Dcom.project.planton.provider.kubernetes.workload.signozkubernetes.v1B\bApiProtoP\x01Z\x83\x01github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/signozkubernetes/v1;signozkubernetesv1\xa2\x02\x06PPPKWS\xaa\x02@Project.Planton.Provider.Kubernetes.Workload.Signozkubernetes.V1\xca\x02@Project\\Planton\\Provider\\Kubernetes\\Workload\\Signozkubernetes\\V1\xe2\x02LProject\\Planton\\Provider\\Kubernetes\\Workload\\Signozkubernetes\\V1\\GPBMetadata\xea\x02FProject::Planton::Provider::Kubernetes::Workload::Signozkubernetes::V1b\x06proto3"

var (
	file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescOnce sync.Once
	file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescData []byte
)

func file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescGZIP() []byte {
	file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDesc), len(file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDesc)))
	})
	return file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDescData
}

var file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_goTypes = []any{
	(*SignozKubernetes)(nil),             // 0: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetes
	(*SignozKubernetesStatus)(nil),       // 1: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStatus
	(*shared.ApiResourceMetadata)(nil),   // 2: project.planton.shared.ApiResourceMetadata
	(*SignozKubernetesSpec)(nil),         // 3: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesSpec
	(*shared.ApiResourceLifecycle)(nil),  // 4: project.planton.shared.ApiResourceLifecycle
	(*shared.ApiResourceAudit)(nil),      // 5: project.planton.shared.ApiResourceAudit
	(*SignozKubernetesStackOutputs)(nil), // 6: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStackOutputs
}
var file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetes.metadata:type_name -> project.planton.shared.ApiResourceMetadata
	3, // 1: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetes.spec:type_name -> project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesSpec
	1, // 2: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetes.status:type_name -> project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStatus
	4, // 3: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStatus.lifecycle:type_name -> project.planton.shared.ApiResourceLifecycle
	5, // 4: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStatus.audit:type_name -> project.planton.shared.ApiResourceAudit
	6, // 5: project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStatus.outputs:type_name -> project.planton.provider.kubernetes.workload.signozkubernetes.v1.SignozKubernetesStackOutputs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_init() }
func file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_init() {
	if File_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto != nil {
		return
	}
	file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_spec_proto_init()
	file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_stack_outputs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDesc), len(file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_msgTypes,
	}.Build()
	File_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto = out.File
	file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_goTypes = nil
	file_project_planton_provider_kubernetes_workload_signozkubernetes_v1_api_proto_depIdxs = nil
}
