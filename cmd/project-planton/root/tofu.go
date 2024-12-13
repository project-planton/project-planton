package root

import (
	"github.com/project-planton/project-planton/cmd/project-planton/root/tofu"
	"github.com/project-planton/project-planton/internal/cli/flag"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var Tofu = &cobra.Command{
	Use:   "tofu",
	Short: "run open-tofu commands",
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get current working directory")
	}

	Tofu.PersistentFlags().String(string(flag.Manifest), "", "path of the deployment-component manifest file")

	Tofu.PersistentFlags().String(string(flag.InputDir), "", "directory containing target.yaml and credential yaml files")
	Tofu.PersistentFlags().String(string(flag.ModuleDir), pwd, "directory containing the terraform module")
	Tofu.PersistentFlags().StringToString(string(flag.Set), map[string]string{}, "override resource manifest values using key=value pairs")

	Tofu.PersistentFlags().String(string(flag.AwsCredential), "", "path of the aws-credential file")
	Tofu.PersistentFlags().String(string(flag.AzureCredential), "", "path of the azure-credential file")
	Tofu.PersistentFlags().String(string(flag.ConfluentCredential), "", "path of the confluent-credential file")
	Tofu.PersistentFlags().String(string(flag.DockerCredential), "", "path of the docker-credential file")
	Tofu.PersistentFlags().String(string(flag.GcpCredential), "", "path of the gcp-credential file")
	Tofu.PersistentFlags().String(string(flag.KubernetesCluster), "", "path of the yaml file containing the kubernetes cluster configuration")
	Tofu.PersistentFlags().String(string(flag.MongodbAtlasCredential), "", "path of the mongodb-atlas-credential file")
	Tofu.PersistentFlags().String(string(flag.SnowflakeCredential), "", "path of the snowflake-credential file")

	Tofu.AddCommand(
		tofu.Apply,
		tofu.Destroy,
		tofu.Init,
		tofu.LoadTfVars,
		tofu.Plan,
		tofu.Refresh,
	)
}
