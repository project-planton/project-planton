package pulumi

import (
	"buf.build/gen/go/plantoncloud/project-planton/protocolbuffers/go/project/planton/shared/pulumi"
	"github.com/plantoncloud/project-planton/internal/cli/flag"
	"github.com/plantoncloud/project-planton/internal/pulumistack"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Preview = &cobra.Command{
	Use:   "preview",
	Short: "run pulumi update preview",
	Run:   previewHandler,
}

func previewHandler(cmd *cobra.Command, args []string) {
	stackFqdn, err := cmd.Flags().GetString(string(flag.Stack))
	flag.HandleFlagErrAndValue(err, flag.Stack, stackFqdn)

	targetManifestPath, err := cmd.Flags().GetString(string(flag.Target))
	flag.HandleFlagErrAndValue(err, flag.Target, targetManifestPath)

	err = pulumistack.Run(stackFqdn, targetManifestPath, pulumi.PulumiOperationType_update, true)
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
