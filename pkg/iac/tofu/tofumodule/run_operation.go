package tofumodule

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/credentials"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfvars"
	"google.golang.org/protobuf/proto"
	"os"
	"os/exec"
	"path/filepath"
)

const TofuCommand = "tofu"

// RunOperation runs a tofu command, optionally adding -json flag and streaming output lines.
func RunOperation(
	tofuModulePath string,
	terraformOperation terraform.TerraformOperationType,
	isAutoApprove bool,
	manifestObject proto.Message,
	isJsonOutput bool,
	linesChan chan string,
	stackInputOptions ...credentials.StackInputCredentialOption,
) error {
	// Gather credential options (currently unused, but left for future usage)
	opts := credentials.StackInputCredentialOptions{}
	for _, opt := range stackInputOptions {
		opt(&opts)
	}

	// Write or update terraform.tfvars
	tfVarsFile := filepath.Join(tofuModulePath, ".terraform", "terraform.tfvars")
	if err := tfvars.WriteVarFile(manifestObject, tfVarsFile); err != nil {
		return errors.Wrapf(err, "failed to write %s file", tfVarsFile)
	}

	// Determine command and arguments
	op := terraformOperation.String()
	args := []string{op, "--var-file", tfVarsFile}

	// Add --auto-approve if needed
	if (terraformOperation == terraform.TerraformOperationType_apply ||
		terraformOperation == terraform.TerraformOperationType_destroy) && isAutoApprove {
		args = append(args, "--auto-approve")
	}

	// If the caller wants JSON output, add the -json flag
	if isJsonOutput {
		args = append(args, "-json")
	}

	tofuCmd := exec.Command(TofuCommand, args...)
	tofuCmd.Dir = tofuModulePath

	// Keep stdin/stderr for interactive prompt or error streaming
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("tofu module directory: %s\n", tofuModulePath)
	fmt.Printf("running command: %s\n", tofuCmd.String())

	if isJsonOutput {
		// If we want JSON output, we’ll capture stdout via a pipe
		stdoutPipe, err := tofuCmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "failed to create stdout pipe")
		}

		// Start the command before we begin reading
		if err := tofuCmd.Start(); err != nil {
			return errors.Wrapf(err, "failed to start tofu command %s", tofuCmd.String())
		}

		// Goroutine to read lines from stdout and push them to linesChan
		go func() {
			scanner := bufio.NewScanner(stdoutPipe)
			for scanner.Scan() {
				line := scanner.Text()
				// If a linesChan was provided, send the line
				if linesChan != nil {
					linesChan <- line
				} else {
					// If no channel, you could log or do something else
					// For now, just print
					fmt.Println(line)
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "error reading tofu output: %v\n", err)
			}
			// If a linesChan was provided, optionally close it after done
			if linesChan != nil {
				close(linesChan)
			}
		}()

		// Wait for the command to finish
		if err := tofuCmd.Wait(); err != nil {
			return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
		}

	} else {
		// If we do NOT want JSON output, simply stream stdout to console
		tofuCmd.Stdout = os.Stdout

		// Run+Wait in one go
		if err := tofuCmd.Run(); err != nil {
			return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
		}
	}

	return nil
}
