// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/provider/azure/azureakscluster/v1/spec.proto

package azureaksclusterv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"
	_ "github.com/project-planton/project-planton/apis/project/planton/shared/options"
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

// Possible network plugin options for the AKS cluster.
type AzureAksClusterNetworkPlugin int32

const (
	AzureAksClusterNetworkPlugin_AZURE_CNI AzureAksClusterNetworkPlugin = 0
	AzureAksClusterNetworkPlugin_KUBENET   AzureAksClusterNetworkPlugin = 1
)

// Enum value maps for AzureAksClusterNetworkPlugin.
var (
	AzureAksClusterNetworkPlugin_name = map[int32]string{
		0: "AZURE_CNI",
		1: "KUBENET",
	}
	AzureAksClusterNetworkPlugin_value = map[string]int32{
		"AZURE_CNI": 0,
		"KUBENET":   1,
	}
)

func (x AzureAksClusterNetworkPlugin) Enum() *AzureAksClusterNetworkPlugin {
	p := new(AzureAksClusterNetworkPlugin)
	*p = x
	return p
}

func (x AzureAksClusterNetworkPlugin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AzureAksClusterNetworkPlugin) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_provider_azure_azureakscluster_v1_spec_proto_enumTypes[0].Descriptor()
}

func (AzureAksClusterNetworkPlugin) Type() protoreflect.EnumType {
	return &file_project_planton_provider_azure_azureakscluster_v1_spec_proto_enumTypes[0]
}

func (x AzureAksClusterNetworkPlugin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AzureAksClusterNetworkPlugin.Descriptor instead.
func (AzureAksClusterNetworkPlugin) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescGZIP(), []int{0}
}

