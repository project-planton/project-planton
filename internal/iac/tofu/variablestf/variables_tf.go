package variablestf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/project-planton/project-planton/pkg/strings/caseconverter"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtoToVariablesTF takes a protobuf message, inspects its structure,
// and returns the content of a `variables.tf` file that defines Terraform
// variables based on the top-level fields (except apiVersion, kind, and status).
//
// Each top-level field (like `metadata`, `spec`) becomes a Terraform variable block.
//
// Terraform types are inferred from the JSON data structure:
// - string -> string
// - bool -> bool
// - number -> number
// - array -> list(...)
// – - map -> object({ ... }) or map(type) if homogeneous primitive map
//
// The function tries to produce a meaningful `description` for each variable,
// but you may customize it further.
func ProtoToVariablesTF(msg proto.Message) (string, error) {
	jsonBytes, err := protojson.MarshalOptions{
		EmitUnpopulated: true,
	}.Marshal(msg)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal proto to json")
	}

	var data map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return "", errors.Wrap(err, "failed to unmarshal json")
	}

	// We skip apiVersion, kind, and status as variables.
	delete(data, "apiVersion")
	delete(data, "kind")
	delete(data, "status")

	// For each remaining top-level key, generate a variable block.
	var buf bytes.Buffer
	for k, val := range data {
		snakeKey := caseconverter.ToSnakeCase(k)
		tfType, err := inferTerraformType(val)
		if err != nil {
			return "", errors.Wrapf(err, "failed to infer type for key %q", k)
		}

		// Construct a variable block
		// You can enhance the description logic to be more meaningful for your schema.
		description := fmt.Sprintf("Description for %s", snakeKey)
		buf.WriteString(fmt.Sprintf(`variable "%s" {
  description = %q
  type = %s
}

`, snakeKey, description, tfType))
	}

	return strings.TrimSpace(buf.String()), nil
}

// inferTerraformType inspects a Go value (from JSON) and returns a Terraform type string.
// Supported conversions:
// - string -> "string"
// - bool -> "bool"
// - float64 -> "number"
// - []interface{} -> "list(<type>)" (type inferred from elements; if heterogeneous, fallback to list(any))
// - map[string]interface{} -> "object({ ... })" or "map(<type>)"
//
// If a map is homogeneous (all values the same primitive type), we use map(type).
// If not homogeneous or contains nested objects/lists, we use object({ ... }).
func inferTerraformType(v interface{}) (string, error) {
	switch val := v.(type) {
	case string:
		return "string", nil
	case bool:
		return "bool", nil
	case float64:
		return "number", nil
	case nil:
		// null could be anything; default to string (or any) if you want.
		// Terraform 0.12+ doesn't have a direct 'any' type in this context,
		// so we must pick a type. Let's pick "string" as a fallback.
		return "string", nil
	case []interface{}:
		// Infer the type of elements. If empty, assume list(any), but we must pick a type.
		if len(val) == 0 {
			return "list(string)", nil
		}
		elementType, err := inferListElementType(val)
		if err != nil {
			return "", err
		}
		return fmt.Sprintf("list(%s)", elementType), nil
	case map[string]interface{}:
		// Check if map is homogeneous and all primitive => map(<type>)
		// Otherwise, object({ ... })
		return inferMapType(val)
	default:
		return "", errors.Errorf("unsupported type: %T", v)
	}
}

// inferListElementType checks the elements of a list.
// If all elements are the same primitive type, returns that type.
// If elements differ or contain objects/lists, fallback to a broader type.
//
// For simplicity, if we have heterogeneous primitives or complex structures,
// we could fallback to list(string) or try object. Realistically, you'd tailor
// this logic to your known schemas.
func inferListElementType(arr []interface{}) (string, error) {
	if len(arr) == 0 {
		return "string", nil
	}

	firstType, err := inferTerraformType(arr[0])
	if err != nil {
		return "", err
	}

	// If first element is primitive (string, bool, number) or complex (object/list),
	// check all others match:
	for _, elem := range arr[1:] {
		elemType, err := inferTerraformType(elem)
		if err != nil {
			return "", err
		}
		if elemType != firstType {
			// If there's a mismatch, you have two choices:
			// 1. Return an error.
			// 2. Fallback to a generic type (like list(string)).
			// Here, we fallback to "string" for simplicity.
			return "string", nil
		}
	}

	return firstType, nil
}

// inferMapType checks a map to determine if it can be a simple map(type) or must be object({ ... }).
func inferMapType(m map[string]interface{}) (string, error) {
	// If empty map, object({}) is fine:
	if len(m) == 0 {
		return "object({})", nil
	}

	// Let's see if we can form a uniform map(type) first:
	var firstValType string
	allSameType := true

	for _, val := range m {
		valType, err := inferTerraformType(val)
		if err != nil {
			return "", err
		}
		// Check if valType is a primitive type (string, bool, number)
		// If it's object(...) or list(...), we must use object({}) instead
		if strings.HasPrefix(valType, "object(") || strings.HasPrefix(valType, "list(") {
			// Complex type found, so we do object({ ... }) and recurse
			return inferObjectType(m)
		}

		if firstValType == "" {
			firstValType = valType
		} else if valType != firstValType {
			allSameType = false
		}
	}

	if allSameType && firstValType != "" {
		// It's a homogeneous map of primitive values:
		return fmt.Sprintf("map(%s)", firstValType), nil
	}

	// If not all same type, fallback to object:
	return inferObjectType(m)
}

// inferObjectType constructs an object({ ... }) type from a map.
// For each key, infer its type and build up the object definition.
func inferObjectType(m map[string]interface{}) (string, error) {
	var fields []string
	for k, val := range m {
		valType, err := inferTerraformType(val)
		if err != nil {
			return "", err
		}
		snakeKey := caseconverter.ToSnakeCase(k)
		fields = append(fields, fmt.Sprintf("%s = %s", snakeKey, valType))
	}
	return fmt.Sprintf("object({ %s })", strings.Join(fields, ", ")), nil
}
