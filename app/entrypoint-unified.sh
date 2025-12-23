#!/bin/bash

set -e

echo "=========================================="
echo "🚀 Project Planton Unified Container"
echo "=========================================="
echo ""

# Create necessary directories if they don't exist
echo "📁 Setting up directories..."
mkdir -p /data/db
mkdir -p /var/log/mongodb
mkdir -p /var/log/supervisor
mkdir -p /home/appuser/.pulumi/state
mkdir -p /home/appuser/go/cache
mkdir -p /home/appuser/go/tmp
mkdir -p /home/appuser/.project-planton

# Set proper permissions
echo "🔒 Setting permissions..."
chown -R mongodb:mongodb /data/db
chown -R mongodb:mongodb /var/log/mongodb
chown -R appuser:root /home/appuser
chown -R appuser:root /app/backend
chown -R appuser:root /app/frontend

echo "✅ Setup complete"
echo ""
echo "Starting services:"
echo "  - MongoDB (port 27017)"
echo "  - Backend gRPC Server (port 50051)"
echo "  - Frontend Next.js (port 3000)"
echo ""

# Execute the command passed to the entrypoint
exec "$@"

