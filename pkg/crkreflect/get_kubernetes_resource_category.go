package crkreflect

import (
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
)

// GetKubernetesResourceCategory returns the (kubernetes_resource_category) option
// recorded on a CloudResourceKind value.  If the option is missing or the
// kind is unknown it returns the “unspecified” sentinel.
func GetKubernetesResourceCategory(
	kind cloudresourcekind.CloudResourceKind,
) cloudresourcekind.KubernetesCloudResourceCategory {
	var unspecified = cloudresourcekind.
		KubernetesCloudResourceCategory_kubernetes_cloud_resource_category_unspecified
	kindMeta, err := KindMeta(kind)
	if err != nil {
		return unspecified
	}
	if kindMeta.KubernetesMeta == nil {
		return unspecified
	}
	return kindMeta.KubernetesMeta.Category
}
