package iacflags

import (
	"github.com/plantonhq/project-planton/internal/cli/flag"
	"github.com/spf13/cobra"
)

// AddTofuApplyFlags adds Tofu/Terraform flags for apply and destroy commands.
func AddTofuApplyFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool(string(flag.AutoApprove), false,
		"Skip interactive approval of plan before applying (Tofu/Terraform)")
}

// AddTofuPlanFlags adds Tofu/Terraform flags for the plan command.
func AddTofuPlanFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().Bool(string(flag.Destroy), false,
		"Create a destroy plan instead of apply plan (Tofu/Terraform)")
}
