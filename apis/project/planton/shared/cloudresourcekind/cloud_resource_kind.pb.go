// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: project/planton/shared/cloudresourcekind/cloud_resource_kind.proto

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

type CloudResourceKind int32

const (
	// 0: Default/unspecified
	CloudResourceKind_unspecified CloudResourceKind = 0
	// 1–49: Test/dev/custom
	CloudResourceKind_FirstTestCloudApiResource  CloudResourceKind = 1
	CloudResourceKind_SecondTestCloudApiResource CloudResourceKind = 2
	CloudResourceKind_ThirdTestCloudApiResource  CloudResourceKind = 3
	// 50–199: saas platform resources
	CloudResourceKind_ConfluentKafka    CloudResourceKind = 50
	CloudResourceKind_MongodbAtlas      CloudResourceKind = 51
	CloudResourceKind_SnowflakeDatabase CloudResourceKind = 52
	// 200–399: AWS resources
	CloudResourceKind_AwsAlb             CloudResourceKind = 200
	CloudResourceKind_AwsCertManagerCert CloudResourceKind = 201
	CloudResourceKind_AwsCloudFront      CloudResourceKind = 202
	CloudResourceKind_AwsDynamodb        CloudResourceKind = 203
	CloudResourceKind_AwsEcrRepo         CloudResourceKind = 204
	CloudResourceKind_AwsEcsCluster      CloudResourceKind = 205
	CloudResourceKind_AwsEcsService      CloudResourceKind = 206
	CloudResourceKind_AwsEksCluster      CloudResourceKind = 207
	CloudResourceKind_AwsIamRole         CloudResourceKind = 208
	CloudResourceKind_AwsLambda          CloudResourceKind = 209
	CloudResourceKind_AwsRdsCluster      CloudResourceKind = 210
	CloudResourceKind_AwsRdsInstance     CloudResourceKind = 211
	CloudResourceKind_AwsRoute53Zone     CloudResourceKind = 212
	CloudResourceKind_AwsS3Bucket        CloudResourceKind = 213
	CloudResourceKind_AwsSecretsManager  CloudResourceKind = 214
	CloudResourceKind_AwsSecurityGroup   CloudResourceKind = 215
	CloudResourceKind_AwsStaticWebsite   CloudResourceKind = 216
	CloudResourceKind_AwsVpc             CloudResourceKind = 217
	// 400–599: Azure resources
	CloudResourceKind_AzureAksCluster CloudResourceKind = 400
	CloudResourceKind_AzureKeyVault   CloudResourceKind = 401
	// 600–799: GCP resources
	CloudResourceKind_GcpArtifactRegistryRepo CloudResourceKind = 600
	CloudResourceKind_GcpCloudCdn             CloudResourceKind = 601
	CloudResourceKind_GcpCloudFunction        CloudResourceKind = 602
	CloudResourceKind_GcpCloudRun             CloudResourceKind = 603
	CloudResourceKind_GcpCloudSql             CloudResourceKind = 604
	CloudResourceKind_GcpDnsZone              CloudResourceKind = 605
	CloudResourceKind_GcpGcsBucket            CloudResourceKind = 606
	CloudResourceKind_GcpGkeAddonBundle       CloudResourceKind = 607
	CloudResourceKind_GcpGkeCluster           CloudResourceKind = 608
	CloudResourceKind_GcpSecretsManager       CloudResourceKind = 609
	CloudResourceKind_GcpStaticWebsite        CloudResourceKind = 610
	CloudResourceKind_GcpProject              CloudResourceKind = 611
	// 800–999: Kubernetes resources
	CloudResourceKind_ArgocdKubernetes         CloudResourceKind = 800
	CloudResourceKind_CronJobKubernetes        CloudResourceKind = 801
	CloudResourceKind_ElasticsearchKubernetes  CloudResourceKind = 802
	CloudResourceKind_GitlabKubernetes         CloudResourceKind = 803
	CloudResourceKind_GrafanaKubernetes        CloudResourceKind = 804
	CloudResourceKind_HelmRelease              CloudResourceKind = 805
	CloudResourceKind_JenkinsKubernetes        CloudResourceKind = 806
	CloudResourceKind_KafkaKubernetes          CloudResourceKind = 807
	CloudResourceKind_KeycloakKubernetes       CloudResourceKind = 808
	CloudResourceKind_KubernetesHttpEndpoint   CloudResourceKind = 809
	CloudResourceKind_LocustKubernetes         CloudResourceKind = 810
	CloudResourceKind_MicroserviceKubernetes   CloudResourceKind = 811
	CloudResourceKind_MongodbKubernetes        CloudResourceKind = 812
	CloudResourceKind_Neo4jKubernetes          CloudResourceKind = 813
	CloudResourceKind_OpenFgaKubernetes        CloudResourceKind = 814
	CloudResourceKind_PostgresKubernetes       CloudResourceKind = 815
	CloudResourceKind_PrometheusKubernetes     CloudResourceKind = 816
	CloudResourceKind_RedisKubernetes          CloudResourceKind = 817
	CloudResourceKind_SignozKubernetes         CloudResourceKind = 818
	CloudResourceKind_SolrKubernetes           CloudResourceKind = 819
	CloudResourceKind_StackJobRunnerKubernetes CloudResourceKind = 820
	CloudResourceKind_TemporalKubernetes       CloudResourceKind = 821
	CloudResourceKind_NatsKubernetes           CloudResourceKind = 822
)

