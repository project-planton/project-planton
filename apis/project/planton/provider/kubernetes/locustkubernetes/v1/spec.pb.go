// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/kubernetes/locustkubernetes/v1/spec.proto

package locustkubernetesv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	kubernetes "github.com/project-planton/project-planton/apis/project/planton/shared/kubernetes"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// **LocustKubernetesSpec** defines the overall configuration for deploying a Locust load testing cluster on Kubernetes.
// This message encapsulates environmental context, Kubernetes deployment specifications, load testing parameters,
// and Helm chart values for customizing the deployment. By configuring these parameters, you can set up a scalable
// and customizable load testing environment to simulate user traffic and measure application performance.
type LocustKubernetesSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The master container specifications for the Locust cluster.
	// This defines the resource allocation and number of replicas for the master node.
	MasterContainer *LocustKubernetesContainer `protobuf:"bytes,1,opt,name=master_container,json=masterContainer,proto3" json:"master_container,omitempty"`
	// The worker container specifications for the Locust cluster.
	// This defines the resource allocation and number of replicas for the worker nodes.
	WorkerContainer *LocustKubernetesContainer `protobuf:"bytes,2,opt,name=worker_container,json=workerContainer,proto3" json:"worker_container,omitempty"`
	// The ingress configuration for the Locust deployment.
	Ingress *kubernetes.IngressSpec `protobuf:"bytes,3,opt,name=ingress,proto3" json:"ingress,omitempty"`
	// The load test parameters, including the main test script, additional library files,
	// and extra Python pip packages needed for test execution.
	// This specifies how the Locust nodes will simulate traffic and interact with the target application.
	LoadTest *LocustKubernetesLoadTest `protobuf:"bytes,4,opt,name=load_test,json=loadTest,proto3" json:"load_test,omitempty"`
	// A map of key-value pairs providing additional customization options for the Helm chart used
	// to deploy the Locust cluster. These values allow for further refinement of the deployment,
	// such as customizing resource limits, setting environment variables, or specifying version tags.
	// For detailed information on the available options, refer to the Helm chart documentation at:
	// https://github.com/deliveryhero/helm-charts/tree/master/stable/locust#values
	HelmValues map[string]string `protobuf:"bytes,5,rep,name=helm_values,json=helmValues,proto3" json:"helm_values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LocustKubernetesSpec) Reset() {
	*x = LocustKubernetesSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocustKubernetesSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocustKubernetesSpec) ProtoMessage() {}

func (x *LocustKubernetesSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocustKubernetesSpec.ProtoReflect.Descriptor instead.
func (*LocustKubernetesSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *LocustKubernetesSpec) GetMasterContainer() *LocustKubernetesContainer {
	if x != nil {
		return x.MasterContainer
	}
	return nil
}

func (x *LocustKubernetesSpec) GetWorkerContainer() *LocustKubernetesContainer {
	if x != nil {
		return x.WorkerContainer
	}
	return nil
}

func (x *LocustKubernetesSpec) GetIngress() *kubernetes.IngressSpec {
	if x != nil {
		return x.Ingress
	}
	return nil
}

func (x *LocustKubernetesSpec) GetLoadTest() *LocustKubernetesLoadTest {
	if x != nil {
		return x.LoadTest
	}
	return nil
}

func (x *LocustKubernetesSpec) GetHelmValues() map[string]string {
	if x != nil {
		return x.HelmValues
	}
	return nil
}

// **LocustKubernetesContainer** specifies the container configuration for Locust master and worker nodes.
// It includes resource allocations for CPU and memory, as well as the number of replicas to deploy.
// Proper configuration ensures optimal performance and scalability of your load testing environment.
type LocustKubernetesContainer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The number of replicas for the container.
	// This determines the level of concurrency and load generation capabilities.
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// The CPU and memory resources allocated to the Locust container.
	Resources *kubernetes.ContainerResources `protobuf:"bytes,2,opt,name=resources,proto3" json:"resources,omitempty"`
}

func (x *LocustKubernetesContainer) Reset() {
	*x = LocustKubernetesContainer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocustKubernetesContainer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocustKubernetesContainer) ProtoMessage() {}

func (x *LocustKubernetesContainer) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocustKubernetesContainer.ProtoReflect.Descriptor instead.
func (*LocustKubernetesContainer) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescGZIP(), []int{1}
}

func (x *LocustKubernetesContainer) GetReplicas() int32 {
	if x != nil {
		return x.Replicas
	}
	return 0
}

func (x *LocustKubernetesContainer) GetResources() *kubernetes.ContainerResources {
	if x != nil {
		return x.Resources
	}
	return nil
}

// **LocustKubernetesLoadTest** defines the specification for a load test using a Locust cluster.
// This message includes the primary Python script for Locust and any additional library files
// necessary to execute the load test. By providing these details, you can define the behavior
// of simulated users and customize the load test according to your application's requirements.
type LocustKubernetesLoadTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A unique identifier or name for this particular load test specification.
	// It is used to reference or distinguish this test configuration among others within a testing suite or environment.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The Python code for the main Locust test script.
	// This script defines the behavior of the simulated users and is crucial for executing the load test.
	MainPyContent string `protobuf:"bytes,2,opt,name=main_py_content,json=mainPyContent,proto3" json:"main_py_content,omitempty"`
	// A map where each entry consists of a filename and its associated Python code content.
	// These files typically contain additional classes or functions required by the main_py_content script.
	// The key of the map is the filename, and the value is the file content.
	LibFilesContent map[string]string `protobuf:"bytes,3,rep,name=lib_files_content,json=libFilesContent,proto3" json:"lib_files_content,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A list of extra Python pip packages that are required for the load test.
	// These packages will be installed in the environment where the load test is executed,
	// allowing for extended functionality or custom dependencies to be included easily.
	PipPackages []string `protobuf:"bytes,4,rep,name=pip_packages,json=pipPackages,proto3" json:"pip_packages,omitempty"`
}

