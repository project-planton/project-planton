package tofumodule

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	terraformbackendcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/terraformbackendcredential/v1"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/credentials"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfbackend"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfvars"
	"google.golang.org/protobuf/proto"
	"os"
	"os/exec"
	"path/filepath"
)

// TofuInit initializes a tofu module with optional JSON streaming.
func TofuInit(
	tofuModulePath string,
	manifestObject proto.Message,
	backendType terraformbackendcredentialv1.TerraformBackendType,
	backendConfigInput []string,
	isJsonOutput bool,
	linesChan chan string,
	stackInputOptions ...credentials.StackInputCredentialOption,
) error {
	// Gather credential options (currently unused in this snippet, but left for future usage)
	opts := credentials.StackInputCredentialOptions{}
	for _, opt := range stackInputOptions {
		opt(&opts)
	}

	// 1. Write the backend file
	if err := tfbackend.WriteBackendFile(tofuModulePath, backendType); err != nil {
		return errors.Wrapf(err, "failed to write backend file")
	}

	// 2. Write or update terraform.tfvars
	tfVarsFile := filepath.Join(tofuModulePath, ".terraform", "terraform.tfvars")
	if err := tfvars.WriteVarFile(manifestObject, tfVarsFile); err != nil {
		return errors.Wrapf(err, "failed to write %s file", tfVarsFile)
	}

	// 3. Build the 'tofu init' command with optional -json
	cmdArgs := []string{
		terraform.TerraformOperationType_init.String(),
		"--var-file", tfVarsFile,
	}
	if isJsonOutput {
		cmdArgs = append(cmdArgs, "-json")
	}

	// 4. Append backend configs if any
	for _, backendConfig := range backendConfigInput {
		cmdArgs = append(cmdArgs, "--backend-config", backendConfig)
	}

	tofuCmd := exec.Command(TofuCommand, cmdArgs...)
	tofuCmd.Dir = tofuModulePath
	// Keep stdin/stderr for interactive usage and error reporting
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("tofu module directory: %s\n", tofuModulePath)
	fmt.Printf("running command: %s\n", tofuCmd.String())

	// 5. Branch: capture lines to channel (if provided) or stream to stdout
	if linesChan != nil {
		// a) linesChan is non-nil → capture line-by-line via a pipe
		stdoutPipe, err := tofuCmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "failed to create stdout pipe")
		}
		// Start command before reading pipe
		if err := tofuCmd.Start(); err != nil {
			return errors.Wrapf(err, "failed to start tofu command %s", tofuCmd.String())
		}

		// Read stdout lines in a separate goroutine
		go func() {
			scanner := bufio.NewScanner(stdoutPipe)
			for scanner.Scan() {
				line := scanner.Text()
				linesChan <- line
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "error reading tofu output: %v\n", err)
			}
		}()

		// Wait for command completion
		if err := tofuCmd.Wait(); err != nil {
			return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
		}

	} else {
		// b) linesChan is nil → just stream stdout directly
		tofuCmd.Stdout = os.Stdout
		if err := tofuCmd.Run(); err != nil {
			return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
		}
	}

	return nil
}
