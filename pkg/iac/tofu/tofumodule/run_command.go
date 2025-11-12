package tofumodule

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/internal/cli/workspace"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
	"github.com/project-planton/project-planton/pkg/iac/tofu/backendconfig"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfbackend"
	log "github.com/sirupsen/logrus"
)

func RunCommand(inputModuleDir, targetManifestPath string,
	terraformOperation terraform.TerraformOperationType,
	valueOverrides map[string]string,
	isAutoApprove, isDestroyPlan bool,
	providerConfigOptions ...stackinputproviderconfig.StackInputProviderConfigOption) error {

	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to override values in target manifest file")
	}

	// Extract backend configuration from manifest labels (optional)
	var backendType terraform.TerraformBackendType = terraform.TerraformBackendType_local
	var backendConfigArgs []string

	tofuBackendConfig, err := backendconfig.ExtractFromManifest(manifestObject)
	if err != nil {
		// Log but don't fail - backend config is optional
		log.Debugf("Could not extract Terraform backend config from manifest labels: %v", err)
	}

	if tofuBackendConfig != nil {
		log.Infof("Using Terraform backend from manifest labels: type=%s, object=%s",
			tofuBackendConfig.BackendType, tofuBackendConfig.BackendObject)

		// Convert backend type string to enum
		backendType = tfbackend.BackendTypeFromString(tofuBackendConfig.BackendType)
		if backendType == terraform.TerraformBackendType_terraform_backend_type_unspecified {
			return errors.Errorf("unsupported backend type from manifest labels: %s", tofuBackendConfig.BackendType)
		}

		// Build backend config arguments based on backend type
		backendConfigArgs = buildBackendConfigArgs(tofuBackendConfig)
	} else {
		log.Debug("No Terraform backend config in manifest labels, using default local backend")
	}

	kindName, err := crkreflect.ExtractKindFromProto(manifestObject)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind name from manifest proto")
	}

	tofuModulePath, err := GetModulePath(inputModuleDir, kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get tofu module directory")
	}

	// Gather credential options
	opts := stackinputproviderconfig.StackInputProviderConfigOptions{}
	for _, opt := range providerConfigOptions {
		opt(&opts)
	}

	stackInputYaml, err := stackinput.BuildStackInputYaml(manifestObject, opts)
	if err != nil {
		return errors.Wrap(err, "failed to build stack input yaml")
	}

	workspaceDir, err := workspace.GetWorkspaceDir()
	if err != nil {
		return errors.Wrap(err, "failed to get workspace directory")
	}

	providerConfigEnvVars, err := GetProviderConfigEnvVars(stackInputYaml, workspaceDir)
	if err != nil {
		return errors.Wrap(err, "failed to get provider config env vars")
	}

	// Initialize tofu with backend configuration
	// This should happen before any operation to ensure backend is properly configured
	err = TofuInit(tofuModulePath, manifestObject, backendType, backendConfigArgs,
		providerConfigEnvVars, false, nil)
	if err != nil {
		return errors.Wrap(err, "failed to initialize tofu module")
	}

	err = RunOperation(tofuModulePath, terraformOperation,
		isAutoApprove,
		isDestroyPlan, manifestObject,
		providerConfigEnvVars, false, nil)
	if err != nil {
		return errors.Wrapf(err, "failed to run tofu operation")
	}
	return nil
}

// buildBackendConfigArgs builds backend configuration arguments based on backend type
func buildBackendConfigArgs(config *backendconfig.TofuBackendConfig) []string {
	var args []string

	switch config.BackendType {
	case "s3":
		// For S3: parse "bucket-name/path/to/state"
		parts := strings.SplitN(config.BackendObject, "/", 2)
		if len(parts) >= 1 {
			args = append(args, fmt.Sprintf("bucket=%s", parts[0]))
		}
		if len(parts) >= 2 {
			args = append(args, fmt.Sprintf("key=%s", parts[1]))
		}

	case "gcs":
		// For GCS: parse "bucket-name/path/to/state"
		parts := strings.SplitN(config.BackendObject, "/", 2)
		if len(parts) >= 1 {
			args = append(args, fmt.Sprintf("bucket=%s", parts[0]))
		}
		if len(parts) >= 2 {
			args = append(args, fmt.Sprintf("prefix=%s", parts[1]))
		}

	case "azurerm":
		// For Azure: parse "container-name/path/to/state"
		parts := strings.SplitN(config.BackendObject, "/", 2)
		if len(parts) >= 1 {
			args = append(args, fmt.Sprintf("container_name=%s", parts[0]))
		}
		if len(parts) >= 2 {
			args = append(args, fmt.Sprintf("key=%s", parts[1]))
		}

	case "local":
		// Local backend doesn't need config args
		// The path is handled by terraform itself
	}

	return args
}
