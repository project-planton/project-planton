package pulumistack

import (
	"buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/shared/pulumi"
	"fmt"
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/pulumimodule"
	"github.com/plantoncloud/project-planton/internal/stackinput"
	"github.com/plantoncloud/project-planton/internal/stackinput/credentials"
	"os"
	"os/exec"
)

func Run(moduleDir, stackFqdn, targetManifestPath string, pulumiOperation pulumi.PulumiOperationType,
	isUpdatePreview bool, stackInputOptions ...credentials.StackInputCredentialOption) error {
	opts := credentials.StackInputCredentialOptions{}
	for _, opt := range stackInputOptions {
		opt(&opts)
	}

	pulumiModuleRepoPath, err := pulumimodule.GetPath(moduleDir, stackFqdn, targetManifestPath)
	if err != nil {
		return errors.Wrapf(err, "failed to get pulumi-module directory")
	}

	pulumiProjectName, err := ExtractProjectName(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to extract project name from %s stack fqdn", stackFqdn)
	}

	stackInputYamlContent, err := stackinput.BuildStackInputYaml(targetManifestPath, opts)
	if err != nil {
		return errors.Wrap(err, "failed to build stack input yaml")
	}

	if err := updateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName); err != nil {
		return errors.Wrapf(err, "failed to update project name in %s/Pulumi.yaml", pulumiModuleRepoPath)
	}

	op := pulumiOperation.String()
	if isUpdatePreview {
		op = "preview"
	}

	pulumiCmd := exec.Command("pulumi", op, "--stack", stackFqdn)

	// Set the STACK_INPUT_YAML environment variable
	pulumiCmd.Env = append(os.Environ(), "STACK_INPUT_YAML="+stackInputYamlContent)

	// Set the working directory to the repository path
	pulumiCmd.Dir = pulumiModuleRepoPath

	// Set stdin, stdout, and stderr to the current terminal to make it an interactive shell
	pulumiCmd.Stdin = os.Stdin
	pulumiCmd.Stdout = os.Stdout
	pulumiCmd.Stderr = os.Stderr

	fmt.Printf("\npulumi module directory: %s \n", pulumiModuleRepoPath)

	if err := pulumiCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute pulumi command %s", op)
	}

	return nil
}
