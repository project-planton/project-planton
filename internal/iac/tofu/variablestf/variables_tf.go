package variablestf

import (
	"bytes"
	"fmt"
	"github.com/project-planton/project-planton/internal/apidocs"
	"strings"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/strings/caseconverter"
	"github.com/pseudomuto/protoc-gen-doc"
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
		fieldLines = append(fieldLines, fmt.Sprintf("%s%s = %s", nextIndent, k, fieldStr))
	}

	return fmt.Sprintf("object({\n%s\n%s})", strings.Join(fieldLines, "\n"), indent)
}

// ProtoToVariablesTF uses proto reflection to determine the Terraform variable schema.
// It now also uses gendoc.Template to include message and field descriptions from proto comments.
func ProtoToVariablesTF(msg proto.Message) (string, error) {
	apiDocsJson, err := apidocs.GetApiDocsJson()
	if err != nil {
		return "", errors.Wrap(err, "failed to get api docs")
	}

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

		tfType, err := fieldDescriptorToTerraformType(fd, md, apiDocsJson)
		if err != nil {
			return "", errors.Wrapf(err, "failed to convert field %q to terraform type", fieldName)
		}

		// Get description from template if available
		desc := findFieldDescription(apiDocsJson, string(md.FullName()), fieldName)

		// If the field is a complex message and we didn't get a field-level doc,
		// try to get the referenced message's doc
		if desc == "" && fd.Kind() == protoreflect.MessageKind {
			desc = findMessageDescription(apiDocsJson, string(fd.Message().FullName()))
		}

		// Fallback logic
		if desc == "" {
			// If still empty, fallback to predefined or generic description
			desc = fieldDescriptions[fieldName]
			if desc == "" {
				desc = fmt.Sprintf("Description for %s", fieldName)
			}
		}

		typeStr := tfType.format(1)

		fmt.Fprintf(&buf, `variable "%s" {
  description = %q
  type = %s
}

`, caseconverter.ToSnakeCase(fieldName), desc, typeStr)
	}

	return strings.TrimSpace(buf.String()), nil
}

func fieldDescriptorToTerraformType(fd protoreflect.FieldDescriptor, parentMsg protoreflect.MessageDescriptor, tmpl *gendoc.Template) (terraformType, error) {
	// If repeated -> list
	if fd.IsList() {
		elemType, err := scalarOrMessageToTFType(parentMsg, fd, tmpl)
		if err != nil {
			return nil, err
		}
		return tfList{elem: elemType}, nil
	}

	return scalarOrMessageToTFType(parentMsg, fd, tmpl)
}

func scalarOrMessageToTFType(parentMsg protoreflect.MessageDescriptor, fd protoreflect.FieldDescriptor, tmpl *gendoc.Template) (terraformType, error) {
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
		return messageToTerraformObject(fd.Message(), fd, tmpl)
	default:
		return nil, fmt.Errorf("unsupported field kind: %v", kind)
	}
}

func messageToTerraformObject(md protoreflect.MessageDescriptor, fd protoreflect.FieldDescriptor, tmpl *gendoc.Template) (terraformType, error) {
	fields := md.Fields()
	obj := tfObject{fields: make(map[string]terraformType)}

	// Skip metadata.version if needed
	shouldSkipVersion := (md.Name() == "Metadata")

	for i := 0; i < fields.Len(); i++ {
		f := fields.Get(i)
		fieldName := string(f.Name())
		if shouldSkipVersion && fieldName == "version" {
			continue
		}

		valType, err := fieldDescriptorToTerraformType(f, md, tmpl)
		if err != nil {
			return nil, err
		}
		snakeKey := caseconverter.ToSnakeCase(fieldName)
		obj.fields[snakeKey] = valType
	}
	return obj, nil
}

// findMessageDescription returns the description of a message from the template
func findMessageDescription(tmpl *gendoc.Template, fullName string) string {
	if tmpl == nil {
		return ""
	}

	for _, f := range tmpl.Files {
		for _, m := range f.Messages {
			if m.FullName == fullName {
				return strings.TrimSpace(m.Description)
			}
		}
	}
	return ""
}

// findFieldDescription returns the description of a field within a message from the template
func findFieldDescription(tmpl *gendoc.Template, messageFullName, fieldName string) string {
	if tmpl == nil {
		return ""
	}

	for _, f := range tmpl.Files {
		for _, m := range f.Messages {
			if m.FullName == messageFullName {
				// found the message, now find the field
				for _, fld := range m.Fields {
					if fld.Name == fieldName {
						return strings.TrimSpace(fld.Description)
					}
				}
			}
		}
	}

	return ""
}
