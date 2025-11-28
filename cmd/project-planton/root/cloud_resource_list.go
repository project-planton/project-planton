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

var CloudResourceListCmd = &cobra.Command{
	Use:   "cloud-resource:list",
	Short: "list all cloud resources from backend",
	Long:  "List all cloud resources from the backend service. Optionally filter by kind.",
	Run:   cloudResourceListHandler,
}

func init() {
	CloudResourceListCmd.Flags().StringP("kind", "k", "", "filter cloud resources by kind")
}

func cloudResourceListHandler(cmd *cobra.Command, args []string) {
	// Get backend URL from configuration
	backendURL, err := GetBackendURL()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Get kind filter if provided
	kind, _ := cmd.Flags().GetString("kind")

	// Create Connect-RPC client
	client := backendv1connect.NewCloudResourceServiceClient(
		http.DefaultClient,
		backendURL,
	)

	// Prepare request
	req := &backendv1.ListCloudResourcesRequest{}
	if kind != "" {
		req.Kind = &kind
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Make the API call
	resp, err := client.ListCloudResources(ctx, connect.NewRequest(req))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeUnavailable {
			fmt.Printf("Error: Cannot connect to backend service at %s. Please check:\n", backendURL)
			fmt.Printf("  1. The backend service is running\n")
			fmt.Printf("  2. The backend URL is correct\n")
			fmt.Printf("  3. Network connectivity\n")
			os.Exit(1)
		}
		fmt.Printf("Error fetching cloud resources: %v\n", err)
		os.Exit(1)
	}

	resources := resp.Msg.Resources

	if len(resources) == 0 {
		if kind != "" {
			fmt.Printf("No cloud resources found with kind '%s'\n", kind)
		} else {
			fmt.Println("No cloud resources found")
		}
		return
	}

	// Display results in table format
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	// Table header
	fmt.Fprintln(w, "ID\tNAME\tKIND\tCREATED")

	// Table rows
	for _, resource := range resources {
		createdAt := "N/A"
		if resource.CreatedAt != nil {
			createdAt = resource.CreatedAt.AsTime().Format("2006-01-02 15:04:05")
		}

		fmt.Fprintf(w, "%s\t%s\t%s\t%s\n",
			resource.Id,
			resource.Name,
			resource.Kind,
			createdAt,
		)
	}

	// Summary
	fmt.Printf("\nTotal: %d cloud resource(s)", len(resources))
	if kind != "" {
		fmt.Printf(" (filtered by kind: %s)", kind)
	}
	fmt.Println()
}

