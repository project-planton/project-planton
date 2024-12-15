package variablestf

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/strings/caseconverter"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Known descriptions for certain top-level variables. You can customize or extend this.
var fieldDescriptions = map[string]string{
	"metadata": "Metadata for the resource, including name and labels",
	"spec":     "Specification for Deployment Component",
}

// ProtoToVariablesTF uses proto reflection to determine the Terraform variable schema from a proto message's fields.
// It produces a `variables.tf`-style output with variable blocks for each top-level field (except apiVersion, kind, and status).
func ProtoToVariablesTF(msg proto.Message) (string, error) {
	md := msg.ProtoReflect().Descriptor()

	// We skip apiVersion, kind, and status as variables.
	skipFields := map[string]bool{
		"api_version": true,
		"kind":        true,
		"status":      true,
	}

	var buf bytes.Buffer
	// Iterate over top-level fields
	fields := md.Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		fieldName := string(fd.Name())

		if skipFields[fieldName] {
			continue
		}

		terraformType, err := fieldDescriptorToTerraformType(fd)
		if err != nil {
			return "", errors.Wrapf(err, "failed to convert field %q to terraform type", fieldName)
		}

		snakeKey := caseconverter.ToSnakeCase(fieldName)
		desc := fieldDescriptions[fieldName]
		if desc == "" {
			desc = fmt.Sprintf("Description for %s", snakeKey)
		}

		fmt.Fprintf(&buf, `variable "%s" {
  description = %q
  type = %s
}

`, snakeKey, desc, terraformType)
	}

	return strings.TrimSpace(buf.String()), nil
}

// fieldDescriptorToTerraformType takes a proto field descriptor and returns a Terraform type string.
func fieldDescriptorToTerraformType(fd protoreflect.FieldDescriptor) (string, error) {
	// Handle repeated fields as lists
	if fd.IsList() {
		elemType, err := scalarOrMessageToTFType(fd)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("list(%s)", elemType), nil
	}

	// For singular fields, just convert directly
	return scalarOrMessageToTFType(fd)
}

// scalarOrMessageToTFType converts either a scalar or message field into a Terraform type.
// If it's a message, recursively build an object schema.
func scalarOrMessageToTFType(fd protoreflect.FieldDescriptor) (string, error) {
	kind := fd.Kind()

	switch kind {
	case protoreflect.StringKind:
		return "string", nil
	case protoreflect.BoolKind:
		return "bool", nil
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind:
		// All integers map to "number" in Terraform
		return "number", nil
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		// Floating-point types map to number
		return "number", nil
	case protoreflect.BytesKind:
		// Bytes can be represented as a string (e.g., base64-encoded)
		return "string", nil
	case protoreflect.EnumKind:
		// Enums map to string (you could refine this if you know the allowed values)
		return "string", nil
	case protoreflect.MessageKind:
		// For messages, we build an object type
		return messageToTerraformObject(fd.Message())
	default:
		return "", fmt.Errorf("unsupported field kind: %v", kind)
	}
}

// messageToTerraformObject takes a message descriptor and constructs a Terraform object({ ... }) type.
func messageToTerraformObject(md protoreflect.MessageDescriptor) (string, error) {
	fields := md.Fields()
	if fields.Len() == 0 {
		return "object({})", nil
	}

	var fieldSpecs []string
	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		fieldName := string(f.Name())
		valType, err := fieldDescriptorToTerraformType(f)
		if err != nil {
			return "", err
		}
		snakeKey := caseconverter.ToSnakeCase(fieldName)
		fieldSpecs = append(fieldSpecs, fmt.Sprintf("%s = %s", snakeKey, valType))
	}

	return fmt.Sprintf("object({ %s })", strings.Join(fieldSpecs, ", ")), nil
}
