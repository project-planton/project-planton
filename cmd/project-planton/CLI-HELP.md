# Project Planton CLI Help

This guide covers the configuration, deployment component management, and cloud resource management commands in the Project Planton CLI.

## Table of Contents

- [Configuration Management](#configuration-management)
- [Deployment Components](#deployment-components)
- [Cloud Resources](#cloud-resources)
- [Stack Jobs](#stack-jobs)
- [Common Workflows](#common-workflows)
- [Troubleshooting](#troubleshooting)

---

## Configuration Management

The Project Planton CLI uses a configuration system similar to Git, allowing you to set and manage settings that persist across commands.

### Commands

#### `project-planton config set <key> <value>`

Set a configuration value.

**Available Keys:**

- `backend-url` - URL of the Project Planton backend service

**Example:**

```bash
project-planton config set backend-url http://localhost:50051
project-planton config set backend-url https://api.project-planton.com
```

**Validation:**

- `backend-url` must start with `http://` or `https://`

#### `project-planton config get <key>`

Get a configuration value.

**Example:**

```bash
project-planton config get backend-url
# Output: http://localhost:50051
```

**Error Handling:**

- Returns exit code 1 if the key is not set
- Prints error message for unknown keys

#### `project-planton config list`

List all configuration values.

**Example:**

```bash
project-planton config list
# Output: backend-url=http://localhost:50051

# If no configuration is set:
# Output: No configuration values set
```

### Configuration Storage

- Configuration is stored in `~/.project-planton/config.yaml`
- The configuration directory is created automatically with permissions `0755`
- The configuration file has permissions `0600` (user read/write only)

---

## Deployment Components

### `project-planton list-deployment-components`

List deployment components from the backend service with optional filtering.

#### Basic Usage

```bash
# List all deployment components
project-planton list-deployment-components
```

**Sample Output:**

```
NAME                KIND                PROVIDER    VERSION  ID PREFIX  SERVICE KIND  CREATED
PostgresKubernetes  PostgresKubernetes  kubernetes  v1       k8spg      Yes           2025-11-25
AwsRdsInstance      AwsRdsInstance      aws         v1       rdsins     Yes           2025-11-25
GcpCloudSql         GcpCloudSql         gcp         v1       gcpsql     Yes           2025-11-25

Total: 3 deployment component(s)
```

#### Filtering by Kind

Use the `--kind` flag to filter deployment components by their kind.

```bash
# Filter by specific kind
project-planton list-deployment-components --kind PostgresKubernetes
project-planton list-deployment-components -k AwsRdsInstance
```

**Sample Output:**

```
NAME                KIND                PROVIDER    VERSION  ID PREFIX  SERVICE KIND  CREATED
PostgresKubernetes  PostgresKubernetes  kubernetes  v1       k8spg      Yes           2025-11-25

Total: 1 deployment component(s) (filtered by kind: PostgresKubernetes)
```

#### Flags

- `--kind, -k` - Filter deployment components by kind (optional)
- `--help, -h` - Show help information

#### Output Format

The command displays results in a table with the following columns:

- **NAME** - Display name of the deployment component
- **KIND** - Type/kind of the deployment component
- **PROVIDER** - Cloud provider (aws, gcp, kubernetes, etc.)
- **VERSION** - API version of the component
- **ID PREFIX** - Prefix used for resource IDs
- **SERVICE KIND** - Whether this component can launch services (Yes/No)
- **CREATED** - Creation date in YYYY-MM-DD format

### Prerequisites

Before using the `list-deployment-components` command, you must configure the backend URL:

```bash
project-planton config set backend-url <your-backend-url>
```

---

## Cloud Resources

Cloud resources represent infrastructure resources that can be created and managed through the Project Planton backend service. You can perform complete lifecycle management including create, read, update, delete, and list operations.

### `project-planton cloud-resource:create`

Create a new cloud resource from a YAML manifest file. This command automatically triggers deployment using credentials resolved from the database.

#### Basic Usage

```bash
project-planton cloud-resource:create --arg=path/to/manifest.yaml
```

**Example:**

```bash
# Create a cloud resource from a YAML file
project-planton cloud-resource:create --arg=my-vpc.yaml
```

**Sample Output:**

```
✅ Cloud resource created successfully!

ID: 507f1f77bcf86cd799439011
Name: my-vpc
Kind: CivoVpc
Created At: 2025-11-28 13:14:12
```

#### Automatic Deployment

When you create a cloud resource, the system automatically:

1. Saves the resource to the database
2. Determines the cloud provider from the resource kind (e.g., `GcpCloudSql` → `gcp`)
3. Resolves credentials from the database based on the provider
4. Triggers a Pulumi deployment with the resolved credentials

**Note:** Credentials are automatically resolved from the database based on the cloud provider. Ensure credentials are configured in the database before creating cloud resources (typically done through the backend API or web console).

#### Flags

- `--arg, -a` - Path to the YAML manifest file (required)
- `--help, -h` - Show help information

#### Manifest Requirements

The YAML manifest must contain:

- `kind` - The type of cloud resource (e.g., `CivoVpc`, `AwsRdsInstance`)
- `metadata.name` - A unique name for the resource

**Example Manifest:**

```yaml
kind: CivoVpc
metadata:
  name: my-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
```

#### Error Handling

**Missing manifest:**

```
Error: --arg flag is required. Provide path to YAML manifest file
Usage: project-planton cloud-resource:create --arg=<yaml-file>
```

**Invalid YAML:**

```
Error: Invalid manifest - invalid YAML format: yaml: line 2: found character that cannot start any token
```

**Duplicate resource:**

```
Error: Invalid manifest - cloud resource with name 'my-vpc' already exists
```

**Connection issues:**

```
Error: Cannot connect to backend service at http://localhost:50051. Please check:
  1. The backend service is running
  2. The backend URL is correct
  3. Network connectivity
```

### `project-planton cloud-resource:list`

List all cloud resources from the backend service with optional filtering by kind.

#### Basic Usage

```bash
# List all cloud resources
project-planton cloud-resource:list
```

**Sample Output:**

```
ID                     NAME      KIND            CREATED
507f1f77bcf86cd799439011  my-vpc   CivoVpc        2025-11-28 13:14:12
507f1f77bcf86cd799439012  my-db    AwsRdsInstance  2025-11-28 13:15:00

Total: 2 cloud resource(s)
```

#### Filtering by Kind

Use the `--kind` flag to filter cloud resources by their kind.

```bash
# Filter by specific kind
project-planton cloud-resource:list --kind CivoVpc
project-planton cloud-resource:list -k AwsRdsInstance
```

**Sample Output:**

```
ID                     NAME      KIND     CREATED
507f1f77bcf86cd799439011  my-vpc   CivoVpc  2025-11-28 13:14:12

Total: 1 cloud resource(s) (filtered by kind: CivoVpc)
```

#### Flags

- `--kind, -k` - Filter cloud resources by kind (optional)
- `--help, -h` - Show help information

#### Output Format

The command displays results in a table with the following columns:

- **ID** - Unique identifier (MongoDB ObjectID)
- **NAME** - Resource name (from metadata.name)
- **KIND** - Resource type/kind (e.g., CivoVpc, AwsRdsInstance)
- **CREATED** - Creation timestamp

### `project-planton cloud-resource:get`

Retrieve detailed information about a specific cloud resource by its ID.

#### Basic Usage

```bash
project-planton cloud-resource:get --id=<resource-id>
```

**Example:**

```bash
# Get a cloud resource by ID
project-planton cloud-resource:get --id=507f1f77bcf86cd799439011
```

**Sample Output:**

```
Cloud Resource Details:
======================
ID:         507f1f77bcf86cd799439011
Name:       my-vpc
Kind:       CivoVpc
Created At: 2025-11-28 13:14:12
Updated At: 2025-11-28 14:05:23

Manifest:
----------
kind: CivoVpc
metadata:
  name: my-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
  description: Production VPC
```

#### Flags

- `--id, -i` - Unique identifier of the cloud resource (required)
- `--help, -h` - Show help information

#### Error Handling

**Missing ID:**

```
Error: --id flag is required. Provide the cloud resource ID
Usage: project-planton cloud-resource:get --id=<resource-id>
```

**Resource not found:**

```
Error: Cloud resource with ID '507f1f77bcf86cd799439011' not found
```

**Invalid ID format:**

```
Error: Invalid manifest - invalid ID format
```

### `project-planton cloud-resource:update`

Update an existing cloud resource by providing a new YAML manifest. The manifest's `name` and `kind` must match the existing resource. This command automatically triggers deployment using credentials resolved from the database.

#### Basic Usage

```bash
project-planton cloud-resource:update --id=<resource-id> --arg=<yaml-file>
```

**Example:**

```bash
# Update a cloud resource
project-planton cloud-resource:update --id=507f1f77bcf86cd799439011 --arg=my-vpc-updated.yaml
```

**Sample Output:**

```
✅ Cloud resource updated successfully!

ID: 507f1f77bcf86cd799439011
Name: my-vpc
Kind: CivoVpc
Updated At: 2025-11-28 14:05:23
```

#### Automatic Deployment

When you update a cloud resource, the system automatically:

1. Updates the resource in the database
2. Determines the cloud provider from the resource kind
3. Resolves credentials from the database based on the provider
4. Triggers a Pulumi deployment with the resolved credentials

**Note:** Credentials are automatically resolved from the database based on the cloud provider. Ensure credentials are configured in the database before updating cloud resources (typically done through the backend API or web console).

#### Flags

- `--id, -i` - Unique identifier of the cloud resource (required)
- `--arg, -a` - Path to the YAML manifest file (required)
- `--help, -h` - Show help information

#### Update Validation

**CRITICAL**: The update operation validates that the manifest's `name` and `kind` match the existing resource to prevent accidental data corruption.

**Validation Rules:**

- Manifest `metadata.name` must match existing resource name
- Manifest `kind` must match existing resource kind
- Resource ID and creation timestamp are preserved

**Example Valid Update:**

```yaml
# Existing resource: name=my-vpc, kind=CivoVpc
# This update will succeed
kind: CivoVpc
metadata:
  name: my-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
  description: Updated description
  tags:
    - production
```

#### Error Handling

**Missing arguments:**

```
Error: --id flag is required. Provide the cloud resource ID
Error: --arg flag is required. Provide path to YAML manifest file
Usage: project-planton cloud-resource:update --id=<resource-id> --arg=<yaml-file>
```

**Resource not found:**

```
Error: Cloud resource with ID '507f1f77bcf86cd799439011' not found
```

**Name mismatch:**

```
Error: Invalid manifest - manifest name 'different-name' does not match existing resource name 'my-vpc'
```

**Kind mismatch:**

```
Error: Invalid manifest - manifest kind 'AwsVpc' does not match existing resource kind 'CivoVpc'
```

**Invalid YAML:**

```
Error: Invalid manifest - invalid YAML format: yaml: line 2: found character that cannot start any token
```

### `project-planton cloud-resource:delete`

Delete a cloud resource by its ID. This operation is irreversible.

#### Basic Usage

```bash
project-planton cloud-resource:delete --id=<resource-id>
```

**Example:**

```bash
# Delete a cloud resource
project-planton cloud-resource:delete --id=507f1f77bcf86cd799439011
```

**Sample Output:**

```
✅ Cloud resource 'my-vpc' deleted successfully
```

#### Flags

- `--id, -i` - Unique identifier of the cloud resource (required)
- `--help, -h` - Show help information

#### Error Handling

**Missing ID:**

```
Error: --id flag is required. Provide the cloud resource ID
Usage: project-planton cloud-resource:delete --id=<resource-id>
```

**Resource not found:**

```
Error: Cloud resource with ID '507f1f77bcf86cd799439011' not found
```

**Connection issues:**

```
Error: Cannot connect to backend service at http://localhost:50051. Please check:
  1. The backend service is running
  2. The backend URL is correct
  3. Network connectivity
```

### `project-planton cloud-resource:apply`

Apply a cloud resource from a YAML manifest file. This command performs an **upsert operation**: if a resource with the same `name` and `kind` already exists, it will be updated; otherwise, a new resource will be created. This command automatically triggers deployment using credentials resolved from the database.

#### Basic Usage

```bash
project-planton cloud-resource:apply --arg=path/to/manifest.yaml
```

**Example:**

```bash
# Apply a cloud resource (create or update)
project-planton cloud-resource:apply --arg=my-vpc.yaml
```

#### Automatic Deployment

When you apply a cloud resource, the system automatically:

1. Creates or updates the resource in the database
2. Determines the cloud provider from the resource kind
3. Resolves credentials from the database based on the provider
4. Triggers a Pulumi deployment with the resolved credentials

**Note:** Credentials are automatically resolved from the database based on the cloud provider. Ensure credentials are configured in the database before applying cloud resources (typically done through the backend API or web console).

#### Key Features

- **Idempotent**: Can be run multiple times safely with the same manifest
- **Declarative**: Declare the desired state, let the system figure out create vs update
- **Name + Kind uniqueness**: Resources are identified by the combination of `metadata.name` and `kind`
- **Kubernetes-style**: Follows the familiar `kubectl apply` pattern

#### Sample Output (Create)

When the resource doesn't exist, it will be created:

```
✅ Cloud resource applied successfully!

Action: Created
ID: 507f1f77bcf86cd799439011
Name: my-vpc
Kind: CivoVpc
Created At: 2025-11-28 13:14:12
Updated At: 2025-11-28 13:14:12
```

#### Sample Output (Update)

When the resource already exists (same name and kind), it will be updated:

```
✅ Cloud resource applied successfully!

Action: Updated
ID: 507f1f77bcf86cd799439011
Name: my-vpc
Kind: CivoVpc
Created At: 2025-11-28 13:14:12
Updated At: 2025-11-28 15:30:45
```

#### Flags

- `--arg, -a` - Path to the YAML manifest file (required)
- `--help, -h` - Show help information

#### Manifest Requirements

The YAML manifest must contain:

- `kind` - The type of cloud resource (e.g., `CivoVpc`, `AwsRdsInstance`)
- `metadata.name` - A unique name for the resource

**Example Manifest:**

```yaml
kind: CivoVpc
metadata:
  name: my-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
  description: Production VPC
```

#### How It Works

1. **Extracts** `metadata.name` and `kind` from the manifest
2. **Queries** the backend for an existing resource with matching name and kind
3. **Creates** the resource if it doesn't exist
4. **Updates** the resource if it already exists (preserves ID and creation timestamp)
5. **Returns** the resource with a flag indicating whether it was created or updated

#### Idempotency

The `apply` command is fully idempotent - you can run it multiple times with the same manifest:

```bash
# First run - creates the resource
$ project-planton cloud-resource:apply --arg=my-vpc.yaml
Action: Created

# Second run - updates the resource (even if nothing changed)
$ project-planton cloud-resource:apply --arg=my-vpc.yaml
Action: Updated

# Third run - still works
$ project-planton cloud-resource:apply --arg=my-vpc.yaml
Action: Updated
```

#### Name + Kind Uniqueness

The combination of `metadata.name` and `kind` uniquely identifies a resource. This means you can have resources with the same name but different kinds:

```bash
# Create a CivoVpc named "my-vpc"
$ cat > civo-vpc.yaml <<EOF
kind: CivoVpc
metadata:
  name: my-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
EOF
$ project-planton cloud-resource:apply --arg=civo-vpc.yaml
Action: Created

# Create an AwsVpc with the same name - this is allowed!
$ cat > aws-vpc.yaml <<EOF
kind: AwsVpc
metadata:
  name: my-vpc
spec:
  region: us-east-1
  cidr: 10.1.0.0/16
EOF
$ project-planton cloud-resource:apply --arg=aws-vpc.yaml
Action: Created

# Now you have TWO resources named "my-vpc" with different kinds
```

#### Use Cases

**1. Initial Resource Creation**

```bash
# Create infrastructure from scratch
project-planton cloud-resource:apply --arg=vpc.yaml
project-planton cloud-resource:apply --arg=database.yaml
project-planton cloud-resource:apply --arg=cache.yaml
```

**2. Configuration Updates**

```bash
# Modify vpc.yaml to change CIDR or add tags
# Then apply the changes
project-planton cloud-resource:apply --arg=vpc.yaml
# The resource is updated automatically
```

**3. GitOps Workflows**

```bash
# In CI/CD pipeline - always apply the latest manifest
git pull origin main
project-planton cloud-resource:apply --arg=manifests/production-vpc.yaml
```

**4. Disaster Recovery**

```bash
# Resources deleted accidentally? Just reapply
project-planton cloud-resource:apply --arg=all-resources/*.yaml
# Creates any missing resources, updates existing ones
```

#### Comparison with Other Commands

**Apply vs Create:**

- `create`: Fails if resource already exists (by name only, regardless of kind)
- `apply`: Creates if not exists, updates if exists (by name AND kind)

**Apply vs Update:**

- `update`: Requires resource ID, fails if resource doesn't exist
- `apply`: No ID needed, works whether resource exists or not

**When to use each:**

- Use `apply` for declarative, idempotent workflows (recommended for most cases)
- Use `create` when you want to ensure a resource doesn't already exist
- Use `update` when you have the resource ID and want explicit update semantics

#### Error Handling

**Missing manifest:**

```
Error: --arg flag is required. Provide path to YAML manifest file
Usage: project-planton cloud-resource:apply --arg=<yaml-file>
```

**Invalid YAML:**

```
Error: Invalid manifest - invalid YAML format: yaml: line 2: found character that cannot start any token
```

**Missing required fields:**

```
Error: Invalid manifest - manifest must contain 'kind' field
Error: Invalid manifest - manifest must contain 'metadata.name' field
```

**Connection issues:**

```
Error: Cannot connect to backend service at http://localhost:50051. Please check:
  1. The backend service is running
  2. The backend URL is correct
  3. Network connectivity
```

### Prerequisites

Before using cloud resource commands, you must configure the backend URL:

```bash
project-planton config set backend-url <your-backend-url>
```

---

## Stack Jobs

Stack jobs represent deployment operations for cloud resources. You can stream real-time output from stack jobs to monitor deployment progress.

### `project-planton stack-job:stream-output`

Stream real-time deployment logs from a stack job. Shows stdout and stderr output as it's generated during deployment.

#### Basic Usage

```bash
# Stream output from a stack job
project-planton stack-job:stream-output --id=<stack-job-id>
```

**Sample Output:**

```
🚀 Streaming output for stack job: 69369e4ec78ad326a6e5aa8b

[15:04:05.123] [stdout] [Seq: 1] Updating (example-env.GcpCloudSql.gcp-postgres-example):
[15:04:05.234] [stdout] [Seq: 2]     pulumi:pulumi:Stack project-planton-examples-example-env.GcpCloudSql.gcp-postgres-example  Compiling the program ...
[15:04:06.456] [stdout] [Seq: 3]     pulumi:pulumi:Stack project-planton-examples-example-env.GcpCloudSql.gcp-postgres-example  Finished compiling
[15:04:07.789] [stdout] [Seq: 4] +  gcp:sql:DatabaseInstance gcp-postgres-example creating (0s)
[15:04:10.123] [stdout] [Seq: 5] +  gcp:sql:DatabaseInstance gcp-postgres-example created (3s)

✅ Stream completed successfully

📊 Total messages received: 5 (last sequence: 5)
```

#### Resuming from a Specific Sequence

If you need to resume streaming from a specific point (e.g., after disconnection), use the `--last-sequence` flag:

```bash
# Resume from sequence 100
project-planton stack-job:stream-output --id=<stack-job-id> --last-sequence=100
```

**Sample Output:**

```
🚀 Streaming output for stack job: 69369e4ec78ad326a6e5aa8b
   Resuming from sequence: 100

[15:05:01.234] [stdout] [Seq: 101] Continuing deployment...
[15:05:02.456] [stdout] [Seq: 102] Finalizing resources...
```

#### Flags

- `--id, -i` - Unique identifier of the stack job (required)
- `--last-sequence, -s` - Last sequence number received (for resuming stream from a specific point, default: 0)
- `--help, -h` - Show help information

#### Output Format

Each stream message displays:

- **Timestamp** - Time when the message was generated (HH:MM:SS.mmm format)
- **Stream Type** - `[stdout]` for standard output or `[stderr]` for error output
- **Sequence** - Sequence number of the message (for ordering and resuming)
- **Content** - The actual log line content

#### Graceful Shutdown

The stream command supports graceful shutdown via interrupt signals:

- Press `Ctrl+C` to cancel the stream
- The command will finish processing the current message and exit cleanly
- A summary showing total messages received and last sequence number will be displayed

**Example:**

```bash
$ project-planton stack-job:stream-output --id=69369e4ec78ad326a6e5aa8b
🚀 Streaming output for stack job: 69369e4ec78ad326a6e5aa8b

[15:04:05.123] [stdout] [Seq: 1] Starting deployment...
[15:04:06.456] [stdout] [Seq: 2] Compiling program...
^C

⚠️  Interrupt received, stopping stream...

⚠️  Stream cancelled

📊 Total messages received: 2 (last sequence: 2)
```

#### Error Handling

**Backend URL Not Configured:**

```
Error: backend URL not configured. Run: project-planton config set backend-url <url>
```

**Solution:**

```bash
project-planton config set backend-url http://localhost:50051
```

**Connection Issues:**

```
❌ Error: Cannot connect to backend service at http://localhost:50051. Please check:
  1. The backend service is running
  2. The backend URL is correct
  3. Network connectivity
```

**Solutions:**

1. **Check if backend service is running:**

   ```bash
   # Check if port is accessible
   curl http://localhost:50051
   ```

2. **Verify backend URL configuration:**

   ```bash
   project-planton config get backend-url
   ```

3. **Update backend URL if needed:**
   ```bash
   project-planton config set backend-url <correct-url>
   ```

**Stack Job Not Found:**

```
❌ Error: Stack job with ID 'invalid-id' not found
```

**Solution:**

- Verify the stack job ID is correct
- Use `project-planton cloud-resource:get` to find the associated cloud resource and its stack jobs

**Stream Error:**

```
❌ Stream error: <error details>
```

**Possible Causes:**

- Backend service disconnected during streaming
- Network interruption
- Backend service error

**Solutions:**

1. Check backend service logs
2. Verify network connectivity
3. Retry the stream command
4. Use `--last-sequence` to resume from the last received sequence number

#### Use Cases

**1. Monitoring Active Deployments**

```bash
# Stream output from an in-progress deployment
project-planton stack-job:stream-output --id=<stack-job-id>
```

**2. Reviewing Completed Deployments**

```bash
# Stream all logs from a completed deployment
project-planton stack-job:stream-output --id=<stack-job-id>
```

**3. Resuming After Disconnection**

```bash
# If disconnected, resume from the last sequence number you saw
project-planton stack-job:stream-output --id=<stack-job-id> --last-sequence=150
```

#### Prerequisites

Before using the stack-job:stream-output command, you must configure the backend URL:

```bash
project-planton config set backend-url <your-backend-url>
```

---

## Common Workflows

### Initial Setup

1. **Configure the backend URL:**

   ```bash
   project-planton config set backend-url http://localhost:50051
   ```

2. **Verify configuration:**

   ```bash
   project-planton config get backend-url
   ```

3. **Test connectivity:**
   ```bash
   project-planton list-deployment-components
   ```

### Daily Usage

1. **List all available deployment components:**

   ```bash
   project-planton list-deployment-components
   ```

2. **Find specific component types:**

   ```bash
   # List all Kubernetes components
   project-planton list-deployment-components --kind PostgresKubernetes

   # List all AWS components
   project-planton list-deployment-components --kind AwsRdsInstance
   ```

3. **Manage cloud resources:**

   ```bash
   # Apply a cloud resource (recommended - works for create and update)
   project-planton cloud-resource:apply --arg=my-vpc.yaml

   # Or use explicit create/update commands
   project-planton cloud-resource:create --arg=my-vpc.yaml

   # Get resource details by ID
   project-planton cloud-resource:get --id=507f1f77bcf86cd799439011

   # Update a resource
   project-planton cloud-resource:update --id=507f1f77bcf86cd799439011 --arg=updated.yaml

   # Delete a resource
   project-planton cloud-resource:delete --id=507f1f77bcf86cd799439011
   ```

4. **List cloud resources:**

   ```bash
   # List all cloud resources
   project-planton cloud-resource:list

   # Filter by kind
   project-planton cloud-resource:list --kind CivoVpc
   ```

5. **Check configuration:**

   ```bash
   project-planton config list
   ```

### Environment-Specific Setup

#### Local Development

```bash
project-planton config set backend-url http://localhost:50051
```

#### Staging Environment

```bash
project-planton config set backend-url https://staging-api.project-planton.com
```

#### Production Environment

```bash
project-planton config set backend-url https://api.project-planton.com
```

---

## Troubleshooting

### Backend URL Not Configured

**Error:**

```
Error: backend URL not configured. Run: project-planton config set backend-url <url>
```

**Solution:**

```bash
project-planton config set backend-url http://localhost:50051
```

### Connection Refused

**Error:**

```
Error: Cannot connect to backend service at http://localhost:50051. Please check:
  1. The backend service is running
  2. The backend URL is correct
  3. Network connectivity
```

**Solutions:**

1. **Check if backend service is running:**

   ```bash
   # If using Docker
   docker ps | grep backend

   # Check if port is accessible
   curl http://localhost:50051
   ```

2. **Verify backend URL configuration:**

   ```bash
   project-planton config get backend-url
   ```

3. **Update backend URL if needed:**
   ```bash
   project-planton config set backend-url <correct-url>
   ```

### Invalid Backend URL

**Error:**

```
Error: backend-url must start with http:// or https://
```

**Solution:**

```bash
# Correct format
project-planton config set backend-url http://localhost:50051
project-planton config set backend-url https://api.example.com
```

### No Results Found

**Output:**

```
No deployment components found
# or
No deployment components found with kind 'YourKind'
```

**Possible Causes:**

1. Backend database is empty
2. Wrong kind filter applied
3. Backend service not properly initialized

**Solutions:**

1. **Check without filters:**

   ```bash
   project-planton list-deployment-components
   ```

2. **Verify backend service logs**

3. **Check available kinds by listing all components first**

### Cloud Resource Creation Errors

**Invalid Manifest:**

```
Error: Invalid manifest - invalid YAML format: yaml: line 2: found character that cannot start any token
```

**Solution:**

- Verify the YAML file is valid
- Ensure `kind` and `metadata.name` fields are present
- Check YAML syntax (indentation, quotes, etc.)

**Duplicate Resource Name:**

```
Error: Invalid manifest - cloud resource with name 'my-vpc' already exists
```

**Solution:**

- Use a different name for the resource
- Check existing resources: `project-planton cloud-resource:list`
- Delete the existing resource if needed: `project-planton cloud-resource:delete --id=<id>`

### Cloud Resource Update Errors

**Name Mismatch:**

```
Error: Invalid manifest - manifest name 'different-name' does not match existing resource name 'my-vpc'
```

**Solution:**

- Ensure the manifest `metadata.name` matches the existing resource name
- Get current resource details: `project-planton cloud-resource:get --id=<id>`
- Update the manifest to use the correct name

**Kind Mismatch:**

```
Error: Invalid manifest - manifest kind 'AwsVpc' does not match existing resource kind 'CivoVpc'
```

**Solution:**

- Ensure the manifest `kind` matches the existing resource kind
- If you need to change the kind, delete and recreate the resource
- Get current resource details: `project-planton cloud-resource:get --id=<id>`

**Resource Not Found:**

```
Error: Cloud resource with ID '507f1f77bcf86cd799439011' not found
```

**Solution:**

- Verify the resource ID is correct
- List all resources: `project-planton cloud-resource:list`
- The resource may have been deleted

### Cloud Resource Deletion Errors

**Resource Not Found:**

```
Error: Cloud resource with ID '507f1f77bcf86cd799439011' not found
```

**Solution:**

- Verify the resource ID is correct
- List all resources: `project-planton cloud-resource:list`
- The resource may have already been deleted

**Empty Results:**

```
No cloud resources found
# or
No cloud resources found with kind 'YourKind'
```

**Possible Causes:**

1. No cloud resources have been created yet
2. Wrong kind filter applied
3. Backend database is empty

**Solutions:**

1. **Check without filters:**

   ```bash
   project-planton cloud-resource:list
   ```

2. **Create a test resource:**

   ```bash
   project-planton cloud-resource:create --arg=test-resource.yaml
   ```

3. **Verify backend service is running and initialized**

### Configuration File Issues

**Error:** Permission denied or file access issues

**Solutions:**

1. **Check file permissions:**

   ```bash
   ls -la ~/.project-planton/
   ```

2. **Reset configuration directory:**
   ```bash
   rm -rf ~/.project-planton/
   project-planton config set backend-url <your-url>
   ```

### Network Connectivity

**Error:** Timeout or DNS resolution issues

**Solutions:**

1. **Test basic connectivity:**

   ```bash
   ping <your-backend-host>
   curl <your-backend-url>
   ```

2. **Check firewall/proxy settings**

3. **Try different backend URL (HTTP vs HTTPS)**

---

## Advanced Usage

### Scripting

The CLI commands are designed to be script-friendly:

```bash
#!/bin/bash

# Check if backend is configured
if ! project-planton config get backend-url > /dev/null 2>&1; then
    echo "Backend not configured"
    exit 1
fi

# Get component count
COMPONENT_COUNT=$(project-planton list-deployment-components | grep "Total:" | grep -o '[0-9]\+')
echo "Found $COMPONENT_COUNT deployment components"

# List specific kinds
for kind in PostgresKubernetes AwsRdsInstance GcpCloudSql; do
    echo "=== $kind ==="
    project-planton list-deployment-components --kind "$kind"
done

# Apply cloud resources from directory (recommended - idempotent)
echo "=== Applying resources ==="
for manifest in resources/*.yaml; do
    echo "Applying resource from $manifest"
    project-planton cloud-resource:apply --arg="$manifest"
done

# Or use explicit create for new resources
for manifest in resources/*.yaml; do
    echo "Creating resource from $manifest"
    project-planton cloud-resource:create --arg="$manifest"
done

# List all cloud resources and get details
project-planton cloud-resource:list | grep -v "^Total:" | tail -n +2 | while read -r id name kind created; do
    echo "=== Resource: $name ($kind) ==="
    project-planton cloud-resource:get --id="$id"
    echo ""
done

# Apply updates to resources (no ID needed!)
echo "=== Applying updates ==="
for manifest in updates/*.yaml; do
    name=$(grep "name:" "$manifest" | awk '{print $2}')
    kind=$(grep "kind:" "$manifest" | awk '{print $2}')
    echo "Applying $kind/$name from $manifest"
    project-planton cloud-resource:apply --arg="$manifest"
done

# Cleanup old resources
echo "=== Cleaning up old resources ==="
project-planton cloud-resource:list --kind TestResource | grep -v "^Total:" | tail -n +2 | while read -r id rest; do
    echo "Deleting test resource: $id"
    project-planton cloud-resource:delete --id="$id"
done
```

### Complete Cloud Resource Lifecycle

#### Using Apply (Recommended - Simpler)

```bash
#!/bin/bash

# Complete workflow example using apply command
set -e

# 1. Apply a resource (creates it)
echo "Applying VPC resource..."
cat > temp-vpc.yaml <<EOF
kind: CivoVpc
metadata:
  name: automation-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
EOF

# First apply creates the resource
OUTPUT=$(project-planton cloud-resource:apply --arg=temp-vpc.yaml)
echo "$OUTPUT"
RESOURCE_ID=$(echo "$OUTPUT" | grep "^ID:" | awk '{print $2}')
echo "Resource ID: $RESOURCE_ID"

# 2. Get resource details
echo "Fetching resource details..."
project-planton cloud-resource:get --id="$RESOURCE_ID"

# 3. Modify and apply again (updates it)
echo "Updating resource..."
cat > temp-vpc.yaml <<EOF
kind: CivoVpc
metadata:
  name: automation-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
  description: Updated via automation
  tags:
    - automated
    - production
EOF

# Apply again - automatically updates the resource
project-planton cloud-resource:apply --arg=temp-vpc.yaml

# 4. Verify update
echo "Verifying update..."
project-planton cloud-resource:get --id="$RESOURCE_ID" | grep "description"

# 5. Apply is idempotent - run it again, still works
echo "Applying again (idempotency test)..."
project-planton cloud-resource:apply --arg=temp-vpc.yaml

# 6. Delete resource
echo "Cleaning up..."
project-planton cloud-resource:delete --id="$RESOURCE_ID"

# Cleanup temp file
rm temp-vpc.yaml

echo "Workflow complete!"
```

#### Using Create/Update (Explicit)

```bash
#!/bin/bash

# Complete workflow example using explicit create/update
set -e

# 1. Create a resource
echo "Creating VPC resource..."
cat > temp-vpc.yaml <<EOF
kind: CivoVpc
metadata:
  name: automation-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
EOF

RESOURCE_ID=$(project-planton cloud-resource:create --arg=temp-vpc.yaml | grep "^ID:" | awk '{print $2}')
echo "Created resource with ID: $RESOURCE_ID"

# 2. Get resource details
echo "Fetching resource details..."
project-planton cloud-resource:get --id="$RESOURCE_ID"

# 3. Update the resource
echo "Updating resource..."
cat > temp-vpc.yaml <<EOF
kind: CivoVpc
metadata:
  name: automation-vpc
spec:
  region: NYC1
  cidr: 10.0.0.0/16
  description: Updated via automation
EOF

project-planton cloud-resource:update --id="$RESOURCE_ID" --arg=temp-vpc.yaml

# 4. Verify update
echo "Verifying update..."
project-planton cloud-resource:get --id="$RESOURCE_ID" | grep "description"

# 5. Delete resource
echo "Cleaning up..."
project-planton cloud-resource:delete --id="$RESOURCE_ID"

# Cleanup temp file
rm temp-vpc.yaml

echo "Workflow complete!"
```

### GitOps-Style Infrastructure Management

```bash
#!/bin/bash

# GitOps workflow - sync infrastructure from Git repository
set -e

MANIFEST_DIR="infrastructure/manifests"

echo "=== Syncing infrastructure from Git ==="

# Pull latest changes
git pull origin main

# Apply all resources (creates new, updates existing)
for manifest in "$MANIFEST_DIR"/*.yaml; do
    echo "Applying $(basename $manifest)..."

    # Parse manifest for name and kind
    name=$(grep "name:" "$manifest" | head -1 | awk '{print $2}')
    kind=$(grep "kind:" "$manifest" | head -1 | awk '{print $2}')

    # Apply the resource
    OUTPUT=$(project-planton cloud-resource:apply --arg="$manifest")

    # Check if created or updated
    if echo "$OUTPUT" | grep -q "Action: Created"; then
        echo "✅ Created $kind/$name"
    else
        echo "✅ Updated $kind/$name"
    fi
done

echo "=== Infrastructure sync complete ==="

# List all resources to verify
echo "Current infrastructure state:"
project-planton cloud-resource:list
```

### JSON Output (Future Enhancement)

Currently, the CLI outputs human-readable tables. JSON output support may be added in future versions:

```bash
# Future enhancement
project-planton list-deployment-components --output json
```

---

## Support

For additional help:

- Check the main CLI help: `project-planton --help`
- Command-specific help: `project-planton <command> --help`
- Project documentation: [Project Planton Documentation](https://project-planton.org)
- GitHub Issues: [Report Issues](https://github.com/project-planton/project-planton/issues)
