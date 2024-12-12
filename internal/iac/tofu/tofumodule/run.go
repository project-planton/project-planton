package tofumodule

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/tofu"
	"github.com/project-planton/project-planton/internal/iac/pulumi/stackinput/credentials"
	"github.com/project-planton/project-planton/internal/iac/tofu/tfvars"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/ulidgen"
	"os"
	"os/exec"
	"path/filepath"
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

	tofuModulePath, err := getModulePath(moduleDir, kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get pulumi-module directory")
	}

	tfvarsString, err := tfvars.ProtoToTFVars(manifestObject)
	if err != nil {
		return errors.Wrap(err, "failed to convert manifest proto to tfvars")
	}

	tfVarsFile, err := writeVarFile(tfvarsString)
	if err != nil {
		return errors.Wrap(err, "failed to write tfvars file")
	}

	op := tofuOperation.String()

	tofuCmd := exec.Command(TofuCommand, op, "--var-file", tfVarsFile)

	// Set the working directory to the repository path
	tofuCmd.Dir = tofuModulePath

	// Set stdin, stdout, and stderr to the current terminal to make it an interactive shell
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stdout = os.Stdout
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("\ntofu module directory: %s \n", tofuModulePath)

	if err := tofuCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute tofu command %s", op)
	}

	return nil
}

func writeVarFile(tfvarsString string) (string, error) {
	tofuWorkspaceDir, err := getWorkspaceDir()
	if err != nil {
		return "", errors.Wrapf(err, "failed to get tofu workspace directory")
	}
	return filepath.Join(tofuWorkspaceDir, ulidgen.NewGenerator().Generate().String(), "terraform.tfvars"), nil
}
