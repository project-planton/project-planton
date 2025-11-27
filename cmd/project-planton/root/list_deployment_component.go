package root

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"text/tabwriter"
	"time"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/internal/backend/proto"
	"github.com/project-planton/project-planton/internal/backend/proto/backendv1connect"
	"github.com/spf13/cobra"
)

var ListDeploymentComponent = &cobra.Command{
	Use:   "list-deployment-components",
	Short: "list deployment components from backend",
	Long:  "List deployment components from the backend service. Optionally filter by kind.",
	Run:   listDeploymentComponentHandler,
}

func init() {
	ListDeploymentComponent.Flags().StringP("kind", "k", "", "filter deployment components by kind")
}

func listDeploymentComponentHandler(cmd *cobra.Command, args []string) {
	// Get backend URL from configuration
	backendURL, err := GetBackendURL()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Get kind filter if provided
	kind, _ := cmd.Flags().GetString("kind")

	// Create Connect-RPC client
	client := backendv1connect.NewDeploymentComponentServiceClient(
		http.DefaultClient,
		backendURL,
	)

	// Prepare request
	req := &backendv1.ListDeploymentComponentsRequest{}
	if kind != "" {
		req.Kind = &kind
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Make the API call
	resp, err := client.ListDeploymentComponents(ctx, connect.NewRequest(req))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeUnavailable {
			fmt.Printf("Error: Cannot connect to backend service at %s. Please check:\n", backendURL)
			fmt.Printf("  1. The backend service is running\n")
			fmt.Printf("  2. The backend URL is correct\n")
			fmt.Printf("  3. Network connectivity\n")
			os.Exit(1)
		}
		fmt.Printf("Error fetching deployment components: %v\n", err)
		os.Exit(1)
	}

	components := resp.Msg.Components

	if len(components) == 0 {
		if kind != "" {
			fmt.Printf("No deployment components found with kind '%s'\n", kind)
		} else {
			fmt.Println("No deployment components found")
		}
		return
	}

	// Display results in table format
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	// Table header
	fmt.Fprintln(w, "NAME\tKIND\tPROVIDER\tVERSION\tID PREFIX\tSERVICE KIND\tCREATED")

	// Table rows
	for _, component := range components {
		isServiceKind := "No"
		if component.IsServiceKind {
			isServiceKind = "Yes"
		}

		createdAt := "N/A"
		if component.CreatedAt != nil {
			createdAt = component.CreatedAt.AsTime().Format("2006-01-02")
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\t%s\t%s\t%s\n",
			component.Name,
			component.Kind,
			component.Provider,
			component.Version,
			component.IdPrefix,
			isServiceKind,
			createdAt,
		)
	}

	// Summary
	fmt.Printf("\nTotal: %d deployment component(s)", len(components))
	if kind != "" {
		fmt.Printf(" (filtered by kind: %s)", kind)
	}
	fmt.Println()
}
