package root

import (
	"fmt"
	"github.com/project-planton/project-planton/internal/cli/version"

	"github.com/spf13/cobra"
)

var Version = &cobra.Command{
	Use:   "version",
	Short: "check the version of the cli",
	Run:   versionHandler,
}

func versionHandler(cmd *cobra.Command, args []string) {
	fmt.Printf("%s\n", version.Version)
}
