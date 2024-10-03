package manifestvalidator

import (
	"fmt"
	"github.com/bufbuild/protovalidate-go"
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/pulumimodule"
	"github.com/plantoncloud/project-planton/internal/stackinput"
	"google.golang.org/protobuf/encoding/protojson"
	"os"
	"sigs.k8s.io/yaml"
)

func Validate(manifestPath string) error {
	manifestYamlBytes, err := os.ReadFile(manifestPath)
	if err != nil {
		return errors.Wrap(err, "failed to read manifest file")
	}
	jsonBytes, err := yaml.YAMLToJSON(manifestYamlBytes)
	if err != nil {
		return errors.Wrap(err, "failed to load yaml to json")
	}

	kindName, err := stackinput.ExtractKindFromTargetManifest(manifestPath)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind from %s stack input yaml", manifestPath)
	}

	manifest := DeploymentComponentMap[DeploymentComponent(pulumimodule.ConvertKindName(kindName))]

	if manifest == nil {
		return errors.Errorf("deployment-component does not contain %s", pulumimodule.ConvertKindName(kindName))
	}

	if err := protojson.Unmarshal(jsonBytes, manifest); err != nil {
		return errors.Wrap(err, "failed to load json into proto message")
	}

	spec, err := ExtractSpec(manifest)
	if err != nil {
		return errors.Wrap(err, "failed to extract spec from manifest")
	}

	v, err := protovalidate.New(
		protovalidate.WithDisableLazy(true),
		protovalidate.WithMessages(spec),
	)
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	return v.Validate(spec)
}
