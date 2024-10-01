package root

import "github.com/spf13/cobra"

var Pulumi = &cobra.Command{
	Use:   "pulumi",
	Short: "Run Pulumi Stacks",
}
