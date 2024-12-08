package tfvars

import (
	"bytes"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"strings"
)

func ProtoToTerraformTFVars(msg proto.Message) (string, error) {
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

// writeHCL writes a map[string]interface{} as HCL variables.
// It recursively handles maps and slices.
// indentLevel controls indentation for pretty printing (optional).
func writeHCL(buf *bytes.Buffer, data interface{}, indentLevel int) error {
	indent := strings.Repeat("  ", indentLevel) // two-space indent

	switch v := data.(type) {
	case map[string]interface{}:
		// For top-level maps, we consider each key as a Terraform variable.
		// Each key = value pair is on its own line.
		// If value is complex (map or list), we nest.
		for k, val := range v {
			switch val.(type) {
			case map[string]interface{}, []interface{}:
				// Complex type: key = { ... } or key = [ ... ]
				// For maps:
				// key = {
				//   nested_key = "value"
				// }
				//
				// For slices:
				// key = [
				//   "value1",
				//   "value2"
				// ]

				buf.WriteString(fmt.Sprintf("%s%s = ", indent, k))
				// Determine if val is a map or slice
				if m, ok := val.(map[string]interface{}); ok {
					buf.WriteString("{\n")
					if err := writeHCL(buf, m, indentLevel+1); err != nil {
						return err
					}
					buf.WriteString(fmt.Sprintf("%s}\n", indent))
				} else if arr, ok := val.([]interface{}); ok {
					buf.WriteString("[\n")
					if err := writeHCL(buf, arr, indentLevel+1); err != nil {
						return err
					}
					buf.WriteString(fmt.Sprintf("%s]\n", indent))
				}
			case string:
				buf.WriteString(fmt.Sprintf("%s%s = %q\n", indent, k, val))
			case bool:
				buf.WriteString(fmt.Sprintf("%s%s = %t\n", indent, k, val))
			case float64:
				// JSON numbers become float64. If needed, you can try to detect if it's int-like.
				// For now, just print as a number.
				buf.WriteString(fmt.Sprintf("%s%s = %v\n", indent, k, val))
			case nil:
				// Terraform doesn't have a true null in tfvars. Usually omitted or use "" or zero values.
				// Decide how you want to handle nil:
				buf.WriteString(fmt.Sprintf("%s%s = null\n", indent, k))
			default:
				return fmt.Errorf("unsupported type for key %q: %T", k, val)
			}
		}
	case []interface{}:
		// Array elements, each line is an element
		for _, element := range v {
			switch elemVal := element.(type) {
			case string:
				buf.WriteString(fmt.Sprintf("%s%q,\n", indent, elemVal))
			case bool:
				buf.WriteString(fmt.Sprintf("%s%t,\n", indent, elemVal))
			case float64:
				buf.WriteString(fmt.Sprintf("%s%v,\n", indent, elemVal))
			case map[string]interface{}:
				buf.WriteString(fmt.Sprintf("%s{\n", indent))
				if err := writeHCL(buf, elemVal, indentLevel+1); err != nil {
					return err
				}
				buf.WriteString(fmt.Sprintf("%s},\n", indent))
			case []interface{}:
				buf.WriteString(fmt.Sprintf("%s[\n", indent))
				if err := writeHCL(buf, elemVal, indentLevel+1); err != nil {
					return err
				}
				buf.WriteString(fmt.Sprintf("%s],\n", indent))
			case nil:
				buf.WriteString(fmt.Sprintf("%snull,\n", indent))
			default:
				return fmt.Errorf("unsupported array element type: %T", elemVal)
			}
		}
	default:
		return fmt.Errorf("top-level must be map[string]interface{}, got %T", data)
	}

	return nil
}
