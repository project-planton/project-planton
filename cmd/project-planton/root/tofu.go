package root

import (
	"github.com/project-planton/project-planton/cmd/project-planton/root/tofu"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/spf13/cobra"
)

var Tofu = &cobra.Command{
	Use:   "tofu",
	Short: "run open-tofu commands",
}

func init() {
	Tofu.PersistentFlags().String(string(flag.Manifest), "", "path of the deployment-component manifest file")
	Tofu.AddCommand(
		tofu.LoadTfVars,
	)
}
