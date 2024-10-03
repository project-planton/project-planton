package manifestprotobuf

import (
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

func SetProtoField(msg proto.Message, fieldPath string, value interface{}) (proto.Message, error) {
	// Reflect over the provided proto.Message
	msgReflect := msg.ProtoReflect()

	// Split the field path by dot notation
	fieldParts := strings.Split(fieldPath, ".")

	for i, part := range fieldParts {
		// Get the descriptor for the current field
		fieldDescriptor := msgReflect.Descriptor().Fields().ByName(protoreflect.Name(part))
		if fieldDescriptor == nil {
			return nil, errors.Errorf("field %s not found in message %s", part, msgReflect.Descriptor().FullName())
		}

		if i == len(fieldParts)-1 {
			// Set the last field
			if !fieldDescriptor.IsList() && !fieldDescriptor.IsMap() {
				// Handle scalar fields
				if err := setProtoScalarField(msgReflect, fieldDescriptor, value); err != nil {
					return nil, err
				}
			} else {
				return nil, errors.Errorf("setting list or map fields is not supported")
			}
		} else {
			// Step into the next nested message
			msgReflect = msgReflect.Mutable(fieldDescriptor).Message()
		}
	}
	return msg, nil
}

func setProtoScalarField(msgReflect protoreflect.Message, fieldDescriptor protoreflect.FieldDescriptor, value interface{}) error {
	var fieldValue protoreflect.Value
	switch fieldDescriptor.Kind() {
	case protoreflect.BoolKind:
		v, ok := value.(bool)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected bool", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfBool(v)
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		v, ok := value.(int32)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected int32", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfInt32(v)
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		v, ok := value.(int64)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected int64", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfInt64(v)
	case protoreflect.Uint32Kind, protoreflect.Fixed32Kind:
		v, ok := value.(uint32)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected uint32", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfUint32(v)
	case protoreflect.Uint64Kind, protoreflect.Fixed64Kind:
		v, ok := value.(uint64)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected uint64", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfUint64(v)
	case protoreflect.FloatKind:
		v, ok := value.(float32)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected float32", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfFloat32(v)
	case protoreflect.DoubleKind:
		v, ok := value.(float64)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected float64", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfFloat64(v)
	case protoreflect.StringKind:
		v, ok := value.(string)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected string", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfString(v)
	case protoreflect.BytesKind:
		v, ok := value.([]byte)
		if !ok {
			return errors.Errorf("incorrect type for field %s: expected []byte", fieldDescriptor.FullName())
		}
		fieldValue = protoreflect.ValueOfBytes(v)
	default:
		return errors.Errorf("unsupported field type for field %s", fieldDescriptor.FullName())
	}
	msgReflect.Set(fieldDescriptor, fieldValue)
	return nil
}
