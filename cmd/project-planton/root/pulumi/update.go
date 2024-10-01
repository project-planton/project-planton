package pulumi

import (
	"buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/shared/pulumi"
	"github.com/plantoncloud/project-planton/internal/cli/flag"
	"github.com/plantoncloud/project-planton/internal/pulumistack"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Update = &cobra.Command{
	Use:   "update",
	Short: "run pulumi update",
	Run:   updateHandler,
}

func updateHandler(cmd *cobra.Command, args []string) {
	stackFqdn, err := cmd.Flags().GetString(string(flag.Stack))
	flag.HandleFlagErrAndValue(err, flag.Stack, stackFqdn)

	stackInputYamlPath, err := cmd.Flags().GetString(string(flag.Input))
	flag.HandleFlagErrAndValue(err, flag.Input, stackInputYamlPath)

	err = pulumistack.Run(stackFqdn, stackInputYamlPath, pulumi.PulumiOperationType_update, false)
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
