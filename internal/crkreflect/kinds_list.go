package crkreflect

import "github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"

func KindsList() []cloudresourcekind.CloudResourceKind {
	resp := make([]cloudresourcekind.CloudResourceKind, 0)
	// Iterate over all the enum values in ApiResourceKind
	for _, enumValue := range cloudresourcekind.CloudResourceKind_value {
		if enumValue == 0 {
			// Skip the zero value, which is usually the "unspecified" value
			continue
		}
		resp = append(resp, cloudresourcekind.CloudResourceKind(enumValue))
	}
	return resp
}
