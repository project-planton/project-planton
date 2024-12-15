package tofu

import (
	"github.com/project-planton/project-planton/internal/apiresourcekind"
	"github.com/project-planton/project-planton/internal/iac/tofu/variablestf"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var GenerateVariables = &cobra.Command{
	Use:   "generate-variables <deployment-component>",
	Short: "Generate Terraform variables for a specified deployment component",
	Long: `The "generate-variables" command takes a specified project-planton 
deployment component type (e.g., "S3Bucket", "RedisKubernetes") and generates 
Terraform variable definitions (variables.tf) and a corresponding 
terraform.tfvars file.

This command instantiates an empty object of the specified component kind 
under the hood, and then converts that empty object into a Terraform-compatible 
variables file. These variables can then be passed into Terraform modules, 
streamlining infrastructure provisioning and ensuring a consistent, 
declarative workflow.`,
	Example: `
  # Generate variables for an S3Bucket deployment component
  project-planton tofu generate-variables S3Bucket

  # Generate variables for a RedisKubernetes deployment component
  project-planton tofu generate-variables RedisKubernetes
`,
	Args: cobra.ExactArgs(1), // "s3-bucket", "redis-kubernetes", etc.
	Run:  generateVariablesHandler,
}

func generateVariablesHandler(cmd *cobra.Command, args []string) {
	kindName := args[0]
	manifestObject := apiresourcekind.DeploymentComponentMap[apiresourcekind.FindMatchingComponent(
		apiresourcekind.ConvertKindName(kindName))]
	variablesTfContent, err := variablestf.ProtoToVariablesTF(manifestObject)
	log.Fatal("failed to generate Terraform variables: ", err)
	println(variablesTfContent)
}
