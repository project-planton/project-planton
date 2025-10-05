package crkreflect

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/cloudresourcekind"
	"google.golang.org/protobuf/proto"
)

// NewInstance creates a new instance of the proto message for the given CloudResourceKind.
// This function returns a fresh instance, not a shared reference from ToMessageMap.
// Use this instead of directly accessing ToMessageMap to avoid shared state issues.
func NewInstance(kind cloudresourcekind.CloudResourceKind) (proto.Message, error) {
	template := ToMessageMap[kind]
	
	if template == nil {
		return nil, errors.Errorf("unsupported cloud resource kind: %s", kind.String())
	}
	
	// Create a new instance using protobuf reflection
	// This ensures each call gets its own independent object
	newInstance := template.ProtoReflect().New().Interface()
	
	return newInstance, nil
}

