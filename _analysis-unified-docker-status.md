# Unified Docker Container Analysis

**Date**: December 15, 2025
**Status**: ❌ BROKEN - Build Failure

## Executive Summary

The unified Docker container approach is currently **non-functional** due to a missing `generate` target in the backend Makefile. The concept and architecture are sound, but the Docker build process fails during the backend build stage.

## What Was Intended (Per Changelog)

Based on `_changelog/2025-12/2025-12-12-091151-single-container-webapp-installation.md`:

### Architecture
```
Single Unified Container (supervisord)
├── MongoDB (priority 1, port 27017)
│   └── localhost-only, no external access
├── Backend (priority 2, port 50051)
│   └── Connect-RPC API + Pulumi integration
└── Frontend (priority 3, port 3000)
    └── Next.js server-side rendering
```

### User Experience
```bash
# Install CLI via Homebrew
brew install project-planton/tap/project-planton

# Initialize web app (one-time setup)
planton webapp init

# Start all services
planton webapp start

# Access at http://localhost:3000
```

### CLI Commands Implemented
✅ `webapp.go` - Main command group
✅ `init.go` - Pull image, create container, configure CLI
✅ `start.go` - Start container and wait for health
✅ `stop.go` - Gracefully stop container
✅ `status.go` - Show container and service health
✅ `logs.go` - View/stream logs with filtering
✅ `restart.go` - Restart all services
✅ `uninstall.go` - Remove container (optionally purge data)

All CLI commands are **implemented and registered** in `cmd/project-planton/root.go` line 45.

---

## Current Status - What's Broken

### 1. Docker Build Failure ❌

**File**: `app/Dockerfile.unified` line 40
**Error**: `make: *** No rule to make target 'generate'.  Stop.`

```dockerfile
# Stage 1: Backend Builder
FROM golang:1.24.7-alpine AS backend-builder
...
WORKDIR /build/app/backend
RUN go mod download
RUN make generate  # ❌ FAILS HERE - Target doesn't exist
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server
```

**Root Cause**: The backend `Makefile` at `app/backend/Makefile` has no `generate` target.

**Current Makefile Contents**:
```makefile
.PHONY: deps
deps:
	@echo "Downloading dependencies..."
	go mod download
	go mod tidy

.PHONY: build
build:
	@echo "Building server..."
	go build -o bin/server ./cmd/server

.PHONY: run
run: build
	@echo "Starting server..."
	./bin/server

# ... other targets (dev, clean, test)
# NO 'generate' TARGET EXISTS
```

### 2. Same Issue in Separate Backend Dockerfile ❌

**File**: `app/backend/Dockerfile` line 35
**Same error**: Would also fail with `make generate`

This suggests the issue was introduced when the backend structure changed, but the Dockerfiles weren't updated.

---

## Why This Wasn't Caught Earlier

### Working Development Setup ✅

The `_cursor/startup.sh` script works because it:
1. Builds CLI at project root: `make local`
2. Builds backend separately: `cd app/backend && go build -o backend-server cmd/server/main.go`
3. Runs backend with environment variables
4. Starts frontend separately with `yarn dev`

**No Docker involved**, so the broken Dockerfiles don't affect local development.

### Code Generation Not Actually Needed ✅

The backend imports generated protobuf code from the **root project**:

```go
// From app/backend/internal/service/stack_job_service.go
import (
    credentialv1 "github.com/project-planton/project-planton/apis/org/project_planton/app/credential/v1"
    stackupdatev1 "github.com/project-planton/project-planton/apis/org/project_planton/app/stackupdate/v1"
    // ... other imports from apis/
)
```

The generated protobuf code already exists in `apis/org/project_planton/` and is committed to the repo:
- `apis/org/project_planton/app/cloudresource/v1/*.pb.go`
- `apis/org/project_planton/app/credential/v1/*.pb.go`
- `apis/org/project_planton/app/stackupdate/v1/*.pb.go`
- `apis/org/project_planton/app/*/v1/*connect/*.connect.go`

**Conclusion**: The backend doesn't need a separate `make generate` step - it just needs access to the root `apis/` directory (which is copied in the Dockerfile).

---

## Analysis of Docker Files

### File Structure
```
app/
├── Dockerfile.unified          # ❌ Broken (line 40: make generate)
├── supervisord.conf            # ✅ OK
├── entrypoint-unified.sh       # ✅ OK
├── backend/
│   ├── Dockerfile              # ❌ Broken (line 35: make generate)
│   ├── Makefile                # ❌ Missing 'generate' target
│   └── cmd/server/main.go      # ✅ OK
└── frontend/
    ├── next.config.js          # ✅ OK (output: 'standalone')
    └── package.json            # ✅ OK
```

### supervisord.conf - Analysis ✅

**File**: `app/supervisord.conf`

```ini
[program:mongodb]
priority=1
user=mongodb
command=/usr/bin/mongod --dbpath /data/db --bind_ip 127.0.0.1

[program:backend]
priority=2
user=appuser
depends_on=mongodb
command=/app/backend/server
environment=MONGODB_URI="mongodb://localhost:27017/project_planton",SERVER_PORT="50051",...

[program:frontend]
priority=3
user=appuser
depends_on=backend
command=node /app/frontend/server.js
environment=NEXT_PUBLIC_API_URL="http://localhost:50051",PORT="3000"
```

