package root

import (
	"fmt"
	"os"

	"github.com/project-planton/project-planton/apis/org/project_planton/shared/iac/pulumi"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/iac/terraform"
	"github.com/project-planton/project-planton/internal/cli/cliprint"
	"github.com/project-planton/project-planton/internal/cli/flag"
	climanifest "github.com/project-planton/project-planton/internal/cli/manifest"
	"github.com/project-planton/project-planton/internal/cli/prompt"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/iac/provisioner"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumistack"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"
	"github.com/project-planton/project-planton/pkg/iac/tofu/tofumodule"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Apply = &cobra.Command{
	Use:   "apply",
	Short: "apply infrastructure changes using the provisioner specified in manifest",
	Long: `Apply infrastructure changes by automatically routing to the appropriate provisioner
(Pulumi, Tofu, or Terraform) based on the manifest label 'project-planton.org/provisioner'.

If the provisioner label is not present, you will be prompted to select one interactively.`,
	Example: `
	# Apply with manifest file
	project-planton apply -f manifest.yaml
	project-planton apply --manifest manifest.yaml

	# Apply with kustomize
	project-planton apply --kustomize-dir _kustomize --overlay prod

	# Apply with field overrides
	project-planton apply -f manifest.yaml --set spec.version=v1.2.3
	`,
	Run: applyHandler,
}

func init() {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("failed to get current working directory")
	}

	// Use StringP to support both --manifest and -f
	Apply.PersistentFlags().StringP(string(flag.Manifest), "f", "", "path of the deployment-component manifest file")

	Apply.PersistentFlags().String(string(flag.InputDir), "", "directory containing target.yaml and credential yaml files")
	Apply.PersistentFlags().String(string(flag.KustomizeDir), "", "directory containing kustomize configuration")
	Apply.PersistentFlags().String(string(flag.Overlay), "", "kustomize overlay to use (e.g., prod, dev, staging)")
	Apply.PersistentFlags().String(string(flag.ModuleDir), pwd, "directory containing the provisioner module")
	Apply.PersistentFlags().StringToString(string(flag.Set), map[string]string{}, "override resource manifest values using key=value pairs")

	// Pulumi-specific flags
	Apply.PersistentFlags().String(string(flag.Stack), "", "pulumi stack fqdn in the format of <org>/<project>/<stack>")
	Apply.PersistentFlags().Bool(string(flag.Yes), false, "Automatically approve and perform the update after previewing it (Pulumi)")
	Apply.PersistentFlags().Bool(string(flag.Diff), false, "Show detailed resource diffs (Pulumi)")

	// Tofu/Terraform-specific flags
	Apply.PersistentFlags().Bool(string(flag.AutoApprove), false, "Skip interactive approval of plan before applying (Tofu/Terraform)")

	// Staging/cleanup flags
	Apply.PersistentFlags().Bool(string(flag.NoCleanup), false, "Do not cleanup the workspace copy after execution (keeps cloned modules)")
	Apply.PersistentFlags().String(string(flag.ModuleVersion), "",
		"Checkout a specific version (tag, branch, or commit SHA) of the IaC modules in the workspace copy.\n"+
			"This allows using a different module version than what's in the staging area without affecting it.")

	// Provider credential flags
	Apply.PersistentFlags().String(string(flag.AtlasProviderConfig), "", "path of the mongodb-atlas-credential file")
	Apply.PersistentFlags().String(string(flag.Auth0ProviderConfig), "", "path of the auth0-credential file")
	Apply.PersistentFlags().String(string(flag.AwsProviderConfig), "", "path of the aws-credential file")
	Apply.PersistentFlags().String(string(flag.AzureProviderConfig), "", "path of the azure-credential file")
	Apply.PersistentFlags().String(string(flag.CloudflareProviderConfig), "", "path of the cloudflare-credential file")
	Apply.PersistentFlags().String(string(flag.ConfluentProviderConfig), "", "path of the confluent-credential file")
	Apply.PersistentFlags().String(string(flag.GcpProviderConfig), "", "path of the gcp-credential file")
	Apply.PersistentFlags().String(string(flag.KubernetesProviderConfig), "", "path of the yaml file containing the kubernetes cluster configuration")
	Apply.PersistentFlags().String(string(flag.SnowflakeProviderConfig), "", "path of the snowflake-credential file")
}

