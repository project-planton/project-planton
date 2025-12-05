package service

import (
	"context"
	"fmt"
	"time"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// StreamingService implements the StreamingService RPC.
type StreamingService struct{}

// NewStreamingService creates a new streaming service instance.
func NewStreamingService() *StreamingService {
	return &StreamingService{}
}

// StreamData streams dummy data to the client.
// This is a server-side streaming RPC that sends data at regular intervals.
func (s *StreamingService) StreamData(
	ctx context.Context,
	req *connect.Request[backendv1.StreamDataRequest],
	stream *connect.ServerStream[backendv1.StreamDataResponse],
) error {
	// Get message count (default: 10 if not specified, 0 means infinite)
	messageCount := int32(10)
	if req.Msg.MessageCount != nil {
		messageCount = *req.Msg.MessageCount
	}

	// Get interval in milliseconds (default: 1000ms)
	intervalMs := int32(1000)
	if req.Msg.IntervalMs != nil {
		intervalMs = *req.Msg.IntervalMs
	}

	// Convert interval to duration
	interval := time.Duration(intervalMs) * time.Millisecond

	// Dummy data templates
	dummyDataTemplates := []string{
		"Processing deployment step 1",
		"Validating cloud resource configuration",
		"Initializing infrastructure components",
		"Deploying resources to cloud provider",
		"Waiting for resource provisioning",
		"Configuring network settings",
		"Setting up security groups",
		"Creating database instances",
		"Configuring load balancers",
		"Finalizing deployment",
	}

	sequence := int32(0)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Stream messages
	for {
		select {
		case <-ctx.Done():
			// Context cancelled, send final message and return
			finalResponse := &backendv1.StreamDataResponse{
				Sequence:  sequence,
				Data:      "Stream cancelled by client",
				Timestamp: timestamppb.Now(),
				Status:    "cancelled",
			}
			_ = stream.Send(finalResponse)
			return ctx.Err()

		case <-ticker.C:
			sequence++

			// Select dummy data based on sequence
			dataIndex := (sequence - 1) % int32(len(dummyDataTemplates))
			data := dummyDataTemplates[dataIndex]

			// Determine status
			status := "streaming"
			if messageCount > 0 && sequence >= messageCount {
				status = "completed"
			}

			// Create response
			response := &backendv1.StreamDataResponse{
				Sequence:  sequence,
				Data:      fmt.Sprintf("%s (message %d)", data, sequence),
				Timestamp: timestamppb.Now(),
				Status:    status,
			}

			// Send response
			if err := stream.Send(response); err != nil {
				return fmt.Errorf("failed to send stream message: %w", err)
			}

			// If we've reached the message count, send completion and return
			if messageCount > 0 && sequence >= messageCount {
				// Send final completion message
				finalResponse := &backendv1.StreamDataResponse{
					Sequence:  sequence,
					Data:      "Stream completed successfully",
					Timestamp: timestamppb.Now(),
					Status:    "completed",
				}
				_ = stream.Send(finalResponse)
				return nil
			}
		}
	}
}
