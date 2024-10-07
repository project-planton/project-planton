package manifest

import (
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/manifest/manifestprotobuf"
	"google.golang.org/protobuf/proto"
)

func LoadWithOverrides(manifestPath string, valueOverrides map[string]string) (proto.Message, error) {
	manifest, err := LoadManifest(manifestPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load manifest")
	}
	for key, value := range valueOverrides {
		manifest, err = manifestprotobuf.SetProtoField(manifest, key, value)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to set %s=%s", key, value)
		}
	}
	return manifest, nil
}
