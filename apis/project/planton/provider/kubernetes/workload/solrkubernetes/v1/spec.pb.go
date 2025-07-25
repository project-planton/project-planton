// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/workload/solrkubernetes/v1/spec.proto

package solrkubernetesv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	kubernetes "github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	_ "github.com/project-planton/project-planton/apis/project/planton/shared/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
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

// *
// **SolrKubernetesSpec** defines the configuration for deploying Apache Solr on a Kubernetes cluster.
// This message includes specifications for the Solr container, Zookeeper container, and ingress settings.
// By configuring these parameters, you can set up a Solr deployment tailored to your application's needs,
// including resource allocation, data persistence, and external access through ingress.
type SolrKubernetesSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The specifications for the Solr container deployment.
	SolrContainer *SolrKubernetesSolrContainer `protobuf:"bytes,1,opt,name=solr_container,json=solrContainer,proto3" json:"solr_container,omitempty"`
	// The Solr-specific configuration options.
	Config *SolrKubernetesSolrConfig `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
	// The specifications for the Zookeeper container deployment.
	ZookeeperContainer *SolrKubernetesZookeeperContainer `protobuf:"bytes,3,opt,name=zookeeper_container,json=zookeeperContainer,proto3" json:"zookeeper_container,omitempty"`
	// The ingress configuration for the Solr deployment.
	Ingress       *kubernetes.IngressSpec `protobuf:"bytes,4,opt,name=ingress,proto3" json:"ingress,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SolrKubernetesSpec) Reset() {
	*x = SolrKubernetesSpec{}
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolrKubernetesSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolrKubernetesSpec) ProtoMessage() {}

func (x *SolrKubernetesSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolrKubernetesSpec.ProtoReflect.Descriptor instead.
func (*SolrKubernetesSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *SolrKubernetesSpec) GetSolrContainer() *SolrKubernetesSolrContainer {
	if x != nil {
		return x.SolrContainer
	}
	return nil
}

func (x *SolrKubernetesSpec) GetConfig() *SolrKubernetesSolrConfig {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *SolrKubernetesSpec) GetZookeeperContainer() *SolrKubernetesZookeeperContainer {
	if x != nil {
		return x.ZookeeperContainer
	}
	return nil
}

func (x *SolrKubernetesSpec) GetIngress() *kubernetes.IngressSpec {
	if x != nil {
		return x.Ingress
	}
	return nil
}

// *
// **SolrKubernetesSolrContainer** specifies the configuration for the Solr container.
// It includes settings such as the number of replicas, container image, resource allocations,
// disk size for data persistence, and Solr-specific configurations.
// Proper configuration ensures optimal performance and data reliability for your Solr deployment.
type SolrKubernetesSolrContainer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The number of Solr pods in the Solr Kubernetes deployment.
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// The CPU and memory resources allocated to the Solr container.
	Resources *kubernetes.ContainerResources `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	// The size of the persistent volume attached to each Solr pod (e.g., "1Gi").
	DiskSize string `protobuf:"bytes,3,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	// The container image for the Solr deployment.
	// Example repository: "solr", example tag: "8.7.0".
	Image         *kubernetes.ContainerImage `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SolrKubernetesSolrContainer) Reset() {
	*x = SolrKubernetesSolrContainer{}
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolrKubernetesSolrContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolrKubernetesSolrContainer) ProtoMessage() {}

func (x *SolrKubernetesSolrContainer) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolrKubernetesSolrContainer.ProtoReflect.Descriptor instead.
func (*SolrKubernetesSolrContainer) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescGZIP(), []int{1}
}

func (x *SolrKubernetesSolrContainer) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *SolrKubernetesSolrContainer) GetResources() *kubernetes.ContainerResources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *SolrKubernetesSolrContainer) GetDiskSize() string {
	if x != nil {
		return x.DiskSize
	}
	return ""
}

func (x *SolrKubernetesSolrContainer) GetImage() *kubernetes.ContainerImage {
	if x != nil {
		return x.Image
	}
	return nil
}

// *
// **SolrKubernetesSolrConfig** specifies the configuration settings for Solr.
// It includes JVM memory settings, custom Solr options, and garbage collection tuning parameters.
type SolrKubernetesSolrConfig struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// JVM memory settings for Solr.
	JavaMem string `protobuf:"bytes,1,opt,name=java_mem,json=javaMem,proto3" json:"java_mem,omitempty"`
	// Custom Solr options (e.g., "-Dsolr.autoSoftCommit.maxTime=10000").
	Opts string `protobuf:"bytes,2,opt,name=opts,proto3" json:"opts,omitempty"`
	// Solr garbage collection tuning configuration (e.g., "-XX:SurvivorRatio=4 -XX:TargetSurvivorRatio=90 -XX:MaxTenuringThreshold=8").
	GarbageCollectionTuning string `protobuf:"bytes,3,opt,name=garbage_collection_tuning,json=garbageCollectionTuning,proto3" json:"garbage_collection_tuning,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *SolrKubernetesSolrConfig) Reset() {
	*x = SolrKubernetesSolrConfig{}
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolrKubernetesSolrConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolrKubernetesSolrConfig) ProtoMessage() {}

