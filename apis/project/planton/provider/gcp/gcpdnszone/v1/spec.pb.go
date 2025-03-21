// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: project/planton/provider/gcp/gcpdnszone/v1/spec.proto

package gcpdnszonev1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	dnsrecordtype "github.com/project-planton/project-planton/apis/project/planton/shared/networking/enums/dnsrecordtype"
	_ "github.com/project-planton/project-planton/apis/project/planton/shared/options"
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

// **GcpDnsZoneSpec** defines the configuration for creating a Google Cloud DNS Managed Zone.
// This message specifies the parameters needed to create and manage a DNS zone within a specified GCP project.
// It includes the project ID, optional service accounts for IAM permissions, and DNS records to be added to the zone.
type GcpDnsZoneSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The ID of the GCP project where the Managed Zone is created.
	ProjectId string `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// An optional list of GCP service accounts that are granted permissions to manage DNS records in the Managed Zone.
	// These accounts are typically workload identities, such as those used by cert-manager,
	// and are added when new environments are created or updated.
	IamServiceAccounts []string `protobuf:"bytes,2,rep,name=iam_service_accounts,json=iamServiceAccounts,proto3" json:"iam_service_accounts,omitempty"`
	// The DNS records to be added to the Managed Zone.
	Records []*GcpDnsRecord `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *GcpDnsZoneSpec) Reset() {
	*x = GcpDnsZoneSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpDnsZoneSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpDnsZoneSpec) ProtoMessage() {}

func (x *GcpDnsZoneSpec) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpDnsZoneSpec.ProtoReflect.Descriptor instead.
func (*GcpDnsZoneSpec) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescGZIP(), []int{0}
}

func (x *GcpDnsZoneSpec) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *GcpDnsZoneSpec) GetIamServiceAccounts() []string {
	if x != nil {
		return x.IamServiceAccounts
	}
	return nil
}

func (x *GcpDnsZoneSpec) GetRecords() []*GcpDnsRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

// **GcpDnsRecord** represents a DNS record to be added to the Managed Zone.
// It includes the record type, name, values, and TTL (Time To Live) settings.
type GcpDnsRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The DNS record type (e.g., A, AAAA, CNAME).
	RecordType dnsrecordtype.DnsRecordType `protobuf:"varint,1,opt,name=record_type,json=recordType,proto3,enum=project.planton.shared.networking.enums.dnsrecordtype.DnsRecordType" json:"record_type,omitempty"`
	// The name of the DNS record (e.g., "example.com." or "dev.example.com.").
	// This value should always end with a dot to signify a fully qualified domain name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The list of values for the DNS record.
	// If the record type is CNAME, each value in the list should end with a dot.
	Values []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
	// The Time To Live (TTL) for the DNS record, in seconds.
	TtlSeconds int32 `protobuf:"varint,4,opt,name=ttl_seconds,json=ttlSeconds,proto3" json:"ttl_seconds,omitempty"`
}

func (x *GcpDnsRecord) Reset() {
	*x = GcpDnsRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcpDnsRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcpDnsRecord) ProtoMessage() {}

func (x *GcpDnsRecord) ProtoReflect() protoreflect.Message {
	mi := &file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcpDnsRecord.ProtoReflect.Descriptor instead.
func (*GcpDnsRecord) Descriptor() ([]byte, []int) {
	return file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescGZIP(), []int{1}
}

func (x *GcpDnsRecord) GetRecordType() dnsrecordtype.DnsRecordType {
	if x != nil {
		return x.RecordType
	}
	return dnsrecordtype.DnsRecordType(0)
}

func (x *GcpDnsRecord) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GcpDnsRecord) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *GcpDnsRecord) GetTtlSeconds() int32 {
	if x != nil {
		return x.TtlSeconds
	}
	return 0
}

var File_project_planton_provider_gcp_gcpdnszone_v1_spec_proto protoreflect.FileDescriptor

