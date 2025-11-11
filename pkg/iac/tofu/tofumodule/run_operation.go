package tofumodule

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project-planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfvars"
	"google.golang.org/protobuf/proto"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
)

const TofuCommand = "tofu"

// RunOperation runs a tofu command, optionally adding -json flag and streaming output lines.
// It also recovers from any panic in the stdout-reading goroutine and returns it as an error.
func RunOperation(
	tofuModulePath string,
	terraformOperation terraform.TerraformOperationType,
	isAutoApprove bool,
	isDestroyPlan bool,
	manifestObject proto.Message,
	providerConfigEnvVars []string,
	isJsonOutput bool,
	jsonLogEventsChan chan string, // channel for streaming output
) (err error) {
	// Write or update terraform.tfvars
	tfVarsFile := filepath.Join(tofuModulePath, ".terraform", "terraform.tfvars")
	if err := tfvars.WriteVarFile(manifestObject, tfVarsFile); err != nil {
		return errors.Wrapf(err, "failed to write %s file", tfVarsFile)
	}

	// Determine command and arguments
	op := terraformOperation.String()
	args := []string{op, "--var-file", tfVarsFile}

	if terraformOperation == terraform.TerraformOperationType_plan {
		args = append(args, "--out", "terraform.tfplan")
		if isDestroyPlan {
			args = append(args, "--destroy")
		}
	}

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
	//https://stackoverflow.com/a/41133244
	tofuCmd.Env = os.Environ()
	tofuCmd.Env = append(tofuCmd.Env, providerConfigEnvVars...)

	// Keep stdin/stderr for interactive prompt or error streaming
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("tofu module directory: %s\n", tofuModulePath)
	fmt.Printf("running command: %s\n", tofuCmd.String())

	// If JSON output, capture stdout in a goroutine with panic recovery
	if isJsonOutput {
		stdoutPipe, err := tofuCmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "failed to create stdout pipe")
		}

		// Start the tofu command
		if err := tofuCmd.Start(); err != nil {
			return errors.Wrapf(err, "failed to start tofu command %s", tofuCmd.String())
		}

		// We'll capture errors (panic or scanner errors) in errChan
		errChan := make(chan error, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					stack := debug.Stack()
					panicErr := fmt.Errorf(
						"panic recovered in RunOperation stdout reader goroutine: %v\nstack trace:\n%s",
						r, string(stack),
					)
					errChan <- panicErr
				}
				close(errChan)
			}()

			scanner := bufio.NewScanner(stdoutPipe)
			for scanner.Scan() {
				line := scanner.Text()
				if jsonLogEventsChan != nil {
					jsonLogEventsChan <- line
				} else {
					fmt.Println(line)
				}
			}
			if err := scanner.Err(); err != nil {
				errChan <- fmt.Errorf("error reading tofu output: %v", err)
			}
		}()

		// Wait for the command to finish
		if err := tofuCmd.Wait(); err != nil {
			return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
		}

		// See if the goroutine reported any error or panic
		if readErr, ok := <-errChan; ok && readErr != nil {
			return readErr
		}

		// Optionally close jsonLogEventsChan if you want to end streaming here
		// close(jsonLogEventsChan)

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