func (x *SolrKubernetesSolrConfig) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolrKubernetesSolrConfig.ProtoReflect.Descriptor instead.
func (*SolrKubernetesSolrConfig) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescGZIP(), []int{2}
}

func (x *SolrKubernetesSolrConfig) GetJavaMem() string {
	if x != nil {
		return x.JavaMem
	}
	return ""
}

func (x *SolrKubernetesSolrConfig) GetOpts() string {
	if x != nil {
		return x.Opts
	}
	return ""
}

func (x *SolrKubernetesSolrConfig) GetGarbageCollectionTuning() string {
	if x != nil {
		return x.GarbageCollectionTuning
	}
	return ""
}

// *
// **SolrKubernetesZookeeperContainer** specifies the configuration for the Zookeeper container used by Solr.
// It includes settings such as the number of replicas, resource allocations, and disk size for data persistence.
// Proper configuration ensures high availability and reliability for your Solr cluster.
type SolrKubernetesZookeeperContainer struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The number of Zookeeper pods in the Zookeeper cluster.
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// The CPU and memory resources allocated to the Zookeeper container.
	Resources *kubernetes.ContainerResources `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
	// The size of the persistent volume attached to each Zookeeper pod (e.g., "1Gi").
	DiskSize      string `protobuf:"bytes,3,opt,name=disk_size,json=diskSize,proto3" json:"disk_size,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SolrKubernetesZookeeperContainer) Reset() {
	*x = SolrKubernetesZookeeperContainer{}
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolrKubernetesZookeeperContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolrKubernetesZookeeperContainer) ProtoMessage() {}

func (x *SolrKubernetesZookeeperContainer) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolrKubernetesZookeeperContainer.ProtoReflect.Descriptor instead.
func (*SolrKubernetesZookeeperContainer) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescGZIP(), []int{3}
}

func (x *SolrKubernetesZookeeperContainer) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *SolrKubernetesZookeeperContainer) GetResources() *kubernetes.ContainerResources {
	if x != nil {
		return x.Resources
	}
	return nil
}

func (x *SolrKubernetesZookeeperContainer) GetDiskSize() string {
	if x != nil {
		return x.DiskSize
	}
	return ""
}

var file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*SolrKubernetesSolrContainer)(nil),
		Field:         540001,
		Name:          "project.planton.provider.kubernetes.workload.solrkubernetes.v1.default_solr_container",
		Tag:           "bytes,540001,opt,name=default_solr_container",
		Filename:      "project/planton/provider/kubernetes/workload/solrkubernetes/v1/spec.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*SolrKubernetesSolrContainer)(nil),
		Field:         540002,
		Name:          "project.planton.provider.kubernetes.workload.solrkubernetes.v1.default_zookeeper_container",
		Tag:           "bytes,540002,opt,name=default_zookeeper_container",
		Filename:      "project/planton/provider/kubernetes/workload/solrkubernetes/v1/spec.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer default_solr_container = 540001;
	E_DefaultSolrContainer = &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_extTypes[0]
	// optional project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer default_zookeeper_container = 540002;
	E_DefaultZookeeperContainer = &file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_extTypes[1]
)

