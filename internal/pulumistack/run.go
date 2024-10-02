package pulumistack

import (
	"buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/shared/pulumi"
	"github.com/pkg/errors"
	"github.com/plantoncloud/project-planton/internal/pulumimodule"
	"github.com/plantoncloud/project-planton/internal/stackinput"
	"github.com/plantoncloud/project-planton/internal/stackinput/credentials"
	cliworkspace "github.com/plantoncloud/project-planton/internal/workspace"
	"os"
	"os/exec"
)

func Run(stackFqdn, targetManifestPath string, pulumiOperation pulumi.PulumiOperationType,
	isUpdatePreview bool, stackInputOptions ...credentials.StackInputCredentialOption) error {
	opts := credentials.StackInputCredentialOptions{}
	for _, opt := range stackInputOptions {
		opt(&opts)
	}

	kindName, err := stackinput.ExtractKindFromTargetManifest(targetManifestPath)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind from %s stack input yaml", targetManifestPath)
	}

	cloneUrl, err := pulumimodule.GetCloneUrl(kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get clone url for %s kind", kindName)
	}

	stackWorkspaceDir, err := cliworkspace.GetWorkspaceDir(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to get %s stack worspace directory", stackFqdn)
	}

	pulumiProjectName, err := ExtractProjectName(stackFqdn)
	if err != nil {
		return errors.Wrapf(err, "failed to extract project name from %s stack fqdn", stackFqdn)
	}

	stackInputYamlContent, err := stackinput.BuildStackInputYaml(targetManifestPath, opts)
	if err != nil {
		return errors.Wrap(err, "failed to build stack input yaml")
	}

	gitRepoName, err := extractGitRepoName(cloneUrl)
	if err != nil {
		return errors.Wrapf(err, "failed to extract git repo name from %s", cloneUrl)
	}

	// Check if the cloned repository directory already exists
	pulumiModuleRepoPath := stackWorkspaceDir + "/" + gitRepoName
	if _, err := os.Stat(pulumiModuleRepoPath); os.IsNotExist(err) {
		gitCloneCommand := exec.Command("git", "clone", cloneUrl, pulumiModuleRepoPath)
		gitCloneCommand.Stdout = os.Stdout
		gitCloneCommand.Stderr = os.Stderr
		if err := gitCloneCommand.Run(); err != nil {
			return errors.Wrapf(err, "failed to clone repository from %s to %s", cloneUrl, stackWorkspaceDir)
		}
	}

	if err := updateProjectNameInPulumiYaml(pulumiModuleRepoPath, pulumiProjectName); err != nil {
		return errors.Wrapf(err, "failed to update project name in %s/Pulumi.yaml", pulumiModuleRepoPath)
	}

	op := pulumiOperation.String()
	if isUpdatePreview {
		op = "preview"
	}

	pulumiCmd := exec.Command("pulumi", op, "--stack", stackFqdn, "--yes")

	// Set the STACK_INPUT_YAML environment variable
	pulumiCmd.Env = append(os.Environ(), "STACK_INPUT_YAML="+stackInputYamlContent)

	// Set the working directory to the repository path
	pulumiCmd.Dir = pulumiModuleRepoPath

	// Set stdin, stdout, and stderr to the current terminal to make it an interactive shell
	pulumiCmd.Stdin = os.Stdin
	pulumiCmd.Stdout = os.Stdout
	pulumiCmd.Stderr = os.Stderr

	if err := pulumiCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute pulumi command %s", op)
	}

	return nil
}
