package stackinput

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"gopkg.in/yaml.v3"
)

// BuildStackInputYaml reads two YAML files, combines their contents,
// and returns a new YAML string with "target" and all the credential keys.
func BuildStackInputYaml(manifestObject proto.Message,
	providerConfigOptions stackinputproviderconfig.StackInputProviderConfigOptions) (string, error) {

	var targetContentMap map[string]interface{}
	targetContent, err := protojson.Marshal(manifestObject)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal manifest object to JSON")
	}

	err = yaml.Unmarshal(targetContent, &targetContentMap)
	if err != nil {
		return "", errors.Wrapf(err, "failed to unmarshal target manifest file")
	}

	stackInputContentMap := map[string]interface{}{
		"target": targetContentMap,
	}

	stackInputContentMap, err = addProviderConfigs(stackInputContentMap, providerConfigOptions)
	if err != nil {
		return "", errors.Wrapf(err, "failed to add credentials to stack-input yaml")
	}

	// Convert map to yaml.Node to control formatting
	var rootNode yaml.Node
	err = rootNode.Encode(stackInputContentMap)
	if err != nil {
		return "", errors.Wrap(err, "failed to encode stack-input to yaml node")
	}

	// Force double-quoted style for serviceAccountKeyBase64 field to prevent line folding
	forceDoubleQuotedStyleForBase64(&rootNode)

	// Marshal the node to YAML
	finalStackInputYaml, err := yaml.Marshal(&rootNode)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal final stack-input yaml")
	}

	return string(finalStackInputYaml), nil
}

// forceDoubleQuotedStyleForBase64 recursively finds and formats serviceAccountKeyBase64 fields
func forceDoubleQuotedStyleForBase64(node *yaml.Node) {
	if node == nil {
		return
	}

	// Handle mapping nodes (objects)
	if node.Kind == yaml.MappingNode {
		for i := 0; i < len(node.Content); i += 2 {
			if i+1 < len(node.Content) {
				keyNode := node.Content[i]
				valueNode := node.Content[i+1]

				// If key is "serviceAccountKeyBase64" (camelCase JSON name), force double-quoted style on the value
				if keyNode.Value == "serviceAccountKeyBase64" && valueNode.Kind == yaml.ScalarNode {
					valueNode.Style = yaml.DoubleQuotedStyle
				}

				// Recursively process nested structures
				forceDoubleQuotedStyleForBase64(valueNode)
			}
		}
	}

	// Handle sequence nodes (arrays)
	if node.Kind == yaml.SequenceNode {
		for _, child := range node.Content {
			forceDoubleQuotedStyleForBase64(child)
		}
	}
}
