package stackinput

import (
	"fmt"
	"os"

	"buf.build/go/protovalidate"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/stackinput/fieldsextractor"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"
)

const (
	PulumiConfigKey   = "planton-cloud:stack-input"
	FilePathEnvVar    = "STACK_INPUT_FILE_PATH"
	YamlContentEnvVar = "STACK_INPUT_YAML"
)

func LoadStackInput(ctx *pulumi.Context, stackInput proto.Message) error {
	stackInputString, ok := ctx.GetConfig(PulumiConfigKey)
	var jsonBytes, stackInputYamlBytes []byte
	var err error

	if !ok {
		yamlContent := os.Getenv(YamlContentEnvVar)
		if yamlContent != "" {
			stackInputYamlBytes = []byte(yamlContent)
			fmt.Printf("DEBUG: Loaded YAML from env var (%d bytes)\n", len(stackInputYamlBytes))
		} else {
			stackInputFilePath := os.Getenv(FilePathEnvVar)
			if stackInputFilePath == "" {
				return errors.Errorf("stack-input not found in pulumi config %s or in %s environment variable",
					PulumiConfigKey, FilePathEnvVar)
			}
			stackInputYamlBytes, err = os.ReadFile(stackInputFilePath)
			if err != nil {
				return errors.Wrap(err, "failed to read input file")
			}
			fmt.Printf("DEBUG: Loaded YAML from file %s (%d bytes)\n", stackInputFilePath, len(stackInputYamlBytes))
		}

		// DEBUG: Log YAML content before conversion
		fmt.Printf("DEBUG: YAML content before YAMLToJSON:\n%s\n", string(stackInputYamlBytes))
		fmt.Printf("DEBUG: YAML bytes length: %d\n", len(stackInputYamlBytes))

		jsonBytes, err = yaml.YAMLToJSON(stackInputYamlBytes)
		if err != nil {
			fmt.Printf("ERROR: YAMLToJSON failed: %v\n", err)
			return errors.Wrap(err, "failed to load yaml to json")
		}

		// DEBUG: Log JSON content after conversion
		fmt.Printf("DEBUG: JSON content after YAMLToJSON (%d bytes):\n%s\n", len(jsonBytes), string(jsonBytes))
	} else {
		fmt.Printf("DEBUG: Loaded YAML from Pulumi config (%d bytes)\n", len(stackInputString))
		fmt.Printf("DEBUG: YAML content from config:\n%s\n", stackInputString)

		jsonBytes, err = yaml.YAMLToJSON([]byte(stackInputString))
		if err != nil {
			fmt.Printf("ERROR: YAMLToJSON failed: %v\n", err)
			return errors.Wrap(err, "failed to load yaml to json")
		}
		fmt.Printf("DEBUG: JSON content after YAMLToJSON (%d bytes):\n%s\n", len(jsonBytes), string(jsonBytes))
	}

	// DEBUG: Log before proto unmarshal
	fmt.Printf("DEBUG: About to unmarshal JSON into proto message\n")
	if len(jsonBytes) > 1000 {
		fmt.Printf("DEBUG: JSON to unmarshal (first 500 chars): %s\n", string(jsonBytes[:500]))
		fmt.Printf("DEBUG: JSON to unmarshal (last 500 chars): %s\n", string(jsonBytes[len(jsonBytes)-500:]))
	} else {
		fmt.Printf("DEBUG: JSON to unmarshal (%d bytes): %s\n", len(jsonBytes), string(jsonBytes))
	}

	if err := protojson.Unmarshal(jsonBytes, stackInput); err != nil {
		fmt.Printf("ERROR: protojson.Unmarshal failed: %v\n", err)
		fmt.Printf("ERROR: JSON that failed to unmarshal (%d bytes):\n%s\n", len(jsonBytes), string(jsonBytes))
		return errors.Wrap(err, "failed to load json into proto message")
	}

	fmt.Printf("DEBUG: Successfully unmarshaled proto message\n")

	targetSpec, err := fieldsextractor.ExtractApiResourceSpecField(stackInput)
	if err != nil {
		return errors.Wrap(err, "failed to extract api resource spec field")
	}

	v, err := protovalidate.New(
		protovalidate.WithDisableLazy(),
		protovalidate.WithMessages((*targetSpec).Interface()),
	)
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate((*targetSpec).Interface()); err != nil {
		return errors.Errorf("%s", err)
	}
	return nil
}