func (x *LocustKubernetesLoadTest) Reset() {
	*x = LocustKubernetesLoadTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocustKubernetesLoadTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocustKubernetesLoadTest) ProtoMessage() {}

func (x *LocustKubernetesLoadTest) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocustKubernetesLoadTest.ProtoReflect.Descriptor instead.
func (*LocustKubernetesLoadTest) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescGZIP(), []int{2}
}

func (x *LocustKubernetesLoadTest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LocustKubernetesLoadTest) GetMainPyContent() string {
	if x != nil {
		return x.MainPyContent
	}
	return ""
}

func (x *LocustKubernetesLoadTest) GetLibFilesContent() map[string]string {
	if x != nil {
		return x.LibFilesContent
	}
	return nil
}

func (x *LocustKubernetesLoadTest) GetPipPackages() []string {
	if x != nil {
		return x.PipPackages
	}
	return nil
}

var file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*LocustKubernetesContainer)(nil),
		Field:         528001,
		Name:          "project.planton.provider.kubernetes.locustkubernetes.v1.default_master_container",
		Tag:           "bytes,528001,opt,name=default_master_container",
		Filename:      "project/planton/provider/kubernetes/locustkubernetes/v1/spec.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*LocustKubernetesContainer)(nil),
		Field:         528002,
		Name:          "project.planton.provider.kubernetes.locustkubernetes.v1.default_worker_container",
		Tag:           "bytes,528002,opt,name=default_worker_container",
		Filename:      "project/planton/provider/kubernetes/locustkubernetes/v1/spec.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer default_master_container = 528001;
	E_DefaultMasterContainer = &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_extTypes[0]
	// optional project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer default_worker_container = 528002;
	E_DefaultWorkerContainer = &file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_extTypes[1]
)

var File_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto protoreflect.FileDescriptor

