package provider

import "github.com/project-planton/project-planton/apis/project/planton/shared"

func GetProvider(kindName KindName) iac.KindProvider {
	for provider, kinds := range ToKindMap {
		for _, kn := range kinds {
			if kn == kindName {
				return provider
			}
		}
	}
	return iac.KindProvider_kind_provider_unspecified
}
