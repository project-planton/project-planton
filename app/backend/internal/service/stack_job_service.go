package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/project-planton/project-planton/app/backend/pkg/models"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/backendconfig"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"

	"connectrpc.com/connect"
	backendv1 "github.com/project-planton/project-planton/app/backend/apis/gen/go/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// StackJobService implements the StackJobService RPC.
type StackJobService struct {
	stackJobRepo      *database.StackJobRepository
	cloudResourceRepo *database.CloudResourceRepository
}

// NewStackJobService creates a new stack job service instance.
func NewStackJobService(
	stackJobRepo *database.StackJobRepository,
	cloudResourceRepo *database.CloudResourceRepository,
) *StackJobService {
	return &StackJobService{
		stackJobRepo:      stackJobRepo,
		cloudResourceRepo: cloudResourceRepo,
	}
}

// DeployCloudResource deploys a cloud resource using Pulumi.
// Fetches the manifest from the cloud resource ID, executes pulumi up, and stores the result in stackjobs table.
func (s *StackJobService) DeployCloudResource(
	ctx context.Context,
	req *connect.Request[backendv1.DeployCloudResourceRequest],
) (*connect.Response[backendv1.DeployCloudResourceResponse], error) {
	cloudResourceID := req.Msg.CloudResourceId
	if cloudResourceID == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("cloud_resource_id cannot be empty"))
	}

	// Fetch cloud resource by ID
	cloudResource, err := s.cloudResourceRepo.FindByID(ctx, cloudResourceID)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to fetch cloud resource: %w", err))
	}

	if cloudResource == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("cloud resource with ID '%s' not found", cloudResourceID))
	}

	// Create stack job with in_progress status
	stackJob := &models.StackJob{
		CloudResourceID: cloudResourceID,
		Status:          "in_progress",
	}

	createdJob, err := s.stackJobRepo.Create(ctx, stackJob)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to create stack job: %w", err))
	}

	// Execute Pulumi deployment asynchronously
	jobID := createdJob.ID.Hex()
	go func() {
		_ = s.deployWithPulumi(context.Background(), jobID, cloudResourceID, cloudResource.Manifest)
	}()

	// Convert to proto
	protoJob := &backendv1.StackJob{
		Id:              createdJob.ID.Hex(),
		CloudResourceId: createdJob.CloudResourceID,
		Status:          createdJob.Status,
		Output:          createdJob.Output,
	}

	if !createdJob.CreatedAt.IsZero() {
		protoJob.CreatedAt = timestamppb.New(createdJob.CreatedAt)
	}
	if !createdJob.UpdatedAt.IsZero() {
		protoJob.UpdatedAt = timestamppb.New(createdJob.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.DeployCloudResourceResponse{
		Job: protoJob,
	}), nil
}

// GetStackJob retrieves a stack job by ID.
func (s *StackJobService) GetStackJob(
	ctx context.Context,
	req *connect.Request[backendv1.GetStackJobRequest],
) (*connect.Response[backendv1.GetStackJobResponse], error) {
	id := req.Msg.Id
	if id == "" {
		return nil, connect.NewError(connect.CodeInvalidArgument, fmt.Errorf("stack job ID cannot be empty"))
	}

	job, err := s.stackJobRepo.FindByID(ctx, id)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to fetch stack job: %w", err))
	}

	if job == nil {
		return nil, connect.NewError(connect.CodeNotFound, fmt.Errorf("stack job with ID '%s' not found", id))
	}

	protoJob := &backendv1.StackJob{
		Id:              job.ID.Hex(),
		CloudResourceId: job.CloudResourceID,
		Status:          job.Status,
		Output:          job.Output,
	}

	if !job.CreatedAt.IsZero() {
		protoJob.CreatedAt = timestamppb.New(job.CreatedAt)
	}
	if !job.UpdatedAt.IsZero() {
		protoJob.UpdatedAt = timestamppb.New(job.UpdatedAt)
	}

	return connect.NewResponse(&backendv1.GetStackJobResponse{
		Job: protoJob,
	}), nil
}

