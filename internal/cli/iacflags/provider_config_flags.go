package iacflags

import (
	"github.com/plantonhq/project-planton/internal/cli/flag"
	"github.com/spf13/cobra"
)

// AddProviderConfigFlags adds all cloud provider credential flags.
func AddProviderConfigFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().String(string(flag.AtlasProviderConfig), "",
		"path of the mongodb-atlas-credential file")

	cmd.PersistentFlags().String(string(flag.Auth0ProviderConfig), "",
		"path of the auth0-credential file")

	cmd.PersistentFlags().String(string(flag.AwsProviderConfig), "",
		"path of the aws-credential file")

	cmd.PersistentFlags().String(string(flag.AzureProviderConfig), "",
		"path of the azure-credential file")

	cmd.PersistentFlags().String(string(flag.CloudflareProviderConfig), "",
		"path of the cloudflare-credential file")

	cmd.PersistentFlags().String(string(flag.ConfluentProviderConfig), "",
		"path of the confluent-credential file")

	cmd.PersistentFlags().String(string(flag.GcpProviderConfig), "",
		"path of the gcp-credential file")

	cmd.PersistentFlags().String(string(flag.KubernetesProviderConfig), "",
		"path of the yaml file containing the kubernetes cluster configuration")

	cmd.PersistentFlags().String(string(flag.SnowflakeProviderConfig), "",
		"path of the snowflake-credential file")

	cmd.PersistentFlags().String(string(flag.OpenFgaProviderConfig), "",
		"path of the openfga-credential file")
}