var File_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto protoreflect.FileDescriptor

const file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDesc = "" +
	"\n" +
	"Iproject/planton/provider/kubernetes/workload/solrkubernetes/v1/spec.proto\x12>project.planton.provider.kubernetes.workload.solrkubernetes.v1\x1a\x1bbuf/validate/validate.proto\x1a2project/planton/shared/kubernetes/kubernetes.proto\x1a/project/planton/shared/kubernetes/options.proto\x1a,project/planton/shared/options/options.proto\x1a google/protobuf/descriptor.proto\"\xd0\x04\n" +
	"\x12SolrKubernetesSpec\x12\xbd\x01\n" +
	"\x0esolr_container\x18\x01 \x01(\v2[.project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainerB9\x8aև\x024\b\x01\x12\x1c\n" +
	"\f\n" +
	"\x051000m\x12\x031Gi\x12\f\n" +
	"\x0350m\x12\x05100Mi\x1a\x031Gi\"\r\n" +
	"\x04solr\x12\x058.7.0R\rsolrContainer\x12p\n" +
	"\x06config\x18\x02 \x01(\v2X.project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrConfigR\x06config\x12\xbd\x01\n" +
	"\x13zookeeper_container\x18\x03 \x01(\v2`.project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesZookeeperContainerB*\x92և\x02%\b\x01\x12\x1c\n" +
	"\f\n" +
	"\x051000m\x12\x031Gi\x12\f\n" +
	"\x0350m\x12\x05100Mi\x1a\x031GiR\x12zookeeperContainer\x12H\n" +
	"\aingress\x18\x04 \x01(\v2..project.planton.shared.kubernetes.IngressSpecR\aingress\"\x96\x03\n" +
	"\x1bSolrKubernetesSolrContainer\x12\x1a\n" +
	"\breplicas\x18\x01 \x01(\x05R\breplicas\x12S\n" +
	"\tresources\x18\x02 \x01(\v25.project.planton.shared.kubernetes.ContainerResourcesR\tresources\x12\xbc\x01\n" +
	"\tdisk_size\x18\x03 \x01(\tB\x9e\x01\xbaH\x9a\x01\xba\x01\x96\x01\n" +
	"!spec.container.disk_size.required\x12\x1aDisk size value is invalid\x1aUthis.matches('^\\\\d+(\\\\.\\\\d+)?\\\\s?(Ki|Mi|Gi|Ti|Pi|Ei|K|M|G|T|P|E)$') && size(this) > 0R\bdiskSize\x12G\n" +
	"\x05image\x18\x04 \x01(\v21.project.planton.shared.kubernetes.ContainerImageR\x05image\"\x85\x01\n" +
	"\x18SolrKubernetesSolrConfig\x12\x19\n" +
	"\bjava_mem\x18\x01 \x01(\tR\ajavaMem\x12\x12\n" +
	"\x04opts\x18\x02 \x01(\tR\x04opts\x12:\n" +
	"\x19garbage_collection_tuning\x18\x03 \x01(\tR\x17garbageCollectionTuning\"\xd2\x02\n" +
	" SolrKubernetesZookeeperContainer\x12\x1a\n" +
	"\breplicas\x18\x01 \x01(\x05R\breplicas\x12S\n" +
	"\tresources\x18\x02 \x01(\v25.project.planton.shared.kubernetes.ContainerResourcesR\tresources\x12\xbc\x01\n" +
	"\tdisk_size\x18\x03 \x01(\tB\x9e\x01\xbaH\x9a\x01\xba\x01\x96\x01\n" +
	"!spec.container.disk_size.required\x12\x1aDisk size value is invalid\x1aUthis.matches('^\\\\d+(\\\\.\\\\d+)?\\\\s?(Ki|Mi|Gi|Ti|Pi|Ei|K|M|G|T|P|E)$') && size(this) > 0R\bdiskSize:\xb2\x01\n" +
	"\x16default_solr_container\x12\x1d.google.protobuf.FieldOptions\x18\xe1\xfa  \x01(\v2[.project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainerR\x14defaultSolrContainer:\xbc\x01\n" +
	"\x1bdefault_zookeeper_container\x12\x1d.google.protobuf.FieldOptions\x18\xe2\xfa  \x01(\v2[.project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainerR\x19defaultZookeeperContainerB\xf1\x03\n" +
	"Bcom.project.planton.provider.kubernetes.workload.solrkubernetes.v1B\tSpecProtoP\x01Z\x7fgithub.com/project-planton/project-planton/apis/project/planton/provider/kubernetes/workload/solrkubernetes/v1;solrkubernetesv1\xa2\x02\x06PPPKWS\xaa\x02>Project.Planton.Provider.Kubernetes.Workload.Solrkubernetes.V1\xca\x02>Project\\Planton\\Provider\\Kubernetes\\Workload\\Solrkubernetes\\V1\xe2\x02JProject\\Planton\\Provider\\Kubernetes\\Workload\\Solrkubernetes\\V1\\GPBMetadata\xea\x02DProject::Planton::Provider::Kubernetes::Workload::Solrkubernetes::V1b\x06proto3"

