package tfvars

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"strings"
)

// ProtoToTFVars converts a given protobuf message into a Terraform tfvars-compatible
// string. The primary use case is to take a structured proto, typically loaded and validated
// from a YAML or JSON input, and produce a corresponding tfvars file that can serve as input
// to Terraform modules.
//
// The conversion process involves:
// 1. Marshaling the proto.Message into JSON using protojson.Marshal.
// 2. Unmarshaling the resulting JSON into a generic map[string]interface{}.
// 3. Recursively processing the map to produce a tfvars-style HCL representation.
//
// The generated tfvars output follows Terraform's variable assignment conventions:
// - Top-level keys map directly to variable names.
// - Values are printed using HCL-compatible syntax:
//   - Strings are quoted: key = "value"
//   - Booleans: key = true/false
//   - Numbers: key = 123 or key = 12.34
//   - Maps are rendered as key = { nested_key = "nested_value" }
//   - Arrays are rendered as key = [ "elem1", "elem2", ... ]
//
// Any nil values are emitted as null. Unsupported types or data structures will result in an error.
//
// Example:
// Given a proto message representing a resource configuration:
//
//	apiVersion: "kubernetes.project.planton/v1"
//	kind: "RedisKubernetes"
//	metadata:
//	  name: "red-one"
//	  labels:
//	    env: "production"
//	spec:
//	  container:
//	    diskSize: "2Gi"
//	    isPersistenceEnabled: true
//	    replicas: 1
//
// This might produce tfvars output:
//
//	apiVersion = "kubernetes.project.planton/v1"
//	kind = "RedisKubernetes"
//	metadata = {
//	  labels = {
//	    env = "production"
//	  }
//	  name = "red-one"
//	}
//	spec = {
//	  container = {
//	    diskSize = "2Gi"
//	    isPersistenceEnabled = true
//	    replicas = 1
//	  }
//	}
//
// Returns:
//   - A string containing the tfvars-formatted representation of the proto message.
//   - An error if the proto cannot be converted to JSON, the JSON cannot be unmarshaled,
//     or if unsupported data types are encountered during conversion.
//
// Typical usage might be:
//
//	tfvarsStr, err := ProtoToTFVars(myProtoMessage)
//	if err != nil {
//	    log.Fatalf("Failed to convert proto to tfvars: %v", err)
//	}
//	ioutil.WriteFile("terraform.tfvars", []byte(tfvarsStr), 0644)
//
// This makes it simple for a pipeline to accept YAML/JSON config, convert it via protobuf (for validation),
// and then produce tfvars for Terraform.
//
// Note: This function assumes the proto message is already validated and contains the expected structure.
func ProtoToTFVars(msg proto.Message) (string, error) {
	// Convert the proto message to JSON
	jsonBytes, err := protojson.Marshal(msg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal proto to json: %w", err)
	}

	// Unmarshal JSON into a generic map
	var data map[string]interface{}
	if err := json.Unmarshal(jsonBytes, &data); err != nil {
		return "", fmt.Errorf("failed to unmarshal json: %w", err)
	}

	// Convert the map into tfvars HCL format
	var buf bytes.Buffer
	if err := writeHCL(&buf, data, 0); err != nil {
		return "", fmt.Errorf("failed to convert map to hcl: %w", err)
	}

	return buf.String(), nil
}

