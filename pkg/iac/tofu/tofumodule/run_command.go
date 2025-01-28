package tofumodule

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func RunCommand(inputModuleDir, targetManifestPath string,
	terraformOperation terraform.TerraformOperationType,
	valueOverrides map[string]string,
	isAutoApprove, isDestroyPlan bool,
	credentialOptions ...stackinputcredentials.StackInputCredentialOption) error {

	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to override values in target manifest file")
	}

	kindName, err := apiresourcekind.ExtractKindFromProto(manifestObject)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind name from manifest proto")
	}

	tofuModulePath, err := GetModulePath(inputModuleDir, kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get tofu module directory")
	}

	// Gather credential options (currently unused, but left for future usage)
	opts := stackinputcredentials.StackInputCredentialOptions{}
	for _, opt := range credentialOptions {
		opt(&opts)
	}

	stackInputYaml, err := stackinput.BuildStackInputYaml(manifestObject, opts)
	if err != nil {
		return errors.Wrap(err, "failed to build stack input yaml")
	}

	credentialEnvVars, err := GetCredentialEnvVars(stackInputYaml)
	if err != nil {
		return errors.Wrap(err, "failed to get credential env vars")
	}

	err = RunOperation(tofuModulePath, terraformOperation,
		isAutoApprove,
		isDestroyPlan, manifestObject,
		credentialEnvVars, false, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to run tofu operation")
	}
	return nil
}
