# Docker Containerization Guide

This guide explains how to containerize and run the Project Planton application using Docker and Docker Compose.

## 🏗️ Architecture

The containerized application consists of:

- **Frontend** (Next.js): Port 3000
- **Backend** (Go): Port 50051
- **MongoDB**: Port 27017

## 🚀 Quick Start

### Prerequisites

- Docker 20.10+
- Docker Compose 2.0+

### Option 1: Using Docker Compose (Recommended)

```bash
# Build and start all services
docker-compose up --build

# Run in background
docker-compose up --build -d

# Stop services
docker-compose down

# Stop and remove volumes
docker-compose down -v
```

### Option 2: Building Individual Services

```bash
# Build backend
docker build -t project-planton-backend ./app/backend

# Build frontend
docker build -t project-planton-frontend ./app/frontend

# Run with external MongoDB
docker run -p 50051:50051 -e MONGODB_URI=mongodb://localhost:27017/project_planton project-planton-backend
docker run -p 3000:3000 project-planton-frontend
```

### Option 3: Building from Project Root

```bash
# Build backend from root (uses go workspace)
docker build -f Dockerfile.backend -t project-planton-backend .
```

## 🔧 Configuration

### Environment Variables

#### Backend
- `MONGODB_URI`: MongoDB connection string (default: `mongodb://localhost:27017`)
- `MONGODB_DATABASE`: Database name (default: `project_planton`)
- `SERVER_PORT`: Server port (default: `50051`)

#### Frontend
- `NODE_ENV`: Environment mode (default: `production`)
- `NEXT_PUBLIC_API_URL`: Backend API URL (default: `http://localhost:50051`)

### Docker Compose Override

Create `docker-compose.override.yml` for local customizations:

```yaml
version: '3.8'
services:
  backend:
    environment:
      - LOG_LEVEL=debug
    ports:
      - "50052:50051"  # Custom port mapping

  frontend:
    environment:
      - NEXT_PUBLIC_API_URL=http://localhost:50052
```

## 📊 Monitoring

### Health Checks

Both services include health check endpoints:

- **Backend**: `http://localhost:50051/health`
- **Frontend**: `http://localhost:3000/api/health`

### Docker Health Status

```bash
# Check container health
docker-compose ps

# View logs
docker-compose logs backend
docker-compose logs frontend
docker-compose logs mongodb
```

## 🔒 Security Features

- **Non-root users**: Both containers run as non-root users
- **Multi-stage builds**: Minimal production images
- **No secrets in images**: Environment variables for configuration
- **Network isolation**: Services communicate via Docker network

## 📁 File Structure

```
project-planton/
├── docker-compose.yml              # Main orchestration file
├── Dockerfile.backend             # Alternative backend build from root
├── .dockerignore                  # Root level ignore file
├── docker/
│   └── mongo-init.js              # MongoDB initialization
└── app/
    ├── backend/
    │   ├── Dockerfile             # Backend container
    │   └── .dockerignore          # Backend ignore file
    └── frontend/
        ├── Dockerfile             # Frontend container
        ├── .dockerignore          # Frontend ignore file
        └── pages/api/health.js    # Health check endpoint
```

## 🐛 Troubleshooting

### Common Issues

1. **Port conflicts**:
   ```bash
   # Check what's using the port
   lsof -ti:3000 | xargs kill -9
   lsof -ti:50051 | xargs kill -9
   ```

2. **Build failures**:
   ```bash
   # Clean build cache
   docker system prune -a
   docker-compose build --no-cache
   ```

3. **Database connection issues**:
   ```bash
   # Check MongoDB logs
   docker-compose logs mongodb

   # Restart MongoDB
   docker-compose restart mongodb
   ```

### Debug Mode

Run with debug logging:

```bash
# Backend debug
docker-compose run -e LOG_LEVEL=debug backend

# Frontend debug
docker-compose run -e NODE_ENV=development frontend
```

## 🚢 Production Deployment

### Build for Production

```bash
# Build optimized images
docker-compose -f docker-compose.yml -f docker-compose.prod.yml build

# Push to registry
docker tag project-planton-backend:latest your-registry/project-planton-backend:v1.0.0
docker tag project-planton-frontend:latest your-registry/project-planton-frontend:v1.0.0

docker push your-registry/project-planton-backend:v1.0.0
docker push your-registry/project-planton-frontend:v1.0.0
```

### Resource Limits

Add resource constraints in production:

```yaml
services:
  backend:
    deploy:
      resources:
        limits:
          cpus: '1.0'
          memory: 512M
        reservations:
          memory: 256M
```

## 📈 Performance Tips

1. **Use .dockerignore** to reduce build context
2. **Multi-stage builds** for smaller images
3. **Layer caching** - copy package files first
4. **Health checks** for proper orchestration
5. **Non-root users** for security

## 🔄 Updates

To update the application:

```bash
# Pull latest changes
git pull origin main

# Rebuild and restart
docker-compose up --build -d
```
