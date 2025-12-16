package root

import (
	"fmt"
	"github.com/spf13/cobra"
)

// ListDeploymentComponent command is temporarily disabled as the DeploymentComponent service
// is not currently available in the proto definitions.
// TODO: Re-enable when DeploymentComponent service is implemented
var ListDeploymentComponent = &cobra.Command{
	Use:   "list-deployment-components",
	Short: "list deployment components from backend (disabled)",
	Long:  "This command is currently disabled as the DeploymentComponent service is not available.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("⚠️  This command is currently disabled.")
		fmt.Println("   The DeploymentComponent service is not available in the current API definitions.")
	},
}
