package pulumistack

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/pulumi"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/backendconfig"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
)

func Run(moduleDir, stackFqdn, targetManifestPath string, pulumiOperation pulumi.PulumiOperationType,
	isUpdatePreview bool, isAutoApprove bool, valueOverrides map[string]string,
	credentialOptions ...stackinputcredentials.StackInputCredentialOption) error {
	opts := stackinputcredentials.StackInputCredentialOptions{}
	for _, opt := range credentialOptions {
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

	// Build pulumi command with optional flags for non-interactive runs
	args := []string{op, "--stack", finalStackFqdn, "--non-interactive"}
	if isAutoApprove {
		args = append(args, "--yes")
		// For 'pulumi up', skip preview to avoid TTY prompts in CI/non-interactive shells
		if op == "up" {
			args = append(args, "--skip-preview")
		}
	}

	pulumiCmd := exec.Command("pulumi", args...)

	// Set the STACK_INPUT_YAML environment variable
	pulumiCmd.Env = append(os.Environ(), "STACK_INPUT_YAML="+stackInputYamlContent)

	// Set the working directory to the repository path
	pulumiCmd.Dir = pulumiModuleRepoPath

	// Stream to terminal and also capture output for error classification
	buf := &bytes.Buffer{}
	mwOut := io.MultiWriter(os.Stdout, buf)
	mwErr := io.MultiWriter(os.Stderr, buf)

	// Set stdin, stdout, and stderr to the current terminal to make it an interactive shell
	pulumiCmd.Stdin = os.Stdin
	pulumiCmd.Stdout = mwOut
	pulumiCmd.Stderr = mwErr

	fmt.Printf("\npulumi module directory: %s \n", pulumiModuleRepoPath)

	if err := pulumiCmd.Run(); err != nil {
		// For preview/update/refresh/destroy, Pulumi can return non-zero even when
		// operation printed a valid plan/result. If no 'error:' diagnostics exist,
		// treat it as success to avoid false negatives in non-interactive mode.
		out := buf.String()
		if !strings.Contains(out, "error:") {
			return nil
		}
		return errors.Wrapf(err, "failed to execute pulumi command %s", op)
	}

	return nil
}
