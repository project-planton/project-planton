package root

import (
	"github.com/plantoncloud/project-planton/cmd/project-planton/root/pulumi"
	"github.com/plantoncloud/project-planton/internal/cli/flag"
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

	Pulumi.PersistentFlags().String(string(flag.InputDir), "", "directory containing target.yaml and credential yaml files")
	Pulumi.PersistentFlags().String(string(flag.ModuleDir), pwd, "directory containing the pulumi module")
	Pulumi.PersistentFlags().String(string(flag.Stack), "", "pulumi stack fqdn in the format of <org>/<project>/<stack>")
	Pulumi.PersistentFlags().String(string(flag.Manifest), "", "path of the deployment-component manifest file")
	Pulumi.PersistentFlags().String(string(flag.AwsCredential), "", "path of the aws-credential file")
	Pulumi.PersistentFlags().String(string(flag.AzureCredential), "", "path of the azure-credential file")
	Pulumi.PersistentFlags().String(string(flag.ConfluentCredential), "", "path of the confluent-credential file")
	Pulumi.PersistentFlags().String(string(flag.DockerCredential), "", "path of the docker-credential file")
	Pulumi.PersistentFlags().String(string(flag.GcpCredential), "", "path of the gcp-credential file")
	Pulumi.PersistentFlags().String(string(flag.KubernetesCluster), "", "path of the yaml file containing the kubernetes cluster configuration")
	Pulumi.PersistentFlags().String(string(flag.MongodbAtlasCredential), "", "path of the mongodb-atlas-credential file")
	Pulumi.PersistentFlags().String(string(flag.SnowflakeCredential), "", "path of the snowflake-credential file")
	Pulumi.AddCommand(
		pulumi.Refresh,
		pulumi.Preview,
		pulumi.Update,
		pulumi.Destroy,
	)
}
