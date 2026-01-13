package manifest

import (
	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/internal/cli/flag"
	"github.com/plantonhq/project-planton/pkg/iac/stackinput"
	"github.com/spf13/cobra"
)

// resolveFromStackInput checks for --stack-input flag and extracts manifest from it.
// Returns empty string if flag not provided.
// When a stack input file is provided, the manifest is extracted from the "target" field
// and written to a temporary file.
func resolveFromStackInput(cmd *cobra.Command) (manifestPath string, isTemp bool, err error) {
	stackInputPath, err := cmd.Flags().GetString(string(flag.StackInput))
	if err != nil {
		return "", false, errors.Wrap(err, "failed to get stack-input flag")
	}

	if stackInputPath == "" {
		return "", false, nil
	}

	manifestPath, err = stackinput.ExtractManifestFromStackInput(stackInputPath)
	if err != nil {
		return "", false, errors.Wrapf(err, "failed to extract manifest from stack input %s", stackInputPath)
	}

	return manifestPath, true, nil
}
