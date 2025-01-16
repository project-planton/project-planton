package fieldsextractor

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// StackInputTargetField Field names in the Protobuf message
const StackInputTargetField = "target"
const StackInputTargetSpecField = "spec"

// ExtractApiResourceSpecField extracts the spec field from stack-input using protobuf reflection and returns it.
func ExtractApiResourceSpecField(stackInput proto.Message) (*protoreflect.Message, error) {
	// Check if stackInput is nil
	if stackInput == nil {
		return nil, errors.New("stack-input is nil")
	}

	// Get the protobuf message descriptor
	stackInputProtoReflect := stackInput.ProtoReflect()

	// Retrieve the target field by name
	targetField := stackInputProtoReflect.Descriptor().Fields().ByName(StackInputTargetField)
	if targetField == nil {
		return nil, errors.Errorf("Field %s not found in stack-input", StackInputTargetField)
	}

	// Get the value of the target field and check if it is nil
	targetValue := stackInputProtoReflect.Get(targetField).Message()
	if targetValue.IsValid() == false {
		return nil, errors.Errorf("Field %s is nil in stack-input", StackInputTargetField)
	}

	// Get the protobuf message descriptor for target
	targetProtoReflect := targetValue.Interface().ProtoReflect()

	// Retrieve the spec field by name
	targetSpecField := targetProtoReflect.Descriptor().Fields().ByName(StackInputTargetSpecField)
	if targetSpecField == nil {
		return nil, errors.Errorf("Field %s not found in target", StackInputTargetSpecField)
	}

	// Get the value of the spec field and check if it is nil
	targetSpecValue := targetProtoReflect.Get(targetSpecField).Message()
	if targetSpecValue.IsValid() == false {
		return nil, errors.Errorf("Field %s is nil in target", StackInputTargetSpecField)
	}

	// Return the spec message
	return &targetSpecValue, nil
}
