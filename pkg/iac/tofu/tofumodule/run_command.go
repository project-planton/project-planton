package tofumodule

import (
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/credentials"
)

func RunCommand(inputModuleDir, targetManifestPath string, terraformOperation terraform.TerraformOperationType,
	valueOverrides map[string]string,
	isAutoApprove bool,
	stackInputOptions ...credentials.StackInputCredentialOption) error {
	opts := credentials.StackInputCredentialOptions{}
	for _, opt := range stackInputOptions {
		opt(&opts)
	}

	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to override values in target manifest file")
	}

	err = RunOperation(inputModuleDir, terraformOperation, isAutoApprove, manifestObject, stackInputOptions...)
	if err != nil {
		return errors.Wrapf(err, "failed to run tofu operation")
	}
	return nil
}
