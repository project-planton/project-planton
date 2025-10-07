package pulumistack

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/backendconfig"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule"
)

// Cancel cancels any in-progress operations on a Pulumi stack.
// This is useful when a stack is locked due to a crashed or interrupted operation.
func Cancel(moduleDir, stackFqdn, targetManifestPath string, valueOverrides map[string]string) error {
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

	if err := updateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName); err != nil {
		return errors.Wrapf(err, "failed to update project name in %s/Pulumi.yaml", pulumiModuleRepoPath)
	}

	// Build pulumi cancel command
	args := []string{"cancel", "--stack", finalStackFqdn, "--yes"}

	pulumiCmd := exec.Command("pulumi", args...)

	// Set the working directory to the repository path
	pulumiCmd.Dir = pulumiModuleRepoPath

	// Stream to terminal and also capture output for error classification
	buf := &bytes.Buffer{}
	mwOut := io.MultiWriter(os.Stdout, buf)
	mwErr := io.MultiWriter(os.Stderr, buf)

	// Set stdin, stdout, and stderr to the current terminal
	pulumiCmd.Stdin = os.Stdin
	pulumiCmd.Stdout = mwOut
	pulumiCmd.Stderr = mwErr

	fmt.Printf("\npulumi module directory: %s\n", pulumiModuleRepoPath)
	fmt.Printf("Canceling in-progress operations for stack: %s\n\n", finalStackFqdn)

	if err := pulumiCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to cancel pulumi stack operation")
	}

	fmt.Printf("\n✓ Successfully canceled in-progress operation for stack: %s\n", finalStackFqdn)
	return nil
}