var file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDesc = []byte{
	0x0a, 0x35, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67,
	0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65,
	0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x4b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f,
	0x6e, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x69, 0x6e, 0x67, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x64, 0x6e, 0x73, 0x72, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x64, 0x6e, 0x73, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x0e,
	0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x5a, 0x6f, 0x6e, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x25,
	0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x69, 0x61, 0x6d, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x12, 0x69, 0x61, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x52, 0x0a, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x67, 0x63, 0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f,
	0x6e, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x07, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0x8e, 0x03, 0x0a, 0x0c,
	0x47, 0x63, 0x70, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x6d, 0x0a, 0x0b,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x6e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x64, 0x6e, 0x73, 0x72,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x44, 0x6e, 0x73, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0xc3, 0x01, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0xae, 0x01, 0xba, 0x48, 0xaa,
	0x01, 0xba, 0x01, 0xa3, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x4e, 0x61, 0x6d,
	0x65, 0x20, 0x73, 0x68, 0x6f, 0x75, 0x6c, 0x64, 0x20, 0x62, 0x65, 0x20, 0x61, 0x6e, 0x79, 0x20,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x44, 0x4e, 0x53, 0x20, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x20, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x71, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x73, 0x28, 0x27, 0x5e, 0x28, 0x3f, 0x3a, 0x5b, 0x2a, 0x5d, 0x5b, 0x2e, 0x5d, 0x29,
	0x3f, 0x28, 0x3f, 0x3a, 0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x28, 0x3f, 0x3a,
	0x5b, 0x5f, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d,
	0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x29, 0x3f, 0x5b, 0x2e, 0x5d, 0x29, 0x2b, 0x28,
	0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x3f, 0x3a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d,
	0x39, 0x2d, 0x5d, 0x7b, 0x30, 0x2c, 0x36, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x29, 0x3f, 0x29, 0x3f, 0x24, 0x27, 0x29, 0xc8, 0x01, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0b, 0x74, 0x74, 0x6c, 0x5f, 0x73, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0x8a, 0xa6, 0x1d, 0x02, 0x36, 0x30,
	0x52, 0x0a, 0x74, 0x74, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x42, 0xf3, 0x02, 0x0a,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70, 0x6c, 0x61,
	0x6e, 0x74, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x67, 0x63,
	0x70, 0x2e, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x76, 0x31, 0x42,
	0x09, 0x53, 0x70, 0x65, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x67, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x70, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x63, 0x70, 0x2f, 0x67, 0x63, 0x70, 0x64, 0x6e,
	0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a,
	0x6f, 0x6e, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x05, 0x50, 0x50, 0x50, 0x47, 0x47, 0xaa, 0x02, 0x2a,
	0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x63, 0x70, 0x2e, 0x47, 0x63, 0x70,
	0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x2a, 0x50, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x64, 0x6e, 0x73,
	0x7a, 0x6f, 0x6e, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x36, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x5c, 0x50, 0x6c, 0x61, 0x6e, 0x74, 0x6f, 0x6e, 0x5c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5c, 0x47, 0x63, 0x70, 0x5c, 0x47, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e,
	0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x2f, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x3a, 0x3a, 0x50, 0x6c, 0x61, 0x6e,
	0x74, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x47,
	0x63, 0x70, 0x3a, 0x3a, 0x47, 0x63, 0x70, 0x64, 0x6e, 0x73, 0x7a, 0x6f, 0x6e, 0x65, 0x3a, 0x3a,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescOnce sync.Once
	file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescData = file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDesc
)

func file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescGZIP() []byte {
	file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescOnce.Do(func() {
		file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescData = protoimpl.X.CompressGZIP(file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescData)
	})
	return file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDescData
}

var file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_goTypes = []any{
	(*GcpDnsZoneSpec)(nil),           // 0: project.planton.provider.gcp.gcpdnszone.v1.GcpDnsZoneSpec
	(*GcpDnsRecord)(nil),             // 1: project.planton.provider.gcp.gcpdnszone.v1.GcpDnsRecord
	(dnsrecordtype.DnsRecordType)(0), // 2: project.planton.shared.networking.enums.dnsrecordtype.DnsRecordType
}
var file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_depIdxs = []int32{
	1, // 0: project.planton.provider.gcp.gcpdnszone.v1.GcpDnsZoneSpec.records:type_name -> project.planton.provider.gcp.gcpdnszone.v1.GcpDnsRecord
	2, // 1: project.planton.provider.gcp.gcpdnszone.v1.GcpDnsRecord.record_type:type_name -> project.planton.shared.networking.enums.dnsrecordtype.DnsRecordType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_init() }
func file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_init() {
	if File_project_planton_provider_gcp_gcpdnszone_v1_spec_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GcpDnsZoneSpec); i {
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
		file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GcpDnsRecord); i {
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
			RawDescriptor: file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_goTypes,
		DependencyIndexes: file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_depIdxs,
		MessageInfos:      file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_msgTypes,
	}.Build()
	File_project_planton_provider_gcp_gcpdnszone_v1_spec_proto = out.File
	file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_rawDesc = nil
	file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_goTypes = nil
	file_project_planton_provider_gcp_gcpdnszone_v1_spec_proto_depIdxs = nil
}