**Status**: ✅ Configuration is correct
- Dependencies are properly ordered
- Environment variables are set
- Ports are configured
- Users are appropriate

### entrypoint-unified.sh - Analysis ✅

**File**: `app/entrypoint-unified.sh`

```bash
#!/bin/bash
set -e

# Create directories
mkdir -p /data/db /var/log/mongodb /var/log/supervisor
mkdir -p /home/appuser/.pulumi/state /home/appuser/go/cache

# Set permissions
chown -R mongodb:mongodb /data/db /var/log/mongodb
chown -R appuser:root /home/appuser /app/backend /app/frontend

# Execute supervisord
exec "$@"
```

**Status**: ✅ Script is correct
- Creates necessary directories
- Sets proper permissions
- Executes supervisord

### Frontend Next.js Config - Analysis ✅

**File**: `app/frontend/next.config.js` line 20

```javascript
module.exports = {
  output: 'standalone',  // ✅ Correct for Docker deployment
  compiler: {
    emotion: { /* ... */ }
  }
};
```

**Status**: ✅ Configured correctly for Docker
- `output: 'standalone'` creates self-contained deployment
- Generates `server.js` in `.next/standalone/`
- Expected by Dockerfile.unified line 133-134

---

## MongoDB Retry Logic - Analysis ✅

**File**: `app/backend/internal/database/mongodb.go`

```go
func Connect(ctx context.Context, uri, databaseName string) (*MongoDB, error) {
	maxRetries := 10
	retryDelay := 3 * time.Second

	for attempt := 1; attempt <= maxRetries; attempt++ {
		// Try connection with 10-second timeout
		connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		client, err := mongo.Connect(connectCtx, clientOptions)

		if err != nil {
			// Retry after delay
			time.Sleep(retryDelay)
			continue
		}

		// Verify with ping
		if err := client.Ping(connectCtx, nil); err != nil {
			time.Sleep(retryDelay)
			continue
		}

		break
	}

	return &MongoDB{Client: client, Database: db}, nil
}
```

**Status**: ✅ Already implemented
- 10 retries with 3-second delays (30 seconds total)
- Handles container startup race conditions
- Logs each attempt with structured logging

---

## What Works vs What's Broken

### ✅ Working Components

1. **CLI Command Structure**
   - All 7 webapp commands implemented
   - Registered in root command
   - Help text and flags configured

2. **Container Configuration**
   - supervisord.conf properly configured
   - entrypoint-unified.sh correct
   - Environment variables set
   - Port mappings defined

3. **Backend Code**
   - MongoDB retry logic implemented
   - Imports protobuf code from root apis/
   - Server starts correctly (verified by _cursor/startup.sh)

4. **Frontend Code**
   - Next.js standalone output configured
   - API client implemented
   - SSR working

5. **Development Setup**
   - _cursor/startup.sh works perfectly
   - Separate backend/frontend processes
   - All services start and communicate

### ❌ Broken Components

1. **Dockerfile.unified** - Build fails at `make generate`
2. **app/backend/Dockerfile** - Same issue
3. **Docker Hub Image** - Cannot build, so doesn't exist
4. **CLI Commands** - Implemented but cannot test without image

---

## Root Cause Analysis

### Timeline of Events (Inferred)

1. **Original State**: Backend had `make generate` target (possibly generated local proto stubs)
2. **Refactoring**: Backend was changed to import from root `apis/` directory
3. **Makefile Update**: Backend Makefile simplified, `generate` target removed
4. **Oversight**: Dockerfiles not updated to remove `make generate` step
5. **Development Works**: Local dev doesn't use Docker, so issue not noticed
6. **Docker Builds Fail**: Both Dockerfiles reference non-existent target

### Why Development Script Works

`_cursor/startup.sh` doesn't use Docker or Makefiles:
```bash
# Direct Go build, no Makefile involved
go build -o backend-server cmd/server/main.go

# Direct execution
./backend-server
```

No `make generate` needed because:
- Protobuf code already exists in `apis/` directory
- Backend imports from root via `go.work` workspace
- All dependencies are in `go.mod`

---

## How to Fix - Solution Options

### Option 1: Remove `make generate` Step (RECOMMENDED) ✅

**Rationale**: Backend doesn't need to generate anything - proto code is in root `apis/`

**Changes Required**:
1. Remove line 40 from `app/Dockerfile.unified`
2. Remove line 35 from `app/backend/Dockerfile`

**Pros**:
- Simple fix
- Reflects actual architecture
- Faster builds (no unnecessary step)

**Cons**:
- None

### Option 2: Add Empty `generate` Target (WORKAROUND) ⚠️

**Changes Required**:
Add to `app/backend/Makefile`:
```makefile
.PHONY: generate
generate:
	@echo "No code generation needed - using root apis/"
```

**Pros**:
- Minimal change
- Dockerfiles unchanged

