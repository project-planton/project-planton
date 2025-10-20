package pulumistack

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/pulumi"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/backendconfig"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
)

func Run(moduleDir, stackFqdn, targetManifestPath string, pulumiOperation pulumi.PulumiOperationType,
	isUpdatePreview bool, isAutoApprove bool, valueOverrides map[string]string, showDiff bool,
	providerConfigOptions ...stackinputproviderconfig.StackInputProviderConfigOption) error {
	opts := stackinputproviderconfig.StackInputProviderConfigOptions{}
	for _, opt := range providerConfigOptions {
		opt(&opts)
	}

	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to override values in target manifest file")
	}

	// Try to extract backend configuration from manifest labels
	// If found, use it instead of the provided stackFqdn
	finalStackFqdn := stackFqdn
	if manifestBackendConfig, err := backendconfig.ExtractFromManifest(manifestObject); err == nil && manifestBackendConfig != nil {
		if manifestBackendConfig.StackFqdn != "" {
			fmt.Printf("Using Pulumi stack from manifest labels: %s\n\n", manifestBackendConfig.StackFqdn)
			finalStackFqdn = manifestBackendConfig.StackFqdn
		}
	}

	// Validate that we have a stack FQDN
	if finalStackFqdn == "" {
		return errors.New("Pulumi stack FQDN is required. Provide it via --stack flag or set pulumi.project-planton.org/stack.fqdn label in manifest")
	}

	kindName, err := crkreflect.ExtractKindFromProto(manifestObject)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind name from manifest proto")
	}

	pulumiModuleRepoPath, err := pulumimodule.GetPath(moduleDir, finalStackFqdn, kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get pulumi-module directory")
	}

	pulumiProjectName, err := ExtractProjectName(finalStackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to extract project name from %s stack fqdn", finalStackFqdn)
	}

	stackInputYamlContent, err := stackinput.BuildStackInputYaml(manifestObject, opts)
	if err != nil {
		return errors.Wrap(err, "failed to build stack input yaml")
	}

	if err := updateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName); err != nil {
		return errors.Wrapf(err, "failed to update project name in %s/Pulumi.yaml", pulumiModuleRepoPath)
	}

	// Map to Pulumi CLI verbs
	op := pulumiOperation.String()
	switch pulumiOperation {
	case pulumi.PulumiOperationType_update:
		op = "up"
	case pulumi.PulumiOperationType_refresh:
		op = "refresh"
	case pulumi.PulumiOperationType_destroy:
		op = "destroy"
	}
	if isUpdatePreview {
		op = "preview"
	}

	// Build pulumi command with optional flags
	args := []string{op, "--stack", finalStackFqdn}
	if isAutoApprove {
		args = append(args, "--yes")
		// For 'pulumi up', skip preview to avoid TTY prompts in CI/non-interactive shells
		if op == "up" {
			args = append(args, "--skip-preview")
		}
	}
	if showDiff {
		args = append(args, "--diff")
	}

	pulumiCmd := exec.Command("pulumi", args...)

	// Set the STACK_INPUT_YAML environment variable
	pulumiCmd.Env = append(os.Environ(), "STACK_INPUT_YAML="+stackInputYamlContent)

	// Set the working directory to the repository path
	pulumiCmd.Dir = pulumiModuleRepoPath

	// Set stdin, stdout, and stderr directly to the terminal for interactive output
	// This allows Pulumi to detect TTY and use the interactive tree view
	pulumiCmd.Stdin = os.Stdin
	pulumiCmd.Stdout = os.Stdout
	pulumiCmd.Stderr = os.Stderr

	fmt.Printf("\npulumi module directory: %s \n", pulumiModuleRepoPath)

	if err := pulumiCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute pulumi command %s", op)
	}

	return nil
}
