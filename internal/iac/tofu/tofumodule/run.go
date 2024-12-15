package tofumodule

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/tofu"
	"github.com/project-planton/project-planton/internal/iac/stackinput/credentials"
	"github.com/project-planton/project-planton/internal/iac/tofu/tfvars"
	"github.com/project-planton/project-planton/internal/manifest"
	"os"
	"os/exec"
	"path/filepath"
)

const TofuCommand = "tofu"

func RunOperation(moduleDir, targetManifestPath string, tofuOperation tofu.TofuOperationType,
	valueOverrides map[string]string,
	isAutoApprove bool,
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
		return errors.Wrapf(err, "failed to get tofu module directory")
	}

	tfVarsFile := filepath.Join(tofuModulePath, ".terraform", "terraform.tfvars")

	if err = tfvars.WriteVarFile(manifestObject, tfVarsFile); err != nil {
		return errors.Wrapf(err, "failed to write %s file", tfVarsFile)
	}

	op := tofuOperation.String()

	tofuCmd := exec.Command(TofuCommand, op, "--var-file", tfVarsFile)

	if (tofuOperation == tofu.TofuOperationType_apply ||
		tofuOperation == tofu.TofuOperationType_destroy) && isAutoApprove {
		tofuCmd.Args = append(tofuCmd.Args, "--auto-approve")
	}

	// Set the working directory to the repository path
	tofuCmd.Dir = tofuModulePath

	// Set stdin, stdout, and stderr to the current terminal to make it an interactive shell
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stdout = os.Stdout
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("\ntofu module directory: %s \n", tofuModulePath)

	fmt.Printf("\nrunning command: %s \n", tofuCmd.String())

	if err := tofuCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
	}

	return nil
}