**Cons**:
- Misleading (suggests generation happens)
- Technical debt
- Doesn't reflect reality

### Option 3: Generate Backend-Specific Stubs (OVERENGINEERED) ❌

**Changes Required**:
- Add buf.gen.yaml to app/backend/
- Copy proto files to backend
- Generate separate stubs
- Update imports

**Pros**:
- Backend becomes standalone

**Cons**:
- Violates Go workspace pattern
- Code duplication
- More complex build
- Against project architecture

---

## Recommended Fix

### 1. Update Dockerfiles

**File**: `app/Dockerfile.unified`

Remove lines 40-41:
```dockerfile
# Before (lines 39-43):
RUN go mod download

# Generate proto code
RUN make generate

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server

# After (lines 39-41):
RUN go mod download

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/server ./cmd/server
```

**File**: `app/backend/Dockerfile`

Same change - remove lines 34-35.

### 2. Test Build

```bash
# Build unified Docker image
docker build -f app/Dockerfile.unified -t project-planton:test .

# Verify size (should be ~500MB)
docker images project-planton:test

# Test run
docker run -d -p 3000:3000 -p 50051:50051 \
  -v project-planton-mongodb-data:/data/db \
  -v project-planton-pulumi-state:/home/appuser/.pulumi \
  -v project-planton-go-cache:/home/appuser/go \
  --name test-container \
  project-planton:test

# Check logs
docker logs -f test-container

# Verify services
curl http://localhost:50051/health  # Backend health
curl http://localhost:3000          # Frontend
```

### 3. Verify Services Start

After container starts:
```bash
# Check supervisord status
docker exec test-container supervisorctl status

# Expected output:
# backend                          RUNNING   pid 123, uptime 0:00:05
# frontend                         RUNNING   pid 124, uptime 0:00:03
# mongodb                          RUNNING   pid 122, uptime 0:00:10
```

### 4. Test CLI Commands

```bash
# Clean up test container
docker rm -f test-container

# Test CLI init
planton webapp init

# Test CLI start
planton webapp start

# Test CLI status
planton webapp status

# Test CLI logs
planton webapp logs -n 50

# Test CLI stop
planton webapp stop
```

---

## Additional Considerations

### Go Version in Dockerfile

**Current**: `golang:1.24.7-alpine`
**Issue**: Go 1.24 doesn't exist yet (latest is 1.23)

**Fix Options**:
1. Change to `golang:1.23-alpine` (if Go 1.23 compatible)
2. Change to `golang:1.22-alpine` (safer, well-tested)
3. Keep as is if 1.24 becomes available before deployment

Check `go.mod` in root and backend:
```bash
grep "^go " go.mod app/backend/go.mod
```

### Frontend Build Output

Verify Next.js standalone build works:
```bash
cd app/frontend
yarn build
ls -la .next/standalone/  # Should contain server.js
```

### MongoDB 8.0 Availability

The Dockerfile uses MongoDB 8.0 from official repository:
```dockerfile
https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/8.0 multiverse
```

Verify this works or fall back to 7.0 if unavailable.

---

## Testing Checklist (Once Fixed)

### Build Tests
- [ ] `docker build -f app/Dockerfile.unified` succeeds
- [ ] Image size is reasonable (~500MB)
- [ ] All files are in expected locations
- [ ] Permissions are correct

### Runtime Tests
- [ ] Container starts successfully
- [ ] MongoDB initializes (check logs)
- [ ] Backend connects to MongoDB (after retries)
- [ ] Frontend starts and connects to backend
- [ ] Port 3000 responds to HTTP requests
- [ ] Port 50051 responds to gRPC requests

### CLI Tests
- [ ] `planton webapp init` pulls image and creates container
- [ ] `planton webapp start` starts all services
- [ ] `planton webapp status` shows all services running
- [ ] `planton webapp logs` displays logs from all services
- [ ] `planton webapp restart` restarts container
- [ ] `planton webapp stop` stops gracefully
- [ ] `planton webapp uninstall` removes container

### Data Persistence Tests
- [ ] Create data via API
- [ ] Stop container
- [ ] Start container
- [ ] Data still exists

### Health Check Tests
- [ ] Container health check passes
- [ ] Backend health endpoint works
- [ ] Frontend health endpoint works

---

## Conclusion

### Current State
- **Development Setup**: ✅ Fully functional
- **Unified Docker**: ❌ Build broken (missing Makefile target)
- **CLI Commands**: ✅ Implemented and ready
- **Configuration**: ✅ All configs correct

### Fix Complexity
- **Very Simple**: Remove 2 lines from Dockerfile.unified (line 40-41)
- **Testing Time**: ~1 hour (build + manual tests)
- **Risk**: Very low (removing unnecessary step)

### Next Steps
1. Remove `make generate` from Dockerfiles
2. Build and test locally
3. Push to Docker Hub: `satishlleftbin/project-planton:latest`
4. Test CLI commands end-to-end
5. Update documentation if needed

### Impact Assessment
- No changes to CLI code needed
- No changes to backend/frontend code needed
- Only Docker build process affected
- Development workflow unchanged