func applyHandler(cmd *cobra.Command, args []string) {
	moduleDir, err := cmd.Flags().GetString(string(flag.ModuleDir))
	flag.HandleFlagErrAndValue(err, flag.ModuleDir, moduleDir)

	valueOverrides, err := cmd.Flags().GetStringToString(string(flag.Set))
	flag.HandleFlagErr(err, flag.Set)

	// Check which manifest source is being used for informative messages
	kustomizeDir, _ := cmd.Flags().GetString(string(flag.KustomizeDir))
	overlay, _ := cmd.Flags().GetString(string(flag.Overlay))

	if kustomizeDir != "" && overlay != "" {
		cliprint.PrintStep(fmt.Sprintf("Building manifest from kustomize overlay: %s", overlay))
	} else {
		cliprint.PrintStep("Loading manifest...")
	}

	// Resolve manifest path with priority: --manifest > --input-dir > --kustomize-dir + --overlay
	targetManifestPath, isTemp, err := climanifest.ResolveManifestPath(cmd)
	if err != nil {
		cliprint.PrintError(fmt.Sprintf("Failed to resolve manifest: %v", err))
		os.Exit(1)
	}
	if isTemp {
		defer os.Remove(targetManifestPath)
	}

	cliprint.PrintSuccess("Manifest loaded")

	// Apply value overrides if any (creates new temp file if overrides exist)
	if len(valueOverrides) > 0 {
		cliprint.PrintStep(fmt.Sprintf("Applying %d field override(s)...", len(valueOverrides)))
	}

	finalManifestPath, isTempOverrides, err := manifest.ApplyOverridesToFile(targetManifestPath, valueOverrides)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if isTempOverrides {
		defer os.Remove(finalManifestPath)
		targetManifestPath = finalManifestPath
		cliprint.PrintSuccess("Overrides applied")
	}

	// Validate manifest before proceeding (after overrides are applied)
	cliprint.PrintStep("Validating manifest...")
	if err := manifest.Validate(targetManifestPath); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	cliprint.PrintSuccess("Manifest validated")

	// Load manifest to extract provisioner
	cliprint.PrintStep("Detecting provisioner...")
	manifestObject, err := manifest.LoadWithOverrides(targetManifestPath, valueOverrides)
	if err != nil {
		cliprint.PrintError(fmt.Sprintf("Failed to load manifest: %v", err))
		os.Exit(1)
	}

	// Extract provisioner from manifest
	provType, err := provisioner.ExtractFromManifest(manifestObject)
	if err != nil {
		cliprint.PrintError(fmt.Sprintf("Invalid provisioner in manifest: %v", err))
		os.Exit(1)
	}

	// If provisioner not specified in manifest, prompt user
	if provType == provisioner.ProvisionerTypeUnspecified {
		cliprint.PrintInfo("Provisioner not specified in manifest")
		provType, err = prompt.PromptForProvisioner()
		if err != nil {
			cliprint.PrintError(fmt.Sprintf("Failed to get provisioner: %v", err))
			os.Exit(1)
		}
	}

	cliprint.PrintSuccess(fmt.Sprintf("Using provisioner: %s", provType.String()))

	// Prepare provider configs
	cliprint.PrintStep("Preparing execution...")
	providerConfigOptions, err := stackinputproviderconfig.BuildWithFlags(cmd.Flags())
	if err != nil {
		cliprint.PrintError(fmt.Sprintf("Failed to build credential options: %v", err))
		os.Exit(1)
	}
	cliprint.PrintSuccess("Execution prepared")

	// Route to appropriate provisioner
	switch provType {
	case provisioner.ProvisionerTypePulumi:
		applyWithPulumi(cmd, moduleDir, targetManifestPath, valueOverrides, providerConfigOptions)
	case provisioner.ProvisionerTypeTofu:
		applyWithTofu(cmd, moduleDir, targetManifestPath, valueOverrides, providerConfigOptions)
	case provisioner.ProvisionerTypeTerraform:
		cliprint.PrintError("Terraform provisioner is not yet implemented. Please use 'tofu' instead.")
		os.Exit(1)
	default:
		cliprint.PrintError("Unknown provisioner type")
		os.Exit(1)
	}
}

func applyWithPulumi(cmd *cobra.Command, moduleDir, targetManifestPath string, valueOverrides map[string]string,
	providerConfigOptions []stackinputproviderconfig.StackInputProviderConfigOption) {

	// Stack can be provided via flag or extracted from manifest
	stackFqdn, err := cmd.Flags().GetString(string(flag.Stack))
	flag.HandleFlagErr(err, flag.Stack)

	showDiff, _ := cmd.Flags().GetBool(string(flag.Diff))
	noCleanup, _ := cmd.Flags().GetBool(string(flag.NoCleanup))
	moduleVersion, _ := cmd.Flags().GetString(string(flag.ModuleVersion))

	err = pulumistack.Run(moduleDir, stackFqdn, targetManifestPath,
		pulumi.PulumiOperationType_update, false, true, valueOverrides, showDiff, moduleVersion, noCleanup, providerConfigOptions...)
	if err != nil {
		cliprint.PrintPulumiFailure()
		os.Exit(1)
	}
	cliprint.PrintPulumiSuccess()
}

func applyWithTofu(cmd *cobra.Command, moduleDir, targetManifestPath string, valueOverrides map[string]string,
	providerConfigOptions []stackinputproviderconfig.StackInputProviderConfigOption) {

	isAutoApprove, err := cmd.Flags().GetBool(string(flag.AutoApprove))
	flag.HandleFlagErr(err, flag.AutoApprove)

	noCleanup, _ := cmd.Flags().GetBool(string(flag.NoCleanup))
	moduleVersion, _ := cmd.Flags().GetString(string(flag.ModuleVersion))

	cliprint.PrintHandoff("OpenTofu")

	err = tofumodule.RunCommand(moduleDir, targetManifestPath, terraform.TerraformOperationType_apply,
		valueOverrides,
		isAutoApprove,
		false,
		moduleVersion, noCleanup,
		providerConfigOptions...)
	if err != nil {
		cliprint.PrintTofuFailure()
		os.Exit(1)
	}
	cliprint.PrintTofuSuccess()
}
