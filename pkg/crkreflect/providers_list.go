package crkreflect

import "github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"

func ProvidersList() []cloudresourcekind.ProjectPlantonCloudResourceProvider {
	resp := make([]cloudresourcekind.ProjectPlantonCloudResourceProvider, 0)
	// Iterate over all the enum values in ApiResourceKind
	for _, enumValue := range cloudresourcekind.ProjectPlantonCloudResourceProvider_value {
		resp = append(resp, cloudresourcekind.ProjectPlantonCloudResourceProvider(enumValue))
	}
	return resp
}
