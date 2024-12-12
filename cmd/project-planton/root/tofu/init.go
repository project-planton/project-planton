package tofu

import (
	"github.com/project-planton/project-planton/apis/project/planton/shared/tofu"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "run tofu init",
	Run:   initHandler,
}

func initHandler(cmd *cobra.Command, args []string) {
	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	err = tofumodule.Run(moduleDir, "", tofu.TofuOperationType_init, map[string]string{})
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
