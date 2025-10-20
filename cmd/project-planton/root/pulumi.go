package root

import (
	"github.com/project-planton/project-planton/cmd/project-planton/root/pulumi"
	"github.com/project-planton/project-planton/internal/cli/flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var Pulumi = &cobra.Command{
	Use:   "pulumi",
	Short: "run a pulumi stack",
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get current working directory")
	}

	Pulumi.PersistentFlags().String(string(flag.Manifest), "", "path of the deployment-component manifest file")

	Pulumi.PersistentFlags().String(string(flag.InputDir), "", "directory containing target.yaml and credential yaml files")
	Pulumi.PersistentFlags().String(string(flag.KustomizeDir), "", "directory containing kustomize configuration")
	Pulumi.PersistentFlags().String(string(flag.Overlay), "", "kustomize overlay to use (e.g., prod, dev, staging)")
	Pulumi.PersistentFlags().String(string(flag.ModuleDir), pwd, "directory containing the pulumi module")
	Pulumi.PersistentFlags().StringToString(string(flag.Set), map[string]string{}, "override resource manifest values using key=value pairs")

	Pulumi.PersistentFlags().String(string(flag.Stack), "", "pulumi stack fqdn in the format of <org>/<project>/<stack>")
	Pulumi.PersistentFlags().Bool(string(flag.Yes), false, "Automatically approve and perform the update after previewing it")
	Pulumi.PersistentFlags().Bool(string(flag.Force), false, "Force removal of stack even if resources exist (use with delete/rm command)")
	Pulumi.PersistentFlags().Bool(string(flag.Diff), false, "Show detailed resource diffs")

	Pulumi.PersistentFlags().String(string(flag.AwsProviderConfig), "", "path of the aws-credential file")
	Pulumi.PersistentFlags().String(string(flag.AzureProviderConfig), "", "path of the azure-credential file")
	Pulumi.PersistentFlags().String(string(flag.ConfluentProviderConfig), "", "path of the confluent-credential file")
	Pulumi.PersistentFlags().String(string(flag.GcpProviderConfig), "", "path of the gcp-credential file")
	Pulumi.PersistentFlags().String(string(flag.KubernetesProviderConfig), "", "path of the yaml file containing the kubernetes cluster configuration")
	Pulumi.PersistentFlags().String(string(flag.AtlasProviderConfig), "", "path of the mongodb-atlas-credential file")
	Pulumi.PersistentFlags().String(string(flag.SnowflakeProviderConfig), "", "path of the snowflake-credential file")

	Pulumi.AddCommand(
		pulumi.Init,
		pulumi.Refresh,
		pulumi.Preview,
		pulumi.Update,
		pulumi.Destroy,
		pulumi.Delete,
		pulumi.Cancel,
	)
}
