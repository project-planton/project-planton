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

var fieldDescriptions = map[string]string{
	"metadata": "Metadata for the resource, including name and labels",
	"spec":     "Specification for Deployment Component",
}

// terraformType is an interface representing a Terraform type.
type terraformType interface {
	// format returns a string representing this type, formatted nicely.
	format(indentLevel int) string
}

type tfPrimitive string

func (p tfPrimitive) format(indentLevel int) string {
	return string(p)
}

type tfList struct {
	elem terraformType
}

func (l tfList) format(indentLevel int) string {
	return fmt.Sprintf("list(%s)", l.elem.format(indentLevel))
}

type tfObject struct {
	fields map[string]terraformType
}

func (o tfObject) format(indentLevel int) string {
	if len(o.fields) == 0 {
		return "object({})"
	}

	indent := strings.Repeat("  ", indentLevel)
	nextIndent := strings.Repeat("  ", indentLevel+1)

	var fieldLines []string
	for k, v := range o.fields {
		fieldStr := v.format(indentLevel + 1)
		if strings.HasPrefix(fieldStr, "object(") || strings.HasPrefix(fieldStr, "list(") {
			// Complex type: put field type in multiline format
			// We'll try to keep formatting tidy:
			if strings.HasPrefix(fieldStr, "object({") {
				// If object is multiline, we can rely on its own formatting.
				fieldLines = append(fieldLines, fmt.Sprintf("%s%s = %s", nextIndent, k, fieldStr))
			} else {
				// just inline for lists or simpler objects
				fieldLines = append(fieldLines, fmt.Sprintf("%s%s = %s", nextIndent, k, fieldStr))
			}
		} else {
			// Primitive type on the same line
			fieldLines = append(fieldLines, fmt.Sprintf("%s%s = %s", nextIndent, k, fieldStr))
		}
	}

	return fmt.Sprintf("object({\n%s\n%s})", strings.Join(fieldLines, "\n"), indent)
}

// ProtoToVariablesTF uses proto reflection to determine the Terraform variable schema.
func ProtoToVariablesTF(msg proto.Message) (string, error) {
	md := msg.ProtoReflect().Descriptor()

	skipFields := map[string]bool{
		"api_version": true,
		"kind":        true,
		"status":      true,
	}

	var buf bytes.Buffer
	fields := md.Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		fieldName := string(fd.Name())

		if skipFields[fieldName] {
			continue
		}

		tfType, err := fieldDescriptorToTerraformType(fd, md)
		if err != nil {
			return "", errors.Wrapf(err, "failed to convert field %q to terraform type", fieldName)
		}

		desc := fieldDescriptions[fieldName]
		if desc == "" {
			desc = fmt.Sprintf("Description for %s", fieldName)
		}

		// Pretty-print the type with indentation
		typeStr := tfType.format(1)

		fmt.Fprintf(&buf, `variable "%s" {
  description = %q
  type = %s
}

`, caseconverter.ToSnakeCase(fieldName), desc, typeStr)
	}

	return strings.TrimSpace(buf.String()), nil
}

func fieldDescriptorToTerraformType(fd protoreflect.FieldDescriptor, parentMsg protoreflect.MessageDescriptor) (terraformType, error) {
	// If repeated -> list
	if fd.IsList() {
		elemType, err := scalarOrMessageToTFType(fd.Message(), fd)
		if err != nil {
			return nil, err
		}
		return tfList{elem: elemType}, nil
	}

	return scalarOrMessageToTFType(fd.Message(), fd)
}

func scalarOrMessageToTFType(parentMsg protoreflect.MessageDescriptor, fd protoreflect.FieldDescriptor) (terraformType, error) {
	kind := fd.Kind()

	switch kind {
	case protoreflect.StringKind:
		return tfPrimitive("string"), nil
	case protoreflect.BoolKind:
		return tfPrimitive("bool"), nil
	case protoreflect.Int32Kind, protoreflect.Int64Kind,
		protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Sint32Kind, protoreflect.Sint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind,
		protoreflect.Sfixed32Kind, protoreflect.Sfixed64Kind,
		protoreflect.FloatKind, protoreflect.DoubleKind:
		return tfPrimitive("number"), nil
	case protoreflect.BytesKind:
		return tfPrimitive("string"), nil
	case protoreflect.EnumKind:
		return tfPrimitive("string"), nil
	case protoreflect.MessageKind:
		return messageToTerraformObject(fd.Message(), fd)
	default:
		return nil, fmt.Errorf("unsupported field kind: %v", kind)
	}
}

func messageToTerraformObject(md protoreflect.MessageDescriptor, fd protoreflect.FieldDescriptor) (terraformType, error) {
	fields := md.Fields()
	obj := tfObject{fields: make(map[string]terraformType)}

	// Skip metadata.version
	shouldSkipVersion := (md.Name() == "Metadata")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		fieldName := string(f.Name())
		if shouldSkipVersion && fieldName == "version" {
			continue
		}

		valType, err := fieldDescriptorToTerraformType(f, md)
		if err != nil {
			return nil, err
		}
		snakeKey := caseconverter.ToSnakeCase(fieldName)
		obj.fields[snakeKey] = valType
	}
	return obj, nil
}
