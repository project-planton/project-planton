package gvk

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared"
)

type GVK struct {
	ApiVersion string                        `yaml:"apiVersion"`
	Kind       string                        `yaml:"kind"`
	Metadata   *shared.CloudResourceMetadata `yaml:"metadata"`
}
