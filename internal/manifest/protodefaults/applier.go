package protodefaults

import (
	"github.com/pkg/errors"
	options_pb "github.com/project-planton/project-planton/apis/project/planton/shared/options"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// ApplyDefaults recursively applies default values from proto field options to a message.
// It traverses all fields in the message and its nested messages, setting defaults
// from the project.planton.shared.options.default option when:
// - The field has a default option defined
// - The field is currently unset/empty
//
// For scalar fields, the default string value is converted to the appropriate type.
// For message fields, the function recurses into the nested message.
func ApplyDefaults(msg proto.Message) error {
	if msg == nil {
		return nil
	}
	return applyDefaultsToMessage(msg.ProtoReflect())
}

// applyDefaultsToMessage recursively applies defaults to a reflected message
func applyDefaultsToMessage(msgReflect protoreflect.Message) error {
	fields := msgReflect.Descriptor().Fields()

	// Iterate through all fields in the message
	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)

		// Skip lists and maps (not supported for defaults)
		if field.IsList() || field.IsMap() {
			continue
		}

		// Check if field is a message type (nested message)
		if field.Kind() == protoreflect.MessageKind {
			// If the field is set, recurse into it
			if msgReflect.Has(field) {
				nestedMsg := msgReflect.Get(field).Message()
				if err := applyDefaultsToMessage(nestedMsg); err != nil {
					return errors.Wrapf(err, "failed to apply defaults to nested field %s", field.FullName())
				}
			}
			continue
		}

		// For scalar fields, apply default if field is not set
		if !msgReflect.Has(field) {
			if err := applyDefaultToField(msgReflect, field); err != nil {
				return errors.Wrapf(err, "failed to apply default to field %s", field.FullName())
			}
		}
	}

	return nil
}

// applyDefaultToField applies the default value to a specific field if it has a default option
func applyDefaultToField(msgReflect protoreflect.Message, field protoreflect.FieldDescriptor) error {
	// Get field options
	options := field.Options()
	if options == nil {
		return nil
	}

	// Extract the default value from field options
	if !proto.HasExtension(options, options_pb.E_Default) {
		return nil // No default defined, skip
	}

	defaultValue, ok := proto.GetExtension(options, options_pb.E_Default).(string)
	if !ok || defaultValue == "" {
		return nil // Default is not a string or is empty, skip
	}

	// Convert the default string to the appropriate field type
	fieldValue, err := ConvertStringToFieldValue(defaultValue, field)
	if err != nil {
		return errors.Wrapf(err, "failed to convert default value '%s'", defaultValue)
	}

	// Set the field value
	msgReflect.Set(field, fieldValue)

	return nil
}
