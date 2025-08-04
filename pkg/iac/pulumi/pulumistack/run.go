package pulumistack

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/pulumi"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
	"os"
	"os/exec"
)

func Run(moduleDir, stackFqdn, targetManifestPath string, pulumiOperation pulumi.PulumiOperationType,
	isUpdatePreview bool, valueOverrides map[string]string,
	credentialOptions ...stackinputcredentials.StackInputCredentialOption) error {
	opts := stackinputcredentials.StackInputCredentialOptions{}
	for _, opt := range credentialOptions {
		opt(&opts)
	}

	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to override values in target manifest file")
	}

	kindName, err := crkreflect.ExtractKindFromProto(manifestObject)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind name from manifest proto")
	}

	pulumiModuleRepoPath, err := pulumimodule.GetPath(moduleDir, stackFqdn, kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get pulumi-module directory")
	}

	pulumiProjectName, err := ExtractProjectName(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to extract project name from %s stack fqdn", stackFqdn)
	}

	stackInputYamlContent, err := stackinput.BuildStackInputYaml(manifestObject, opts)
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
