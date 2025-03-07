package tofumodule

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfbackend"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfvars"
	"google.golang.org/protobuf/proto"
	"os"
	"os/exec"
	"path/filepath"
	"runtime/debug"
)

// TofuInit initializes a tofu module with optional JSON streaming.
func TofuInit(
	tofuModulePath string,
	manifestObject proto.Message,
	backendType terraform.TerraformBackendType,
	backendConfigInput []string,
	credentialEnvVars []string,
	isJsonOutput bool,
	jsonLogEventsChan chan string, // channel for streaming output
) (err error) {
	// (1) Process backend & tfvars as usual...
	if err := tfbackend.WriteBackendFile(tofuModulePath, backendType); err != nil {
		return errors.Wrapf(err, "failed to write backend file")
	}

	tfVarsFile := filepath.Join(tofuModulePath, ".terraform", "terraform.tfvars")
	if err := tfvars.WriteVarFile(manifestObject, tfVarsFile); err != nil {
		return errors.Wrapf(err, "failed to write %s file", tfVarsFile)
	}

	// (2) Build the Tofu 'init' command
	cmdArgs := []string{
		terraform.TerraformOperationType_init.String(),
		"--var-file", tfVarsFile,
	}
	if isJsonOutput {
		cmdArgs = append(cmdArgs, "-json")
	}
	for _, backendConfig := range backendConfigInput {
		cmdArgs = append(cmdArgs, "--backend-config", backendConfig)
	}

	tofuCmd := exec.Command(TofuCommand, cmdArgs...)
	tofuCmd.Dir = tofuModulePath
	//https://stackoverflow.com/a/41133244
	tofuCmd.Env = os.Environ()
	tofuCmd.Env = append(tofuCmd.Env, credentialEnvVars...)

	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("tofu module directory: %s\n", tofuModulePath)
	fmt.Printf("running command: %s\n", tofuCmd.String())

	// (3) If jsonLogEventsChan is provided, read stdout in a goroutine with panic recovery.
	if jsonLogEventsChan != nil {
		stdoutPipe, err := tofuCmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "failed to create stdout pipe")
		}

		if err := tofuCmd.Start(); err != nil {
			return errors.Wrapf(err, "failed to start tofu command %s", tofuCmd.String())
		}

		// We'll use errChan to propagate any panic or scanning error back to this function.
		errChan := make(chan error, 1)

		go func() {
			defer func() {
				if r := recover(); r != nil {
					// Convert panic to an error, including stack trace
					stack := debug.Stack()
					panicErr := fmt.Errorf(
						"panic recovered in TofuInit stdout reader goroutine: %v\nstack trace:\n%s",
						r, string(stack),
					)
					// Send the panic error to errChan
					errChan <- panicErr
				}
				// Close errChan so we can detect completion
				close(errChan)
			}()

			scanner := bufio.NewScanner(stdoutPipe)
			for scanner.Scan() {
				line := scanner.Text()
				jsonLogEventsChan <- line
			}
			if err := scanner.Err(); err != nil {
				// Send scanner error (only if there's no panic)
				errChan <- fmt.Errorf("error reading tofu output: %v", err)
			}
		}()

		// Wait for tofuCmd to finish executing.
		if err := tofuCmd.Wait(); err != nil {
			return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
		}

		// After tofuCmd finishes, read from errChan to see if the goroutine reported any error.
		if readErr, ok := <-errChan; ok && readErr != nil {
			// This will be either the recovered panic error or the scanner error
			return readErr
		}

		// Optionally close jsonLogEventsChan if TofuInit owns it
		// close(jsonLogEventsChan)

	} else {
		// (4) No channel → just stream stdout to console
		tofuCmd.Stdout = os.Stdout
		if err := tofuCmd.Run(); err != nil {
			return errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
		}
	}

	return nil
}
