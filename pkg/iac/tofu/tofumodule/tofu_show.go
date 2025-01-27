package tofumodule

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// TofuShowJsonOutput runs `tofu show -json <planFile>` in the specified module path,
// captures the JSON output, and returns it as a string.
func TofuShowJsonOutput(tofuModulePath, planFile string, credentialEnvVars []string) (string, error) {
	// Example command: tofu show -json terraform.tfplan
	args := []string{"show", "-json", planFile}

	tofuCmd := exec.Command(TofuCommand, args...)
	tofuCmd.Dir = tofuModulePath

	// Inherit all environment variables and append any credentials or custom vars
	tofuCmd.Env = os.Environ()
	tofuCmd.Env = append(tofuCmd.Env, credentialEnvVars...)

	// Optionally attach stdin/stderr, so tofu can prompt or display errors
	tofuCmd.Stdin = os.Stdin
	tofuCmd.Stderr = os.Stderr

	fmt.Printf("tofu module directory: %s\n", tofuModulePath)
	fmt.Printf("running command: %s\n", tofuCmd.String())

	// Capture the JSON output from stdout
	outputBytes, err := tofuCmd.Output()
	if err != nil {
		return "", errors.Wrapf(err, "failed to execute tofu command %s", tofuCmd.String())
	}

	return string(outputBytes), nil
}
