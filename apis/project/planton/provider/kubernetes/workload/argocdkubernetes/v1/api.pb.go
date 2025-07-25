// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/workload/argocdkubernetes/v1/api.proto

package argocdkubernetesv1

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

// argocd-kubernetes
type ArgocdKubernetes struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// api-version
	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// resource-kind
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// metadata
	Metadata *shared.ApiResourceMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// spec
	Spec *ArgocdKubernetesSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	// status
	Status        *ArgocdKubernetesStatus `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArgocdKubernetes) Reset() {
	*x = ArgocdKubernetes{}
	mi := &file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArgocdKubernetes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgocdKubernetes) ProtoMessage() {}

func (x *ArgocdKubernetes) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgocdKubernetes.ProtoReflect.Descriptor instead.
func (*ArgocdKubernetes) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *ArgocdKubernetes) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *ArgocdKubernetes) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *ArgocdKubernetes) GetMetadata() *shared.ApiResourceMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *ArgocdKubernetes) GetSpec() *ArgocdKubernetesSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *ArgocdKubernetes) GetStatus() *ArgocdKubernetesStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

// argocd-kubernetes status.
type ArgocdKubernetesStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// lifecycle
	Lifecycle *shared.ApiResourceLifecycle `protobuf:"bytes,99,opt,name=lifecycle,proto3" json:"lifecycle,omitempty"`
	// audit-info
	Audit *shared.ApiResourceAudit `protobuf:"bytes,98,opt,name=audit,proto3" json:"audit,omitempty"`
	// stack-job id
	StackJobId string `protobuf:"bytes,97,opt,name=stack_job_id,json=stackJobId,proto3" json:"stack_job_id,omitempty"`
	// stack-outputs
	// argocd-kubernetes stack-outputs
	Outputs       *ArgocdKubernetesStackOutputs `protobuf:"bytes,1,opt,name=outputs,proto3" json:"outputs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ArgocdKubernetesStatus) Reset() {
	*x = ArgocdKubernetesStatus{}
	mi := &file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ArgocdKubernetesStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArgocdKubernetesStatus) ProtoMessage() {}

func (x *ArgocdKubernetesStatus) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArgocdKubernetesStatus.ProtoReflect.Descriptor instead.
func (*ArgocdKubernetesStatus) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *ArgocdKubernetesStatus) GetLifecycle() *shared.ApiResourceLifecycle {
	if x != nil {
		return x.Lifecycle
	}
	return nil
}

func (x *ArgocdKubernetesStatus) GetAudit() *shared.ApiResourceAudit {
	if x != nil {
		return x.Audit
	}
	return nil
}

func (x *ArgocdKubernetesStatus) GetStackJobId() string {
	if x != nil {
		return x.StackJobId
	}
	return ""
}

func (x *ArgocdKubernetesStatus) GetOutputs() *ArgocdKubernetesStackOutputs {
	if x != nil {
		return x.Outputs
	}
	return nil
}

var File_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto protoreflect.FileDescriptor

const file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDesc = "" +
	"\n" +
	"Jproject/planton/provider/kubernetes/workload/argocdkubernetes/v1/api.proto\x12@project.planton.provider.kubernetes.workload.argocdkubernetes.v1\x1a\x1bbuf/validate/validate.proto\x1aKproject/planton/provider/kubernetes/workload/argocdkubernetes/v1/spec.proto\x1aTproject/planton/provider/kubernetes/workload/argocdkubernetes/v1/stack_outputs.proto\x1a#project/planton/shared/status.proto\x1a%project/planton/shared/metadata.proto\"\xc1\x03\n" +
	"\x10ArgocdKubernetes\x12I\n" +
	"\vapi_version\x18\x01 \x01(\tB(\xbaH%r#\n" +
	"!kubernetes.project-planton.org/v1R\n" +
	"apiVersion\x12+\n" +
	"\x04kind\x18\x02 \x01(\tB\x17\xbaH\x14r\x12\n" +
	"\x10ArgocdKubernetesR\x04kind\x12O\n" +
	"\bmetadata\x18\x03 \x01(\v2+.project.planton.shared.ApiResourceMetadataB\x06\xbaH\x03\xc8\x01\x01R\bmetadata\x12r\n" +
	"\x04spec\x18\x04 \x01(\v2V.project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesSpecB\x06\xbaH\x03\xc8\x01\x01R\x04spec\x12p\n" +
	"\x06status\x18\x05 \x01(\v2X.project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStatusR\x06status\"\xc0\x02\n" +
	"\x16ArgocdKubernetesStatus\x12J\n" +
	"\tlifecycle\x18c \x01(\v2,.project.planton.shared.ApiResourceLifecycleR\tlifecycle\x12>\n" +
	"\x05audit\x18b \x01(\v2(.project.planton.shared.ApiResourceAuditR\x05audit\x12 \n" +
	"\fstack_job_id\x18a \x01(\tR\n" +
	"stackJobId\x12x\n" +
	"\aoutputs\x18\x01 \x01(\v2^.project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStackOutputsR\aoutputsB\xff\x03\n" +
	"Dcom.project.planton.provider.kubernetes.workload.argocdkubernetes.v1B\bApiProtoP\x01Z\x83\x01github.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/argocdkubernetes/v1;argocdkubernetesv1\xa2\x02\x06PPPKWA\xaa\x02@Project.Planton.Provider.Kubernetes.Workload.Argocdkubernetes.V1\xca\x02@Project\\Planton\\Provider\\Kubernetes\\Workload\\Argocdkubernetes\\V1\xe2\x02LProject\\Planton\\Provider\\Kubernetes\\Workload\\Argocdkubernetes\\V1\\GPBMetadata\xea\x02FProject::Planton::Provider::Kubernetes::Workload::Argocdkubernetes::V1b\x06proto3"

var (
	file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescOnce sync.Once
	file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescData []byte
)

func file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescGZIP() []byte {
	file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDesc), len(file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDesc)))
	})
	return file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDescData
}

var file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_goTypes = []any{
	(*ArgocdKubernetes)(nil),             // 0: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetes
	(*ArgocdKubernetesStatus)(nil),       // 1: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStatus
	(*shared.ApiResourceMetadata)(nil),   // 2: project.planton.shared.ApiResourceMetadata
	(*ArgocdKubernetesSpec)(nil),         // 3: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesSpec
	(*shared.ApiResourceLifecycle)(nil),  // 4: project.planton.shared.ApiResourceLifecycle
	(*shared.ApiResourceAudit)(nil),      // 5: project.planton.shared.ApiResourceAudit
	(*ArgocdKubernetesStackOutputs)(nil), // 6: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStackOutputs
}
var file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetes.metadata:type_name -> project.planton.shared.ApiResourceMetadata
	3, // 1: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetes.spec:type_name -> project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesSpec
	1, // 2: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetes.status:type_name -> project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStatus
	4, // 3: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStatus.lifecycle:type_name -> project.planton.shared.ApiResourceLifecycle
	5, // 4: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStatus.audit:type_name -> project.planton.shared.ApiResourceAudit
	6, // 5: project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStatus.outputs:type_name -> project.planton.provider.kubernetes.workload.argocdkubernetes.v1.ArgocdKubernetesStackOutputs
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_init() }
func file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_init() {
	if File_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto != nil {
		return
	}
	file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_spec_proto_init()
	file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_stack_outputs_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDesc), len(file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_msgTypes,
	}.Build()
	File_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto = out.File
	file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_goTypes = nil
	file_project_planton_provider_kubernetes_workload_argocdkubernetes_v1_api_proto_depIdxs = nil
}