var file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x42, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x37, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c,
	0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62,
	0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x32, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x73, 0x68,
	0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe5, 0x05, 0x0a, 0x14, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0xa4, 0x01, 0x0a, 0x10, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x75,
	0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0x25, 0x8a, 0xe8, 0x81, 0x02, 0x20,
	0x08, 0x01, 0x12, 0x1c, 0x0a, 0x0c, 0x0a, 0x05, 0x31, 0x30, 0x30, 0x30, 0x6d, 0x12, 0x03, 0x31,
	0x47, 0x69, 0x12, 0x0c, 0x0a, 0x03, 0x35, 0x30, 0x6d, 0x12, 0x05, 0x31, 0x30, 0x30, 0x4d, 0x69,
	0x52, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65,
	0x72, 0x12, 0xa4, 0x01, 0x0a, 0x10, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x42, 0x25, 0x92, 0xe8, 0x81, 0x02, 0x20, 0x08, 0x01, 0x12, 0x1c, 0x0a, 0x0c, 0x0a, 0x05, 0x31,
	0x30, 0x30, 0x30, 0x6d, 0x12, 0x03, 0x31, 0x47, 0x69, 0x12, 0x0c, 0x0a, 0x03, 0x35, 0x30, 0x6d,
	0x12, 0x05, 0x31, 0x30, 0x30, 0x4d, 0x69, 0x52, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x07, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x65, 0x64, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x49, 0x6e,
	0x67, 0x72, 0x65, 0x73, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x69, 0x6e, 0x67, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x76, 0x0a, 0x09, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x51, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x75,
	0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x4c, 0x6f, 0x61, 0x64, 0x54, 0x65, 0x73, 0x74, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x08, 0x6c, 0x6f, 0x61, 0x64, 0x54, 0x65, 0x73, 0x74, 0x12, 0x7e, 0x0a, 0x0b, 0x68, 0x65,
	0x6c, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x5d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72,
	0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74,
	0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x48,
	0x65, 0x6c, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0a,
	0x68, 0x65, 0x6c, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x3d, 0x0a, 0x0f, 0x48, 0x65,
	0x6c, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8c, 0x01, 0x0a, 0x19, 0x4c, 0x6f,
	0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x72, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x73, 0x12, 0x53, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x35, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61,
	0x69, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x52, 0x09, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xea, 0x02, 0x0a, 0x18, 0x4c, 0x6f, 0x63,
	0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4c, 0x6f, 0x61,
	0x64, 0x54, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x2e, 0x0a, 0x0f, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x70, 0x79, 0x5f, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x0d, 0x6d, 0x61, 0x69, 0x6e, 0x50, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x9a, 0x01, 0x0a, 0x11, 0x6c, 0x69, 0x62, 0x5f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x66, 0x2e,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x4c, 0x6f, 0x61, 0x64, 0x54, 0x65, 0x73, 0x74,
	0x2e, 0x4c, 0x69, 0x62, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0f, 0x6c,
	0x69, 0x62, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x21,
	0x0a, 0x0c, 0x70, 0x69, 0x70, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x69, 0x70, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x73, 0x1a, 0x42, 0x0a, 0x14, 0x4c, 0x69, 0x62, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x3a, 0xad, 0x01, 0x0a, 0x18, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x81, 0x9d, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x16, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x3a, 0xad, 0x01, 0x0a, 0x18, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e,
	0x65, 0x72, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x82, 0x9d, 0x20, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e,
	0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e,
	0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x52, 0x16, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x42, 0xc7, 0x03, 0x0a, 0x3b, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2e, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x7a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x6c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x6c, 0x6f, 0x63, 0x75,
	0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x05, 0x50, 0x50, 0x50, 0x4b, 0x4c, 0xaa, 0x02, 0x37, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x4c, 0x6f, 0x63,
	0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x37, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74,
	0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x43, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x5c, 0x4c, 0x6f, 0x63, 0x75, 0x73, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x3c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x4c, 0x6f, 0x63, 0x75, 0x73,
	0x74, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescData = file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDesc
)

func file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescData)
	})
	return file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDescData
}

var file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_goTypes = []any{
	(*LocustKubernetesSpec)(nil),          // 0: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec
	(*LocustKubernetesContainer)(nil),     // 1: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer
	(*LocustKubernetesLoadTest)(nil),      // 2: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesLoadTest
	nil,                                   // 3: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec.HelmValuesEntry
	nil,                                   // 4: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesLoadTest.LibFilesContentEntry
	(*kubernetes.IngressSpec)(nil),        // 5: project.planton.shared.kubernetes.IngressSpec
	(*kubernetes.ContainerResources)(nil), // 6: project.planton.shared.kubernetes.ContainerResources
	(*descriptorpb.FieldOptions)(nil),     // 7: google.protobuf.FieldOptions
}
var file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_depIdxs = []int32{
	1,  // 0: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec.master_container:type_name -> project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer
	1,  // 1: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec.worker_container:type_name -> project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer
	5,  // 2: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec.ingress:type_name -> project.planton.shared.kubernetes.IngressSpec
	2,  // 3: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec.load_test:type_name -> project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesLoadTest
	3,  // 4: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec.helm_values:type_name -> project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesSpec.HelmValuesEntry
	6,  // 5: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer.resources:type_name -> project.planton.shared.kubernetes.ContainerResources
	4,  // 6: project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesLoadTest.lib_files_content:type_name -> project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesLoadTest.LibFilesContentEntry
	7,  // 7: project.planton.provider.kubernetes.locustkubernetes.v1.default_master_container:extendee -> google.protobuf.FieldOptions
	7,  // 8: project.planton.provider.kubernetes.locustkubernetes.v1.default_worker_container:extendee -> google.protobuf.FieldOptions
	1,  // 9: project.planton.provider.kubernetes.locustkubernetes.v1.default_master_container:type_name -> project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer
	1,  // 10: project.planton.provider.kubernetes.locustkubernetes.v1.default_worker_container:type_name -> project.planton.provider.kubernetes.locustkubernetes.v1.LocustKubernetesContainer
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	9,  // [9:11] is the sub-list for extension type_name
	7,  // [7:9] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_init() }
func file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_init() {
	if File_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*LocustKubernetesSpec); i {
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
		file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LocustKubernetesContainer); i {
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
		file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LocustKubernetesLoadTest); i {
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
			RawDescriptor: file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_msgTypes,
		ExtensionInfos:    file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_extTypes,
	}.Build()
	File_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto = out.File
	file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_rawDesc = nil
	file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_goTypes = nil
	file_project_planton_provider_kubernetes_locustkubernetes_v1_spec_proto_depIdxs = nil
}