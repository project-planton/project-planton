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
) cloudresourcekind.ProjectPlantonKubernetesResourceType {
	var unspecified = cloudresourcekind.
		ProjectPlantonKubernetesResourceType_project_planton_kubernetes_resource_type_unspecified

	desc := kind.Descriptor()
	if desc == nil {
		return unspecified
	}

	valDesc := desc.Values().ByNumber(protoreflect.EnumNumber(kind))
	if valDesc == nil {
		return unspecified
	}

	if !proto.HasExtension(valDesc.Options(), cloudresourcekind.E_KubernetesResourceType) {
		return unspecified
	}

	rt, ok := proto.GetExtension(
		valDesc.Options(),
		cloudresourcekind.E_KubernetesResourceType,
	).(cloudresourcekind.ProjectPlantonKubernetesResourceType)
	if !ok {
		return unspecified
	}
	return rt
}