var (
	file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescData []byte
)

func file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDesc), len(file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDesc)))
	})
	return file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDescData
}

var file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_goTypes = []any{
	(*SolrKubernetesSpec)(nil),               // 0: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSpec
	(*SolrKubernetesSolrContainer)(nil),      // 1: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer
	(*SolrKubernetesSolrConfig)(nil),         // 2: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrConfig
	(*SolrKubernetesZookeeperContainer)(nil), // 3: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesZookeeperContainer
	(*kubernetes.IngressSpec)(nil),           // 4: project.planton.shared.kubernetes.IngressSpec
	(*kubernetes.ContainerResources)(nil),    // 5: project.planton.shared.kubernetes.ContainerResources
	(*kubernetes.ContainerImage)(nil),        // 6: project.planton.shared.kubernetes.ContainerImage
	(*descriptorpb.FieldOptions)(nil),        // 7: google.protobuf.FieldOptions
}
var file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_depIdxs = []int32{
	1,  // 0: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSpec.solr_container:type_name -> project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer
	2,  // 1: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSpec.config:type_name -> project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrConfig
	3,  // 2: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSpec.zookeeper_container:type_name -> project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesZookeeperContainer
	4,  // 3: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSpec.ingress:type_name -> project.planton.shared.kubernetes.IngressSpec
	5,  // 4: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer.resources:type_name -> project.planton.shared.kubernetes.ContainerResources
	6,  // 5: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer.image:type_name -> project.planton.shared.kubernetes.ContainerImage
	5,  // 6: project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesZookeeperContainer.resources:type_name -> project.planton.shared.kubernetes.ContainerResources
	7,  // 7: project.planton.provider.kubernetes.workload.solrkubernetes.v1.default_solr_container:extendee -> google.protobuf.FieldOptions
	7,  // 8: project.planton.provider.kubernetes.workload.solrkubernetes.v1.default_zookeeper_container:extendee -> google.protobuf.FieldOptions
	1,  // 9: project.planton.provider.kubernetes.workload.solrkubernetes.v1.default_solr_container:type_name -> project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer
	1,  // 10: project.planton.provider.kubernetes.workload.solrkubernetes.v1.default_zookeeper_container:type_name -> project.planton.provider.kubernetes.workload.solrkubernetes.v1.SolrKubernetesSolrContainer
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	9,  // [9:11] is the sub-list for extension type_name
	7,  // [7:9] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_init() }
func file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_init() {
	if File_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDesc), len(file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_msgTypes,
		ExtensionInfos:    file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_extTypes,
	}.Build()
	File_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto = out.File
	file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_goTypes = nil
	file_project_planton_provider_kubernetes_workload_solrkubernetes_v1_spec_proto_depIdxs = nil
}
