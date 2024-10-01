package root

import (
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var Pulumi = &cobra.Command{
	Use:   "pulumi",
	Short: "Run Pulumi Stacks",
	Run:   pulumiHandler,
}

func pulumiHandler(cmd *cobra.Command, args []string) {
	pulumiCommand := exec.Command("pulumi", args...)
	pulumiCommand.Stdout = os.Stdout
	pulumiCommand.Stderr = os.Stderr
	if err := pulumiCommand.Run(); err != nil {
		log.Fatalf("pulumi command failed: %s", err.Error())
	}
}
