package root

import (
	"github.com/plantoncloud/project-planton/cmd/project-planton/root/pulumi"
	"github.com/plantoncloud/project-planton/internal/cli/flag"
	"github.com/spf13/cobra"
)

var Pulumi = &cobra.Command{
	Use:   "pulumi",
	Short: "run a pulumi stack",
}

func init() {
	Pulumi.PersistentFlags().String(string(flag.Stack), "", "pulumi stack fqdn in the format of <org>/<project>/<stack>")
	Pulumi.PersistentFlags().String(string(flag.Input), "", "path of the stack input file")
	Pulumi.AddCommand(
		pulumi.Refresh,
		pulumi.Preview,
		pulumi.Update,
		pulumi.Destroy,
	)
}
