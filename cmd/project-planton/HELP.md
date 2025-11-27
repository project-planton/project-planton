# Project Planton CLI Help

This guide covers the configuration and deployment component management commands in the Project Planton CLI.

## Table of Contents

- [Configuration Management](#configuration-management)
- [Deployment Components](#deployment-components)
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

3. **Check configuration:**
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
