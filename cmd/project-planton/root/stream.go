package root

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"github.com/project-planton/project-planton/app/backend/apis/gen/go/proto/backendv1connect"
	"github.com/spf13/cobra"
)

var StreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "stream data from backend",
	Long:  "Stream dummy data from the backend service. Supports configurable message count and interval.",
	Run:   streamHandler,
}

func init() {
	StreamCmd.Flags().Int32P("count", "c", 10, "number of messages to stream (0 = infinite)")
	StreamCmd.Flags().Int32P("interval", "i", 1000, "interval between messages in milliseconds")
}

func streamHandler(cmd *cobra.Command, args []string) {
	// Get flags
	messageCount, _ := cmd.Flags().GetInt32("count")
	intervalMs, _ := cmd.Flags().GetInt32("interval")

	// Get backend URL from configuration
	backendURL, err := GetBackendURL()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Create Connect-RPC client
	client := backendv1connect.NewStreamingServiceClient(
		http.DefaultClient,
		backendURL,
	)

	// Prepare request
	req := &backendv1.StreamDataRequest{}
	if messageCount > 0 {
		req.MessageCount = &messageCount
	}
	if intervalMs > 0 {
		req.IntervalMs = &intervalMs
	}

	// Create context with timeout (or no timeout for infinite streams)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle interrupt signal for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sigChan
		fmt.Println("\n\n⚠️  Interrupt received, stopping stream...")
		cancel()
	}()

	// Start streaming
	fmt.Printf("🚀 Starting stream...\n")
	if messageCount > 0 {
		fmt.Printf("   Message count: %d\n", messageCount)
	} else {
		fmt.Printf("   Message count: infinite (press Ctrl+C to stop)\n")
	}
	fmt.Printf("   Interval: %dms\n\n", intervalMs)

	stream, err := client.StreamData(ctx, connect.NewRequest(req))
	if err != nil {
		if connect.CodeOf(err) == connect.CodeUnavailable {
			fmt.Printf("❌ Error: Cannot connect to backend service at %s. Please check:\n", backendURL)
			fmt.Printf("  1. The backend service is running\n")
			fmt.Printf("  2. The backend URL is correct\n")
			fmt.Printf("  3. Network connectivity\n")
			os.Exit(1)
		}
		fmt.Printf("❌ Error starting stream: %v\n", err)
		os.Exit(1)
	}

	// Iterate over the stream
	messageNum := 0
	for stream.Receive() {
		messageNum++
		response := stream.Msg()

		// Format timestamp
		timestamp := "N/A"
		if response.Timestamp != nil {
			timestamp = response.Timestamp.AsTime().Format("15:04:05.000")
		}

		// Log the message
		fmt.Printf("[%s] [Seq: %d] [Status: %s] %s\n",
			timestamp,
			response.Sequence,
			response.Status,
			response.Data,
		)

		// Check if stream is completed
		if response.Status == "completed" {
			fmt.Printf("\n✅ Stream completed successfully\n")
			break
		}

		// Check if context was cancelled
		if ctx.Err() != nil {
			fmt.Printf("\n⚠️  Stream cancelled\n")
			break
		}
	}

	// Check for stream errors
	if err := stream.Err(); err != nil {
		if err == io.EOF {
			// EOF is normal when stream completes
			fmt.Printf("\n📊 Total messages received: %d\n", messageNum)
			return
		}
		if connect.CodeOf(err) == connect.CodeCanceled {
			fmt.Printf("\n⚠️  Stream was cancelled\n")
			fmt.Printf("📊 Total messages received: %d\n", messageNum)
			return
		}
		fmt.Printf("\n❌ Stream error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n📊 Total messages received: %d\n", messageNum)
}
