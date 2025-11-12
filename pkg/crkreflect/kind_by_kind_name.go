package crkreflect

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
)

func KindByKindName(kindName string) (cloudresourcekind.CloudResourceKind, error) {
	// Iterate over the ApiResourceKind enum values
	for _, enumValue := range cloudresourcekind.CloudResourceKind_value {
		kind := cloudresourcekind.CloudResourceKind(enumValue)
		kindMeta, err := KindMeta(kind)
		if err != nil {
			continue
		}
		// Compare the kind_name in the metadata with the message kind name
		if kindMeta.Name == kindName {
			// If it matches, return the corresponding ApiResourceKind
			return kind, nil
		}
	}
	return cloudresourcekind.CloudResourceKind_unspecified,
		errors.Errorf("no matching CloudResourceKind found for kind: %s", kindName)
}
