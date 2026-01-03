package gvk

import (
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared"
)

type GVK struct {
	ApiVersion string                        `yaml:"apiVersion"`
	Kind       string                        `yaml:"kind"`
	Metadata   *shared.CloudResourceMetadata `yaml:"metadata"`
}