// Enum value maps for CloudResourceKind.
var (
	CloudResourceKind_name = map[int32]string{
		0:   "unspecified",
		1:   "FirstTestCloudApiResource",
		2:   "SecondTestCloudApiResource",
		3:   "ThirdTestCloudApiResource",
		50:  "ConfluentKafka",
		51:  "MongodbAtlas",
		52:  "SnowflakeDatabase",
		200: "AwsAlb",
		201: "AwsCertManagerCert",
		202: "AwsCloudFront",
		203: "AwsDynamodb",
		204: "AwsEcrRepo",
		205: "AwsEcsCluster",
		206: "AwsEcsService",
		207: "AwsEksCluster",
		208: "AwsIamRole",
		209: "AwsLambda",
		210: "AwsRdsCluster",
		211: "AwsRdsInstance",
		212: "AwsRoute53Zone",
		213: "AwsS3Bucket",
		214: "AwsSecretsManager",
		215: "AwsSecurityGroup",
		216: "AwsStaticWebsite",
		217: "AwsVpc",
		400: "AzureAksCluster",
		401: "AzureKeyVault",
		600: "GcpArtifactRegistryRepo",
		601: "GcpCloudCdn",
		602: "GcpCloudFunction",
		603: "GcpCloudRun",
		604: "GcpCloudSql",
		605: "GcpDnsZone",
		606: "GcpGcsBucket",
		607: "GcpGkeAddonBundle",
		608: "GcpGkeCluster",
		609: "GcpSecretsManager",
		610: "GcpStaticWebsite",
		611: "GcpProject",
		800: "ArgocdKubernetes",
		801: "CronJobKubernetes",
		802: "ElasticsearchKubernetes",
		803: "GitlabKubernetes",
		804: "GrafanaKubernetes",
		805: "HelmRelease",
		806: "JenkinsKubernetes",
		807: "KafkaKubernetes",
		808: "KeycloakKubernetes",
		809: "KubernetesHttpEndpoint",
		810: "LocustKubernetes",
		811: "MicroserviceKubernetes",
		812: "MongodbKubernetes",
		813: "Neo4jKubernetes",
		814: "OpenFgaKubernetes",
		815: "PostgresKubernetes",
		816: "PrometheusKubernetes",
		817: "RedisKubernetes",
		818: "SignozKubernetes",
		819: "SolrKubernetes",
		820: "StackJobRunnerKubernetes",
		821: "TemporalKubernetes",
		822: "NatsKubernetes",
	}
	CloudResourceKind_value = map[string]int32{
		"unspecified":                0,
		"FirstTestCloudApiResource":  1,
		"SecondTestCloudApiResource": 2,
		"ThirdTestCloudApiResource":  3,
		"ConfluentKafka":             50,
		"MongodbAtlas":               51,
		"SnowflakeDatabase":          52,
		"AwsAlb":                     200,
		"AwsCertManagerCert":         201,
		"AwsCloudFront":              202,
		"AwsDynamodb":                203,
		"AwsEcrRepo":                 204,
		"AwsEcsCluster":              205,
		"AwsEcsService":              206,
		"AwsEksCluster":              207,
		"AwsIamRole":                 208,
		"AwsLambda":                  209,
		"AwsRdsCluster":              210,
		"AwsRdsInstance":             211,
		"AwsRoute53Zone":             212,
		"AwsS3Bucket":                213,
		"AwsSecretsManager":          214,
		"AwsSecurityGroup":           215,
		"AwsStaticWebsite":           216,
		"AwsVpc":                     217,
		"AzureAksCluster":            400,
		"AzureKeyVault":              401,
		"GcpArtifactRegistryRepo":    600,
		"GcpCloudCdn":                601,
		"GcpCloudFunction":           602,
		"GcpCloudRun":                603,
		"GcpCloudSql":                604,
		"GcpDnsZone":                 605,
		"GcpGcsBucket":               606,
		"GcpGkeAddonBundle":          607,
		"GcpGkeCluster":              608,
		"GcpSecretsManager":          609,
		"GcpStaticWebsite":           610,
		"GcpProject":                 611,
		"ArgocdKubernetes":           800,
		"CronJobKubernetes":          801,
		"ElasticsearchKubernetes":    802,
		"GitlabKubernetes":           803,
		"GrafanaKubernetes":          804,
		"HelmRelease":                805,
		"JenkinsKubernetes":          806,
		"KafkaKubernetes":            807,
		"KeycloakKubernetes":         808,
		"KubernetesHttpEndpoint":     809,
		"LocustKubernetes":           810,
		"MicroserviceKubernetes":     811,
		"MongodbKubernetes":          812,
		"Neo4jKubernetes":            813,
		"OpenFgaKubernetes":          814,
		"PostgresKubernetes":         815,
		"PrometheusKubernetes":       816,
		"RedisKubernetes":            817,
		"SignozKubernetes":           818,
		"SolrKubernetes":             819,
		"StackJobRunnerKubernetes":   820,
		"TemporalKubernetes":         821,
		"NatsKubernetes":             822,
	}
)

