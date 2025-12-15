# Docker Image Deployment Guide

**Last Updated:** December 12, 2025
**Purpose:** Build and deploy the unified Project Planton Docker image to Docker Hub

---

## Overview

This guide covers building the unified Docker image (MongoDB + Backend + Frontend) and publishing it to Docker Hub so users can install it via the CLI with `planton webapp init`.

**Docker Hub Repository:** `satishlleftbin/project-planton`
**Image Tag Strategy:** `latest` for stable releases, versioned tags for specific releases

---

## Prerequisites

### 1. Docker Hub Account

**Create an account:**
```bash
# Sign up at https://hub.docker.com if you don't have an account
```

**Login to Docker Hub:**
```bash
docker login

# Enter your Docker Hub username and password
# Output: Login Succeeded
```

### 2. Docker Engine

**Verify Docker is running:**
```bash
docker --version
# Output: Docker version 24.0.x or higher

docker info
# Should show server information
```

### 3. Sufficient Disk Space

**Check available space:**
```bash
df -h

# You need at least 10GB free for building
# - Build cache: ~3GB
# - Intermediate layers: ~4GB
# - Final image: ~500MB
```

---

## Building the Unified Docker Image

### Step 1: Navigate to Project Root

```bash
cd /Volumes/Others/Work/crafts/leftbin/planton/project-planton
```

### Step 2: Build the Image

**Build for current architecture (x86_64):**

```bash
# Build the unified image
docker build -f app/Dockerfile.unified -t satishlleftbin/project-planton:latest .

# This will take 5-10 minutes depending on your machine
# You'll see output from all three build stages:
#   [1/3] Building backend...
#   [2/3] Building frontend...
#   [3/3] Creating final image...
```

**Build output you'll see:**
```
[+] Building 450.5s (45/45) FINISHED
 => [backend-builder 1/8] FROM golang:1.24.7-alpine
 => [backend-builder 2/8] RUN apk add --no-cache git make bash curl wget
 => [backend-builder 3/8] RUN curl -fsSL https://github.com/bufbuild/buf/...
 => [backend-builder 4/8] WORKDIR /build
 => [backend-builder 5/8] COPY go.work go.work.sum ./
 => [backend-builder 6/8] COPY go.mod go.sum ./
 => [backend-builder 7/8] COPY pkg/ ./pkg/
 => [backend-builder 8/8] RUN CGO_ENABLED=0 GOOS=linux go build...
 => [frontend-builder 1/6] FROM node:20-alpine
 => [frontend-builder 2/6] RUN corepack enable
 => [frontend-builder 3/6] COPY package.json yarn.lock ./
 => [frontend-builder 4/6] RUN yarn install --frozen-lockfile
 => [frontend-builder 5/6] COPY . .
 => [frontend-builder 6/6] RUN yarn build
 => [stage-2 1/15] FROM ubuntu:22.04
 => [stage-2 2/15] RUN apt-get update && apt-get install -y...
 => [stage-2 3/15] RUN curl -fsSL https://deb.nodesource.com/setup_20.x...
 => [stage-2 4/15] COPY --from=backend-builder /usr/local/go /usr/local/go
 => [stage-2 5/15] RUN wget -q "https://get.pulumi.com/releases/sdk/pulumi...
 => [stage-2 6/15] RUN useradd -u 1001 -r -g 0 -d /home/appuser...
 => [stage-2 7/15] COPY --from=backend-builder /build/app/backend/bin/server...
 => [stage-2 8/15] COPY --from=frontend-builder /app/public /app/frontend/public
 => exporting to image
 => => naming to docker.io/satishlleftbin/project-planton:latest
```

### Step 3: Verify the Build

**Check image size:**
```bash
docker images satishlleftbin/project-planton

# Expected output:
# REPOSITORY                         TAG       IMAGE ID       CREATED          SIZE
# satishlleftbin/project-planton    latest    abc123def456   2 minutes ago    520MB
```

**Expected image size:** ~500-550MB

**Test the image locally:**
```bash
# Create a test container
docker run -d \
  --name test-planton \
  -p 3000:3000 \
  -p 50051:50051 \
  satishlleftbin/project-planton:latest

# Wait 60 seconds for services to start
sleep 60

# Check if services are running
curl http://localhost:3000
curl http://localhost:50051

# Check logs
docker logs test-planton

# Clean up
docker stop test-planton
docker rm test-planton
```

---

