package tofumodule

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/credential/terraformbackendcredential/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/terraform"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/credentials"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfbackend"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfvars"
	"google.golang.org/protobuf/proto"
	"os"
	"os/exec"
	"path/filepath"
)

func TofuInit(moduleDir string, manifestObject proto.Message,
	backendType terraformbackendcredentialv1.TerraformBackendType,
	backendConfigInput []string,
	stackInputOptions ...credentials.StackInputCredentialOption) error {
	opts := credentials.StackInputCredentialOptions{}
	for _, opt := range stackInputOptions {
		opt(&opts)
	}

	kindName, err := apiresourcekind.ExtractKindFromProto(manifestObject)
	if err != nil {
		return errors.Wrapf(err, "failed to extract kind name from manifest proto")
	}

	tofuModulePath, err := getModulePath(moduleDir, kindName)
	if err != nil {
		return errors.Wrapf(err, "failed to get tofu module directory")
	}

	if err := tfbackend.WriteBackendFile(tofuModulePath, backendType); err != nil {
		return errors.Wrapf(err, "failed to write backend file")
	}

	tfVarsFile := filepath.Join(tofuModulePath, ".terraform", "terraform.tfvars")

	if err = tfvars.WriteVarFile(manifestObject, tfVarsFile); err != nil {
		return errors.Wrapf(err, "failed to write %s file", tfVarsFile)
	}

	tofuCmd := exec.Command(TofuCommand, terraform.TerraformOperationType_init.String(), "--var-file", tfVarsFile)

	// Set the working directory to the repository path
	tofuCmd.Dir = tofuModulePath

	// Set stdin, stdout, and stderr to the current terminal to make it an interactive shell
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stdout = os.Stdout
	tofuCmd.Stderr = os.Stderr

	for _, backendConfig := range backendConfigInput {
		tofuCmd.Args = append(tofuCmd.Args, "--backend-config", backendConfig)
	}

	fmt.Printf("\ntofu module directory: %s \n", tofuModulePath)

	fmt.Printf("\nrunning command: %s \n", tofuCmd.String())

	if err := tofuCmd.Run(); err != nil {
		return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
	}

	return nil
}