func (x CloudResourceKind) Enum() *CloudResourceKind {
	p := new(CloudResourceKind)
	*p = x
	return p
}

func (x CloudResourceKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CloudResourceKind) Descriptor() protoreflect.EnumDescriptor {
	return file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_enumTypes[0].Descriptor()
}

func (CloudResourceKind) Type() protoreflect.EnumType {
	return &file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_enumTypes[0]
}

func (x CloudResourceKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CloudResourceKind.Descriptor instead.
func (CloudResourceKind) EnumDescriptor() ([]byte, []int) {
	return file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDescGZIP(), []int{0}
}

var File_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto protoreflect.FileDescriptor

const file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDesc = "" +
	"\n" +
	"Bproject/planton/shared/cloudresourcekind/cloud_resource_kind.proto\x12(project.planton.shared.cloudresourcekind*\xf0\n" +
	"\n" +
	"\x11CloudResourceKind\x12\x0f\n" +
	"\vunspecified\x10\x00\x12\x1d\n" +
	"\x19FirstTestCloudApiResource\x10\x01\x12\x1e\n" +
	"\x1aSecondTestCloudApiResource\x10\x02\x12\x1d\n" +
	"\x19ThirdTestCloudApiResource\x10\x03\x12\x12\n" +
	"\x0eConfluentKafka\x102\x12\x10\n" +
	"\fMongodbAtlas\x103\x12\x15\n" +
	"\x11SnowflakeDatabase\x104\x12\v\n" +
	"\x06AwsAlb\x10\xc8\x01\x12\x17\n" +
	"\x12AwsCertManagerCert\x10\xc9\x01\x12\x12\n" +
	"\rAwsCloudFront\x10\xca\x01\x12\x10\n" +
	"\vAwsDynamodb\x10\xcb\x01\x12\x0f\n" +
	"\n" +
	"AwsEcrRepo\x10\xcc\x01\x12\x12\n" +
	"\rAwsEcsCluster\x10\xcd\x01\x12\x12\n" +
	"\rAwsEcsService\x10\xce\x01\x12\x12\n" +
	"\rAwsEksCluster\x10\xcf\x01\x12\x0f\n" +
	"\n" +
	"AwsIamRole\x10\xd0\x01\x12\x0e\n" +
	"\tAwsLambda\x10\xd1\x01\x12\x12\n" +
	"\rAwsRdsCluster\x10\xd2\x01\x12\x13\n" +
	"\x0eAwsRdsInstance\x10\xd3\x01\x12\x13\n" +
	"\x0eAwsRoute53Zone\x10\xd4\x01\x12\x10\n" +
	"\vAwsS3Bucket\x10\xd5\x01\x12\x16\n" +
	"\x11AwsSecretsManager\x10\xd6\x01\x12\x15\n" +
	"\x10AwsSecurityGroup\x10\xd7\x01\x12\x15\n" +
	"\x10AwsStaticWebsite\x10\xd8\x01\x12\v\n" +
	"\x06AwsVpc\x10\xd9\x01\x12\x14\n" +
	"\x0fAzureAksCluster\x10\x90\x03\x12\x12\n" +
	"\rAzureKeyVault\x10\x91\x03\x12\x1c\n" +
	"\x17GcpArtifactRegistryRepo\x10\xd8\x04\x12\x10\n" +
	"\vGcpCloudCdn\x10\xd9\x04\x12\x15\n" +
	"\x10GcpCloudFunction\x10\xda\x04\x12\x10\n" +
	"\vGcpCloudRun\x10\xdb\x04\x12\x10\n" +
	"\vGcpCloudSql\x10\xdc\x04\x12\x0f\n" +
	"\n" +
	"GcpDnsZone\x10\xdd\x04\x12\x11\n" +
	"\fGcpGcsBucket\x10\xde\x04\x12\x16\n" +
	"\x11GcpGkeAddonBundle\x10\xdf\x04\x12\x12\n" +
	"\rGcpGkeCluster\x10\xe0\x04\x12\x16\n" +
	"\x11GcpSecretsManager\x10\xe1\x04\x12\x15\n" +
	"\x10GcpStaticWebsite\x10\xe2\x04\x12\x0f\n" +
	"\n" +
	"GcpProject\x10\xe3\x04\x12\x15\n" +
	"\x10ArgocdKubernetes\x10\xa0\x06\x12\x16\n" +
	"\x11CronJobKubernetes\x10\xa1\x06\x12\x1c\n" +
	"\x17ElasticsearchKubernetes\x10\xa2\x06\x12\x15\n" +
	"\x10GitlabKubernetes\x10\xa3\x06\x12\x16\n" +
	"\x11GrafanaKubernetes\x10\xa4\x06\x12\x10\n" +
	"\vHelmRelease\x10\xa5\x06\x12\x16\n" +
	"\x11JenkinsKubernetes\x10\xa6\x06\x12\x14\n" +
	"\x0fKafkaKubernetes\x10\xa7\x06\x12\x17\n" +
	"\x12KeycloakKubernetes\x10\xa8\x06\x12\x1b\n" +
	"\x16KubernetesHttpEndpoint\x10\xa9\x06\x12\x15\n" +
	"\x10LocustKubernetes\x10\xaa\x06\x12\x1b\n" +
	"\x16MicroserviceKubernetes\x10\xab\x06\x12\x16\n" +
	"\x11MongodbKubernetes\x10\xac\x06\x12\x14\n" +
	"\x0fNeo4jKubernetes\x10\xad\x06\x12\x16\n" +
	"\x11OpenFgaKubernetes\x10\xae\x06\x12\x17\n" +
	"\x12PostgresKubernetes\x10\xaf\x06\x12\x19\n" +
	"\x14PrometheusKubernetes\x10\xb0\x06\x12\x14\n" +
	"\x0fRedisKubernetes\x10\xb1\x06\x12\x15\n" +
	"\x10SignozKubernetes\x10\xb2\x06\x12\x13\n" +
	"\x0eSolrKubernetes\x10\xb3\x06\x12\x1d\n" +
	"\x18StackJobRunnerKubernetes\x10\xb4\x06\x12\x17\n" +
	"\x12TemporalKubernetes\x10\xb5\x06\x12\x13\n" +
	"\x0eNatsKubernetes\x10\xb6\x06B\xe4\x02\n" +
	",com.project.planton.shared.cloudresourcekindB\x16CloudResourceKindProtoP\x01ZXgithub.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind\xa2\x02\x04PPSC\xaa\x02(Project.Planton.Shared.Cloudresourcekind\xca\x02(Project\\Planton\\Shared\\Cloudresourcekind\xe2\x024Project\\Planton\\Shared\\Cloudresourcekind\\GPBMetadata\xea\x02+Project::Planton::Shared::Cloudresourcekindb\x06proto3"

var (
	file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDescOnce sync.Once
	file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDescData []byte
)

func file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDescGZIP() []byte {
	file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDescOnce.Do(func() {
		file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDesc), len(file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDesc)))
	})
	return file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDescData
}

var file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_goTypes = []any{
	(CloudResourceKind)(0), // 0: project.planton.shared.cloudresourcekind.CloudResourceKind
}
var file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_init() }
func file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_init() {
	if File_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDesc), len(file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_goTypes,
		DependencyIndexes: file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_depIdxs,
		EnumInfos:         file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_enumTypes,
	}.Build()
	File_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto = out.File
	file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_goTypes = nil
	file_project_planton_shared_cloudresourcekind_cloud_resource_kind_proto_depIdxs = nil
}
