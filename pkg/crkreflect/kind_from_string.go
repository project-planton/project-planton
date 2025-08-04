package crkreflect

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"strings"
)

var AliasMap = map[cloudresourcekind.CloudResourceKind][]string{}

func KindFromString(cloudResourceKindString string) cloudresourcekind.CloudResourceKind {
	for kind, aliases := range AliasMap {
		for _, alias := range aliases {
			if alias == cloudResourceKindString {
				return kind
			}
		}
	}

	cloudResourceKindString = strings.ReplaceAll(cloudResourceKindString, "-", "")
	cloudResourceKindString = strings.ReplaceAll(cloudResourceKindString, "_", "")

	for _, k := range KindsList() {
		if strings.EqualFold(k.String(), cloudResourceKindString) {
			return k
		}
	}

	return cloudresourcekind.CloudResourceKind_unspecified
}