// ListStackJobs lists stack jobs with optional filters and pagination.
func (s *StackJobService) ListStackJobs(
	ctx context.Context,
	req *connect.Request[backendv1.ListStackJobsRequest],
) (*connect.Response[backendv1.ListStackJobsResponse], error) {
	opts := &database.StackJobListOptions{}

	if req.Msg.CloudResourceId != nil {
		id := *req.Msg.CloudResourceId
		opts.CloudResourceID = &id
	}

	if req.Msg.Status != nil {
		s := *req.Msg.Status
		opts.Status = &s
	}

	// Apply pagination with defaults (page=0, size=20) if not provided
	var pageNum int32 = 0
	var pageSize int32 = 20
	if req.Msg.PageInfo != nil {
		pageNum = req.Msg.PageInfo.Num
		pageSize = req.Msg.PageInfo.Size
	}
	opts.PageNum = &pageNum
	opts.PageSize = &pageSize

	// Calculate total pages
	totalCount, err := s.stackJobRepo.Count(ctx, opts)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to count stack jobs: %w", err))
	}

	var totalPages int32
	if pageSize > 0 {
		totalPages = int32((totalCount + int64(pageSize) - 1) / int64(pageSize))
	}

	jobs, err := s.stackJobRepo.List(ctx, opts)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, fmt.Errorf("failed to list stack jobs: %w", err))
	}

	protoJobs := make([]*backendv1.StackJob, 0, len(jobs))
	for _, job := range jobs {
		protoJob := &backendv1.StackJob{
			Id:              job.ID.Hex(),
			CloudResourceId: job.CloudResourceID,
			Status:          job.Status,
			Output:          job.Output,
		}

		if !job.CreatedAt.IsZero() {
			protoJob.CreatedAt = timestamppb.New(job.CreatedAt)
		}
		if !job.UpdatedAt.IsZero() {
			protoJob.UpdatedAt = timestamppb.New(job.UpdatedAt)
		}

		protoJobs = append(protoJobs, protoJob)
	}

	response := &backendv1.ListStackJobsResponse{
		Jobs:       protoJobs,
		TotalPages: totalPages,
	}

	return connect.NewResponse(response), nil
}

