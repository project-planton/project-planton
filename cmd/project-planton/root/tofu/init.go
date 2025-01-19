package tofu

import (
	terraformbackendcredentialv1 "github.com/project-planton/project-planton/apis/project/planton/credential/terraformbackendcredential/v1"
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/internal/cli/flag"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/credentials"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tfbackend"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Init = &cobra.Command{
	Use:   "init",
	Short: "run tofu init",
	Run:   initHandler,
}

func init() {
	Init.PersistentFlags().StringArray(string(flag.BackendConfig), []string{},
		"Configuration to be merged with what is in the\n                          "+
			"configuration file's 'backend' block. "+
			"This can be\n                          either a path to an HCL file with key/value\n                          "+
			"assignments (same format as terraform.tfvars) or a\n                          'key=value' format, and can be "+
			"specified multiple\n                          times. The backend type must be in the "+
			"configuration\n                          itself.")

	Init.PersistentFlags().String(string(flag.BackendType), terraformbackendcredentialv1.TerraformBackendType_local.String(),
		"Specifies the backend type that Terraform will use to store and manage the state.\n"+
			"This must match one of the supported Terraform backends, such as 'local', 's3', 'gcs',\n"+
			"'azurerm', 'remote', 'consul', 'http', 'etcdv3', 'manta', 'swift', 'artifactory', or\n"+
			"'oss'. By default, it uses 'local', which stores the Terraform state on the local\n"+
			"filesystem.\n\n"+
			"If you choose a different backend (e.g., 's3'), you can then supply additional\n"+
			"configuration parameters using the '--backend-config' flag. For example, when using\n"+
			"'s3', you might provide a bucket name, key, region, and a DynamoDB table for locking,\n"+
			"either via a path to an HCL file or via key-value pairs.\n\n"+
			"This option can be used multiple times if you need to override certain backend\n"+
			"attributes. The backend type itself, however, must be declared in your Terraform\n"+
			"configuration using a 'terraform { backend \"<type>\" {} }' block. The '--backend-type'\n"+
			"flag will then instruct Terraform which backend configuration block to use.\n\n"+
			"Example:\n"+
			"  --backend-type=s3 --backend-config=bucket=my-terraform-bucket --backend-config=key=state.tfstate")

}

func initHandler(cmd *cobra.Command, args []string) {
	inputDir, err := cmd.Flags().GetString(string(flag.InputDir))
	flag.HandleFlagErr(err, flag.InputDir)

	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	backendTypeString, err := cmd.Flags().GetString(string(flag.BackendType))
	flag.HandleFlagErrAndValue(err, flag.BackendType, backendTypeString)

	backendConfigList, err := cmd.Flags().GetStringArray(string(flag.BackendConfig))
	flag.HandleFlagErr(err, flag.BackendConfig)

	backendType := tfbackend.BackendTypeFromString(backendTypeString)

	credentialOptions := make([]credentials.StackInputCredentialOption, 0)
	targetManifestPath := inputDir + "/target.yaml"

	if inputDir == "" {
		targetManifestPath, err = cmd.Flags().GetString(string(flag.Manifest))
		flag.HandleFlagErrAndValue(err, flag.Manifest, targetManifestPath)

		credentialOptions, err = credentials.BuildWithFlags(cmd.Flags())
		if err != nil {
			log.Fatalf("failed to build credentiaal options: %v", err)
		}
	} else {
		credentialOptions, err = credentials.BuildWithFlags(cmd.Flags())
		if err != nil {
			log.Fatalf("failed to build credentiaal options: %v", err)
		}
	}

	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		log.Fatalf("failed to override values in target manifest file")
	}

	kindName, err := apiresourcekind.ExtractKindFromProto(manifestObject)
	if err != nil {
		log.Fatalf("failed to extract kind name from manifest proto %v", err)
	}

	tofuModulePath, err := tofumodule.GetModulePath(moduleDir, kindName)
	if err != nil {
		log.Fatalf("failed to get tofu module directory %v", err)
	}

	err = tofumodule.TofuInit(tofuModulePath, manifestObject,
		backendType,
		backendConfigList,
		false, nil, credentialOptions...)
	if err != nil {
		log.Fatalf("failed to run tofu operation: %v", err)
	}
}