// AzureAksClusterSpec defines the specification required to deploy an Azure Kubernetes Service (AKS) cluster.
// This minimal spec covers essential configurations to achieve a production-ready environment while avoiding extraneous complexity.
type AzureAksClusterSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Azure region in which to create the AKS cluster (e.g., "eastus").
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// The Azure resource ID of the Virtual Network subnet to use for cluster nodes.
	// This should reference the subnet created by an AzureVirtualNetwork resource.
	VnetSubnetId *v1.StringValueOrRef `protobuf:"bytes,2,opt,name=vnet_subnet_id,json=vnetSubnetId,proto3" json:"vnet_subnet_id,omitempty"`
	// Networking plugin for the AKS cluster: "azure_cni" for Azure CNI (advanced networking) or "kubenet" for basic networking.
	// Defaults to Azure CNI if not specified.
	NetworkPlugin AzureAksClusterNetworkPlugin `protobuf:"varint,3,opt,name=network_plugin,json=networkPlugin,proto3,enum=project.planton.provider.azure.azureakscluster.v1.AzureAksClusterNetworkPlugin" json:"network_plugin,omitempty"`
	// Kubernetes version for the cluster control plane.
	// If not specified, Azure will default to a supported version. It is recommended to explicitly set a version (e.g., "1.30") for production clusters.
	KubernetesVersion string `protobuf:"bytes,4,opt,name=kubernetes_version,json=kubernetesVersion,proto3" json:"kubernetes_version,omitempty"`
	// Deploy the cluster as a private cluster (no public API server endpoint).
	// When set to true, the API server endpoint will be private. When false (default), a public endpoint is created.
	PrivateClusterEnabled bool `protobuf:"varint,5,opt,name=private_cluster_enabled,json=privateClusterEnabled,proto3" json:"private_cluster_enabled,omitempty"`
	// Authorized IP address ranges (CIDR blocks) that are allowed to access the API server.
	// This is applicable only if the cluster has a public endpoint. Leave empty to allow all (0.0.0.0/0) or for private clusters.
	AuthorizedIpRanges []string `protobuf:"bytes,6,rep,name=authorized_ip_ranges,json=authorizedIpRanges,proto3" json:"authorized_ip_ranges,omitempty"`
	// Disable Azure Active Directory integration for Kubernetes RBAC.
	// By default, AKS clusters have Azure AD integration enabled (this field is false). Set to true to disable Azure AD RBAC integration.
	DisableAzureAdRbac bool `protobuf:"varint,7,opt,name=disable_azure_ad_rbac,json=disableAzureAdRbac,proto3" json:"disable_azure_ad_rbac,omitempty"`
	// The Azure resource ID of a Log Analytics Workspace for AKS monitoring integration.
	// If provided, the AKS cluster will send logs and metrics to this Log Analytics workspace.
	LogAnalyticsWorkspaceId string `protobuf:"bytes,8,opt,name=log_analytics_workspace_id,json=logAnalyticsWorkspaceId,proto3" json:"log_analytics_workspace_id,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *AzureAksClusterSpec) Reset() {
	*x = AzureAksClusterSpec{}
	mi := &file_project_planton_provider_azure_azureakscluster_v1_spec_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AzureAksClusterSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AzureAksClusterSpec) ProtoMessage() {}

func (x *AzureAksClusterSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_azure_azureakscluster_v1_spec_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AzureAksClusterSpec.ProtoReflect.Descriptor instead.
func (*AzureAksClusterSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *AzureAksClusterSpec) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AzureAksClusterSpec) GetVnetSubnetId() *v1.StringValueOrRef {
	if x != nil {
		return x.VnetSubnetId
	}
	return nil
}

func (x *AzureAksClusterSpec) GetNetworkPlugin() AzureAksClusterNetworkPlugin {
	if x != nil {
		return x.NetworkPlugin
	}
	return AzureAksClusterNetworkPlugin_AZURE_CNI
}

func (x *AzureAksClusterSpec) GetKubernetesVersion() string {
	if x != nil {
		return x.KubernetesVersion
	}
	return ""
}

func (x *AzureAksClusterSpec) GetPrivateClusterEnabled() bool {
	if x != nil {
		return x.PrivateClusterEnabled
	}
	return false
}

func (x *AzureAksClusterSpec) GetAuthorizedIpRanges() []string {
	if x != nil {
		return x.AuthorizedIpRanges
	}
	return nil
}

func (x *AzureAksClusterSpec) GetDisableAzureAdRbac() bool {
	if x != nil {
		return x.DisableAzureAdRbac
	}
	return false
}

func (x *AzureAksClusterSpec) GetLogAnalyticsWorkspaceId() string {
	if x != nil {
		return x.LogAnalyticsWorkspaceId
	}
	return ""
}

var File_project_planton_provider_azure_azureakscluster_v1_spec_proto protoreflect.FileDescriptor

const file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDesc = "" +
	"\n" +
	"<project/planton/provider/azure/azureakscluster/v1/spec.proto\x121project.planton.provider.azure.azureakscluster.v1\x1a\x1bbuf/validate/validate.proto\x1a6project/planton/shared/foreignkey/v1/foreign_key.proto\x1a,project/planton/shared/options/options.proto\"\xc5\x06\n" +
	"\x13AzureAksClusterSpec\x12\x1e\n" +
	"\x06region\x18\x01 \x01(\tB\x06\xbaH\x03\xc8\x01\x01R\x06region\x12\x8b\x01\n" +
	"\x0evnet_subnet_id\x18\x02 \x01(\v26.project.planton.shared.foreignkey.v1.StringValueOrRefB-\xbaH\x03\xc8\x01\x01\x88\xd4a\x95\x03\x92\xd4a\x1estatus.outputs.nodes_subnet_idR\fvnetSubnetId\x12v\n" +
	"\x0enetwork_plugin\x18\x03 \x01(\x0e2O.project.planton.provider.azure.azureakscluster.v1.AzureAksClusterNetworkPluginR\rnetworkPlugin\x127\n" +
	"\x12kubernetes_version\x18\x04 \x01(\tB\b\x92\xa6\x1d\x041.30R\x11kubernetesVersion\x126\n" +
	"\x17private_cluster_enabled\x18\x05 \x01(\bR\x15privateClusterEnabled\x12\xb7\x01\n" +
	"\x14authorized_ip_ranges\x18\x06 \x03(\tB\x84\x01\xbaH\x80\x01\x92\x01}\"{ry2w^(?:25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])(?:\\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])){3}(?:\\/(?:3[0-2]|[12]?[0-9]))?$R\x12authorizedIpRanges\x121\n" +
	"\x15disable_azure_ad_rbac\x18\a \x01(\bR\x12disableAzureAdRbac\x12\xa9\x01\n" +
	"\x1alog_analytics_workspace_id\x18\b \x01(\tBl\xbaHirg2e^/subscriptions/[^/]+/resourceGroups/[^/]+/providers/Microsoft\\.OperationalInsights/workspaces/[^/]+$R\x17logAnalyticsWorkspaceId*:\n" +
	"\x1cAzureAksClusterNetworkPlugin\x12\r\n" +
	"\tAZURE_CNI\x10\x00\x12\v\n" +
	"\aKUBENET\x10\x01B\xa2\x03\n" +
	"5com.project.planton.provider.azure.azureakscluster.v1B\tSpecProtoP\x01Zsgithub.com/project-planton/project-planton/apis/project/planton/provider/azure/azureakscluster/v1;azureaksclusterv1\xa2\x02\x05PPPAA\xaa\x021Project.Planton.Provider.Azure.Azureakscluster.V1\xca\x021Project\\Planton\\Provider\\Azure\\Azureakscluster\\V1\xe2\x02=Project\\Planton\\Provider\\Azure\\Azureakscluster\\V1\\GPBMetadata\xea\x026Project::Planton::Provider::Azure::Azureakscluster::V1b\x06proto3"

var (
	file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescData []byte
)

func file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDesc), len(file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDesc)))
	})
	return file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDescData
}

var file_project_planton_provider_azure_azureakscluster_v1_spec_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_project_planton_provider_azure_azureakscluster_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_project_planton_provider_azure_azureakscluster_v1_spec_proto_goTypes = []any{
	(AzureAksClusterNetworkPlugin)(0), // 0: project.planton.provider.azure.azureakscluster.v1.AzureAksClusterNetworkPlugin
	(*AzureAksClusterSpec)(nil),       // 1: project.planton.provider.azure.azureakscluster.v1.AzureAksClusterSpec
	(*v1.StringValueOrRef)(nil),       // 2: project.planton.shared.foreignkey.v1.StringValueOrRef
}
var file_project_planton_provider_azure_azureakscluster_v1_spec_proto_depIdxs = []int32{
	2, // 0: project.planton.provider.azure.azureakscluster.v1.AzureAksClusterSpec.vnet_subnet_id:type_name -> project.planton.shared.foreignkey.v1.StringValueOrRef
	0, // 1: project.planton.provider.azure.azureakscluster.v1.AzureAksClusterSpec.network_plugin:type_name -> project.planton.provider.azure.azureakscluster.v1.AzureAksClusterNetworkPlugin
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_project_planton_provider_azure_azureakscluster_v1_spec_proto_init() }
func file_project_planton_provider_azure_azureakscluster_v1_spec_proto_init() {
	if File_project_planton_provider_azure_azureakscluster_v1_spec_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDesc), len(file_project_planton_provider_azure_azureakscluster_v1_spec_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_azure_azureakscluster_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_azure_azureakscluster_v1_spec_proto_depIdxs,
		EnumInfos:         file_project_planton_provider_azure_azureakscluster_v1_spec_proto_enumTypes,
		MessageInfos:      file_project_planton_provider_azure_azureakscluster_v1_spec_proto_msgTypes,
	}.Build()
	File_project_planton_provider_azure_azureakscluster_v1_spec_proto = out.File
	file_project_planton_provider_azure_azureakscluster_v1_spec_proto_goTypes = nil
	file_project_planton_provider_azure_azureakscluster_v1_spec_proto_depIdxs = nil
}
