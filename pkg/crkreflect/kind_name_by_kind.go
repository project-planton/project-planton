package crkreflect

import (
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
	log "github.com/sirupsen/logrus"
)

func ExtractKindNameByKind(kind cloudresourcekind.CloudResourceKind) string {
	kindMeta, err := KindMeta(kind)
	if err != nil {
		log.Errorf("failed to extract kind meta by kind %s", kind.String())
		return ""
	}
	return kindMeta.Name
}
