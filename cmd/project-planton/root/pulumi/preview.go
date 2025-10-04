package pulumi

import (
	"os"

	"github.com/project-planton/project-planton/apis/project/planton/shared/iac/pulumi"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/cli/manifest"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumistack"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputcredentials"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Preview = &cobra.Command{
	Use:   "preview",
	Short: "run pulumi update preview",
	Run:   previewHandler,
}

func previewHandler(cmd *cobra.Command, args []string) {
	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	stackFqdn, err := cmd.Flags().GetString(string(flag.Stack))
	flag.HandleFlagErrAndValue(err, flag.Stack, stackFqdn)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	// Resolve manifest path with priority: --manifest > --input-dir > --kustomize-dir + --overlay
	targetManifestPath, isTemp, err := manifest.ResolveManifestPath(cmd)
	if err != nil {
		log.Fatalf("failed to resolve manifest: %v", err)
	}
	if isTemp {
		defer os.Remove(targetManifestPath)
	}

	credentialOptions, err := stackinputcredentials.BuildWithFlags(cmd.Flags())
	if err != nil {
		log.Fatalf("failed to build credentiaal options: %v", err)
	}

	err = pulumistack.Run(moduleDir, stackFqdn, targetManifestPath,
		pulumi.PulumiOperationType_update, true, false, valueOverrides, credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run pulumi: %v", err)
	}
}
