package tofumodule

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/tofu"
	"github.com/project-planton/project-planton/internal/iac/pulumi/stackinput/credentials"
	"github.com/project-planton/project-planton/internal/manifest"
	"os"
	"os/exec"
)

const TofuCommand = "tofu"

func Run(moduleDir, targetManifestPath string, tofuOperation tofu.TofuOperationType, valueOverrides map[string]string,
	stackInputOptions ...credentials.StackInputCredentialOption) error {
	opts := credentials.StackInputCredentialOptions{}
	for _, opt := range stackInputOptions {
		opt(&opts)
	}

	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		return errors.Wrapf(err, "failed to override values in target manifest file")
	}

	kindName, err := manifest.ExtractKindFromProto(manifestObject)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind name from manifest proto")
	}

	//todo: update this to get the tofu module path

	tofuModuleRepoPath, err := getModulePath(moduleDir, "stackFqdn", kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get pulumi-module directory")
	}

	//todo: replce this logic with generating tfvars and saving to to a file and passing the location of the file to tofu command

	//stackInputYamlContent, err := stackinput.BuildStackInputYaml(manifestObject, opts)
	//if err != nil {
	//	return errors.Wrap(err, "failed to build stack input yaml")
	//}

	tfVarsFile := ""

	op := tofuOperation.String()

	tofuCmd := exec.Command(TofuCommand, op, "--var-file", tfVarsFile)

	// Set the working directory to the repository path
	tofuCmd.Dir = tofuModuleRepoPath

	// Set stdin, stdout, and stderr to the current terminal to make it an interactive shell
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stdout = os.Stdout
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("\ntofu module directory: %s \n", tofuModuleRepoPath)

	if err := tofuCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute tofu command %s", op)
	}

	return nil
}
