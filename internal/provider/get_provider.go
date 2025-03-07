package provider

import "github.com/project-planton/project-planton/apis/project/planton/shared"

func GetProvider(kindName KindName) shared.KindProvider {
	for provider, kinds := range ToKindMap {
		for _, kn := range kinds {
			if kn == kindName {
				return provider
			}
		}
	}
	return shared.KindProvider_kind_provider_unspecified
}
