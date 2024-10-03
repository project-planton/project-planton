package manifestvalidator

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

const (
	SpecProtoFieldName = "spec"
)

func ExtractSpec(manifestObject proto.Message) (proto.Message, error) {
	// Get the protobuf message descriptor for target
	manifestProtoReflect := manifestObject.ProtoReflect()

	// Retrieve the spec field by name
	specField := manifestProtoReflect.Descriptor().Fields().ByName(SpecProtoFieldName)
	if specField == nil {
		return nil, errors.Errorf("Field %s not found in manifest", SpecProtoFieldName)
	}

	// Get the value of the spec field and check if it is nil
	specValue := manifestProtoReflect.Get(specField).Message()
	if specValue.IsValid() == false {
		return nil, errors.Errorf("field %s is nil in manifest", SpecProtoFieldName)
	}

	// Return the spec message
	return specValue.Interface(), nil
}
