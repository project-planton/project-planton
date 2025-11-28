package root

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"github.com/project-planton/project-planton/app/backend/apis/gen/go/proto/backendv1connect"
	"github.com/spf13/cobra"
)

var CloudResourceDeleteCmd = &cobra.Command{
	Use:   "cloud-resource:delete",
	Short: "delete a cloud resource by ID",
	Long:  "Delete a cloud resource by providing its unique ID. This action is irreversible.",
	Run:   cloudResourceDeleteHandler,
}

func init() {
	CloudResourceDeleteCmd.Flags().StringP("id", "i", "", "unique identifier of the cloud resource (required)")
	CloudResourceDeleteCmd.MarkFlagRequired("id")
}

func cloudResourceDeleteHandler(cmd *cobra.Command, args []string) {
	// Get resource ID
	id, _ := cmd.Flags().GetString("id")
	if id == "" {
		fmt.Println("Error: --id flag is required. Provide the cloud resource ID")
		fmt.Println("Usage: project-planton cloud-resource:delete --id=<resource-id>")
		os.Exit(1)
	}

	// Get backend URL from configuration
	backendURL, err := GetBackendURL()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Create Connect-RPC client
	client := backendv1connect.NewCloudResourceServiceClient(
		http.DefaultClient,
		backendURL,
	)

	// Prepare request
	req := &backendv1.DeleteCloudResourceRequest{
		Id: id,
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Make the API call
	resp, err := client.DeleteCloudResource(ctx, connect.NewRequest(req))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeUnavailable {
			fmt.Printf("Error: Cannot connect to backend service at %s. Please check:\n", backendURL)
			fmt.Printf("  1. The backend service is running\n")
			fmt.Printf("  2. The backend URL is correct\n")
			fmt.Printf("  3. Network connectivity\n")
			os.Exit(1)
		}
		if connect.CodeOf(err) == connect.CodeNotFound {
			fmt.Printf("Error: Cloud resource with ID '%s' not found\n", id)
			os.Exit(1)
		}
		fmt.Printf("Error deleting cloud resource: %v\n", err)
		os.Exit(1)
	}

	// Display success message
	fmt.Println("✅ " + resp.Msg.Message)
}

