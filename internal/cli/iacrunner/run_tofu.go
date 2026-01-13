package iacrunner

import (
	"os"

	"github.com/pkg/errors"
	"github.com/plantonhq/project-planton/apis/org/project_planton/shared/iac/terraform"
	"github.com/plantonhq/project-planton/internal/cli/cliprint"
	"github.com/plantonhq/project-planton/internal/cli/flag"
	"github.com/plantonhq/project-planton/pkg/iac/tofu/tofumodule"
	"github.com/spf13/cobra"
)

// RunTofu executes a Tofu/Terraform operation using the resolved context.
func RunTofu(ctx *Context, cmd *cobra.Command, operation terraform.TerraformOperationType) error {
	isAutoApprove, err := cmd.Flags().GetBool(string(flag.AutoApprove))
	if err != nil {
		return errors.Wrap(err, "failed to get auto-approve flag")
	}

	// For plan operation, check if it's a destroy plan
	isDestroyPlan := false
	if operation == terraform.TerraformOperationType_plan {
		isDestroyPlan, _ = cmd.Flags().GetBool(string(flag.Destroy))
		// Plan is always auto-approve (non-interactive)
		isAutoApprove = true
	}

	cliprint.PrintHandoff("OpenTofu")

	err = tofumodule.RunCommand(
		ctx.ModuleDir,
		ctx.ManifestPath,
		operation,
		ctx.ValueOverrides,
		isAutoApprove,
		isDestroyPlan,
		ctx.ModuleVersion,
		ctx.NoCleanup,
		ctx.KubeContext,
		ctx.ProviderConfigOpts...,
	)
	if err != nil {
		cliprint.PrintTofuFailure()
		os.Exit(1)
	}
	cliprint.PrintTofuSuccess()
	return nil
}