// deployWithPulumi executes pulumi up and stores output in stackjobs table
func (s *StackJobService) deployWithPulumi(ctx context.Context, jobID string, cloudResourceID string, manifestYaml string) error {
	// Write manifest to temp file
	tmpFile, err := os.CreateTemp("", "manifest-*.yaml")
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to create temp file: %w", err))
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(manifestYaml); err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to write manifest: %w", err))
	}
	tmpFile.Close()

	// Load manifest
	manifestObject, err := manifest.LoadManifest(tmpFile.Name())
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to load manifest: %w", err))
	}

	// Extract stack FQDN from manifest labels (best effort - continue even if fails)
	var backendConfig *backendconfig.PulumiBackendConfig
	var stackFqdn string
	backendConfig, err = backendconfig.ExtractFromManifest(manifestObject)
	if err != nil {
		// Continue - let Pulumi report the error
	} else if backendConfig != nil && backendConfig.StackFqdn != "" {
		stackFqdn = backendConfig.StackFqdn
	}

	// Extract kind name (best effort - continue even if fails)
	var kindName string
	kindName, err = crkreflect.ExtractKindFromProto(manifestObject)
	if err != nil {
		// Continue - let Pulumi report the error
	}

	// Get Pulumi module directory
	moduleDir := os.Getenv("PULUMI_MODULE_DIR")
	if moduleDir == "" {
		moduleDir = "." // Default to current directory
	}

	// Try to get Pulumi module path (best effort)
	var pulumiModulePath string
	if stackFqdn != "" && kindName != "" {
		pulumiModulePath, err = pulumimodule.GetPath(moduleDir, stackFqdn, kindName)
		if err != nil {
			// Continue - let Pulumi report the error
			pulumiModulePath = moduleDir // Use fallback
		}
	} else {
		pulumiModulePath = moduleDir // Use fallback
	}

	// Build stack input YAML (best effort)
	var stackInputYaml string
	stackInputYaml, err = stackinput.BuildStackInputYaml(manifestObject, stackinputproviderconfig.StackInputProviderConfigOptions{})
	if err != nil {
		// Continue - let Pulumi report the error
		stackInputYaml = "" // Empty fallback
	}

	// ALWAYS execute Pulumi - let it report errors for invalid configurations
	var pulumiArgs []string
	if stackFqdn != "" {
		pulumiArgs = []string{
			"up",
			"--stack", stackFqdn,
			"--yes",
			"--skip-preview",
		}
	} else {
		// Try without stack - Pulumi will report the error
		pulumiArgs = []string{
			"up",
			"--yes",
			"--skip-preview",
		}
	}

	// Execute Pulumi command directly
	timeout := 600 * time.Second // 10 minutes
	cmdCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Build command
	args := []string{"pulumi"}
	args = append(args, pulumiArgs...)

	cmd := exec.CommandContext(cmdCtx, args[0], args[1:]...)

	// Set working directory if provided
	if pulumiModulePath != "" {
		cmd.Dir = pulumiModulePath
	}

	// Set environment variables
	cmd.Env = os.Environ()
	if stackInputYaml != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("STACK_INPUT_YAML=%s", stackInputYaml))
	}

	// Capture output
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute command
	err = cmd.Run()
	exitCode := 0
	success := true

	if err != nil {
		success = false
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			exitCode = -1
		}
	}

	// Prepare response-like structure
	var respStdout, respStderr string
	var respExitCode int32
	if stdout.Len() > 0 {
		respStdout = stdout.String()
	}
	if stderr.Len() > 0 {
		respStderr = stderr.String()
	}
	respExitCode = int32(exitCode)

	// Prepare deployment output as JSON - always use real Pulumi output
	deploymentOutput := map[string]interface{}{
		"timestamp":  time.Now().Format(time.RFC3339),
		"error_type": "pulumi", // This is always a real Pulumi execution
	}

	// Include stack_fqdn if available
	if stackFqdn != "" {
		deploymentOutput["stack_fqdn"] = stackFqdn
	}

	var status string
	if err != nil {
		// Deployment failed with execution error (rare - usually Pulumi returns exit code)
		status = "failed"
		deploymentOutput["status"] = "failed"
		deploymentOutput["exit_code"] = exitCode
		deploymentOutput["stdout"] = respStdout
		deploymentOutput["stderr"] = respStderr
		// Use real Pulumi stderr as error, fallback to exec error
		if respStderr != "" {
			deploymentOutput["error"] = respStderr
		} else if respStdout != "" {
			deploymentOutput["error"] = respStdout
		} else {
			deploymentOutput["error"] = err.Error()
		}
	} else if !success {
		// Deployment failed (non-zero exit code) - use real Pulumi error output
		status = "failed"
		deploymentOutput["status"] = "failed"
		deploymentOutput["stdout"] = respStdout
		deploymentOutput["stderr"] = respStderr
		deploymentOutput["exit_code"] = respExitCode
		// Always use real Pulumi stderr as the error message
		if respStderr != "" {
			deploymentOutput["error"] = respStderr
		} else if respStdout != "" {
			// Fallback to stdout if stderr is empty
			deploymentOutput["error"] = respStdout
		} else {
			deploymentOutput["error"] = fmt.Sprintf("Pulumi command failed with exit code %d", respExitCode)
		}
	} else {
		// Deployment succeeded - use real Pulumi success output
		status = "success"
		deploymentOutput["status"] = "success"
		deploymentOutput["stdout"] = respStdout
		deploymentOutput["stderr"] = respStderr
		deploymentOutput["exit_code"] = respExitCode
		// No error field for success
	}

	// Convert to JSON string and update stack job
	outputJSON, jsonErr := json.Marshal(deploymentOutput)
	if jsonErr != nil {
		// Store error status if JSON marshaling fails
		errorOutput := map[string]interface{}{
			"status":    "failed",
			"error":     fmt.Sprintf("Failed to marshal deployment output: %v", jsonErr),
			"timestamp": time.Now().Format(time.RFC3339),
		}
		errorJSON, _ := json.Marshal(errorOutput)
		updateJob := &models.StackJob{
			Status: "failed",
			Output: string(errorJSON),
		}
		_, _ = s.stackJobRepo.Update(ctx, jobID, updateJob)
		return fmt.Errorf("failed to marshal deployment output: %w", jsonErr)
	}

	// Update stack job with deployment output
	updateJob := &models.StackJob{
		Status: status,
		Output: string(outputJSON),
	}
	_, updateErr := s.stackJobRepo.Update(ctx, jobID, updateJob)
	if updateErr != nil {
		return fmt.Errorf("failed to update stack job: %w", updateErr)
	}

	return nil
}

// updateJobWithError updates a stack job with an error status
func (s *StackJobService) updateJobWithError(ctx context.Context, jobID string, err error) error {
	errorOutput := map[string]interface{}{
		"status":    "failed",
		"error":     err.Error(),
		"timestamp": time.Now().Format(time.RFC3339),
	}
	outputJSON, _ := json.Marshal(errorOutput)
	updateJob := &models.StackJob{
		Status: "failed",
		Output: string(outputJSON),
	}
	_, _ = s.stackJobRepo.Update(ctx, jobID, updateJob)
	return err
}
