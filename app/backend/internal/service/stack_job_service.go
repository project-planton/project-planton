package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/project-planton/project-planton/app/backend/internal/database"
	"github.com/project-planton/project-planton/app/backend/pkg/models"
	"github.com/project-planton/project-planton/internal/manifest"
	"github.com/project-planton/project-planton/pkg/crkreflect"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/backendconfig"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumistack"
	"github.com/project-planton/project-planton/pkg/iac/stackinput"
	"github.com/project-planton/project-planton/pkg/iac/stackinput/stackinputproviderconfig"

	"connectrpc.com/connect"
	atlasv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/atlas"
	awsv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/aws"
	azurev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/azure"
	cloudflarev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/cloudflare"
	confluentv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/confluent"
	gcpv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/gcp"
	kubernetesv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/kubernetes"
	snowflakev1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/snowflake"
	cloudresourcekind "github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"
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

	// Extract user-provided credentials from request (if provided)
	var userProviderConfig *backendv1.ProviderConfig
	if req.Msg.ProviderConfig != nil {
		userProviderConfig = req.Msg.ProviderConfig
	}

	// Execute Pulumi deployment asynchronously
	jobID := createdJob.ID.Hex()
	go func() {
		_ = s.deployWithPulumi(context.Background(), jobID, cloudResourceID, cloudResource.Manifest, userProviderConfig)
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
// This function performs all required setup steps before executing Pulumi:
// 1. Loads and validates manifest
// 2. Extracts stack FQDN and kind
// 3. Gets Pulumi module path
// 4. Initializes stack if needed
// 5. Updates Pulumi.yaml project name
// 6. Builds stack input YAML (with credentials)
// 7. Executes pulumi up with user-provided credentials
func (s *StackJobService) deployWithPulumi(ctx context.Context, jobID string, cloudResourceID string, manifestYaml string, userProviderConfig *backendv1.ProviderConfig) error {
	// Step 1: Write manifest to temp file
	tmpFile, err := os.CreateTemp("", "manifest-*.yaml")
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to create temp file: %w", err))
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(manifestYaml); err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to write manifest: %w", err))
	}
	if err := tmpFile.Close(); err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to close temp file: %w", err))
	}

	// Step 2: Load manifest and validate
	manifestObject, err := manifest.LoadManifest(tmpFile.Name())
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to load manifest: %w", err))
	}

	// Step 3: Extract stack FQDN from manifest labels (REQUIRED - fail if missing)
	backendConfig, err := backendconfig.ExtractFromManifest(manifestObject)
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to extract backend config from manifest: %w", err))
	}
	if backendConfig == nil || backendConfig.StackFqdn == "" {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("stack FQDN not found in manifest labels. Add 'pulumi.project-planton.org/stack.fqdn' label"))
	}
	stackFqdn := backendConfig.StackFqdn

	// Step 4: Extract kind name (REQUIRED - fail if missing)
	kindName, err := crkreflect.ExtractKindFromProto(manifestObject)
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to extract kind from manifest: %w", err))
	}
	if kindName == "" {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("kind field is required in manifest"))
	}

	// Step 4.1: Get provider from kind enum (for credential validation)
	kindEnum, err := crkreflect.KindByKindName(kindName)
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to get kind enum for '%s': %w", kindName, err))
	}
	provider := crkreflect.GetProvider(kindEnum)

	// Step 5: Get Pulumi module directory
	moduleDir := os.Getenv("PULUMI_MODULE_DIR")
	if moduleDir == "" {
		moduleDir = "." // Default to current directory
	}

	// Step 6: Get Pulumi module path (REQUIRED - fail if missing)
	pulumiModulePath, err := pulumimodule.GetPath(moduleDir, stackFqdn, kindName)
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to get Pulumi module path: %w", err))
	}

	// Step 7: Extract project name from stack FQDN
	pulumiProjectName, err := pulumistack.ExtractProjectName(stackFqdn)
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to extract project name from stack FQDN '%s': %w", stackFqdn, err))
	}

	// Step 8: Update Pulumi.yaml project name (REQUIRED - fail if error)
	if err := pulumistack.UpdateProjectNameInPulumiYaml(pulumiModulePath, pulumiProjectName); err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to update Pulumi.yaml project name: %w", err))
	}

	// Step 9: Initialize stack if it doesn't exist (idempotent)
	// Note: pulumistack.Init writes to os.Stdout/Stderr, but we need to capture output
	// So we'll handle stack initialization manually with output capture
	if err := s.ensureStackInitialized(ctx, jobID, moduleDir, stackFqdn, tmpFile.Name(), pulumiModulePath); err != nil {
		// Check if error is "stack already exists" - that's OK
		if !strings.Contains(err.Error(), "already exists") {
			return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to initialize stack: %w", err))
		}
		// Stack already exists, continue
	}

	// Step 10: Build provider config options from user-provided credentials
	// Credentials must be provided via provider_config in the API request
	if userProviderConfig == nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("provider_config is required in API request for resource '%s'. Please provide credentials via provider_config field", kindName))
	}

	// Convert user-provided credentials to files (same pattern as CLI)
	var awsConfig *awsv1.AwsProviderConfig
	var gcpConfig *gcpv1.GcpProviderConfig
	var azureConfig *azurev1.AzureProviderConfig
	var atlasConfig *atlasv1.AtlasProviderConfig
	var cloudflareConfig *cloudflarev1.CloudflareProviderConfig
	var confluentConfig *confluentv1.ConfluentProviderConfig
	var snowflakeConfig *snowflakev1.SnowflakeProviderConfig
	var kubernetesConfig *kubernetesv1.KubernetesProviderConfig

	// Extract provider configs from oneof
	switch cfg := userProviderConfig.Config.(type) {
	case *backendv1.ProviderConfig_Aws:
		// Convert backend proto to provider proto
		awsConfig = &awsv1.AwsProviderConfig{
			AccountId:       cfg.Aws.AccountId,
			AccessKeyId:     cfg.Aws.AccessKeyId,
			SecretAccessKey: cfg.Aws.SecretAccessKey,
		}
		if cfg.Aws.Region != nil {
			region := *cfg.Aws.Region
			awsConfig.Region = &region
		}
		if cfg.Aws.SessionToken != nil {
			awsConfig.SessionToken = *cfg.Aws.SessionToken
		}
	case *backendv1.ProviderConfig_Gcp:
		gcpConfig = &gcpv1.GcpProviderConfig{
			ServiceAccountKeyBase64: cfg.Gcp.ServiceAccountKeyBase64,
		}
	case *backendv1.ProviderConfig_Azure:
		azureConfig = &azurev1.AzureProviderConfig{
			ClientId:       cfg.Azure.ClientId,
			ClientSecret:   cfg.Azure.ClientSecret,
			TenantId:       cfg.Azure.TenantId,
			SubscriptionId: cfg.Azure.SubscriptionId,
		}
	case *backendv1.ProviderConfig_Atlas:
		atlasConfig = &atlasv1.AtlasProviderConfig{
			PublicKey:  cfg.Atlas.PublicKey,
			PrivateKey: cfg.Atlas.PrivateKey,
		}
	case *backendv1.ProviderConfig_Cloudflare:
		cloudflareConfig = &cloudflarev1.CloudflareProviderConfig{
			AuthScheme: cloudflarev1.CloudflareAuthScheme(cfg.Cloudflare.AuthScheme),
		}
		if cfg.Cloudflare.ApiToken != nil {
			cloudflareConfig.ApiToken = *cfg.Cloudflare.ApiToken
		}
		if cfg.Cloudflare.ApiKey != nil {
			cloudflareConfig.ApiKey = *cfg.Cloudflare.ApiKey
		}
		if cfg.Cloudflare.Email != nil {
			cloudflareConfig.Email = *cfg.Cloudflare.Email
		}
	case *backendv1.ProviderConfig_Confluent:
		confluentConfig = &confluentv1.ConfluentProviderConfig{
			ApiKey:    cfg.Confluent.ApiKey,
			ApiSecret: cfg.Confluent.ApiSecret,
		}
	case *backendv1.ProviderConfig_Snowflake:
		snowflakeConfig = &snowflakev1.SnowflakeProviderConfig{
			Account:  cfg.Snowflake.Account,
			Region:   cfg.Snowflake.Region,
			Username: cfg.Snowflake.Username,
			Password: cfg.Snowflake.Password,
		}
	case *backendv1.ProviderConfig_Kubernetes:
		kubernetesConfig = &kubernetesv1.KubernetesProviderConfig{
			Provider: kubernetesv1.KubernetesProvider(cfg.Kubernetes.Provider),
		}
		if cfg.Kubernetes.GcpGke != nil {
			kubernetesConfig.GcpGke = &kubernetesv1.KubernetesProviderConfigGcpGke{
				ClusterEndpoint:         cfg.Kubernetes.GcpGke.ClusterEndpoint,
				ClusterCaData:           cfg.Kubernetes.GcpGke.ClusterCaData,
				ServiceAccountKeyBase64: cfg.Kubernetes.GcpGke.ServiceAccountKeyBase64,
			}
		}
		if cfg.Kubernetes.DigitalOceanDoks != nil {
			kubernetesConfig.DigitalOceanDoks = &kubernetesv1.KubernetesProviderConfigDigitalOceanDoks{
				KubeConfig: cfg.Kubernetes.DigitalOceanDoks.KubeConfig,
			}
		}
	}

	providerConfigOptions, cleanupProviderConfigs, err := stackinputproviderconfig.BuildProviderConfigOptionsFromUserCredentials(
		awsConfig,
		gcpConfig,
		azureConfig,
		atlasConfig,
		cloudflareConfig,
		confluentConfig,
		snowflakeConfig,
		kubernetesConfig,
	)
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to build provider config from user credentials: %w", err))
	}
	defer cleanupProviderConfigs()

	// Validate that required credentials are provided based on provider enum
	if err := s.validateProviderCredentials(provider, providerConfigOptions, kindName); err != nil {
		return s.updateJobWithError(ctx, jobID, err)
	}

	// Debug: Log which provider configs were found
	if providerConfigOptions.AwsProviderConfig != "" {
		fmt.Printf("DEBUG: AWS provider config file created: %s\n", providerConfigOptions.AwsProviderConfig)
	}
	if providerConfigOptions.GcpProviderConfig != "" {
		fmt.Printf("DEBUG: GCP provider config file created: %s\n", providerConfigOptions.GcpProviderConfig)
	}
	if providerConfigOptions.AzureProviderConfig != "" {
		fmt.Printf("DEBUG: Azure provider config file created: %s\n", providerConfigOptions.AzureProviderConfig)
	}

	// Step 11: Build stack input YAML (REQUIRED - fail if error)
	stackInputYaml, err := stackinput.BuildStackInputYaml(manifestObject, providerConfigOptions)
	if err != nil {
		return s.updateJobWithError(ctx, jobID, fmt.Errorf("failed to build stack input YAML: %w", err))
	}

	// Debug: Check if provider config is in the stack input YAML
	if strings.Contains(stackInputYaml, "provider_config") {
		fmt.Printf("DEBUG: Provider config found in stack input YAML\n")
		// Extract and log provider config details (without sensitive values)
		if strings.Contains(stackInputYaml, "access_key_id") {
			fmt.Printf("DEBUG: AWS access_key_id field found in provider_config\n")
		}
		if strings.Contains(stackInputYaml, "region") {
			// Extract region value for debugging
			regionStart := strings.Index(stackInputYaml, "region:")
			if regionStart > 0 {
				regionLine := stackInputYaml[regionStart : strings.Index(stackInputYaml[regionStart:], "\n")+regionStart]
				fmt.Printf("DEBUG: %s\n", strings.TrimSpace(regionLine))
			}
		}
	} else {
		fmt.Printf("DEBUG: WARNING: Provider config NOT found in stack input YAML\n")
	}

	// Step 12: Cancel any existing locks on the stack (idempotent - safe to run even if no lock exists)
	// This handles cases where a previous deployment was interrupted and left a lock
	if err := s.cancelStackLock(ctx, pulumiModulePath, stackFqdn); err != nil {
		// Log the error but don't fail - the lock might not exist or might be from an active process
		// If there's actually a lock from an active process, pulumi up will fail with a clear error
		fmt.Printf("Warning: Failed to cancel stack lock (this is OK if no lock exists): %v\n", err)
	}

	// Step 13: Execute Pulumi command
	// Increased timeout to 30 minutes to account for plugin downloads and large deployments
	timeout := 1800 * time.Second // 30 minutes
	cmdCtx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	pulumiArgs := []string{
		"up",
		"--stack", stackFqdn,
		"--yes",
		"--skip-preview",
	}

	cmd := exec.CommandContext(cmdCtx, "pulumi", pulumiArgs...)
	cmd.Dir = pulumiModulePath

	// Set environment variables
	cmd.Env = os.Environ()
	// Set STACK_INPUT_YAML (required by Pulumi modules)
	if stackInputYaml != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("STACK_INPUT_YAML=%s", stackInputYaml))
	}
	// Set PROJECT_PLANTON_MANIFEST (some modules may read this)
	cmd.Env = append(cmd.Env, fmt.Sprintf("PROJECT_PLANTON_MANIFEST=%s", manifestYaml))

	// Capture output
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute command
	err = cmd.Run()
	exitCode := 0
	var wasKilled bool
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			// Check if context was cancelled (timeout or cancellation)
			if cmdCtx.Err() == context.DeadlineExceeded {
				exitCode = -1
				wasKilled = true
			} else if cmdCtx.Err() == context.Canceled {
				exitCode = -1
				wasKilled = true
			} else {
				exitCode = -1
			}
		}
	}

	// Step 14: Prepare deployment output
	respStdout := stdout.String()
	respStderr := stderr.String()
	respExitCode := int32(exitCode)

	deploymentOutput := map[string]interface{}{
		"timestamp":  time.Now().Format(time.RFC3339),
		"error_type": "pulumi",
		"stack_fqdn": stackFqdn,
		"stdout":     respStdout,
		"stderr":     respStderr,
		"exit_code":  respExitCode,
	}

	var status string
	if err != nil || exitCode != 0 {
		status = "failed"
		deploymentOutput["status"] = "failed"

		// Check if error is due to stack lock
		combinedOutput := respStderr + respStdout
		isLockError := strings.Contains(combinedOutput, "currently locked") ||
			strings.Contains(combinedOutput, "lock file") ||
			strings.Contains(combinedOutput, "pulumi cancel")

		// Determine error message based on failure type
		var errorMsg string
		if isLockError {
			errorMsg = fmt.Sprintf("Stack is locked by another Pulumi process. This usually happens when:\n"+
				"1. A previous deployment is still running\n"+
				"2. A previous deployment was interrupted and didn't clean up\n"+
				"3. Multiple deployments are trying to run simultaneously\n\n"+
				"Solution: Wait for the other process to finish, or manually cancel the lock:\n"+
				"  pulumi cancel --stack %s --yes\n\n"+
				"Original error:\n%s", stackFqdn, combinedOutput)
		} else if wasKilled {
			if cmdCtx.Err() == context.DeadlineExceeded {
				errorMsg = fmt.Sprintf("Pulumi deployment timed out after %.0f minutes. The process was killed. This may happen if:\n- Plugin downloads take too long\n- The deployment is very large\n- Network issues slow down the process\n\nLast output:\n%s", timeout.Minutes(), respStdout)
			} else {
				errorMsg = fmt.Sprintf("Pulumi deployment was interrupted (process killed). Exit code: %d\n\nLast output:\n%s", respExitCode, respStdout)
			}
		} else if respStderr != "" {
			errorMsg = respStderr
		} else if respStdout != "" {
			// If stdout contains error-like content, use it
			if strings.Contains(respStdout, "error") || strings.Contains(respStdout, "Error") || strings.Contains(respStdout, "failed") {
				errorMsg = respStdout
			} else {
				// Otherwise, prepend a message about the failure
				errorMsg = fmt.Sprintf("Pulumi command failed with exit code %d\n\nOutput:\n%s", respExitCode, respStdout)
			}
		} else if err != nil {
			errorMsg = err.Error()
		} else {
			errorMsg = fmt.Sprintf("Pulumi command failed with exit code %d", respExitCode)
		}
		deploymentOutput["error"] = errorMsg
	} else {
		status = "success"
		deploymentOutput["status"] = "success"
	}

	// Step 13: Convert to JSON and update stack job
	outputJSON, jsonErr := json.Marshal(deploymentOutput)
	if jsonErr != nil {
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

// ensureStackInitialized ensures the Pulumi stack exists, initializing it if needed
func (s *StackJobService) ensureStackInitialized(ctx context.Context, jobID, moduleDir, stackFqdn, manifestPath, pulumiModulePath string) error {
	// Check if stack exists by trying to select it
	checkCmd := exec.CommandContext(ctx, "pulumi", "stack", "select", stackFqdn)
	checkCmd.Dir = pulumiModulePath
	checkCmd.Env = os.Environ()
	var checkStderr bytes.Buffer
	checkCmd.Stderr = &checkStderr

	err := checkCmd.Run()
	if err == nil {
		// Stack exists, nothing to do
		return nil
	}

	// Stack doesn't exist, initialize it
	// Use pulumistack.Init but capture output
	initCmd := exec.CommandContext(ctx, "pulumi", "stack", "init", stackFqdn)
	initCmd.Dir = pulumiModulePath
	initCmd.Env = os.Environ()

	var initStdout, initStderr bytes.Buffer
	initCmd.Stdout = &initStdout
	initCmd.Stderr = &initStderr

	initErr := initCmd.Run()
	if initErr != nil {
		// Check if error is "stack already exists" (race condition)
		output := initStderr.String() + initStdout.String()
		if strings.Contains(output, "already exists") || strings.Contains(output, "stack already exists") {
			return nil // Stack exists, that's OK
		}
		return fmt.Errorf("failed to initialize stack: %w, stderr: %s", initErr, initStderr.String())
	}

	return nil
}

// cancelStackLock cancels any existing locks on the Pulumi stack.
// This is safe to call even if no lock exists - it will simply do nothing.
func (s *StackJobService) cancelStackLock(ctx context.Context, pulumiModulePath, stackFqdn string) error {
	// Build pulumi cancel command
	args := []string{"cancel", "--stack", stackFqdn, "--yes"}
	cancelCmd := exec.CommandContext(ctx, "pulumi", args...)
	cancelCmd.Dir = pulumiModulePath
	cancelCmd.Env = os.Environ()

	// Capture output (but don't fail if cancel fails - lock might not exist)
	var stdout, stderr bytes.Buffer
	cancelCmd.Stdout = &stdout
	cancelCmd.Stderr = &stderr

	err := cancelCmd.Run()
	if err != nil {
		// If cancel fails, it might be because:
		// 1. No lock exists (this is fine)
		// 2. Lock is from an active process (pulumi up will handle this)
		// 3. Some other error (we'll let pulumi up handle it)
		// So we return the error but don't fail the deployment
		return fmt.Errorf("pulumi cancel failed (this is OK if no lock exists): %w, stderr: %s", err, stderr.String())
	}

	return nil
}

// validateProviderCredentials validates that required credentials are provided based on provider enum
func (s *StackJobService) validateProviderCredentials(
	provider cloudresourcekind.CloudResourceProvider,
	providerConfigOptions stackinputproviderconfig.StackInputProviderConfigOptions,
	kindName string,
) error {
	switch provider {
	case cloudresourcekind.CloudResourceProvider_aws:
		if providerConfigOptions.AwsProviderConfig == "" {
			return fmt.Errorf(
				"AWS credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_gcp:
		if providerConfigOptions.GcpProviderConfig == "" {
			return fmt.Errorf(
				"GCP credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_azure:
		if providerConfigOptions.AzureProviderConfig == "" {
			return fmt.Errorf(
				"Azure credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_atlas:
		if providerConfigOptions.AtlasProviderConfig == "" {
			return fmt.Errorf(
				"Atlas credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_cloudflare:
		if providerConfigOptions.CloudflareProviderConfig == "" {
			return fmt.Errorf(
				"Cloudflare credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_confluent:
		if providerConfigOptions.ConfluentProviderConfig == "" {
			return fmt.Errorf(
				"Confluent credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_snowflake:
		if providerConfigOptions.SnowflakeProviderConfig == "" {
			return fmt.Errorf(
				"Snowflake credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_kubernetes:
		if providerConfigOptions.KubernetesProviderConfig == "" {
			return fmt.Errorf(
				"Kubernetes credentials required for resource '%s'. Provide credentials via provider_config in API request",
				kindName,
			)
		}
	case cloudresourcekind.CloudResourceProvider_cloud_resource_provider_unspecified:
		// No credentials needed for unspecified provider
		return nil
	default:
		// For other providers (civo, digitalocean, etc.), credentials are optional
		// They may be provided but are not required
		return nil
	}
	return nil
}

// ifEmpty returns "SET" if value is not empty, otherwise returns defaultValue
func ifEmpty(value, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return "SET"
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
