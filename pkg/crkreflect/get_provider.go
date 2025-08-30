package crkreflect

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// GetProvider returns the Cloud‑resource **provider** recorded in the
// (provider) enum‑value option of the given CloudResourceKind.
//
// If the kind is unknown or the option is absent, the function returns the
// “unspecified” sentinel.
func GetProvider(
	kind cloudresourcekind.CloudResourceKind,
) cloudresourcekind.CloudResourceProvider {
	desc := kind.Descriptor()
	if desc == nil {
		return cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified
	}

	valDesc := desc.Values().ByNumber(protoreflect.EnumNumber(kind))
	if valDesc == nil {
		return cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified
	}

	if !proto.HasExtension(valDesc.Options(), cloudresourcekind.E_Provider) {
		return cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified
	}

	prov, ok := proto.GetExtension(
		valDesc.Options(),
		cloudresourcekind.E_Provider,
	).(cloudresourcekind.CloudResourceProvider)
	if !ok {
		return cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified
	}
	return prov
}
