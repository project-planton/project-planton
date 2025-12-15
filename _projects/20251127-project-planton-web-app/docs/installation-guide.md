# Project Planton Web App - Installation Guide

**Last Updated:** December 11, 2025
**Status:** Internal Preview (Not yet public release)

---

## Overview

The Project Planton Web App provides a unified web interface for managing cloud resources and deployments. Everything runs in a single Docker container including MongoDB, backend API, and frontend web UI.

## Architecture

```
┌─────────────────────────────────────────┐
│  Unified Docker Container              │
│                                         │
│  ┌───────────────────────────────────┐ │
│  │  MongoDB (port 27017)             │ │
│  │  - Data persistence               │ │
│  └───────────────────────────────────┘ │
│                                         │
│  ┌───────────────────────────────────┐ │
│  │  Backend API (port 50051)         │ │
│  │  - Connect-RPC server             │ │
│  │  - Pulumi deployments             │ │
│  └───────────────────────────────────┘ │
│                                         │
│  ┌───────────────────────────────────┐ │
│  │  Frontend (port 3000)             │ │
│  │  - Next.js web interface          │ │
│  └───────────────────────────────────┘ │
└─────────────────────────────────────────┘
```

---

## Prerequisites

### Required

- **Docker Engine** - The container runtime
  - macOS: `brew install docker` or [Docker Desktop](https://docker.com/products/docker-desktop)
  - Linux: [Docker Engine Installation Guide](https://docs.docker.com/engine/install/)
  - Windows: [Docker Desktop for Windows](https://docs.docker.com/desktop/install/windows-install/)

### Verification

After installing Docker, verify it's working:

```bash
docker --version
docker info
```

---

## Installation

### Step 1: Install CLI

Install the Project Planton CLI using Homebrew:

```bash
brew install project-planton/tap/project-planton
```

Verify the installation:

```bash
project-planton version
```

### Step 2: Initialize Web App

Initialize the web app (this pulls the Docker image and sets up the container):

```bash
project-planton webapp init
```

This command will:
- ✅ Check Docker availability
- ✅ Pull the unified Docker image (~500MB)
- ✅ Create Docker volumes for data persistence
- ✅ Create the container with proper configuration
- ✅ Configure CLI to use the local backend

**Time required:** 2-5 minutes (depending on internet speed)

### Step 3: Start the Web App

Start the web app services:

```bash
project-planton webapp start
```

This command will:
- ✅ Start the container
- ✅ Wait for all services to be healthy
- ✅ Display access URLs

**Time required:** 30-60 seconds

### Step 4: Access the Web Interface

Once started, access the web app:

- **Frontend (Web UI):** http://localhost:3000
- **Backend API:** http://localhost:50051

---

## Daily Usage

### Starting the Web App

```bash
project-planton webapp start
```

### Stopping the Web App

```bash
project-planton webapp stop
```

Data is preserved when stopped. Start again anytime.

### Checking Status

```bash
project-planton webapp status
```

Shows container and service status.

### Viewing Logs

```bash
# View last 100 lines
project-planton webapp logs

# Follow logs in real-time
project-planton webapp logs -f

# Show more lines
project-planton webapp logs -n 500
```

Press `Ctrl+C` to stop following logs.

### Restarting

```bash
project-planton webapp restart
```

Useful after configuration changes or if services become unresponsive.

---

## Data Persistence

All data is stored in Docker volumes and persists across container restarts:

| Volume | Purpose | Location |
|--------|---------|----------|
| `project-planton-mongodb-data` | MongoDB database | `/data/db` |
| `project-planton-pulumi-state` | Pulumi state files | `/home/appuser/.pulumi` |
| `project-planton-go-cache` | Go build cache | `/home/appuser/go` |

### Backing Up Data

```bash
# Backup MongoDB
docker run --rm -v project-planton-mongodb-data:/data \
  -v $(pwd):/backup ubuntu tar czf /backup/mongodb-backup.tar.gz /data

# Backup Pulumi state
docker run --rm -v project-planton-pulumi-state:/data \
  -v $(pwd):/backup ubuntu tar czf /backup/pulumi-backup.tar.gz /data
```

---

## Uninstallation

### Keep Data (Recommended)

```bash
project-planton webapp uninstall
```

This removes the container but keeps data volumes. You can reinstall later with existing data.

### Complete Removal (Delete Everything)

```bash
project-planton webapp uninstall --purge-data
```

⚠️ **Warning:** This deletes all data including MongoDB database and Pulumi state. Cannot be undone!

---

## Troubleshooting

### Docker Not Found

**Error:**
```
❌ Error: Docker Engine is not installed or not running
```

**Solution:**
1. Install Docker (see Prerequisites section)
2. Verify: `docker info`

### Port Already in Use

**Error:**
```
Error response from daemon: driver failed programming external connectivity...
Bind for 0.0.0.0:3000 failed: port is already allocated
```

**Solution:**
1. Check what's using the port: `lsof -i :3000` or `lsof -i :50051`
2. Stop the conflicting service
3. Or modify ports in container creation (advanced)

### Services Not Starting

**Check logs:**
```bash
project-planton webapp logs -f
```

**Common issues:**
- MongoDB taking longer to initialize (wait 60-90 seconds)
- Insufficient disk space (MongoDB needs ~1GB)
- Docker resource limits (increase Docker memory/CPU allocation)

### Container Already Exists

**Error:**
```
⚠️ Container 'project-planton-webapp' already exists.
```

**Solution:**
```bash
# If you want to start existing container
project-planton webapp start

# If you want to start fresh
project-planton webapp uninstall
project-planton webapp init
```

---

## Configuration

### CLI Configuration

The CLI stores configuration in `~/.project-planton/config.yaml`:

```yaml
backend-url: http://localhost:50051
webapp-container-id: project-planton-webapp
webapp-version: latest
```

### Environment Variables

The container uses these environment variables (configured automatically):

| Variable | Value | Purpose |
|----------|-------|---------|
| `MONGODB_URI` | `mongodb://localhost:27017/project_planton` | MongoDB connection |
| `SERVER_PORT` | `50051` | Backend API port |
| `PORT` | `3000` | Frontend port |
| `PULUMI_HOME` | `/home/appuser/.pulumi` | Pulumi state location |

---

## Next Steps

Once the web app is running:

1. **Explore the Dashboard** - http://localhost:3000
2. **Create Cloud Resources** - Use the web interface to define infrastructure
3. **Deploy Resources** - Deploy to actual cloud providers
4. **Manage Credentials** - Store cloud provider credentials securely

---

## Getting Help

- View all commands: `project-planton webapp --help`
- Check service status: `project-planton webapp status`
- View logs: `project-planton webapp logs -f`

---

## System Requirements

| Resource | Minimum | Recommended |
|----------|---------|-------------|
| RAM | 2GB | 4GB+ |
| Disk Space | 2GB | 5GB+ |
| CPU | 2 cores | 4+ cores |
| Docker Version | 20.10+ | Latest |

---

## Security Notes

- The web app runs on localhost only (not exposed to network)
- MongoDB has no authentication (localhost only)
- Pulumi state is stored locally (not in cloud backend)
- For production use, additional security hardening is required

---

**Status:** This is an internal preview release. Not recommended for production use yet.