## Pushing to Docker Hub

### Step 1: Tag the Image (Optional - for versioning)

**Tag with version number:**
```bash
# Get current version (or manually set)
VERSION="1.0.0"

# Tag with version
docker tag satishlleftbin/project-planton:latest satishlleftbin/project-planton:${VERSION}

# Tag with major version
docker tag satishlleftbin/project-planton:latest satishlleftbin/project-planton:v1

# Verify tags
docker images satishlleftbin/project-planton
```

### Step 2: Push to Docker Hub

**Push the latest tag:**
```bash
docker push satishlleftbin/project-planton:latest

# Output:
# The push refers to repository [docker.io/satishlleftbin/project-planton]
# abc123def456: Pushed
# ...
# latest: digest: sha256:xyz789... size: 4567
```

**Push versioned tags (if created):**
```bash
docker push satishlleftbin/project-planton:${VERSION}
docker push satishlleftbin/project-planton:v1
```

**Expected push time:** 5-10 minutes (depending on internet speed for ~500MB)

### Step 3: Verify on Docker Hub

**Check the repository:**
1. Visit https://hub.docker.com/r/satishlleftbin/project-planton
2. Verify the `latest` tag is present
3. Check the size (~520MB)
4. Verify the last push timestamp

**Or use CLI:**
```bash
# Pull to verify (on a different machine or after removing local image)
docker pull satishlleftbin/project-planton:latest
```

---

## Building Multi-Architecture Images (Optional)

For supporting both Intel (x86_64) and Apple Silicon (ARM64) machines:

### Setup Buildx

```bash
# Create a new builder
docker buildx create --name multiarch --use

# Inspect the builder
docker buildx inspect --bootstrap
```

### Build and Push Multi-Arch

```bash
# Build for both amd64 and arm64
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -f app/Dockerfile.unified \
  -t satishlleftbin/project-planton:latest \
  --push \
  .

# This will:
# - Build for both architectures
# - Create a multi-arch manifest
# - Push directly to Docker Hub
```

**Build time:** 15-20 minutes (building twice - once for each architecture)

**Verification:**
```bash
# Check manifest
docker manifest inspect satishlleftbin/project-planton:latest

# Should show both platforms:
# - linux/amd64
# - linux/arm64
```

---

## Version Management Strategy

### Semantic Versioning

Use semantic versioning: `MAJOR.MINOR.PATCH`

**Example tags:**
```bash
# Latest stable (always updated)
satishlleftbin/project-planton:latest

# Specific version
satishlleftbin/project-planton:1.0.0
satishlleftbin/project-planton:1.0.1
satishlleftbin/project-planton:1.1.0

# Major version (updated for minor/patch)
satishlleftbin/project-planton:v1

# Development/preview
satishlleftbin/project-planton:dev
satishlleftbin/project-planton:preview
```

### Tagging Workflow

**For a new release:**

```bash
VERSION="1.0.1"

# Build
docker build -f app/Dockerfile.unified -t satishlleftbin/project-planton:latest .

# Tag with version
docker tag satishlleftbin/project-planton:latest satishlleftbin/project-planton:${VERSION}
docker tag satishlleftbin/project-planton:latest satishlleftbin/project-planton:v1

# Push all tags
docker push satishlleftbin/project-planton:latest
docker push satishlleftbin/project-planton:${VERSION}
docker push satishlleftbin/project-planton:v1
```

### CLI Configuration

The CLI currently pulls from `latest`:

```go
// cmd/project-planton/root/webapp/init.go
const (
    DockerImageName  = "satishlleftbin/project-planton"
    DockerImageTag   = "latest"
)
```

**To allow version selection in future:**
```go
// Add flag to init command
var imageTag string

func init() {
    InitCmd.Flags().StringVar(&imageTag, "version", "latest", "Docker image version to install")
}
```

---

## Automated CI/CD (Future)

### GitHub Actions Workflow

Create `.github/workflows/docker-build.yml`:

```yaml
name: Build and Push Docker Image

on:
  push:
    branches: [main]
    tags: ['v*']
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v2

      - name: Login to Docker Hub
        uses: docker/login-action@v2
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_TOKEN }}

      - name: Extract metadata
        id: meta
        uses: docker/metadata-action@v4
        with:
          images: satishlleftbin/project-planton
          tags: |
            type=ref,event=branch
            type=semver,pattern={{version}}
            type=semver,pattern={{major}}.{{minor}}
            type=semver,pattern={{major}}

      - name: Build and push
        uses: docker/build-push-action@v4
        with:
          context: .
          file: ./app/Dockerfile.unified
          platforms: linux/amd64,linux/arm64
          push: ${{ github.event_name != 'pull_request' }}
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```

