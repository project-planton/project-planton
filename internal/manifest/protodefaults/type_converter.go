package protodefaults

import (
	"math"
	"strconv"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ConvertStringToFieldValue converts a string default value to the appropriate protoreflect.Value
// based on the field's type descriptor. Returns an error if the conversion fails.
func ConvertStringToFieldValue(defaultStr string, field protoreflect.FieldDescriptor) (protoreflect.Value, error) {
	kind := field.Kind()

	switch kind {
	case protoreflect.StringKind:
		return protoreflect.ValueOfString(defaultStr), nil

	case protoreflect.Int32Kind:
		val, err := strconv.ParseInt(defaultStr, 10, 32)
		if err != nil {
			return protoreflect.Value{}, errors.Wrapf(err, "failed to convert '%s' to int32 for field %s", defaultStr, field.FullName())
		}
		return protoreflect.ValueOfInt32(int32(val)), nil

	case protoreflect.Int64Kind:
		val, err := strconv.ParseInt(defaultStr, 10, 64)
		if err != nil {
			return protoreflect.Value{}, errors.Wrapf(err, "failed to convert '%s' to int64 for field %s", defaultStr, field.FullName())
		}
		return protoreflect.ValueOfInt64(val), nil

	case protoreflect.Uint32Kind:
		val, err := strconv.ParseUint(defaultStr, 10, 32)
		if err != nil {
			return protoreflect.Value{}, errors.Wrapf(err, "failed to convert '%s' to uint32 for field %s", defaultStr, field.FullName())
		}
		return protoreflect.ValueOfUint32(uint32(val)), nil

	case protoreflect.Uint64Kind:
		val, err := strconv.ParseUint(defaultStr, 10, 64)
		if err != nil {
			return protoreflect.Value{}, errors.Wrapf(err, "failed to convert '%s' to uint64 for field %s", defaultStr, field.FullName())
		}
		return protoreflect.ValueOfUint64(val), nil

	case protoreflect.BoolKind:
		val, err := strconv.ParseBool(defaultStr)
		if err != nil {
			return protoreflect.Value{}, errors.Wrapf(err, "failed to convert '%s' to bool for field %s", defaultStr, field.FullName())
		}
		return protoreflect.ValueOfBool(val), nil

	case protoreflect.FloatKind:
		val, err := strconv.ParseFloat(defaultStr, 32)
		if err != nil {
			return protoreflect.Value{}, errors.Wrapf(err, "failed to convert '%s' to float32 for field %s", defaultStr, field.FullName())
		}
		// Check for overflow
		if val > math.MaxFloat32 || val < -math.MaxFloat32 {
			return protoreflect.Value{}, errors.Errorf("value '%s' overflows float32 for field %s", defaultStr, field.FullName())
		}
		return protoreflect.ValueOfFloat32(float32(val)), nil

	case protoreflect.DoubleKind:
		val, err := strconv.ParseFloat(defaultStr, 64)
		if err != nil {
			return protoreflect.Value{}, errors.Wrapf(err, "failed to convert '%s' to float64 for field %s", defaultStr, field.FullName())
		}
		return protoreflect.ValueOfFloat64(val), nil

	case protoreflect.EnumKind:
		enumDescriptor := field.Enum()
		enumValue := enumDescriptor.Values().ByName(protoreflect.Name(defaultStr))
		if enumValue == nil {
			return protoreflect.Value{}, errors.Errorf("enum value '%s' not found in enum %s for field %s", defaultStr, enumDescriptor.FullName(), field.FullName())
		}
		return protoreflect.ValueOfEnum(enumValue.Number()), nil

	default:
		return protoreflect.Value{}, errors.Errorf("unsupported field type %s for field %s", kind, field.FullName())
	}
}