// writeHCL is a helper function that formats a given data structure into
// HCL-compatible syntax suitable for Terraform tfvars. It handles recursion
// into maps and arrays, prints primitives (string, bool, number, null) with
// appropriate quoting or keywords, and uses indentation for readability.
//
// Parameters:
// - buf: a *bytes.Buffer to which the HCL content will be written
// - data: the current data fragment (map, slice, or scalar)
// - indentLevel: the current depth of indentation for pretty-printing
//
// Supported types within data are:
// - map[string]interface{}: rendered as key = { ... } blocks
// - []interface{}: rendered as arrays [ "val", "val2", ... ]
// - string: quoted "string"
// - bool: true/false
// - float64: numeric values as-is
// - nil: rendered as null
//
// Any unsupported type will produce an error.
//
// Example nested formatting:
//
//	key = {
//	  nested_key = "value"
//	  arr_key = [
//	    "elem1",
//	    "elem2",
//	  ]
//	}
//
// This function is only intended for internal use by ProtoToTFVars.
func writeHCL(buf *bytes.Buffer, data interface{}, indentLevel int) error {
	indent := strings.Repeat("  ", indentLevel) // two-space indent

	switch v := data.(type) {

	case map[string]interface{}:
		// If we have a map at this level, we treat each key-value pair as a Terraform variable (or nested attribute).
		//
		// For each entry:
		// - If the value is a map or array, we print the key and open a block with either { ... } or [ ... ].
		// - If the value is a primitive (string, bool, number, null), we print key = value directly.
		for k, val := range v {
			switch val.(type) {

			case map[string]interface{}, []interface{}:
				// Complex types (maps and arrays) require nested formatting:
				//
				// For maps:
				// key = {
				//   nested_key = "value"
				//   ...
				// }
				//
				// For arrays:
				// key = [
				//   "elem1",
				//   "elem2",
				// ]
				buf.WriteString(fmt.Sprintf("%s%s = ", indent, k))
				if m, ok := val.(map[string]interface{}); ok {
					// Nested map block
					buf.WriteString("{\n")
					if err := writeHCL(buf, m, indentLevel+1); err != nil {
						return err
					}
					buf.WriteString(fmt.Sprintf("%s}\n", indent))
				} else if arr, ok := val.([]interface{}); ok {
					// Nested array block
					buf.WriteString("[\n")
					if err := writeHCL(buf, arr, indentLevel+1); err != nil {
						return err
					}
					buf.WriteString(fmt.Sprintf("%s]\n", indent))
				}

			case string:
				// If the value is a string, we must quote it.
				// Example: key = "some string"
				buf.WriteString(fmt.Sprintf("%s%s = %q\n", indent, k, val))

			case bool:
				// Boolean values are printed as true or false.
				// Example: key = true
				buf.WriteString(fmt.Sprintf("%s%s = %t\n", indent, k, val))

			case float64:
				// Numbers (in JSON) are float64. We just print them as-is.
				// Example: key = 100 or key = 3.14
				buf.WriteString(fmt.Sprintf("%s%s = %v\n", indent, k, val))

			case nil:
				// If the value is null, we print null.
				// Terraform supports null values in tfvars.
				// Example: key = null
				buf.WriteString(fmt.Sprintf("%s%s = null\n", indent, k))

			default:
				// If we encounter a type that is not handled above, return an error.
				return fmt.Errorf("unsupported type for key %q: %T", k, val)
			}
		}

	case []interface{}:
		// If we have an array at this level, we treat it as a list of elements.
		// Each element could be a primitive or a nested structure.
		//
		// For each element:
		// - Strings, bools, numbers, null print as "val", true/false, number, or null followed by a comma.
		// - Maps print as a nested block { ... },
		// - Arrays print as nested [ ... ] structures.
		//
		// Example:
		// key = [
		//   "elem1",
		//   true,
		//   123,
		//   {
		//     nested_key = "value"
		//   },
		// ]
		for _, element := range v {
			switch elemVal := element.(type) {

			case string:
				// String elements are quoted and followed by a comma.
				// "value",
				buf.WriteString(fmt.Sprintf("%s%q,\n", indent, elemVal))

			case bool:
				// Boolean elements: true/false,
				buf.WriteString(fmt.Sprintf("%s%t,\n", indent, elemVal))

			case float64:
				// Numeric elements are printed as-is.
				// 123,
				buf.WriteString(fmt.Sprintf("%s%v,\n", indent, elemVal))

			case map[string]interface{}:
				// Map elements are printed as { ... } blocks inside the array.
				// {
				//   nested_key = "value"
				// },
				buf.WriteString(fmt.Sprintf("%s{\n", indent))
				if err := writeHCL(buf, elemVal, indentLevel+1); err != nil {
					return err
				}
				buf.WriteString(fmt.Sprintf("%s},\n", indent))

			case []interface{}:
				// Nested arrays within arrays:
				// [
				//   "inner1",
				//   "inner2",
				// ],
				buf.WriteString(fmt.Sprintf("%s[\n", indent))
				if err := writeHCL(buf, elemVal, indentLevel+1); err != nil {
					return err
				}
				buf.WriteString(fmt.Sprintf("%s],\n", indent))

			case nil:
				// Null element: null,
				buf.WriteString(fmt.Sprintf("%snull,\n", indent))

			default:
				// Unsupported element type in the array.
				return fmt.Errorf("unsupported array element type: %T", elemVal)
			}
		}

	default:
		// The top-level data structure must be a map or an array of supported types.
		// If we get a different type here, it's invalid.
		return fmt.Errorf("top-level must be map[string]interface{}, got %T", data)
	}

	return nil
}