---

## Troubleshooting

### Build Fails

**Issue: "no space left on device"**

```bash
# Clean up Docker
docker system prune -a --volumes

# Remove unused images
docker image prune -a

# Check space
df -h
```

**Issue: "failed to solve: process ... did not complete successfully"**

```bash
# Check which stage failed
# Common issues:
# - Network issues (npm/go dependencies)
# - Missing files in COPY commands
# - Build errors in backend/frontend

# Build with verbose output
docker build -f app/Dockerfile.unified --progress=plain -t test-image .
```

### Push Fails

**Issue: "denied: requested access to the resource is denied"**

```bash
# Re-login to Docker Hub
docker logout
docker login

# Verify credentials
docker login --username satishlleftbin
```

**Issue: "connection timeout"**

```bash
# Check internet connection
ping hub.docker.com

# Try again with retry
docker push satishlleftbin/project-planton:latest
```

### Image Too Large

**Current size: ~520MB**

If you need to reduce size:

```dockerfile
# In Dockerfile.unified, consider:
# 1. Use Alpine instead of Ubuntu (complex due to MongoDB)
# 2. Remove unnecessary packages
# 3. Combine RUN commands to reduce layers
# 4. Use .dockerignore to exclude files
```

Create `.dockerignore`:
```
node_modules
.git
.next
*.log
*.md
.DS_Store
```

---

## Testing After Push

### Test Fresh Install on New Machine

**Simulate fresh install:**
```bash
# Remove local images
docker rmi satishlleftbin/project-planton:latest

# Test CLI installation flow
planton webapp init
planton webapp start

# Verify it pulled from Docker Hub
docker images | grep satishlleftbin

# Test the web app
open http://localhost:3000
```

### Test on Different Architectures

**Intel Mac / Linux:**
```bash
docker pull satishlleftbin/project-planton:latest
# Should pull linux/amd64 variant
```

**Apple Silicon Mac:**
```bash
docker pull satishlleftbin/project-planton:latest
# Should pull linux/arm64 variant (if multi-arch built)
```

---

## Quick Reference

### Full Build and Push Workflow

```bash
# 1. Navigate to project
cd /Volumes/Others/Work/crafts/leftbin/planton/project-planton

# 2. Login to Docker Hub
docker login

# 3. Build the image
docker build -f app/Dockerfile.unified -t satishlleftbin/project-planton:latest .

# 4. Test locally
docker run -d --name test-planton -p 3000:3000 -p 50051:50051 satishlleftbin/project-planton:latest
sleep 60
curl http://localhost:3000
docker stop test-planton && docker rm test-planton

# 5. Tag with version (optional)
VERSION="1.0.0"
docker tag satishlleftbin/project-planton:latest satishlleftbin/project-planton:${VERSION}

# 6. Push to Docker Hub
docker push satishlleftbin/project-planton:latest
docker push satishlleftbin/project-planton:${VERSION}

# 7. Verify on Docker Hub
open https://hub.docker.com/r/satishlleftbin/project-planton

# 8. Test pull
docker rmi satishlleftbin/project-planton:latest
docker pull satishlleftbin/project-planton:latest
```

### Multi-Architecture Build (One Command)

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -f app/Dockerfile.unified \
  -t satishlleftbin/project-planton:latest \
  --push \
  .
```

---

## Next Steps

After pushing the image to Docker Hub:

1. ✅ **Test the CLI installation flow** on a clean machine
2. ✅ **Document the version** in release notes
3. ✅ **Update changelog** with Docker image information
4. ✅ **Announce the release** to users
5. ✅ **Monitor Docker Hub** download statistics

---

## Support

**Docker Hub Repository:** https://hub.docker.com/r/satishlleftbin/project-planton
**Docker Hub Username:** `satishlleftbin`
**Image Name:** `project-planton`
**Default Tag:** `latest`

For issues:
- Check Docker Hub repository page
- Verify image digest matches
- Test pull on different machines
- Check Docker Hub build logs (if using automated builds)

---

**Last Build:** To be documented after first successful push
**Current Version:** To be tagged based on release schedule
**Multi-Arch Support:** TBD (requires buildx setup)

