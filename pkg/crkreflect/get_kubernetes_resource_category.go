package crkreflect

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GetKubernetesResourceType returns the (kubernetes_resource_type) option
// recorded on a CloudResourceKind value.  If the option is missing or the
// kind is unknown it returns the “unspecified” sentinel.
func GetKubernetesResourceType(
	kind cloudresourcekind.CloudResourceKind,
) cloudresourcekind.KubernetesCloudResourceCategory {
	var unspecified = cloudresourcekind.
		KubernetesCloudResourceCategory_kubernetes_cloud_resource_category_unspecified

	desc := kind.Descriptor()
	if desc == nil {
		return unspecified
	}

	valDesc := desc.Values().ByNumber(protoreflect.EnumNumber(kind))
	if valDesc == nil {
		return unspecified
	}

	if !proto.HasExtension(valDesc.Options(), cloudresourcekind.E_Meta) {
		return unspecified
	}

	cloudResourceKindMeta, ok := proto.GetExtension(
		valDesc.Options(),
		cloudresourcekind.E_Meta,
	).(*cloudresourcekind.CloudResourceKindMeta)
	if !ok || cloudResourceKindMeta == nil {
		return unspecified
	}
	if cloudResourceKindMeta.KubernetesMeta == nil {
		return unspecified
	}
	return cloudResourceKindMeta.KubernetesMeta.Category
}
