package pulumistack

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared/iac/pulumi"
	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/manifest"
	"github.com/plantonhq/project-planton/pkg/crkreflect"
	"github.com/plantonhq/project-planton/pkg/iac/pulumi/backendconfig"
	"github.com/plantonhq/project-planton/pkg/iac/pulumi/pulumimodule"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
	log "github.com/sirupsen/logrus"
)

func Run(moduleDir, stackFqdn, targetManifestPath string, pulumiOperation pulumi.PulumiOperationType,
	isUpdatePreview bool, isAutoApprove bool, valueOverrides map[string]string, showDiff bool, moduleVersion string, noCleanup bool,
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
			cyan := color.New(color.FgCyan).SprintFunc()
			fmt.Printf("\nDetected Stack from Labels: %s\n\n", cyan(manifestBackendConfig.StackFqdn))
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

	pathResult, err := pulumimodule.GetPath(moduleDir, finalStackFqdn, kindName, moduleVersion, noCleanup)
	if err != nil {
		return errors.Wrapf(err, "failed to get pulumi-module directory")
	}

	// Setup cleanup to run after execution
	if pathResult.ShouldCleanup {
		defer func() {
			if cleanupErr := pathResult.CleanupFunc(); cleanupErr != nil {
				fmt.Printf("Warning: failed to cleanup workspace copy: %v\n", cleanupErr)
			}
		}()
	}

	pulumiModuleRepoPath := pathResult.ModulePath

	pulumiProjectName, err := ExtractProjectName(finalStackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to extract project name from %s stack fqdn", finalStackFqdn)
	}

	stackInputYamlContent, err := stackinput.BuildStackInputYaml(manifestObject, opts)
	if err != nil {
		return errors.Wrap(err, "failed to build stack input yaml")
	}

	// Update project name in Pulumi.yaml
	// For binary mode, we regenerate the Pulumi.yaml with the correct project name
	if err := UpdateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName); err != nil {
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

	// Log execution mode and directory info (debug level only)
	if pathResult.UseBinary {
		log.Debugf("execution mode: binary (no compilation)")
		log.Debugf("binary path: %s", pathResult.BinaryPath)
	} else {
		log.Debugf("execution mode: source (compilation required)")
	}
	log.Debugf("workspace directory: %s", pulumiModuleRepoPath)
	fmt.Println()

	// Print handoff message after all setup is complete
	cliprint.PrintHandoff("Pulumi")

	if err := pulumiCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute pulumi command %s", op)
	}

	return nil
}
